# Core Commands

The primary user-facing `brew` commands: installing, searching, listing, and querying packages. For the formula and cask DSLs, see [formula.md](formula.md) and [cask.md](cask.md). For dependency inspection, see [dependency.md](dependency.md).

## Package Installation and Removal

### `brew install` [options] `<formula|cask>` [...]

Installs a formula or cask. Unless configured otherwise, also checks for outdated dependents and runs cleanup.

| Flag | Purpose |
|------|---------|
| `--formula` | Treat all named arguments as formulae |
| `--cask` | Treat all named arguments as casks |
| `-s, --build-from-source` | Compile from source instead of using a bottle |
| `--force-bottle` | Install from a bottle even if it wouldn't normally be used |
| `--HEAD` | Install the HEAD (main/trunk) version |
| `-d, --debug` | Open interactive debugging on failure |
| `-f, --force` | Skip keg-only/non-migrated checks; overwrite files for casks |
| `-v, --verbose` | Print verification and post-install steps |
| `-n, --dry-run` | Show what would be installed without installing |
| `-y, --no-ask` | Skip confirmation prompts |
| `--ignore-dependencies` | Skip installing dependencies (dev only) |
| `--only-dependencies` | Install only the dependencies, not the formula itself |
| `--cc <compiler>` | Use a specific compiler (e.g. `gcc-9`) |
| `--keep-tmp` | Retain temporary files |
| `--debug-symbols` | Generate debug symbols on build |
| `--build-bottle` | Prepare for eventual bottling |
| `--skip-post-install` | Skip post-install steps |
| `--skip-link` | Skip linking into prefix |
| `--overwrite` | Delete conflicting files in prefix while linking |
| `--adopt` | Adopt identical existing artifacts (casks) |
| `--skip-cask-deps` | Skip cask dependencies |
| `--zap` | For cask reinstall, remove all associated files |
| `--[no-]binaries` | Enable/disable linking helper executables (casks) |
| `--require-sha` | Require checksums for all casks |

When a formula and a cask share the same name, use `--formula` or `--cask` to disambiguate. Without these flags, Homebrew tries formula first, then cask.

### `brew uninstall|remove|rm` [options] `<installed_formula|cask>` [...]

Uninstalls a formula or cask.

| Flag | Purpose |
|------|---------|
| `-f, --force` | Delete all versions; uninstall even if not installed |
| `--zap` | Remove ALL files associated with a cask (may affect shared files) |
| `--ignore-dependencies` | Don't fail if formula is a dependency of others |
| `--formula` / `--cask` | Filter by type |

### `brew reinstall` [options] `<formula|cask>` [...]

Uninstalls then reinstalls using the same original install options. Supports most `brew install` flags including `--build-from-source`, `--cask`, `--zap`, `--adopt`.

### `brew upgrade` [options] [`<installed_formula|cask>` ...]

Upgrades outdated, unpinned packages using original install options. With no arguments, upgrades all outdated packages.

| Flag | Purpose |
|------|---------|
| `--formula` / `--cask` | Upgrade only formulae or only casks |
| `-s, --build-from-source` | Build from source instead of using bottles |
| `--force-bottle` | Use a bottle even if it wouldn't normally be used |
| `--fetch-HEAD` | Check upstream for HEAD version updates |
| `-g, --greedy` | Include `version :latest` and `auto_updates true` casks |
| `--greedy-latest` | Include `version :latest` casks only |
| `--greedy-auto-updates` | Include `auto_updates true` casks only |
| `--minimum-version <version>` | Only upgrade if below minimum |
| `--no-quit` | Don't quit running apps during cask upgrades |
| `--overwrite` | Delete conflicting files |
| `--skip-cask-deps` | Skip cask dependencies |
| `--[no-]binaries` / `--require-sha` | Cask-specific |
| `-n, --dry-run` / `-d, --debug` / `-f, --force` / `-v, --verbose` / `-y, --no-ask` | Standard flags |
| `--keep-tmp` / `--debug-symbols` | Build-related |

### `brew migrate` [options] `<installed_formula|cask>` [...]

