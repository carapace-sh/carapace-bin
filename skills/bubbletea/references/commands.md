# Commands — Async I/O and Side Effects

In-depth reference for Bubble Tea's command system: `Cmd`, `Batch`, `Sequence`, built-in commands, and how commands are dispatched.

## The Cmd Type

```go
type Cmd func() Msg
```

A `Cmd` is a function that takes no arguments and returns a `Msg`. It represents an asynchronous side effect. When `nil`, it's a no-op.

**Key principle**: Commands are the **only** way to perform I/O in Bubble Tea. `Update` must be a pure function — all side effects go through returned commands.

## Built-in Commands

### Quit, Suspend, Interrupt

```go
func Quit() Msg       { return QuitMsg{} }
func Suspend() Msg    { return SuspendMsg{} }
func Interrupt() Msg  { return InterruptMsg{} }
```

These are synchronous — they return a `Msg` directly (not a `Cmd`). Use them as return values from `Update`:

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyPressMsg:
        if msg.String() == "ctrl+c" {
            return m, tea.Interrupt
        }
        if msg.String() == "q" {
            return m, tea.Quit
        }
    }
    return m, nil
}
```

### Tick — Single Timer

```go
func Tick(d time.Duration, fn func(time.Time) Msg) Cmd
```

Fires once after `d` duration. The callback receives the actual time and returns a message.

```go
// Flash a message for 2 seconds, then clear it
func flashCmd() tea.Cmd {
    return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
        return clearMsg{}
    })
}
```

Implementation: creates a `time.NewTimer`, blocks on its channel, drains extras.

### Every — Repeating Timer

```go
func Every(duration time.Duration, fn func(time.Time) Msg) Cmd
```

Fires once, aligned to the system clock (truncated to duration boundary). **Not truly periodic** — you must re-issue the command in `Update` to repeat:

```go
func (m model) Init() tea.Cmd {
    return tea.Every(time.Second, tickMsg)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tickMsg:
        m.count++
        return m, tea.Every(time.Second, tickMsg)  // re-issue
    }
    return m, nil
}
```

Implementation: calculates the delay to the next aligned boundary, creates a `time.NewTimer`.

### RequestWindowSize

```go
func RequestWindowSize() Msg { return windowSizeMsg{} }
```

Triggers a `WindowSizeMsg` to be sent. Useful when you need the current terminal size outside of the normal resize handler.

### Println / Printf

```go
func Println(args ...any) Cmd
func Printf(template string, args ...any) Cmd
```

Prints a line **above** the program (in the terminal's scrollback). Useful for status messages that should persist after the program exits.

```go
return m, tea.Println("Download complete:", filename)
```

### Exec / ExecProcess — Run Subprocesses

```go
type ExecCommand interface {
    Run() error
    SetStdin(io.Reader)
    SetStdout(io.Writer)
    SetStderr(io.Writer)
}

type ExecCallback func(error) Msg

func Exec(c ExecCommand, fn ExecCallback) Cmd
func ExecProcess(c *exec.Cmd, fn ExecCallback) Cmd
```

Runs an external command. The program **releases the terminal** (restores normal mode), runs the command, then **restores the terminal**. The callback receives the error result.

**This blocks the event loop** — no other messages are processed while the command runs.

```go
func openEditorCmd(filename string) tea.Cmd {
    cmd := exec.Command("vim", filename)
    return tea.ExecProcess(cmd, func(err error) tea.Msg {
        return editorFinishedMsg{err: err}
    })
}
```

The `ExecCommand` interface allows custom implementations for testing.

## Batch — Concurrent Commands

```go
func Batch(cmds ...Cmd) Cmd
type BatchMsg []Cmd
```

Runs all commands **concurrently** in separate goroutines. All goroutines are waited on before the batch completes. Each resulting message is sent to the event loop independently.

```go
func (m model) Init() tea.Cmd {
    return tea.Batch(
        fetchUsersCmd(),
        fetchReposCmd(),
        tea.Every(time.Second, tickMsg),
    )
}
```

**Internal behavior** (`execBatchMsg`):
1. Launches each cmd in its own goroutine
2. Each goroutine calls `cmd()` to get a `Msg`
3. If the result is a `BatchMsg` or `sequenceMsg`, it's handled recursively
4. Otherwise, the message is sent via `p.Send(msg)`
5. `sync.WaitGroup` ensures all goroutines complete

**Ordering**: Messages from concurrent commands arrive in **non-deterministic order**. If ordering matters, use `Sequence`.

## Sequence — Sequential Commands

```go
func Sequence(cmds ...Cmd) Cmd
// sequenceMsg is unexported
```

Runs commands **sequentially** — each cmd blocks until complete before the next starts. No goroutines are spawned for individual commands.

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case saveCompleteMsg:
        return m, tea.Sequence(
            tea.Println("Saved!"),
            tea.Quit,
        )
    }
    return m, nil
}
```

