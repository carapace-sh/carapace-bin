# Go Core Concepts

The foundational mental model for the `go` command — modules, packages, import paths, GOPATH vs module-aware mode, and package patterns.

> **Source of truth**: `go help modules`, `go help packages`, `go help gopath`, `go help importpath`. For **build flags and compilation**, see [build.md](build.md). For **module operations**, see [mod.md](mod.md). For **environment variables**, see [environment.md](environment.md).

## Modules vs GOPATH Mode

The `go` command operates in one of two modes, controlled by `GO111MODULE`:

| Mode | `GO111MODULE` | Description |
|------|---------------|-------------|
| **Module-aware** | `on` or `auto` (with `go.mod` present) | Uses `go.mod` for dependency management. The default since Go 1.16. |
| **GOPATH** | `off` | Uses `GOPATH` directory structure for package resolution. Legacy mode. |

In `auto` mode (the default before Go 1.16), the go command enables module-aware mode when a `go.mod` is found in the current directory or any parent, and falls back to GOPATH mode otherwise. Since Go 1.16, `GO111MODULE` defaults to `on` — module-aware mode is always enabled, even without a `go.mod` (in which case most commands error).

### Module-Aware Mode

A **module** is a collection of Go packages released, versioned, and distributed together. A module is defined by a `go.mod` file at its root:

```
module example.com/myproject

go 1.23

require (
    github.com/spf13/cobra v1.8.1
    golang.org/x/net v0.30.0
)
```

The `go` command searches for `go.mod` by walking up from the current directory. The directory containing `go.mod` is the **module root**.

### GOPATH Mode (Legacy)

In GOPATH mode, source code lives under `$GOPATH/src`. The import path is determined by the directory structure relative to `src`:

```
$GOPATH/
  src/
    example.com/myproject/
      main.go
    github.com/spf13/cobra/
      ...
  bin/          # installed commands
  pkg/          # compiled package objects (GOOS_GOARCH subdirs)
```

GOPATH mode is rarely used with modern Go. See `go help gopath` for full details.

## Packages and Import Paths

An **import path** is a string that uniquely identifies a package. The `go` command resolves import paths to directories on disk:

| Import Path Form | Resolution |
|-----------------|------------|
| `fmt` | Standard library package in `$GOROOT/src/fmt` |
| `example.com/user/repo` | Module from a remote repository (module-aware) |
| `./foo/bar` | Relative file system path (must be within a module) |
| `github.com/user/repo/sub` | Subdirectory of a remote module |

### Reserved Package Names

Five names have special meaning in package patterns:

| Name | Meaning |
|------|---------|
| `main` | Top-level package in a standalone executable |
| `all` | All packages in the main module + dependencies (including test dependencies). In GOPATH mode, all packages in all GOPATH trees. |
| `std` | All packages in the standard Go library |
| `cmd` | Go repository's commands and their internal libraries |
| `tool` | Tools defined in the current module's `go.mod` file (`tool` directives) |

### Package Patterns

Import paths can be **patterns** containing `...` wildcards:

```bash
go test ./...          # all packages in current module (recursively)
go build example.com/... # all packages under example.com
go test net/...        # net and all its subpackages
```

Special pattern rules:

- `...` matches any string, including slashes and the empty string
- `net/...` matches both `net` and `net/http` (trailing `/...` can match empty)
- A wildcard element never matches a `vendor` path element: `./...` skips `./vendor/` subdirectories
- `./vendor/...` explicitly matches vendored packages

### File List as Package

If the package arguments are a list of `.go` files from a single directory, the go command treats them as a single synthesized package, **ignoring build constraints** and any other files in the directory:

```bash
go build foo.go bar.go    # builds a package from just these two files
```

## Package Naming Conventions

### The `main` Package

A package named `main` is compiled into an executable, not a library. The executable is named after the last non-major-version component of the import path:

```bash
go build example.com/sam        # produces 'sam' or 'sam.exe'
go build example.com/foo/v2     # produces 'foo' or 'foo.exe', not 'v2.exe'
```

When building from `.go` files, the executable is named after the first source file:

```bash
go build ed.go rx.go    # produces 'ed' or 'ed.exe'
```

### The `documentation` Package

A package named `documentation` is treated as documentation for a non-Go program in the directory. Files in such a package are ignored by the go command.

## Test Packages

Test files (`*_test.go`) can declare a package with the `_test` suffix, creating a **separate test package** that is compiled and linked alongside the main test binary:

```go
// In foo.go:
package foo

// In foo_test.go (external test package):
package foo_test

import (
    "testing"
    "github.com/example/foo"  // imports the package under test
)
```

