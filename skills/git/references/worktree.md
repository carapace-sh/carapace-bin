# Git Worktrees and Submodules

Multiple working trees, submodule integration, subtree merges, sparse-checkout, and partial clone strategies for managing large or multi-repo projects.

> **Source of truth**: <https://git-scm.com/docs/git-worktree>, <https://git-scm.com/docs/git-submodule>. For **core concepts**, see [concepts.md](concepts.md). For **reference types**, see [refs.md](refs.md).

## Worktrees

A worktree is an additional working tree attached to the same repository. Each worktree checks out a different branch, but they share the same object database.

### Why Worktrees

- Work on multiple branches simultaneously without stashing
- Run long-running tasks (builds, tests) in one tree while editing in another
- Avoid the overhead of full clones for parallel work

### Creating Worktrees

```bash
git worktree add ../feature-worktree feature    # Check out feature branch
git worktree add -b new-feature ../new-feature main  # Create branch + worktree
git worktree add --detach ../detached abc1234   # Detached HEAD worktree
git worktree add --lock ../locked feature       # Locked (not pruned)
```

### Listing Worktrees

```bash
git worktree list
git worktree list --porcelain    # Machine-readable format
```

### Removing Worktrees

```bash
git worktree remove ../feature-worktree
git worktree remove --force ../feature-worktree  # Force if dirty
git worktree prune                # Prune deleted worktree admin data
```

### Worktree Internals

Each worktree has its own:
- Working tree directory
- Index file (`.git/worktrees/<name>/index`)
- HEAD (`.git/worktrees/<name>/HEAD`)
- Per-worktree config (`.git/config.worktree`)

The main `.git/` directory stores admin data:

```
.git/
├── worktrees/
│   └── feature-worktree/
│       ├── HEAD           # HEAD for this worktree
│       ├── index          # Index for this worktree
│       ├── commondir      # Path to shared .git (relative)
│       └── gitdir         # Path to worktree's .git file
```

The worktree directory contains a `.git` **file** (not directory):

```
gitdir: /path/to/main/.git/worktrees/feature-worktree
```

### Worktree Restrictions

- **No duplicate branches**: The same branch cannot be checked out in multiple worktrees (would cause index/HEAD conflicts)
- **Linked worktrees share objects**: No separate object database — all worktrees use the same `.git/objects/`
- **Worktree config**: Requires `extensions.worktreeConfig = true` to use per-worktree config

## Submodules

A submodule embeds one Git repository inside another at a specific path, pinned to a specific commit.

### Adding Submodules

```bash
git submodule add https://github.com/user/lib.git vendor/lib
git submodule add -b main https://github.com/user/lib.git vendor/lib  # Track branch
```

This creates:
- A `.gitmodules` file (tracked)
- An entry in the index with mode `160000` (gitlink)
- A clone in `vendor/lib/`

### .gitmodules Format

```ini
[submodule "vendor/lib"]
    path = vendor/lib
    url = https://github.com/user/lib.git
    branch = main              # Optional: branch to track
    fetchRecurseSubmodules = true  # Optional: auto-fetch
    ignore = dirty              # Optional: ignore changes
```

### Initializing and Updating

```bash
git submodule init                # Register submodule URLs from .gitmodules
git submodule update             # Check out recorded commits
git submodule update --init       # Init + update
git submodule update --remote    # Update to latest remote tip (not recorded commit)
git submodule update --init --recursive  # Init + update + nested submodules
```

### Cloning with Submodules

```bash
git clone --recurse-submodules <url>
git clone --recurse-submodules=<pathspec> <url>  # Only specific submodules
```

After a regular clone:

```bash
git submodule update --init --recursive
```

### Submodule Status

```bash
git submodule status              # Show current state
git submodule status --recursive  # Include nested submodules
git submodule foreach 'git status'  # Run command in each submodule
git submodule foreach --recursive 'git pull origin main'
```

### Submodule Sync

```bash
git submodule sync               # Update local config from .gitmodules
git submodule sync --recursive   # Include nested
```

### Absorbing Submodule Git Dir

```bash
git submodule absorbgitdirs      # Move .git into parent's .git/modules/
```

This moves the submodule's `.git` directory from `vendor/lib/.git/` to `.git/modules/vendor/lib/`, leaving a gitdir file behind. This is the default since Git 2.12.

### Updating Submodule References

```bash
cd vendor/lib
git pull origin main
cd ../..
git add vendor/lib               # Stage the new commit OID
git commit -m "Update lib submodule"
```

### Removing Submodules

```bash
git submodule deinit -f vendor/lib    # Unregister + clean working tree
git rm -f vendor/lib                  # Remove from index + working tree
rm -rf .git/modules/vendor/lib        # Remove internal data
```

Or with Git 2.17+:

```bash
git rm --submodule=vendor/lib
```

### Submodule Pitfalls

- **Detached HEAD**: Submodules are checked out at a specific commit, not a branch. `git submodule update` puts them in detached HEAD.
- **Dirty submodules**: Uncommitted changes in a submodule are not tracked by the parent. `git diff --submodule` shows submodule changes.
- **Branch tracking**: `git submodule update --remote` updates to the remote branch tip, but the parent records a specific commit. This can be confusing.
- **Nested submodules**: Deeply nested submodules require `--recursive` for all operations.
- **Clone overhead**: `--recurse-submodules` clones all submodules, which can be slow for large projects.

