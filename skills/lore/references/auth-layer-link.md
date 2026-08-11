# Auth, Layer, Link, Lock, and Logfile Commands

Commands for authentication, layer management, repository linking, file locking, and logfile inspection.

> **Source of truth**: <https://epicgames.github.io/lore/reference/lore-cli-commands/>

## Auth

### `lore auth login`

**Usage:** `lore auth login [OPTIONS] [remote-url]`

Authenticate the CLI.

| Option | Description |
|--------|-------------|
| `--token-type <TOKEN_TYPE>` | Token type for non-interactive login (`api-key`, `eg1`, `lore`) |
| `--token <TOKEN>` | Token value for non-interactive login (requires `--token-type`) |
| `--auth-url <AUTH_URL>` | Auth service URL with scheme |
| `--no-browser` | Avoid opening a browser to login |

### `lore auth info`

**Usage:** `lore auth info [OPTIONS] [user-id]...`

Display identity information. Omit user-id for current user.

| Option | Description |
|--------|-------------|
| `--with-token` | Include cached tokens in the output |

### `lore auth list`

**Usage:** `lore auth list [OPTIONS]`

List all stored authentication identities.

| Option | Description |
|--------|-------------|
| `--with-token` | Include cached tokens in the output |

### `lore auth logout`

**Usage:** `lore auth logout [OPTIONS]`

Remove stored authentication and authorization tokens.

| Option | Description |
|--------|-------------|
| `--auth-url <auth-url>` | Auth service URL (omit to use current repository's auth URL) |
| `--resource <resource>` | Resource ID to remove a specific authorization |
| `--user-id <user-id>` | User ID to remove (omit to remove all identities) |

### `lore auth clear`

**Usage:** `lore auth clear`

Clear all stored authentication data.

## `lore login`

**Usage:** `lore login [OPTIONS] [remote-url]`

Shortcut for CLI authentication (same options as `lore auth login`).

## Layer

### `lore layer add`

**Usage:** `lore layer add [OPTIONS] <path> <repository> <path>`

Add a repository layer. Arguments: path in current repo, repository to add as layer, path in layer repo.

| Option | Description |
|--------|-------------|
| `--metadata <metadata>` | Metadata key to use for matching revisions |

### `lore layer remove`

**Usage:** `lore layer remove [OPTIONS] <path> [repository]`

Remove a repository layer.

| Option | Description |
|--------|-------------|
| `--purge` | Also delete untracked files and all directories inside the layer mount |

### `lore layer list`

**Usage:** `lore layer list`

List repository layers.

## Link

### `lore link add`

**Usage:** `lore link add [OPTIONS] <link_path> <link_url> <source_path>`

Link to the given point in the repository and subpath from the given repository.

| Option | Description |
|--------|-------------|
| `--pin <pin>` | Branch or specific revision to pin the link to (default: latest on main branch) |
| `--disable-branching` | Disable automatic branch creation in the linked repository |

### `lore link remove`

**Usage:** `lore link remove <link_path>`

Remove the link at the given point in the repository.

### `lore link update`

**Usage:** `lore link update [OPTIONS] <link_path>`

Update the link to a new pin.

| Option | Description |
|--------|-------------|
| `--pin <pin>` | Branch or specific revision to pin the link to (default: latest on current branch) |

### `lore link list`

**Usage:** `lore link list [OPTIONS]`

List all links in the repository.

| Option | Description |
|--------|-------------|
| `--staged` | Only show links with staged changes |

## Lock

### `lore lock acquire`

**Usage:** `lore lock acquire <paths|--branch <branch>>`

Acquire lock on file(s).

| Option | Description |
|--------|-------------|
| `--branch <branch>` | Branch where lock is to be acquired |

### `lore lock status`

**Usage:** `lore lock status [OPTIONS] [paths]...`

Get lock status on file(s).

| Option | Description |
|--------|-------------|
| `--branch <branch>` | Branch where lock was acquired |

### `lore lock query`

**Usage:** `lore lock query [OPTIONS]`

Query lock status by branch, owner, or path.

| Option | Description |
|--------|-------------|
| `--branch <branch-name>` | Branch to query locks on |
| `--owner <owner-id>` | Owner to query locks belonging to them |
| `--path <path>` | Path to query lock information on |

### `lore lock release`

**Usage:** `lore lock release [OPTIONS] [paths]...`

Release lock on file(s).

| Option | Description |
|--------|-------------|
| `--branch <branch>` | Branch where lock was acquired |
| `--owner <owner>` | Owner of the lock |

## Logfile

### `lore logfile info`

**Usage:** `lore logfile info`

Show logfile information.