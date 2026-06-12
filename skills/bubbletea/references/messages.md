# Messages — Input Events and Custom Messages

In-depth reference for Bubble Tea's message system: built-in message types, the `Msg` type alias, key/mouse messages, and message filtering.

## The Msg Type

```go
type Msg = uv.Event  // alias to github.com/charmbracelet/ultraviolet Event
```

`Msg` is a type alias to `uv.Event` from the ultraviolet package. Any value can be a message — it's `interface{}`. Messages trigger `Update`, which produces new state + optional command.

## Built-in Message Types

### Lifecycle Messages

| Type | When Sent |
|------|-----------|
| `QuitMsg struct{}` | Program should exit |
| `InterruptMsg struct{}` | SIGINT / Ctrl+C received |
| `SuspendMsg struct{}` | Program should suspend (Ctrl+Z, Unix only — ignored on Windows) |
| `ResumeMsg struct{}` | Program resumed from suspend |

### Terminal Messages

| Type | Fields | When Sent |
|------|--------|-----------|
| `WindowSizeMsg` | `Width, Height int` | Terminal resized (and on startup) |
| `ColorProfileMsg` | `colorprofile.Profile` (embedded) | Color profile detected/changed |
| `EnvMsg` | `uv.Environ` (type alias — Getenv/LookupEnv) | On startup — environment variables |
| `CapabilityMsg` | `Content string` | Terminal capability report (e.g., "RGB", "Tc") |
| `ModeReportMsg` | `Mode ansi.Mode`, `Value ansi.ModeValue` | Terminal mode report response |
| `KeyboardEnhancementsMsg` | `Flags ansi.KeyboardFlags` | Terminal keyboard enhancement support report |
| `TerminalVersionMsg` | `Name string` | Terminal version/identification |
| `ForegroundColorMsg` | `color.Color` (embedded) | Terminal foreground color report |
| `BackgroundColorMsg` | `color.Color` (embedded) | Terminal background color report |
| `CursorColorMsg` | `color.Color` (embedded) | Terminal cursor color report |
| `CursorPositionMsg` | `X, Y int` | Terminal cursor position report |
| `ClipboardMsg` | `Content string`, `Selection byte` | Clipboard content from terminal |

### Focus Messages

| Type | When Sent |
|------|-----------|
| `FocusMsg` | Terminal gained focus (requires `View.ReportFocus = true`) |
| `BlurMsg` | Terminal lost focus (requires `View.ReportFocus = true`) |

### Paste Messages

| Type | Fields | When Sent |
|------|--------|-----------|
| `PasteStartMsg` | `struct{}` | Bracketed paste started |
| `PasteEndMsg` | `struct{}` | Bracketed paste ended |
| `PasteMsg` | `Content string` | Paste content |

## Key Messages

### KeyMsg Interface

```go
type KeyMsg interface {
    fmt.Stringer
    Key() Key
}
```

`KeyMsg` is an **interface** in v2 (was a struct in v1). Two implementations:

| Type | Meaning |
|------|---------|
| `KeyPressMsg` | Key pressed down |
| `KeyReleaseMsg` | Key released (requires `KeyboardEnhancements.ReportEventTypes`) |

### Key Struct

```go
type Key struct {
    Text        string  // printable characters (e.g., "a", "1")
    Mod         KeyMod  // modifier bitmask
    Code        rune    // key code (special or printable)
    ShiftedCode rune    // shifted key code
    BaseCode    rune    // key code per standard PC-101 layout
    IsRepeat    bool    // key repeat (requires ReportEventTypes)
}
```

### KeyMod — Modifier Bitmask

```go
type KeyMod = uv.KeyMod
```

| Modifier | Value |
|----------|-------|
| `ModCtrl` | Ctrl key |
| `ModAlt` | Alt/Option key |
| `ModShift` | Shift key |
| `ModMeta` | Meta key |
| `ModHyper` | Hyper key |
| `ModSuper` | Super/Windows key |
| `ModCapsLock` | Caps Lock state |
| `ModNumLock` | Num Lock state |
| `ModScrollLock` | Scroll Lock state (Windows API only) |

Check modifiers:

```go
if msg.Mod.Contains(tea.ModAlt) { /* Alt was held */ }
if msg.Mod.Contains(tea.ModCtrl) { /* Ctrl was held */ }
```

### Key Constants

**Navigation:**

| Constant | Key |
|----------|-----|
| `KeyUp` | Up arrow |
| `KeyDown` | Down arrow |
| `KeyLeft` | Left arrow |
| `KeyRight` | Right arrow |
| `KeyHome` | Home |
| `KeyEnd` | End |
| `KeyPgUp` | Page Up |
| `KeyPgDown` | Page Down |
| `KeyInsert` | Insert |
| `KeyDelete` | Delete |

**Special (C0):**

| Constant | Key |
|----------|-----|
| `KeyBackspace` | Backspace |
| `KeyTab` | Tab |
| `KeyEnter` | Enter |
| `KeyReturn` | Enter (alias) |
| `KeyEscape` | Escape |
| `KeyEsc` | Escape (alias) |

**Special (G0):**

| Constant | Key |
|----------|-----|
| `KeySpace` | Space |

**Function keys:** `KeyF1` through `KeyF48`

**Keypad:** `KeyKpEnter`, `KeyKpUp`, `KeyKpDown`, `KeyKpLeft`, `KeyKpRight`, `KeyKpPgUp`, `KeyKpPgDown`, `KeyKpHome`, `KeyKpEnd`, `KeyKpInsert`, `KeyKpDelete`, `KeyKpBegin`

