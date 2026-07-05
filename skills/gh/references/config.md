# Configuration and Environment

How to configure `gh` settings, use environment variables, set up shell completion, and understand exit codes.

## Configuration Commands

### `gh config get`

Print the value of a configuration key. Supports per-host settings with `--host`.

```sh
$ gh config get git_protocol
$ gh config get git_protocol --host github.com
```

### `gh config set`

Update a configuration key with a value. Supports per-host settings.

```sh
$ gh config set editor vim
$ gh config set editor "code --wait"
$ gh config set git_protocol ssh --host github.com
$ gh config set prompt disabled
```

### `gh config list`

Print all configuration keys and values. Aliased as `gh config ls`. Supports `--host` for per-host configuration.

### `gh config clear-cache`

Clear the gh configuration cache. Useful when configuration changes aren't being picked up.

## Configurable Settings

| Key | Values | Default | Notes |
|-----|--------|---------|-------|
| `git_protocol` | `https` \| `ssh` | `https` | Protocol for git clone/push. Can be set per-host. |
| `editor` | string | (system default) | Text editor for authoring text |
| `prompt` | `enabled` \| `disabled` | `enabled` | Toggle interactive prompting |
| `prefer_editor_prompt` | `enabled` \| `disabled` | `disabled` | Preference for editor-based interactive prompting |
| `pager` | string | (system default) | Terminal pager program |
| `http_unix_socket` | string | — | Path to a Unix socket for HTTP connection |
| `browser` | string | (system default) | Web browser for opening URLs |
| `color_labels` | `enabled` \| `disabled` | `disabled` | Display labels using RGB hex colors (truecolor terminals) |
| `accessible_colors` | `enabled` \| `disabled` | `disabled` | Use 4-bit accessible colors |
| `accessible_prompter` | `enabled` \| `disabled` | `disabled` | Use accessible prompter (screen reader compatible) |
| `spinner` | `enabled` \| `disabled` | `enabled` | Animated spinner progress indicator |
| `telemetry` | `enabled` \| `disabled` \| `log` | `disabled` | Telemetry mode (`log` prints to stderr) |

## Environment Variables

### Authentication

| Variable | Description |
|----------|-------------|
| `GH_TOKEN` | Auth token for `github.com` / `ghe.com` subdomains (highest precedence) |
| `GITHUB_TOKEN` | Auth token for `github.com` (second precedence) |
| `GH_ENTERPRISE_TOKEN` | Auth token for GHES hosts (highest precedence) |
| `GITHUB_ENTERPRISE_TOKEN` | Auth token for GHES (second precedence) |
| `GH_HOST` | Default GitHub hostname when not inferable from context |

See [auth.md](auth.md) for authentication details.

### Repository and Context

| Variable | Description |
|----------|-------------|
| `GH_REPO` | Repository in `[HOST/]OWNER/REPO` format for commands operating on a local repository |

### Editor and Browser

| Variable | Description |
|----------|-------------|
| `GH_EDITOR` | Editor tool (highest precedence) |
| `GIT_EDITOR` | Editor (second) |
| `VISUAL` | Editor (third) |
| `EDITOR` | Editor (fourth) |
| `GH_BROWSER` | Web browser (highest precedence) |
| `BROWSER` | Web browser (second) |

### Output and Display

