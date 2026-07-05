# Developer Commands

Commands for formula/cask authors and Homebrew contributors: auditing, style checking, testing, version bumping, and CI tooling.

## Code Quality

### `brew audit` [options] [`<formula|cask>` ...]

Checks formulae/casks for Homebrew coding style violations and common issues.

| Flag | Purpose |
|------|---------|
| `--strict` | Run additional, stricter style checks |
| `--git` | Check git history for issues |
| `--online` | Check URLs and other online resources |
| `--new` | Check if new formula/cask is eligible for homebrew-core |
| `--fix` | Auto-fix violations |
| `--only <method>` | Run only specific audit methods |
| `--except <method>` | Skip specific audit methods |
| `--only-cops <cops>` | Run only specific RuboCop cops |
| `--except-cops <cops>` | Skip specific RuboCop cops |
| `--formula` / `--cask` | Filter by type |
| `--tap <user/repo>` | Audit all formulae in a tap |
| `--installed` | Audit all installed formulae |
| `--changed` | Audit only changed formulae |

### `brew style` [options] [`<file|tap|formula|cask>` ...]

Checks conformance to Homebrew style guidelines using RuboCop.

| Flag | Purpose |
|------|---------|
| `--fix` | Auto-fix style violations |
| `--todo` | Add `rubocop:todo` comments for remaining violations (requires `--fix`) |
| `--reset-cache` | Reset RuboCop cache |
| `--changed` | Only check changed files |
| `--formula` / `--cask` | Filter by type |
| `--only-cops <cops>` | Run only specific cops |
| `--except-cops <cops>` | Skip specific cops |

### `brew typecheck|tc` [options] [`<tap>` ...]

Checks for typechecking errors using Sorbet.

| Flag | Purpose |
|------|---------|
| `--fix` | Auto-fix type errors |
| `-q, --quiet` | Quiet output |
| `--update` | Update RBI files |
| `--update-all` | Update all RBI files |
| `--suggest-typed` | Suggest `typed` levels |
| `--lsp` | Start LSP server |
| `--dir <path>` / `--file <path>` | Target specific paths |
| `--ignore <pattern>` | Ignore files |

Related: `HOMEBREW_SORBET_RUNTIME` and `HOMEBREW_SORBET_RECURSIVE` environment variables.

### `brew rubocop`

Installs, configures, and runs Homebrew's RuboCop.

## Testing

### `brew test` [options] `<installed_formula>` [...]

Runs the `test do` block provided by an installed formula.

| Flag | Purpose |
|------|---------|
| `-f, --force` | Test formulae even if they are unlinked |
| `--HEAD` | Test HEAD version |
| `--keep-tmp` | Retain temporary files |
| `--retry` | Retry on failure |

### `brew tests` [options]

Runs Homebrew's unit and integration tests.

| Flag | Purpose |
|------|---------|
| `--coverage` | Generate coverage report |
| `--generic` | Run only OS-agnostic tests |
| `--online` | Run online tests |
| `--debug` | Debug mode |
| `--changed` | Only run changed tests |
| `--fail-fast` | Stop on first failure |
| `--no-parallel` | Run tests serially |
| `--only <script>` | Run only specific test script |
| `--seed <value>` | Set random seed |
| `--profile <n>` | Profile N tests |
| `--stackprof` / `--vernier` / `--ruby-prof` | Use specific profiler |

### `brew test-bot` [options] [`<formula>`]

Tests the full lifecycle of a Homebrew change (CI). This is the primary CI tool used by Homebrew's GitHub Actions.

| Flag | Purpose |
|------|---------|
| `--dry-run` | Show what would be done |
| `--cleanup` | Clean up after |
| `--skip-setup` | Skip setup steps |
| `--build-from-source` | Build all from source |
| `--build-dependents-from-source` | Build dependents from source |
| `--junit` | Generate JUnit XML |
| `--keep-old` | Keep old bottles |
| `--skip-relocation` | Skip bottle relocation |
| `--only-json-tab` | Only JSON tab |
| `--local` | Run locally |
| `--tap <tap>` | Test specific tap |
| `--fail-fast` | Stop on first failure |
| `-v, --verbose` | Verbose |
| `--publish` | Publish bottles |
| `--root-url <url>` | Bottle root URL |
| `--git-name <name>` / `--git-email <email>` | Git identity |
| `--skip-*` / `--only-*` | Various step selection flags |

## Version Management

### `brew livecheck|lc` [options] [`<formula|cask>` ...]

Checks for newer versions from upstream sources.

| Flag | Purpose |
|------|---------|
| `--full-name` | Show fully-qualified names |
| `--tap <tap>` | Check specific tap |
| `--installed` | Only check installed formulae |
| `--newer-only` | Only show newer versions |
| `--json` | JSON output |
| `-r, --resources` | Check resources too |
| `-q, --quiet` | Quiet output |
| `--formula` / `--cask` | Filter by type |
| `--extract-plist` | Use extract-plist strategy |
| `--autobump` | Include autobumped packages |

