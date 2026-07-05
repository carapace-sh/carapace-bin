# Go CLI — Generate, Get, Install, Run, Clean, Doc, Fmt, Version, Telemetry

Reference for the remaining `go` subcommands: `go generate`, `go get`, `go install`, `go run`, `go clean`, `go doc`, `go fmt`, `go fix`, `go version`, `go telemetry`, `go bug`, and `go list`.

> **Source of truth**: `go help generate`, `go help get`, `go help install`, `go help run`, `go help clean`, `go help doc`, `go help fmt`, `go help fix`, `go help version`, `go help telemetry`, `go help bug`, `go help list`. For **build/test commands**, see [build.md](build.md) and [test.md](test.md). For **`go vet`/`go fix`** tools, see [tool.md](tool.md).

## `go generate`

```
go generate [-run regexp] [-skip regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]
```

Runs commands described by `//go:generate` directives within source files. These commands can run any process, but the intent is to create or update Go source files.

**Never run automatically** by `go build`, `go test`, etc. — must be invoked explicitly.

### Generate Directives

A directive is a line of the form:

```go
//go:generate command argument...
```

Rules:
- No leading spaces before `//go:generate`
- No space in `//go`
- `command` is an executable (in PATH, fully qualified path, or a `-command` alias)
- Arguments are space-separated tokens or double-quoted strings
- Quoted strings use Go syntax, evaluated before execution

```go
//go:generate stringer -type=Pill
//go:generate go tool cover -html=coverage.out
//go:generate mockgen -destination=mock.go -package=mocks example.com/foo Service
```

### Environment Variables Set During Generation

| Variable | Value |
|----------|-------|
| `$GOARCH` | Execution architecture (`arm`, `amd64`, etc.) |
| `$GOOS` | Execution OS (`linux`, `windows`, etc.) |
| `$GOFILE` | Base name of the file containing the directive |
| `$GOLINE` | Line number of the directive in the source file |
| `$GOPACKAGE` | Package name of the file containing the directive |
| `$GOROOT` | GOROOT directory of the invoking go command |
| `$DOLLAR` | A literal dollar sign (`$`) |
| `$PATH` | `$PATH` of the parent process, with `$GOROOT/bin` prepended |

### Variable Expansion

As a last step before running, all environment variable references (`$NAME`) are expanded throughout the command line — including inside quoted strings. Undefined variables expand to the empty string.

```go
//go:generate echo $GOFILE generates from line $GOLINE
```

### Command Aliases (`-command`)

A directive of the form:

```go
//go:generate -command xxx args...
```

Defines an alias `xxx` for the remainder of this source file only:

```go
//go:generate -command foo go tool foo
//go:generate foo -arg1 -arg2    # runs: go tool foo -arg1 -arg2
```

### Processing Order

1. Packages are processed in the order given on the command line, one at a time
2. Within a package, source files are processed in **file name order**, one at a time
3. Within a source file, generators run in the order they appear, one at a time
4. The generator runs in the package's source directory

If any generator returns a non-zero exit status, `go generate` skips all further processing for that package.

### Flags

| Flag | Description |
|------|-------------|
| `-run regexp` | Only run directives whose full source text matches the regexp |
| `-skip regexp` | Skip directives whose full source text matches the regexp (overrides `-run` if both match) |
| `-n` | Print commands that would be executed (dry run) |
| `-v` | Print names of packages and files as they are processed |
| `-x` | Print commands as they are executed |

Also accepts standard build flags.

### The `generate` Build Tag

`go generate` sets the build tag `generate` so that files may be examined by `go generate` but ignored during normal builds:

```go
//go:build generate
// +build generate

package main
```

### Generated File Convention

To convey that code is generated, generated source should have a line matching:

```
^// Code generated .* DO NOT EDIT\.$
```

This line must appear before the first non-comment, non-blank text in the file.

### Invalid Code Handling

For packages with invalid code, `go generate` processes only source files with a valid package clause.

## `go get`

```
go get [-t] [-u] [-tool] [build flags] [packages]
```

Resolves command-line arguments to packages at specific module versions, updates `go.mod` to require those versions, and downloads source code into the module cache.

