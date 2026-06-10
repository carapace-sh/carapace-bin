# Git Core Concepts

The foundational mental model for understanding Git — the three states, the content-addressable object store, the DAG, and the relationship between working tree, index, and repository.

> **Source of truth**: <https://git-scm.com/book/en/v2/Git-Basics-Getting-a-Git-Repository>, <https://git-scm.com/docs/git>. For **object format details**, see [objects.md](objects.md). For **reference types**, see [refs.md](refs.md). For **index internals**, see [index.md](index.md).

## The Three States

Git tracks three logical areas for every file:

|| State | Location | Description |
||-------|----------|-------------|
|| **Committed** | `.git/objects/` | Safely stored in the local object database |
|| **Modified** | Working tree | Changed but not yet staged |
|| **Staged** | `.git/index` | Marked to go into the next commit |

This gives rise to the three "areas" developers work with:

```
Working Tree          Index (Staging)        Repository
(files on disk)       (.git/index)           (.git/objects/)
     │                     │                      │
     ├── git add ─────────►│                      │
     │                     ├── git commit ───────►│
     │◄── git checkout ────┤                      │
     │◄──────── git reset ─┤◄── git reset --soft ─┤
```

### Working Tree

The actual files on disk. This is the one area that is **not** a Git-specific data structure — it's the directory the user sees and edits. The working tree is checked out from the index (or directly from a commit for `git checkout <tree-ish> -- <path>`).

### Index (Staging Area)

A flat list of file entries mapping path → object ID + stat info. The index is the **proposed next commit**. See [index.md](index.md) for format details.

### Repository

The `.git/` directory containing the object database and refs. A **bare repository** has no working tree — only the `.git/` contents at the top level (used on servers).

## Content-Addressable Storage

Git is fundamentally a **content-addressable filesystem**. Every piece of data is identified by its SHA-1 hash (or SHA-256 in experimental mode):

```
hash = SHA-1(type + " " + size + "\0" + content)
```

Key properties:

- **Same content → same hash**: Identical files stored once, regardless of path
- **Immutable**: Objects never change once created (they are created by hashing content)
- **Integrity**: Any corruption is detectable — `git fsck` verifies all object hashes
- **Deduplication**: Files with identical content share the same blob object

See [objects.md](objects.md) for the four object types and their format.

## The Directed Acyclic Graph (DAG)

Git history is a **directed acyclic graph** of commit objects:

```
    C1 ← C2 ← C3 ← C5  (main)
                ↘
                 C4 ← C6  (feature)
```

- Each commit points to **one or more parent commits** (merge commits have 2+ parents)
- Edges point **backward** in time (child → parent)
- The **root commit** has no parents
- There is no inherent "main branch" — branches are just movable pointers

### Commit Graph Properties

|| Property | Description |
||----------|-------------|
|| **Acyclic** | No commit can be its own ancestor (enforced by hash uniqueness) |
|| **Directed** | Parent pointers are one-way (child → parent) |
|| **Reachability** | From any ref, all ancestors are reachable; unreachable objects are garbage |
|| **Merge commits** | Commits with ≥2 parents represent merged histories |

## HEAD — The Current Commit

`HEAD` is a symbolic reference pointing to the current commit the working tree is based on:

|| State | HEAD points to | Behavior |
||-------|---------------|----------|
|| **Attached** | A branch name (e.g. `refs/heads/main`) | Moving HEAD moves the branch |
|| **Detached** | A specific commit SHA | HEAD does not move any branch |

### Detached HEAD

When HEAD points directly to a commit (not a branch), you are in "detached HEAD" state. This happens when:

- Checking out a commit hash: `git checkout abc1234`
- Checking out a tag: `git checkout v1.0`
- Checking out a remote branch without creating a local one
- During `git rebase` or `git bisect` operations

New commits made in detached HEAD are **orphaned** when you switch away — they become unreachable unless you create a branch or the reflog preserves them. See [refs.md](refs.md) for reflog recovery.

### HEAD File

```
# Attached (symbolic ref)
.git/HEAD → ref: refs/heads/main

# Detached
.git/HEAD → abc1234def567890...
```

