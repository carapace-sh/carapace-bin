# Component Composition and Charm Ecosystem

In-depth reference for composing Bubble Tea applications: embedding sub-models, focus management, layout patterns, building custom components, and integrating with the Charm ecosystem.

## Embedding Sub-Models

The standard pattern for composing multiple components in a single Bubble Tea model:

### 1. Embed as Value-Type Fields

```go
type MainModel struct {
    textInput textinput.Model
    viewport  viewport.Model
    spinner   spinner.Model
    list      list.Model
    table     table.Model
    help      help.Model
}
```

Components are **value types** (not pointers). This is important — you must reassign after `Update`.

### 2. Initialize with New() Constructors

```go
func initialModel() MainModel {
    m := MainModel{
        textInput: textinput.New(),
        viewport:  viewport.New(width, height),
        spinner:   spinner.New(),
        list:      list.New(items, delegate, width, height),
    }
    m.textInput.Placeholder = "Type here..."
    m.textInput.Focus()
    return m
}
```

### 3. Forward Update Calls

```go
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd

    // Handle parent-level messages first
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.viewport.SetWidth(msg.Width)
        m.viewport.SetHeight(msg.Height - statusBarHeight)
    case tea.KeyPressMsg:
        switch msg.String() {
        case "tab":
            m.switchFocus()
        case "ctrl+c":
            return m, tea.Quit
        }
    }

    // Forward to focused sub-model
    var cmd tea.Cmd
    m.textInput, cmd = m.textInput.Update(msg)  // MUST reassign
    cmds = append(cmds, cmd)

    m.spinner, cmd = m.spinner.Update(msg)
    cmds = append(cmds, cmd)

    return m, tea.Batch(cmds...)
}
```

**Critical**: `m.textInput, cmd = m.textInput.Update(msg)` — the reassignment is required because components are value types. If you don't reassign, the update is lost.

### 4. Compose Views

```go
func (m MainModel) View() tea.View {
    content := lipgloss.JoinVertical(lipgloss.Left,
        m.headerView(),
        m.textInput.View(),
        m.viewport.View(),
        m.footerView(),
    )
    v := tea.NewView(content)
    v.AltScreen = true
    v.MouseMode = tea.MouseModeCellMotion
    return v
}
```

## Focus Management

### Focus/Blur Pattern

Most bubbles components implement `Focus()` and `Blur()`:

```go
// Focus returns a Cmd (typically for cursor blink initialization)
func (m *Model) Focus() tea.Cmd {
    m.focus = true
    return m.virtualCursor.Focus()
}

// Blur is void — no return Cmd
func (m *Model) Blur() {
    m.focus = false
    m.virtualCursor.Blur()
}
```

### Update Guard on Focus

Components skip processing when blurred:

```go
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
    if !m.focus {
        return m, nil  // textinput, textarea, table all do this
    }
    // ... process msg
}
```

### Focus Switching in Parent

```go
type focusState int

const (
    inputFocus focusState = iota
    viewportFocus
    listFocus
)

func (m *MainModel) switchFocus() tea.Cmd {
    var cmds []tea.Cmd

    // Blur current
    switch m.focus {
    case inputFocus:
        m.textInput.Blur()
    case listFocus:
        m.list.Blur()
    }

    // Advance to next
    m.focus = (m.focus + 1) % 3

    // Focus new
    switch m.focus {
    case inputFocus:
        cmds = append(cmds, m.textInput.Focus())
    case listFocus:
        cmds = append(cmds, m.list.Focus())
    }

    return tea.Batch(cmds...)
}
```

### Conditional Forwarding

Only forward messages to the focused component:

```go
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch m.focus {
    case inputFocus:
        m.textInput, cmd = m.textInput.Update(msg)
    case viewportFocus:
        m.viewport, cmd = m.viewport.Update(msg)
    case listFocus:
        m.list, cmd = m.list.Update(msg)
    }
    return m, cmd
}
```

## Layout Patterns

### Vertical Stack (Most Common)

```go
func (m Model) View() tea.View {
    return tea.NewView(lipgloss.JoinVertical(lipgloss.Left,
        m.titleView(),
        m.contentView(),
        m.statusView(),
        m.helpView(),
    ))
}
```

### Horizontal Split

```go
func (m Model) View() tea.View {
    sidebar := lipgloss.NewStyle().Width(20).Render(m.sidebarView())
    main := lipgloss.NewStyle().Width(m.width - 20).Render(m.mainView())
    return tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, sidebar, main))
}
```

### Status Bar at Bottom

