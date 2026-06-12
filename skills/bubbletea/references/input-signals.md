# Input Handling and Signals

In-depth reference for Bubble Tea's input system: raw mode, input parsing, keyboard protocols, bracketed paste, signal handling, suspend/resume, and platform-specific code.

## Raw Mode

### initTerminal()

Called during `Run()` startup. If the renderer is not disabled:

1. Calls `initInput()` (platform-specific)
2. Enters raw mode on the TTY
3. Saves previous terminal state for restoration

### Unix (tty_unix.go)

Build tags: `darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || aix || zos`

```go
func (p *Program) initInput() error {
    // Check if input is a terminal
    if !term.IsTerminal(p.ttyInput.Fd()) {
        return nil
    }
    // Enter raw mode, save previous state
    p.previousTtyInputState, err = term.MakeRaw(p.ttyInput.Fd())
    // Set up output TTY if output is a terminal
    // Check for optimized movements (hard tabs, backspace)
    p.checkOptimizedMovements()
    return nil
}
```

### Windows (tty_windows.go)

Build tags: `windows`

```go
func (p *Program) initInput() error {
    // Enter raw mode
    p.previousTtyInputState, err = term.MakeRaw(p.ttyInput.Fd())
    // Enable virtual terminal input mode
    windows.SetConsoleMode(p.ttyInput.Fd(), ENABLE_VIRTUAL_TERMINAL_INPUT)
    // Enable virtual terminal processing on output
    windows.SetConsoleMode(p.ttyOutput.Fd(), ENABLE_VIRTUAL_TERMINAL_PROCESSING | DISABLE_NEWLINE_AUTO_RETURN)
    // Always enable hard tabs + backspace on Windows
    p.useHardTabs = true
    p.useBackspace = true
    return nil
}
```

## Input Reader

### initInputReader

```go
func (p *Program) initInputReader(cancel bool)
```

1. Creates a `uv.NewCancelReader(p.input)` — cancellable reader from ultraviolet
2. Wraps it in `uv.NewTerminalReader(p.cancelReader, term)` — parses terminal input sequences
3. Starts `p.readLoop()` goroutine that calls `p.inputScanner.StreamEvents(p.ctx, p.msgs)`
4. Parsed events are streamed directly into the program's message channel

### CancelReader

The `muesli/cancelreader` (wrapped by ultraviolet) provides a cancellable blocking read. This is essential because:

- Normal `os.Stdin.Read()` blocks indefinitely
- When the program needs to shut down, it must cancel the read
- The cancel reader provides `Cancel()` method that unblocks the read

### waitForReadLoop

After cancelling the reader, `waitForReadLoop()` waits up to 500ms for the read loop goroutine to finish. If it times out, the cancel reader reported success but couldn't actually cancel the read (edge case on some platforms).

## Input Parsing (ultraviolet)

Bubble Tea delegates all input parsing to `charmbracelet/ultraviolet`:

### Keyboard Protocols

| Protocol | Description |
|----------|-------------|
| Kitty Keyboard Protocol | CSI u sequences — full key disambiguation, modifiers, repeat, release |
| modifyOtherKeys (mode 4) | XTerm extension — disambiguates modified keys |
| Legacy escape sequences | VT100/VT220/xterm standard sequences |

The terminal is queried for Kitty keyboard support at startup. If supported, it's enabled via `ansi.KittyKeyboard(flags)`. Otherwise, `modifyOtherKeys2` is tried.

### Key Sequence Processing

Ultraviolet's `TerminalReader` handles:
- Single-byte keys (printable characters, control characters)
- Multi-byte UTF-8 sequences
- CSI sequences (arrow keys, function keys, etc.)
- SSGR mouse encoding
- Kitty keyboard protocol CSI u sequences
- Paste sequences (bracketed paste start/end)
- Focus/blur sequences
- Window size reports
- Clipboard sequences
- Capability reports

All parsed events are `uv.Event` values, which are type-aliased to `tea.Msg`.

## translateInputEvent

Converts ultraviolet events to tea messages via type switch:

