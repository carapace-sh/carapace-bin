# Git CLI Reference

Command-line reference for Git's porcelain commands — the high-level user-facing interface. For plumbing commands, see [internals.md](internals.md).

> **Source of truth**: `git --help`, <https://git-scm.com/docs>. For **concepts**, see [concepts.md](concepts.md). For **branching and merging**, see [branching-merging.md](branching-merging.md). For **rebase**, see [rebase.md](rebase.md). For **remotes**, see [remotes.md](remotes.md).

## Setup and Init

|| Command | Description |
||---------|-------------|
|| `git init [--initial-branch=<name>]` | Create new repository |
|| `git clone <url> [<dir>]` | Clone remote repository |
|| `git clone --depth 1 <url>` | Shallow clone (no history) |
|| `git clone --filter=blob:none <url>` | Partial clone (blobless) |
|| `git clone --recurse-submodules <url>` | Clone with submodules |
|| `git clone --bare <url>` | Clone as bare repository |
|| `git clone --mirror <url>` | Mirror clone (bare + all refs) |

## Snapshot Operations

|| Command | Description |
||---------|-------------|
|| `git add <path>` | Stage file(s) |
|| `git add -p` | Interactive patch-mode staging |
|| `git add -A` / `git add --all` | Stage all changes |
|| `git add -u` | Stage modified/deleted only |
|| `git status` | Show working tree status |
|| `git status --short` / `git status -s` | Short format |
|| `git status --branch` / `git status -b` | Show branch info |
|| `git commit [-m <msg>]` | Create commit |
|| `git commit -a` | Stage tracked files and commit |
|| `git commit --amend` | Amend the last commit |
|| `git commit --allow-empty` | Create empty commit |
|| `git commit --fixup=<commit>` | Create fixup commit for rebase --autosquash |
|| `git commit --squash=<commit>` | Create squash commit for rebase --autosquash |
|| `git reset [<commit>]` | Reset HEAD (mixed: reset index + working tree) |
|| `git reset --soft [<commit>]` | Reset HEAD only (keep index + working tree) |
|| `git reset --mixed [<commit>]` | Reset HEAD + index (keep working tree) |
|| `git reset --hard [<commit>]` | Reset HEAD + index + working tree |
|| `git reset [<commit>] -- <path>` | Unstage specific file(s) |
|| `git restore <path>` | Restore working tree file (from index) |
|| `git restore --staged <path>` | Unstage file (from HEAD) |
|| `git restore -s <tree> <path>` | Restore from specific tree-ish |
|| `git rm <path>` | Remove from index + working tree |
|| `git rm --cached <path>` | Remove from index only |
|| `git mv <old> <new>` | Rename file (stage the rename) |
|| `git clean -fd` | Remove untracked files and directories |
|| `git clean -fdx` | Also remove ignored files |
|| `git clean -n` | Dry run |

## Branching and Switching

|| Command | Description |
||---------|-------------|
|| `git branch` | List local branches |
|| `git branch -a` | List all branches (local + remote) |
|| `git branch -v` | List with last commit |
|| `git branch <name>` | Create branch at HEAD |
|| `git branch <name> <start>` | Create branch at specific commit |
|| `git branch -d <name>` | Delete merged branch |
|| `git branch -D <name>` | Force delete branch |
|| `git branch -m <old> <new>` | Rename branch |
|| `git branch --set-upstream-to=<remote>/<branch> <branch>` | Set tracking |
|| `git switch <branch>` | Switch to branch |
|| `git switch -c <branch>` | Create and switch to branch |
|| `git switch -` | Switch to previous branch |
|| `git switch --detach <commit>` | Detach HEAD at commit |
|| `git checkout <branch>` | Switch to branch (legacy) |
|| `git checkout -b <branch>` | Create and switch (legacy) |
|| `git checkout --detach <commit>` | Detach HEAD (legacy) |

## Inspection and History