| Variable | Description |
|----------|-------------|
| `GH_PAGER` | Terminal pager (highest precedence) |
| `PAGER` | Terminal pager (second) |
| `GLAMOUR_STYLE` | Markdown rendering style (see [glamour styles](https://github.com/charmbracelet/glamour#styles)) |
| `NO_COLOR` | Disable ANSI color output |
| `CLICOLOR` | Set to `0` to disable ANSI colors |
| `CLICOLOR_FORCE` | Non-`0` to force ANSI colors even when piped |
| `GH_COLOR_LABELS` | Display labels with RGB hex colors (truecolor) |
| `GH_FORCE_TTY` | Force terminal-style output; number = column count, percentage = applied to viewport |
| `GH_MDWIDTH` | Default max width for markdown render wrapping (max 120) |

### Debug and Development

| Variable | Description |
|----------|-------------|
| `GH_DEBUG` | Truthy = verbose stderr; `api` = also log HTTP traffic |
| `DEBUG` | (deprecated) `1`/`true`/`yes` = verbose stderr |
| `GH_PATH` | Path to gh executable (useful for cygwin) |

### Update and Extension Notifications

| Variable | Description |
|----------|-------------|
| `GH_NO_UPDATE_NOTIFIER` | Disable gh update notifications (checked every 24h) |
| `GH_NO_EXTENSION_UPDATE_NOTIFIER` | Disable extension update notifications |

### Config Directory

`GH_CONFIG_DIR` — where gh stores configuration files. Default (in precedence order):
1. `$XDG_CONFIG_HOME/gh` (if `$XDG_CONFIG_HOME` is set)
2. `$AppData/GitHub CLI` (Windows, if `$AppData` is set)
3. `$HOME/.config/gh`

### Prompting

| Variable | Description |
|----------|-------------|
| `GH_PROMPT_DISABLED` | Disable interactive prompting |
| `GH_ACCESSIBLE_PROMPTER` | Enable screen reader / braille compatible prompts (preview) |

### Telemetry

| Variable | Description |
|----------|-------------|
| `GH_TELEMETRY` | `log` = print to stderr; `false`/`0` = disable; takes precedence over `DO_NOT_TRACK` |
| `DO_NOT_TRACK` | `true`/`1` = disable telemetry (ignored when `GH_TELEMETRY` is set) |

### Spinner

| Variable | Description |
|----------|-------------|
| `GH_SPINNER_DISABLED` | Replace spinner animation with textual progress indicator |

## Shell Completion

`gh completion` generates shell completion scripts.

```sh
# bash
eval "$(gh completion -s bash)"

# zsh — generate _gh and place in $fpath
gh completion -s zsh > /usr/local/share/zsh/site-functions/_gh
# Ensure ~/.zshrc has:
#   autoload -U compinit
#   compinit -i

# fish
gh completion -s fish > ~/.config/fish/completions/gh.fish

# PowerShell
Invoke-Expression -Command $(gh completion -s powershell | Out-String)
```

| Flag | Description |
|------|-------------|
| `-s, --shell string` | Shell type: `bash`, `zsh`, `fish`, `powershell` |

Zsh version 5.7+ is recommended. When installed via package manager (e.g. Homebrew), shell configuration may be automatic.

## Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Command completed successfully |
| `1` | Command failed for any reason |
| `2` | Command was running but got cancelled |
| `4` | Command requires authentication |

> Some commands may have additional exit codes — check per-command documentation.

## Aliases

`gh alias` manages command shortcuts.

### `gh alias set`

Define a word that expands to a full gh command. Supports positional placeholders (`$1`, `$2`, ...) and shell expressions.

```sh
$ gh alias set pv 'pr view'
$ gh pv -w 123  #=> gh pr view -w 123

$ gh alias set bugs 'issue list --label=bugs'
$ gh alias set 'issue mine' 'issue list --mention @me'
$ gh alias set epicsBy 'issue list --author="$1" --label="epic"'
$ gh epicsBy vilmibm  #=> gh issue list --author="vilmibm" --label="epic"

# Shell expression (starts with ! or --shell)
$ gh alias set --shell igrep 'gh issue list --label="$1" | grep "$2"'
$ gh igrep epic foo  #=> gh issue list --label="epic" | grep "foo"
```

| Flag | Description |
|------|-------------|
| `--clobber` | Overwrite existing aliases of the same name |
| `-s, --shell` | Declare alias as a shell expression (evaluated via `sh`) |

Use `-` as the expansion argument to read from stdin (avoids quoting issues).

### `gh alias list`

Print all configured aliases. Aliased as `gh alias ls`.

### `gh alias delete`

Delete one or all aliases.

```sh
$ gh alias delete pv          # delete one
$ gh alias delete --all       # delete all
```

### `gh alias import`

Import aliases from a YAML file. Keys are alias names, values are expansions.

```yaml
bugs: issue list --label=bug
igrep: '!gh issue list --label="$1" | grep "$2"'
features: |-
    issue list
    --label=enhancement
```

```sh
$ gh alias import aliases.yml
$ gh alias import -            # from stdin
```

| Flag | Description |
|------|-------------|
| `--clobber` | Overwrite existing aliases of the same name |

## Accessibility

gh provides accessible experiences configurable via `gh config set` or environment variables:

| Feature | Config Key | Env Var |
|---------|-----------|---------|
| 4-bit accessible colors | `accessible_colors` | `GH_ACCESSIBLE_COLORS` |
| Truecolor label colors | `color_labels` | `GH_COLOR_LABELS` |
| Screen reader compatible prompts | `accessible_prompter` | `GH_ACCESSIBLE_PROMPTER` |
| Text-based spinner (no animation) | `spinner disabled` | `GH_SPINNER_DISABLED` |

## Related

- For authentication details, see [auth.md](auth.md).
- For JSON output formatting (`--json`, `--jq`, `--template`), see [output-formatting.md](output-formatting.md).
