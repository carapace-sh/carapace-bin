# Revision Commands

Commands for managing Lore revisions — history, info, commit, amend, sync, bisect, diff, find, restore, cherry-pick, revert, and metadata.

> **Source of truth**: <https://epicgames.github.io/lore/reference/lore-cli-commands/>

## `lore revision history`

**Usage:** `lore revision history [OPTIONS] [LENGTH]`

List revisions of a repository.

| Option | Description |
|--------|-------------|
| `--revision <revision>` | Start listing from the specified revision (default: current branch latest) |
| `--branch <branch>` | Show branch revisions |
| `--only-branch` | Stop when reaching a revision on a different branch |
| `--oneline` | Output each revision on one line only |

## `lore revision info`

**Usage:** `lore revision info [OPTIONS] [revision]`

Get info about a revision.

| Option | Description |
|--------|-------------|
| `--delta` | Show delta information |
| `--metadata` | Show file metadata information |

## `lore revision commit`

**Usage:** `lore revision commit [OPTIONS] <MESSAGE>`

Commit the staged state.

| Option | Description |
|--------|-------------|
| `--stats` | Print stats |
| `--link <LINK>` | Commit only changes in this linked repository |
| `--link-message <PATH>` | Per-link commit message (can be specified multiple times) |
| `--layer <LAYER>` | Commit only changes in this layer |
| `--layer-message <PATH>` | Per-layer commit message (can be specified multiple times) |

## `lore revision amend`

**Usage:** `lore revision amend [OPTIONS] <MESSAGE>`

Amend the latest commit's message.

| Option | Description |
|--------|-------------|
| `--stats` | Print stats |

## `lore revision sync`

**Usage:** `lore revision sync [OPTIONS] [revision]`

**Alias:** `synchronize`

Synchronize to a given state of a repository. Revision can be a partial hash signature.

| Option | Description |
|--------|-------------|
| `--forward-changes` | Fast forward any local changes if syncing to a local revision |
| `--reset` | Reset any local modified files to match the incoming revision |
| `--root-file <path>` | Root files for dependency-based selective sync |
| `--dependency-tag <tag>` | Tags to filter dependencies by |
| `--dependency-recursive` | Follow transitive dependencies recursively |
| `--dependency-depth-limit <depth>` | Max dependency traversal depth (0 = unlimited) |

## `lore revision bisect`

**Usage:** `lore revision bisect --start <start_revision> --end <end_revision>`

Binary search for a change introduced between start (exclusive) and end (inclusive).

| Option | Description |
|--------|-------------|
| `--start <start_revision>` | The latest revision known to not have the change |
| `--end <end_revision>` | The earliest revision known to have the change |

## `lore revision diff`

**Usage:** `lore revision diff [OPTIONS] <revision_source>`

Diff two revisions.

| Option | Description |
|--------|-------------|
| `--target <revision_target>` | Target revision to compare (default: current revision) |
| `--path <PATH>` | Optional path in repository |
| `--targets <file>` | Path to a targets file |

## `lore revision find`

**Usage:** `lore revision find <COMMAND>`

Find revision by metadata or number.

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `metadata` | `lore revision find metadata <key> [value]` | Find revision by metadata key/value |
| `number` | `lore revision find number <NUMBER>` | Find revision by revision number |

## `lore revision restore`

**Usage:** `lore revision restore [MESSAGE]`

Restore current revision as latest revision.

## `lore revision cherry-pick`

**Usage:** `lore revision cherry-pick [OPTIONS] <revision>`

Cherry-pick a revision onto the currently synced revision.

| Option | Description |
|--------|-------------|
| `--message <MESSAGE>` | Change the message for committing when no conflicts arise |
| `--no-commit` | Disable auto commits even if no conflicts arise |

### Subcommands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `unresolve` | `lore revision cherry-pick unresolve <paths\|--targets <file>>` | Marks the cherry-pick unresolved |
| `restart` | `lore revision cherry-pick restart <paths\|--targets <file>>` | Restart the cherry-pick for specific files |
| `resolve` | `lore revision cherry-pick resolve <paths\|--targets <file>>` | Resolve conflicts |
| `abort` | `lore revision cherry-pick abort` | Abort a cherry-pick |

### `lore revision cherry-pick resolve` Subcommands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `mine` | `resolve mine <paths\|--targets <file>>` | Resolve using my changes |
| `theirs` | `resolve theirs <paths\|--targets <file>>` | Resolve using the incoming changes |

## `lore revision revert`

**Usage:** `lore revision revert [OPTIONS] <revision>`

Revert a revision from the currently synced revision.

| Option | Description |
|--------|-------------|
| `--message <MESSAGE>` | Change the message for committing when no conflicts arise |
| `--no-commit` | Disable auto commits even if no conflicts arise |

### Subcommands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `unresolve` | `lore revision revert unresolve <paths\|--targets <file>>` | Marks the revert unresolved |
| `restart` | `lore revision revert restart <paths\|--targets <file>>` | Restart the revert for specific files |
| `resolve` | `lore revision revert resolve <paths\|--targets <file>>` | Resolve conflicts |
| `abort` | `lore revision revert abort` | Abort a revert |

### `lore revision revert resolve` Subcommands

Same as cherry-pick: `mine` and `theirs` subcommands.

## `lore revision metadata`

Manage metadata on a revision.

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `clear` | `lore revision metadata clear` | Clear metadata for a staged revision |
| `get` | `lore revision metadata get [OPTIONS] [key]` | Get metadata from a revision |
| `set` | `lore revision metadata set [OPTIONS] [pairs]...` | Set metadata for a staged revision |

Options for `get`: `--revision <revision>`.
Options for `set`: `--binary` (values are paths to files).