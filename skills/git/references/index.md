# The Git Index (Staging Area)

The index is Git's staging area — a flat, binary-format file that represents the proposed next commit. It tracks which file versions will be included in the next snapshot.

> **Source of truth**: <https://git-scm.com/book/en/v2/Git-Internals-The-Index>, <https://git-scm.com/docs/gitformat-index>. For **core concepts**, see [concepts.md](concepts.md). For **object types**, see [objects.md](objects.md).

## What the Index Is

The index is a **binary file** at `.git/index` that maps pathnames to object IDs and stat information:

```
path → (mode, OID, stat_info)
```

It is **not** a diff — it is a complete snapshot of the tree. Every tracked file has an entry, even if unchanged from HEAD.

### The Three Trees

At any point, there are three "tree" states:

```
HEAD commit    Index    Working Tree
   │             │          │
   │  (last)     │ (next)   │ (current)
   │  commit     │ commit   │ files
```

- `git diff` compares index vs working tree
- `git diff --cached` compares HEAD vs index
- `git diff HEAD` compares HEAD vs working tree

## Index Format

The `.git/index` file is a binary structure:

```
DIRC header (4 bytes "DIRC" + version + number of entries)
Index entries (sorted by pathname)
Extensions (optional)
SHA-1 checksum (20 bytes)
```

### Index Versions

| Version | Changes |
|---------|---------|
| 2 | Base format |
| 3 | Extended flags (skip-worktree, assume-unchanged, intent-to-add) |
| 4 | Variable-width path compression (prefix compression) |

### Entry Structure (v2)

Each entry contains:

| Field | Size | Description |
|-------|------|-------------|
| ctime | 8 bytes | File change time (sec + nsec) |
| mtime | 8 bytes | File modification time (sec + nsec) |
| dev | 4 bytes | Device number |
| ino | 4 bytes | Inode number |
| mode | 4 bytes | File mode (040000, 100644, 100755, 120000, 160000) |
| uid | 4 bytes | User ID |
| gid | 4 bytes | Group ID |
| file size | 4 bytes | File size on disk |
| OID | 20 bytes | SHA-1 of the blob object |
| flags | 2 bytes | Name length + stage + extended flag |
| name | Variable | Pathname (NUL-terminated) |

### Stage Numbers

The 2-bit stage field is used during merge conflicts:

| Stage | Meaning |
|-------|---------|
| 0 | Normal entry (no conflict) |
| 1 | Base (common ancestor) |
| 2 | Ours (current branch) |
| 3 | Theirs (merged branch) |

Unmerged entries have stages 1–3 simultaneously for the same path. `git ls-files -u` shows unmerged entries.

## Index Operations

### Staging Files

```bash
git add <path>           # Stage file (create blob, update index entry)
git add -p <path>        # Interactive patch-mode staging
git add -A               # Stage all changes (new, modified, deleted)
git add -u               # Stage modified and deleted (not new files)
git add --intent-to-add <path>  # Mark path for future addition (stage 0, no OID)
```

### Unstaging

```bash
git reset HEAD <path>    # Unstage file (copy entry from HEAD to index)
git restore --staged <path>  # Same (Git 2.23+)
git reset               # Unstage everything
```

### Removing from Index

```bash
git rm <path>            # Remove from index AND working tree
git rm --cached <path>   # Remove from index only (keep working tree)
```

### Inspecting the Index

```bash
git status               # Show index vs HEAD and index vs working tree
git ls-files             # List index entries
git ls-files -s          # Show stage numbers and OIDs
git ls-files -u          # Show unmerged entries
git ls-files -v          # Show assume-unchanged/skip-worktree flags
git diff-index HEAD      # Compare index to HEAD
git diff-files           # Compare working tree to index
```

## Special Index Flags

### skip-worktree

Marks a file to be **ignored for changes** in the working tree. Git pretends the working tree version matches the index:

```bash
git update-index --skip-worktree <path>    # Set
git update-index --no-skip-worktree <path>  # Unset
```

Use cases:
- Local config files that differ per developer
- Files that should not be tracked for changes (e.g., IDE settings)

### assume-unchanged

Tells Git to **assume the file hasn't changed** (skip stat check for performance):

```bash
git update-index --assume-unchanged <path>    # Set
git update-index --no-assume-unchanged <path>  # Unset
```

Difference from skip-worktree:
- `assume-unchanged` is a **performance hint** — Git may still detect changes
- `skip-worktree` is a **user intent** — Git should not touch the file

### intent-to-add

Marks a path as "will be added later" without staging content:

```bash
git add --intent-to-add <path>
```

The file shows as "new file" in `git status` but has no blob OID in the index. `git diff` will show the entire file content as new.

## Sparse-Checkout

Sparse-checkout limits which files appear in the working tree by controlling which index entries are checked out:

```bash
git sparse-checkout init --cone    # Enable cone-mode sparse checkout
git sparse-checkout set src/ docs/ # Only these directories
git sparse-checkout list           # Show current patterns
git sparse-checkout disable        # Restore full checkout
```

### Cone Mode

In cone mode, only directory-level patterns are allowed — much faster than pattern matching:

```
/*              # Root-level files
/src/           # Everything under src/
/docs/          # Everything under docs/
```

### Non-Cone Mode

Full `.gitignore`-style patterns:

```
*.md
!README.md
src/**/*.go
```

The index still contains all entries, but entries outside the sparse-checkout cone have the `SKIP_WORKTREE` flag set.

## Index Extensions

The index format supports extensions after the entries:

| Extension | Code | Purpose |
|-----------|------|---------|
| Tree cache | `TREE` | Caches tree objects for faster `git status` |
| Resolve undo | `REUC` | Records merge conflict resolutions for `git rerere` |
| Link | `link` | Points to a separate index file (for split index) |
| Untracked cache | `UNTR` | Caches untracked file info for faster `git status` |
| FS Monitor | `FSMN` | Integration with fsmonitor-watchman for fast status |
| Sparse directories | `sdir` | Lists directories skipped by sparse-checkout |

## Split Index

For large indices, Git can split the index into a shared base + per-session extension:

```bash
git update-index --split-index
```

The `link` extension in the index points to a shared base index file (`.git/sharedindex.*`). This avoids rewriting the entire index on every `git add`. Note: `core.splitIndex` is deprecated in Git 2.37+; use `git update-index --split-index` instead.

## Edge Cases and Known Issues

- **Index corruption**: If `.git/index` is corrupted, `git read-tree HEAD` rebuilds it from the HEAD commit. You lose staged changes.
- **Line ending conversion**: `core.autocrlf` and `core.safecrlf` affect what goes into the index. The index stores the "normalized" form; the working tree gets the OS-specific form.
- **Large repos**: The index is a single file that must be rewritten on every `git add`. For repos with millions of files, this can be slow. Split index and FSMonitor help.
- **Stage conflicts**: During a merge conflict, the same path has entries at stages 1, 2, and 3. `git add` on the path resolves the conflict by replacing all stages with a single stage-0 entry.
- **Intent-to-add and diff**: `git diff` with intent-to-add files shows the entire file as new content, which can be confusing.
- **Index version 4**: Path compression in v4 can cause compatibility issues with older Git versions.

## References

- <https://git-scm.com/docs/gitformat-index>
- <https://git-scm.com/book/en/v2/Git-Internals-The-Index>

## Related Skills

- For core concepts (three states), see [concepts.md](concepts.md)
- For object types referenced by index entries, see [objects.md](objects.md)
- For sparse-checkout and partial clone, see [worktree.md](worktree.md)