**Note**: In earlier Go versions, `go get` was used to build and install packages. Now `go get` is dedicated to adjusting dependencies in `go.mod`. Use `go install` to build and install commands.

### Common Operations

```bash
go get example.com/pkg                    # add/upgrade to latest
go get example.com/pkg@v1.2.3             # specific version
go get example.com/pkg@none               # remove dependency
go get go@latest                          # upgrade minimum Go version
go get toolchain@patch                    # upgrade Go toolchain to latest patch
```

### Flags

| Flag | Description |
|------|-------------|
| `-t` | Consider modules needed to build tests of specified packages |
| `-u` | Update dependencies to newer minor or patch releases |
| `-u=patch` | Update dependencies, defaulting to patch releases |
| `-tool` | Add/remove a `tool` line in `go.mod` for each listed package |
| `-x` | Print commands as executed (useful for debugging VCS downloads) |

When `-t` and `-u` are used together, test dependencies are also updated.

### `-tool` Flag

`-tool` adds a `tool` directive to `go.mod` for each listed package:

```bash
go get -tool golang.org/x/tools/cmd/stringer    # adds tool directive
go get -tool golang.org/x/tools/cmd/stringer@none  # removes tool directive
```

See [mod.md](mod.md) for the `tool` directive and [tool.md](tool.md) for `go tool` usage.

## `go install`

```
go install [build flags] [packages]
```

Compiles and installs packages. Executables go to `GOBIN` (default: `$GOPATH/bin` or `$HOME/go/bin`). Executables in `$GOROOT` are installed to `$GOROOT/bin` or `$GOTOOLDIR` instead of `$GOBIN`. Cross-compiled binaries go to `$GOOS_$GOARCH` subdirectories.

### Module-Aware Install with Version Suffixes

```bash
go install example.com/pkg@latest       # install latest, ignoring local go.mod
go install example.com/pkg@v1.2.3       # install specific version
```

When version suffixes are used:
- Runs in module-aware mode, **ignoring** the `go.mod` in the current directory
- Useful for installing tools without affecting the main module's dependencies

### Constraints for Version-Suffixed Install

- Arguments must be package paths or patterns (with `...`), not stdlib packages, meta-patterns (`std`, `cmd`, `all`), or relative/absolute file paths
- All arguments must have the **same version suffix**
- All arguments must refer to packages in the **same module** at the same version
- Package path arguments must refer to **main** packages; patterns only match main packages
- No module is the "main" module; the target module's `go.mod` must not have `replace`/`exclude` that would change interpretation
- Vendor directories are not used

### Without Version Suffixes

Without version suffixes, `go install` runs in the context of the main module (module-aware mode) or GOPATH mode, depending on `GO111MODULE` and `go.mod` presence.

Non-main packages are:
- **Module-aware mode**: built and cached, not installed
- **GOPATH mode**: installed to `$GOPATH/pkg/$GOOS_$GOARCH`

### Standard Library

Since Go 1.20, the standard library is built and cached but not installed to `$GOROOT/pkg/`. Set `GODEBUG=installgoroot=all` to restore the old behavior.

## `go run`

```
go run [build flags] [-exec xprog] package [arguments...]
```

Compiles and runs the named main Go package. The package can be specified as:
- A list of `.go` source files from a single directory
- An import path
- A file system path
- A pattern matching a single package (`go run .` or `go run my/cmd`)

### Version Suffix

```bash
go run example.com/cmd@latest     # run from remote module, ignoring local go.mod
go run example.com/cmd@v1.2.3     # run specific version
```

### Execution

- By default, runs the compiled binary directly: `a.out arguments...`
- With `-exec xprog`: invokes `xprog a.out arguments...`
- Without `-exec`, if `GOOS`/`GOARCH` differ from the system default and a program named `go_$GOOS_$GOARCH_exec` is on `PATH`, uses that (e.g., `go_js_wasm_exec` for WASM)
- Compiles **without** debugger info (to reduce build time). Use `go build` for debuggable binaries.
- Exit status of `go run` is **not** the exit status of the compiled binary

## `go clean`

```
go clean [-i] [-r] [-cache] [-testcache] [-modcache] [-fuzzcache] [build flags] [packages]
```

Removes object files from package source directories. The go command builds most objects in a temporary directory, so `go clean` is mainly for files left by other tools or manual `go build` invocations.

