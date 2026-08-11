# Operational Commands

Top-level operational shortcuts and infrastructure commands — status, clone, stage, dirty, unstage, reset, diff, history, commit, sync, push, service, notification, completions, and shared-store.

> **Source of truth**: <https://epicgames.github.io/lore/reference/lore-cli-commands/>

## Top-Level Shortcuts

These commands are shortcuts that delegate to the corresponding `lore repository` / `lore file` / `lore revision` subcommands. They share the same options.

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

## `lore status`

**Usage:** `lore status [OPTIONS] [PATH]...`

Same options as `lore repository status`.

## `lore clone`

**Usage:** `lore clone [OPTIONS] <url> [path]`

Same options as `lore repository clone`.

## `lore stage`

**Usage:** `lore stage [OPTIONS] <paths|--targets <file>>`

Same options as `lore file stage`, including `move` and `merge` subcommands.

## `lore dirty`

**Usage:** `lore dirty [OPTIONS] [paths]...`

Same options as `lore file dirty`, including `move` and `copy` subcommands.

## `lore unstage`

**Usage:** `lore unstage <paths|--targets <file>>`

Same options as `lore file unstage`.

## `lore reset`

**Usage:** `lore reset [OPTIONS] <paths|--targets <file>>`

Same options as `lore file reset`.

## `lore diff`

**Usage:** `lore diff [OPTIONS] [paths]...`

Same options as `lore file diff`.

## `lore history`

**Usage:** `lore history [OPTIONS] [LENGTH]`

Same options as `lore revision history`.

## `lore commit`

**Usage:** `lore commit [OPTIONS] <MESSAGE>`

Same options as `lore revision commit`.

## `lore sync`

**Usage:** `lore sync [OPTIONS] [revision]`

**Alias:** `synchronize`

Same options as `lore revision sync`.

## `lore push`

**Usage:** `lore push [OPTIONS] [branch]`

Same options as `lore branch push`.

## Service

### `lore service run`

**Usage:** `lore service run`

Run this process as the service.

### `lore service start`

**Usage:** `lore service start`

Start service for a repository.

### `lore service stop`

**Usage:** `lore service stop [all]`

Stop service for a repository. Pass `all` to stop servicing all repositories.

## Notification

### `lore notification subscribe`

**Usage:** `lore notification subscribe [seconds]`

Subscribe to events on the given repository. Optionally specify duration in seconds.

## Completions

### `lore completions`

**Usage:** `lore completions <shell> [path]`

Generate terminal autocompletions.

| Argument | Description |
|----------|-------------|
| `<shell>` | Shell to generate for: `bash`, `elvish`, `fish`, `powershell`, `zsh` |
| `<path>` | Directory path to write the autocompletion script to |

## Shared Store

### `lore shared-store create`

**Usage:** `lore shared-store create [OPTIONS] <remote-url>`

Create a shared store backed by a remote URL.

| Option | Description |
|--------|-------------|
| `--path <path>` | Where to create the shared store |
| `--make-default <MAKE_DEFAULT>` | Set as default in global config (default: true; values: `true`, `false`) |

### `lore shared-store info`

**Usage:** `lore shared-store info`

Show shared store information.

### `lore shared-store set-use-automatically`

**Usage:** `lore shared-store set-use-automatically <enabled>`

Set whether to automatically use the shared store (values: `true`, `false`).