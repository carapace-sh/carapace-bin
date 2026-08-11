# CLI Command Reference

Every `lore` command, subcommand, argument, and flag.

> **Source of truth**: <https://epicgames.github.io/lore/reference/lore-cli-commands/>

## Global Options

| Flag | Description |
|------|-------------|
| `--repository <path>` | Use given path as repository path |
| `--log-level <level>` | Set the logging level |
| `-d`, `--debug` | Enable debug output |
| `-f`, `--force` | Force the operation if possible |
| `--dry-run` | Dry run mode, only report what would have changed |
| `-P`, `--no-pager` | Disable pagination |
| `--offline` | Force offline mode |
| `--remote` | Use remote data |
| `--local` | Use local data |
| `--identity <IDENTITY>` | Use given identity |
| `--max-connections <MAX>` | Set max parallel connections |
| `--file-count-limit <count>` | Set max parallel files opened |
| `--file-size-limit <size>` | Set max total bytes of parallel files opened |
| `--compress-limit <count>` | Set max parallel compress operations |
| `--search-limit <SEARCH_LIMIT>` | Set max revisions to search when matching/finding |
| `--search-nearest` | Search for nearest match when matching revisions |
| `--gc` | Run automatic garbage collection on local store in background |
| `--sync-data` | Force sync data to storage media during flush |
| `--non-interactive` | Disable interactive prompts |

## Top-Level Commands

| Command | Delegates to | Description |
|---------|-------------|-------------|
| `lore status` | `lore repository status` | Show current repository status |
| `lore clone` | `lore repository clone` | Clone a remote repository |
| `lore stage` | `lore file stage` | Stage changes for commit |
| `lore dirty` | `lore file dirty` | Mark files as dirty |
| `lore unstage` | `lore file unstage` | Unstage changes |
| `lore reset` | `lore file reset` | Reset changes |
| `lore diff` | `lore file diff` | Show differences |
| `lore history` | `lore revision history` | List revisions |
| `lore commit` | `lore revision commit` | Commit the staged revision |
| `lore sync` | `lore revision sync` | Synchronize to a repository state |
| `lore push` | `lore branch push` | Push commits to remote |

## `lore repository` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `status` | `status [OPTIONS] [PATH]...` | Show current repository status |
| `info` | `info [url]` | Get info about a repository |
| `list` | `list <url>` | List repositories on a remote |
| `create` | `create [OPTIONS] <url>` | Create a repository |
| `clone` | `clone [OPTIONS] <url> [path]` | Clone a remote repository |
| `delete` | `delete <url>` | Delete a repository |
| `verify` | `verify [OPTIONS] [COMMAND]` | Verify repository state consistency |
| `dump` | `dump [OPTIONS]` | Dump repository state information |
| `gc` | `gc` | Run full garbage collection |
| `store` | `store immutable query [OPTIONS] <ADDRESS>` | Access the repository data store |
| `metadata` | `metadata get/set/clear` | Repository metadata operations |
| `instance` | `instance list/prune` | Instance management |
| `config` | `config get <KEY>` | Read a configuration value |
| `update-path` | `update-path` | Update stored path for this instance |

### `lore repository status` Options

| Option | Description |
|--------|-------------|
| `--scan` | Walk the filesystem and reconcile every file against the current revision |
| `--check-dirty` | Verify already-dirty files without a full scan |
| `--reset` | Drop existing staged anchor before computing status |
| `--revision-only` | Only show revision info, skip all diffs |
| `--count` | Count directories and files |
| `--targets <file>` | Path to a targets file |

### `lore repository create` Options

| Option | Description |
|--------|-------------|
| `--description <description>` | Optional description |
| `--id <id>` | Optional ID |
| `--use-shared-store` | Use shared store |
| `--shared-store-path <PATH>` | Path to shared store |

### `lore repository clone` Options

| Option | Description |
|--------|-------------|
| `--view <view>` | Client-side view filter file |
| `--revision <revision>` | Revision to sync |
| `--branch <branch>` | Branch to sync |
| `--bare` | Clone without files, only fetch latest revision tree |
| `--virtual` | Clone virtually using split-write filesystem |
| `--direct-file-write` | Write directly to destination |
| `--layer <repository>` | Layer to add |
| `--layer-metadata <key>` | Metadata key for layer revision matching |
| `--prefetch <file>` | File containing list of files to prefetch |
| `--use-shared-store` | Use shared store |
| `--shared-store-path <PATH>` | Shared store path |
| `--no-tracking` | Clone without local repository tracking (memory-only) |
| `--root-file <path>` | Root files for dependency-based selective clone |
| `--dependency-tag <tag>` | Tags to filter dependencies by |
| `--dependency-recursive` | Follow transitive dependencies recursively |
| `--dependency-depth-limit <depth>` | Max dependency traversal depth (0 = unlimited) |

