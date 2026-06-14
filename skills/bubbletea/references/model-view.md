# Model, Update, and View — The Elm Architecture in Go

In-depth reference for Bubble Tea's core architecture: the `Model` interface, the `Update`/`View` cycle, the declarative `View` struct, and how views are composed and diffed.

## The Model Interface

Every Bubble Tea program implements a single interface:

```go
type Model interface {
    Init() Cmd
    Update(Msg) (Model, Cmd)
    View() View
}
```

- **`Init()`** — Called once when the program starts. Returns an initial `Cmd` (or `nil` for no-op). Use it to kick off async work (HTTP requests, timers, etc.).
- **`Update(msg)`** — Called for every message. Receives the current model and a message, returns a new model and an optional `Cmd`. This is a **pure state transition** — no side effects except via returned commands.
- **`View()`** — Called after every `Update`. Returns a `View` struct that declaratively describes what the terminal should look like. **Must be side-effect-free** — no I/O, no mutations.

### The Cycle

```
Init() → Cmd → Msg → Update(Msg) → (Model, Cmd) → View() → render
                ↑                                              |
                └──────────── Cmd() → Msg ────────────────────┘
```

1. `Init()` returns a `Cmd`
2. The `Cmd` executes asynchronously and produces a `Msg`
3. `Update(Msg)` processes the message, returns new `Model` + optional `Cmd`
4. `View()` renders the model into a `View` struct
5. The renderer diffs the new `View` against the previous one and emits ANSI updates
6. Any returned `Cmd` executes and the cycle repeats

## The View Struct (v2)

The `View` struct is the central declarative type in v2. It replaces v1's approach of using commands and program options to control terminal state.

```go
type View struct {
    Content                   string
    AltScreen                 bool
    ReportFocus               bool
    DisableBracketedPasteMode bool
    MouseMode                 MouseMode
    KeyboardEnhancements      KeyboardEnhancements
    OnMouse                   func(MouseMsg) Cmd
    Cursor                    *Cursor
    BackgroundColor           color.Color
    ForegroundColor           color.Color
    WindowTitle               string
    ProgressBar               *ProgressBar
}
```

### Constructor

```go
func NewView(s string) View {
    var view View
    view.SetContent(s)
    return view
}
```

### Content

`Content` is a `string` containing ANSI-styled text (typically produced by lipgloss). The renderer parses it into a cell buffer for diffing against the previous frame.

```go
func (m model) View() tea.View {
    v := tea.NewView(lipgloss.NewStyle().Bold(true).Render("Hello"))
    return v
}
```

### AltScreen

Set `AltScreen = true` to enter the alternate screen buffer. The renderer automatically enters on `true` and exits on `false` (or when the program quits).

```go
v.AltScreen = true
```

In v1, this was `tea.EnterAltScreen` / `tea.ExitAltScreen` commands or `tea.WithAltScreen()` program option.

### Cursor

The `Cursor` field controls cursor visibility, position, shape, and color. When `nil`, the cursor is hidden.

```go
type Cursor struct {
    Position
    Color color.Color
    Shape CursorShape
    Blink bool
}

type Position struct{ X, Y int }
```

**CursorShape constants:**

| Constant | Shape |
|----------|-------|
| `CursorBlock` | Block cursor (default) |
| `CursorUnderline` | Underline cursor |
| `CursorBar` | Vertical bar cursor |

**Usage:**

```go
// Show blinking block cursor at (5, 3)
v.Cursor = tea.NewCursor(5, 3)

// Show non-blinking underline cursor
v.Cursor = &tea.Cursor{
    Position: tea.Position{X: 5, Y: 3},
    Shape:    tea.CursorUnderline,
    Blink:    false,
}

// Hide cursor
v.Cursor = nil
```

In v1, cursor was implicit from string position, and `tea.HideCursor`/`tea.ShowCursor` were commands.

### MouseMode

Controls what mouse events the terminal reports:

| Constant | Events |
|----------|--------|
| `MouseModeNone` | No mouse events |
| `MouseModeCellMotion` | Click, release, wheel, drag (button held) |
| `MouseModeAllMotion` | All events including motion with no button |

```go
v.MouseMode = tea.MouseModeCellMotion
```

In v1, this was `tea.EnableMouseCellMotion` / `tea.EnableMouseAllMotion` commands or `tea.WithMouseCellMotion()` program option.

### KeyboardEnhancements

Requests advanced keyboard features from the terminal (Kitty keyboard protocol + modifyOtherKeys2):

```go
type KeyboardEnhancements struct {
    ReportEventTypes           bool  // key repeat + release events
    ReportAlternateKeys        bool  // alternate key values
    ReportAllKeysAsEscapeCodes bool  // all keys as escape codes
    ReportAssociatedText       bool  // text associated with key events
}
```

