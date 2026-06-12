# Program — Lifecycle, Options, and Event Loop

In-depth reference for Bubble Tea's `Program` type: creation, options, the event loop, startup/shutdown, and subprocess execution.

## Creating a Program

```go
func NewProgram(model Model, opts ...ProgramOption) *Program
```

Creates a `*Program` with the model and applies all options. Internal channels are initialized (`msgs`, `errs`, `rendererDone`). Also sets up: context (from `externalCtx` or `context.Background`), output (defaults to `os.Stdout`), environ (defaults to `os.Environ()`), fps clamping (default 60, max 120), and `TEA_TRACE` logger if the env var is set.

```go
p := tea.NewProgram(myModel,
    tea.WithAltScreen,  // v1 only — use View.AltScreen in v2
)
```

## Program Options

```go
type ProgramOption func(*Program)
```

| Option | Signature | Description |
|--------|-----------|-------------|
| `WithContext` | `(ctx context.Context) ProgramOption` | External context; cancel → `ErrProgramKilled` |
| `WithOutput` | `(output io.Writer) ProgramOption` | Override output (default: stdout) |
| `WithInput` | `(input io.Reader) ProgramOption` | Override input (default: stdin); `nil` disables input |
| `WithEnvironment` | `(env []string) ProgramOption` | Set env vars (SSH sessions) |
| `WithoutSignalHandler` | `() ProgramOption` | Disable signal handler (`disableSignalHandler = true`) |
| `WithoutCatchPanics` | `() ProgramOption` | Disable panic recovery |
| `WithoutSignals` | `() ProgramOption` | Set `ignoreSignals = 1` (atomic) — signals ignored but handler still runs |
| `WithoutRenderer` | `() ProgramOption` | Disable renderer (daemon mode) |
| `WithFilter` | `(func(Model, Msg) Msg) ProgramOption` | Message filter |
| `WithFPS` | `(fps int) ProgramOption` | Max FPS (default 60, capped 120) |
| `WithColorProfile` | `(profile colorprofile.Profile) ProgramOption` | Force color profile |
| `WithWindowSize` | `(width, height int) ProgramOption` | Initial window size (testing) |

**Removed in v2** (now declarative View fields): `WithAltScreen`, `WithMouseCellMotion`, `WithMouseAllMotion`, `WithReportFocus`, `WithoutBracketedPaste`, `WithInputTTY`, `WithANSICompressor`.

## Program.Run()

```go
func (p *Program) Run() (Model, error)
```

Blocks until the program exits. Returns the final model and an error (if any).

### Return Values

| Condition | Model | Error |
|-----------|-------|-------|
| Graceful quit (`QuitMsg`) | Final model | `nil` |
| Interrupt (`InterruptMsg` / SIGINT) | Final model | `ErrInterrupted` |
| External context cancelled | Final model | `fmt.Errorf("%w: %w", ErrProgramKilled, ctx.Err())` |
| Internal context cancelled | Final model | `ErrProgramKilled` |
| Panic recovered | Final model | `fmt.Errorf("%w: %w", ErrProgramKilled, ErrProgramPanic)` |
| Error + killed | Final model | `fmt.Errorf("%w: %w", ErrProgramKilled, err)` |

### Error Values

```go
var ErrProgramPanic  = errors.New("program experienced a panic")
var ErrProgramKilled = errors.New("program was killed")
var ErrInterrupted   = errors.New("program was interrupted")
```

Check with `errors.Is`:

```go
_, err := p.Run()
if errors.Is(err, tea.ErrInterrupted) {
    // user pressed Ctrl+C
}
```

### Debugging

```go
func LogToFile(path string, prefix string) (*os.File, error)
```

Sets up logging to a file (useful since the TUI occupies the terminal). Don't forget to close the file. Also available: `TEA_TRACE` env var to enable internal trace logging.

## Startup Sequence

`Run()` executes these steps in order:

1. Check `initialModel != nil`
2. Initialize `handlers`, `cmds` channel, `finished` channel
3. Set up context (derive from `externalCtx` or `context.Background`)
4. Set up input (nil → disable, or open TTY)
5. Handle signals (if not disabled)
6. Set up panic recovery (if not disabled)
7. `initTerminal()` — enter raw mode, hide cursor
8. Get initial window size (from `WithWindowSize` or TTY query)
9. Create renderer, send `WindowSizeMsg`, `EnvMsg`
10. `initInputReader()` — start input parsing
11. `startRenderer()` — ticker-based rendering loop
12. Query synchronized output support (mode 2026) and unicode core (mode 2027)
13. Execute `model.Init()` command
14. Render initial view
15. Start resize handler, command handler goroutines
16. Enter `eventLoop()` — blocks until quit/killed
17. After loop: determine if killed or graceful
18. If graceful: render final state
19. `shutdown()` — restore terminal, stop renderer, wait for handlers

## The Event Loop

```go
func (p *Program) eventLoop(model Model, cmds chan Cmd) (Model, error)
```

The core loop processes messages one at a time (single-threaded):

```
for {
    select {
    case <-ctx.Done(): return model, nil
    case err := <-errs: return model, err
    case msg := <-msgs:
        msg = translateInputEvent(msg)     // ultraviolet → tea
        if filter != nil: msg = filter(model, msg)
        if msg == nil: continue

        switch msg := msg.(type) {
        case QuitMsg:       return model, nil
        case InterruptMsg:  return model, ErrInterrupted
        case SuspendMsg:    if suspendSupported { p.suspend() }
        case execMsg:       p.exec(msg.cmd, msg.fn)  // BLOCKS
        case BatchMsg:      go p.execBatchMsg(msg)
        case sequenceMsg:   go p.execSequenceMsg(msg)
        case printLineMessage: p.renderer.insertAbove(msg.messageBody)
        // ... internal messages (WindowSizeMsg, ModeReportMsg, etc.)
        }

        model, cmd = model.Update(msg)  // user Update
        cmds <- cmd                      // dispatch command
        p.render(model)                 // render view
    }
}
```

