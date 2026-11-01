# DNF5 Command-Line Aliases

DNF5 supports custom aliases for commands and options using TOML configuration files. Aliases abbreviate longer command and option sequences.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/misc/aliases.7.html>.

## Configuration Locations

Aliases are loaded from these drop-in directories (in order):

1. `/usr/share/dnf5/aliases.d/` — Distribution aliases
2. `/etc/dnf/dnf5-aliases.d/` — System-wide user aliases
3. `$XDG_CONFIG_HOME/dnf5/aliases.d/` — Per-user aliases

## File Format

TOML format. Must begin with the `version` attribute:

```toml
version = '1.1'
```

Each alias is defined in a separate TOML section. The section name uniquely identifies the alias and defines its scope using dot-separated paths:

```toml
['group.list.in']    # alias only within 'group list' subcommand
```

## Alias Types

### 1. `command` — Alias for a command

Defines an alias for a command or subcommand sequence.

**Required keys**: `type`, `attached_command`

**Optional keys**: `descr.<locale>` (v1.1), `descr` (deprecated v1.1), `group_id`, `complete` (default: `false`), `required_values` (v1.1), `attached_named_args`

```toml
['grouplist']
type = 'command'
attached_command = 'group.list'

['list-fedora-all']
type = 'command'
attached_command = 'list'
attached_named_args = [
    { id_path = 'repo', value = 'fedora' },
    { id_path = 'list.showduplicates' }
]
```

### 2. `cloned_named_arg` — Alias for an existing option

Defines another name for a given option.

**Required keys**: `type`, either `long_name` or `short_name`, `source`

```toml
['nobest']
type = 'cloned_named_arg'
long_name = 'nobest'
source = 'no-best'

['repoquery.list']
type = 'cloned_named_arg'
long_name = 'list'
short_name = 'l'
source = 'repoquery.files'
```

### 3. `named_arg` — Alias that replaces multiple options

Defines an alias option that can substitute multiple options and define values for each.

**Required keys**: `type`, either `long_name` or `short_name`

**Optional keys**: `descr.<locale>` (v1.1), `descr` (deprecated), `has_value` (default: `false`), `value_help.<locale>` (v1.1), `value_help` (deprecated), `const_value`, `group_id`, `complete`, `attached_named_args`

The `value` in `attached_named_args` can contain `${}` placeholders that are substituted with the alias option's value.

```toml
['list.all-available']
type = 'named_arg'
long_name = 'all-available'
attached_named_args = [
    { id_path = 'list.showduplicates' },
    { id_path = 'list.available' }
]

['download.dest']
type = 'named_arg'
long_name = 'dest'
has_value = true
value_help.C = 'DESTDIR'
attached_named_args = [
    { id_path = 'download.destdir', value = '${}' }
]

['settsflags']
type = 'named_arg'
long_name = 'settsflags'
has_value = true
value_help.C = 'TS_FLAGS'
attached_named_args = [
    { id_path = 'setopt', value = 'tsflags=${}' }
]
```

### 4. `group` — Groups multiple aliases together

Defines a group header for organizing aliases in help output.

**Required keys**: `type`, `header.<locale>` (v1.1) or `header` (deprecated)

```toml
['repo.query-aliases']
type = 'group'
header.C = 'Query Aliases:'
```

Aliases reference the group via `group_id = 'repo.query-aliases'`.

## Value Placeholders

In `attached_named_args`, the `value` field supports placeholders:

| Placeholder | Meaning |
|-------------|---------|
| `${}` | The value of the alias option (when `has_value = true`) |
| `${1}`, `${2}`, ... | Consumed positional values (index 1 = first value) |

## `required_values` (v1.1)

Defines positional arguments consumed by the alias:

```toml
['whatrequires']
type = 'command'
attached_command = 'repoquery'
required_values = [
    { value_help.C = 'PACKAGE' }
]
attached_named_args = [
    { id_path = 'repoquery.installed' },
    { id_path = 'repoquery.whatrequires', value = '${1}' }
]
```

## `attached_named_args` Structure

An array of tables, each must contain `id_path` (the path to the option being attached). Optional `value` field:

```toml
attached_named_args = [
    { id_path = 'repo', value = 'fedora' },
    { id_path = 'list.showduplicates' }
]
```

- `id_path` — Dot-separated path to the option being set
- `value` — Value for the option. Can contain `${}` and `${index}` placeholders.

## `complete` Flag

When `complete = true`, bash autocompletion is enabled for the alias. Default: `false`.

## Localization

v1.1 added locale support:
- `descr.<locale>` — Description in specific locale (e.g., `descr.en_US`)
- `descr.C` — Fallback description (used when locale-specific is not available)
- `descr` — Deprecated in v1.1, no locale support, kept for backward compatibility

Same pattern applies to `value_help.<locale>` and `header.<locale>`.
