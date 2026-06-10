# Git Branching and Merging

How Git branches work, merge strategies, merge bases, conflict handling, and the mechanics of combining histories.

> **Source of truth**: <https://git-scm.com/book/en/v2/Git-Branching-Basic-Branching-and-Merging>, <https://git-scm.com/docs/git-merge>. For **reference types**, see [refs.md](refs.md). For **rebase**, see [rebase.md](rebase.md). For **conflict resolution**, see [diff-patch.md](diff-patch.md).

## Branch Fundamentals

A branch is a lightweight movable pointer to a commit. Creating a branch is O(1) — just writing a 41-byte file.

### Branch Creation

```bash
git branch feature          # Create at HEAD
git branch feature abc123   # Create at specific commit
git switch -c feature       # Create and switch
git checkout -b feature     # Legacy form
```

### Branch Movement

Branches move **automatically** when you commit while attached:

```
Before:  main → C3
After:   main → C4 → C3   (main advanced to new commit C4)
```

In detached HEAD, commits are made but no branch moves — the new commits are only reachable via the reflog.

## Merge Base

The **merge base** is the best common ancestor of two or more commits — the point where the histories diverged.

```
      C1 ← C2 ← C3 ← C5  (main)
           ↘
            C4 ← C6       (feature)

Merge base of main and feature = C2
```

### Finding the Merge Base

```bash
git merge-base main feature          # Single merge base
git merge-base --all main feature    # All merge bases (for criss-cross merges)
git merge-base --is-ancestor A B    # Exit 0 if A is ancestor of B
```

### Criss-Cross Merges

When two branches merge each other in both directions, there can be **multiple merge bases**:

```
    A ← M1 ← B        (M1 merges B into A's history)
    A ← M2 ← B        (M2 merges A into B's history)
```

In this case, `git merge-base --all` returns both M1 and B (or A and M2, depending on topology). The `ort` merge strategy picks one as the "best" merge base.

## Merge Strategies

Git supports several merge strategies:

