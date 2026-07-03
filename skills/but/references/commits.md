# Commits and Staging

Commands for creating, editing, and managing commits in the GitButler workspace.

> **Source of truth**: <https://docs.gitbutler.com/cli-guides/cli-tutorial/branching-and-commiting> and <https://docs.gitbutler.com/cli-guides/cli-tutorial/editing-commits>

## `but commit` — Create a Commit

Commit changes to a stack (virtual branch). This is the primary commit command.

```bash
but commit -m "message" <branch>
but commit --ai -m "instructions" <branch>
but commit -p h0,i0 -m "message" <branch>
but commit --only -m "message" <branch>
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--after` | | Insert the commit after this commit or branch |
| `--ai` | `-i` | Generate commit message using AI with optional user summary |
| `--all` | `-a` | No-op compatibility flag for `git commit -a` |
| `--before` | | Insert the commit before this commit or branch |
| `--changes` | `-p` | Uncommitted file or hunk CLI IDs to include (comma-separated or repeated) |
| `--create` | `-c` | Create a new branch for this commit |
| `--diff` | | Always show diff inside the editor |
| `--message` | `-m` | Commit message |
| `--message-file` | | Read commit message from file |
| `--no-diff` | | Never show the diff inside the editor |
| `--no-hooks` | `-n` | Bypass pre-commit hooks |
| `--no-verify` | | Bypass pre-commit hooks (hidden alias for `--no-hooks`) |
| `--only` | `-o` | Only commit staged files, not unstaged files |

The `--ai` flag uses `NoOptDefVal = " "`, so `--ai` by itself generates a message, or `--ai="your instructions"` provides guidance.

The `--after` and `--before` flags insert commits at specific positions in a stack, enabling fine-grained commit ordering without rebase.

### Positional Argument

The positional argument is the branch name to commit to. If `--create` is used and the branch doesn't exist, it's created.

### Committing Specific Files/Hunks

Use `--changes` (`-p`) with CLI IDs to commit specific files or hunks:

```bash
but commit -p h0,i0 -m "message" <branch>     # Specific hunks by CLI ID
but commit -p app/models/foo.rb -m "msg"      # By file path (via git changes)
```

CLI IDs are multi-part identifiers (e.g. `h0`, `i0`, `c3`) shown in `but status -f`. See [inspection.md](inspection.md) for the CLI ID system.

## `but commit empty` — Insert Blank Commit

Insert a blank commit before or after a specified commit. Used for splitting commits.

```bash
but commit empty --after <sha>           # Insert blank commit after
but commit empty --before <sha>          # Insert blank commit before
but commit empty --after <sha> -m "msg"  # With a message (since 0.20.4)
```

| Flag | Short | Description |
|------|-------|-------------|
| `--after` | | Insert the blank commit after this commit or branch |
| `--before` | | Insert the blank commit before this commit or branch |
| `--message` | `-m` | Commit message for the inserted blank commit |

After creating a blank commit, use `but rub` to move changes into it. See [rubbing.md](rubbing.md).

## `but stage` — Stage Changes to a Branch

Assign file or hunk changes to a specific branch. This is like `git add` but for multiple virtual branches.

```bash
but stage <file> <branch>
but stage -b <branch>                     # Interactive mode (specify branch first)
but stage                                  # Interactive hunk picker (no args)
```

| Flag | Short | Description |
|------|-------|-------------|
| `--branch` | `-b` | Branch to stage to (for interactive mode) |

The positional arguments are: changes (CLI IDs or file paths, comma-separated) and branch name.

Running `but stage` without arguments opens the interactive hunk picker in the TUI.

## `but amend` — Amend Changes into a Commit

Amend one or more file changes into a specific commit. Rebases any dependent commits.

```bash
but amend <commit>
but amend -p h0,i0 <commit>               # Amend specific changes
```

| Flag | Short | Description |
|------|-------|-------------|
| `--changes` | `-p` | Uncommitted file or hunk CLI IDs to amend into the commit |

The positional argument is the commit to amend into (commit SHA or CLI ID).

## `but squash` — Squash Commits Together

Squash two commits into one.

```bash
but squash <source-commit> <target-commit>
but squash --ai <source> <target>          # AI-generated combined message
but squash -d <source> <target>            # Drop source message, keep target's
```

| Flag | Short | Description |
|------|-------|-------------|
| `--ai` | `-i` | Generate commit message using AI |
| `--drop-message` | `-d` | Drop source commit messages, keep only target's |
| `--message` | `-m` | Provide a new commit message for the resulting commit |

Two positional arguments: source commit and target commit (commit SHAs or CLI IDs).

## `but reword` — Edit Commit Message

Edit the commit message of the specified commit. Opens an editor.

```bash
but reword <commit>
```

Can also rename a branch:

```bash
but reword -m <new-name> <branch>
```

## `but move` — Move a Commit or Branch

Move a commit or branch to a different location.

```bash
but move <commit> <target>
```

| Flag | Short | Description |
|------|-------|-------------|
| `--after` | | Move after this commit or branch |

Two positional arguments: the entity to move (commit) and the target (commit or branch).

## `but pick` — Cherry-Pick a Commit

Cherry-pick a commit from an unapplied branch into an applied virtual branch.

```bash
but pick <commit> <branch>
```

Two positional arguments: the commit to pick and the target branch.

## `but absorb` — Auto-Amend Changes

Amends changes into the appropriate commits where they belong. GitButler automatically determines which commits each change should be amended into.

```bash
but absorb
but absorb --dry-run                       # Show the absorption plan without changes
```

| Flag | Description |
|------|-------------|
| `--dry-run` | Show the absorption plan without making any changes |

## `but uncommit` — Uncommit Changes

Uncommit changes from a commit to the unstaged area.

```bash
but uncommit <commit>
but uncommit --diff <commit>               # Show resulting uncommitted diff (since 0.20.4)
but uncommit --discard <commit>            # Discard instead of moving to uncommitted
```

| Flag | Short | Description |
|------|-------|-------------|
| `--diff` | | Show the resulting uncommitted diff after uncommitting |
| `--discard` | `-d` | Discard the selected committed changes instead of moving them to uncommitted |

## `but discard` — Discard Uncommitted Changes

Discard uncommitted changes from the worktree.

```bash
but discard <file>
```

Positional arguments are file paths or CLI IDs of changes to discard.

## Related

- [rubbing.md](rubbing.md) — The `but rub` command for combining entities
- [inspection.md](inspection.md) — CLI IDs and status output
- [branches.md](branches.md) — Branch management
- [oplog.md](oplog.md) — Undo/redo for commit operations