**Internal behavior** (`execSequenceMsg`):
1. Iterates over commands in order
2. Calls `cmd()` synchronously
3. If the result is a `BatchMsg` or `sequenceMsg`, it's handled recursively
4. Otherwise, the message is sent via `p.Send(msg)`

## compactCmds — Optimization

Both `Batch` and `Sequence` use `compactCmds[T ~[]Cmd]()`:

- Filters out `nil` commands
- Returns `nil` if no valid commands remain (no-op)
- Returns the single command directly if only one remains (avoids wrapping)
- Returns `T(validCmds)` only if 2+ commands remain

This means `tea.Batch(singleCmd)` is equivalent to just `singleCmd`, and `tea.Batch(nil, nil)` is a no-op.

## Command Dispatch Flow

```
Update(msg) → (Model, Cmd)
                ↓
         cmds channel
                ↓
      handleCommands goroutine
                ↓
         go cmd()        ← each cmd runs in its own goroutine
                ↓
         p.Send(msg)     ← result sent back to event loop
                ↓
      eventLoop processes msg
                ↓
      Update(msg) → (Model, Cmd)  ← cycle repeats
```

**Special dispatch in event loop** (before reaching `Update`):

| Message Type | Action |
|---|---|
| `BatchMsg` | `go p.execBatchMsg(msg)` — concurrent execution |
| `sequenceMsg` | `go p.execSequenceMsg(msg)` — sequential execution |
| `execMsg` | `p.exec(msg.cmd, msg.fn)` — **blocks event loop** |
| `printLineMessage` | `p.renderer.insertAbove(msg.messageBody)` — prints above program |
| `QuitMsg` | Returns from event loop |
| `InterruptMsg` | Returns `ErrInterrupted` |
| `SuspendMsg` | `p.suspend()` — suspends process |

## Custom Commands

Any function with signature `func() tea.Msg` is a valid command:

```go
// HTTP request
func fetchUserCmd(id int) tea.Cmd {
    return func() tea.Msg {
        resp, err := http.Get(fmt.Sprintf("/api/users/%d", id))
        if err != nil {
            return errMsg{err}
        }
        defer resp.Body.Close()
        var user User
        json.NewDecoder(resp.Body).Decode(&user)
        return userLoadedMsg{user}
    }
}

// File read
func readFileCmd(path string) tea.Cmd {
    return func() tea.Msg {
        data, err := os.ReadFile(path)
        if err != nil {
            return errMsg{err}
        }
        return fileLoadedMsg{data: data}
    }
}
```

## Common Patterns

### Multiple Commands from Update

```go
var cmds []tea.Cmd
m.spinner, cmd = m.spinner.Update(msg)
cmds = append(cmds, cmd)
m.viewport, cmd = m.viewport.Update(msg)
cmds = append(cmds, cmd)
return m, tea.Batch(cmds...)
```

### Conditional Command

```go
var cmd tea.Cmd
if m.needsRefresh {
    cmd = refreshCmd()
}
return m, cmd
```

### Command Chain (Sequence)

```go
return m, tea.Sequence(
    saveDataCmd(),
    tea.Println("Data saved"),
    tea.Quit,
)
```

### No-Op Command

```go
return m, nil  // nil Cmd is a no-op
```
