# Git References

The reference system — branches, tags, HEAD, reflogs, remote-tracking branches, refspecs, and the packed-refs file.

> **Source of truth**: <https://git-scm.com/book/en/v2/Git-Internals-Git-References>, <https://git-scm.com/docs/gitrevisions>. For **object types**, see [objects.md](objects.md). For **core concepts**, see [concepts.md](concepts.md). For **remotes and refspecs**, see [remotes.md](remotes.md).

## Reference Types

A reference (ref) is a named pointer to a Git object, typically a commit:

|| Type | Location | Points to | Mutable |
||------|----------|-----------|---------|
|| **Branch** | `refs/heads/<name>` | Commit | Yes (moves on commit) |
|| **Tag (lightweight)** | `refs/tags/<name>` | Commit | No (fixed) |
|| **Tag (annotated)** | `refs/tags/<name>` | Tag object | No (fixed) |
|| **Remote-tracking** | `refs/remotes/<remote>/<branch>` | Commit | Updated by fetch |
|| **HEAD** | `.git/HEAD` | Branch or commit | Yes |
|| **Notes** | `refs/notes/<name>` | Tree (note mapping) | Yes |
|| **Replace** | `refs/replace/<OID>` | Commit (replacement) | Yes |
|| **Stash** | `refs/stash` | Commit (stash entry) | Yes |

## Branches

A branch is a movable pointer to a commit. When you commit while on a branch, the branch advances to the new commit.

### Branch Storage

```
.git/refs/heads/main       → abc1234def...
.git/refs/heads/feature    → 567890abc...
```

Each file contains a single line: the 40-character hex OID of the commit it points to.

### Branch Operations

```bash
git branch                  # List local branches
git branch <name>           # Create branch at HEAD
git branch <name> <commit>  # Create branch at specific commit
git branch -d <name>        # Delete branch (safe — checks merge)
git branch -D <name>        # Force delete
git branch -m <old> <new>   # Rename branch
git branch -f <name> <commit>  # Force move branch to commit
```

### Default Branch

The initial branch name is controlled by `init.defaultBranch` (default `main` since Git 2.28):

```bash
git config --global init.defaultBranch main
```

## Tags

Tags are fixed pointers, typically used for releases.

### Lightweight Tags

A simple pointer to a commit — just a ref file:

```bash
git tag v1.0              # Create lightweight tag
```

### Annotated Tags

A full object with message, tagger, and optional signature:

```bash
git tag -a v1.0 -m "Release 1.0"    # Create annotated tag
git tag -s v1.0 -m "Release 1.0"    # Create GPG-signed tag
```

See [objects.md](objects.md) for the tag object format.

### Tag Operations

```bash
git tag                    # List tags
git tag -l "v1.*"          # List tags matching pattern
git tag -d v1.0            # Delete tag
git tag -n                 # Show tag messages
git push origin v1.0       # Push single tag
git push origin --tags     # Push all tags
git push origin --follow-tags  # Push annotated tags reachable from pushed refs
```

## HEAD

HEAD is the symbolic reference indicating the current commit. See [concepts.md](concepts.md) for the detached HEAD discussion.

### HEAD Reflog

HEAD has its own reflog (`.git/logs/HEAD`) recording every checkout and commit:

```bash
git reflog show HEAD       # Show HEAD reflog
git reflog                 # Shorthand for HEAD reflog
```

## Reflogs

A **reflog** records the history of where a reference has pointed. Every time a branch is updated (commit, reset, rebase, etc.), the previous value is recorded.

### Reflog Storage

```
.git/logs/HEAD                          # HEAD reflog
.git/logs/refs/heads/main              # Branch reflog
.git/logs/refs/stash                    # Stash reflog
```

### Reflog Format

Each line:

```
<old-OID> <new-OID> <name> <email> <timestamp> <tz> <action>\t<message>
```

Example:

```
abc1234 def5678 Alice alice@example.com 1700000000 +0000 commit: fix bug
```

### Reflog Commands

```bash
git reflog show main       # Show reflog for branch
git reflog show HEAD       # Show HEAD reflog (default)
git reflog expire --expire=90.days.ago --all  # Prune old entries
git reflog delete main@{2}  # Delete specific entry
```

### Reflog References

Reflog entries can be referenced by position:

```
main@{0}       # Current value of main
main@{1}       # Previous value of main
main@{2}       # Two moves ago
main@{yesterday}  # Value of main at yesterday
main@{2.hours.ago} # Value of main 2 hours ago
HEAD@{5}       # 5th previous HEAD position
```

### Reflog Expiry

|| Config | Default | Description |
||--------|---------|-------------|
|| `gc.reflogExpire` | 90 days | When to expire reflog entries for unreachable commits |
|| `gc.reflogExpireUnreachable` | 30 days | When to expire entries for unreachable commits |
|| `gc.reflogExpireIn` | Never | For `--expire=<date>` |

## Remote-Tracking Branches

Remote-tracking branches are local mirrors of remote branches, updated by `git fetch`:

```
refs/remotes/origin/main     → abc1234...
refs/remotes/origin/feature  → def5678...
```

### Tracking Relationship

A local branch can **track** a remote-tracking branch:

```bash
git branch --track feature origin/feature    # Set up tracking
git branch --set-upstream-to=origin/feature feature  # Change upstream
git branch -vv                                # Show tracking info
```

The tracking relationship configures:
- Default push/pull target (`git pull` without arguments)
- `branch.<name>.remote` and `branch.<name>.merge` config entries
- Status display (`git status` shows ahead/behind counts)

