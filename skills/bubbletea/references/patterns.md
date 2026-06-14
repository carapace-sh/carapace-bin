# Design Patterns and Best Practices

Battle-tested patterns from production Bubble Tea applications. Covers architecture, performance, state management, external event bridging, testing, and rendering strategies.

## Architecture Patterns

### Pointer-Based Model with Draw/View Split

For complex applications, use a **pointer-based model** (`*Model` not `Model`) with a `Draw(scr uv.Screen, area uv.Rectangle)` method that renders to an `uv.ScreenBuffer`, and a `View()` that wraps the buffer into a `tea.View`. This separates layout computation from the declarative View struct.

```go
type AppModel struct { /* fields */ }

// Draw renders to a screen buffer — the real rendering logic lives here.
func (m *AppModel) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
    layout := m.generateLayout(area.Dx(), area.Dy())
    screen.Clear(scr)

    switch m.state {
    case stateMain:
        m.mainView.Draw(scr, layout.main)
        m.statusBar.Draw(scr, layout.status)
    }

    if m.dialog.HasDialogs() {
        return m.dialog.Draw(scr, scr.Bounds())
    }
    return nil
}

// View creates the declarative View struct from the screen buffer.
func (m *AppModel) View() tea.View {
    var v tea.View
    v.AltScreen = true
    v.BackgroundColor = m.styles.Background
    v.MouseMode = tea.MouseModeCellMotion

    canvas := uv.NewScreenBuffer(m.width, m.height)
    v.Cursor = m.Draw(canvas, canvas.Bounds())
    v.Content = canvas.Render()
    return v
}
```

**Why**: The `Draw` method gives pixel-level control over screen regions (sidebar, main content, editor, status bar) using `uv.Rectangle` positioning. `View()` is a thin wrapper that sets declarative flags and serializes the buffer. This is the pattern used by Charm's own production apps.

### Shared Dependencies Struct

Inject shared dependencies (styles, config, services) through a single struct rather than passing them individually:

```go
type Shared struct {
    Config *config.Config
    Styles *styles.Styles
    Store  *store.Store  // any app-specific data store
}
```

Every component receives `*Shared` at construction:

```go
func NewChat(s *Shared) *Chat { ... }
func NewStatusBar(s *Shared, km help.KeyMap) *StatusBar { ... }
func NewSidebar(s *Shared) *Sidebar { ... }
```

**Why**: Avoids prop drilling. When a component needs config, styles, or data store access, it goes through the shared struct. Adding a new dependency only requires updating `Shared`, not every constructor signature.

### State Machine for UI Modes

Use an explicit state enum to drive rendering and input handling:

```go
type appState uint8

const (
    stateSetup appState = iota
    stateLoading
    stateMain
    stateHelp
)

type focusState uint8

const (
    focusNone focusState = iota
    focusInput
    focusContent
)
```

In `Update`, switch on state to determine which messages to process:

```go
func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch m.state {
        case stateMain:
            // handle main view keys
        case stateHelp:
            // handle help view keys
        }
    }
}
```

In `Draw`, switch on state to determine which regions to render:

```go
func (m *AppModel) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
    switch m.state {
    case stateSetup:
        m.drawSetup(scr, layout.main)
    case stateMain:
        m.mainView.Draw(scr, layout.main)
        m.statusBar.Draw(scr, layout.status)
    }
}
```

**Why**: Prevents key bindings from one mode leaking into another. Makes the rendering path explicit and easy to reason about.

### Feature Methods on the Main Model

For features that need access to multiple sub-components (e.g., command history needs both the input field and the history list), put the logic as methods on the main model in a dedicated file:

```
model/
  app.go          # core Update/View/Draw
  history.go      # command history navigation
  navigation.go   # session/view loading
  sidebar.go      # sidebar rendering
  toolbar.go      # toolbar/todos rendering
```

Each file defines private message types and `tea.Cmd` factories for its feature:

```go
// history.go
type historyLoadedMsg struct {
    entries []string
}

func (m *AppModel) loadHistory() tea.Cmd {
    return func() tea.Msg {
        // async load
        return historyLoadedMsg{entries: entries}
    }
}

func (m *AppModel) handleHistoryUp() tea.Cmd {
    // mutate m.history and m.textInput
}
```

