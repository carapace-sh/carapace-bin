# Design Patterns and Best Practices

Battle-tested patterns from production Bubble Tea applications (derived from real-world codebases like Charm's Crush). Covers architecture, performance, state management, external event bridging, testing, and rendering strategies.

## Architecture Patterns

### Pointer-Based Model with Draw/View Split

For complex applications, use a **pointer-based model** (`*UI` not `UI`) with a `Draw(scr uv.Screen, area uv.Rectangle)` method that renders to an `uv.ScreenBuffer`, and a `View()` that wraps the buffer into a `tea.View`. This separates layout computation from the declarative View struct.

```go
type UI struct { /* fields */ }

// Draw renders to a screen buffer — the real rendering logic lives here.
func (m *UI) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
    layout := m.generateLayout(area.Dx(), area.Dy())
    screen.Clear(scr)

    switch m.state {
    case stateChat:
        m.chat.Draw(scr, layout.main)
        m.status.Draw(scr, layout.status)
    }

    if m.dialog.HasDialogs() {
        return m.dialog.Draw(scr, scr.Bounds())
    }
    return nil
}

// View creates the declarative View struct from the screen buffer.
func (m *UI) View() tea.View {
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

**Why**: The `Draw` method gives pixel-level control over screen regions (sidebar, chat, editor, status bar) using `uv.Rectangle` positioning. `View()` is a thin wrapper that sets declarative flags and serializes the buffer. This is the pattern used by Charm's own production apps.

### Common Struct for Shared Dependencies

Inject shared dependencies (styles, workspace, config) through a `Common` struct rather than passing them individually:

```go
type Common struct {
    Workspace workspace.Workspace
    Styles    *styles.Styles
}

func (c *Common) Config() *config.Config {
    return c.Workspace.Config()
}
```

Every component receives `*Common` at construction:

```go
func NewChat(com *common.Common) *Chat { ... }
func NewStatus(com *common.Common, km help.KeyMap) *Status { ... }
func newHeader(com *common.Common) *header { ... }
```

**Why**: Avoids prop drilling. When a component needs config, styles, or workspace access, it goes through `com`. Adding a new dependency only requires updating `Common`, not every constructor signature.

### State Machine for UI Modes

Use an explicit state enum to drive rendering and input handling:

```go
type uiState uint8

const (
    uiOnboarding uiState = iota
    uiInitialize
    uiLanding
    uiChat
)

type uiFocusState uint8

const (
    uiFocusNone uiFocusState = iota
    uiFocusEditor
    uiFocusMain
)
```

In `Update`, switch on state to determine which messages to process:

```go
func (m *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        switch m.state {
        case uiChat:
            // handle chat keys
        case uiLanding:
            // handle landing keys
        }
    }
}
```

In `Draw`, switch on state to determine which regions to render:

```go
func (m *UI) Draw(scr uv.Screen, area uv.Rectangle) *tea.Cursor {
    switch m.state {
    case uiOnboarding:
        m.drawHeader(scr, layout.header)
    case uiChat:
        m.chat.Draw(scr, layout.main)
        m.status.Draw(scr, layout.status)
    }
}
```

**Why**: Prevents key bindings from one mode leaking into another. Makes the rendering path explicit and easy to reason about.

### Feature Methods on the Main Model

For features that need access to multiple sub-components (e.g., prompt history needs both the textarea and the history list), put the logic as methods on the main model in a dedicated file:

```
model/
  ui.go          # core Update/View/Draw
  history.go     # prompt history navigation
  session.go     # session loading
  sidebar.go     # sidebar rendering
  pills.go       # pills/todos rendering
```

Each file defines private message types and `tea.Cmd` factories for its feature:

```go
// history.go
type promptHistoryLoadedMsg struct {
    messages []string
}

func (m *UI) loadPromptHistory() tea.Cmd {
    return func() tea.Msg {
        // async load
        return promptHistoryLoadedMsg{messages: msgs}
    }
}