| ultraviolet Event | tea Message |
|---|---|
| `ClipboardEvent` | `ClipboardMsg` |
| `ForegroundColorEvent` | `ForegroundColorMsg` |
| `BackgroundColorEvent` | `BackgroundColorMsg` |
| `CursorColorEvent` | `CursorColorMsg` |
| `CursorPositionEvent` | `CursorPositionMsg` |
| `FocusEvent` | `FocusMsg` |
| `BlurEvent` | `BlurMsg` |
| `KeyPressEvent` | `KeyPressMsg` |
| `KeyReleaseEvent` | `KeyReleaseMsg` |
| `MouseClickEvent` | `MouseClickMsg` |
| `MouseMotionEvent` | `MouseMotionMsg` |
| `MouseReleaseEvent` | `MouseReleaseMsg` |
| `MouseWheelEvent` | `MouseWheelMsg` |
| `PasteEvent` | `PasteMsg` |
| `PasteStartEvent` | `PasteStartMsg` |
| `PasteEndEvent` | `PasteEndMsg` |
| `WindowSizeEvent` | `WindowSizeMsg` |
| `CapabilityEvent` | `CapabilityMsg` |
| `TerminalVersionEvent` | `TerminalVersionMsg` |
| `KeyboardEnhancementsEvent` | `KeyboardEnhancementsMsg` |
| `ModeReportEvent` | `ModeReportMsg` |

Unrecognized events pass through as-is.

## Bracketed Paste Mode

Enabled by default. When the user pastes text, the terminal wraps it in special escape sequences:

```
ESC[200~  <pasted content>  ESC[201~
```

This allows the program to distinguish between pasted text and typed text. Without bracketed paste, a paste of "rm -rf /\n" would execute immediately.

### Messages

| Type | When |
|------|------|
| `PasteStartMsg` | Paste begins |
| `PasteMsg` | Paste content |
| `PasteEndMsg` | Paste ends |

### Disable

```go
v.DisableBracketedPasteMode = true
```

### Renderer Handling

The renderer enables/disables bracketed paste mode based on `View.DisableBracketedPasteMode`:

- Enable: `ansi.SetModeBracketedPaste`
- Disable: `ansi.ResetModeBracketedPaste`
- On cleanup: always resets to disabled

## Signal Handling

### Unix (signals_unix.go)

Build tags: same as tty_unix.go

**SIGWINCH** — Terminal resize:

```go
func (p *Program) listenForResize() chan struct{} {
    // Register for SIGWINCH
    // On each SIGWINCH: p.checkResize()
    // checkResize: get terminal size, send WindowSizeMsg
}
```

**SIGINT + SIGTERM** — Interrupt/quit:

```go
func (p *Program) handleSignals() chan struct{} {
    // SIGINT  → send InterruptMsg
    // SIGTERM → send QuitMsg
    // Other signals → send QuitMsg
}
```

**Important**: In raw mode, Ctrl+C is captured as a `KeyPressMsg` (not SIGINT). The SIGINT handler only fires when input is not a TTY (e.g., piped input).

### Windows (signals_windows.go)

Build tags: `windows`

- `listenForResize()` is a no-op (Windows has no SIGWINCH)
- Signal handling uses Windows console APIs

### ignoreSignals

An atomic flag (`uint32`) that temporarily suppresses signal handling:

- Set to `1` during `ReleaseTerminal()` / `RestoreTerminal()` cycle
- Set to `1` during suspend/resume
- Set to `0` when terminal control returns to the program

## Suspend and Resume

### Suspend Flow

When `SuspendMsg` is received (or user presses Ctrl+Z):

1. `releaseTerminal(true)` — restore terminal, cancel input reader, reset renderer
2. `suspendProcess()` — **Unix only**:
   - Register for `SIGCONT`
   - Send `SIGTSTP` to process group 0 (`syscall.Kill(0, syscall.SIGTSTP)`)
   - **Block** until `SIGCONT` is received
3. `RestoreTerminal()` — re-enter raw mode, restart input reader and renderer
4. Send `ResumeMsg{}` to the program

### Platform Support

