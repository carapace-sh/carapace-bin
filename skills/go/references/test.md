# Go Test — Testing, Benchmarks, Fuzzing, Coverage, and Profiling

Reference for `go test` — the command that automates testing Go packages, including test functions, benchmarks, fuzz tests, examples, coverage analysis, profiling, and test result caching.

> **Source of truth**: `go help test`, `go help testflag`, `go help testfunc`. For **build flags** shared with `go test`, see [build.md](build.md). For **vet integration**, see [tool.md](tool.md).

## Overview

```
go test [build/test flags] [packages] [build/test flags & test binary flags]
```

`go test` recompiles each package along with `*_test.go` files and runs the resulting test binary. Output format:

```
ok   archive/tar   0.011s
FAIL archive/zip   0.022s
ok   compress/gzip 0.033s
```

Followed by detailed output for each failed package.

### What Gets Compiled

- Files matching `*_test.go` are compiled alongside the package
- Files whose names begin with `_` or `.` are ignored (including `_test.go` — only `*_test.go` is recognized)
- Test files declaring a package with `_test` suffix are compiled as a separate package and linked with the test binary (external test package)
- Directories named `testdata` are ignored (available for test fixtures)

### Automatic `go vet`

`go test` runs `go vet` on the package and its test source files before running the test binary. A curated subset of vet checks is used:

| Check | Description |
|-------|-------------|
| `atomic` | Common misuse of sync/atomic |
| `bool` | Common mistakes with boolean operators |
| `buildtags` | Ill-formed or misplaced `//go:build` constraints |
| `directive` | Mistakes with `//go:` directives |
| `errorsas` | Invalid argument to `errors.As` |
| `ifaceassert` | Suspicious interface assertions |
| `nilfunc` | Comparison of function to nil |
| `printf` | Common mistakes with printf-like format strings |
| `stringintconv` | String conversion from non-integer types |
| `tests` | Common mistakes in test functions |

Disable with `-vet=off`. Run all checks with `-vet=all`. See [tool.md](tool.md) for vet details.

### Environment

The go command places `$GOROOT/bin` at the beginning of `$PATH` in the test's environment, so tests that execute `go` commands use the same `go` as the parent.

All test output and summary lines are printed to the go command's **stdout**, even if the test printed to its own stderr. The go command's stderr is reserved for build errors.

## Two Execution Modes

### Local Directory Mode

Triggered when `go test` is invoked with **no package arguments** (e.g., `go test` or `go test -v`):

- Compiles the package in the current directory
- Runs the test binary
- **Caching is disabled**
- Prints summary: `ok`/`FAIL`, package name, elapsed time

### Package List Mode

Triggered with **explicit package arguments** (e.g., `go test math`, `go test ./...`, `go test .`):

- Compiles and tests each listed package
- Passing packages: only the final `ok` summary line is printed
- Failing packages: full test output is printed
- With `-bench` or `-v`: full output even for passing packages
- **Caching is enabled** for successful results

## Test Caching

In package list mode, successful test results are cached. When a cached result matches, `go test` reprints the previous output with `(cached)` instead of the elapsed time.

### Cacheable Flags

A test result is cached only if all flags come from this restricted set:

| Flag | Description |
|------|-------------|
| `-benchtime` | Benchmark duration |
| `-coverprofile` | Coverage profile output |
| `-cpu` | GOMAXPROCS list |
| `-failfast` | Stop on first failure |
| `-fullpath` | Full file names in errors |
| `-list` | List matching tests |
| `-outputdir` | Output directory |
| `-parallel` | Parallel test count |
| `-run` | Test filter |
| `-short` | Short mode |
| `-skip` | Skip filter |
| `-timeout` | Timeout duration |
| `-v` | Verbose |

Any flag or argument outside this set disables caching. The idiomatic way to explicitly disable caching is `-count=1`.

### Cache Key Inputs

Tests that open files within the package's module or consult environment variables match future runs only if the files and env vars are unchanged.

### Cached Results and Timeout

A cached test result is treated as executing in no time, so a successful result is cached and reused regardless of the `-timeout` setting.

## `go test` Flags

These flags are handled by `go test` itself (not the test binary):