This pattern prevents import cycles in tests and forces tests to use only the package's exported API. See [test.md](test.md) for testing details.

## The `testdata` Directory

Any directory named `testdata` is ignored by the go command for build purposes. It is available to hold ancillary data needed by tests:

```
myproject/
  foo.go
  foo_test.go
  testdata/
    input.txt        # test fixture, not compiled
    golden.txt       # expected output
```

## File Types

The go command recognizes files by extension. Only these extensions are examined:

| Extension | Type | Notes |
|-----------|------|-------|
| `.go` | Go source | Compiled by the Go compiler |
| `.c`, `.h` | C source | Compiled with OS-native compiler if cgo/SWIG used |
| `.cc`, `.cpp`, `.cxx`, `.hh`, `.hpp`, `.hxx` | C++ source | Only useful with cgo/SWIG |
| `.m` | Objective-C | Only useful with cgo |
| `.s`, `.S`, `.sx` | Assembly | Go assembler (or OS-native if cgo/SWIG) |
| `.swig` | SWIG definition | Passed to SWIG |
| `.swigcxx` | SWIG C++ definition | Passed to SWIG with `-c++` |
| `.syso` | System object | Added to archive |

Files of each type (except `.syso`) may contain build constraints, but the go command stops scanning for constraints at the first non-blank, non-comment line. See [build.md](build.md) for build constraint syntax.

## Ignored Directories and Files

The go command ignores:

- Directories beginning with `.` or `_`
- Files beginning with `.` or `_` (including `_test.go` — only `*_test.go` is recognized as test files)
- Directories named `testdata`

## Source File Naming and Implicit Build Constraints

If a file's name (after stripping the extension and a possible `_test` suffix) matches `*_GOOS`, `*_GOARCH`, or `*_GOOS_GOARCH`, the file has an **implicit build constraint** requiring those terms:

```
source_windows.go         # only built when GOOS=windows
math_386.s                # only built when GOARCH=386
source_windows_amd64.go   # only built when GOOS=windows && GOARCH=amd64
```

This is equivalent to writing `//go:build windows` (or the corresponding constraint) at the top of the file. See [build.md](build.md) for full build constraint details.

## Vendor Directories

A `vendor` directory at the module root contains vendored copies of dependencies. When a `vendor` directory is present and the `go` version in `go.mod` is 1.14+, the go command acts as if `-mod=vendor` were set, using vendored packages instead of downloading them.

Key vendor rules:

- `./...` patterns do not match packages in `vendor` subdirectories
- `./vendor/...` explicitly matches vendored packages
- A directory named `vendor` that contains code but is not at the module root is **not** a vendored package — it's a regular package named `vendor`
- The `vendor/modules.txt` file records the module versions included

See [mod.md](mod.md) for `go mod vendor` details.

## GOPATH and GOROOT

| Variable | Purpose |
|----------|---------|
| `GOROOT` | Root of the Go installation tree (compiler, standard library, tools) |
| `GOPATH` | Workspace root for Go code (legacy) and installed binaries (`$GOPATH/bin`) |

In module-aware mode, `GOPATH` is mostly used for:

- `$GOPATH/bin` — where `go install` puts executables (unless `GOBIN` is set)
- `$GOPATH/pkg` — legacy compiled package cache (not used in module mode)
- Module cache location (defaults to `$GOPATH/pkg/mod`, overridden by `GOMODCACHE`)

If `GOPATH` is unset, it defaults to `$HOME/go` (unless that directory holds a Go distribution). Run `go env GOPATH` to see the effective value.

## Module Discovery

The go command discovers the main module by searching for `go.mod` in the current directory and then successive parent directories. The directory containing `go.mod` becomes the module root.

When no `go.mod` is found and `GO111MODULE=auto`, the go command falls back to GOPATH mode. With `GO111MODULE=on` (the default since Go 1.16), an error is reported if no `go.mod` is found.

Use `go env GOMOD` to see the path to the active `go.mod` file:

```bash
go env GOMOD    # /path/to/module/go.mod or /dev/null if no module
```

## References

- `go help modules` — module system overview
- `go help packages` — package lists and patterns
- `go help gopath` — GOPATH structure
- `go help importpath` — import path syntax
- `go help filetype` — recognized file extensions
- <https://go.dev/ref/mod> — Go modules reference

## Related

- [build.md](build.md) — build flags, build constraints, build modes
- [mod.md](mod.md) — `go mod` subcommands and module operations
- [environment.md](environment.md) — `GO111MODULE`, `GOPATH`, `GOROOT`, `GOMODCACHE`