Related: `HOMEBREW_LIVECHECK_AUTOBUMP` and `HOMEBREW_LIVECHECK_WATCHLIST`.

### `brew bump` [options] [`<formula|cask>` ...]

Displays out-of-date packages and latest versions.

| Flag | Purpose |
|------|---------|
| `--full-name` | Fully-qualified names |
| `--no-pull-requests` | Don't show PR status |
| `--no-autobump` | Exclude autobumped packages |
| `--formula` / `--cask` / `--tap` | Filter |
| `--installed` | Only installed packages |
| `--open-pr` | Open a PR for out-of-date packages |
| `--no-fork` | Don't fork the repository |
| `--start-with <letter>` | Start from specific letter |
| `--bump-synced` | Bump additional formulae marked as synced with the given formulae |

### `brew bump-formula-pr` [options] [`<formula>`]

Creates a PR to update a formula's URL or tag.

| Flag | Purpose |
|------|---------|
| `-n, --dry-run` | Show what would be done |
| `--write-only` | Write without committing |
| `--commit` | Commit changes |
| `--no-audit` | Skip audit |
| `--strict` / `--online` | Additional checks |
| `--no-browse` / `--no-fork` | PR options |
| `--url <url>` | New URL |
| `--sha256 <hash>` | New SHA-256 |
| `--tag <tag>` | New git tag |
| `--revision <rev>` | New git revision |
| `--version <ver>` | New version |
| `--mirror <url>` | Mirror URL |
| `--fork-org <org>` | Fork organization |
| `--message <msg>` | PR message |
| `--install-dependencies` | Install deps before building |
| `--python-package-name` / `--python-extra-packages` / `--python-exclude-packages` | Python resource options |

### `brew bump-cask-pr` [options] `<cask>`

Creates a PR to update a cask version.

| Flag | Purpose |
|------|---------|
| `-n, --dry-run` / `--write-only` / `--commit` | Execution control |
| `--no-audit` / `--no-style` / `--no-browse` / `--no-fork` | Skip checks |
| `--version <ver>` | New version |
| `--version-arm <ver>` / `--version-intel <ver>` | Architecture-specific versions |
| `--url <url>` | New URL |
| `--sha256 <hash>` | New SHA-256 |
| `--message <msg>` / `--fork-org <org>` | PR options |

### `brew bump-revision` [options] `<formula>` [...]

Increments the `revision` of a formula.

| Flag | Purpose |
|------|---------|
| `-n, --dry-run` | Show what would be done |
| `--remove-bottle-block` | Remove bottle block |
| `--write-only` | Write without committing |
| `--message <msg>` | Commit message |

### `brew bump-unversioned-casks` [options] `<cask|tap>` [...]

Checks unversioned URL casks for updates.

| Flag | Purpose |
|------|---------|
| `-n, --dry-run` | Show what would be done |
| `--limit <minutes>` | Time limit |
| `--state-file <path>` | State file |

### `brew bump-compatibility-version` [options] `<formula>` [...]

Increments the `compatibility_version` of a formula.

## Formula Creation and Editing

### `brew create` [options] `<URL>`

Generates a formula (or cask with `--cask`) from a URL.

| Flag | Purpose |
|------|---------|
| `--autotools` / `--cabal` / `--cmake` / `--crystal` / `--go` / `--meson` / `--node` / `--perl` / `--python` / `--ruby` / `--rust` / `--zig` | Build system template |
| `--cask` | Generate a cask instead of a formula |
| `--no-fetch` | Don't fetch the URL |
| `--HEAD` | Generate HEAD version |
| `--set-name <name>` / `--set-version <ver>` / `--set-license <lic>` | Set metadata |
| `--tap <user/repo>` | Target tap |
| `-f, --force` | Overwrite existing |

### `brew edit` [options] [`<formula|cask|tap>` ...]

Opens a formula/cask/tap in the editor (`HOMEBREW_EDITOR`).

| Flag | Purpose |
|------|---------|
| `--formula` / `--cask` | Filter by type |
| `--print-path` | Print path instead of opening |

### `brew cat` [--formula|--cask] `<formula|cask>` [...]

Displays the source of a formula or cask. Uses `bat` if `HOMEBREW_BAT` is set.

### `brew extract` [options] `<formula> <tap>`

Extracts a historic version of a formula into a tap.

| Flag | Purpose |
|------|---------|
| `--version <ver>` | Specific version to extract |
| `--git-revision <rev>` | Git revision to extract from |
| `-f, --force` | Overwrite existing |

### `brew unpack` [options] `<formula|cask>` [...]

Unpacks formula/cask source files into subdirectories.

| Flag | Purpose |
|------|---------|
| `--destdir <path>` | Destination directory |
| `--patch` | Apply patches |
| `-g, --git` | Initialize git in unpacked dir |
| `-f, --force` | Overwrite existing |

## Ruby Environment

### `brew irb` [--examples] [--pry]