`git symbolic-ref HEAD` reads/writes the symbolic form. `git update-ref` can set it directly.

## Plumbing vs Porcelain

Git commands are divided into two categories:

|| Category | Purpose | Examples |
||----------|---------|---------|
|| **Porcelain** | User-facing, high-level | `git add`, `git commit`, `git log`, `git merge` |
|| **Plumbing** | Low-level, scriptable | `git hash-object`, `git cat-file`, `git update-index` |

Plumbing commands are building blocks that porcelain commands are composed from. They are stable for scripting and expose the raw object model. See [internals.md](internals.md) for the full plumbing reference.

## The .git Directory

A standard `.git/` directory layout:

```
.git/
├── HEAD              # Current branch or commit
├── config            # Repository-local config
├── description       # For gitweb (rarely used)
├── hooks/            # Hook scripts
├── index             # The staging area
├── info/
│   ├── exclude       # Local exclude patterns (like .gitignore)
│   └── refs          # Dumb HTTP protocol info
├── objects/
│   ├── pack/         # Packfiles and indices
│   └── ??/          # Loose objects (first 2 hex chars = dir)
├── refs/
│   ├── heads/        # Branches
│   ├── tags/         # Tags
│   └── remotes/      # Remote-tracking refs
└── logs/             # Reflogs
    └── refs/
        └── heads/    # Branch reflogs
```

### Gitfile (worktrees)

When using `git worktree` or `git clone --separate-git-dir`, the `.git` entry can be a **file** instead of a directory:

```
gitdir: /path/to/actual/.git
```

This redirects Git to the real object database. See [worktree.md](worktree.md).

## Repository Initialization

### `git init`

Creates a new `.git/` directory with:
- An empty object database
- HEAD pointing to `refs/heads/main` (or `master` per `init.defaultBranch`)
- No commits yet (the branch ref doesn't exist until the first commit)

### `git clone`

Copies an existing repository:
1. Creates `.git/` directory
2. Sets up remote `origin` with the source URL
3. Fetches all objects and refs
4. Checks out the default branch (HEAD → remote HEAD)

### Bare Repositories

A bare repository (`git init --bare`) has **no working tree** — the `.git/` contents are at the top level. Used for:
- Server-side repositories (push targets)
- Mirror clones (`git clone --mirror`)
- CI/CD reference repositories

## Hash Algorithms

|| Algorithm | Status | Hash length | Example |
||-----------|--------|-------------|---------|
|| **SHA-1** | Default | 40 hex chars | `da39a3ee5e6b4b0d3255bfef95601890afd80709` |
|| **SHA-256** | Experimental | 64 hex chars | Enabled via `init.defaultHash` |

SHA-1 collision attacks exist (SHAttered, 2017) but are not practical against Git's hash format because the collision must match the `type size\0content` prefix. SHA-256 support is being rolled out for future-proofing.

## Edge Cases and Known Issues

- **Empty repositories have no commits**: `HEAD` points to a branch that doesn't exist yet. `git log` fails until the first commit.
- **Orphan branches**: `git checkout --orphan <branch>` creates a branch with no history — the first commit has no parent.
- **`.git` can be a file**: In worktrees and separate-git-dir setups, `.git` is a pointer file, not a directory.
- **Case sensitivity**: Git stores paths as-is but the index and tree objects are case-sensitive. On case-insensitive filesystems (macOS, Windows), two files differing only in case cause conflicts.
- **Symlinks**: Git stores symlinks as blob objects with the link target as content. The `core.symlinks` config can disable them (Windows).
- **Filemode changes**: `git diff` shows mode changes by default. `core.fileMode = false` ignores them (useful on FAT/exFAT).

## References

- <https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain>
- <https://git-scm.com/docs/gitrepository-layout>

## Related Skills

- For object format and packfile details, see [objects.md](objects.md)
- For reference types and reflogs, see [refs.md](refs.md)
- For index format and staging, see [index.md](index.md)
- For plumbing commands, see [internals.md](internals.md)