## Subtree Merge

An alternative to submodules that merges another repository's history into the current one:

```bash
git remote add lib https://github.com/user/lib.git
git fetch lib
git merge -s ours --no-commit --allow-unrelated-histories lib/main
git read-tree --prefix=vendor/lib/ -u lib/main
git commit -m "Add lib as subtree"
```

### Updating a Subtree

```bash
git fetch lib
git merge -s subtree lib/main -m "Update lib subtree"
```

### Subtree vs Submodule

|| Aspect | Submodule | Subtree |
||--------|-----------|---------|
|| Repository | Separate repo | Merged into parent |
|| History | Separate | Shared with parent |
|| Clone | Requires `--recurse-submodules` | No special flags |
|| Update | `git submodule update` | `git merge -s subtree` |
|| Size | Small (reference only) | Large (full history) |
|| Bidirectional | Can push back to source | Difficult |
|| Complexity | High (many pitfalls) | Lower |

## git-subtree (Extension)

A contributed command that simplifies subtree operations:

```bash
git subtree add --prefix=vendor/lib lib main --squash
git subtree pull --prefix=vendor/lib lib main --squash
git subtree push --prefix=vendor/lib lib main    # Push back to source
git subtree merge --prefix=vendor/lib lib/main
git subtree split --prefix=vendor/lib -b lib-history  # Extract subtree history
```

`--squash` merges the subtree as a single commit without importing the full history.

## Sparse-Checkout

Limit which files appear in the working tree:

```bash
git sparse-checkout init --cone     # Enable cone mode
git sparse-checkout set src/ docs/  # Only these directories
git sparse-checkout add tests/      # Add more directories
git sparse-checkout list            # Show current patterns
git sparse-checkout disable         # Restore full checkout
git sparse-checkout reapply         # Reapply after changes
```

### Cone Mode

Fast, directory-level patterns only:

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

### Sparse-Checkout Internals

The index still contains **all** entries, but files outside the sparse-checkout cone have the `SKIP_WORKTREE` flag set. The sparse-checkout patterns are stored in `.git/info/sparse-checkout`.

## Partial Clone

Download only some objects initially, fetching the rest on demand:

```bash
git clone --filter=blob:none <url>       # No blobs (blobless clone)
git clone --filter=tree:0 <url>          # No trees or blobs (treeless clone)
git clone --filter=blob:limit=1m <url>   # Only small blobs
git clone --filter=depth:1 <url>        # Shallow + filter
```

### Filter Types

|| Filter | What's Excluded | Use Case |
||--------|-----------------|----------|
|| `blob:none` | All blobs | CI builds, code review |
|| `tree:0` | All trees and blobs | Minimal initial clone |
|| `blob:limit=<size>` | Blobs larger than size | Mixed small/large files |
|| `depth:1` | History beyond depth | Shallow clone |

### Promisor Remotes

In a partial clone, the remote is marked as a **promisor remote**. When Git needs a missing object, it fetches it from the promisor:

```bash
git config remote.origin.promisor true
git config remote.origin.partialclonefilter blob:none
```

### Lazy Fetch

When a command needs a missing object (e.g., `git checkout` needs a blob), Git automatically fetches it:

```
fatal: cannot checkout 'README.md' — object is missing
```

This is handled transparently if the promisor remote is configured.

## Shallow Clones

Truncate history at a certain depth:

```bash
git clone --depth 1 <url>                    # Only latest commit
git clone --depth 1 --single-branch <url>    # Only one branch
git clone --shallow-submodules <url>         # Shallow submodules too
git fetch --unshallow                         # Convert to full clone
git fetch --depth=10                          # Deepen to 10 commits
git fetch --unshallow                         # Remove depth limit
```

### Shallow File

`.git/shallow` lists the boundary commit OIDs. Commits below the boundary are not present.

### Limitations

- `git log` cannot show history beyond the shallow point
- `git blame` may fail for lines changed before the shallow point
- `git push` from a shallow repo may be rejected
- Some merge operations require full history

## Edge Cases and Known Issues

- **Worktree branch conflicts**: You cannot check out the same branch in two worktrees. Use different branches or detached HEAD.
- **Submodule URL changes**: After changing a submodule URL in `.gitmodules`, run `git submodule sync` to update the local config.
- **Sparse-checkout and add**: `git add` on a file outside the sparse-checkout cone adds it to the index but doesn't check it out to the working tree.
- **Partial clone and offline**: Missing objects cannot be fetched offline. Partial clones require network access for some operations.
- **Submodule recursive operations**: Many Git commands need `--recurse-submodules` to affect submodules. Without it, submodules are treated as opaque directories.
- **Worktree pruning**: If a worktree directory is deleted manually, `git worktree prune` cleans up the stale admin data.

## References

- <https://git-scm.com/docs/git-worktree>
- <https://git-scm.com/docs/git-submodule>
- <https://git-scm.com/book/en/v2/Git-Tools-Submodules>

## Related Skills

- For core concepts, see [concepts.md](concepts.md)
- For reference types (gitlink mode 160000), see [refs.md](refs.md)
- For index internals (SKIP_WORKTREE flag), see [index.md](index.md)
- For remotes (promisor remotes), see [remotes.md](remotes.md)
