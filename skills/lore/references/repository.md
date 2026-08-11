# Repository Commands

Commands for managing Lore repositories — creation, cloning, status, verification, garbage collection, store queries, metadata, instances, and configuration.

> **Source of truth**: <https://epicgames.github.io/lore/reference/lore-cli-commands/>

## `lore repository status`

**Usage:** `lore repository status [OPTIONS] [PATH]...`

Show current repository status.

| Option | Description |
|--------|-------------|
| `--scan` | Walk the filesystem under the given paths and reconcile every file against the current revision |
| `--check-dirty` | Verify already-dirty files against the filesystem without a full scan |
| `--reset` | Drop the existing staged anchor before computing status |
| `--revision-only` | Only show revision info, skip all diffs |
| `--count` | Count directories and files |
| `--targets <file>` | Path to a targets file |

## `lore repository info`

**Usage:** `lore repository info [url]`

Get info about a repository. Optional URL for remote lookup.

## `lore repository list`

**Usage:** `lore repository list <url>`

List repositories on a remote.

## `lore repository create`

**Usage:** `lore repository create [OPTIONS] <url>`

Create a repository in the given directory.

| Option | Description |
|--------|-------------|
| `--description <description>` | Optional description of repository |
| `--id <id>` | Optional ID of repository |
| `--use-shared-store` | Use the shared store rather than create a local immutable store |
| `--shared-store-path <SHARED_STORE_PATH>` | Use this path rather than the system default as the shared store location |

## `lore repository clone`

**Usage:** `lore repository clone [OPTIONS] <url> [path]`

Clone a remote repository into the given path.

| Option | Description |
|--------|-------------|
| `--view <view>` | Optional client side view filter file |
| `--revision <revision>` | Optional revision to sync |
| `--branch <branch>` | Optional branch to sync |
| `--bare` | Clone without files, only fetch latest revision tree |
| `--virtual` | Clone virtually using split-write filesystem |
| `--direct-file-write` | Write directly to the destination file instead of temp+move |
| `--layer <repository>` | Layer to add |
| `--layer-metadata <key>` | Metadata key to link layer revisions with |
| `--prefetch <file>` | File containing list of files to prefetch |
| `--use-shared-store` | Use the shared store rather than create a local immutable store |
| `--shared-store-path <SHARED_STORE_PATH>` | Use this path rather than the system default |
| `--no-tracking` | Clone without local repository tracking (memory-only stores) |
| `--root-file <path>` | Root files for dependency-based selective clone |
| `--dependency-tag <tag>` | Tags to filter dependencies by |
| `--dependency-recursive` | Follow transitive dependencies recursively |
| `--dependency-depth-limit <depth>` | Max dependency traversal depth (0 = unlimited) |

## `lore repository delete`

**Usage:** `lore repository delete <url>`

Delete a repository.

## `lore repository verify`

**Usage:** `lore repository verify [OPTIONS] [COMMAND]`

Verify repository state consistency.

| Subcommand | Description |
|------------|-------------|
| `state` | Verify repository state consistency (default) |
| `fragment` | Verify a specific fragment in the local store |

| Option | Description |
|--------|-------------|
| `--path <path>` | Optional path in the repository to start verification from |
| `--heal` | Attempt to heal discrepancies found in a new staged state |

### `lore repository verify state`

**Usage:** `lore repository verify state [OPTIONS]`

Options: `--path <path>`, `--heal`.

### `lore repository verify fragment`

**Usage:** `lore repository verify fragment [OPTIONS] <HASH>`

| Option | Description |
|--------|-------------|
| `--context <CONTEXT>` | Context part of the address to verify |
| `--heal` | Attempt to heal if verification fails (remote only) |

## `lore repository dump`

**Usage:** `lore repository dump [OPTIONS]`

Dump repository state information.

| Option | Description |
|--------|-------------|
| `--path <path>` | Optional path to start dumping from |
| `--revision <revision>` | Optional revision to dump |
| `--max-depth <max-depth>` | Optional max depth of tree dump |

## `lore repository gc`

**Usage:** `lore repository gc`

Run a full garbage collection pass on the local repository store.

## `lore repository store`

**Usage:** `lore repository store immutable query [OPTIONS] <ADDRESS>`

Access the repository data store.

| Subcommand | Description |
|------------|-------------|
| `immutable query` | Query the immutable store by fragment address |

### `lore repository store immutable query`

| Option | Description |
|--------|-------------|
| `--recurse` | Recurse into subfragments |

## `lore repository metadata`

Manage metadata on a repository.

| Subcommand | Usage | Description |
|------------|-------|-------------|
| `get` | `lore repository metadata get [key]` | Get metadata (omit key to list all) |
| `set` | `lore repository metadata set [OPTIONS] [pairs]...` | Set metadata key/value pairs |
| `clear` | `lore repository metadata clear [keys]...` | Clear metadata (omit keys to clear all user-defined) |

Options for `set`: `--binary` (values are paths to binary files), `--numeric` (values are numeric u64).

## `lore repository instance`

Manage registered instances for this repository.

| Subcommand | Description |
|------------|-------------|
| `list` | List all registered instances |
| `prune` | Remove stale instance entries |

## `lore repository config`

**Usage:** `lore repository config get <KEY>`

Read a configuration value.

## `lore repository update-path`

**Usage:** `lore repository update-path`

Update the stored path for this instance.