**Why**: Keeps `app.go` focused on the core Update/View dispatch. Feature logic is co-located but doesn't bloat the main file.

## External Event Bridging

### Channel → program.Send() Bridge

The only way to inject external events into Bubble Tea's single-threaded update loop is `program.Send()`. Build a bridge goroutine that reads from a channel and calls `Send()`:

```go
func bridgeEvents(ctx context.Context, events <-chan Event, program *tea.Program) {
    for {
        select {
        case <-ctx.Done():
            return
        case ev, ok := <-events:
            if !ok { return }
            program.Send(ev)  // ← bridge: channel → tea.Msg
        }
    }
}
```

Start the bridge **before** `program.Run()`:

```go
go bridgeEvents(ctx, eventCh, program)
if _, err := program.Run(); err != nil { ... }
```

### Fan-In Architecture

When multiple sources produce events, use a central channel with fan-in goroutines:

```
Source goroutines (background workers, file watchers, network events)
    │ send to per-source channels
    ▼
fan-in goroutines
    │ forward into central event channel
    ▼
bridge goroutine
    │ reads <-chan and calls program.Send(msg)
    ▼
tea.Program (Bubble Tea event loop)
```

### Lossy vs Must-Deliver Publishing

For high-frequency streaming (e.g., progress updates, streaming text), use **lossy** publishing that drops events when the subscriber channel is full. For terminal events (e.g., operation complete), use **must-deliver** with bounded blocking:

```go
// Lossy: drops if channel full (high-frequency streaming)
select {
case events <- msg:
default:
    // channel full, drop event
}

// Must-deliver: blocks briefly before dropping
ctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
defer cancel()
select {
case events <- msg:
case <-ctx.Done():
    // timed out, log and drop
}
```

**Why**: The Bubble Tea update loop is single-threaded. If it's busy rendering, the channel fills up. Lossy delivery prevents backpressure from blocking background workers; must-deliver ensures critical state transitions aren't lost.

## Performance Patterns

### Versioned Render Cache (Freeze Finished Items)

For lists of items where most are static (log entries, chat messages, table rows), use a version counter to avoid re-rendering unchanged items:

```go
type Versioned struct {
    v uint64
}

func (vc *Versioned) Bump() { vc.v++ }
func (vc *Versioned) Version() uint64 { return vc.v }

type Item interface {
    Render(width int) string
    Version() uint64
    Finished() bool  // true → output won't change, safe to freeze
}
```

The list cache checks `(pointer, width, version)`:

```go
type listCacheEntry struct {
    width   int
    version uint64
    frozen  bool    // once Finished() && rendered, skip Render() calls
    content string
    lines   []string
    height  int
}
```

**Why**: In a scrollable list, old items rarely change. Freezing their rendered output means the list only calls `Render()` on the newest (still-mutating) item. This turns O(n) rendering into O(1) for stable frames.

### Per-Section Render Cache

When a composite item has multiple independently-changing sections (header, body, footer), cache each section separately so updating one doesn't invalidate another:

```go
type sectionCache struct {
    width   int
    srcHash uint64  // hash of source text
    extra   uint64  // other state that affects render
    out     string
    h       int
    valid   bool
}

func (s *sectionCache) hit(width int, srcHash, extra uint64) bool {
    return s.valid && s.width == width && s.srcHash == srcHash && s.extra == extra
}
```

**Why**: Streaming one section (e.g., a progress indicator) shouldn't force re-rendering another (which may involve expensive markdown/code processing). Per-section caches isolate invalidation.

### Screen Buffer Caching (Draw Cache)

When rendering to `uv.ScreenBuffer`, cache the decoded buffer to skip ANSI re-parsing on frames with identical content:

```go
type drawCache struct {
    rendered string
    method   ansi.Method
    buf      uv.ScreenBuffer
}

func (m *Component) Draw(scr uv.Screen, area uv.Rectangle) {
    rendered := m.list.Render()
    method, ok := scr.WidthMethod().(ansi.Method)
    if m.cache == nil || m.cache.rendered != rendered || m.cache.method != method {
        m.cache = newDrawCache(rendered, method)
    }
    drawCachedBuffer(scr, area, m.cache.buf)
}
```

