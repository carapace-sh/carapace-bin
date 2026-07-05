# Go Tool — Bundled Tools and Vet

Reference for `go tool` — the command that runs Go's built-in toolchain tools (asm, cgo, compile, cover, fix, link, preprofile, vet) and tools defined in `go.mod`. Covers `go vet` and `go fix` as driver commands.

> **Source of truth**: `go help tool`, `go help vet`, `go help fix`, `go doc cmd/vet`, `go tool vet help`. For **build flags** passed through to tools, see [build.md](build.md).

## `go tool` Command

```
go tool [-n] command [args...]
```

Runs the go tool command identified by the arguments. With no arguments, prints the list of known tools.

### Flags

| Flag | Description |
|------|-------------|
| `-n` | Print the command that would be executed but do not execute it |
| `-modfile=file.mod` | Use an alternate go.mod file |
| `-C dir` | Change to directory before running |
| `-overlay file` | Use overlay config for build operations |
| `-modcacherw` | Leave module cache directories read-write |

### Tool Resolution

Go ships with **built-in tools**. Additional tools may be defined in the current module's `go.mod` via `tool` directives. When run, `go tool <name>` resolves `<name>` in this order:

1. **Built-in tools** (see table below)
2. **Module-defined tools** (`tool` directives in `go.mod`)

For details on each built-in tool, run `go doc cmd/<command>`.

## Built-in Tools