### Remote-Tracking Refs Are Read-Only

You cannot commit onto a remote-tracking branch. They are updated only by:

- `git fetch` (or `git remote update`)
- `git pull` (which does a fetch)
- `git push` (updates the remote, then fetch updates the tracking ref)

## Packed Refs

When a repository has many refs, Git packs them into a single file for efficiency:

```
.git/packed-refs
```

Format:

```
# pack-refs with: peeled fully-peeled sorted 
abc1234def567890... refs/heads/main
^def567890abc1234...                   # Peeled tag (commit the tag points to)
567890abcdef1234... refs/tags/v1.0
```

- Lines starting with `^` are "peeled" — the commit an annotated tag points to
- When a packed ref is updated, Git writes the new value to the loose file in `refs/`, which **overrides** the packed version
- `git pack-refs --all` packs all loose refs

## Refspecs

A **refspec** maps remote refs to local refs during fetch and push:

```
[+]<src>:<dst>
```

- `+` prefix enables force update
- `<src>` is the source ref (remote for fetch, local for push)
- `<dst>` is the destination ref

### Default Refspecs

Configured in `.git/config` when a remote is added:

```
[remote "origin"]
    url = https://github.com/user/repo.git
    fetch = +refs/heads/*:refs/remotes/origin/*
```

This means: fetch all branches from the remote and store them as remote-tracking branches.

### Custom Refspecs

```bash
git fetch origin refs/heads/feature:refs/heads/feature  # Fetch directly into local branch
git fetch origin +refs/heads/main:refs/remotes/origin/main  # Force-update
git push origin refs/heads/main:refs/heads/main  # Push main to main
```

### Refspec Patterns

Wildcards in refspecs:

```
refs/heads/*:refs/remotes/origin/*    # All branches
refs/tags/*:refs/tags/*              # All tags
```

## Special Refs

|| Ref | Purpose |
||-----|---------|
|| `HEAD` | Current commit/branch |
|| `FETCH_HEAD` | Last fetched tip (`.git/FETCH_HEAD`) |
|| `MERGE_HEAD` | Commit being merged into HEAD (`.git/MERGE_HEAD`) |
|| `REBASE_HEAD` | Commit being rebased (during rebase) |
|| `CHERRY_PICK_HEAD` | Commit being cherry-picked |
|| `REVERT_HEAD` | Commit being reverted |
|| `ORIG_HEAD` | Previous HEAD before dangerous operation (merge, rebase) |

## Revision Syntax

Git supports many ways to refer to commits:

|| Syntax | Meaning |
||--------|---------|
|| `abc1234` | Full or abbreviated SHA |
|| `main` | Branch name |
|| `v1.0` | Tag name |
|| `main@{yesterday}` | Reflog — value of main at that time |
|| `main@{3}` | Reflog — 3rd previous value |
|| `main~3` | 3rd generation ancestor (first parent only) |
|| `main^2` | 2nd parent (for merge commits) |
|| `main^` / `main^1` | First parent |
|| `main~` / `main~1` | First parent (same as `^`) |
|| `main^^` / `main~2` | Grandparent (first parent of first parent) |
|| `HEAD^..HEAD` | Range (see below) |
|| `abc..def` | Commits reachable from def but not abc |
|| `abc...def` | Symmetric difference (reachable from either, not both) |
|| `abc^@` | All parents of abc |
|| `abc^!` | abc itself, excluding all its parents |

## Reference Update Commands

|| Command | Purpose |
||---------|---------|
|| `git update-ref <ref> <new-oid>` | Update a ref (safe — takes lock) |
|| `git update-ref -d <ref>` | Delete a ref |
|| `git symbolic-ref HEAD` | Read symbolic ref |
|| `git symbolic-ref HEAD refs/heads/main` | Set symbolic ref |
|| `git for-each-ref` | Iterate refs (scriptable) |
|| `git for-each-ref --format='%(refname) %(objectname)' refs/heads/` | Custom format |
|| `git show-ref` | List all refs |
|| `git pack-refs --all` | Pack loose refs |

## Edge Cases and Known Issues

- **Stale remote-tracking branches**: After a remote branch is deleted, the local tracking ref persists. Use `git remote prune origin` or `git fetch --prune` to clean up.
- **Reflog recovery**: Even after `git reset --hard` or a force push, the reflog preserves the old commit for `gc.reflogExpire` days. Use `git reflog` to find and recover.
- **Packed refs override**: A loose ref always takes precedence over a packed ref with the same name. This can cause confusion if the loose file is stale.
- **Tag deletion on remote**: `git push origin :refs/tags/v1.0` deletes a tag on the remote. The colon syntax pushes "nothing" to the destination.
- **Namespace conflicts**: A branch named `feature` and a tag named `feature` can coexist, but `git rev-parse feature` resolves in a defined priority order (see [concepts.md](concepts.md) symbol resolution).
- **Refname restrictions**: Refs cannot contain `..`, `~`, `^`, `:`, space, or control characters. They cannot end with `.lock` or `@{`.

## References

- <https://git-scm.com/docs/gitrevisions>
- <https://git-scm.com/docs/git-update-ref>
- <https://git-scm.com/docs/git-for-each-ref>

## Related Skills

- For object types that refs point to, see [objects.md](objects.md)
- For core concepts (HEAD, detached HEAD), see [concepts.md](concepts.md)
- For remotes and refspecs in practice, see [remotes.md](remotes.md)
- For reflog-based recovery, see [rebase.md](rebase.md)
