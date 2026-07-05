# Go Work — Workspaces and Multi-Module Builds

Reference for `go work` — the workspace system that allows working with multiple local modules simultaneously. Covers the `go.work` file format, workspace subcommands, and how workspaces interact with builds.

> **Source of truth**: `go help work`, `go help work init`, `go help work use`, `go help work sync`, `go help work edit`, `go help work vendor`. For **modules** (single-module context), see [mod.md](mod.md). For **`GOWORK`**, see [environment.md](environment.md).

## Overview

A **workspace** is a set of local modules used together as the main modules for builds. Workspaces are defined by a `go.work` file and allow developing across multiple modules without publishing versions or editing each module's `go.mod`.

Support for workspaces is built into many go commands, not just `go work`. When a `go.work` file is found, all `use`-d modules become main modules.

### When to Use Workspaces

| Scenario | Without Workspace | With Workspace |
|----------|-------------------|----------------|
| Develop two local modules together | Use `replace` in each `go.mod` pointing to local paths | Add both to `go.work` with `use` |
| Test a bugfix in a dependency | Fork, `replace` to local fork, rebuild | Add local fork to workspace |
| Monorepo with multiple modules | Complex `replace` chains | Single `go.work` with `use` directives |

## The `go.work` File

A `go.work` file is line-oriented, with each line holding a single directive (keyword + arguments):

```
go 1.18

use ../foo/bar
use ./baz

replace example.com/foo v1.2.3 => example.com/bar v1.4.5
```

The leading keyword can be factored out of adjacent lines to create a block (like Go imports):

```
use (
  ../foo/bar
  ./baz
)
```

### Directives

| Directive | Description |
|-----------|-------------|
| `go <version>` | Go version the file was written at |
| `use <dir>` | Add a module directory to the workspace |
| `replace <old> => <new>` | Replace a module (same syntax as `go.mod` replace) |
| `toolchain <name>` | Go toolchain to use |
| `godebug <key>=<value>` | GODEBUG key-value pair |

### Replace Precedence

`replace` directives in `go.work` take precedence over `replace` directives in any module's `go.mod`. This is primarily intended to override conflicting replaces in different workspace modules.

### Use Directive

The `use` directive specifies a module to include in the workspace. The argument is a **directory path** containing the module's `go.mod` file:

```
use ./my-module      # relative path
use /abs/path/mod    # absolute path
```

All `use`-d modules become **main modules** — their `replace` and `exclude` directives are all honored.

## Workspace Discovery

The go command searches for `go.work` in the current directory and then parent directories (when `GOWORK=auto`, the default). Use `go env GOWORK` to see the active workspace file:

```bash
go env GOWORK    # /path/to/go.work or empty if no workspace
```

### Disabling Workspace Mode

```bash
GOWORK=off go build ./...    # ignore go.work for this command
```

### Empty Workspace

A `go.work` with no `use` directives cannot be used for builds from local modules. `go work init` with no arguments creates an empty workspace (useful as a starting point).

## `go work` Subcommands

### `go work init`

```
go work init [moddirs]
```

Creates a new `go.work` in the current directory. Optionally accepts module directory paths as arguments, which become `use` directives.

```bash
go work init ./foo ./bar    # create workspace with foo and bar
go work init                # create empty workspace
```

The current Go version is listed in the `go` directive.

### `go work use`

```
go work use [-r] [moddirs]
```

Adds or removes modules from `go.work`:

- If a directory exists and contains a `go.mod`, a `use` directive is **added**
- If a directory does not exist, its `use` directive is **removed**
- Fails if any remaining `use` directives refer to non-existent modules

| Flag | Description |
|------|-------------|
| `-r` | Search recursively for modules in the argument directories |

```bash
go work use ./new-module          # add a module
go work use -r ./monorepo         # add all modules found recursively
go work use ./removed-module      # if dir doesn't exist, removes the use directive
```

With no arguments, `go work use` updates the `go` line in `go.work` to be at least as new as all used modules' `go` lines.

### `go work sync`

```
go work sync
```

Syncs the workspace's build list back to the workspace's modules.

The workspace's **build list** is the set of versions of all transitive dependency modules used for builds. `go work sync`:

1. Computes the build list using Minimal Version Selection (MVS)
2. Upgrades each dependency module in each workspace module to the build list's version (if not already the same)

MVS guarantees the build list's version is always ≥ the version in each workspace module, so `sync` only **upgrades** dependencies.

See [mod.md](mod.md) for MVS details.

### `go work edit`

```
go work edit [editing flags] [go.work]
```

Command-line interface for editing `go.work`, primarily for tools/scripts. Only reads `go.work` — does not look up module information. If no file is specified, searches for `go.work` in the current directory and parents.