| Strategy | Flag | Description |
|----------|------|-------------|
| **ort** (Ostensibly Recursive's Twin) | Default | Recursive three-way merge with rename detection |
| **recursive** | `-s recursive` | Legacy default (pre-Git 2.34) |
| **resolve** | `-s resolve` | Three-way merge, single merge base only |
| **octopus** | `-s octopus` | For merging more than two heads (fails on conflicts) |
| **ours** | `-s ours` | Auto-resolve all conflicts as "our" side |
| **subtree** | `-s subtree` | Like `ort` but adjusts tree paths for subtree merges |

### ort Strategy

The default since Git 2.34. Key features:

- Handles rename detection across all sides
- Supports criss-cross merges (multiple merge bases)
- Low-memory footprint (uses in-memory trees)
- Options: `--no-renames`, `--diff-algorithm`, `--find-renames`

### ours Strategy

**Not** "keep our changes" — it **discards all changes from the other side**:

```bash
git merge -s ours feature   # Result tree is identical to HEAD
```

Use case: recording a merge without incorporating changes (e.g., marking a branch as integrated).

### octopus Strategy

For merging 3+ branches at once. Fails if any conflicts exist — it cannot handle conflicts. Used by default when `git merge` is given multiple branch names.

## Fast-Forward Merges

When the current branch is a direct ancestor of the target, Git can **fast-forward** — just move the branch pointer:

```
Before:  main → C3,  feature → C4 → C3
After:   main → C4 → C3,  feature → C4 → C3
         (main moved to C4, no new commit created)
```

```bash
git merge feature           # Fast-forward if possible
git merge --ff-only feature # Fail if not fast-forward
git merge --no-ff feature   # Always create merge commit
```

### When to Use --no-ff

- Preserves the branch topology (shows that a feature was merged)
- Makes `git log --first-parent` useful (follows mainline)
- Allows reverting the entire feature with one revert

## Merge Commits

A merge commit has two or more parents:

```
    C3 ← C5 ← M    (M has parents C5 and C6)
         ↘   ↗
          C4 ← C6
```

### First-Parent Convention

The **first parent** is the branch you were on when you merged. The **second parent** is the branch you merged in. This convention enables:

```bash
git log --first-parent    # Follow only mainline
git log --merges          # Show only merge commits
git log --no-merges       # Show only non-merge commits
```

## Squash Merges

`git merge --squash` applies all changes from the source branch but does **not** create a merge commit:

```bash
git merge --squash feature
git commit -m "Feature: add X"
```

The result is a single commit with all the feature's changes, but **no merge topology** — the branch history is lost.

### When to Use Squash Merge

- Clean up a messy feature branch history
- Policy: only clean commits on main
- One logical change per merge

### When Not to Use

- Need to revert the entire feature later (no single merge commit to revert)
- Want to preserve authorship of individual commits
- Need bisectability of individual changes

## Merge Conflict Resolution

When two sides modify the same area, Git marks the conflict:

```
<<<<<<< HEAD
content from current branch
=======
content from merged branch
>>>>>>> feature
```

### Resolution Steps

1. Open the conflicted file
2. Edit to resolve the conflict (remove markers, choose content)
3. `git add <resolved-file>`
4. `git merge --continue` (or `git commit`)

### Conflict Markers

| Marker | Meaning |
|--------|---------|
| `<<<<<<<` | Start of conflict |
| `=======` | Separator between sides |
| `>>>>>>>` | End of conflict |

For merge commits with 3+ parents, markers may include `|||||| <base>` (with `merge.conflictstyle = diff3`):

```
<<<<<<< HEAD
ours
|||||| base
original
=======
theirs
>>>>>>> feature
```

### Merge Drivers

Custom merge drivers can be defined in `.gitattributes`:

```
*.json merge=json-driver
```

```bash
git config merge.json-driver.name "JSON merge driver"
git config merge.json-driver.driver "json-merge %O %A %B %L"
```

- `%O` — base (common ancestor)
- `%A` — ours (current branch, modified in-place)
- `%B` — theirs (merged branch)
- `%L` — conflict marker size

The driver should write the merge result to `%A` and exit 0 for success, non-zero for conflict.

### Built-in Merge Drivers

| Driver | Behavior |
|--------|----------|
| `text` | Default three-way merge with conflict markers |
| `binary` | Keep ours version (no attempt to merge) |
| `union` | Auto-resolve by concatenating both sides (no conflict markers) |

## Merge Abort and Undo

```bash
git merge --abort          # Abort in-progress merge, restore pre-merge state
git reset --hard HEAD      # Same effect after a failed merge
git revert -m 1 <merge>   # Revert a merge commit (specify parent to keep)
```

### Reverting a Merge Commit

A merge commit has multiple parents. `git revert` needs to know which parent to follow:

```bash
git revert -m 1 <merge-commit>   # -m 1 = keep first parent (mainline)
```

This creates a new commit that undoes the changes from the merged branch. **Caution**: this can make future merges of the same branch difficult, as Git thinks those changes are already in the history.

## Orphan Branches

```bash
git checkout --orphan <branch>   # Create branch with no history
git rm -rf .                     # Clear index (optional)
git commit --allow-empty -m "Initial"  # First commit has no parent
```

Use cases: `gh-pages` for GitHub Pages, separate documentation sites, completely independent histories.

## Edge Cases and Known Issues

- **Merge conflicts in generated files**: Files that are auto-generated (e.g., lock files, compiled output) should be handled with custom merge drivers or `.gitattributes` merge strategies.
- **Renaming conflicts**: If one side renames a file and the other modifies it, the `ort` strategy handles this, but older strategies may not.
- **Submodule conflicts**: Submodule conflicts show the two commit OIDs — you must choose which to check out.
- **Large merges**: Merging many branches at once (octopus) fails on any conflict. Merge pairwise instead.
- **Re-merging after revert**: After `git revert -m 1 <merge>`, re-merging the same branch may not bring back the changes. You need to revert the revert first.
- **Fast-forward and branch topology**: Fast-forward merges lose the information that a feature branch existed. Use `--no-ff` to preserve topology.

## References

- <https://git-scm.com/docs/git-merge>
- <https://git-scm.com/docs/git-merge-base>
- <https://git-scm.com/book/en/v2/Git-Branching-Basic-Branching-and-Merging>

## Related Skills

- For rebase (alternative to merge), see [rebase.md](rebase.md)
- For conflict markers and diff details, see [diff-patch.md](diff-patch.md)
- For branch internals, see [refs.md](refs.md)
- For custom merge drivers, see [diff-patch.md](diff-patch.md)
