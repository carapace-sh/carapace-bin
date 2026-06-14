# Renderer — Terminal Rendering Pipeline

In-depth reference for Bubble Tea's renderer: the `cursedRenderer`, screen buffer diffing, ANSI escape emission, and rendering optimizations.

## Renderer Interface

```go
type renderer interface {
    start()
    close() error
    render(View)
    flush(closing bool) error
    reset()
    insertAbove(string) error
    setSyncdUpdates(bool)
    setWidthMethod(ansi.Method)
    resize(int, int)
    setColorProfile(colorprofile.Profile)
    clearScreen()
    writeString(string) (int, error)
    onMouse(MouseMsg) Cmd
}
```

Two implementations:

| Type | File | Purpose |
|------|------|---------|
| `cursedRenderer` | `cursed_renderer.go` | Real renderer — diffs views, emits ANSI |
| `nilRenderer` | `nil_renderer.go` | No-op — used with `WithoutRenderer()` |

## cursedRenderer Struct

```go
type cursedRenderer struct {
    w             io.Writer          // output writer
    buf           bytes.Buffer       // ANSI output buffer
    scr           *uv.TerminalRenderer  // ultraviolet terminal renderer
    cellbuf       uv.ScreenBuffer       // ultraviolet screen buffer
    lastView      *View               // previous view for diffing
    env           []string            // environment variables
    term          string              // TERM value
    width, height int                 // terminal dimensions
    mu            sync.Mutex          // protects view/lastView
    profile       colorprofile.Profile
    logger        uv.Logger
    view          View                // pending view
    hardTabs      bool                // use hard tabs for cursor movement
    backspace     bool                // use backspace for cursor movement
    mapnl         bool                // map NL to CR+LF
    syncdUpdates  bool                // synchronized output (mode 2026)
    starting      bool                // renderer starting flag
}
```

## Rendering Pipeline

### 1. render(View) — Store Pending View

Called from the event loop after each `Update`. Note: `Program.render()` calls `model.View()` first, then passes the result to `renderer.render()`:

```go
func (p *Program) render(model Model) {
    if p.renderer != nil {
        p.renderer.render(model.View())
    }
}

func (s *cursedRenderer) render(v View) {
    s.mu.Lock()
    s.view = v
    s.mu.Unlock()
}
```

Just stores the view — no actual rendering yet. Thread-safe via mutex.

### 2. flush(closing bool) — Render Frame

Called by a ticker (default 60fps). This is where the actual rendering happens:

```
1. Read pending view (mutex-protected)
2. Short-circuit: if viewEquals(lastView, &view) && frame area unchanged → skip
3. Parse Content into uv.StyledString
4. Compute frame area (inline vs alt screen)
5. Draw content into cellbuf via content.Draw()
6. Handle terminal state transitions:
   a. Alt screen enter/exit
   b. Bracketed paste mode enable/disable
   c. Focus reporting enable/disable
   d. Mouse mode changes
   e. Window title changes
   f. Keyboard enhancements (Kitty protocol + modifyOtherKeys)
   g. Cursor color/shape/visibility
   h. Foreground/background color
   i. Progress bar
7. Call scr.Render(cellbuf.RenderBuffer) — compute diff, write ANSI
8. Store lastView = &view
```

### 3. View Diffing (viewEquals)

Compares all View fields to determine if rendering is needed:

```go
func viewEquals(a, b *View) bool {
    if a.Content != b.Content ||
       a.AltScreen != b.AltScreen ||
       a.DisableBracketedPasteMode != b.DisableBracketedPasteMode ||
       a.ReportFocus != b.ReportFocus ||
       a.MouseMode != b.MouseMode ||
       a.WindowTitle != b.WindowTitle ||
       a.ForegroundColor != b.ForegroundColor ||
       a.BackgroundColor != b.BackgroundColor ||
       a.KeyboardEnhancements != b.KeyboardEnhancements {
        return false
    }
    // Cursor: nil check + field comparison (X, Y, Shape, Blink, Color)
    if (a.Cursor == nil) != (b.Cursor == nil) { return false }
    if a.Cursor != nil && b.Cursor != nil {
        if a.Cursor.X != b.Cursor.X || a.Cursor.Y != b.Cursor.Y ||
           a.Cursor.Shape != b.Cursor.Shape || a.Cursor.Blink != b.Cursor.Blink ||
           a.Cursor.Color != b.Cursor.Color { return false }
    }
    // ProgressBar: nil check + value comparison
    if (a.ProgressBar == nil) != (b.ProgressBar == nil) { return false }
    if a.ProgressBar != nil && b.ProgressBar != nil {
        if *a.ProgressBar != *b.ProgressBar { return false }
    }
    return true
}
```

