# Operations Log — Undo, Redo, and Snapshots

The operations log tracks every workspace operation, enabling undo/redo and restoration to any previous state.

> **Source of truth**: <https://docs.gitbutler.com/cli-guides/cli-tutorial/operations-log>

## Concept

Every operation that modifies the workspace (commits, amends, squashes, moves, branch operations, discards) is recorded as an **oplog entry**. Each entry has an ID, timestamp, operation type, and title. You can restore the workspace to any previous entry.

## `but oplog` — Show Operation History

```bash
but oplog                # Show operation history
```

Parent command for oplog subcommands. Without a subcommand, lists operations.

### `but oplog list`

List operation history entries.

```bash
but oplog list
but oplog list --since <entry-id>
```

| Flag | Description |
|------|-------------|
| `--since` | Show entries since this oplog entry |

### `but oplog restore`

Restore the workspace to a specific oplog snapshot.

```bash
but oplog restore <entry-id>
```

Positional argument is the oplog entry ID to restore to.

### `but oplog snapshot`

Create a manual savepoint in the operation history.

```bash
but oplog snapshot
```

Useful before large history edits (squashing, moving multiple commits) as a recovery point.

## `but undo` — Undo Last Operation

```bash
but undo
```

Reverts to the previous snapshot. Undo the last operation by reverting to the previous state. Multiple `but undo` calls walk backwards through history.

## `but redo` — Redo Last Undo

```bash
but redo
```

Redo the last undo. Walks forward through history after undoing.

## Operation Types

Each oplog entry has an operation type:

| Operation | Style | Description |
|-----------|-------|-------------|
| `CreateCommit` | green | New commit created |
| `CreateBranch` | magenta | New branch created |
| `AmendCommit` | yellow | Commit amended |
| `UndoCommit` | red | Commit undone |
| `SquashCommit` | default | Commits squashed |
| `UpdateCommitMessage` | yellow | Commit message edited (reword) |
| `MoveCommit` | cyan | Commit moved |
| `RestoreFromSnapshot` | red | Restored from snapshot |
| `ReorderCommit` | cyan | Commit reordered |
| `InsertBlankCommit` | default | Blank commit inserted |
| `MoveHunk` | cyan | Hunk moved between commits |
| `ReorderBranches` | default | Branches reordered |
| `UpdateWorkspaceBase` | default | Workspace base updated |
| `UpdateBranchName` | default | Branch renamed |
| `GenericBranchUpdate` | default | Branch updated |
| `ApplyBranch` | default | Branch applied |
| `UnapplyBranch` | default | Branch unapplied |
| `DeleteBranch` | default | Branch deleted |
| `DiscardChanges` | default | Changes discarded |

## JSON Output

```bash
but --format json oplog
```

Returns an array of entries with `id`, `createdAt`, and `details` (containing `operation`, `title`, `body`, and `trailers`).

## Workflow

```
but oplog snapshot          # Create savepoint before risky edit
but squash <a> <b>          # Perform the edit
but undo                    # If something went wrong, undo
but oplog restore <id>      # Or restore to specific point
but redo                    # Redo if you undid too far
```

## Related

- [commits.md](commits.md) — Commit operations that generate oplog entries
- [conflicts.md](conflicts.md) — Conflict resolution operations
