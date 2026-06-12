# v1 to v2 Migration Guide

In-depth reference for migrating from Bubble Tea v1 to v2: breaking changes, new patterns, and a complete mapping of v1 → v2 equivalents.

## Module Path Change

| v1 | v2 |
|----|-----|
| `github.com/charmbracelet/bubbletea` | `charm.land/bubbletea/v2` |

The v2 module uses a custom domain (`charm.land`) instead of the GitHub path. This is a hard breaking change — you must update all import paths.

```go
// v1
import tea "github.com/charmbracelet/bubbletea"

// v2
import tea "charm.land/bubbletea/v2"
```

**Companion ecosystem:**

| Library | v1 | v2 |
|---------|-----|-----|
| lipgloss | `github.com/charmbracelet/lipgloss` | `charm.land/lipgloss/v2` |
| bubbles | `github.com/charmbracelet/bubbles` | `charm.land/bubbles/v2` |
| glamour | `github.com/charmbracelet/glamour` | `charm.land/glamour/v2` |
| huh | `github.com/charmbracelet/huh` | `charm.land/huh/v2` |

## View() Return Type

The biggest architectural change: `View()` now returns `tea.View` instead of `string`.

```go
// v1
func (m model) View() string {
    return "Hello, World"
}

// v2
func (m model) View() tea.View {
    return tea.NewView("Hello, World")
}
```

This enables the **declarative view** pattern — terminal state (alt screen, mouse mode, cursor, etc.) is now set on the View struct instead of through commands.

## Key Messages

### KeyMsg Type

| v1 | v2 |
|----|-----|
| `tea.KeyMsg` (struct) | `tea.KeyMsg` (interface) |
| `tea.KeyMsg.Type` (KeyType) | `tea.KeyPressMsg` / `tea.KeyReleaseMsg` |
| `msg.Type` | `msg.Code` (rune) |
| `msg.Runes` ([]rune) | `msg.Text` (string) |
| `msg.Alt` (bool) | `msg.Mod.Contains(tea.ModAlt)` |

### Key Matching

```go
// v1
switch msg := msg.(type) {
case tea.KeyMsg:
    switch msg.Type {
    case tea.KeyEnter:
    case tea.KeyCtrlC:
    default:
        switch msg.String() {
        case "j":
        case "k":
        }
    }
}

// v2
switch msg := msg.(type) {
case tea.KeyPressMsg:
    switch msg.String() {
    case "enter":
    case "ctrl+c":
    case "j":
    case "k":
    }
}
```

### Space Key

```go
// v1
case " ":  // matches space

// v2
case "space":  // space is now "space" in String()
```

### Modifier Keys

```go
// v1
if msg.Alt { /* alt was held */ }

// v2
if msg.Mod.Contains(tea.ModAlt) { /* alt was held */ }
if msg.Mod.Contains(tea.ModCtrl) { /* ctrl was held */ }
if msg.Mod.Contains(tea.ModShift) { /* shift was held */ }
```

## Mouse Messages

| v1 | v2 |
|----|-----|
| `tea.MouseMsg` (struct) | `tea.MouseMsg` (interface) |
| `msg.Type` (MouseType) | Type switch: `MouseClickMsg`, `MouseReleaseMsg`, `MouseWheelMsg`, `MouseMotionMsg` |
| `msg.X`, `msg.Y` | `msg.Mouse().X`, `msg.Mouse().Y` |
| `msg.Button` (MouseButton) | `msg.Mouse().Button` |
| `msg.Alt` | `msg.Mouse().Mod.Contains(tea.ModAlt)` |
| `tea.MouseLeft` | `tea.MouseLeft` (same name) |
| `tea.MouseButtonLeft` | `tea.MouseLeft` (renamed) |

```go
// v1
case tea.MouseMsg:
    if msg.Type == tea.MouseLeft {
        x, y := msg.X, msg.Y
    }

// v2
case tea.MouseClickMsg:
    m := msg.Mouse()
    if m.Button == tea.MouseLeft {
        x, y := m.X, m.Y
    }
```

## Terminal State — Commands → View Fields

The biggest conceptual change: terminal state is now **declarative** on the View struct instead of **imperative** via commands.

### Alt Screen

```go
// v1 — command
func (m model) Init() tea.Cmd {
    return tea.EnterAltScreen
}
// or program option
p := tea.NewProgram(model, tea.WithAltScreen())

// v2 — View field
func (m model) View() tea.View {
    v := tea.NewView(m.content)
    v.AltScreen = true
    return v
}
```