**Why**: `uv.StyledString.Draw` re-parses ANSI sequences on every call. For content that hasn't changed between frames, this is pure waste. Caching the decoded buffer turns the draw into an O(cells) copy.

### Message Throttling via WithFilter

Use `tea.WithFilter` to throttle high-frequency events (mouse motion, rapid updates) before they reach `Update`:

```go
func throttleFilter(m tea.Model, msg tea.Msg) tea.Msg {
    switch msg.(type) {
    case tea.MouseWheelMsg, tea.MouseMotionMsg:
        now := time.Now()
        if now.Sub(lastMouseEvent) < 15*time.Millisecond {
            return nil  // drop the message
        }
        lastMouseEvent = now
    }
    return msg
}

p := tea.NewProgram(model, tea.WithFilter(throttleFilter))
```

**Why**: Mouse motion events can fire hundreds of times per second. Without throttling, the update loop spends all its time processing mouse events instead of rendering.

### Animation Visibility Tracking

Pause animations for items scrolled out of view, restart when they become visible:

```go
type ScrollView struct {
    pausedAnimations map[string]struct{}
}

func (s *ScrollView) Animate(msg stepMsg) tea.Cmd {
    // Only tick animations for visible items
    for id := range s.pausedAnimations {
        if s.isItemVisible(id) {
            delete(s.pausedAnimations, id)
            return tickAnimation(id)
        }
    }
}
```

**Why**: Running animations for off-screen items wastes CPU and generates unnecessary render cycles.

## Component Patterns

### Dialog Overlay Stack

Manage multiple dialogs as a stack with a grace period for async-opened dialogs:

```go
type Dialog interface {
    ID() string
    HandleMsg(msg tea.Msg) Action
    Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor
}

type Overlay struct {
    dialogs          []Dialog
    graceOpenedAt    time.Time
    graceLastInputAt time.Time
}
```

**Grace period**: When a dialog opens asynchronously (e.g., a confirmation prompt from a background worker), absorb keystrokes for a short window to prevent in-flight keys from acting on the dialog before the user sees it:

```go
func (d *Overlay) Update(msg tea.Msg) tea.Msg {
    if _, ok := msg.(tea.KeyPressMsg); ok && d.inGracePeriod() {
        d.graceLastInputAt = time.Now()
        return nil  // absorb the keystroke
    }
    // ... dispatch to front dialog
}
```

**Why**: Without the grace period, a user typing in the main view could accidentally accept or dismiss a dialog that appeared mid-keystroke.

### Interface-Based Item Types

Define item capabilities as small, composable interfaces rather than a single monolithic type:

```go
type ListItem interface {
    list.Item
    list.RawRenderable
    Identifiable
}

type HighlightableItem interface {
    ListItem
    list.Highlightable
}

type FocusableItem interface {
    ListItem
    list.Focusable
}

type Animatable interface {
    StartAnimation() tea.Cmd
    Animate(msg stepMsg) tea.Cmd
}

type Expandable interface {
    ToggleExpanded() bool
}
```

**Why**: Items opt into only the capabilities they need. The list can check `if a, ok := item.(Animatable); ok` without forcing all items to implement animation. New capabilities are added as new interfaces without breaking existing types.

### Standalone Renderer Components (Builder Pattern)

For rendering components that don't participate in the Update loop (diff views, formatted output, static panels), use a builder pattern with a `String()` method:

```go
dv := diffview.New().
    Before("main.go", beforeContent).
    After("main.go", afterContent).
    Width(80).
    Height(24).
    Style(styles.DefaultDarkStyle())

output := dv.String()  // render to string
```

**Why**: Not every rendering component needs to be a `tea.Model`. Pure renderers are simpler to test and compose. They can be called inline from a `View()` or `Draw()` method.

### Status Bar with TTL-Based Messages

Use a status bar that auto-clears messages after a TTL:

