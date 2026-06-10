# Git Rebase, Cherry-Pick, and History Rewriting

How Git rewrites history — interactive rebase, cherry-pick, commit amending, and the recovery mechanisms that make rewriting safe.

> **Source of truth**: <https://git-scm.com/docs/git-rebase>, <https://git-scm.com/docs/git-cherry-pick>. For **branching and merging**, see [branching-merging.md](branching-merging.md). For **reflog recovery**, see [refs.md](refs.md). For **conflict resolution**, see [diff-patch.md](diff-patch.md).

## Rebase Overview

Rebase replays commits from one branch onto another base:

```
Before:          After rebase:
  A←B←C (main)     A←B←C←D'←E' (main)
       ↘                ↑
        D←E (feature)   D←E (feature, now orphaned)
```

The original commits D and E still exist (in the reflog) but are no longer reachable from the feature branch.

### Basic Rebase

```bash
git rebase main              # Rebase current branch onto main
git rebase main feature      # Rebase feature onto main (checks out feature first)
git rebase --onto new A B    # Rebase commits in (A..B] onto new
```

### The --onto Flag

`--onto` enables precise control over which commits are rebased and where:

```bash
git rebase --onto main feature~3 feature
```

This takes the last 3 commits of `feature` and replays them onto `main`, leaving earlier feature commits behind.

```
Before:  main → C3,  feature → D → E → F → G → C3
After:   main → C3 → E' → F' → G',  feature → D → E → F → G → C3 (orphaned)
```

## Interactive Rebase

`git rebase -i` opens an editor with a "todo list" of commits to replay:

```bash
git rebase -i main
git rebase -i HEAD~5          # Rebase last 5 commits
git rebase -i --root           # Rebase from the root commit
```

### Todo List Format

```
pick abc1234 First commit
reword def5678 Second commit
edit 9012345 Third commit
squash 6789012 Fourth commit
fixup 3456789 Fifth commit
drop  0123456 Sixth commit
```

### Commands

|| Command | Short | Action |
||---------|-------|--------|
|| `pick` | `p` | Use commit as-is |
|| `reword` | `r` | Use commit but edit the message |
|| `edit` | `e` | Use commit but stop for amending |
|| `squash` | `s` | Meld into previous commit (combine messages) |
|| `fixup` | `f` | Meld into previous commit (discard this message) |
|| `fixup -C` | `F` | Meld into previous commit (use this message instead) |
|| `fixup -c` | `C` | Like fixup -C but edit message |
|| `drop` | `d` | Delete commit |
|| `break` | `b` | Stop here (continue with `git rebase --continue`) |
|| `exec` | `x` | Run shell command |
|| `label` | `l` | Label current HEAD (for merge rebase) |
|| `merge` | `m` | Create merge commit (for `--rebase-merges`) |
|| `reset` | `t` | Reset to label (for `--rebase-merges`) |

### Reordering

Move lines up or down to reorder commits. Commits are applied top-to-bottom.

### Splitting a Commit

1. Mark the commit as `edit`
2. When rebase stops: `git reset HEAD^`
3. Stage and commit changes in pieces
4. `git rebase --continue`

## Autosquash

Commits created with `--fixup` or `--squash` are automatically reordered:

```bash
git commit --fixup=abc1234     # Create fixup commit
git commit --squash=abc1234    # Create squash commit
git rebase -i --autosquash     # Auto-arrange fixup/squash after their targets
```

The todo list is automatically rearranged:

```
pick abc1234 Original commit
fixup def5678 fixup! Original commit
```

## Rebase Merges

`--rebase-merges` preserves merge commits during rebase:

```bash
git rebase -i --rebase-merges main
```

This generates `label`, `reset`, and `merge` commands in the todo list to recreate the merge topology.

## Update Refs

`--update-refs` automatically updates any branches that point to commits being rebased:

```bash
git rebase -i --update-refs main
```

Without this flag, other branches pointing to the original commits become orphaned. With it, they are moved to the corresponding new commits.

## Cherry-Pick

Cherry-pick applies a commit from one branch onto another:

```bash
git cherry-pick abc1234           # Apply single commit
git cherry-pick abc1234..def5678  # Apply range
git cherry-pick --continue        # Continue after conflict
git cherry-pick --abort           # Abort
git cherry-pick -n abc1234        # Apply but don't commit (no commit created)
git cherry-pick -x abc1234        # Add "(cherry picked from commit ...)" to message
git cherry-pick -s abc1234        # Add Signed-off-by
```

### Cherry-Pick Internals

Cherry-pick is essentially a three-way merge:
1. Base = parent of the cherry-picked commit
2. Ours = HEAD
3. Theirs = cherry-picked commit

The resulting commit has HEAD as its parent (not the original parent).

### Cherry-Pick vs Merge

Cherry-pick creates a **new commit** with a different OID. The original commit still exists. This means:
- `git log` shows the cherry-picked commit as a different commit
- `git cherry` can identify which commits have been cherry-picked
- Future merges may detect the changes as already applied (via patch-id)

## Commit Amending

```bash
git commit --amend                 # Amend last commit (message + content)
git commit --amend -m "New msg"    # Amend with new message
git commit --amend --no-edit       # Amend content only, keep message
git commit --amend --author="..."  # Change author
git commit --amend --date="..."    # Change author date
```