| Editing Flag | Description |
|--------------|-------------|
| `-fmt` | Reformat `go.work` without other changes |
| `-use=path` | Add a use directive |
| `-dropuse=path` | Drop a use directive |
| `-replace=old[@v]=new[@v]` | Add/replace a replacement |
| `-dropreplace=old[@v]` | Drop a replacement |
| `-go=version` | Set the Go language version |
| `-toolchain=name` | Set the Go toolchain |
| `-godebug=key=value` | Add/replace a godebug line |
| `-dropgodebug=key` | Drop a godebug line |
| `-print` | Print final go.work text instead of writing back |
| `-json` | Print final go.work as JSON |

Editing flags may be **repeated**, with changes applied in order.

### `go work edit -json` Output

```go
type GoWork struct {
    Go        string
    Toolchain string
    Godebug   []Godebug
    Use       []Use
    Replace   []Replace
}

type Use struct {
    DiskPath   string
    ModulePath string
}

type Replace struct {
    Old Module
    New Module
}

type Module struct {
    Path    string
    Version string
}
```

### `go work vendor`

```
go work vendor [-e] [-v] [-o outdir]
```

Resets the workspace's vendor directory to include all packages needed to build and test all workspace packages. Does not include test code for vendored packages.

| Flag | Description |
|------|-------------|
| `-v` | Print names of vendored modules and packages |
| `-e` | Proceed despite errors loading packages |
| `-o outdir` | Create vendor directory at given path (default `vendor`). Only `vendor` at module root is used by the go command. |

## How Workspaces Affect Builds

When a workspace is active:

1. All `use`-d modules become **main modules** simultaneously
2. The **build list** is computed from all main modules' requirements using MVS
3. `replace` directives from `go.work` take precedence over those in any `go.mod`
4. `exclude` directives from all main modules are honored
5. Builds can reference packages from any main module

### Interaction with `go.mod`

| Aspect | Without Workspace | With Workspace |
|--------|-------------------|----------------|
| Main module | One (the one with `go.mod` in scope) | All `use`-d modules |
| `replace` precedence | `go.mod` replaces | `go.work` replaces override `go.mod` replaces |
| `go get` | Modifies `go.mod` | Modifies `go.mod` in the relevant module |
| `go mod tidy` | Per module | Per module (workspaces don't change this) |

### `go get` in Workspaces

`go get` in a workspace modifies the `go.mod` of the module containing the current directory (not the `go.work`). The workspace's build list is then recomputed.

### `go mod tidy` in Workspaces

`go mod tidy` operates on individual modules, not the workspace as a whole. Run it in each module directory. The `-compat` flag uses the module's `go` directive, not the workspace's.

## Edge Cases and Gotchas

### `GOWORK=off` Disables Workspace

Set `GOWORK=off` to ignore the workspace for a single command. This is useful when you want to build as if the workspace doesn't exist:

```bash
GOWORK=off go build ./...
```

### Empty Workspace Can't Build

A `go.work` with no `use` directives cannot be used for builds from local modules. `go work init` with no arguments creates an empty workspace — add modules with `go work use`.

### `go work sync` Only Upgrades

`go work sync` can only **upgrade** dependencies (never downgrade) because MVS guarantees the build list version is ≥ the version in each module. To downgrade, use `go get` in the individual module.

### `go work use` Removes Non-Existent Directories

If you pass a directory path to `go work use` that doesn't exist, the corresponding `use` directive is **removed** from `go.work`. This is a convenient way to remove modules:

```bash
# Remove ./old-module from workspace
go work use ./old-module    # if ./old-module doesn't exist, it's removed
```

But this means accidentally passing a typo can remove an intended module.

### Workspace `replace` vs `go.mod` `replace`

A `replace` in `go.work` applies to all modules in the workspace and overrides any conflicting `replace` in their `go.mod` files. This is intentional for resolving conflicts, but can cause surprises if you forget a workspace replace exists.

### `go.work` Is Not Published

The `go.work` file should generally **not** be committed to version control for libraries. It's a local development artifact. For applications (not importable by others), committing `go.work` may be acceptable.

## References

- `go help work` — workspace overview and subcommand list
- `go help work init` — init details
- `go help work use` — use details
- `go help work sync` — sync details
- `go help work edit` — edit details
- `go help work vendor` — vendor details
- <https://go.dev/ref/mod#workspaces> — workspace reference
- <https://go.dev/doc/tutorial/workspaces> — workspace tutorial

## Related

- [mod.md](mod.md) — module system, `go.mod` format, MVS, `replace` directives
- [environment.md](environment.md) — `GOWORK`, `GOTOOLCHAIN`
- [concepts.md](concepts.md) — module discovery, main modules
