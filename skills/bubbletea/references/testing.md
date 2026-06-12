# Testing Bubble Tea Programs

In-depth reference for testing Bubble Tea programs: test patterns, program options for testing, testing commands, and testing components.

## Core Testing Pattern

Use `bytes.Buffer` for I/O and send quit from a goroutine:

```go
func TestMyProgram(t *testing.T) {
    t.Parallel()
    var buf bytes.Buffer
    var in bytes.Buffer

    m := &myModel{}
    p := tea.NewProgram(m,
        tea.WithInput(&in),
        tea.WithOutput(&buf),
    )
    go func() {
        for {
            time.Sleep(time.Millisecond)
            if m.ready.Load() {
                p.Quit()
                return
            }
        }
    }()

    if _, err := p.Run(); err != nil {
        t.Fatal(err)
    }
}
```

## Program Options for Testing

| Option | Purpose |
|--------|---------|
| `tea.WithInput(&bytes.Buffer{})` | Provide input from a buffer (no real TTY) |
| `tea.WithOutput(&bytes.Buffer{})` | Capture output to a buffer |
| `tea.WithoutSignals()` | Ignore OS signals (prevents SIGINT from killing test) |
| `tea.WithoutRenderer()` | Disable rendering entirely (non-TUI testing) |
| `tea.WithContext(ctx)` | Use cancellable context for test control |
| `tea.WithWindowSize(w, h)` | Set initial terminal size |
| `tea.WithColorProfile(p)` | Force a specific color profile |

## Testing with Context Cancellation

```go
func TestWithContext(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    p := tea.NewProgram(model, tea.WithContext(ctx))
    _, err := p.Run()
    if errors.Is(err, tea.ErrProgramKilled) {
        // Expected: context cancelled
    }
}
```

## Testing with Send

```go
func TestWithSend(t *testing.T) {
    p := tea.NewProgram(model,
        tea.WithInput(&bytes.Buffer{}),
        tea.WithOutput(&bytes.Buffer{}),
    )
    go func() {
        p.Send(customMsg{})
        p.Send(tea.Quit())
    }()
    _, err := p.Run()
    if err != nil {
        t.Fatal(err)
    }
}
```

**Note**: `p.Send()` blocks if the program hasn't started yet. Always call from a goroutine.

## Testing Commands Directly

Since `Cmd` is `func() Msg`, you can call commands directly:

```go
func TestTickCommand(t *testing.T) {
    cmd := tea.Tick(10*time.Millisecond, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
    msg := cmd()  // blocks until timer fires
    if _, ok := msg.(tickMsg); !ok {
        t.Errorf("expected tickMsg, got %T", msg)
    }
}
```

## Testing Batch and Sequence

```go
func TestBatch(t *testing.T) {
    var count atomic.Int32
    cmd := tea.Batch(
        func() tea.Msg { count.Add(1); return msg1{} },
        func() tea.Msg { count.Add(1); return msg2{} },
    )
    // Batch is a BatchMsg — need to execute through a program
    // Or test individual commands directly
}
```

## Testing Update Logic

Since `Update` is a pure function (given model + message → new model + command), you can test it directly without a `Program`:

```go
func TestUpdate(t *testing.T) {
    m := myModel{count: 0}
    newM, cmd := m.Update(tickMsg{})

    // Check state change
    if newM.(myModel).count != 1 {
        t.Errorf("expected count=1, got %d", newM.(myModel).count)
    }

    // Check command
    if cmd != nil {
        msg := cmd()
        if _, ok := msg.(tickMsg); !ok {
            t.Errorf("expected tickMsg from cmd, got %T", msg)
        }
    }
}
```

## Testing View Output

```go
func TestView(t *testing.T) {
    m := myModel{items: []string{"foo", "bar"}}
    v := m.View()

    // Check View struct fields
    if !v.AltScreen {
        t.Error("expected alt screen")
    }

    // Check content
    if !strings.Contains(v.Content, "foo") {
        t.Error("expected 'foo' in view content")
    }
}
```

## Testing with Output Capture

```go
func TestOutputCapture(t *testing.T) {
    var buf bytes.Buffer
    p := tea.NewProgram(model,
        tea.WithInput(&bytes.Buffer{}),
        tea.WithOutput(&buf),
    )
    go p.Send(tea.Quit())

    _, err := p.Run()
    if err != nil {
        t.Fatal(err)
    }

    // Check rendered output
    output := buf.String()
    if !strings.Contains(output, "expected text") {
        t.Errorf("unexpected output: %s", output)
    }
}
```

## Testing Error Conditions

```go
func TestInterrupt(t *testing.T) {
    p := tea.NewProgram(model,
        tea.WithInput(&bytes.Buffer{}),
        tea.WithOutput(&bytes.Buffer{}),
    )
    go p.Send(tea.Interrupt())

    _, err := p.Run()
    if !errors.Is(err, tea.ErrInterrupted) {
        t.Errorf("expected ErrInterrupted, got %v", err)
    }
}

func TestKill(t *testing.T) {
    p := tea.NewProgram(model,
        tea.WithInput(&bytes.Buffer{}),
        tea.WithOutput(&bytes.Buffer{}),
    )
    go p.Kill()

    _, err := p.Run()
    if !errors.Is(err, tea.ErrProgramKilled) {
        t.Errorf("expected ErrProgramKilled, got %v", err)
    }
}
```

## Testing with Filter

```go
func TestWithFilter(t *testing.T) {
    var filtered atomic.Bool
    filter := func(m tea.Model, msg tea.Msg) tea.Msg {
        if _, ok := msg.(tea.QuitMsg); ok {
            filtered.Store(true)
            return nil  // block quit
        }
        return msg
    }

    p := tea.NewProgram(model,
        tea.WithInput(&bytes.Buffer{}),
        tea.WithOutput(&bytes.Buffer{}),
        tea.WithFilter(filter),
    )
    go func() {
        p.Send(tea.Quit())
        // Quit was blocked, need to kill instead
        time.Sleep(100 * time.Millisecond)
        p.Kill()
    }()

    p.Run()
    if !filtered.Load() {
        t.Error("expected quit to be filtered")
    }
}
```

## Testing Components (Bubbles)

```go
func TestTextInput(t *testing.T) {
    ti := textinput.New()
    ti.SetValue("hello")

    // Test Update directly
    newTi, cmd := ti.Update(tea.KeyPressMsg{Code: 'a', Text: "a"})
    if newTi.Value() != "helloa" {
        t.Errorf("expected 'helloa', got %s", newTi.Value())
    }
    _ = cmd
}
```

## Test Helpers from Source

The Bubble Tea test suite uses a `testModel` struct with `atomic.Value` fields for thread-safe state checks:

```go
type testModel struct {
    executed atomic.Value
    counter atomic.Int32
}

func (m *testModel) Init() tea.Cmd { return nil }
func (m *testModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    m.executed.Store(true)
    m.counter.Add(1)
    return m, nil
}
func (m *testModel) View() tea.View {
    return tea.NewView("")
}
```

## Common Test Patterns Summary

| Pattern | When to Use |
|---------|-------------|
| Direct `Update()` call | Testing state transitions |
| Direct `cmd()` call | Testing command results |
| `WithInput`/`WithOutput` + `go p.Quit()` | Basic integration test |
| `WithContext` + `cancel()` | Test cancellation/shutdown |
| `WithFilter` | Test message filtering |
| `WithoutRenderer` | Test non-TUI logic |
| `WithWindowSize` | Test resize handling |
| `p.Send(msg)` from goroutine | Test specific message handling |
| `p.Kill()` | Test forced shutdown |
| Output buffer inspection | Test rendered output |