**Key properties:**
- **Single-threaded**: Messages are processed sequentially. `Update` is never called concurrently.
- **Blocking on exec**: `execMsg` blocks the entire loop. No messages are processed while a subprocess runs.
- **Non-blocking on Batch/Sequence**: These are dispatched to goroutines. The loop continues processing other messages.
- **Filter runs first**: `WithFilter` intercepts messages before any processing.

## Program Methods

### Send — Inject Messages

```go
func (p *Program) Send(msg Msg)
```

Thread-safe message injection from outside the event loop. Blocks if the program hasn't started. No-op if the program has already terminated.

```go
p := tea.NewProgram(model)
go func() {
    time.Sleep(5 * time.Second)
    p.Send(timeoutMsg{})
}()
p.Run()
```

### Quit — Graceful Shutdown

```go
func (p *Program) Quit()
```

Sends a `QuitMsg`. The event loop exits, final state is rendered, terminal is restored.

### Kill — Immediate Shutdown

```go
func (p *Program) Kill()
```

Stops immediately. Skips final render. Returns `ErrProgramKilled`. Use when you need to terminate without cleanup.

### Wait — Block Until Done

```go
func (p *Program) Wait()
```

Blocks on the `finished` channel until the program has fully shut down.

## Shutdown

### Graceful Quit

1. `eventLoop` returns `nil` error
2. `killed = false` → final `p.render(model)` is called
3. `p.shutdown(false)`:
   - Cancel internal context
   - Wait for all handler goroutines (`handlers.shutdown()`)
   - Cancel input reader, wait for read loop
   - Flush renderer (final frame)
   - Restore terminal state

### Kill

1. `p.shutdown(true)` called directly
2. Cancel internal context
3. Wait for handler goroutines
4. Cancel input reader (don't wait for read loop)
5. Stop renderer (skip flush)
6. Restore terminal state

### Interrupt

1. `eventLoop` returns `ErrInterrupted`
2. `killed = true` → no final render
3. `p.shutdown(true)` — same as kill path

## channelHandlers

```go
type channelHandlers struct {
    handlers []chan struct{}
    mu       sync.RWMutex
}
```

Tracks background goroutines. Each goroutine returns a `chan struct{}` that closes when done. `shutdown()` waits for all of them concurrently using a `sync.WaitGroup`.

## ReleaseTerminal / RestoreTerminal

### ReleaseTerminal

```go
func (p *Program) ReleaseTerminal() error
```

Restores terminal to normal mode, cancels input reader, stops renderer. Use when you need the terminal for something else (e.g., running an editor).

Internal steps (calls `releaseTerminal(false)` — the `false` means "don't reset renderer state", allowing resumption):
1. Set `ignoreSignals = 1` (suppress SIGINT/SIGTERM)
2. Cancel the cancelReader
3. Wait for read loop to finish (`waitForReadLoop`)
4. Stop the renderer (flush last frame, but don't reset)
5. Restore terminal state (`restoreTerminalState`)

### RestoreTerminal

```go
func (p *Program) RestoreTerminal() error
```

Re-initializes the terminal after `ReleaseTerminal`. Re-enters raw mode, restarts input reader and renderer, checks for resize.

Internal steps:
1. Set `ignoreSignals = 0` (re-enable signal handling)
2. `initTerminal()` — re-enter raw mode
3. `initInputReader(false)` — restart input parsing
4. Start renderer (`startRenderer`)
5. `checkResize()` in goroutine (terminal may have resized while released)
6. Repaint current view

### Exec Pattern

The `Exec`/`ExecProcess` commands use this cycle internally:

```
releaseTerminal(false) → run command → RestoreTerminal()
```

If the command fails, `RestoreTerminal` is still called. The event loop is **blocked** during this entire cycle.

## Embedding in Larger Applications

### Run in Goroutine, Control Externally

```go
errCh := make(chan error, 1)
go func() {
    _, err := p.Run()
    errCh <- err
}()
// Do other work...
p.Quit()  // or p.Kill()
err := <-errCh
```

### With Context

```go
ctx, cancel := context.WithCancel(context.Background())
p := tea.NewProgram(model, tea.WithContext(ctx))
go p.Run()
// Later: cancel() → p.Run() returns ErrProgramKilled
```

### Send Messages from Outside

```go
p := tea.NewProgram(model)
go func() {
    p.Send(customMsg{})
}()
p.Run()
```

### Without Renderer (Daemon Mode)

```go
p := tea.NewProgram(model, tea.WithoutRenderer())
// Update still runs, but nothing is rendered to the terminal
```

### Run Subprocess Mid-Program

```go
func openEditorCmd(filename string) tea.Cmd {
    cmd := exec.Command("vim", filename)
    return tea.ExecProcess(cmd, func(err error) tea.Msg {
        return editorFinishedMsg{err: err}
    })
}
```

Or manually with `ReleaseTerminal`/`RestoreTerminal`:

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case openEditorMsg:
        return m, func() tea.Msg {
            p.ReleaseTerminal()
            exec.Command("vim", m.filename).Run()
            p.RestoreTerminal()
            return editorDoneMsg{}
        }
    }
    return m, nil
}
```