### `lore repository verify` Options

| Option | Description |
|--------|-------------|
| `--path <path>` | Path to start verification from |
| `--heal` | Attempt to heal discrepancies |

Subcommands: `state` (default), `fragment <HASH>`.

### `lore repository metadata`

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `get` | `get [key]` | Get metadata (omit key to list all) |
| `set` | `set [OPTIONS] [pairs]...` | Set metadata key/value pairs |
| `clear` | `clear [keys]...` | Clear metadata (omit keys to clear all user-defined) |

Options for `set`: `--binary`, `--numeric`.

## `lore branch` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `list` | `list [OPTIONS]` | List available branches `--archived` |
| `info` | `info [branch]` | Get info about a branch |
| `create` | `create [OPTIONS] <branch>` | Create a new branch (`--id <id>`) |
| `switch` | `switch [OPTIONS] <branch> [revision]` | Switch to a different branch |
| `push` | `push [OPTIONS] [branch]` | Push commits to remote |
| `merge` | `merge [OPTIONS] <branch\|--id>` | Merge source branch into current |
| `diff` | `diff [OPTIONS] <target>` | Diff two branches |
| `archive` | `archive <branch>` | Archive an existing branch |
| `reset` | `reset [OPTIONS] <revision>` | Reset local latest pointer |
| `protect` | `protect <branch>` | Protect branch from direct pushes |
| `unprotect` | `unprotect <branch>` | Remove push protection |
| `latest` | `latest list [OPTIONS] [LIMIT]` | Branch latest list |
| `metadata` | `metadata get/set/clear` | Branch metadata operations |

### `lore branch switch` Options

| Option | Description |
|--------|-------------|
| `--dry-run` | Dry run, only report what would change |
| `--local` | Keep last local latest revision, do not sync from remote |
| `--reset` | Reset modified files to match incoming revision |
| `--bare` | Only update anchor tracking |

### `lore branch merge` Subcommands

| Subcommand | Description |
|------------|-------------|
| `unresolve <paths\|--targets>` | Marks the merge unresolved |
| `into [OPTIONS] <branch\|--id> <MESSAGE>` | Merge current branch into target |
| `start [OPTIONS] <branch\|--id>` | Start a merge process |
| `restart <paths\|--targets>` | Restart merge for specific files |
| `resolve [OPTIONS] [paths]...` | Resolve merge conflicts |
| `abort [OPTIONS]` | Abort a merge process |

Merge `resolve` subcommands: `mine` (use my changes), `theirs` (use their changes).

## `lore revision` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `history` | `history [OPTIONS] [LENGTH]` | List revisions of a repository |
| `info` | `info [OPTIONS] [revision]` | Get info about a revision |
| `commit` | `commit [OPTIONS] <MESSAGE>` | Commit the staged state |
| `amend` | `amend [OPTIONS] <MESSAGE>` | Amend the latest commit's message |
| `sync` | `sync [OPTIONS] [revision]` | Synchronize to a given state (alias: `synchronize`) |
| `bisect` | `bisect --start <start> --end <end>` | Binary search for a change |
| `diff` | `diff [OPTIONS] <revision_source>` | Diff two revisions |
| `find` | `find metadata/number` | Find revision by metadata or number |
| `restore` | `restore [MESSAGE]` | Restore current revision as latest |
| `cherry-pick` | `cherry-pick [OPTIONS] <revision>` | Cherry-pick a revision |
| `revert` | `revert [OPTIONS] <revision>` | Revert a revision |
| `metadata` | `metadata get/set/clear` | Revision metadata operations |

### `lore revision commit` Options

| Option | Description |
|--------|-------------|
| `--stats` | Print stats |
| `--link <LINK>` | Commit only changes in linked repository |
| `--link-message <PATH>` | Per-link commit message (repeatable) |
| `--layer <LAYER>` | Commit only changes in this layer |
| `--layer-message <PATH>` | Per-layer commit message (repeatable) |

### `lore revision sync` Options

| Option | Description |
|--------|-------------|
| `--forward-changes` | Fast forward local changes if syncing to local revision |
| `--reset` | Reset modified files to match incoming revision |
| `--root-file <path>` | Root files for dependency-based selective sync |
| `--dependency-tag <tag>` | Tags to filter dependencies by |
| `--dependency-recursive` | Follow transitive dependencies |
| `--dependency-depth-limit <depth>` | Max traversal depth (0 = unlimited) |