|| Command | Description |
||---------|-------------|
|| `git log` | Show commit history |
|| `git log --oneline` | One line per commit |
|| `git log --graph` | ASCII graph of branch structure |
|| `git log --all` | All branches |
|| `git log --decorate` | Show ref names |
|| `git log -p` | Show patches |
|| `git log --stat` | Show file stats |
|| `git log --author=<pattern>` | Filter by author |
|| `git log --since=<date>` | Filter by date |
|| `git log -n <count>` | Limit number of commits |
|| `git log --first-parent` | Follow only first parent (mainline) |
|| `git log --follow <path>` | Follow renames across history |
|| `git log -L :<func>:<file>` | Show history of function |
|| `git log --format=<format>` | Custom format string |
|| `git show <commit>` | Show commit details |
|| `git show <commit>:<path>` | Show file at commit |
|| `git show <commit>^:<path>` | Show file at parent commit |
|| `git diff` | Index vs working tree |
|| `git diff --cached` / `git diff --staged` | HEAD vs index |
|| `git diff HEAD` | HEAD vs working tree |
|| `git diff <a> <b>` | Between two commits/trees |
|| `git diff <a>...<b>` | Changes on b's side since merge base |
|| `git diff <a>..<b>` | Changes between a and b |
|| `git diff --stat` | Summary stats |
|| `git diff --word-diff` | Word-level diff |
|| `git shortlog` | Summary by author |
|| `git describe` | Describe commit using tags |
|| `git describe --tags` | Use lightweight tags too |
|| `git blame <file>` | Show last change per line |
|| `git blame -w` | Ignore whitespace |
|| `git blame -C` | Detect file copies |
|| `git blame -M` | Detect file moves |
|| `git grep <pattern>` | Search in tracked files |
|| `git grep -n <pattern>` | Show line numbers |
|| `git rev-list <range>` | List commit OIDs |
|| `git rev-parse <ref>` | Resolve ref to OID |

## Merging and Rebasing

|| Command | Description |
||---------|-------------|
|| `git merge <branch>` | Merge branch into current |
|| `git merge --no-ff <branch>` | Create merge commit even if fast-forward possible |
|| `git merge --ff-only <branch>` | Only merge if fast-forward possible |
|| `git merge --squash <branch>` | Squash branch into working tree (no commit) |
|| `git merge --abort` | Abort in-progress merge |
|| `git merge --continue` | Continue after conflict resolution |
|| `git rebase <upstream>` | Rebase current branch onto upstream |
|| `git rebase -i <upstream>` | Interactive rebase |
|| `git rebase --onto <new-base> <upstream> <branch>` | Rebase onto different base |
|| `git rebase --abort` | Abort rebase |
|| `git rebase --continue` | Continue rebase |
|| `git rebase --skip` | Skip current patch |
|| `git cherry-pick <commit>` | Apply commit on current branch |
|| `git cherry-pick --continue` | Continue after conflict |
|| `git cherry-pick --abort` | Abort cherry-pick |
|| `git revert <commit>` | Create new commit that undoes <commit> |
|| `git revert -n <commit>` | Revert but don't commit |
|| `git rerere` | Reuse recorded conflict resolutions |

See [branching-merging.md](branching-merging.md) and [rebase.md](rebase.md) for details.

## Remote Operations

|| Command | Description |
||---------|-------------|
|| `git remote` | List remotes |
|| `git remote -v` | List with URLs |
|| `git remote add <name> <url>` | Add remote |
|| `git remote remove <name>` | Remove remote |
|| `git remote set-url <name> <url>` | Change URL |
|| `git remote set-head <name> --auto` | Update remote HEAD ref |
|| `git fetch <remote>` | Download objects and refs |
|| `git fetch --prune` | Fetch and delete stale remote-tracking refs |
|| `git fetch --all` | Fetch from all remotes |
|| `git pull` | Fetch + merge (or rebase per config) |
|| `git pull --rebase` | Fetch + rebase |
|| `git pull --ff-only` | Only pull if fast-forward |
|| `git push <remote> <branch>` | Push branch |
|| `git push -u <remote> <branch>` | Push and set upstream |
|| `git push --force` | Force push (overwrite history) |
|| `git push --force-with-lease` | Force push with safety check |
|| `git push --all` | Push all branches |
|| `git push --tags` | Push all tags |
|| `git push --follow-tags` | Push annotated tags reachable from pushed refs |
|| `git push --dry-run` | Simulate push |
|| `git ls-remote <remote>` | List remote refs without fetching |

See [remotes.md](remotes.md) for protocol details.

## Stash

|| Command | Description |
||---------|-------------|
|| `git stash` | Stash working tree changes |
|| `git stash -u` | Include untracked files |
|| `git stash --all` | Include untracked + ignored |
|| `git stash -m <msg>` | Stash with message |
|| `git stash list` | List stashes |
|| `git stash pop` | Apply and drop latest stash |
|| `git stash pop stash@{2}` | Apply and drop specific stash |
|| `git stash apply` | Apply without dropping |
|| `git stash drop stash@{0}` | Drop stash |
|| `git stash clear` | Drop all stashes |
|| `git stash show` | Show stash summary |
|| `git stash show -p` | Show stash diff |
|| `git stash branch <name>` | Create branch from stash |

