# Go Mod — Module Management and Dependencies

Reference for `go mod` subcommands, the `go.mod` file format, module versioning, the module proxy protocol, checksum verification, and dependency resolution via Minimal Version Selection (MVS).

> **Source of truth**: `go help mod`, `go help modules`, `go help go.mod`, `go help module-auth`, `go help goproxy`, `go help private`, `go help vcs`. For **environment variables** controlling module behavior, see [environment.md](environment.md). For **`go get`** (dependency adjustment), see [cli.md](cli.md).

## Module System Overview

A **module** is a collection of Go packages released, versioned, and distributed together. Modules are downloaded from module proxy servers or directly from version control repositories.

Support for modules is built into **all** go commands, not just `go mod`. Day-to-day dependency management uses `go get` (add, remove, upgrade, downgrade). `go mod` provides lower-level operations.

## The `go.mod` File

The `go.mod` file is the module definition. It is located at the module root and is discovered by walking up from the current directory.

### Directives

| Directive | Description |
|-----------|-------------|
| `module <path>` | Module path (e.g., `module example.com/myproject`) |
| `go <version>` | Expected Go language version |
| `toolchain <name>` | Go toolchain to use (e.g., `toolchain go1.23.4`) |
| `godebug <key>=<value>` | GODEBUG key-value pairs |
| `require <path> <version>` | Dependency on a module at a specific version |
| `replace <old> => <new>` | Replace a module with another (local path or different version) |
| `exclude <path> <version>` | Exclude a specific module version |
| `retract <version>` | Retract a published version (with optional rationale) |
| `tool <path>` | Tool dependency (path to a tool's main package) |
| `ignore <path>` | Ignore a path |

### Example

```
module example.com/myproject

go 1.23

toolchain go1.23.4

require (
    github.com/spf13/cobra v1.8.1
    golang.org/x/net v0.30.0
)

require github.com/google/uuid v1.6.0 // indirect

replace example.com/old => example.com/new v1.2.0

replace example.com/local => ../local-module

exclude example.com/broken v1.5.0

retract [v1.1.0, v1.1.9] // security vulnerability in this range

tool golang.org/x/tools/cmd/goimports
```

### The `go` Directive and Language Version

The `go` directive specifies:
- The **Go language version** used to compile the module
- Which module graph loading semantics apply (Go 1.17+ uses lazy module loading, retaining more requirements in `go.mod`)

### Indirect Dependencies

Dependencies needed only transitively are marked `// indirect`:

```
require github.com/google/uuid v1.6.0 // indirect
```

`go mod tidy` adds `// indirect` markers. Direct dependencies (imported by the module's own packages) have no marker.

### The `replace` Directive

`replace` redirects a module to a different source:

```
replace example.com/old v1.2.0 => example.com/new v1.4.5    # specific version
replace example.com/old => example.com/new v1.4.5           # all versions
replace example.com/old => ../local-path                     # local directory
```

Local replaces use a **directory path**, not a module path. The directory must contain a `go.mod` file.

### The `retract` Directive

`retract` marks a published version as withdrawn:

```
retract v1.1.0                         # single version
retract [v1.1.0, v1.1.9]              # version range
retract [v1.1.0, v1.1.9]              # with rationale on next line
// Security vulnerability in this range
```

The `go` command won't automatically select retracted versions and warns when they're in use.

### The `tool` Directive

The `tool` directive declares a tool dependency — a main package that can be run with `go tool`:

```
tool golang.org/x/tools/cmd/goimports
tool golang.org/x/tools/cmd/stringer
```

Tools are installed to the module cache and run via `go tool <name>`. `go get -tool <path>` adds a tool directive. `go get -tool <path>@none` removes it.

## The `go.sum` File

`go.sum` records cryptographic hashes of module contents to verify integrity:

```
github.com/spf13/cobra v1.8.1 h1:...
github.com/spf13/cobra v1.8.1/go.mod h1:...
```

Two hashes per version:
- Module zip hash (`h1:...`)
- `go.mod` hash (`/go.mod h1:...`)

Hashes are verified against the checksum database (GOSUMDB) and/or computed locally. See [Module Authentication](#module-authentication) below.

## `go mod` Subcommands

### `go mod init`

```
go mod init [module-path]
```

Creates a new `go.mod` in the current directory. The `go.mod` must not already exist.

If the module path is omitted, `init` attempts to infer it from import comments in `.go` files and the current directory (if in GOPATH).

### `go mod tidy`

```
go mod tidy [-e] [-v] [-x] [-diff] [-go=version] [-compat=version]
```

Ensures `go.mod` matches the source code:
- Adds missing modules needed to build packages and dependencies
- Removes unused modules that don't provide relevant packages
- Adds missing `go.sum` entries and removes unnecessary ones

| Flag | Description |
|------|-------------|
| `-v` | Print information about removed modules |
| `-e` | Proceed despite errors loading packages |
| `-diff` | Print changes as unified diff instead of modifying files. Non-zero exit if diff is non-empty. |
| `-go=version` | Update the `go` directive (may change which requirements are retained) |
| `-compat=version` | Preserve checksums needed for the indicated Go release. Default: version prior to the `go` directive. |
| `-x` | Print commands download executes |

### `go mod download`

```
go mod download [-x] [-json] [-reuse=old.json] [modules]
```

Downloads modules to the local cache. With no arguments, downloads all modules needed to build and test the main module.

Arguments can be module patterns (dependencies of the main module) or module queries (`path@version`).

The `-json` flag prints a sequence of JSON objects:

```go
type Module struct {
    Path     string // module path
    Query    string // version query
    Version  string // module version
    Error    string // error loading module
    Info     string // path to cached .info file
    GoMod    string // path to cached .mod file
    Zip      string // path to cached .zip file
    Dir      string // path to cached source root
    Sum      string // checksum (as in go.sum)
    GoModSum string // checksum for go.mod
    Origin   any    // provenance
    Reuse    bool   // reuse of old info is safe
}
```

The `-reuse` flag accepts a file containing JSON output from a previous `go mod download -json` to avoid redownloading unchanged modules.

### `go mod edit`

```
go mod edit [editing flags] [-fmt|-print|-json] [go.mod]
```

Command-line interface for editing `go.mod`, primarily for tools/scripts. Reads only `go.mod` — does not look up module information.

| Editing Flag | Description |
|--------------|-------------|
| `-fmt` | Reformat `go.mod` without other changes |
| `-module path` | Change the module path |
| `-godebug=key=value` | Add/replace a godebug line |
| `-dropgodebug=key` | Drop a godebug line |
| `-require=path@version` | Add/replace a requirement |
| `-droprequire=path` | Drop a requirement |
| `-go=version` | Set the expected Go version |
| `-toolchain=version` | Set the Go toolchain |
| `-exclude=path@version` | Add an exclusion |
| `-dropexclude=path@version` | Drop an exclusion |
| `-replace=old[@v]=new[@v]` | Add/replace a replacement |
| `-dropreplace=old[@v]` | Drop a replacement |
| `-retract=version` | Add a retraction (single version or `[low,high]` interval) |
| `-dropretract=version` | Drop a retraction |
| `-tool=path` | Add a tool declaration |
| `-droptool=path` | Drop a tool declaration |
| `-ignore=path` | Add an ignore declaration |
| `-dropignore=path` | Drop an ignore declaration |
| `-print` | Print final go.mod text instead of writing back |
| `-json` | Print final go.mod as JSON |

Editing flags may be **repeated**, with changes applied in order.

The `-json` output corresponds to:

```go
type GoMod struct {
    Module    ModPath
    Go        string
    Toolchain string
    Godebug   []Godebug
    Require   []Require
    Exclude   []Module
    Replace   []Replace
    Retract   []Retract
    Tool      []Tool
    Ignore    []Ignore
}

type Require struct {
    Path     string
    Version  string
    Indirect bool
}

type Retract struct {
    Low       string
    High      string
    Rationale string
}
```

### `go mod graph`

```
go mod graph [-go=version] [-x]
```

Prints the module requirement graph (with replacements applied). Each line: `module@version dependency@version`. The main module has no `@version` suffix.

The `-go` flag reports the graph as loaded by the given Go version instead of the `go` directive's version.

### `go mod vendor`

```
go mod vendor [-e] [-v] [-o outdir]
```

Resets the `vendor` directory to include all packages needed to build and test the main module's packages. Does not include test code for vendored packages.

| Flag | Description |
|------|-------------|
| `-v` | Print names of vendored modules and packages |
| `-e` | Proceed despite errors loading packages |
| `-o outdir` | Create vendor directory at given path (default `vendor`). Only `vendor` at module root is used by the go command. |

The `vendor/modules.txt` file records the module versions included.

### `go mod verify`

```
go mod verify
```

Checks that dependencies in the local cache have not been modified since download. Prints "all modules verified." on success, or reports changed modules with non-zero exit.

### `go mod why`

```
go mod why [-m] [-vendor] packages...
```

Shows the shortest path in the import graph from the main module to each listed package.

| Flag | Description |
|------|-------------|
| `-m` | Treat arguments as module names (find path to any package in each module) |
| `-vendor` | Exclude tests of dependencies |

Output is stanzas separated by blank lines:

```
# golang.org/x/text/language
rsc.io/quote
rsc.io/sampler
golang.org/x/text/language

# golang.org/x/text/encoding
(main module does not need package golang.org/x/text/encoding)
```

## Version Queries

Module versions use semantic versioning (`vMAJOR.MINOR.PATCH`). Special queries:

| Query | Description |
|--------|-------------|
| `@latest` | Latest released version |
| `@none` | Remove the dependency (downgrade dependents) |
| `@patch` | Latest patch release of current minor |
| `@v1.2.3` | Specific version |
| `@branch-name` | Latest commit on a branch |
| `@commit-hash` | Specific commit |

## Minimal Version Selection (MVS)

The go command uses **Minimal Version Selection** to resolve the module graph. MVS selects the **minimum** version of each module that satisfies all requirements. Key properties:

- **Deterministic**: The same `go.mod` always produces the same build list
- **Reproducible**: No "latest" resolution at build time — versions are pinned in `go.mod`
- **No solver**: Unlike SAT-solver-based systems, MVS doesn't backtrack

### Lazy Module Loading (Go 1.17+)

With `go 1.17+`, the go command uses lazy module loading:
- Only the direct dependencies of the main module are loaded initially
- The full transitive dependency graph is loaded lazily when needed
- More requirements are retained in `go.mod` to enable this

This means `go.mod` is larger (more `// indirect` entries) but builds are faster and more reproducible.

## Module Authentication

### Checksum Database (GOSUMDB)

The go command validates downloaded modules against a checksum database. Default: `sum.golang.org`.

The checksum database provides a transparent log of module hashes. The go command:
1. Downloads the module zip and/or `go.mod`
2. Computes a cryptographic hash
3. Compares against the known hash in `go.sum` and/or the checksum database

### `go.sum` Entries

Each entry in `go.sum` is a line:

```
<module-path> <version> h1:<hash>
<module-path> <version>/go.mod h1:<hash>
```

The `h1:` prefix indicates SHA-256. Hashes are base64-encoded.

### Private Modules (GOPRIVATE)

The `GOPRIVATE` variable controls which modules are treated as private (bypassing proxy and checksum database):

```
GOPRIVATE=*.corp.example.com,rsc.io/private
```

Glob patterns (Go `path.Match` syntax) match module path prefixes.

### Fine-Grained Control

| Variable | Controls |
|----------|----------|
| `GOPRIVATE` | Both proxy bypass AND checksum bypass (shorthand for `GONOPROXY` + `GONOSUMDB`) |
| `GONOPROXY` | Bypass the module proxy (download directly from VCS) |
| `GONOSUMDB` | Bypass the checksum database |
| `GOINSECURE` | Fetch in insecure manner (no HTTPS). Does **not** disable checksum validation. |

### VCS Restrictions (GOVCS)

The `GOVCS` variable controls which version control commands may be used:

- Default: only `git` and `hg` for public servers; all VCS for private (GOPRIVATE) servers
- Format: `pattern:vcslist` rules, comma-separated
- `vcslist` is pipe-separated, or `all` or `off`

```
GOVCS=github.com:git,*.corp.example.com:all
```

## The Module Proxy Protocol

A Go module proxy is a web server responding to specific GET requests. The protocol:

| URL | Response |
|-----|----------|
| `/<module>/@v/list` | List of known versions |
| `/<module>/@v/<version>.info` | Version metadata JSON |
| `/<module>/@v/<version>.mod` | The `go.mod` file |
| `/<module>/@v/<version>.zip` | The module zip file |

Even a file system serving fixed files can be a proxy (no query parameters needed).

### GOPROXY Values

```
GOPROXY=https://proxy.golang.org,direct    # try proxy first, then direct VCS
GOPROXY=direct                              # always use VCS directly
GOPROXY=off                                 # no downloads allowed
GOPROXY=https://proxy.example.com           # custom proxy
```

Multiple values are separated by commas (`,`) or pipes (`|`):
- **Comma** — fall back to the next entry only on 404 or 410 errors; all other errors are terminal
- **Pipe** — fall back to the next entry on any error (including timeouts)

`direct` means fetch from VCS. `off` disallows all downloads.

## The Module Cache

Downloaded modules are stored in `GOMODCACHE` (default: `$GOPATH/pkg/mod`).

- Cache directories are **read-only** by default (prevents accidental modification)
- Use `-modcacherw` to leave them read-write
- `go clean -modcache` removes the entire cache
- `go mod verify` checks cache integrity

### Cache Structure

```
$GOMODCACHE/
  cache/
    download/    # downloaded .info, .mod, .zip files
  <module-path>@<version>/   # unpacked source
```

## Edge Cases and Gotchas

### `go get` vs `go mod tidy`

- `go get` adjusts a specific dependency (and its dependents)
- `go mod tidy` reconciles all dependencies with the source code

Users should prefer `go get` for add/upgrade/downgrade. `go mod edit` is for tools/scripts that understand the module graph.

### Vendor Mode Auto-Activation

If a `vendor` directory exists and `go.mod` declares `go 1.14+`, the go command acts as if `-mod=vendor` were set. To override, use `-mod=mod` explicitly.

### `-compat` in `go mod tidy`

`-compat=version` preserves checksums needed for an older Go release to load the module. By default, `tidy` uses the version prior to the `go` directive. This affects which `go.sum` entries are kept.

### `go mod edit` Doesn't Resolve

`go mod edit` only modifies `go.mod` text — it does not download modules or check that versions exist. Use `go get` for real dependency management.

### Replace with Local Path

A local `replace` must point to a directory containing a `go.mod` file, not a module path:

```
replace example.com/foo => ../foo    # ../foo must contain go.mod
```

### Module Path Conventions

Module paths should start with a unique prefix (domain name). Common patterns:
- `github.com/user/repo` — GitHub-hosted
- `example.com/project` — custom domain
- `golang.org/x/...` — Go extended packages

Major version changes require a path suffix (`/v2`, `/v3`) for modules at v2+.

## References

- `go help mod` — `go mod` subcommand list
- `go help modules` — module system overview
- `go help go.mod` — `go.mod` file format
- `go help module-auth` — checksum verification
- `go help goproxy` — module proxy protocol
- `go help private` — private module configuration
- `go help vcs` — VCS restrictions
- <https://go.dev/ref/mod> — Go modules reference (comprehensive)

## Related

- [concepts.md](concepts.md) — module discovery, vendor directories, package patterns
- [environment.md](environment.md) — `GOPROXY`, `GOSUMDB`, `GOPRIVATE`, `GOMODCACHE`, `GOVCS`, `GOINSECURE`
- [work.md](work.md) — workspaces (multi-module builds)
- [cli.md](cli.md) — `go get`, `go install` with version suffixes
