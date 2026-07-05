# Bundle and Brewfile

`brew bundle` manages dependencies from Homebrew and many other package managers via a declarative `Brewfile`. This document covers the Brewfile format, entry types, options, and all `brew bundle` subcommands.

## Brewfile Format

Brewfiles are Ruby-evaluated files that declare desired system state. Each line specifies a package type and name, with optional arguments. Ruby logic (conditionals, variables) is fully supported.

```ruby
# Brewfile
tap "homebrew/cask-fonts"
brew "openssl@3"
brew "postgresql@16", restart_service: true
cask "firefox"
mas "Refined GitHub", id: 1519867270
vscode "editorconfig.editorconfig"
```

## Entry Types

### `tap` — Homebrew Tap

```ruby
tap "user/tap-repository"
tap "user/tap-repository", "https://user@bitbucket.org/user/homebrew-tap-repository.git"
tap "user/tap-repository", trusted: true
tap "user/tap-repository", trusted: { formula: "foo", cask: "bar" }
```

| Option | Purpose |
|--------|---------|
| Custom URL (2nd arg) | Non-default remote URL |
| `trusted: true` | Trust the entire tap |
| `trusted: { formula:, cask:, command: }` | Trust specific items only |

### `brew` — Homebrew Formula

```ruby
brew "ruby"
brew "postgresql@16", restart_service: true
brew "denji/nginx/nginx-full", link: :overwrite, args: ["with-rmtp"], restart_service: :always
brew "mysql@5.6", restart_service: :changed, link: true, conflicts_with: ["mysql"]
brew "postgresql@16", postinstall: "${HOMEBREW_PREFIX}/opt/postgresql@16/bin/postgres -D ..."
brew "ruby", version_file: ".ruby-version"
brew "gnupg" if OS.mac?
brew "glibc" if OS.linux?
```

| Option | Value | Description |
|--------|-------|-------------|
| `restart_service:` | `true` / `:always` | Restart service after install/upgrade (always) |
| `restart_service:` | `:changed` | Only restart if formula was installed or upgraded |
| `start_service:` | `true` | Start service after install |
| `link:` | `true` | Run `brew link` |
| `link:` | `:overwrite` | Run `brew link --overwrite` |
| `link:` | `false` | Prevent linking |
| `args:` | array | Extra arguments for `brew install` (e.g. `["with-rmtp"]`) |
| `args:` | hash | Extra arguments as key-value (e.g. `{ appdir: "..." }`) |
| `conflicts_with:` | array | Formula names to `brew unlink` |
| `postinstall:` | string | Shell command to run after install/upgrade |
| `version_file:` | string | Write installed version to file (e.g. `.ruby-version`) |
| `trusted:` | `true` | Trust the formula |

### `cask` — macOS Application via Cask

```ruby
cask "firefox"
cask "firefox", args: { appdir: "~/my-apps/Applications" }
cask "opera", greedy: true
cask "google-cloud-sdk", postinstall: "${HOMEBREW_PREFIX}/bin/gcloud components update"
```

| Option | Description |
|--------|-------------|
| `args:` | Hash passed to `brew install --cask` (e.g. `{ appdir:, require_sha: }`) |
| `greedy: true` | Force upgrade of auto-updating/unversioned casks |
| `postinstall:` | Shell command after install/upgrade |
| `trusted: true` | Trust the cask |

### `mas` — Mac App Store

```ruby
mas "Refined GitHub", id: 1519867270
```

| Option | Required | Description |
|--------|----------|-------------|
| `id:` | Yes | Mac App Store item identifier (integer) |

### `vscode` — VS Code Extension

```ruby
vscode "editorconfig.editorconfig"
```

Works with VS Code and its forks/variants.

### `go` — Go Package

```ruby
go "github.com/charmbracelet/crush"
```

### `cargo` — Rust Cargo Package

```ruby
cargo "ripgrep"
```

### `uv` — Python Tool via uv

```ruby
uv "mkdocs"
```

### `flatpak` — Flatpak Package (Linux)

```ruby
flatpak "com.visualstudio.code"
flatpak "org.godotengine.Godot", remote: "flathub-beta", url: "https://dl.flathub.org/beta-repo/"
```

| Option | Description |
|--------|-------------|
| `remote:` | Flatpak remote name |
| `url:` | Repository URL for the remote |

### `krew` — kubectl Plugin via Krew

```ruby
krew "ctx"
```

### `winget` — WinGet Package (Windows/WSL)

```ruby
winget "Steam", id: "Valve.Steam"
winget "PowerToys", id: "XP89DCGQ3K6VLD", source: "msstore"
```

| Option | Required | Description |
|--------|----------|-------------|
| `id:` | Yes | WinGet package identifier |
| `source:` | No | WinGet source (e.g. `"msstore"`) |

### Other Entry Types

| Type | Description |
|------|-------------|
| `npm` | Node.js packages |
| `whalebrew` | Docker-based commands |
| `appstore` | Mac App Store (alias for `mas`) |

## Global Directives

### `cask_args` — Default Cask Arguments

Sets default arguments for all `cask` entries in the Brewfile:

```ruby
cask_args appdir: "~/Applications", require_sha: true
```

### `ENV` — Environment Variables

```ruby
ENV["SOME_ENV_VAR"] = "some_value"
```