### Stash Internals

Stashes are stored as special merge commits in `refs/stash`:

```
stash commit
├── parent 1: HEAD at time of stash (index state)
├── parent 2: index state (staged changes)
└── parent 3: untracked files (only with -u)
```

## Bisect

Binary search to find the commit that introduced a bug:

```bash
git bisect start
git bisect bad                  # Current commit is bad
git bisect good <commit>        # Known good commit
# Git checks out middle commit; test it
git bisect good                 # Mark as good
git bisect bad                  # Mark as bad
# Repeat until found
git bisect reset                # Return to original branch
```

Automated bisect:

```bash
git bisect run <test-command>   # Auto-bisect using test script
```

The test command should exit 0 for good, 1 for bad, 125 for skip.

## Tagging

|| Command | Description |
||---------|-------------|
|| `git tag <name>` | Create lightweight tag |
|| `git tag -a <name> -m <msg>` | Create annotated tag |
|| `git tag -s <name> -m <msg>` | Create signed tag |
|| `git tag -v <name>` | Verify tag signature |
|| `git tag -d <name>` | Delete local tag |
|| `git push origin :refs/tags/<name>` | Delete remote tag |
|| `git tag -l <pattern>` | List tags matching pattern |
|| `git tag -n` | Show tag messages |

See [refs.md](refs.md) for tag internals.

## Notes

```bash
git notes add -m "note" <commit>    # Add note to commit
git notes show <commit>             # Show note
git notes list                      # List all notes
git notes remove <commit>           # Remove note
git notes edit <commit>             # Edit note
```

Notes are stored in `refs/notes/commits` as a tree mapping commit OIDs to note blobs. They are **not** pushed by default.

## Range-Diff

Compare two commit ranges (e.g., before and after rebase):

```bash
git range-diff <old-range> <new-range>
git range-diff main@{1}...main@{0}  # Compare before/after force push
```

## Maintenance

|| Command | Description |
||---------|-------------|
|| `git gc` | Garbage collect + repack |
|| `git gc --aggressive` | Slower, better compression |
|| `git gc --auto` | Auto-gc (run only if needed) |
|| `git maintenance start` | Start scheduled maintenance (Git 2.30+) |
|| `git maintenance run` | Run maintenance task |
|| `git prune` | Remove unreachable loose objects |
|| `git fsck` | Verify object database integrity |
|| `git repack -a -d` | Repack all objects |
|| `git count-objects -v` | Show object statistics |

## Global Flags

|| Flag | Description |
||------|-------------|
|| `-C <path>` | Run as if started in <path> |
|| `-c <key>=<value>` | Override config for this command |
|| `--git-dir=<path>` | Set .git directory path |
|| `--work-tree=<path>` | Set working tree path |
|| `--no-replace-objects` | Disable replace objects |
|| `--literal-pathspecs` | Treat pathspecs literally |
|| `--glob-pathspecs` | Treat pathspecs as globs |

## Edge Cases and Known Issues

- **`git commit -a` does not add new files**: It only stages modifications and deletions to already-tracked files.
- **`git checkout` is overloaded**: Git 2.23+ splits it into `git switch` (branches) and `git restore` (files). The old form still works.
- **`git pull` is a compound operation**: It does `git fetch` + `git merge` (or `git rebase` with `--rebase`). The merge/rebase choice depends on config and can cause unexpected merge commits.
- **`git reset --hard` is destructive**: It discards both staged and unstaged changes. The reflog preserves the old HEAD for recovery.
- **`git clean` does not affect ignored files by default**: Use `-x` to also remove ignored files.

## References

- <https://git-scm.com/docs>
- <https://git-scm.com/book/en/v2/Git-Basics-Recording-Changes-to-the-Repository>

## Related Skills

- For branching and merging details, see [branching-merging.md](branching-merging.md)
- For rebase and history rewriting, see [rebase.md](rebase.md)
- For remotes and protocols, see [remotes.md](remotes.md)
- For diff and conflict details, see [diff-patch.md](diff-patch.md)
- For plumbing commands, see [internals.md](internals.md)