### Mouse Mode

```go
// v1 — command
func (m model) Init() tea.Cmd {
    return tea.EnableMouseCellMotion
}
// or program option
p := tea.NewProgram(model, tea.WithMouseCellMotion())

// v2 — View field
func (m model) View() tea.View {
    v := tea.NewView(m.content)
    v.MouseMode = tea.MouseModeCellMotion
    return v
}
```

### Cursor

```go
// v1 — commands
tea.HideCursor
tea.ShowCursor

// v2 — View field
v.Cursor = nil              // hide
v.Cursor = tea.NewCursor(5, 3)  // show at position
```

### Window Title

```go
// v1 — command
tea.SetWindowTitle("My App")

// v2 — View field
v.WindowTitle = "My App"
```

### Focus Reporting

```go
// v1 — command
tea.ReportFocus()

// v2 — View field
v.ReportFocus = true
```

### Bracketed Paste

```go
// v1 — program option
p := tea.NewProgram(model, tea.WithoutBracketedPaste)

// v2 — View field
v.DisableBracketedPasteMode = true
```

## Program API Changes

### Run Method

```go
// v1
p := tea.NewProgram(model)
err := p.Start()
// or
model, err := p.StartReturningModel()

// v2 — single method
model, err := p.Run()
```

### Removed Program Options

| v1 Option | v2 Equivalent |
|-----------|---------------|
| `tea.WithAltScreen()` | `View.AltScreen = true` |
| `tea.WithMouseCellMotion()` | `View.MouseMode = tea.MouseModeCellMotion` |
| `tea.WithMouseAllMotion()` | `View.MouseMode = tea.MouseModeAllMotion` |
| `tea.WithReportFocus()` | `View.ReportFocus = true` |
| `tea.WithoutBracketedPaste()` | `View.DisableBracketedPasteMode = true` |
| `tea.WithInputTTY()` | Removed — v2 always opens TTY automatically |
| `tea.WithANSICompressor()` | Removed — new renderer handles optimization automatically |

### New Program Options

| v2 Option | Description |
|-----------|-------------|
| `tea.WithFPS(int)` | Control max FPS (default 60, max 120) |
| `tea.WithColorProfile(Profile)` | Force color profile |
| `tea.WithWindowSize(w, h)` | Set initial window size |

## Command Changes

### Renamed Commands

| v1 | v2 |
|----|-----|
| `tea.Sequentially(cmds...)` | `tea.Sequence(cmds...)` |
| `tea.WindowSize()` (returns `Cmd`) | `tea.RequestWindowSize()` (returns `Msg` directly, not a `Cmd`) |

### Removed Commands

| v1 Command | v2 Equivalent |
|------------|---------------|
| `tea.EnterAltScreen` | `View.AltScreen = true` |
| `tea.ExitAltScreen` | `View.AltScreen = false` |
| `tea.EnableMouseCellMotion` | `View.MouseMode = tea.MouseModeCellMotion` |
| `tea.EnableMouseAllMotion` | `View.MouseMode = tea.MouseModeAllMotion` |
| `tea.DisableMouse` | `View.MouseMode = tea.MouseModeNone` |
| `tea.HideCursor` | `View.Cursor = nil` |
| `tea.ShowCursor` | `View.Cursor = tea.NewCursor(x, y)` |
| `tea.SetWindowTitle(title)` | `View.WindowTitle = title` |
| `tea.ReportFocus()` | `View.ReportFocus = true` |

### New Commands

| v2 Command | Description |
|------------|-------------|
| `tea.Suspend()` | Suspend the program (Ctrl+Z) |
| `tea.Interrupt()` | Interrupt the program (Ctrl+C) |
| `tea.ExecProcess(cmd, callback)` | Run subprocess with terminal release/restore |

## New Message Types

