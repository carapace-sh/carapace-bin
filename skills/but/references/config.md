# Configuration, Aliases, and AI Settings

Commands for viewing and managing GitButler configuration, aliases, user settings, forge auth, AI providers, and UI preferences.

> **Source of truth**: <https://docs.gitbutler.com/cli-guides/cli-tutorial/configuration>

## `but config` — View and Manage Configuration

Parent command for configuration management.

```bash
but config                    # View all configuration
```

### `but config user`

User name, email, and editor settings.

```bash
but config user set <key> <value>     # Set a user config value
but config user unset <key>           # Unset a user config value
```

Config keys:

| Key | Description |
|-----|-------------|
| `name` | User name for commits |
| `email` | User email for commits |
| `editor` | Editor command |

### `but config target`

View or change the target branch (merge base).

```bash
but config target                    # Show current target
but config target <branch>           # Change target branch
```

Positional argument is a remote branch name (e.g. `origin/main`).

### `but config forge`

Forge (GitHub) settings and authentication.

```bash
but config forge                     # View forge config
but config forge auth                # Authenticate to forge
but config forge forget              # Forget stored auth
but config forge list-users          # List authenticated users
```

`but config forge forget` accepts a positional argument: the forge user to forget.

### `but config metrics`

Enable or disable metrics collection.

```bash
but config metrics enable
but config metrics disable
```

Positional argument is `enable` or `disable` (styled as keywords).

### `but config ui`

UI configuration settings.

```bash
but config ui set <key> <value>       # Set a UI config value
but config ui unset <key>             # Unset a UI config value
```

UI config keys:

| Key | Values | Description |
|-----|--------|-------------|
| `tui` | `true`, `false` | Enable TUI mode for diffs by default |

### `but config ai`

AI provider configuration. Parent command for AI provider subcommands.

#### `but config ai anthropic`

Configure Anthropic (Claude) as the AI provider.

| Flag | Description |
|------|-------------|
| `--api-key-env` | Environment variable name containing the API key |
| `--key-option` | Key option (provider-specific) |
| `--model` | Model name to use |

#### `but config ai openai`

Configure OpenAI as the AI provider.

| Flag | Description |
|------|-------------|
| `--api-key-env` | Environment variable name containing the API key |

#### `but config ai openrouter`

Configure OpenRouter as the AI provider.

| Flag | Description |
|------|-------------|
| `--api-key-env` | Environment variable name containing the API key |

#### `but config ai ollama`

Configure Ollama (local) as the AI provider.

#### `but config ai lmstudio`

Configure LM Studio (local) as the AI provider.

#### `but config ai show`

Show current AI configuration.

## Aliases

### `but alias` — Manage Command Aliases

Parent command for alias management.

```bash
but alias                     # List aliases (default subcommand)
```

### `but alias list`

List all configured aliases. Shows both user-defined and default aliases.

```bash
but alias list
```

### `but alias add`

Add a new alias.

```bash
but alias add <name> <expansion>
but alias add --global <name> <expansion>    # Store in ~/.gitconfig
```

| Flag | Short | Description |
|------|-------|-------------|
| `--global` | `-g` | Store globally (in `~/.gitconfig`) instead of locally |

Positional arguments:
1. Alias name
2. Expansion (the command the alias expands to)

### `but alias remove`

Remove an existing alias.

```bash
but alias remove <name>
but alias rm <name>                           # Alias for remove
```

| Flag | Short | Description |
|------|-------|-------------|
| `--global` | `-g` | Remove from global config instead of local |

Positional argument is the alias name.

### Default Aliases

GitButler ships with built-in default aliases:

| Alias | Expansion |
|-------|-----------|
| `default` | `status` |
| `st` | `status` |
| `stf` | `status --files` |

The `default` alias runs when you type just `but` with no arguments. User aliases override defaults.

### JSON Output

```bash
but --format json alias list
```

Returns a JSON object with `User` (array of `{name, value, scope}`) and `Default` (map of name→value).

## Related

- [ai-agents.md](ai-agents.md) — AI agent setup and skills
- [forge.md](forge.md) — Forge authentication details
- [misc.md](misc.md) — Root command global flags