```go
const statusBarHeight = 1

func (m Model) View() tea.View {
    return tea.NewView(lipgloss.JoinVertical(lipgloss.Left,
        m.viewport.View(),
        m.statusBarView(),
    ))
}

// On WindowSizeMsg:
m.viewport.SetHeight(msg.Height - statusBarHeight)
```

### Fixed Header + Scrollable Body

```go
func (m Model) View() tea.View {
    header := m.headerView()  // fixed height
    body := m.viewport.View() // scrollable, fills remaining space
    return tea.NewView(header + "\n" + body)
}
```

### Bordered Panels

```go
var (
    borderStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("63")).
        Padding(1, 2)
)

func (m Model) View() tea.View {
    left := borderStyle.Width(m.width/2 - 2).Render(m.leftView())
    right := borderStyle.Width(m.width/2 - 2).Render(m.rightView())
    return tea.NewView(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
}
```

## Key Bindings

### key.Binding

```go
type Binding struct {
    keys     []string
    help     Help
    disabled bool
}
```

Create with functional options:

```go
var UpKey = key.NewBinding(
    key.WithKeys("k", "up"),
    key.WithHelp("↑/k", "move up"),
)

var QuitKey = key.NewBinding(
    key.WithKeys("q", "ctrl+c"),
    key.WithHelp("q", "quit"),
)
```

### KeyMap Struct

```go
type KeyMap struct {
    Up    key.Binding
    Down  key.Binding
    Enter key.Binding
    Quit  key.Binding
}

func DefaultKeyMap() KeyMap {
    return KeyMap{
        Up:    key.NewBinding(key.WithKeys("k", "up"), key.WithHelp("↑/k", "up")),
        Down:  key.NewBinding(key.WithKeys("j", "down"), key.WithHelp("↓/j", "down")),
        Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "select")),
        Quit:  key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
    }
}
```

### Matching Keys

```go
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch {
        case key.Matches(msg, m.KeyMap.Up):
            m.cursorUp()
        case key.Matches(msg, m.KeyMap.Down):
            m.cursorDown()
        case key.Matches(msg, m.KeyMap.Enter):
            return m, m.selectItem()
        case key.Matches(msg, m.KeyMap.Quit):
            return m, tea.Quit
        }
    }
    return m, nil
}
```

### Dynamic Enable/Disable

```go
// Disable a binding
m.KeyMap.Quit.SetEnabled(false)

// Re-enable
m.KeyMap.Quit.SetEnabled(true)

// Check if enabled
if m.KeyMap.Quit.Enabled() { ... }
```

### Help Integration

Implement `help.KeyMap` interface:

```go
func (m Model) ShortHelp() []key.Binding {
    return []key.Binding{m.KeyMap.Up, m.KeyMap.Down, m.KeyMap.Enter}
}

func (m Model) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {m.KeyMap.Up, m.KeyMap.Down},
        {m.KeyMap.Enter, m.KeyMap.Quit},
    }
}
```

Render help:

```go
m.help.View(m.KeyMap)  // or m.help.View(m) if m implements help.KeyMap
```

## Building Custom Components

Follow the conventions used by all bubbles components:

```go
package mycomponent

import (
    tea "charm.land/bubbletea/v2"
    "charm.land/bubbles/v2/key"
    "charm.land/lipgloss/v2"
)

// 1. KeyMap
type KeyMap struct {
    Up    key.Binding
    Down  key.Binding
    Enter key.Binding
}

func DefaultKeyMap() KeyMap {
    return KeyMap{
        Up:    key.NewBinding(key.WithKeys("k", "up"), key.WithHelp("↑/k", "up")),
        Down:  key.NewBinding(key.WithKeys("j", "down"), key.WithHelp("↓/j", "down")),
        Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "select")),
    }
}

// 2. Model (value type)
type Model struct {
    focus  bool
    KeyMap KeyMap
    Style  lipgloss.Style
    items  []string
    cursor int
    width  int
    height int
}

// 3. Constructor
func New() Model {
    return Model{
        KeyMap: DefaultKeyMap(),
        Style:  lipgloss.NewStyle(),
    }
}

// 4. Focus/Blur
func (m *Model) Focus() tea.Cmd {
    m.focus = true
    return nil
}

func (m *Model) Blur() {
    m.focus = false
}

// 5. Init
func (m Model) Init() tea.Cmd { return nil }

// 6. Update
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
    if !m.focus {
        return m, nil
    }

    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch {
        case key.Matches(msg, m.KeyMap.Up):
            if m.cursor > 0 {
                m.cursor--
            }
        case key.Matches(msg, m.KeyMap.Down):
            if m.cursor < len(m.items)-1 {
                m.cursor++
            }
        case key.Matches(msg, m.KeyMap.Enter):
            return m, func() tea.Msg { return selectMsg{index: m.cursor} }
        }
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    }
    return m, nil
}

// 7. View
func (m Model) View() string {
    var s strings.Builder
    for i, item := range m.items {
        if i == m.cursor {
            s.WriteString(m.Style.Bold(true).Render("> " + item))
        } else {
            s.WriteString("  " + item)
        }
        s.WriteString("\n")
    }
    return s.String()
}

// 8. Help integration
func (m Model) ShortHelp() []key.Binding {
    return []key.Binding{m.KeyMap.Up, m.KeyMap.Down, m.KeyMap.Enter}
}

func (m Model) FullHelp() [][]key.Binding {
    return [][]key.Binding{
        {m.KeyMap.Up, m.KeyMap.Down},
        {m.KeyMap.Enter},
    }
}
```

