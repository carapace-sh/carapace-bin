# Go Build — Compilation, Flags, and Build Constraints

Reference for `go build`, `go install`, and the shared build flags used across multiple commands. Covers build constraints (tags), build modes, cross-compilation, profile-guided optimization, and JSON output.

> **Source of truth**: `go help build`, `go help buildconstraint`, `go help buildmode`, `go help buildjson`. For **build cache internals**, see [environment.md](environment.md) (`GOCACHE`). For **package patterns**, see [concepts.md](concepts.md).

## `go build`

```
go build [-o output] [build flags] [packages]
```

Compiles packages and their dependencies, but does **not** install the results. Key behaviors:

- Compiling a single `main` package → writes executable named after the last non-major-version component of the import path
- Compiling from `.go` files → executable named after the first source file
- Compiling multiple packages or a non-main package → compiles but discards output (syntax/type check)
- Ignores `*_test.go` files
- `.exe` suffix added automatically on Windows

The `-o` flag forces a specific output file or directory:

```bash
go build -o myapp              # name the output binary
go build -o bin/               # write executables to bin/
```

### See Also

- `go install` — compile and install (see [cli.md](cli.md))
- `go run` — compile and run (see [cli.md](cli.md))
- `go clean` — remove build artifacts (see [cli.md](cli.md))

## Shared Build Flags

These flags are shared by `go build`, `go clean`, `go get`, `go install`, `go list`, `go run`, and `go test`:

### General Flags

| Flag | Description |
|------|-------------|
| `-C dir` | Change to `dir` before running. **Must be the first flag.** Files named on the command line are interpreted after changing directories. |
| `-a` | Force rebuilding of packages that are already up-to-date |
| `-n` | Print commands but do not run them |
| `-p n` | Number of programs (build commands, test binaries) to run in parallel. Default is `GOMAXPROCS`. |
| `-v` | Print names of packages as they are compiled |
| `-work` | Print the temporary work directory name and do not delete it on exit |
| `-x` | Print the commands being executed |

### Sanitizer and Race Detection

| Flag | Description |
|------|-------------|
| `-race` | Enable data race detector. Supported: `darwin/amd64`, `darwin/arm64`, `freebsd/amd64`, `linux/amd64`, `linux/arm64` (48-bit VMA), `linux/ppc64le`, `linux/riscv64`, `windows/amd64`. |
| `-msan` | Enable memory sanitizer interoperation. Supported: `linux/amd64`, `linux/arm64`, `linux/loong64`, `freebsd/amd64` (Clang/LLVM only). PIE build mode used except on `linux/amd64`. |
| `-asan` | Enable address sanitizer interoperation. Supported: `linux/arm64`, `linux/amd64`, `linux/loong64`. Requires GCC 7+ or Clang/LLVM 9+ (16+ for loong64). |

### Coverage

| Flag | Description |
|------|-------------|
| `-cover` | Enable coverage instrumentation |
| `-covermode set,count,atomic` | Coverage mode. Default `set` (or `atomic` with `-race`). `set`: bool, `count`: int, `atomic`: int, thread-safe. |
| `-coverpkg pattern1,pattern2,...` | Apply coverage to packages matching patterns. Default: packages in the main module. |

### Compiler and Linker Flags

| Flag | Description |
|------|-------------|
| `-asmflags '[pattern=]arg list'` | Arguments for each `go tool asm` invocation |
| `-compiler name` | Compiler to use: `gc` (default) or `gccgo` |
| `-gccgoflags '[pattern=]arg list'` | Arguments for each gccgo compiler/linker invocation |
| `-gcflags '[pattern=]arg list'` | Arguments for each `go tool compile` invocation |
| `-ldflags '[pattern=]arg list'` | Arguments for each `go tool link` invocation |

The `-asmflags`, `-gccgoflags`, `-gcflags`, and `-ldflags` flags accept a space-separated list. A package pattern and `=` can prefix the list to restrict it to matching packages. The latest match on the command line wins:

```bash
go build -gcflags=-S fmt                    # disassembly only for fmt
go build -gcflags=all=-S fmt                # disassembly for fmt and all dependencies
```

To embed spaces in an element, surround it with single or double quotes.

### Module and Dependency Flags

| Flag | Description |
|------|-------------|
| `-mod mode` | Module download mode: `readonly`, `vendor`, or `mod`. Default: `vendor` if vendor dir present and go ≥1.14, else `readonly`. |
| `-modcacherw` | Leave newly-created module cache directories read-write (instead of read-only) |
| `-modfile file` | Use alternate `go.mod` (a `go.mod` must still exist to determine module root). Derived `go.sum` path: trim `.mod`, append `.sum`. |
| `-linkshared` | Build against shared libraries created with `-buildmode=shared` |
| `-installsuffix suffix` | Suffix for package installation directory. Auto-appends `_race`, `_msan`, `_asan` for respective flags. |

### Output and Path Flags