Migrates renamed packages to new names. Flags: `-f, --force`, `-n, --dry-run`, `--formula`, `--cask`.

## Update Commands

### `brew update|up` [options]

Fetches the newest version of Homebrew and all formulae from GitHub. In API mode (default since 4.0), updates the JSON formula/cask cache. In git-clone mode, pulls all tapped repositories.

| Flag | Purpose |
|------|---------|
| `--auto-update` | Run on auto-updates (skips slower steps) |
| `-f, --force` | Always do a full update check |
| `-v, --verbose` | Print directories checked and git operations |
| `-d, --debug` | Trace all shell commands |

### `brew update-if-needed`

Runs `brew update --auto-update` only if needed. Fast no-op alternative for scripts.

### `brew update-reset` [`<repository>` ...]

Fetches and resets Homebrew and all taps to latest `origin/HEAD` — destroys all uncommitted changes.

## Search and Query

### `brew search|-S` [options] `<text>|/<regex>/` [...]

Substring search of cask tokens and formula names. Flank with `/` for regex.

| Flag | Purpose |
|------|---------|
| `--formula` / `--cask` | Limit search to formulae or casks |
| `--desc` | Search descriptions instead of names |
| `--pull-request` | Search GitHub PRs |
| `--open` / `--closed` | Filter PRs by status |
| `--alpine` / `--repology` / `--macports` / `--fink` / `--opensuse` / `--fedora` / `--archlinux` / `--debian` / `--ubuntu` | Search external package databases |

### `brew info|abv` [options] [`<formula|cask>` ...]

Displays statistics for Homebrew installation or info about specific formulae/casks.

| Flag | Purpose |
|------|---------|
| `--analytics` | List analytics data |
| `--days <30|90|365>` | Analytics timeframe |
| `--category <install|install-on-request|build-error|cask-install|os-version>` | Analytics category |
| `--github` | Open GitHub source page |
| `--json` | Print as JSON (`v1` for formulae, `v2` for both) |
| `--installed` | Inventory of installed packages |
| `--variations` | Include variations hash in JSON |
| `--sizes` | Show size of installed packages |
| `-v, --verbose` / `--formula` / `--cask` | Standard flags |

### `brew desc` [options] `<formula|cask|text|/regex/>` [...]

Displays formula/cask name and one-line description.

| Flag | Purpose |
|------|---------|
| `-s, --search` | Search both names and descriptions |
| `-n, --name` | Search names only |
| `-d, --description` | Search descriptions only |
| `--formula` / `--cask` | Filter by type |

### `brew home|homepage` [--formula|--cask] [`<formula|cask>` ...]

Opens a formula/cask's homepage in a browser. With no arguments, opens `https://brew.sh`.

### `brew options` [options] [`<formula>` ...]

Shows install options specific to a formula. Flags: `--compact`, `--installed`.

### `brew fetch` [options] `<formula|cask>` [...]

Downloads bottle or source packages without installing.

| Flag | Purpose |
|------|---------|
| `--os <os>` / `--arch <arch>` | Download for specific platforms |
| `--all-platforms` | Download for all platforms |
| `--bottle-tag <tag>` | Specific bottle tag |
| `--HEAD` | Fetch HEAD version |
| `-f, --force` | Re-fetch even if already cached |
| `-v, --verbose` / `--retry` | Output/control flags |
| `--deps` | Also download dependencies |
| `-s, --build-from-source` / `--build-bottle` / `--force-bottle` | Build mode |
| `--formula` / `--cask` | Filter by type |

## Listing

### `brew list|ls` [options] [`<installed_formula|cask>` ...]

Lists all installed formulae and casks, or files within a specific keg.

| Flag | Purpose |
|------|---------|
| `--formula` / `--cask` | List only formulae or only casks |
| `--full-name` | Fully-qualified names (tap/name) |
| `--versions` | Show version numbers |
| `--json` | JSON output (requires `--versions`) |
| `--multiple` | Only formulae with multiple versions |
| `--pinned` | Only pinned packages |
| `--installed-on-request` | Only manually installed |
| `--no-installed-on-request` | Only dependency-installed |
| `--poured-from-bottle` | Only bottle-installed |
| `--built-from-source` | Only source-built |
| `-1` / `-l` / `-r` / `-t` | Passed to `ls(1)` for formatting/sorting |