**Other:** `KeyExtended`, `KeyBegin`, `KeyFind`, `KeySelect`

### Key Matching Patterns

**By string (convenient):**

```go
switch msg := msg.(type) {
case tea.KeyPressMsg:
    switch msg.String() {
    case "enter":
    case "ctrl+c":
    case "shift+tab":
    case "alt+enter":
    case "q":
    }
}
```

**By code (type-safe):**

```go
case tea.KeyMsg:
    switch key := msg.Key(); key.Code {
    case tea.KeyEnter:
    case tea.KeyUp:
    default:
        switch key.Text {
        case "j":
        case "k":
        }
    }
```

**By modifier:**

```go
case tea.KeyPressMsg:
    if msg.Mod.Contains(tea.ModCtrl) && msg.Code == tea.KeyC {
        // Ctrl+C
    }
```

### v1 vs v2 Key Differences

| v1 | v2 |
|----|-----|
| `tea.KeyMsg` (struct) | `tea.KeyMsg` (interface) |
| `msg.Type` (tea.KeyType) | `msg.Code` (rune) |
| `msg.Runes` ([]rune) | `msg.Text` (string) |
| `msg.Alt` (bool) | `msg.Mod.Contains(tea.ModAlt)` |
| `case "q":` matches directly | `case "q":` still works via String() |
| `case " ":` matches space | `case "space":` in v2 |

## Mouse Messages

### MouseMsg Interface

```go
type MouseMsg interface {
    fmt.Stringer
    Mouse() Mouse
}
```

Four implementations:

| Type | When Sent |
|------|-----------|
| `MouseClickMsg` | Mouse button clicked |
| `MouseReleaseMsg` | Mouse button released |
| `MouseWheelMsg` | Scroll wheel |
| `MouseMotionMsg` | Mouse moved (requires `MouseModeAllMotion`) |

### Mouse Struct

```go
type Mouse struct {
    X, Y   int
    Button MouseButton
    Mod    KeyMod
}
```

Access via the `Mouse()` method:

```go
case tea.MouseClickMsg:
    m := msg.Mouse()
    fmt.Printf("Click at (%d, %d) button=%v\n", m.X, m.Y, m.Button)
```

### MouseButton Constants

| Constant | Button |
|----------|--------|
| `MouseLeft` | Left button |
| `MouseMiddle` | Middle button |
| `MouseRight` | Right button |
| `MouseWheelUp` | Scroll up |
| `MouseWheelDown` | Scroll down |
| `MouseWheelLeft` | Scroll left |
| `MouseWheelRight` | Scroll right |
| `MouseBackward` | Back button |
| `MouseForward` | Forward button |
| `MouseButton10` | Button 10 |
| `MouseButton11` | Button 11 |

### MouseMode (on View)

| Mode | Events |
|------|--------|
| `MouseModeNone` | No mouse events |
| `MouseModeCellMotion` | Click, release, wheel, drag |
| `MouseModeAllMotion` | All including motion |

### View.OnMouse Handler

Intercepts mouse events before they reach `Update`:

```go
v.OnMouse = func(msg tea.MouseMsg) tea.Cmd {
    m := msg.Mouse()
    if m.Button == tea.MouseLeft && m.Y == 0 {
        return func() tea.Msg { return headerClickMsg{x: m.X} }
    }
    return nil  // pass through to Update
}
```

## Custom Messages

Any type can be a message. Convention: define a type per event.

```go
// Simple message
type tickMsg time.Time

// Message with data
type userLoadedMsg struct {
    user User
    err  error
}

// Enum-style message
type statusMsg string

const (
    statusReady  statusMsg = "ready"
    statusLoading statusMsg = "loading"
    statusError   statusMsg = "error"
)
```

### Type-Switch Pattern

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        // handle key
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    case userLoadedMsg:
        if msg.err != nil {
            return m, tea.Quit
        }
        m.user = msg.user
    case tickMsg:
        m.count++
    }
    return m, nil
}
```

## Message Filtering

### WithFilter Program Option

```go
func WithFilter(filter func(Model, Msg) Msg) ProgramOption
```

Called **before** the event loop processes each message. Return:
- The original `msg` → pass through normally
- A **different** `Msg` → substitute the message
- `nil` → **drop** the message entirely

### Prevent Quit with Unsaved Changes

```go
func filter(m tea.Model, msg tea.Msg) tea.Msg {
    if _, ok := msg.(tea.QuitMsg); !ok {
        return msg
    }
    model := m.(myModel)
    if model.hasUnsavedChanges {
        return nil  // block quit
    }
    return msg
}

p := tea.NewProgram(model, tea.WithFilter(filter))
```

### Transform Messages

```go
func filter(m tea.Model, msg tea.Msg) tea.Msg {
    // Convert all mouse motions to a no-op when not in mouse mode
    if _, ok := msg.(tea.MouseMotionMsg); ok && !m.(myModel).mouseEnabled {
        return nil
    }
    // Remap Ctrl+C to custom interrupt
    if k, ok := msg.(tea.KeyPressMsg); ok && k.String() == "ctrl+c" {
        return customInterruptMsg{}
    }
    return msg
}
```

### Filter Order

The filter runs **after** `translateInputEvent` but **before** the event loop's internal message handling. This means:
- Raw ultraviolet events are already translated to tea messages
- The filter can intercept `QuitMsg`, `InterruptMsg`, etc.
- The filter can intercept `BatchMsg` and `sequenceMsg` (but this is unusual)

## Input Event Translation

The `translateInputEvent` function converts ultraviolet events to tea messages:

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
