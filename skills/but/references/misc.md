# Miscellaneous — Root Command, Global Flags, Update, Completions

The root `but` command, global flags, command groups, update management, and shell completions.

> **Source of truth**: <https://docs.gitbutler.com/cli-overview>

## Root Command

```bash
but [global flags] [command] [args]
```

The root command is `but` — "A GitButler CLI tool". Running `but` with no arguments runs the `default` alias (which expands to `status`).

### Global Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--current-dir` | `-C` | Run as if `but` was started in PATH instead of the current working directory |
| `--format` | | Explicitly control how output should be formatted (persistent) |
| `--help` | `-h` | Print help |
| `--version` | `-V` | Print version |

The `--current-dir` flag changes the working directory for all operations. The `--format` flag is persistent (available on all subcommands).

### Output Formats

| Format | Description |
|--------|-------------|
| `human` | Default human-readable output |
| `agent` | Formatted for AI agent consumption |
| `shell` | Shell-friendly output |
| `json` | Machine-parseable JSON |
| `none` | No output |

### Positional Arguments

The root command accepts positional arguments without a subcommand. These are CLI IDs, directory paths, or branch names — enabling direct operations like `but <cli-id> <target>` which routes through the alias/default system.

## Command Groups

Commands are organized into groups shown in `--help` output:

| Group ID | Commands |
|----------|----------|
| `inspection` | `status`, `diff`, `show` |
| `branching and committing` | `apply`, `unapply`, `branch`, `clean`, `commit`, `discard`, `merge`, `pick`, `resolve`, `stage` |
| `editing commits` | `absorb`, `amend`, `move`, `reword`, `rub`, `squash`, `uncommit` |
| `operation history` | `oplog`, `redo`, `undo` |
| `server interactions` | `pr`, `pull`, `push` |
| `rules` | _(empty — mark/unmark removed in 0.20.4)_ |
| `other commands` | `alias`, `config`, `gui`, `setup`, `skill`, `teardown`, `tui`, `update`, `agent`, `completions` |

## Dynamic Aliases

User-defined and default aliases are registered as commands at runtime. This means `but st` works exactly like `but status`.

## `but completions` — Generate Shell Completions

```bash
but completions
```

Generate shell completion scripts for the `but` command. Now appears in `--help` output (since 0.20.4), making it discoverable when using the `but` binary without an installer.

## `but update` — Manage Updates

Parent command for CLI and app update management.

### `but update check`

Check if a new version of the GitButler CLI is available.

```bash
but update check
```

### `but update install`

Install or update the GitButler desktop application.

```bash
but update install
```

### `but update suppress`

Suppress update notifications temporarily.

```bash
but update suppress
```

## `but gui` — Open Desktop GUI

```bash
but gui
but gui <path>
```

Open the GitButler desktop GUI for the current project. Accepts an optional project path as a positional argument.

## `but merge` — Merge Branch

Merge a branch into your local target branch.

```bash
but merge <branch>
```

Positional argument is the local branch name to merge.

## Related

- [config.md](config.md) — Configuration and aliases
- [workspace.md](workspace.md) — Workspace model and setup/teardown
- [inspection.md](inspection.md) — `--format` flag and output formats
