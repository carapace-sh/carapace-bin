# Maintenance and Diagnostics

Commands for cleaning up, diagnosing problems, managing package state (pinning, outdated), and inspecting installed packages.

## Cleanup

### `brew cleanup` [options] [`<formula|cask>` ...]

Removes stale lock files, outdated downloads, and old versions from the Cellar.

| Flag | Purpose |
|------|---------|
| `--prune <days>` | Remove cache files older than N days |
| `--prune=all` | Remove all cache files |
| `-n, --dry-run` | Show what would be removed |
| `-s, --scrub` | Scrub cache including latest versions |
| `--prune-prefix` | Only prune symlinks from prefix |

`brew install`, `brew upgrade`, and `brew reinstall` automatically run cleanup every `HOMEBREW_CLEANUP_PERIODIC_FULL_DAYS` (default 30) days.

### `brew autoremove` [--dry-run]

Uninstalls formulae that were only installed as dependencies and are no longer needed.

| Flag | Purpose |
|------|---------|
| `-n, --dry-run` | List what would be removed without removing |

Automatic autoremove during `brew cleanup` and `brew uninstall` is disabled by `HOMEBREW_NO_AUTOREMOVE`. The `brew autoremove` command itself can still be run manually regardless.

## Diagnostics

### `brew doctor`

Checks system for potential problems. Exits non-zero if problems found.

| Flag | Purpose |
|------|---------|
| `--list-checks` | List all audit methods |
| `-D, --audit-debug` | Enable debugging and profiling |

### `brew missing` [--hide=`<hidden>`] [`<formula|cask>` ...]

Checks for missing dependencies in installed formulae. The `--hide` flag ignores specific formulae during the check.

## Package State Management

### `brew pin` [--formula|--cask] `<installed_formula|cask>` [...]

Pins a package to prevent it from being upgraded. Pinned packages are skipped by `brew upgrade`.

### `brew unpin` [--formula|--cask] `<installed_formula|cask>` [...]

Unpins a package to allow upgrades.

### `brew outdated` [options] [`<formula|cask>` ...]

Lists installed packages with newer versions available.

| Flag | Purpose |
|------|---------|
| `-q, --quiet` | Quiet output (just names) |
| `-v, --verbose` | Verbose output |
| `--formula` / `--cask` | Filter by type |
| `--json` | JSON output (`v2`; `v1` is no longer supported) |
| `--minimum-version <version>` | Filter by minimum version |
| `--fetch-HEAD` | Check HEAD versions upstream |
| `-g, --greedy` | Include `version :latest` and `auto_updates true` casks |
| `--greedy-latest` | Include `version :latest` casks only |
| `--greedy-auto-updates` | Include `auto_updates true` casks only |

### `brew migrate` [options] `<installed_formula|cask>` [...]

Migrates renamed packages to new names. Flags: `-f, --force`, `-n, --dry-run`, `--formula`, `--cask`.

## Inspection Commands

### `brew leaves` [options]

Lists installed formulae that are not dependencies of another installed formula — the "top-level" manually installed packages.

| Flag | Purpose |
|------|---------|
| `-r, --installed-on-request` | Only manually installed leaves |
| `-p, --installed-as-dependency` | Only dependency-installed leaves |

### `brew linkage` [options] [`<installed_formula>` ...]

Checks library links from formula kegs. Useful for detecting broken dependencies after upgrades.

| Flag | Purpose |
|------|---------|
| `--test` | Test for broken linkage |
| `--strict` | Treat warnings as errors |
| `--reverse` | Show reverse linkage (what depends on this) |
| `--cached` | Use cached linkage data |

### `brew tab` [options] `<installed_formula|cask>` [...]

Edits tab information for installed formulae/casks — controls autoremove behavior.

| Flag | Purpose |
|------|---------|
| `--installed-on-request` | Mark as manually installed |
| `--no-installed-on-request` | Mark as dependency-installed |
| `--formula` / `--cask` | Filter by type |

## Tab Metadata (INSTALL_RECEIPT.json)

Each installed keg has a tab — a JSON file at `Cellar/<formula>/<version>/INSTALL_RECEIPT.json` recording:

- How the formula was installed (source vs bottle)
- Compiler used and version
- Options used during installation
- Whether it was `installed_on_request` (manually) or as a dependency
- Runtime dependencies
- Install time

The tab is used by:
- `brew autoremove` — to determine if a formula was installed as a dependency
- `brew leaves` — to filter manually-installed vs dependency-installed
- `brew reinstall` — to reproduce original install options
- `brew list --installed-on-request` / `--poured-from-bottle` / `--built-from-source`

## Linking

### `brew link|ln` [options] `<installed_formula>` [...]

Symlinks formula files from the Cellar into Homebrew's prefix. Normally done automatically on install, but can be run manually (e.g., after `brew unlink`).

| Flag | Purpose |
|------|---------|
| `--overwrite` | Delete existing conflicting files |
| `-n, --dry-run` | List files to be linked/deleted |
| `-f, --force` | Allow linking keg-only formulae |
| `--HEAD` | Link HEAD version |

### `brew unlink` [--dry-run] `<installed_formula>` [...]

Removes symlinks for a formula from Homebrew's prefix — temporarily disables the formula without uninstalling it.

| Flag | Purpose |
|------|---------|
| `-n, --dry-run` | List files that would be unlinked |

See [cellar.md](cellar.md) for details on how keg linking works.

## Environment Variables for Maintenance

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_CLEANUP_MAX_AGE_DAYS` | `120` | Max age for cached files |
| `HOMEBREW_CLEANUP_PERIODIC_FULL_DAYS` | `30` | Full cleanup interval |
| `HOMEBREW_NO_INSTALL_CLEANUP` | — | Never auto-cleanup |
| `HOMEBREW_NO_CLEANUP_FORMULAE` | — | Formulae exempt from cleanup |
| `HOMEBREW_NO_AUTOREMOVE` | — | Don't auto-remove unused dependents |
| `HOMEBREW_NO_INSTALLED_DEPENDENTS_CHECK` | — | Don't check dependents after install |

## References

- [Homebrew Manpage](https://docs.brew.sh/Manpage)
- For the Cellar: [cellar.md](cellar.md)
- For environment variables: [environment.md](environment.md)
- For dependency commands: [dependency.md](dependency.md)