```go
type StatusBar struct {
    msg  StatusMsg
    help help.Model
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case StatusMsg:
        m.statusBar.SetMessage(msg)
        ttl := msg.TTL
        if ttl <= 0 { ttl = DefaultStatusTTL }
        cmds = append(cmds, clearStatusMsgCmd(ttl))
}

func clearStatusMsgCmd(ttl time.Duration) tea.Cmd {
    return tea.Tick(ttl, func(t time.Time) tea.Msg {
        return clearStatusMsg{}
    })
}
```

**Why**: Status messages (success, error, info) should auto-dismiss. Using `tea.Tick` for TTL keeps the cleanup in the update loop rather than requiring a separate timer mechanism.

### Double-Click Detection with Delayed Single-Click

Distinguish single and double clicks by delaying the single-click action:

```go
case tea.MouseReleaseMsg:
    if m.handleMouseUp(x, y) && m.hasSelection() {
        cmds = append(cmds, tea.Tick(doubleClickThreshold, func(t time.Time) tea.Msg {
            if time.Since(m.lastClickTime) >= doubleClickThreshold {
                return singleClickMsg{}  // single click action
            }
            return nil  // double click occurred, skip single-click action
        }))
    }
```

**Why**: The terminal protocol doesn't distinguish single/double clicks. Delaying the single-click action by the double-click threshold lets you handle both correctly.

## Testing Patterns

### Direct Model Testing (No tea.Program)

For complex UIs, test models directly without running a `tea.Program`:

```go
func newTestModel() *AppModel {
    u := &AppModel{
        shared:    NewShared(nil),
        statusBar: NewStatusBar(NewShared(nil), nil),
        mainView:  NewMainView(NewShared(nil)),
        textInput: textinput.New(),
        state:     stateMain,
        focus:     focusInput,
        width:     140,
        height:    45,
    }
    return u
}

func TestLayout(t *testing.T) {
    m := newTestModel()
    m.updateLayoutAndSize()
    // assert on layout regions
}
```

**Why**: `tea.Program` requires a real TTY or extensive mocking. Direct model testing is faster, more deterministic, and tests the actual logic without the framework overhead.

### Screen Buffer Rendering Assertions

Render to `uv.ScreenBuffer` and assert on the output:

```go
func renderToBuffer(t *testing.T, c *Component, w, h int) string {
    scr := uv.NewScreenBuffer(w, h)
    c.Draw(scr, testArea(w, h))
    return scr.Render()
}
```

**Why**: Tests the actual rendering pipeline, not just the model state. Catches layout bugs that unit tests on model fields would miss.

### Version Bump Contract Tests

Assert that every state-mutating method bumps the version counter (so the render cache invalidates):

```go
func requireBump(t *testing.T, name string, item versionedItem, mutate func()) {
    before := item.Version()
    mutate()
    after := item.Version()
    require.Greaterf(t, after, before, "%s must bump Version()", name)
}

func requireNoBump(t *testing.T, name string, item versionedItem, mutate func()) {
    before := item.Version()
    mutate()
    after := item.Version()
    require.Equal(t, before, after, "%s must not bump on no-op", name)
}
```

**Why**: The render cache is only correct if every observable mutation bumps the version. Contract tests ensure new mutators don't forget to bump, and no-ops don't cause unnecessary re-renders.

### Render Hit Counting (Cache Verification)

Wrap items to count `Render()` calls and verify cache behavior:

```go
type trackedItem struct {
    *Versioned
    renderHits int
}

func (t *trackedItem) Render(width int) string {
    t.renderHits++
    return t.body
}

func TestCacheFreezesFinishedItem(t *testing.T) {
    item := &trackedItem{Versioned: NewVersioned(), body: "hello"}
    item.Finish()  // mark as finished
    l := NewList(item)
    l.SetSize(80, 24)
    l.Render()
    l.Render()
    assert.Equal(t, 1, item.renderHits)  // Render called once, cached after
}
```

### Golden File Testing for Visual Regression

Use golden files for complex visual output (diffs, formatted markdown):