If `viewEquals` returns `true` and the frame area hasn't changed, the frame is **skipped entirely** — no ANSI output. This is the primary performance optimization.

### 4. Cell Buffer Diffing

When content changes, the renderer uses ultraviolet's `ScreenBuffer`:

1. `uv.NewStyledString(view.Content)` — parses ANSI-styled string into cells
2. `content.Draw(cellbuf, cellbuf.Bounds())` — draws cells into the buffer
3. `scr.Render(cellbuf.RenderBuffer)` — compares against previous frame, emits only changed cells

The `TerminalRenderer` from ultraviolet handles:
- Cursor movement optimization (move to changed cell, write character)
- ANSI style changes (only emit style changes when needed)
- Line clearing and scrolling

## Terminal State Transitions

The renderer handles transitions between terminal states by comparing the current view against `lastView`:

### Alt Screen

```go
// Enter alt screen
if view.AltScreen && (lastView == nil || !lastView.AltScreen) {
    buf.WriteString(ansi.SetAltScreenBuffer)
}

// Exit alt screen
if lastView != nil && lastView.AltScreen && !view.AltScreen {
    buf.WriteString(ansi.ResetAltScreenBuffer)
}
```

### Mouse Mode

```go
// Enable mouse
switch {
case view.MouseMode == MouseModeCellMotion && (lastView == nil || lastView.MouseMode != MouseModeCellMotion):
    buf.WriteString(ansi.SetModeMouseButtonEvent + ansi.SetModeMouseExtSgr)
case view.MouseMode == MouseModeAllMotion && (lastView == nil || lastView.MouseMode != MouseModeAllMotion):
    buf.WriteString(ansi.SetModeMouseAnyEvent + ansi.SetModeMouseExtSgr)
}

// Disable mouse
if lastView != nil && lastView.MouseMode != MouseModeNone && view.MouseMode == MouseModeNone {
    buf.WriteString(ansi.ResetModeMouseButtonEvent + ansi.ResetModeMouseAnyEvent + ansi.ResetModeMouseExtSgr)
}
```

### Keyboard Enhancements

```go
flags := keyboardEnhancementsFlags(view.KeyboardEnhancements)
buf.WriteString(ansi.KittyKeyboard(flags))
// On disable:
buf.WriteString(ansi.ResetModifyOtherKeys)
buf.WriteString(ansi.KittyKeyboard(0))
```

### Cursor

```go
// Hide cursor
if view.Cursor == nil {
    buf.WriteString(ansi.ResetModeTextCursorEnable)
}

// Show cursor with style
if view.Cursor != nil {
    buf.WriteString(ansi.SetModeTextCursorEnable)
    buf.WriteString(ansi.SetCursorStyle(shape, blink))
    if view.Cursor.Color != nil {
        buf.WriteString(ansi.SetCursorColor(color))
    }
}
```

### Window Title

```go
if view.WindowTitle != "" && (lastView == nil || lastView.WindowTitle != view.WindowTitle) {
    buf.WriteString(ansi.SetWindowTitle(view.WindowTitle))
}
```

### Foreground/Background Color

```go
if view.ForegroundColor != nil {
    buf.WriteString(ansi.SetForegroundColor(color))
}
if view.BackgroundColor != nil {
    buf.WriteString(ansi.SetBackgroundColor(color))
}
```

### Progress Bar

```go
if view.ProgressBar != nil {
    buf.WriteString(ansi.SetProgressBar(state, value))
}
```

## Synchronized Output (Mode 2026)

When the terminal supports it, the renderer wraps frame updates in synchronized output mode:

```
ANSI: Begin Synchronized Update → render diff → End Synchronized Update
```

This prevents partial frame rendering (tearing/flickering). Support is detected at startup by querying the terminal.