Enters the interactive Homebrew Ruby shell. Uses Pry if `HOMEBREW_PRY` is set.

### `brew ruby` [options] (-e `<text>` | `<file>`)

Runs a Ruby instance with Homebrew's libraries loaded.

| Flag | Purpose |
|------|---------|
| `-r <file>` | Require a file |
| `-e <text>` | Execute Ruby code |

### `brew sh` [options] [`<file>`]

Enters an interactive Homebrew build environment shell.

| Flag | Purpose |
|------|---------|
| `-r, --ruby` | Include Ruby environment |
| `--env` | Show environment |
| `-c, --cmd <cmd>` | Run command instead of interactive shell |

### `brew prof` [--stackprof] [--vernier] `<command>` [...]

Runs Homebrew with a Ruby profiler.

### `brew exec|x` [--formulae=`<formulae>`] [--sandbox=`<path>`] [--deny-network] [--] `<command>` [args...]

Runs a command in an environment populated by Homebrew formulae.

## CI and Release Tooling

### `brew readall` [options] [`<tap>` ...]

Imports all items from taps for debugging/validation.

| Flag | Purpose |
|------|---------|
| `--os <os>` / `--arch <arch>` | Target platform |
| `--aliases` | Check aliases |
| `--syntax` | Syntax check only |
| `--no-simulate` | Don't simulate installation |

### `brew generate-man-completions` [--no-exit-code]

Generates manpages and shell completions from the brew source.

### `brew generate-zap` [--name] `<cask_or_name>`

Generates a `zap` stanza for a cask by scanning for associated files.

### `brew vendor-gems` [--update=] [--no-commit]

Installs and commits vendored gems.

### `brew update-python-resources` [options] `<formula>` [...]

Updates PyPI resource blocks in a formula.

| Flag | Purpose |
|------|---------|
| `-p, --print-only` | Print without writing |
| `--ignore-errors` | Continue on errors |
| `--install-dependencies` | Install deps first |
| `--version <ver>` | Specific version |
| `--package-name <name>` | PyPI package name |
| `--extra-packages <pkgs>` / `--exclude-packages <pkgs>` | Package filtering |

### `brew update-perl-resources` [options] `<formula>` [...]

Updates CPAN resource blocks. Flags: `-p, --print-only`, `--ignore-errors`.

### `brew update-test` [options]

Runs a test of `brew update`. Flags: `--to-tag`, `--keep-tmp`, `--commit`, `--before <date>`.

## Developer Mode

### `brew developer` [subcommand]

Controls Homebrew's developer mode.

| Subcommand | Purpose |
|------------|---------|
| `brew developer` / `brew developer state` | Show current state |
| `brew developer on` | Enable developer mode |
| `brew developer off` | Disable developer mode |

When enabled, sets `HOMEBREW_DEVELOPER` — turns warnings into errors, uses faster auto-update intervals, and unblocks developer-only commands.

### `brew lgtm` [--online]

Runs typecheck, style, tests, audit, and test checks in one go. "Looks Good To Me" — a pre-PR sanity check.

### `brew debugger` [--open] `<command>` [...]

Runs a Homebrew command in debug mode.

### `brew gist-logs` [--new-issue] [--private] `<formula>`

Uploads build failure logs to a new GitHub Gist.

| Flag | Purpose |
|------|---------|
| `-n, --new-issue` | Auto-create a GitHub issue |
| `-p, --private` | Create a private Gist |

## Aliases

### `brew alias` [--edit] [`<alias>` | `<alias>=<command>`]

Shows or edits command aliases.

### `brew unalias` `<alias>` [...]

Removes aliases.

## Miscellaneous

| Command | Purpose |
|---------|---------|
| `brew contributions` [options] | Summarise contributions to Homebrew repos |
| `brew rubydoc` [--only-public] [--open] | Generate Homebrew's RubyDoc documentation |
| `brew nodenv-sync` | Symlink installed Node versions into `~/.nodenv/versions` |
| `brew pyenv-sync` | Symlink installed Python versions into `~/.pyenv/versions` |
| `brew rbenv-sync` | Symlink installed Ruby versions into `~/.rbenv/versions` |
| `brew install-bundler-gems` [--groups=] [--add-groups=] | Install Homebrew's Bundler gems |
| `brew setup-ruby` [`<command>`] | Install and configure Homebrew's Ruby |
| `brew setup-sandbox` | Set up Homebrew sandbox (requires sudo, Linux only) |
| `brew sandbox-exec` [--deny-network] `<writable-path>` [--] `<command>` [args...] | Run in Homebrew sandbox |
| `brew mcp-server` [--debug] | Start Homebrew MCP server |

## References

- [Homebrew Manpage](https://docs.brew.sh/Manpage)
- [Formula Cookbook](https://docs.brew.sh/Formula-Cookbook)
- For formula DSL: [formula.md](formula.md)
- For cask DSL: [cask.md](cask.md)
- For bottles: [bottle.md](bottle.md)
- For environment variables: [environment.md](environment.md)