### `brew formulae`

Lists all locally installable formulae including short names.

### `brew casks`

Lists all locally installable casks including short names.

### `brew leaves` [options]

Lists installed formulae that are not dependencies of another installed formula.

| Flag | Purpose |
|------|---------|
| `-r, --installed-on-request` | Only manually installed leaves |
| `-p, --installed-as-dependency` | Only dependency-installed leaves |

## Path and Metadata Queries

These are invoked as flags on the `brew` command itself (`brew --prefix`, not `brew prefix`).

| Command | Purpose |
|---------|---------|
| `brew --version` / `brew -v` | Print version numbers of Homebrew and core taps |
| `brew --prefix` [`<formula>`] | Display install prefix (e.g. `/opt/homebrew`); with formula, shows its Cellar path |
| `brew --prefix --unbrewed` | List unbrewed header files in prefix |
| `brew --prefix --installed` | Only show prefix if formula is installed |
| `brew --cellar` [`<formula>`] | Display Cellar path |
| `brew --caskroom` [`<cask>`] | Display Caskroom path |
| `brew --cache` [options] [`<formula|cask>`] | Display download cache path |
| `brew --repository` / `brew --repo` [`<tap>`] | Display Homebrew Git repository location |
| `brew --taps` | Display Taps directory path |
| `brew --env` / `brew --environment` [`<formula>`] | Summarise build environment |
| `brew --config` | Show Homebrew and system configuration info |

### `brew shellenv` [`<shell>` ...]

Prints export statements to add Homebrew to `PATH`, `MANPATH`, and `INFOPATH`. This is the standard way to set up brew in a shell profile.

Supported shells: `bash`, `csh`, `fish`, `pwsh`, `sh`, `tcsh`, `zsh`.

```sh
eval "$(/opt/homebrew/bin/brew shellenv)"
```

## Utility Commands

| Command | Purpose |
|---------|---------|
| `brew command <cmd>` | Display the path to the file used when invoking `brew <cmd>` |
| `brew commands` [--quiet] [--include-aliases] | List built-in and external commands |
| `brew help` [`<command>`] | Show usage instructions for a brew command |
| `brew which-formula` [--explain] `<command>` | Show which formula provides a given command |
| `brew source` [`<formula>`] | Open a formula's source repository in a browser |
| `brew docs` | Open Homebrew's online documentation in a browser |
| `brew cat` [--formula\|--cask] `<formula\|cask>` | Display the source of a formula or cask |
| `brew formula` `<formula>` | Display the path where a formula is located |
| `brew log` [options] [`<formula\|cask>`] | Show git log for a formula/cask or the Homebrew repository |
| `brew completions` [subcommand] | Control automatic linking of external tap shell completions |

### `brew completions`

Controls automatic linking of external tap shell completions.

| Subcommand | Purpose |
|------------|---------|
| `brew completions` / `brew completions state` | Show current state |
| `brew completions link` | Link external tap completions |
| `brew completions unlink` | Unlink external tap completions |

## Global Options

These flags apply to most subcommands:

| Flag | Purpose |
|------|---------|
| `-d, --debug` | Display debugging information |
| `-q, --quiet` | Make output more quiet |
| `-v, --verbose` | Make output more verbose |
| `-h, --help` | Show help message |

## Command Abbreviations

Many commands have short aliases:

| Alias | Full Command |
|-------|-------------|
| `rm` / `remove` | `uninstall` |
| `up` | `update` |
| `abv` | `info` |
| `ls` | `list` |
| `ln` | `link` |
| `lc` | `livecheck` |
| `tc` | `typecheck` |
| `x` | `exec` |
| `-S` | `search` |

## References

- [Homebrew Manpage](https://docs.brew.sh/Manpage) — official command reference
- For formula DSL: [formula.md](formula.md)
- For cask DSL: [cask.md](cask.md)
- For dependency commands: [dependency.md](dependency.md)
- For cleanup/doctor: [maintenance.md](maintenance.md)