### Component Conventions

| Convention | Description |
|-----------|-------------|
| Value-type `Model` | Not a pointer — reassign after `Update` |
| `Focus()` returns `tea.Cmd` | For cursor blink initialization |
| `Blur()` is void | No return value |
| `Update` early-returns when blurred | Skip processing when not focused |
| `KeyMap` struct with `key.Binding` | Define all keybindings |
| `DefaultKeyMap()` constructor | Provide sensible defaults |
| `key.Matches(msg, binding)` | Key dispatch pattern |
| `SetEnabled(bool)` | Dynamic keybinding toggling |
| `lipgloss.Style.Render()` | All styled output |
| Component IDs | Disambiguate messages when multiple instances exist |

## Charm Ecosystem Integration

### Lipgloss (Styling/Layout)

```go
import "charm.land/lipgloss/v2"
```

- `lipgloss.NewStyle().Bold(true).Foreground(color).Render(text)` — styled text
- `lipgloss.JoinVertical(alignment, sections...)` — vertical layout
- `lipgloss.JoinHorizontal(alignment, sections...)` — horizontal layout
- `lipgloss.Width(s)` / `lipgloss.Height(s)` — measure rendered size
- `lipgloss.NewStyle().Width(w).Height(h)` — constrain dimensions

### Bubbles (Component Library)

```go
import "charm.land/bubbles/v2"
```

| Component | Import | Key Features |
|-----------|--------|-------------|
| textinput | `bubbles/textinput` | Single-line text input with placeholder, cursor, suggestions |
| textarea | `bubbles/textarea` | Multi-line text editor with line numbers |
| viewport | `bubbles/viewport` | Scrollable content area with keyboard/mouse scrolling |
| spinner | `bubbles/spinner` | Loading spinner with multiple styles |
| list | `bubbles/list` | Selectable/filterable item list with delegates |
| table | `bubbles/table` | Columnar data table with sorting |
| help | `bubbles/help` | Keybinding help display |
| paginator | `bubbles/paginator` | Page navigation |
| progress | `bubbles/progress` | Visual progress bar |
| cursor | `bubbles/cursor` | Cursor blink management |

### Glamour (Markdown Rendering)

```go
import "charm.land/glamour/v2"
```

Not a Bubble Tea model — renders markdown to ANSI string:

```go
r, _ := glamour.NewTermRenderer(glamour.WithWordWrap(80))
out, _ := r.Render(markdownContent)
m.viewport.SetContent(out)
```

### Huh (Form Library)

```go
import "charm.land/huh/v2"
```

Built on Bubble Tea. Can run standalone or embedded:

```go
// Standalone
form := huh.NewForm(
    huh.NewGroup(
        huh.NewInput().Title("Name"),
        huh.NewSelect().Options(huh.NewOptions("a", "b", "c")...),
    ),
)
err := form.Run()

// Embedded in Bubble Tea app
// huh.Field interface: Model + Focus/Blur + Error/Skip/Zoom + KeyBinds
```

### Harmonica (Spring Animations)

```go
import "github.com/charmbracelet/harmonica"
```

Spring physics for smooth animations:

```go
spring := harmonica.NewSpring(harmonica.Frequency(2.0), harmonica.Damping(0.5))
position, velocity := spring.Update(target, current, velocity)
```

### BubbleZone (Mouse Tracking)

```go
import "github.com/lrstanley/bubblezone"
```

Easy mouse event tracking for components:

```go
zone := bubblezone.New()
zone.Mark("button", styledButton)
// In Update: if zone.Get("button").InBounds(msg) { ... }
```