```go
func TestDiffView(t *testing.T) {
    dv := diffview.New().Before("main.go", before).After("main.go", after)
    output := dv.String()
    golden.RequireEqual(t, []byte(output))
}
```

**Why**: Snapshot testing catches visual regressions that string-contains assertions miss. Golden files can be reviewed visually and updated with `go test -update`.

### ANSI Stripping for Content Comparison

Strip ANSI sequences before comparing rendered content:

```go
func stripANSI(s string) string {
    return ansi.Strip(s)
}

func TestRenderedContent(t *testing.T) {
    rendered := item.Render(80)
    visible := stripANSI(rendered)
    assert.Contains(t, visible, "expected text")
}
```

**Why**: ANSI escape codes make string comparisons fragile. Stripping them lets you assert on visible content without coupling to styling details.

## Key Binding Patterns

### Structured KeyMap with Sub-Maps

Group key bindings by UI region:

```go
type KeyMap struct {
    Input struct {
        AddFile     key.Binding
        Submit      key.Binding
        Newline     key.Binding
    }
    Content struct {
        NewView  key.Binding
        Cancel   key.Binding
        Up       key.Binding
        Down     key.Binding
    }
    // Global
    Quit     key.Binding
    Help     key.Binding
    Commands key.Binding
}
```

**Why**: Flat key maps become unwieldy in large apps. Sub-structs group related bindings and make the help view more organized.

### Dynamic Help Text

Update help text based on state:

```go
func (m *AppModel) ShortHelp() []key.Binding {
    if m.isBusy() {
        cancelBinding := m.keyMap.Content.Cancel
        if m.confirmingCancel {
            cancelBinding.SetHelp("esc", "press again to confirm")
        } else {
            cancelBinding.SetHelp("esc", "cancel")
        }
        binds = append(binds, cancelBinding)
    }
    // ...
}
```

**Why**: Static help text is misleading when bindings change meaning (e.g., "esc" means "cancel" vs "press again to confirm" vs "close dialog").

### Keyboard Enhancement Adaptation

Adjust key bindings based on terminal capabilities:

```go
case tea.KeyboardEnhancementsMsg:
    if msg.SupportsKeyDisambiguation() {
        m.keyMap.Input.Newline.SetHelp("shift+enter", "newline")
    }
```

**Why**: Terminals without Kitty keyboard protocol can't distinguish `Enter` from `Shift+Enter`. Adapting help text prevents user confusion.

## Style and Theme Patterns

### Centralized Style Struct

Define all styles in a single struct with sub-structs for each UI region:

```go
type Styles struct {
    Header struct {
        Title      lipgloss.Style
        Subtitle   lipgloss.Style
    }
    Input struct {
        Textarea          textarea.Styles
        PromptFocused     lipgloss.Style
        PromptBlurred     lipgloss.Style
    }
    Status struct {
        Help lipgloss.Style
        Info lipgloss.Style
    }
    // ...
}
```

**Why**: A single `Styles` struct makes theme switching trivial (swap the whole struct). Components access styles through `shared.Styles.Header.Title` rather than defining their own.

### Theme Selection Based on Context

Choose theme dynamically based on runtime context (user preference, environment, feature flags):

```go
func ThemeForContext(ctx AppContext) Styles {
    switch ctx.Theme {
    case "dark":
        return darkTheme()
    case "light":
        return lightTheme()
    default:
        return defaultTheme()
    }
}
```

**Why**: Different contexts (user preferences, environment variables, feature flags) may require different visual styles. Theme selection at construction time keeps the rest of the code context-agnostic.

### Icon Constants

Define UI icons as constants for consistency:

```go
const (
    CheckIcon       = "✓"
    SpinnerIcon     = "⋯"
    PendingIcon     = "●"
    SuccessIcon     = "✓"
    ErrorIcon       = "×"
    RadioOn         = "◉"
    RadioOff        = "○"
    BorderThin      = "│"
    BorderThick     = "▌"
    ScrollbarThumb  = "┃"
    ScrollbarTrack  = "│"
)
```

**Why**: Centralized icon definitions ensure visual consistency and make it easy to swap icon sets (e.g., ASCII fallbacks for limited terminals).