### Cherry-pick/Revert Subcommands

`unresolve`, `restart`, `resolve` (with `mine`/`theirs` subcommands), `abort`.

## `lore file` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `info` | `info [OPTIONS] <paths\|--targets>` | Get info about a file or directory |
| `metadata` | `metadata clear/get/set` | File metadata operations |
| `dependency` | `dependency add/remove/list` | File dependency management |
| `stage` | `stage [OPTIONS] <paths\|--targets>` | Stage changes for commit |
| `dirty` | `dirty [OPTIONS] [paths]...` | Mark files as dirty |
| `unstage` | `unstage <paths\|--targets>` | Unstage changes |
| `reset` | `reset [OPTIONS] <paths\|--targets>` | Reset changes to current revision |
| `obliterate` | `obliterate <--address\|--path>` | Obliterate a file or fragment |
| `history` | `history [OPTIONS] <PATH> [LENGTH]` | List revisions of a file |
| `diff` | `diff [OPTIONS] [paths]...` | Show differences between revisions |
| `write` | `write [OPTIONS] --output <OUTPUT>` | Write data to a specific location |
| `hash` | `hash [OPTIONS] [paths]...` | Hash a local file |

### `lore file stage` Options

| Option | Description |
|--------|-------------|
| `--case <case>` | Case change handling (`error`, `keep`, `rename`) |
| `--scan` | Walk filesystem to detect changes |
| `--targets <file>` | Path to a targets file |

Subcommands: `move <from> <to>`, `merge <paths\|--targets>`.

### `lore file reset` Options

| Option | Description |
|--------|-------------|
| `--purge` | Delete untracked files |
| `--targets <file>` | Path to a targets file |
| `--revision <revision>` | Revision to reset files to |
| `--last-merged-from <branch>` | Reset to last merge point from this branch |

### `lore file diff` Options

| Option | Description |
|--------|-------------|
| `--source <revision_source>` | Source revision (default: current revision) |
| `--target <revision_target>` | Target revision (default: current filesystem state) |
| `--diff3` | Three-way merge output with conflict markers |
| `-U`, `--context <n>` | Context lines (default: 3) |
| `--ignore-space-at-eol` | Ignore trailing whitespace |
| `--ignore-space-change` | Collapse internal whitespace |
| `--targets <file>` | Path to a targets file |

## `lore auth` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `login` | `login [OPTIONS] [remote-url]` | Authenticate the CLI |
| `info` | `info [OPTIONS] [user-id]...` | Display identity information |
| `list` | `list [OPTIONS]` | List stored identities |
| `logout` | `logout [OPTIONS]` | Remove auth tokens |
| `clear` | `clear` | Clear all auth data |

Auth `login` options: `--token-type` (`api-key`, `eg1`, `lore`), `--token`, `--auth-url`, `--no-browser`.

## `lore layer` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `add` | `add [OPTIONS] <path> <repository> <path>` | Add a repository layer |
| `remove` | `remove [OPTIONS] <path> [repository]` | Remove a repository layer |
| `list` | `list` | List repository layers |

## `lore link` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `add` | `add [OPTIONS] <link_path> <link_url> <source_path>` | Link a repository |
| `remove` | `remove <link_path>` | Remove a link |
| `update` | `update [OPTIONS] <link_path>` | Update a link pin |
| `list` | `list [OPTIONS]` | List all links |

Link `add`/`update` options: `--pin <pin>` (branch or revision), `--disable-branching`.

## `lore lock` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `acquire` | `acquire <paths\|--branch>` | Acquire lock on file(s) |
| `status` | `status [OPTIONS] [paths]...` | Get lock status |
| `query` | `query [OPTIONS]` | Query locks by branch, owner, or path |
| `release` | `release [OPTIONS] [paths]...` | Release lock on file(s) |

## `lore login`

**Usage:** `lore login [OPTIONS] [remote-url]`

Shortcut for CLI authentication (same options as `lore auth login`).

## `lore service` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `run` | `run` | Run this process as the service |
| `start` | `start` | Start service for a repository |
| `stop` | `stop [all]` | Stop service for a repository |

## `lore notification`

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `subscribe` | `subscribe [seconds]` | Subscribe to repository events |

## `lore completions`

**Usage:** `lore completions <shell> [path]`

Generate terminal autocompletions. Shell: `bash`, `elvish`, `fish`, `powershell`, `zsh`.

## `lore shared-store` Commands

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `create` | `create [OPTIONS] <remote-url>` | Create a shared store |
| `info` | `info` | Show shared store information |
| `set-use-automatically` | `set-use-automatically <enabled>` | Toggle automatic use |