| v2 Message | Description |
|------------|-------------|
| `tea.SuspendMsg` | Program should suspend |
| `tea.ResumeMsg` | Program resumed from suspend |
| `tea.InterruptMsg` | Program interrupted (Ctrl+C) |
| `tea.KeyReleaseMsg` | Key released (requires KeyboardEnhancements) |
| `tea.MouseClickMsg` | Mouse click event |
| `tea.MouseReleaseMsg` | Mouse release event |
| `tea.MouseWheelMsg` | Mouse wheel event |
| `tea.MouseMotionMsg` | Mouse motion event |
| `tea.PasteMsg` | Paste content |
| `tea.PasteStartMsg` | Paste start |
| `tea.PasteEndMsg` | Paste end |
| `tea.ClipboardMsg` | Clipboard content |
| `tea.CapabilityMsg` | Terminal capability report |
| `tea.ModeReportMsg` | Terminal mode report |
| `tea.KeyboardEnhancementsMsg` | Keyboard enhancement support |
| `tea.TerminalVersionMsg` | Terminal version |
| `tea.ForegroundColorMsg` | Terminal foreground color |
| `tea.BackgroundColorMsg` | Terminal background color |
| `tea.CursorColorMsg` | Terminal cursor color |
| `tea.CursorPositionMsg` | Terminal cursor position |

## New View Features

| Feature | How to Use |
|---------|-----------|
| `View.Cursor` | `v.Cursor = tea.NewCursor(x, y)` — cursor position, shape, blink, color |
| `View.ProgressBar` | `v.ProgressBar = tea.NewProgressBar(state, value)` |
| `View.BackgroundColor` | `v.BackgroundColor = color.RGBA{...}` |
| `View.ForegroundColor` | `v.ForegroundColor = color.RGBA{...}` |
| `View.KeyboardEnhancements` | `v.KeyboardEnhancements = tea.KeyboardEnhancements{ReportEventTypes: true}` |
| `View.OnMouse` | `v.OnMouse = func(msg tea.MouseMsg) tea.Cmd { ... }` |

## Complete v1 → v2 Quick Reference

| v1 | v2 |
|----|-----|
| `View() string` | `View() tea.View` |
| `tea.KeyMsg` (struct) | `tea.KeyPressMsg` for presses, `tea.KeyMsg` (interface) for both |
| `msg.Type` | `msg.Code` (rune) |
| `msg.Runes` | `msg.Text` (string, not `[]rune`) |
| `msg.Alt` | `msg.Mod.Contains(tea.ModAlt)` |
| `tea.MouseMsg` (struct) | `tea.MouseMsg` (interface) — call `.Mouse()` |
| `tea.MouseButtonLeft` | `tea.MouseLeft` |
| `tea.WithAltScreen()` | `view.AltScreen = true` |
| `tea.WithMouseCellMotion()` | `view.MouseMode = tea.MouseModeCellMotion` |
| `tea.WithMouseAllMotion()` | `view.MouseMode = tea.MouseModeAllMotion` |
| `tea.EnterAltScreen` | `view.AltScreen = true` |
| `tea.ExitAltScreen` | `view.AltScreen = false` |
| `tea.HideCursor` | `view.Cursor = nil` |
| `tea.ShowCursor` | `view.Cursor = &tea.Cursor{...}` |
| `tea.SetWindowTitle("...")` | `view.WindowTitle = "..."` |
| `tea.ReportFocus()` | `view.ReportFocus = true` |
| `p.Start()` | `p.Run()` |
| `p.StartReturningModel()` | `p.Run()` |
| `tea.Sequentially(...)` | `tea.Sequence(...)` |
| `tea.WindowSize()` (returns `Cmd`) | `tea.RequestWindowSize()` (returns `Msg`) |
| `tea.WithInputTTY()` | Removed — automatic |
| `tea.WithANSICompressor()` | Removed — automatic |
| `case " ":` | `case "space":` |

## Migration Checklist

1. Update import paths: `github.com/charmbracelet/bubbletea` → `charm.land/bubbletea/v2`
2. Update companion imports: lipgloss, bubbles, glamour, huh
3. Change `View() string` → `View() tea.View`
4. Wrap View return values with `tea.NewView(content)`
5. Move terminal state from commands/options to View fields
6. Replace `tea.KeyMsg` with `tea.KeyPressMsg` / `tea.KeyReleaseMsg`
7. Update key field access: `msg.Type` → `msg.Code`, `msg.Runes` → `msg.Text`, `msg.Alt` → `msg.Mod.Contains()`
8. Update mouse handling: `tea.MouseMsg` struct → interface with `.Mouse()` method
9. Replace `p.Start()` / `p.StartReturningModel()` with `p.Run()`
10. Replace `tea.Sequentially` with `tea.Sequence`
11. Replace `tea.WindowSize()` with `tea.RequestWindowSize()`
12. Update space key matching: `case " ":` → `case "space":`
13. Remove `tea.WithInputTTY()` — v2 handles TTY automatically
14. Remove `tea.WithANSICompressor()` — v2 renderer handles optimization
