# Dependency Inspection

Commands for inspecting the dependency graph: forward dependencies (`brew deps`), reverse dependencies (`brew uses`), and related queries.

## `brew deps` [options] [`<formula|cask>` ...]

Shows dependencies for a formula. For installed formulae, shows actual runtime dependencies (not all possible dependencies from the DSL).

| Flag | Purpose |
|------|---------|
| `-n, --topological` | Topological sort order |
| `-1, --direct` | Show direct dependencies only (non-recursive) |
| `--union` | Show union of dependencies across multiple formulae |
| `--full-name` | Show fully-qualified tap/name |
| `--tree` | Show as tree |
| `--graph` | Show as directed graph |
| `--dot` | Show as Graphviz dot format |
| `--annotate` | Mark build/test/implicit/optional/recommended deps |
| `--installed` | Only show installed dependencies |
| `--missing` | Show missing dependencies |
| `--include-build` | Include build dependencies |
| `--include-optional` | Include optional dependencies |
| `--include-test` | Include test dependencies |
| `--include-implicit` | Include implicit dependencies |
| `--skip-recommended` | Skip recommended dependencies |
| `--for-each` | Show deps for each formula separately |
| `--HEAD` | Use HEAD version deps |
| `--os <os>` | Dependencies for specific OS |
| `--arch <arch>` | Dependencies for specific architecture |
| `--formula` / `--cask` | Filter by type |

### Examples

```bash
# Show all dependencies of foo
brew deps foo

# Show direct dependencies only
brew deps --direct foo

# Show as a tree
brew deps --tree foo

# Show what's missing
brew deps --missing foo

# Show deps for multiple formulae as union
brew deps --union foo bar baz

# Annotate dependency types
brew deps --annotate foo
# Output: pkgconf (build), jpeg (runtime), readline (recommended)
```

## `brew uses` [options] `<formula>` [...]

Shows formulae and casks that depend on the given formula(s) — the reverse dependency graph.

| Flag | Purpose |
|------|---------|
| `--recursive` | Include indirect dependents |
| `--installed` | Only show installed formulae |
| `--missing` | Only show formulae with missing dependencies |
| `--include-build` | Include build dependents |
| `--include-optional` | Include optional dependents |
| `--include-test` | Include test dependents |
| `--include-implicit` | Include implicit dependents |
| `--skip-recommended` | Skip recommended dependents |
| `--formula` / `--cask` | Filter by type |

### Examples

```bash
# What formulae depend on openssl@3?
brew uses openssl@3

# What installed formulae depend on openssl@3?
brew uses --installed openssl@3

# Recursive — includes indirect dependents
brew uses --recursive openssl@3

# What's missing a dependency on openssl@3?
brew uses --missing openssl@3
```

## Dependency Types in the Formula DSL

The `depends_on` method in formulae specifies dependency types that affect what `brew deps` shows by default:

| Symbol | Default in `brew deps`? | Description |
|--------|------------------------|-------------|
| _(none)_ | Yes | Runtime dependency |
| `:build` | No (use `--include-build`) | Build-time only |
| `:test` | No (use `--include-test`) | Test-time only |
| `:recommended` | Yes (use `--skip-recommended` to exclude) | Installed by default |
| `:optional` | No (use `--include-optional`) | Opt-in |

See [formula.md](formula.md) for the full `depends_on` DSL reference.

## `brew missing` [--hide=`<hidden>`] [`<formula|cask>` ...]

Checks for missing dependencies in installed formulae. Reports any formula that has a declared dependency which is not installed.

## Dependency-Related Environment Variables

| Variable | Purpose |
|----------|---------|
| `HOMEBREW_NO_INSTALLED_DEPENDENTS_CHECK` | Don't check for broken/outdated dependents after install/upgrade |

## References

- [Homebrew Manpage](https://docs.brew.sh/Manpage)
- For the formula DSL: [formula.md](formula.md)
- For maintenance: [maintenance.md](maintenance.md)