| Flag | Description |
|------|-------------|
| `-json` | Emit build output as JSON (see [Build JSON Output](#build-json-output) below) |
| `-overlay file` | Read a JSON overlay config mapping disk file paths to backing file paths |
| `-trimpath` | Remove file system paths from the executable. Paths become `module@version` or plain import path. |
| `-pkgdir dir` | Install and load all packages from `dir` instead of usual locations |
| `-toolexec 'cmd args'` | Program to invoke toolchain programs (asm, vet, etc.). Sets `TOOLEXEC_IMPORTPATH` env var. |

### Version Control and Tags

| Flag | Description |
|------|-------------|
| `-buildvcs` | Stamp binary with VCS info (`true`, `false`, or `auto`). Default `auto`: stamps if main package, main module, and current directory are in the same repo. |
| `-tags tag,list` | Comma-separated build tags to consider satisfied. (Older space-separated form is deprecated but recognized.) |
| `-pgo file` | Profile-guided optimization profile. `auto` (default): uses `default.pgo` in package dir. `off`: disable PGO. |

## Build Constraints (Build Tags)

A **build constraint** (or build tag) is a condition under which a file should be included in the package. It appears as a line comment:

```go
//go:build (linux && 386) || (darwin && !cgo)
```

### Syntax and Placement

- Must begin with `//go:build`
- Must appear near the top of the file, preceded only by blank lines and other comments
- In Go files, must appear **before** the package clause
- Should be followed by a blank line (to distinguish from package documentation)
- A file may have only **one** `//go:build` line
- Expressions use `||`, `&&`, `!`, and parentheses (same meaning as Go)

### Satisfied Tags During a Build

| Tag Category | Examples |
|--------------|----------|
| Target OS | `linux`, `darwin`, `windows` (from `GOOS`/`runtime.GOOS`) |
| Target arch | `amd64`, `arm64`, `386` (from `GOARCH`/`runtime.GOARCH`) |
| Arch features | `amd64.v2`, `arm.7`, `arm64.v8.1` (from `GOAMD64`, `GOARM`, etc.) |
| Unix tag | `unix` — satisfied if GOOS is a Unix or Unix-like system |
| Compiler | `gc` or `gccgo` |
| cgo | `cgo` — if `CGO_ENABLED=1` and cgo is supported |
| Go version | `go1.1`, `go1.12`, `go1.23` — through the current version |
| Additional tags | Any tags passed via `-tags` flag |

### Common Build Constraint Patterns

```go
//go:build ignore
// Exclude a file from all builds (conventional)

//go:build linux && amd64
// Only on Linux amd64

//go:build !windows
// On everything except Windows

//go:build cgo && (linux || darwin)
// Only with cgo on Linux or macOS

//go:build purego
// Pure Go implementation (convention for assembly alternatives)
```

Each `//go:build` line is a separate comment — the explanatory comment on the following line is not part of the constraint.

### Implicit Build Constraints from File Names

File names matching `*_GOOS`, `*_GOARCH`, or `*_GOOS_GOARCH` (after stripping extension and `_test` suffix) create implicit constraints:

```
dns_windows.go         → //go:build windows
math_386.s             → //go:build 386
source_windows_amd64.go → //go:build windows && amd64
```

### OS Compatibility Tags

Some GOOS values imply additional tags:

| GOOS | Also matches |
|------|-------------|
| `android` | `linux` tags and files |
| `illumos` | `solaris` tags and files |
| `ios` | `darwin` tags and files |

### Architecture Feature Tags

Architecture feature tags cascade — a higher level sets all lower levels too:

| GOARCH | Variable | Feature Tags |
|--------|----------|-------------|
| `amd64` | `GOAMD64` | `v1` (default), `v2`, `v3`, `v4` — e.g., `GOAMD64=v2` sets both `amd64.v1` and `amd64.v2` |
| `arm` | `GOARM` | `5`, `6`, `7` — e.g., `arm.5`, `arm.6`, `arm.7` |
| `arm64` | `GOARM64` | `v8.0` (default) through `v9.5` |
| `mips/mipsle` | `GOMIPS` | `hardfloat` (default), `softfloat` |
| `mips64/mips64le` | `GOMIPS64` | `hardfloat`, `softfloat` |
| `ppc64/ppc64le` | `GOPPC64` | `power8` (default), `power9`, `power10` |
| `riscv64` | `GORISCV64` | `rva20u64` (default), `rva22u64`, `rva23u64` |
| `wasm` | `GOWASM` | `satconv`, `signext` |

Use negation for code handling the absence of a feature:

```go
//go:build !amd64.v2
```

### Language Version Downgrade

In modules with Go version 1.21+, if a file's build constraint has a term for a Go major release (e.g., `go1.20`), the language version used to compile that file is the minimum version implied by the constraint.

### Legacy `// +build` Syntax

Go 1.16 and earlier used `// +build` syntax. `gofmt` will add an equivalent `//go:build` constraint when encountering the older syntax. Both forms are recognized, but `//go:build` is preferred.

## Build Modes

The `-buildmode` flag controls what kind of object file is produced:

| Mode | Description |
|------|-------------|
| `archive` | Build non-main packages into `.a` files (main packages ignored) |
| `c-archive` | Build main package into a C archive. Callable symbols are cgo `//export` functions. Requires exactly one main package. |
| `c-shared` | Build main package into a C shared library. Callable symbols are cgo `//export` functions. On wasip1, builds a WASI reactor (uses `//go:wasmexport`). |
| `default` | Main packages → executables; non-main → `.a` files (default behavior) |
| `shared` | Combine non-main packages into a shared library (for use with `-linkshared`) |
| `exe` | Build main packages and all imports into executables (non-main ignored) |
| `pie` | Build position independent executables (PIE) |
| `plugin` | Build main packages into Go plugins (non-main ignored) |

### AIX Note

On AIX, when linking a C program that uses a Go archive built with `-buildmode=c-archive`, pass `-Wl,-bnoobjreorder` to the C compiler.

## Cross-Compilation

Set `GOOS` and `GOARCH` environment variables to cross-compile:

```bash
GOOS=linux GOARCH=arm64 go build     # Linux ARM64 binary
GOOS=windows GOARCH=amd64 go build   # Windows AMD64 binary
GOOS=darwin GOARCH=arm64 go build    # macOS Apple Silicon
```

Cross-compiled binaries are installed in `$GOOS_$GOARCH` subdirectories by `go install`.

See [environment.md](environment.md) for all supported `GOOS`/`GOARCH` combinations and architecture-specific variables.

## Profile-Guided Optimization (PGO)

PGO uses a runtime profile to optimize hot paths in the compiled binary:

```bash
# 1. Build and run with profiling to generate default.pgo
go build -pgo=auto      # uses default.pgo in package dir if present

# 2. Explicit profile
go build -pgo=profile.pgo

# 3. Disable PGO
go build -pgo=off
```

The default is `auto` — for each main package, the go command selects `default.pgo` in the package's directory if it exists, and applies it to the transitive dependencies of that main package.

## The `-overlay` Flag

The overlay flag reads a JSON config file that provides a virtual file system overlay for build operations:

```json
{
    "Replace": {
        "path/on/disk.go": "actual/file.go",
        "nonexistent.go": ""
    }
}
```

- Maps disk file paths to backing file paths
- If the backing path is empty, the disk file is treated as non-existent
- Limitations: cgo files included from outside the include path must be in the same directory as the Go package; overlays do not appear when running binaries via `go run` and `go test`; files beneath `GOMODCACHE` cannot be replaced

## Build JSON Output

`go build -json`, `go install -json`, and `go test -json` emit build output as a newline-separated stream of `BuildEvent` objects:

```go
type BuildEvent struct {
    ImportPath string
    Action     string    // "build-output" or "build-fail"
    Output     string    // set when Action == "build-output"
}
```

- `ImportPath` matches `Package.ImportPath` from `go list -json`
- Concatenating all `Output` fields reconstructs the full build output
- For `go test -json`, parsers can distinguish `BuildEvent` from `TestEvent` by inspecting the `Action` field
- Non-JSON error text may still appear on stderr (indicates early, serious errors)

## Edge Cases and Gotchas

### Build Cache and C Libraries

The build cache does **not** detect changes to C libraries imported with cgo. After modifying C libraries, run `go clean -cache` or use `-a` to force rebuilding.

### `-C` Must Be First

The `-C` flag must be the first flag on the command line. Any files named are interpreted after changing directories.

### `-modfile` Still Requires `go.mod`

When using `-modfile`, a `go.mod` file must still exist in the module root to determine the module root directory — it's just not accessed. The `.sum` file is derived by trimming `.mod` and appending `.sum`.

### `go build` vs `go install`

- `go build` compiles but does not install. For non-main packages, it serves as a compile check.
- `go install` compiles and installs — executables to `$GOBIN` (or `$GOPATH/bin`), libraries are cached.

### Using Lower-Level Tools

For projects that can't follow go conventions, lower-level invocations bypass some overhead:

```bash
go tool compile foo.go    # direct compiler invocation
go tool link foo.o        # direct linker invocation
```

See [tool.md](tool.md) for details on `go tool` subcommands.

## References

- `go help build` — full build flag reference
- `go help buildconstraint` — build constraint syntax
- `go help buildmode` — build modes
- `go help buildjson` — JSON build output format
- `go help c` — calling between Go and C

## Related

- [concepts.md](concepts.md) — packages, import paths, module discovery
- [environment.md](environment.md) — `GOCACHE`, `GOOS`, `GOARCH`, `GODEBUG`
- [tool.md](tool.md) — `go tool compile`, `go tool link`, `go tool asm`
- [cli.md](cli.md) — `go install`, `go run`, `go clean`
