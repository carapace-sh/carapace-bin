# Go Environment Variables

Reference for all environment variables consulted by the `go` command — general-purpose, cgo, architecture-specific, code coverage, and special-purpose variables. Covers `go env` for inspecting and persisting settings.

> **Source of truth**: `go help environment`, `go help env`. For **module-specific variables** (`GOPROXY`, `GOSUMDB`, `GOPRIVATE`, etc.), see also [mod.md](mod.md).

## `go env` Command

```
go env [-json] [-changed] [-u] [-w] [var ...]
```

Prints Go environment information.

| Flag | Description |
|------|-------------|
| (none) | Print all env vars as shell script (batch file on Windows) |
| `var ...` | Print value of each named variable on its own line |
| `-json` | Print environment as JSON |
| `-changed` | Print only settings that differ from the default |
| `-w NAME=VALUE` | Persist a default setting to the env config file |
| `-u NAME` | Unset a previously `-w`-set default |

### The Env Config File

`go env -w` writes to a per-user config file in `os.UserConfigDir`. Its location is controlled by `GOENV`:

```bash
go env GOENV          # print the config file path
GOENV=off             # disable the config file (use only env vars)
```

`GOENV` cannot be changed with `go env -w` — set it in the shell environment.

### Common Usage

```bash
go env GOPATH                          # check GOPATH
go env -w GOPRIVATE=*.corp.example.com # persist GOPRIVATE
go env -w GOBIN=$HOME/.local/bin       # persist GOBIN
go env -u GOBIN                        # unset persisted GOBIN
go env -json                           # all vars as JSON
go env -changed                        # only non-defaults
```

## General-Purpose Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `GCCGO` | gccgo command for `go build -compiler=gccgo` | `gccgo` |
| `GO111MODULE` | Module-aware mode: `off`, `on`, or `auto` | `on` (since Go 1.16; was `auto` before) |
| `GOARCH` | Target architecture (`amd64`, `arm64`, `386`, etc.) | Host architecture |
| `GOAUTH` | Authentication for go-import and HTTPS module mirror | `netrc` |
| `GOBIN` | Directory where `go install` puts commands | `$GOPATH/bin` |
| `GOCACHE` | Build cache directory (absolute path) | User cache dir `/go-build` |
| `GOCACHEPROG` | External build cache program (with flags) | (none) |
| `GODEBUG` | Debugging facilities (comma-separated). Cannot be set with `go env -w`. | (none) |
| `GOENV` | Location of the env config file. Cannot be set with `go env -w`. `off` disables config file. | User config dir |
| `GOFLAGS` | Space-separated `-flag=value` settings applied by default | (none) |
| `GOINSECURE` | Comma-separated glob patterns for insecure (non-HTTPS) fetching | (none) |
| `GOMODCACHE` | Module cache directory | `$GOPATH/pkg/mod` |
| `GOOS` | Target operating system (`linux`, `darwin`, `windows`, etc.) | Host OS |
| `GOPATH` | Workspace root (legacy) and bin/pkg locations | `$HOME/go` |
| `GOPRIVATE` | Comma-separated globs for private modules (bypasses proxy + checksum) | (none) |
| `GONOPROXY` | Comma-separated globs to bypass proxy (overrides `GOPRIVATE` for proxy) | (none) |
| `GONOSUMDB` | Comma-separated globs to bypass checksum DB (overrides `GOPRIVATE` for checksum) | (none) |
| `GOPROXY` | Module proxy URL(s), comma-separated | `https://proxy.golang.org,direct` |
| `GOROOT` | Root of the Go installation tree | Auto-detected |
| `GOSUMDB` | Checksum database name and optional key/URL | `sum.golang.org` |
| `GOTMPDIR` | Temporary directory for the go command and testing | OS temp dir (`/tmp`) |
| `GOTOOLCHAIN` | Which Go toolchain to use | `auto` |
| `GOVCS` | Allowed VCS commands per module pattern | `public:git|hg,private:all` |
| `GOWORK` | Workspace file path. `auto` searches up; `off` disables. | `auto` |

### `GOFLAGS`

Space-separated flags applied to every go command (when the flag is known to that command):

```
GOFLAGS=-mod=vendor -trimpath
```

Each entry must be a standalone flag. Values must not contain spaces. Command-line flags override `GOFLAGS`.

### `GODEBUG`

Enables debugging facilities. Key values include:

| GODEBUG | Description |
|---------|-------------|
| `gocacheverify=1` | Bypass cache, rebuild everything, verify results match cached |
| `gocachehash=1` | Print inputs for all content hashes (verbose) |
| `gocachetest=1` | Print details of cached test reuse decisions |
| `installgoroot=all` | Restore `$GOROOT/pkg/GOOS_GOARCH` installation (pre-1.20 behavior) |

See <https://go.dev/doc/godebug> for the full list. `GODEBUG` cannot be set with `go env -w` — use the shell environment or `go.mod`'s `godebug` directive.

### `GOTOOLCHAIN`

Controls which Go toolchain is used. The setting takes the form `<name>`, `<name>+auto`, or `<name>+path`:

| Value | Description |
|-------|-------------|
| `auto` (default) | Shorthand for `local+auto` — use the local toolchain, but auto-switch to a newer toolchain if `go.mod` requires it (with download fallback) |
| `path` | Shorthand for `local+path` — use the local toolchain, search PATH for newer ones, no download fallback |
| `local` | Always use the bundled (local) Go toolchain — no switching |
| `go1.23.4` | Use a specific toolchain version — no switching |
| `go1.23.4+auto` | Default to this version, but still auto-switch upward if needed |
| `go1.23.4+path` | Default to this version, search PATH for newer ones, no download fallback |

The `go` command can download and switch toolchains automatically. See <https://go.dev/doc/toolchain>.

## Cgo Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `AR` | Command for library archives (gccgo builds) | `ar` |
| `CC` | C compiler command | Platform default (e.g., `gcc`) |
| `CGO_CFLAGS` | Flags passed to C compiler | (none) |
| `CGO_CFLAGS_ALLOW` | Regex of additional allowed `#cgo CFLAGS` directives | (none) |
| `CGO_CFLAGS_DISALLOW` | Regex of disallowed `#cgo CFLAGS` directives | (none) |
| `CGO_CPPFLAGS` | C preprocessor flags | (none) |
| `CGO_CPPFLAGS_ALLOW` | Regex of allowed CPPFLAGS directives | (none) |
| `CGO_CPPFLAGS_DISALLOW` | Regex of disallowed CPPFLAGS directives | (none) |
| `CGO_CXXFLAGS` | C++ compiler flags | (none) |
| `CGO_CXXFLAGS_ALLOW` | Regex of allowed CXXFLAGS directives | (none) |
| `CGO_CXXFLAGS_DISALLOW` | Regex of disallowed CXXFLAGS directives | (none) |
| `CGO_ENABLED` | Whether cgo is supported: `0` or `1` | `1` (if C compiler available) |
| `CGO_FFLAGS` | Fortran compiler flags | (none) |
| `CGO_FFLAGS_ALLOW` | Regex of allowed FFLAGS directives | (none) |
| `CGO_FFLAGS_DISALLOW` | Regex of disallowed FFLAGS directives | (none) |
| `CGO_LDFLAGS` | Linker flags for cgo | (none) |
| `CGO_LDFLAGS_ALLOW` | Regex of allowed LDFLAGS directives | (none) |
| `CGO_LDFLAGS_DISALLOW` | Regex of disallowed LDFLAGS directives | (none) |
| `CXX` | C++ compiler command | Platform default |
| `FC` | Fortran compiler command | Platform default |
| `PKG_CONFIG` | Path to pkg-config tool | `pkg-config` |

The `_ALLOW` and `_DISALLOW` regexes apply to `#cgo` source code directives, **not** to the corresponding environment variables.

## Architecture-Specific Variables

| Variable | GOARCH | Description | Valid Values |
|----------|--------|-------------|--------------|
| `GO386` | `386` | Floating point implementation | `sse2` (default), `softfloat` |
| `GOAMD64` | `amd64` | Microarchitecture level | `v1` (default), `v2`, `v3`, `v4` |
| `GOARM` | `arm` | ARM architecture | `5`, `6`, `7`. Default: `7` when cross-compiling; based on build system when on arm. Optional `,softfloat` or `,hardfloat` suffix. |
| `GOARM64` | `arm64` | ARM64 architecture | `v8.0` (default) through `v9.5`. Optional `,lse`, `,crypto` suffixes. |
| `GOMIPS` | `mips`, `mipsle` | Floating point | `hardfloat` (default), `softfloat` |
| `GOMIPS64` | `mips64`, `mips64le` | Floating point | `hardfloat` (default), `softfloat` |
| `GOPPC64` | `ppc64`, `ppc64le` | Target ISA | `power8` (default), `power9`, `power10` |
| `GORISCV64` | `riscv64` | RISC-V profile | `rva20u64` (default), `rva22u64`, `rva23u64` |
| `GOWASM` | `wasm` | Experimental WebAssembly features | `satconv`, `signext` (comma-separated) |

### Architecture Feature Cascade

For `amd64`, `arm`, `ppc64`, `ppc64le`, and `riscv64`, a feature level sets all lower levels too. For example, `GOAMD64=v2` sets both `amd64.v1` and `amd64.v2` build tags. See [build.md](build.md) for build constraint details.

## Code Coverage Variable

| Variable | Description |
|----------|-------------|
| `GOCOVERDIR` | Directory for coverage data files from `go build -cover` binaries |

When running a binary built with `go build -cover`, coverage data is written to `GOCOVERDIR`:

```bash
go build -cover -o myapp
GOCOVERDIR=./coverage-data ./myapp
go tool cover -func=./coverage-data/*  # analyze
```

## Special-Purpose Variables