| Flag | Description |
|------|-------------|
| `-args` | Pass the remainder of the command line to the test binary, uninterpreted. The package list must appear before this flag. |
| `-c` | Compile the test binary to `pkg.test` in the current directory but do not run it. Output name/dir can be changed with `-o`. |
| `-exec xprog` | Run the test binary using `xprog`. Same behavior as `go run`. |
| `-json` | Convert test output to JSON (see [Test JSON Output](#test-json-output) below). Also emits build output in JSON. |
| `-o file` | Save a copy of the test binary to `file`. The test still runs unless `-c` is specified. If `file` ends with `/` or is a directory, the binary is written as `pkg.test` in that directory. |

## Test Binary Flags

These flags control execution of the test and are recognized by `go test` (passed to the test binary). See `go help testflag` for full details.

### Execution Control

| Flag | Description |
|------|-------------|
| `-run regexp` | Run only tests matching the regexp. Split by `/` for subtest hierarchy. Parents of matches are run too. |
| `-skip regexp` | Skip tests matching the regexp. Split by `/` for subtest hierarchy. |
| `-count n` | Run each test, benchmark, and fuzz seed n times (default 1). Examples always run once. Does not apply to `-fuzz`. |
| `-failfast` | Do not start new tests after the first failure |
| `-fullpath` | Show full file names in error messages |
| `-timeout d` | Panic if test binary runs longer than duration `d`. Default `10m`. `0` disables. |
| `-v` | Verbose: log all tests as they run. Print all Log/Logf output even on success. |
| `-short` | Tell long-running tests to shorten their run time |
| `-shuffle off,on,N` | Randomize execution order. `on` seeds from clock. Integer `N` uses `N` as seed. Seed is reported. |
| `-parallel n` | Max parallel tests calling `t.Parallel`. Default: `GOMAXPROCS`. Only applies within a single test binary. |

### Benchmarking

| Flag | Description |
|------|-------------|
| `-bench regexp` | Run benchmarks matching the regexp. Default: no benchmarks. `-bench .` runs all. Split by `/` for sub-benchmarks. |
| `-benchtime t` | Run enough iterations to take `t`. Default `1s`. Special `Nx` runs N times (e.g., `-benchtime 100x`). |
| `-benchmem` | Print memory allocation statistics for benchmarks. C `malloc` allocations not counted. |
| `-cpu 1,2,4` | List of GOMAXPROCS values to run tests/benchmarks with. Default: current GOMAXPROCS. Does not apply to `-fuzz`. |

### Fuzzing

| Flag | Description |
|------|-------------|
| `-fuzz regexp` | Run the fuzz test matching the regexp. Must match exactly one package in the main module and one fuzz test. Runs after tests, benchmarks, seed corpora, and examples. |
| `-fuzztime t` | Run fuzz target for duration `t`. Default: run forever. `Nx` runs N times. |
| `-fuzzminimizetime t` | Duration for each minimization attempt. Default `60s`. `Nx` runs N times. |

Fuzzing writes failures to `testdata/fuzz/`. The fuzzing engine caches inputs that expand code coverage in the build cache. `go clean -fuzzcache` removes these (may make fuzzing less effective temporarily).

### Listing

| Flag | Description |
|------|-------------|
| `-list regexp` | List tests, benchmarks, fuzz tests, or examples matching the regexp. No tests run. Only top-level tests listed (no subtests). |

### Output

| Flag | Description |
|------|-------------|
| `-json` | Log verbose output and test results as JSON (machine-readable). |
| `-outputdir directory` | Place profiling output and test artifacts in this directory. Default: current directory. |
| `-artifacts` | Save test artifacts in the directory specified by `-outputdir`. See `go doc testing.T.ArtifactDir`. |

### Vet Integration

| Flag | Description |
|------|-------------|
| `-vet list` | Configure `go vet` during `go test`. Empty: curated checks. `off`: no vet. Comma-separated list: specific checks. |

### Coverage

| Flag | Description |
|------|-------------|
| `-cover` | Enable coverage analysis |
| `-covermode set,count,atomic` | Coverage mode. Default `set` (or `atomic` with `-race`). |
| `-coverpkg pattern1,pattern2,...` | Apply coverage to packages matching patterns. Default: only the package being tested. |
| `-coverprofile cover.out` | Write coverage profile to file after all tests pass. Sets `-cover`. |

### Profiling

These flags generate profiles for use with `go tool pprof`. Each also leaves the test binary as `pkg.test` for profile analysis.

| Flag | Description |
|------|-------------|
| `-blockprofile block.out` | Write goroutine blocking profile |
| `-blockprofilerate n` | Sample one blocking event every `n` nanoseconds. Default (without flag): all events. |
| `-cpuprofile cpu.out` | Write CPU profile before exiting |
| `-memprofile mem.out` | Write allocation profile after all tests pass |
| `-memprofilerate n` | Set `runtime.MemProfileRate`. `1` profiles all allocations. |
| `-mutexprofile mutex.out` | Write mutex contention profile |
| `-mutexprofilefraction n` | Sample 1 in `n` stack traces of goroutines holding a contended mutex |
| `-trace trace.out` | Write execution trace before exiting |

## Test JSON Output

`go test -json` produces a newline-separated stream of `TestEvent` objects (see `go doc test2json`). Combined with `-json` on build commands, both `TestEvent` and `BuildEvent` objects are emitted. Parsers can distinguish them by inspecting the `Action` field.

See [build.md](build.md) for the `BuildEvent` struct definition.

## Testing Functions

`*_test.go` files can contain four kinds of functions (see `go help testfunc`):

### Test Functions

```go
func TestXxx(t *testing.T) { ... }
```

`Xxx` does not start with a lowercase letter. Uses `*testing.T` to report success/failure.

### Benchmark Functions

```go
func BenchmarkXxx(b *testing.B) { ... }
```

Run with `-bench`. The `b.N` field controls iteration count.

### Fuzz Tests

```go
func FuzzXxx(f *testing.F) { ... }
```

Run with `-fuzz`. Uses seed corpus from `testdata/fuzz/` and generates new inputs to find panics or failures.

### Example Functions

```go
func ExamplePrintln() {
    Println("The output of\nthis example.")
    // Output: The output of
    // this example.
}
```

- `Output:` comment → output compared exactly
- `Unordered output:` comment → output compared, order ignored
- No comment → compiled but not executed
- Empty `Output:` → compiled, executed, expected to produce no output

Example naming conventions:

| Name | Documents |
|------|-----------|
| `ExamplePrintln` | function `Println` |
| `ExampleT_M` | method `M` of type `T` |
| `ExamplePrintln_suffix` | additional example of `Println` (suffix doesn't start with uppercase) |

## Flag Passing to the Test Binary

`go test` rewrites recognized flags before invoking the test binary. Flags known to `go test` are rewritten with a `test.` prefix:

```bash
go test -v -myflag testdata -cpuprofile=prof.out -x
# runs as: pkg.test -test.v -myflag testdata -test.cpuprofile=prof.out
```

The `-x` flag is removed (applies only to the go command, not the test).

### Direct Test Binary Invocation

When running a generated test binary directly (`go test -c` output), the `test.` prefix is **mandatory**:

```bash
./pkg.test -test.v -test.run TestFoo
```

### `-args` for Uninterpreted Passing

```bash
go test -v -args -x -v
# runs as: pkg.test -test.v -x -v
```

The `-args` flag passes everything after it to the test binary unchanged.

## Working Directory

`go test` runs the test binary from within the corresponding package's source code directory. This may be in the module cache (read-only, checksum-verified). Tests must not write to the module directory unless explicitly requested (e.g., `-fuzz` writes to `testdata/fuzz`).

## Race Detector

The `-race` flag enables data race detection. When enabled:

- `-covermode` defaults to `atomic` (instead of `set`)
- `-installsuffix` gets `_race` appended
- Significant runtime overhead

For the list of supported platforms, see [build.md](build.md) (`-race` flag).

## Edge Cases and Gotchas

### Caching Disabled by Extra Flags

Any flag outside the cacheable set disables caching. This includes custom flags passed via `-args`. The idiomatic explicit cache-bust is `-count=1`.

### Package List Position

The package list must appear before any flag not known to `go test`:

```bash
go test -v -myflag testdata -cpuprofile=prof.out
# package list could go on either side of -v, but must be before -myflag
```

### Coverage Line Numbers

Coverage works by annotating source code before compilation. Compilation failures with coverage enabled may report line numbers that don't match the original source.

### `-parallel` vs `-p`

- `-parallel n` — max parallel tests within a single test binary (calling `t.Parallel`)
- `-p n` — max parallel test binaries across packages (build flag, default `GOMAXPROCS`)

### Fuzz Target Constraints

`-fuzz` must match exactly one package in the main module and exactly one fuzz test within that package. Fuzzing runs after all other tests, benchmarks, and examples.

## References

- `go help test` — `go test` command reference
- `go help testflag` — all test flags
- `go help testfunc` — testing function signatures
- `go doc test2json` — JSON test output encoding
- `go doc testing` — the testing package API
- <https://go.dev/doc/fuzz> — Go fuzzing guide

## Related

- [build.md](build.md) — shared build flags (`-race`, `-cover`, `-covermode`, `-coverpkg`)
- [tool.md](tool.md) — `go tool cover`, `go tool pprof`, vet checks
- [concepts.md](concepts.md) — test packages, `testdata`, `*_test.go` naming
- [environment.md](environment.md) — `GOCOVERDIR`, `GOTMPDIR`