| Tool | Description | Common Usage |
|------|-------------|--------------|
| `asm` | Go assembler | Assembles `.s` files; invoked by `go build` |
| `cgo` | C interop generator | Generates Go/C glue code; invoked when package imports `"C"` |
| `compile` | Go compiler | Compiles `.go` files to object files; invoked by `go build` |
| `cover` | Coverage tool | Formats coverage profiles; see [Coverage](#coverage-go-tool-cover) |
| `fix` | Code fixer | Applies fixes for API changes; see [go fix](#go-fix) |
| `link` | Go linker | Links object files into executables; invoked by `go build` |
| `preprofile` | Pre-optimization profiler | Used internally for PGO |
| `vet` | Static analyzer | Reports likely mistakes; see [go vet](#go-vet) |

### Direct Tool Invocation

Tools are normally invoked automatically by `go build`, `go test`, etc. Direct invocation is useful for debugging or non-standard build setups:

```bash
go tool compile -S foo.go    # compile and print assembly
go tool link -s foo.o        # link with symbols stripped
go tool asm -o foo.o foo.s   # assemble
```

### The `-toolexec` Flag

Build commands accept `-toolexec 'cmd args'` to wrap tool invocations:

```bash
go build -toolexec 'my-wrapper' ./...
```

The wrapper receives the tool path and its arguments. The `TOOLEXEC_IMPORTPATH` environment variable is set, matching `go list -f {{.ImportPath}}` for the package being built.

## `go vet`

```
go vet [build flags] [-vettool prog] [vet flags] [packages]
```

Runs `go tool vet` (cmd/vet) on the named packages and reports diagnostics. `go test` also runs vet automatically (a curated subset).

### Flags

| Flag | Description |
|------|-------------|
| `-c int` | Display offending line with this many lines of context (default -1) |
| `-json` | Emit JSON output |
| `-fix` | Apply each diagnostic's first fix (if any) instead of printing |
| `-diff` | Print fixes as unified diff instead of applying; non-zero exit if non-empty |
| `-vettool prog` | Use an alternative analysis tool |

### Supported Build Flags

`go vet` supports build flags that control package resolution and execution: `-C`, `-n`, `-x`, `-v`, `-tags`, `-toolexec`. See [build.md](build.md).

### The Default Vet Tool

The default is `go tool vet` (cmd/vet). For help on checkers and their flags:

```bash
go tool vet help          # list all checkers
go tool vet help printf   # details on the printf checker
```

### Default Vet Checks (used by `go test`)

When `go test` runs vet automatically, it uses this curated subset:

| Check | Description |
|-------|-------------|
| `atomic` | Common misuse of `sync/atomic` |
| `bool` | Common mistakes with boolean operators |
| `buildtags` | Ill-formed or misplaced `//go:build` constraints |
| `directive` | Mistakes with `//go:` directives |
| `errorsas` | Invalid argument to `errors.As` |
| `ifaceassert` | Suspicious interface assertions |
| `nilfunc` | Comparison of function to nil |
| `printf` | Common mistakes with printf-like format strings |
| `stringintconv` | String conversion from non-integer types |
| `tests` | Common mistakes in test functions |

Control via `-vet` flag on `go test`:
- Empty list → curated checks (default)
- `off` → no vet
- Comma-separated list → specific checks
- `all` → all checks

### Alternative Vet Tools

The `-vettool` flag selects a different analysis tool. Alternative tools should be built atop `golang.org/x/tools/go/analysis/unitchecker`:

```bash
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
go vet -vettool=$(which shadow) ./...
```

### `go vet` vs `go tool vet`

`go vet` is the driver command that resolves packages and invokes the vet tool. `go tool vet` is the underlying tool. `go vet` adds package resolution, build flag support, and the `-vettool` mechanism.

## `go fix`

```
go fix [build flags] [-fixtool prog] [fix flags] [packages]
```

Runs `go tool fix` (cmd/fix) on the named packages and applies suggested fixes for API changes between Go versions.

### Flags

| Flag | Description |
|------|-------------|
| `-diff` | Print fixes as unified diff instead of applying; non-zero exit if non-empty |
| `-fixtool prog` | Use an alternative fixer tool |

### The Default Fix Tool

The default is `go tool fix` (cmd/fix). For help on fixers:

```bash
go tool fix help           # list all fixers
go tool fix help hostport  # details on the hostport fixer
```

### Supported Build Flags

`go fix` supports: `-C`, `-n`, `-x`, `-v`, `-tags`, `-toolexec`. See [build.md](build.md).

## Coverage (`go tool cover`)

`go tool cover` formats coverage profiles generated by `go test -coverprofile`:

```bash
# Generate coverage profile
go test -coverprofile=coverage.out ./...

# View text summary
go tool cover -func=coverage.out

# Generate HTML report
go tool cover -html=coverage.out

# Output to a specific file
go tool cover -html=coverage.out -o coverage.html
```

For `go build -cover` binaries (not test-based coverage), use `GOCOVERDIR` to collect coverage data at runtime. See [environment.md](environment.md).

## Profiling (`go tool pprof`)

`go tool pprof` analyzes profiles generated by `go test` profiling flags or by `runtime/pprof`:

```bash
# CPU profile from test
go test -cpuprofile=cpu.out ./...
go tool pprof cpu.out

# Interactive top view
go tool pprof -top cpu.out

# Web interface
go tool pprof -http=:8080 cpu.out

# Memory profile
go test -memprofile=mem.out ./...
go tool pprof -alloc_objects mem.out
```

Key pprof options:
- `-sample_index=alloc_space` — show allocation by space
- `-sample_index=alloc_objects` — show allocation by object count
- `-show_bytes` — show bytes instead of objects

For full pprof usage, run `go tool pprof -h` or see the [pprof docs](https://pkg.go.dev/runtime/pprof).

## Module-Defined Tools

The `tool` directive in `go.mod` declares tool dependencies that can be run with `go tool`:

```go
// go.mod
tool golang.org/x/tools/cmd/goimports
tool golang.org/x/tools/cmd/stringer
```

```bash
go tool stringer -type=Pill     # run the stringer tool
go tool goimports -l .          # run goimports
```

### Managing Tool Dependencies

| Action | Command |
|--------|---------|
| Add a tool | `go get -tool golang.org/x/tools/cmd/stringer` |
| Remove a tool | `go get -tool golang.org/x/tools/cmd/stringer@none` |
| List available tools | `go tool` (with no arguments) |

The `tool` pattern in package lists expands to tools defined in the current module's `go.mod`:

```bash
go list tool     # list tool packages
```

## Edge Cases and Gotchas

### `go vet` is Run by `go test`

`go test` automatically runs a subset of vet checks. If vet finds problems, the test binary is not run. Use `-vet=off` to disable, or `-vet=all` for all checks.

### `-fix` Modifies Source

`go vet -fix` and `go fix` modify source files in place. Use `-diff` first to preview changes.

### Tools Are Platform-Specific

Built-in tools like `compile` and `link` are specific to the target platform (`GOOS`/`GOARCH`). Cross-compilation uses the target platform's tools.

### `go tool` With No Arguments Lists Tools

Running `go tool` with no arguments prints all known tools (built-in + module-defined). This is useful for discovering what's available.

### TOOLEXEC_IMPORTPATH

When using `-toolexec`, the wrapper can read `TOOLEXEC_IMPORTPATH` to know which package is being built. This matches `go list -f {{.ImportPath}}`.

## References

- `go help tool` — `go tool` command reference
- `go help vet` — `go vet` command reference
- `go help fix` — `go fix` command reference
- `go tool vet help` — vet checker list and details
- `go tool fix help` — fixer list and details
- `go doc cmd/vet` — vet documentation
- `go doc cmd/fix` — fix documentation
- `go doc cmd/cgo` — cgo documentation
- `go doc cmd/compile` — compiler documentation
- `go doc cmd/link` — linker documentation

## Related

- [build.md](build.md) — `-toolexec`, `-gcflags`, `-ldflags`, `-asmflags` (flags passed to tools)
- [test.md](test.md) — `go test` vet integration, coverage profiles, profiling flags
- [mod.md](mod.md) — `tool` directive in `go.mod`
- [environment.md](environment.md) — `GOTOOLDIR`, `GOCOVERDIR`