Detection logic (`shouldQuerySynchronizedOutput`):
- Returns true if: no `TERM_PROGRAM` and no `SSH_TTY`, or `WT_SESSION` is set, or `TERM_PROGRAM` is set (except Apple Terminal) and not SSH, or `TERM` contains "ghostty"/"wezterm"/"alacritty"/"kitty"/"rio"

## Unicode Core (Mode 2027)

When supported, the renderer uses grapheme cluster width for cell measurement instead of simple rune width. This correctly handles emoji, combining characters, and wide glyphs.

Detected via mode report response. When `ModeSet`/`ModePermanentlySet` is received for `ModeUnicodeCore`, the width method is set to `ansi.GraphemeWidth`.

## Rendering Optimizations

### Hard Tabs

When the terminal's `TABDLY` termios flag is `TAB0` (no delay), the renderer can use `\t` (tab) characters for horizontal cursor movement. This is faster than CSI cursor sequences for large horizontal jumps.

Detected in `checkOptimizedMovements()` on Unix (termios) and always enabled on Windows.

### Backspace

When the terminal's `BSDLY` termios flag is `BS0` (no delay), the renderer can use `\b` (backspace) for leftward cursor movement. Faster than CSI sequences for small leftward moves.

Detected on Unix (Linux, Darwin, Solaris, AIX). Not available on BSD (only `TABDLY` is checked). Always enabled on Windows.

### NL-to-CR+LF Mapping

When `mapnl` is true, newlines in content are mapped to `\r\n` (CR+LF). This is a workaround for some terminal emulators (originally for Wish SSH sessions) that don't handle bare `\n` correctly.

### FPS Control

The renderer runs on a ticker. Default: 60fps. Max: 120fps. Controlled by `tea.WithFPS(n)`.

Higher FPS = smoother animations but more CPU usage. Lower FPS = less CPU but choppier updates.

### Frame Skip

If `viewEquals` returns true and the frame area hasn't changed, the entire frame is skipped. This means idle programs consume zero CPU for rendering.

## ANSI Escape Sequences Used

| Sequence | Purpose |
|----------|---------|
| `ansi.SetAltScreenBuffer` | Enter alt screen |
| `ansi.ResetAltScreenBuffer` | Exit alt screen |
| `ansi.SetModeBracketedPaste` | Enable bracketed paste |
| `ansi.ResetModeBracketedPaste` | Disable bracketed paste |
| `ansi.SetModeFocusEvent` | Enable focus reporting |
| `ansi.SetModeMouseButtonEvent` | Enable mouse button events |
| `ansi.SetModeMouseAnyEvent` | Enable all mouse events |
| `ansi.SetModeMouseExtSgr` | Enable SGR mouse encoding |
| `ansi.SetWindowTitle` | Set window title |
| `ansi.SetCursorStyle` | Set cursor shape + blink |
| `ansi.SetCursorColor` | Set cursor color |
| `ansi.ResetCursorColor` | Reset cursor color |
| `ansi.SetForegroundColor` | Set terminal foreground |
| `ansi.ResetForegroundColor` | Reset foreground |
| `ansi.SetBackgroundColor` | Set terminal background |
| `ansi.ResetBackgroundColor` | Reset background |
| `ansi.KittyKeyboard(flags)` | Enable Kitty keyboard protocol |
| `ansi.ResetModifyOtherKeys` | Disable modifyOtherKeys |
| `ansi.SetModeTextCursorEnable` | Show cursor |
| `ansi.ResetModeTextCursorEnable` | Hide cursor |
| `ansi.SetProgressBar` | Set progress bar |
| `ansi.SetModeUnicodeCore` | Enable grapheme width mode |
| `ansi.SetModeSynchronizedOutput` | Begin synchronized update |
| `ansi.ResetModeSynchronizedOutput` | End synchronized update |

## Renderer Lifecycle

```
start() → ticker goroutine starts → flush() called on each tick
                                    ↓
                              view changed?
                              ├── yes → diff + emit ANSI
                              └── no  → skip frame
                                    ↓
close()/stop() → flush final frame (if graceful) → restore terminal state
```

On graceful shutdown: renderer flushes the final frame (last view rendered).
On kill: renderer skips the final flush.

For application-level render caching patterns (versioned cache, per-section cache, screen buffer caching), see [patterns.md](patterns.md#performance-patterns).