When `ReportEventTypes` is enabled, you receive `KeyReleaseMsg` and `KeyPressMsg` with `IsRepeat` field set.

```go
v.KeyboardEnhancements = tea.KeyboardEnhancements{
    ReportEventTypes:    true,
    ReportAlternateKeys: true,
}
```

### OnMouse

A callback that intercepts mouse events before they reach `Update`. Return a `Cmd` to produce a message, or `nil` for no-op.

```go
v.OnMouse = func(msg tea.MouseMsg) tea.Cmd {
    m := msg.Mouse()
    if m.Y == 0 && m.Button == tea.MouseLeft {
        return func() tea.Msg { return headerClickMsg{x: m.X} }
    }
    return nil
}
```

### ReportFocus

When `true`, the program receives `FocusMsg` and `BlurMsg` when the terminal gains or loses focus.

```go
v.ReportFocus = true
```

### WindowTitle

Sets the terminal window/tab title:

```go
v.WindowTitle = "My App"
```

### BackgroundColor / ForegroundColor

Set terminal default colors:

```go
v.BackgroundColor = color.RGBA{R: 30, G: 30, B: 46, A: 255}
v.ForegroundColor = color.RGBA{R: 205, G: 214, B: 244, A: 255}
```

### ProgressBar

Controls the terminal's native progress bar (supported terminals only):

```go
type ProgressBar struct {
    State ProgressBarState
    Value int  // 0-100
}
```

**ProgressBarState constants:**

| Constant | Meaning |
|----------|---------|
| `ProgressBarNone` | No progress bar |
| `ProgressBarDefault` | Normal progress |
| `ProgressBarError` | Error state |
| `ProgressBarIndeterminate` | Indeterminate/spinning |
| `ProgressBarWarning` | Warning state |

```go
v.ProgressBar = tea.NewProgressBar(tea.ProgressBarDefault, 75)
```

Bracketed paste is enabled by default. Set to `true` to disable:

```go
v.DisableBracketedPasteMode = true
```

## View Diffing

The renderer compares the new `View` against `lastView` using `viewEquals()`, which checks all fields:

- `Content` (string equality)
- `AltScreen`, `ReportFocus`, `DisableBracketedPasteMode` (bool equality)
- `MouseMode`, `KeyboardEnhancements` (struct equality)
- `WindowTitle` (string equality)
- `ForegroundColor`, `BackgroundColor` (nil + deep equality)
- `Cursor` (nil check + deep equality)
- `ProgressBar` (nil check + deep equality)

If nothing changed and the frame area is the same, the renderer **skips the frame entirely** — no ANSI output. This is the primary performance optimization.

When content changes, the renderer:
1. Parses `Content` into a `uv.StyledString`
2. Draws it into a `uv.ScreenBuffer` (cell buffer)
3. Compares the cell buffer against the previous frame
4. Emits only the changed cells as ANSI cursor-movement + character sequences

## View Composition Patterns

For the Draw/View split pattern (rendering to a ScreenBuffer with pixel-level layout control), see [patterns.md](patterns.md#pointer-based-model-with-drawview-split).

### Simple View

```go
func (m model) View() tea.View {
    return tea.NewView(m.text)
}
```

### Styled View with Lipgloss

```go
func (m model) View() tea.View {
    style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
    return tea.NewView(style.Render("Hello, World"))
}
```

### Multi-Section Layout

```go
func (m model) View() tea.View {
    content := lipgloss.JoinVertical(lipgloss.Left,
        m.headerView(),
        m.bodyView(),
        m.footerView(),
    )
    v := tea.NewView(content)
    v.AltScreen = true
    v.Cursor = tea.NewCursor(m.cursorX, m.cursorY)
    v.MouseMode = tea.MouseModeCellMotion
    return v
}
```

### Conditional Alt Screen

```go
func (m model) View() tea.View {
    v := tea.NewView(m.content)
    if m.fullscreen {
        v.AltScreen = true
    }
    return v
}
```

### Dynamic Terminal State

All View fields are evaluated every frame. This means terminal state changes (alt screen, mouse mode, cursor, etc.) are **reactive** — they change when the model state changes, without needing explicit commands.

```go
func (m model) View() tea.View {
    v := tea.NewView(m.renderContent())
    v.MouseMode = tea.MouseModeCellMotion
    if m.showCursor {
        v.Cursor = tea.NewCursor(m.cursorX, m.cursorY)
    }
    if m.focused {
        v.WindowTitle = "My App — Editing"
    } else {
        v.WindowTitle = "My App"
    }
    return v
}
```