func (m *UI) handleHistoryUp() tea.Cmd {
    // mutate m.promptHistory and m.textarea
}
```

**Why**: Keeps `ui.go` focused on the core Update/View dispatch. Feature logic is co-located but doesn't bloat the main file.

## External Event Bridging

### PubSub → program.Send() Bridge

The only way to inject external events into Bubble Tea's single-threaded update loop is `program.Send()`. Build a bridge goroutine that reads from a channel and calls `Send()`:

```go
func (app *App) Subscribe(program *tea.Program) {
    events := app.broker.Subscribe(ctx)
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
go ws.Subscribe(program)
if _, err := program.Run(); err != nil { ... }
```

### Fan-In Architecture

When multiple services produce events, use a central broker with fan-in goroutines:

```
Service goroutines (agent, LSP, MCP, permissions)
    │ publish to per-service pubsub.Broker[T]
    ▼
setupSubscriber goroutines (fan-in)
    │ re-publish into central Broker[tea.Msg]
    ▼
Subscribe goroutine
    │ reads <-chan and calls program.Send(msg)
    ▼
tea.Program (Bubble Tea event loop)
```

### Lossy vs Must-Deliver Publishing

For high-frequency streaming (e.g., token deltas), use **lossy** publishing that drops events when the subscriber channel is full. For terminal events (e.g., run complete), use **must-deliver** with bounded blocking:

```go
// Lossy: drops if channel full (streaming tokens)
broker.Publish(pubsub.UpdatedEvent, msg)

// Must-deliver: blocks up to 50ms per subscriber before dropping
broker.PublishMustDeliver(ctx, pubsub.UpdatedEvent, msg)
```

**Why**: The Bubble Tea update loop is single-threaded. If it's busy rendering, the channel fills up. Lossy delivery prevents backpressure from blocking the agent; must-deliver ensures critical state transitions aren't lost.

## Performance Patterns

### Versioned Render Cache (Freeze Finished Items)

For lists of items where most are static (chat messages), use a version counter to avoid re-rendering unchanged items:

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

**Why**: In a chat UI, old messages never change. Freezing their rendered output means the list only calls `Render()` on the newest (streaming) item. This turns O(n) rendering into O(1) for stable frames.

### Per-Section Render Cache

When a message has multiple independently-changing sections (thinking, content, error), cache each section separately so streaming one doesn't invalidate another:

```go
type assistantSection struct {
    width   int
    srcHash uint64  // FNV-64 of source text
    extra   uint64  // other state that affects render
    out     string
    h       int
    valid   bool
}

func (s *assistantSection) hit(width int, srcHash, extra uint64) bool {
    return s.valid && s.width == width && s.srcHash == srcHash && s.extra == extra
}
```

**Why**: Streaming a thinking block shouldn't force re-rendering the content block (which may involve expensive glamour/markdown processing). Per-section caches isolate invalidation.

### Screen Buffer Caching (Draw Cache)

When rendering to `uv.ScreenBuffer`, cache the decoded buffer to skip ANSI re-parsing on frames with identical content:

```go
type chatDrawCache struct {
    rendered string
    method   ansi.Method
    buf      uv.ScreenBuffer
}

func (m *Chat) Draw(scr uv.Screen, area uv.Rectangle) {
    rendered := m.list.Render()
    method, ok := scr.WidthMethod().(ansi.Method)
    if m.drawCache == nil || m.drawCache.rendered != rendered || m.drawCache.method != method {
        m.drawCache = newChatDrawCache(rendered, method)
    }
    drawCachedBuffer(scr, area, m.drawCache.buf)
}
```

**Why**: `uv.StyledString.Draw` re-parses ANSI sequences on every call. For a chat list that hasn't changed between frames, this is pure waste. Caching the decoded buffer turns the draw into an O(cells) copy.

### Message Throttling via WithFilter

Use `tea.WithFilter` to throttle high-frequency events (mouse motion, streaming tokens) before they reach `Update`:

```go
func MouseEventFilter(m tea.Model, msg tea.Msg) tea.Msg {
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

p := tea.NewProgram(model, tea.WithFilter(MouseEventFilter))
```

**Why**: Mouse motion events can fire hundreds of times per second. Without throttling, the update loop spends all its time processing mouse events instead of rendering.

### Animation Visibility Tracking

Pause animations for items scrolled out of view, restart when they become visible:

```go
type Chat struct {
    pausedAnimations map[string]struct{}
}

func (c *Chat) Animate(msg anim.StepMsg) tea.Cmd {
    // Only tick animations for visible items
    for id := range c.pausedAnimations {
        if c.isItemVisible(id) {
            delete(c.pausedAnimations, id)
            return anim.New(settings).Animate()
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
    dialogs []Dialog
    graceOpenedAt    time.Time
    graceLastInputAt time.Time
}
```

**Grace period**: When a dialog opens asynchronously (e.g., permission prompt from agent), absorb keystrokes for a short window to prevent in-flight keys from acting on the dialog before the user sees it:

```go
func (d *Overlay) Update(msg tea.Msg) tea.Msg {
    if _, ok := msg.(tea.KeyPressMsg); ok && d.inGracePeriod() {
        d.graceLastInputAt = time.Now()
        return nil  // absorb the keystroke
    }
    // ... dispatch to front dialog
}
```

**Why**: Without the grace period, a user typing in the editor could accidentally accept or dismiss a permission dialog that appeared mid-keystroke.

### Interface-Based Item Types

Define item capabilities as small, composable interfaces rather than a single monolithic type:

```go
type MessageItem interface {
    list.Item
    list.RawRenderable
    Identifiable
}

type HighlightableMessageItem interface {
    MessageItem
    list.Highlightable
}

type FocusableMessageItem interface {
    MessageItem
    list.Focusable
}

type Animatable interface {
    StartAnimation() tea.Cmd
    Animate(msg anim.StepMsg) tea.Cmd
}

type Expandable interface {
    ToggleExpanded() bool
}
```

**Why**: Items opt into only the capabilities they need. The list can check `if animatable, ok := item.(chat.Animatable); ok` without forcing all items to implement animation. New capabilities are added as new interfaces without breaking existing types.

### Standalone Renderer Components (Builder Pattern)

For rendering components that don't participate in the Update loop (diff views, formatted output), use a builder pattern with a `String()` method:

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
type Status struct {
    msg      util.InfoMsg
    help     help.Model
}

func (m *UI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    case util.InfoMsg:
        m.status.SetInfoMsg(msg)
        ttl := msg.TTL
        if ttl <= 0 { ttl = DefaultStatusTTL }
        cmds = append(cmds, clearInfoMsgCmd(ttl))
}

func clearInfoMsgCmd(ttl time.Duration) tea.Cmd {
    return tea.Tick(ttl, func(t time.Time) tea.Msg {
        return util.ClearStatusMsg{}
    })
}
```

**Why**: Status messages (success, error, info) should auto-dismiss. Using `tea.Tick` for TTL keeps the cleanup in the update loop rather than requiring a separate timer mechanism.

### Double-Click Detection with Delayed Single-Click

Distinguish single and double clicks by delaying the single-click action:

```go
case tea.MouseReleaseMsg:
    if m.chat.HandleMouseUp(x, y) && m.chat.HasHighlight() {
        cmds = append(cmds, tea.Tick(doubleClickThreshold, func(t time.Time) tea.Msg {
            if time.Since(m.lastClickTime) >= doubleClickThreshold {
                return copyChatHighlightMsg{}  // single click → copy
            }
            return nil  // double click occurred, skip
        }))
    }
```

**Why**: The terminal protocol doesn't distinguish single/double clicks. Delaying the single-click action by the double-click threshold lets you handle both correctly.

## Testing Patterns

### Direct Model Testing (No tea.Program)

For complex UIs, test models directly without running a `tea.Program`:

```go
func newTestUI() *UI {
    com := common.DefaultCommon(nil)
    u := &UI{
        com:      com,
        status:   NewStatus(com, nil),
        chat:     NewChat(com),
        textarea: textarea.New(),
        state:    uiChat,
        focus:    uiFocusEditor,
        width:    140,
        height:   45,
    }
    return u
}

func TestLayout(t *testing.T) {
    u := newTestUI()
    u.updateLayoutAndSize()
    // assert on layout regions
}
```

**Why**: `tea.Program` requires a real TTY or extensive mocking. Direct model testing is faster, more deterministic, and tests the actual logic without the framework overhead.

### Screen Buffer Rendering Assertions

Render to `uv.ScreenBuffer` and assert on the output:

```go
func renderToBuffer(t *testing.T, c *Chat, w, h int) string {
    scr := uv.NewScreenBuffer(w, h)
    c.Draw(scr, drawTestArea(w, h))
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
    Editor struct {
        AddFile     key.Binding
        SendMessage key.Binding
        Newline     key.Binding
    }
    Chat struct {
        NewSession key.Binding
        Cancel     key.Binding
        Up         key.Binding
        Down       key.Binding
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
func (m *UI) ShortHelp() []key.Binding {
    if m.isAgentBusy() {
        cancelBinding := m.keyMap.Chat.Cancel
        if m.isCanceling {
            cancelBinding.SetHelp("esc", "press again to cancel")
        } else {
            cancelBinding.SetHelp("esc", "cancel")
        }
        binds = append(binds, cancelBinding)
    }
    // ...
}
```

**Why**: Static help text is misleading when bindings change meaning (e.g., "esc" means "cancel" vs "press again to cancel" vs "close dialog").

### Keyboard Enhancement Adaptation

Adjust key bindings based on terminal capabilities:

```go
case tea.KeyboardEnhancementsMsg:
    if msg.SupportsKeyDisambiguation() {
        m.keyMap.Editor.Newline.SetHelp("shift+enter", "newline")
    }
```

**Why**: Terminals without Kitty keyboard protocol can't distinguish `Enter` from `Shift+Enter`. Adapting help text prevents user confusion.

## Style and Theme Patterns

### Centralized Style Struct

Define all styles in a single struct with sub-structs for each UI region:

```go
type Styles struct {
    Header struct {
        Charm     lipgloss.Style
        WorkingDir lipgloss.Style
    }
    Editor struct {
        Textarea textarea.Styles
        PromptNormalFocused lipgloss.Style
    }
    Status struct {
        Help lipgloss.Style
        Info lipgloss.Style
    }
    // ...
}
```

**Why**: A single `Styles` struct makes theme switching trivial (swap the whole struct). Components access styles through `com.Styles.Header.Charm` rather than defining their own.

### Theme Selection Based on Provider

Choose theme dynamically based on runtime context:

```go
func ThemeForProvider(providerID string) Styles {
    switch providerID {
    case "hyper":
        return hyperTheme()
    default:
        return defaultTheme()
    }
}
```

**Why**: Different providers (OpenAI, Anthropic, Hyper) may have different brand colors. Theme selection at construction time keeps the rest of the code provider-agnostic.

### Icon Constants

Define UI icons as constants for consistency:

```go
const (
    CheckIcon       = "✓"
    SpinnerIcon     = "⋯"
    ToolPending     = "●"
    ToolSuccess     = "✓"
    ToolError       = "×"
    RadioOn         = "◉"
    RadioOff        = "○"
    BorderThin      = "│"
    BorderThick     = "▌"
    ScrollbarThumb  = "┃"
    ScrollbarTrack  = "│"
)
```

**Why**: Centralized icon definitions ensure visual consistency and make it easy to swap icon sets (e.g., ASCII fallbacks for limited terminals).