| Platform | `suspendSupported` | Behavior |
|----------|-------------------|----------|
| Unix | `true` | Full suspend/resume via SIGTSTP/SIGCONT |
| Windows | `false` | `suspendProcess()` is a no-op; `SuspendMsg` is ignored |

### ResumeMsg

Sent after the program resumes from suspend. Use it to refresh state that may have changed while suspended:

```go
case tea.ResumeMsg:
    // Re-check terminal size, refresh display, etc.
```

## Terminal Capability Detection

### Color Profile

```go
colorprofile.Detect(p.output, p.environ)
```

Checks `COLORTERM`, `TERM`, and terminal queries to determine:
- `ANSI` (4-bit, 16 colors)
- `ANSI256` (8-bit, 256 colors)
- `TrueColor` (24-bit, 16M colors)

Sent as `ColorProfileMsg` on startup and when capabilities are reported.

### Synchronized Output (Mode 2026)

`shouldQuerySynchronizedOutput()` checks environment variables:

Returns true if:
- No `TERM_PROGRAM` and no `SSH_TTY`, or
- `WT_SESSION` is set (Windows Terminal), or
- `TERM_PROGRAM` is set (except Apple Terminal) and not SSH, or
- `TERM` contains "ghostty", "wezterm", "alacritty", "kitty", or "rio"

If true, queries the terminal for mode 2026 support. If the terminal responds with `ModeSet`, synchronized updates are enabled on the renderer.

### Unicode Core (Mode 2027)

Queried alongside mode 2026. When supported, the renderer uses grapheme cluster width for cell measurement. This correctly handles:
- Emoji (including multi-codepoint sequences like family emoji)
- Combining characters (accents, diacritics)
- Wide CJK characters
- Zero-width joiner sequences

### Capability Reports

When the terminal reports capability "RGB" or "Tc", the color profile is upgraded to `TrueColor` and a `ColorProfileMsg` is sent.

## Platform-Specific Code

### Build Tags and Files

| File | Build Tags | Purpose |
|------|-----------|---------|
| `tty_unix.go` | `darwin \|\| dragonfly \|\| freebsd \|\| linux \|\| netbsd \|\| openbsd \|\| solaris \|\| aix \|\| zos` | Raw mode via termios, suspend via SIGTSTP/SIGCONT |
| `tty_windows.go` | `windows` | Raw mode + VT console modes, no suspend |
| `signals_unix.go` | same as tty_unix.go | SIGWINCH resize, SIGINT/SIGTERM handling |
| `signals_windows.go` | `windows` | No SIGWINCH, Windows signal handling |
| `termios_unix.go` | `darwin \|\| linux \|\| solaris \|\| aix` | Check `TABDLY` + `BSDLY` for optimizations |
| `termios_bsd.go` | `dragonfly \|\| freebsd` | Check only `TABDLY` |
| `termios_windows.go` | `windows` | Always enable hard tabs + backspace |
| `termios_other.go` | `!windows && !darwin && !dragonfly && !freebsd && !linux && !solaris && !aix` | No optimization detection |

### Key Differences

| Feature | Unix | Windows | Other |
|---------|------|---------|-------|
| Raw mode | `term.MakeRaw` via termios | `term.MakeRaw` + VT console modes | Varies |
| Suspend | SIGTSTP/SIGCONT | Not supported | Not supported |
| Resize | SIGWINCH | No SIGWINCH | Varies |
| Hard tabs | Detected from termios `TABDLY` | Always enabled | Not detected |
| Backspace | Detected from termios `BSDLY` | Always enabled | Not detected |
| Signal handling | SIGINT → InterruptMsg, SIGTERM → QuitMsg | Windows console APIs | Varies |

## restoreTerminalState

Called during shutdown to restore the terminal to its original state:

1. Flush output buffer
2. Restore input terminal state (via `term.Restore`)
3. Restore output terminal state (via `term.Restore`)
4. Reset bracketed paste mode
5. Reset mouse mode
6. Reset cursor visibility
7. Reset keyboard enhancements
8. Exit alt screen (if active)
9. Reset foreground/background colors
10. Reset window title
