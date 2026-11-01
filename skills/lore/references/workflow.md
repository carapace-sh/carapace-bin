# Workflows

Common Lore workflows — creating a repository, staging, committing, branching, merging, syncing, and more.

> **Source of truth**: <https://epicgames.github.io/lore/tutorials/quickstart/>

## Quickstart Workflow

### 1. Create a Repository

```bash
lore repository create lore://127.0.0.1:41337/my-project
```

This creates the repo on the server AND sets up a local working tree with `.lore/` directory.

### 2. Add Files and Stage

```bash
# Stage new files, edits, or deletions
lore stage hello.txt sample.bin

# Stage a rename (preserves file identity)
lore stage move old.txt new.txt

# Check what's staged
lore status --scan
```

`lore stage` covers adds, edits, and deletes in a single command. `A` prefix means newly added.

### 3. Commit

```bash
lore commit "Initial revision"
```

Staging and committing work **fully offline** — no server round-trip needed.

### 4. Push

```bash
lore push
```

Uploads local revisions to the server. If push fails with a conflict, run `lore sync`, resolve, commit, and push again.

**Push protocol** (two-phase):
1. Upload fragments: client enumerates, queries remote for missing, uploads the rest.
2. Advance latest pointer: single conditional put on the branch's latest pointer.

### 5. Set Up a Shared Store and Clone

```bash
# Create a shared store (URL has host+port only, no repo path)
lore shared-store create lore://127.0.0.1:41337

# Clone into a sibling directory
lore clone lore://127.0.0.1:41337/my-project my-project-b --use-shared-store
```

### 6. Branch

```bash
# Create and switch to a new branch
lore branch create my-feature

# Switch back to main
lore branch switch main

# Switch to a specific revision on a branch
lore branch switch my-feature <revision>
```

### 7. Merge

```bash
# Merge source branch into current branch
lore branch merge my-feature --message "Merge my-feature into main"
```

Clean merges auto-commit. If conflicts occur, Lore stages the merge:

```bash
# Resolve conflicts
lore branch merge resolve --targets conflicts.txt
# Or use mine/theirs
lore branch merge resolve mine
lore branch merge resolve theirs

# Abort
lore branch merge abort
```

### 8. Sync

```bash
cd ~/my-project-b
lore sync
```

Pulls remote changes. Delta calculation computes what changed between revisions.

## Common Operations

### Amend a Commit Message

```bash
lore revision amend "Updated commit message"
# Or via shortcut
lore commit --amend "Updated commit message"  # if available
```

### View History

```bash
# Repository history
lore history --oneline

# File history
lore file history main.cpp --oneline

# Revision history starting from a specific point
lore revision history --revision <hash> --only-branch
```

### View Diffs

```bash
# Working tree diff
lore diff

# Diff between revisions
lore revision diff <source> --target <target>

# Branch diff
lore branch diff main --source my-feature

# File diff with context
lore file diff main.cpp -U5
```

### Cherry-Pick

```bash
lore revision cherry-pick <revision>
```

Supports `unresolve`, `restart`, `resolve` (with `mine`/`theirs`), and `abort` subcommands.

### Revert

```bash
lore revision revert <revision>
```

Same conflict resolution subcommands as cherry-pick.

### Bisect

```bash
lore revision bisect --start <good> --end <bad>
```

Binary search for a change introduced between start (exclusive) and end (inclusive).

### Find a Revision

```bash
# By metadata
lore revision find metadata <key> [value]

# By revision number
lore revision find number <NUMBER>
```

### Restore

```bash
lore revision restore "Restore to current revision"
```

### Obliterate Content

```bash
# By file path
lore file obliterate --path path/to/file.ext

# By blob address
lore file obliterate --address <hash>
```

### Reset Files

```bash
# Reset working tree files to current revision
lore reset file.txt

# Reset to a specific revision
lore reset --revision <hash> file.txt

# Reset to the last merge point from a branch
lore reset --last-merged-from main file.txt

# Purge untracked files
lore reset --purge .
```

### File Dependencies

```bash
# Add dependency edges
lore file dependency add source.txt dep1.txt dep2.txt --tag build

# List dependencies
lore file dependency list source.txt --recursive --tag build

# List dependents (reverse)
lore file dependency list source.txt --reverse

# Remove dependencies
lore file dependency remove source.txt dep1.txt --tag build
```

### File Locking

```bash
# Acquire lock
lore lock acquire file.bin --branch main

# Check lock status
lore lock status file.bin

# Query locks by branch, owner, or path
lore lock query --branch main --owner <user-id>

# Release lock
lore lock release file.bin
```

### File Info

```bash
lore file info path/to/file --revision <hash>
lore file info path/to/file --local    # local filesystem hash/size
lore file info path/to/file --filtered # filtered repo size
```

### File Hash

```bash
lore file hash file.bin
```

### Write Data from Store

```bash
lore file write --address <hash> --output /path/to/output
lore file write --path file.txt --revision <hash> --output /path/to/output
```

## Sparse Checkouts

Files outside the view filter (`.lore/view`) are not materialized on disk. Only the parts of the tree the view asks for are walked. Fragments are fetched on demand:

- Reading a 4 MiB range from a multi-gigabyte file fetches only the overlapping fragments.
- The view filter is local to a client — it doesn't travel with a clone.

## Ignore Files

`.loreignore` is an outbound filter — paths matching patterns in it are excluded from staging and committing. It's combined with the view filter: operations on committed state consult the view only; operations on local state consult both.

## Identity

Set commit identity at create/clone time:

```bash
lore repository create lore://host:port/repo --identity you@example.com
```

If unset, Lore fails with: *"No commit identity configured; pass --identity or set identity in .lore/config.toml"*.