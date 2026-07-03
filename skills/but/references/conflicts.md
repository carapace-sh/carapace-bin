# Conflict Resolution

How GitButler handles conflicts during rebase, pull, and branch update operations.

> **Source of truth**: <https://docs.gitbutler.com/cli-guides/cli-tutorial/conflict-resolution>

## Key Difference from Git

Rebases **always succeed** in GitButler. Conflicts don't halt the operation — instead, conflicted commits are marked as conflicted (shown with the `◐` icon in status output). You resolve conflicts one commit at a time without stopping the rebase.

## Conflict Markers

GitButler uses **zdiff3** style conflict markers, which show ours/theirs/ancestor:

```
<<<<<<< ours
our version
||||||| ancestor
original version
=======
their version
>>>>>>> theirs
```

## `but resolve` — Conflict Resolution Mode

Parent command for conflict resolution operations.

### `but resolve`

Enter conflict resolution mode. This begins the interactive resolution process for conflicted commits.

### `but resolve finish`

Finalize the conflict resolution. Commits the resolved state and marks the conflict as resolved.

```bash
but resolve finish
```

### `but resolve cancel`

Cancel conflict resolution mode. Reverts to the pre-resolution state.

```bash
but resolve cancel
```

### `but resolve status`

Show the current conflict resolution status — which commits are conflicted and what remains to resolve.

```bash
but resolve status
```

## Workflow

```
but pull                           # Pull triggers a rebase with conflicts
but status                         # See which commits are conflicted (◐ icon)
but resolve                        # Enter resolution mode
  # edit files to resolve conflicts
but resolve finish                 # Finalize resolution
  # or
but resolve cancel                 # Cancel and revert
```

## Conflict Detection in Other Commands

`but branch show --check` can check if a branch merges cleanly into upstream and identify conflicting commits:

```bash
but branch show --check <branch>
```

## Related

- [oplog.md](oplog.md) — Undo conflict resolution with `but undo`
- [branches.md](branches.md) — `but branch update` can trigger conflicts
- [forge.md](forge.md) — `but pull` can trigger conflicts