Amending creates a **new commit** with a different OID. The old commit becomes unreachable (but preserved in the reflog).

## Reset

`git reset` moves the current branch to a specified commit:

|| Mode | HEAD | Index | Working Tree |
||------|------|-------|-------------|
|| `--soft` | Moved | Unchanged | Unchanged |
|| `--mixed` (default) | Moved | Reset to HEAD | Unchanged |
|| `--hard` | Moved | Reset to HEAD | Reset to HEAD |
|| `--merge` | Moved | Reset to HEAD | Reset to HEAD (only for unmerged) |
|| `--keep` | Moved | Unchanged | Reset only if no local changes |

### Reset Patterns

```bash
git reset HEAD~1                  # Undo last commit, keep changes staged
git reset --soft HEAD~1           # Undo last commit, keep changes in index
git reset --hard HEAD~1           # Undo last commit, discard all changes
git reset HEAD <path>             # Unstage file
git reset --hard origin/main      # Reset to remote state
```

## Revert

`git revert` creates a new commit that undoes a previous commit:

```bash
git revert abc1234                # Revert single commit
git revert abc1234..def5678       # Revert range
git revert -n abc1234             # Revert but don't commit
git revert -m 1 <merge-commit>   # Revert merge (specify parent to keep)
```

Unlike reset, revert is **safe for shared history** — it creates a new commit rather than rewriting history.

## Recovering Lost Commits

### Using the Reflog

```bash
git reflog                        # Find the lost commit OID
git reset --hard <oid>            # Restore to that commit
git cherry-pick <oid>             # Apply just that commit
git branch recovered <oid>        # Create branch at the lost commit
```

### Using `git fsck`

```bash
git fsck --lost-found             # Find dangling commits
git fsck --unreachable            # List all unreachable objects
```

Dangling commits are written to `.git/lost-found/commit/`.

### Using ORIG_HEAD

After a merge or rebase, `ORIG_HEAD` points to the previous HEAD:

```bash
git reset --hard ORIG_HEAD        # Undo the merge/rebase
```

## History Rewriting at Scale

### filter-branch (Deprecated)

```bash
git filter-branch --tree-filter 'rm -f passwords.txt' HEAD
git filter-branch --subdirectory-filter src HEAD
git filter-branch --env-filter 'export GIT_AUTHOR_EMAIL="new@example.com"' HEAD
```

**Deprecated** due to performance and safety issues. Use `git filter-repo` instead.

### filter-repo

```bash
git filter-repo --path passwords.txt --invert-paths  # Remove file from history
git filter-repo --subdirectory-filter src             # Extract subdirectory
git filter-repo --email new@example.com               # Change author email
git filter-repo --replace-text expressions.txt        # Replace text in all blobs
```

`git filter-repo` is faster, safer, and handles all edge cases that `filter-branch` misses. It requires a fresh clone (or `--force`).

## The Golden Rule of Rebase

**Never rebase commits that have been pushed and shared with others.** Rebase rewrites history — if someone else has based work on the original commits, their history diverges.

Exception: When all collaborators agree to force-push (e.g., feature branches with a single owner, or `--force-with-lease` safety checks).

## Rebase vs Merge

|| Aspect | Rebase | Merge |
||--------|--------|-------|
|| History | Linear (no merge commits) | Preserves branch topology |
|| Commit OIDs | Changed (new commits) | Preserved |
|| Conflict resolution | One commit at a time | All at once |
|| Traceability | Loses branch context | Shows which branch was merged |
|| Revert | Must revert individual commits | Revert the merge commit |
|| Shared branches | Dangerous (rewrites history) | Safe (no rewriting) |

## Edge Cases and Known Issues

- **Rebase conflict cascade**: A conflict in an early commit may cause conflicts in subsequent commits that depend on the same code. Interactive rebase lets you resolve one at a time.
- **Empty commits after rebase**: If a commit becomes empty after rebase (changes already in the base), Git drops it by default. Use `--keep-empty` to preserve.
- **Rebase and submodules**: Submodule entries may conflict during rebase. Resolve by checking out the desired submodule commit.
- **Amend after push**: Amending a pushed commit requires force push. Use `--force-with-lease` for safety.
- **Rebase on dirty tree**: Rebase requires a clean working tree (or `--autostash` to auto-stash changes).
- **`git rebase --autostash`**: Automatically stashes dirty changes before rebase and pops them after. Enabled by default with `rebase.autoStash = true`.
- **Cherry-pick and patch-id**: `git cherry` compares patch-ids to determine if a commit has been cherry-picked. Whitespace changes can cause false negatives.

## References

- <https://git-scm.com/docs/git-rebase>
- <https://git-scm.com/docs/git-cherry-pick>
- <https://git-scm.com/book/en/v2/Git-Branching-Rebasing>

## Related Skills

- For merging (alternative to rebase), see [branching-merging.md](branching-merging.md)
- For conflict resolution details, see [diff-patch.md](diff-patch.md)
- For reflog-based recovery, see [refs.md](refs.md)
- For hooks triggered during rebase, see [hooks.md](hooks.md)