Set environment variables usable inside `brew bundle exec`, `brew bundle sh`, or `system` commands.

## Conditional Logic

Brewfiles are Ruby, so conditionals work:

```ruby
brew "gnupg" if OS.mac?
brew "glibc" if OS.linux?
cask "java" unless system "/usr/libexec/java_home", "--failfast"
```

> `brew bundle list` will not output entries hidden by Ruby conditionals on other platforms.

## `brew bundle` Subcommands

### `brew bundle` / `brew bundle install`

Reads the Brewfile and installs/upgrades all dependencies.

```bash
brew bundle                    # same as brew bundle install
brew bundle install
brew bundle --no-upgrade       # skip upgrades
brew bundle --global           # use ~/.Brewfile
brew bundle --file=/path/to/Brewfile
```

### `brew bundle upgrade`

Explicitly triggers upgrades (useful when `--no-upgrade` is the default).

### `brew bundle dump`

Snapshots all installed supported packages into a Brewfile.

```bash
brew bundle dump --global --force
brew bundle dump --file=/path/to/Brewfile --force
```

### `brew bundle cleanup`

Uninstalls supported dependencies that are installed but NOT listed in the Brewfile.

```bash
brew bundle cleanup --global --force
```

### `brew bundle check`

Checks if all Brewfile dependencies are satisfied. Returns exit code 0 if satisfied, non-zero if not.

```bash
brew bundle check
brew bundle check --verbose
brew bundle check || brew bundle install
```

### `brew bundle exec`

Runs a command in an environment customized by the Brewfile — ensures all paths to Brewfile dependencies are set up.

```bash
brew bundle exec which node
brew bundle exec --check       # run check before exec
brew bundle exec --install     # run install before exec
brew bundle exec --services    # start all services during exec
```

Sets `HOMEBREW_INSIDE_BUNDLE=1`.

### `brew bundle sh`

Like `brew bundle exec` but drops into an interactive shell with the Brewfile environment.

```bash
brew bundle sh
brew bundle sh --check
brew bundle sh --install
brew bundle sh --services
```

### `brew bundle list`

Lists all entries in the Brewfile.

```bash
brew bundle list
brew bundle list --cask
brew bundle list --taps
```

### `brew bundle add`

Adds an entry to the Brewfile.

```bash
brew bundle add wget
```

### `brew bundle remove`

Removes an entry from the Brewfile.

```bash
brew bundle remove wget
```

### `brew bundle edit`

Opens the Brewfile in the default text editor.

### `brew bundle env`

Dumps environment variables for Brewfile deps in a form suitable for shell evaluation.

```bash
brew bundle env
eval "$(brew bundle env)"
```

## Common Flags Across Subcommands

| Flag | Purpose |
|------|---------|
| `--file <path>` | Use a specific Brewfile |
| `-g, --global` | Use `~/.Brewfile` |
| `--install` | Run install first (exec, sh, env) |
| `--upgrade` | Run upgrade |
| `--no-upgrade` | Skip upgrades |
| `-v, --verbose` / `-f, --force` | Standard flags |
| `--formula` / `--cask` / `--tap` / `--mas` / `--vscode` / `--go` / `--cargo` / `--uv` / `--flatpak` / `--winget` / `--krew` / `--npm` | Filter by type |
| `--jobs <n>` | Parallel installs (default 1) |
| `--zap` | Use `zap` instead of `uninstall` for casks (cleanup) |
| `--no-describe` | Don't add description comments (dump) |
| `--no-restart` | Don't restart services (dump) |

## Environment Variables

| Variable | Purpose |
|----------|---------|
| `HOMEBREW_BUNDLE_FILE` | Default Brewfile path |
| `HOMEBREW_BUNDLE_FILE_GLOBAL` | Global Brewfile path (for `--global`) |
| `HOMEBREW_BUNDLE_JOBS` | Parallel installs (`auto` = max 4) |
| `HOMEBREW_BUNDLE_NO_JOBS` | Disable parallel jobs |
| `HOMEBREW_BUNDLE_NO_UPGRADE` | Skip all upgrades, install only missing |
| `HOMEBREW_BUNDLE_NO_DESCRIBE` | Don't enable description comments |
| `HOMEBREW_BUNDLE_SERVICES` | Start services during `brew bundle sh` |
| `HOMEBREW_BUNDLE_CHECK` | Check deps before `brew bundle sh` |
| `HOMEBREW_BUNDLE_FORCE_INSTALL_CLEANUP` | Run cleanup after install |
| `HOMEBREW_BUNDLE_SECRETS` | Disable default secret scrubbing |
| `HOMEBREW_BUNDLE_USER_CACHE` | Custom bundle cache directory |
| `HOMEBREW_BUNDLE_DUMP_NO_*` | Skip specific types during dump |
| `HOMEBREW_BUNDLE_CLEANUP_NO_*` | Skip specific types during cleanup |
| `HOMEBREW_INSIDE_BUNDLE` | Set automatically inside `exec`/`sh`/`env` |

## References

- [Brew Bundle and Brewfile](https://docs.brew.sh/Brew-Bundle-and-Brewfile)
- For services: [services.md](services.md)
- For casks: [cask.md](cask.md)
- For environment variables: [environment.md](environment.md)