| Variable | Description |
|----------|-------------|
| `GCCGOTOOLDIR` | Where to find gccgo tools (e.g., cgo). Default: based on gccgo config. |
| `GOEXPERIMENT` | Comma-separated toolchain experiments to enable/disable. See `GOROOT/src/internal/goexperiment/flags.go`. For Go toolchain development only. |
| `GOFIPS140` | FIPS-140 cryptography mode. Default `off`. Other values enable compliance. See <https://go.dev/doc/security/fips140>. |
| `GO_EXTLINK_ENABLED` | Whether linker uses external linking mode with `-linkmode=auto` + cgo. `0` disable, `1` enable. |
| `GIT_ALLOW_PROTOCOL` | Defined by Git. Colon-separated allowed schemes for `git fetch/clone`. Cannot be set with `go env -w`. |

## Read-Only Variables (available from `go env` but not set from environment)

| Variable | Description |
|----------|-------------|
| `GOEXE` | Executable suffix (`.exe` on Windows, `` elsewhere) |
| `GOGCCFLAGS` | Space-separated args supplied to the CC command |
| `GOHOSTARCH` | Architecture of the Go toolchain binaries |
| `GOHOSTOS` | OS of the Go toolchain binaries |
| `GOMOD` | Absolute path to the main module's `go.mod`. `/dev/null` if module-aware but no go.mod. Empty if GOPATH mode. |
| `GOTELEMETRY` | Current telemetry mode (`off`, `local`, or `on`) |
| `GOTELEMETRYDIR` | Directory where telemetry data is written |
| `GOTOOLDIR` | Directory where go tools (compile, cover, etc.) are installed |
| `GOVERSION` | Version of the installed Go tree (`runtime.Version()`) |

## Build Cache Debugging

The build cache (controlled by `GOCACHE`) can be debugged with `GODEBUG`:

```bash
GODEBUG=gocacheverify=1 go build ./...    # bypass cache, verify
GODEBUG=gocachehash=1 go build ./...      # print hash inputs (verbose)
GODEBUG=gocachetest=1 go test ./...       # print test cache decisions
```

The go command periodically deletes unused cached data. `go clean -cache` removes all cached data. See `go help cache` for details.

### Cgo Cache Invalidation

The build cache does **not** detect changes to C libraries imported with cgo. After modifying C libraries:

```bash
go clean -cache          # clear entire cache
# or
go build -a ./...        # force rebuild
```

## Cross-Compilation

Set `GOOS` and `GOARCH` to cross-compile:

```bash
GOOS=linux GOARCH=arm64 go build
GOOS=windows GOARCH=amd64 go build
GOOS=darwin GOARCH=arm64 go build
```

Cross-compiled binaries from `go install` go to `$GOOS_$GOARCH` subdirectories.

### Supported GOOS Values

`aix`, `android`, `darwin`, `dragonfly`, `freebsd`, `illumos`, `ios`, `js`, `linux`, `netbsd`, `openbsd`, `plan9`, `solaris`, `wasip1`, `windows`

### Supported GOARCH Values

`386`, `amd64`, `arm`, `arm64`, `loong64`, `mips`, `mips64`, `mips64le`, `mipsle`, `ppc64`, `ppc64le`, `riscv64`, `s390x`, `wasm`

## Edge Cases and Gotchas

### `GODEBUG` Cannot Use `go env -w`

`GODEBUG` must be set in the shell environment or via the `godebug` directive in `go.mod`/`go.work`. `go env -w GODEBUG=...` will fail.

### `GOENV` Cannot Use `go env -w`

`GOENV` must be set in the shell environment. `go env -w GOENV=...` will fail. Set `GOENV=off` to disable the config file entirely.

### `GOFLAGS` Space-Separated

`GOFLAGS` is space-separated (not comma-separated). Each entry must be a standalone flag. Values must not contain spaces:

```
GOFLAGS=-mod=vendor -trimpath    # correct
GOFLAGS=-mod=vendor,-trimpath    # wrong (not comma-separated)
```

### `CGO_ENABLED` and Cross-Compilation

When cross-compiling, `CGO_ENABLED` defaults to `0` (since a C cross-compiler is typically not available). Enable explicitly if you have a cross C compiler:

```bash
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build
```

### `GIT_ALLOW_PROTOCOL` Is Git's Variable

`GIT_ALLOW_PROTOCOL` is defined by Git, not Go. The go command reads it but cannot set a default with `go env -w`.

## References

- `go help environment` — full environment variable reference
- `go help env` — `go env` command reference
- `go help cache` — build cache details
- <https://go.dev/doc/godebug> — GODEBUG settings
- <https://go.dev/doc/toolchain> — GOTOOLCHAIN documentation
- <https://go.dev/doc/security/fips140> — GOFIPS140 documentation

## Related

- [mod.md](mod.md) — `GOPROXY`, `GOSUMDB`, `GOPRIVATE`, `GONOPROXY`, `GONOSUMDB`, `GOINSECURE`, `GOVCS`, `GOMODCACHE`
- [build.md](build.md) — cross-compilation, build cache, `-trimpath`, build constraints
- [work.md](work.md) — `GOWORK`
- [concepts.md](concepts.md) — `GOPATH`, `GOROOT`, `GO111MODULE`
