# Rubbing — The `but rub` Command

`but rub` is GitButler's powerful commit editing tool. "Rubbing" is combining two entities together — depending on what you rub into what, different operations happen automatically.

> **Source of truth**: <https://docs.gitbutler.com/cli-guides/cli-tutorial/rubbing>

## Concept

The key insight is that all commit editing operations (amend, squash, move, stage, uncommit) are fundamentally the same operation: **combine entity A with entity B**. The type of each entity determines the operation.

## Operation Matrix

| Source (arg 1) | Target (arg 2) | Operation |
|----------------|----------------|-----------|
| File/hunk change | Branch | **Stage** — assign changes to a branch |
| File/hunk change | `zz` | **Unassign** — move changes back to unassigned |
| File/hunk change | Commit | **Amend** — add changes into an existing commit |
| Commit | Commit | **Squash** — merge two commits into one |
| Commit | `zz` | **Uncommit** — put commit changes back to unassigned |
| Commit | Branch | **Move** — move a commit to another branch |
| File change ID | Commit | **Move file** — move a specific file change from one commit to another |

## Usage

```bash
but rub <source> <target>
```

Two positional arguments, both can be CLI IDs or names:

- **Arg 1 (source)**: CLI IDs (changes, commits, branches) or local branch names
- **Arg 2 (target)**: CLI IDs (branches, stacks, commits) or local branch names

### The `zz` Target

`zz` is the special CLI ID for the **unassigned area** — changes that aren't assigned to any branch. Rubbing something into `zz` moves it to unassigned. Rubbing `zz` as a source isn't valid — use `but stage` or `but rub <change> <branch>` instead.

## Common Workflows

### Stage a File to a Branch

```bash
but rub <file-id> <branch>
# Equivalent to: but stage <file> <branch>
```

### Amend a Commit

```bash
but rub <file-id> <commit-sha>
# Equivalent to: but amend -p <file-id> <commit-sha>
```

### Squash Two Commits

```bash
but rub <source-commit> <target-commit>
# Equivalent to: but squash <source> <target>
```

### Move a Commit to Another Branch

```bash
but rub <commit-sha> <branch>
# Equivalent to: but move <commit> <branch>
```

### Uncommit Changes

```bash
but rub <commit-sha> zz
# Equivalent to: but uncommit <commit-sha>
```

### Unassign Changes

```bash
but rub <file-id> zz
# Moves changes back to unassigned
```

### Split a Commit

```bash
but commit empty --after <sha>     # Create blank commit
but rub <file-id> <blank-commit>   # Move specific changes into it
```

## CLI IDs in Rub

`but rub` uses CLI IDs extensively. These are short identifiers (e.g. `h0`, `i0`, `c3`, `b1`) shown in `but status -f` output. The `--format json` output also includes `cliId` fields. See [inspection.md](inspection.md) for the full CLI ID reference.

## Related

- [commits.md](commits.md) — Dedicated commit commands (amend, squash, move, etc.)
- [inspection.md](inspection.md) — CLI IDs and status output
- [tui.md](tui.md) — Interactive rubbing in the TUI