### Files Removed Per Package

When a package argument is given or `-i`/`-r` is set:

| File Pattern | Source |
|--------------|--------|
| `_obj/` | Old Makefile object directory |
| `_test/` | Old Makefile test directory |
| `_testmain.go` | Old gotest file |
| `test.out` | Old test log |
| `build.out` | Old test log |
| `*.[568ao]` | Object files |
| `DIR(.exe)` | From `go build` |
| `DIR.test(.exe)` | From `go test -c` |
| `MAINFILE(.exe)` | From `go build MAINFILE.go` |
| `*.so` | From SWIG |

`DIR` = final path element; `MAINFILE` = base name of a Go source file not included in the package build.

### Flags

| Flag | Description |
|------|-------------|
| `-i` | Also remove installed archive or binary (what `go install` would create) |
| `-n` | Print remove commands but don't execute |
| `-r` | Apply recursively to all dependencies |
| `-x` | Print remove commands as executed |
| `-cache` | Remove the **entire** go build cache |
| `-testcache` | Expire all test results in the build cache (keeps cached builds) |
| `-modcache` | Remove the **entire** module download cache (including unpacked source) |
| `-fuzzcache` | Remove cached fuzz testing files from the build cache |

## `go doc`

```
go doc [doc flags] [package|[package.]symbol[.methodOrField]]
```

Prints documentation comments for a package, const, func, type, var, method, or struct field, followed by a one-line summary of first-level items.

### Usage Forms

```bash
go doc                              # current package
go doc Foo                          # symbol Foo in current package (capital = symbol, not package)
go doc encoding/json                # package docs
go doc json                         # shorthand for encoding/json
go doc json.Number                  # type documentation
go doc json.Number.Int64            # method documentation
go doc json.Decoder.Decode          # method (multiple forms work)
go doc text/template new            # two-argument form: package + symbol
```

### Matching Rules

- Lowercase letters match either case; uppercase letters match exactly
- Package path must be a qualified path or proper suffix (not `.` or `...`)
- Packages are scanned in breadth-first lexical order; GOROOT before GOPATH
- If no package specified, current directory's package is used

### Flags

| Flag | Description |
|------|-------------|
| `-all` | Show all documentation for the package |
| `-c` | Respect case when matching symbols |
| `-cmd` | Treat `package main` like a regular package (show exported symbols) |
| `-http` | Serve HTML documentation over HTTP |
| `-short` | One-line representation for each symbol |
| `-src` | Show full source code for the symbol |
| `-u` | Show unexported symbols, methods, and fields |

## `go fmt`

```
go fmt [-n] [-x] [packages]
```

Runs `gofmt -l -w` on the named packages. Prints names of modified files.

| Flag | Description |
|------|-------------|
| `-n` | Print commands that would be executed |
| `-x` | Print commands as executed |
| `-mod` | Set module download mode (`readonly` or `vendor`) |

To run `gofmt` with specific options, run `gofmt` directly:

```bash
gofmt -d -s .        # show simplification diffs
gofmt -w .           # write changes
```

## `go version`

```
go version [-m] [-v] [-json] [file ...]
```

Prints build information for Go binary files. With no files, prints its own version info.

| Flag | Description |
|------|-------------|
| `-m` | Print embedded module version information |
| `-v` | Report unrecognized files during directory scan |
| `-json` | Output `runtime/debug.BuildInfo` as JSON (requires `-m`) |

If a directory is named, walks it recursively looking for Go binaries. Without `-v`, unrecognized files are not reported.

```bash
go version                    # go version go1.23.4 linux/amd64
go version -m ./myapp         # show version + embedded module info
go version -json -m ./myapp   # JSON format
```

## `go telemetry`

```
go telemetry [off|local|on]
```

Manages Go telemetry data and settings. Three modes:

| Mode | Local Collection | Upload |
|------|-----------------|--------|
| `off` | No | No |
| `local` | Yes | No |
| `on` | Yes | Yes (to telemetry.go.dev) |

```bash
go telemetry           # print current mode
go telemetry off       # disable all collection
go telemetry local     # collect locally, don't upload
go telemetry on        # collect and upload
```

