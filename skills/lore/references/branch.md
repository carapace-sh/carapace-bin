# Branch Commands

Commands for managing Lore branches — listing, creating, switching, merging, diffing, archiving, protecting, and metadata.

> **Source of truth**: <https://epicgames.github.io/lore/reference/lore-cli-commands/>

## `lore branch list`

**Usage:** `lore branch list [OPTIONS]`

List available branches.

| Option | Description |
|--------|-------------|
| `--archived` | Include archived local branches |

## `lore branch info`

**Usage:** `lore branch info [branch]`

Get info about the given branch.

## `lore branch create`

**Usage:** `lore branch create [OPTIONS] <branch>`

Create a new branch.

| Option | Description |
|--------|-------------|
| `--id <id>` | Optional explicit branch ID (hex-encoded 16-byte identifier) |

## `lore branch switch`

**Usage:** `lore branch switch [OPTIONS] <branch> [revision]`

Switch to a different branch. Optionally specify a revision to switch to.

| Option | Description |
|--------|-------------|
| `--dry-run` | Do a dry run sync and only report what changes would be done |
| `--local` | Keep last local latest revision, do not sync latest from remote |
| `--reset` | Reset any local modified files to match the incoming revision |
| `--bare` | Only update anchor tracking without modifying or verifying files |

## `lore branch push`

**Usage:** `lore branch push [OPTIONS] [branch]`

Push commits to remote. Pushes current branch if not specified.

| Option | Description |
|--------|-------------|
| `--fast-forward-merge` | Allow the server to fast-forward merge if the target branch head has moved |

## `lore branch merge`

**Usage:** `lore branch merge [OPTIONS] <branch|--id <branch-id>>`

Merge two branches. Source branch is merged into the current branch.

| Option | Description |
|--------|-------------|
| `--id <branch-id>` | ID of the source branch |
| `--message <MESSAGE>` | Change the message for committing when no conflicts arise |

### Subcommands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `unresolve` | `lore branch merge unresolve <paths\|--targets <file>>` | Marks the merge unresolved |
| `into` | `lore branch merge into [OPTIONS] <branch\|--id> <MESSAGE>` | Merge current branch into target branch |
| `start` | `lore branch merge start [OPTIONS] <branch\|--id>` | Start a merge process |
| `restart` | `lore branch merge restart <paths\|--targets <file>>` | Restart the merge for specific files |
| `resolve` | `lore branch merge resolve [OPTIONS] [paths]...` | Resolve merge conflicts |
| `abort` | `lore branch merge abort [OPTIONS]` | Abort a merge process |

### `lore branch merge start` Options

| Option | Description |
|--------|-------------|
| `--id <branch-id>` | ID of the source branch |
| `--message <MESSAGE>` | Change the message for committing when no conflicts arise |
| `--no-commit` | Disable auto commits even if no conflicts arise |
| `--dry-run` | Dry run, only report what would change |
| `--link <LINK>` | Merge only a specific linked repository at the given mount path |
| `--ignore-links` | Merge only the main repository, skipping all linked repositories |

### `lore branch merge into` Options

| Option | Description |
|--------|-------------|
| `--id <branch-id>` | ID of the target branch |
| `--link <LINK>` | Merge only a specific linked repository at the given mount path |
| `--ignore-links` | Merge only the main repository, skipping all linked repositories |

### `lore branch merge resolve` Subcommands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `mine` | `lore branch merge resolve mine <paths\|--targets <file>>` | Resolve using my changes |
| `theirs` | `lore branch merge resolve theirs <paths\|--targets <file>>` | Resolve using their changes |

### `lore branch merge abort` Options

| Option | Description |
|--------|-------------|
| `--link <LINK>` | Abort only a specific linked repository merge |
| `--ignore-links` | Abort only the main repository merge, keeping link pin updates |

## `lore branch diff`

**Usage:** `lore branch diff [OPTIONS] <target>`

Diff two branches using the common ancestor base revision.

| Option | Description |
|--------|-------------|
| `--source <source>` | Name of the source branch |
| `--auto-resolve` | Attempt to auto resolve conflicts if true |

## `lore branch archive`

**Usage:** `lore branch archive <branch>`

Archive an existing branch.

## `lore branch reset`

**Usage:** `lore branch reset [OPTIONS] <revision>`

Reset local latest pointer for a branch.

| Option | Description |
|--------|-------------|
| `--branch <branch>` | Branch to reset, or the current branch if not set |

## `lore branch protect`

**Usage:** `lore branch protect <branch>`

Protect a branch from direct pushes.

## `lore branch unprotect`

**Usage:** `lore branch unprotect <branch>`

Remove push protection from a branch.

## `lore branch latest`

**Usage:** `lore branch latest list [OPTIONS] [LIMIT]`

Branch latest related commands.

| Option | Description |
|--------|-------------|
| `--branch <branch>` | Branch to query |

## `lore branch metadata`

Manage metadata on a branch.

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `get` | `lore branch metadata get [OPTIONS] [key]` | Get metadata (omit key to list all) |
| `set` | `lore branch metadata set [OPTIONS] [pairs]...` | Set metadata key/value pairs |
| `clear` | `lore branch metadata clear [OPTIONS] [keys]...` | Clear metadata (omit keys to clear all) |

Options for all: `--branch <branch>` (branch name, uses current if not specified).
Options for `set`: `--binary`, `--numeric`.