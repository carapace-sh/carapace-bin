# File Commands

Commands for managing files in a Lore repository — info, metadata, dependencies, stage, dirty, unstage, reset, obliterate, history, diff, write, and hash.

> **Source of truth**: <https://epicgames.github.io/lore/reference/lore-cli-commands/>

## `lore file info`

**Usage:** `lore file info [OPTIONS] <paths|--targets <file>>`

Get info about the given file or directory.

| Option | Description |
|--------|-------------|
| `--targets <file>` | Path to a targets file |
| `--revision <revision>` | Revision to get info from |
| `--local` | Calculate the local file system size and hash based on the current local filter |
| `--filtered` | Calculate the repository size based on the current local filter |

## `lore file metadata`

Manage metadata on a file.

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `clear` | `lore file metadata clear <PATH>` | Clear metadata for a staged file |
| `get` | `lore file metadata get [OPTIONS] <PATH> [key]` | Get metadata from a file |
| `set` | `lore file metadata set [OPTIONS] <PATH> [pairs]...` | Set metadata on a staged file |

Options for `get`: `--revision <revision>`.
Options for `set`: `--binary` (values are paths to files).

## `lore file dependency`

Manage file dependency edges.

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `add` | `lore file dependency add [OPTIONS] <SOURCE> [dependencies]...` | Add dependency edges from source to dependencies |
| `remove` | `lore file dependency remove [OPTIONS] <SOURCE> [dependencies]...` | Remove dependency edges |
| `list` | `lore file dependency list [OPTIONS] [paths]...` | List dependencies or dependents |

### `lore file dependency add` Options

| Option | Description |
|--------|-------------|
| `--tag <tag>` | Tags to apply to all added dependency edges |
| `--force` | Skip cycle detection |

### `lore file dependency remove` Options

| Option | Description |
|--------|-------------|
| `--tag <tag>` | Remove only specific tags instead of entire edges |

### `lore file dependency list` Options

| Option | Description |
|--------|-------------|
| `--reverse` | List dependents instead of dependencies |
| `--recursive` | Recursively resolve transitive dependencies |
| `--tag <tag>` | Filter by tag |
| `--depth <limit>` | Maximum recursion depth (0 = unlimited) |
| `--revision <revision>` | Revision to query (defaults to staged/current) |

## `lore file stage`

**Usage:** `lore file stage [OPTIONS] <paths|--targets <file>>`

Stage changes for commit.

| Option | Description |
|--------|-------------|
| `--case <case>` | Case change handling (`error`, `keep`, `rename`) |
| `--scan` | Walk the filesystem to detect modified, added, and deleted files |
| `--targets <file>` | Path to a targets file |

### Subcommands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `move` | `lore file stage move <from> <to>` | Move or rename a file or directory |
| `merge` | `lore file stage merge <paths\|--targets <file>>` | Stage as a merge |

## `lore file dirty`

**Usage:** `lore file dirty [OPTIONS] [paths]...`

Mark files as dirty (modified in working tree without staging).

| Option | Description |
|--------|-------------|
| `--targets <file>` | Path to a targets file |

### Subcommands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `move` | `lore file dirty move <from> <to>` | Mark a file as moved (dirty) |
| `copy` | `lore file dirty copy <from> <to>` | Mark a file as copied (dirty) |

## `lore file unstage`

**Usage:** `lore file unstage <paths|--targets <file>>`

Unstage changes to a file or directory.

| Option | Description |
|--------|-------------|
| `--targets <file>` | Path to a targets file |

## `lore file reset`

**Usage:** `lore file reset [OPTIONS] <paths|--targets <file>>`

Reset changes to a path or file to the current revision.

| Option | Description |
|--------|-------------|
| `--purge` | Delete untracked files |
| `--targets <file>` | Path to a targets file |
| `--revision <revision>` | Revision to reset files to |
| `--last-merged-from <branch>` | Reset to the last point of merge from this branch |

## `lore file obliterate`

**Usage:** `lore file obliterate <--address <ADDRESS>|--path <PATH>>`

Obliterate a file or fragment.

| Option | Description |
|--------|-------------|
| `--address <ADDRESS>` | Address of a blob |
| `--path <PATH>` | Path to a file |

## `lore file history`

**Usage:** `lore file history [OPTIONS] <PATH> [LENGTH]`

List revisions of a file.

| Option | Description |
|--------|-------------|
| `--revision <revision>` | Revision to start from |
| `--branch <branch>` | Show branch revisions |
| `--depth <depth>` | Number of revisions to search initially |
| `--oneline` | Output each revision on one line only |

## `lore file diff`

**Usage:** `lore file diff [OPTIONS] [paths]...`

Show differences between two revisions of a file.

| Option | Description |
|--------|-------------|
| `--source <revision_source>` | Source revision to diff from (default: current revision) |
| `--target <revision_target>` | Target revision to diff to (default: current file system state) |
| `--diff3` | Produce three-way merge output with conflict markers |
| `-U`, `--context <n>` | Number of unchanged context lines (default: 3) |
| `--ignore-space-at-eol` | Ignore trailing whitespace differences |
| `--ignore-space-change` | Collapse internal whitespace runs before comparing |
| `--targets <file>` | Path to a targets file |

## `lore file write`

**Usage:** `lore file write [OPTIONS] --output <OUTPUT>`

Write data to a specific location from a blob address or file path.

| Option | Description |
|--------|-------------|
| `--address <ADDRESS>` | Address of a blob |
| `--path <PATH>` | Path to a file |
| `--revision <REVISION>` | Revision specifier |
| `--output <OUTPUT>` | Path to a destination (required) |

## `lore file hash`

**Usage:** `lore file hash [OPTIONS] [paths]...`

Hash a local file.

| Option | Description |
|--------|-------------|
| `--targets <file>` | Path to a targets file |