The current mode is available as `go env GOTELEMETRY`. The data directory is `go env GOTELEMETRYDIR`. See <https://go.dev/doc/telemetry> for details.

## `go bug`

```
go bug
```

Opens the default browser and starts a new bug report with useful system information.

## `go list`

```
go list [-f format] [-json] [-m] [list flags] [build flags] [packages]
```

Lists packages, one per line. The `-f` flag controls output format using `text/template` syntax. The struct passed to the template is a `Package` struct (documented in `go help list`) with fields for Dir, ImportPath, Name, GoFiles, Deps, Module, etc.

### Common Usage

```bash
go list ./...                    # all package import paths
go list -f '{{.ImportPath}} {{.Name}}' ./...   # custom format
go list -json ./...              # JSON output
go list -m all                   # list modules (not packages)
go list -m -json all             # modules as JSON
```

### `-f` Template

The default output is equivalent to `-f '{{.ImportPath}}'`. Key Package fields:

| Field | Description |
|-------|-------------|
| `.ImportPath` | Import path of the package |
| `.Name` | Package name |
| `.Dir` | Directory containing package sources |
| `.GoFiles` | `.go` source files (excluding Cgo, Test, XTest) |
| `.Deps` | Import paths of dependencies |
| `.Module` | Info about containing module (can be nil) |
| `.Standard` | Whether part of the standard library |
| `.Stale` | Whether `go install` would do anything |

### List Flags

| Flag | Description |
|------|-------------|
| `-f format` | Output format (template syntax) |
| `-json` | JSON output |
| `-m` | List modules instead of packages |
| `-compiled` | Include `CompiledGoFiles` in output |
| `-deps` | Include dependencies (not just matched packages) |
| `-e` | Tolerate erroneous packages (set `Error` field instead of failing) |
| `-export` | Include `Export` and `BuildID` fields |
| `-find` | Identify packages but don't load dependencies or resolve imports |
| `-test` | Report test binaries too (adds `.test` suffix to package import paths) |
| `-u` | Add available upgrade information (with `-m`) |
| `-versions` | Set `Versions` field to all known module versions (with `-m`) |
| `-retracted` | Report retracted version information (with `-m`) |

For the full Package struct and all list flags, run `go help list`.

## Edge Cases and Gotchas

### `go generate` Doesn't Parse Files

`go generate` scans for directives line-by-line — it doesn't parse the file. Lines that look like directives in comments or multiline strings will be treated as directives:

```go
/*
//go:generate echo this runs even inside a comment block
*/
```

### `go get` vs `go install`

- `go get` — adjusts `go.mod` dependencies (no longer builds/installs)
- `go install` — builds and installs (can use `@version` to ignore local `go.mod`)

### `go run` Exit Status

The exit status of `go run` is **not** the exit status of the compiled binary. If the binary exits with code 42, `go run` may still exit with 0 (or vice versa for build errors).

### `go run` No Debug Info

`go run` compiles without debugger information to reduce build time. For debugging, use `go build` then run the binary directly.

### `go clean -modcache` Is Destructive

`go clean -modcache` removes the **entire** module download cache, including unpacked source. All modules will need to be re-downloaded. Use `-testcache` for a less destructive cache clear.

### `go fmt` Runs `gofmt -l -w`

`go fmt` runs `gofmt -l -w` (list modified files, write changes). For other gofmt options (like `-s` for simplification or `-d` for diff), run `gofmt` directly.

## References

- `go help generate` — generate directive reference
- `go help get` — dependency management
- `go help install` — install command
- `go help run` — run command
- `go help clean` — clean command
- `go help doc` — doc command
- `go help fmt` — fmt command
- `go help fix` — fix command
- `go help version` — version command
- `go help telemetry` — telemetry command
- `go help bug` — bug command
- `go help list` — list command and Package struct
- `go doc cmd/gofmt` — gofmt documentation

## Related

- [build.md](build.md) — shared build flags, build modes, cross-compilation
- [test.md](test.md) — `go test` command and flags
- [mod.md](mod.md) — `go mod` subcommands, `go.mod` format, `go get` details
- [tool.md](tool.md) — `go tool`, `go vet`, `go fix`
- [environment.md](environment.md) — `GOBIN`, `GOPATH`, `GOTELEMETRY`, `GOTELEMETRYDIR`
