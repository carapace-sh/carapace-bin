# Git Internals — Plumbing Commands

The low-level plumbing commands that expose Git's object model directly. These are the building blocks that porcelain commands are composed from.

> **Source of truth**: <https://git-scm.com/docs>, <https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain>. For **object types**, see [objects.md](objects.md). For **core concepts**, see [concepts.md](concepts.md). For **reference types**, see [refs.md](refs.md).

## .git Directory Layout

```
.git/
├── HEAD                    # Current branch or commit
├── config                  # Repository-local config
├── description             # Gitweb description (rarely used)
├── index                   # The staging area (binary)
├── info/
│   ├── exclude             # Local gitignore (not tracked)
│   ├── refs                # Dumb HTTP protocol info
│   └── grafts              # Graft points (deprecated, use replace)
├── hooks/                  # Hook scripts
├── logs/
│   ├── HEAD                # HEAD reflog
│   └── refs/
│       ├── heads/          # Branch reflogs
│       └── stash           # Stash reflog
├── objects/
│   ├── pack/               # Packfiles and indices
│   ├── info/
│   │   └── packs           # List of packfiles (dumb HTTP)
│   └── ??/                 # Loose objects (first 2 hex chars)
├── refs/
│   ├── heads/              # Branches
│   ├── tags/               # Tags
│   ├── remotes/            # Remote-tracking refs (legacy)
│   ├── replace/            # Replace objects
│   └── stash               # Stash ref
├── packed-refs             # Packed references file
├── shallow                 # Shallow clone boundary
├── MERGE_HEAD              # Commit being merged
├── MERGE_MSG               # Merge commit message
├── MERGE_MODE              # Merge options
├── CHERRY_PICK_HEAD        # Commit being cherry-picked
├── REVERT_HEAD             # Commit being reverted
├── REBASE_HEAD             # Commit being rebased
├── ORIG_HEAD              # Previous HEAD (before merge/rebase)
├── FETCH_HEAD              # Last fetched tips
└── modules/                # Submodule .git directories
```

## Object Manipulation

### git hash-object

Compute the object ID for content, optionally writing it to the object store:

```bash
git hash-object <file>              # Compute OID (don't store)
git hash-object -w <file>           # Compute OID and store as blob
git hash-object --stdin             # Read from stdin
git hash-object -t commit <file>    # Specify object type
```

### git cat-file

Inspect objects in the store:

```bash
git cat-file -t <OID>               # Object type
git cat-file -s <OID>               # Object size (bytes)
git cat-file -p <OID>               # Pretty-print (human-readable)
git cat-file blob <OID>             # Raw blob content
git cat-file tree <OID>            # Raw tree content (binary)
git cat-file commit <OID>          # Raw commit content
git cat-file tag <OID>             # Raw tag content
git cat-file -p <OID>^{tree}       # Tree of a commit
git cat-file -e <OID>              # Exit 0 if object exists
```

### git mktree

Build a tree object from stdin:

```bash
echo -e "100644 blob e69de29bb2d1d6434b8b29ae775ad8c2e48c5391\tREADME" | git mktree
```

Input format: `mode type OID\tname`

### git mktag

Create a tag object from stdin:

```bash
git mktag <<EOF
object abc1234...
type commit
tag v1.0
tagger Alice <alice@example.com> 1700000000 +0000

Release v1.0
EOF
```

### git commit-tree

Create a commit object from a tree:

```bash
git commit-tree <tree-oid> -p <parent-oid> -m "message"
echo "message" | git commit-tree <tree-oid> -p <parent-oid>
```

Low-level alternative to `git commit`. Used by `git rebase` internally.

## Index Manipulation

### git update-index

Modify the index directly:

```bash
git update-index --add <file>        # Add file to index
git update-index --remove <file>     # Remove from index
git update-index --force-remove <file>  # Force remove
git update-index --chmod=+x <file>  # Set executable bit
git update-index --skip-worktree <file>  # Set skip-worktree
git update-index --no-skip-worktree <file>  # Unset
git update-index --assume-unchanged <file>  # Set assume-unchanged
git update-index --refresh           # Refresh stat info
git update-index --again             # Update entries that differ from HEAD
```

### git read-tree

Read tree objects into the index:

```bash
git read-tree HEAD                  # Replace index with HEAD tree
git read-tree --reset HEAD          # Reset index + working tree
git read-tree -m HEAD MERGE_HEAD    # Merge trees into index (for merge)
git read-tree --prefix=src/ HEAD    # Read tree under prefix
```

### git write-tree

Write the index as a tree object:

```bash
git write-tree                      # Write index to object store, return OID
git write-tree --missing-ok         # Allow missing objects
```

### git ls-files

List index entries:

```bash
git ls-files                        # List tracked files
git ls-files -s                     # Show mode, OID, stage, path
git ls-files -u                     # Show unmerged entries
git ls-files -v                     # Show assume-unchanged/skip-worktree flags
git ls-files -d                     # Show deleted files
git ls-files -m                     # Show modified files
git ls-files -o                     # Show untracked files
git ls-files -o --exclude-standard  # Untracked, respect .gitignore
```

### git ls-tree

List tree object entries:

```bash
git ls-tree HEAD                    # Top-level entries
git ls-tree -r HEAD                 # Recursive
git ls-tree -r -t HEAD              # Recursive, show tree entries
git ls-tree -l HEAD                 # Long format (with size)
git ls-tree HEAD -- src/            # Only entries under src/
```

## Reference Manipulation

### git update-ref

Update a reference safely (takes a lock file):

```bash
git update-ref refs/heads/main abc1234     # Set to OID
git update-ref refs/heads/main abc1234 def5678  # Only if current value is def5678
git update-ref -d refs/heads/feature       # Delete ref
git update-ref -m "create branch" refs/heads/new abc1234  # With reflog message
```

### git symbolic-ref

Read or write symbolic references:

```bash
git symbolic-ref HEAD              # Read: refs/heads/main
git symbolic-ref HEAD refs/heads/feature  # Write
git symbolic-ref -q HEAD           # Quiet (exit 1 if not symbolic)
```

### git for-each-ref

Iterate over references (scriptable):

```bash
git for-each-ref                    # List all refs
git for-each-ref refs/heads/        # Only branches
git for-each-ref --format='%(refname:short) %(objectname:short)' refs/tags/
git for-each-ref --sort=-committerdate refs/heads/  # Sort by date
git for-each-ref --count=5 refs/heads/  # Limit count
```

### Format Fields

|| Field | Description |
||-------|-------------|
|| `%(refname)` | Full ref name |
|| `%(refname:short)` | Short name (e.g., `main` not `refs/heads/main`) |
|| `%(objectname)` | Full OID |
|| `%(objectname:short)` | Abbreviated OID |
|| `%(objecttype)` | Object type |
|| `%(objectsize)` | Object size |
|| `%(authorname)` | Author name |
|| `%(authoremail)` | Author email |
|| `%(authordate)` | Author date |
|| `%(committerdate)` | Committer date |
|| `%(upstream)` | Upstream ref name |
|| `%(upstream:short)` | Short upstream name |
|| `%(ahead)` | Commits ahead of upstream |
|| `%(behind)` | Commits behind upstream |

## Revision Parsing

### git rev-parse

Parse and resolve revision expressions:

```bash
git rev-parse HEAD                  # Full OID of HEAD
git rev-parse --short HEAD          # Abbreviated OID
git rev-parse --abbrev-ref HEAD     # Branch name (e.g., main)
git rev-parse --show-toplevel       # Repository root
git rev-parse --git-dir             # .git directory path
git rev-parse --is-inside-work-tree # true/false
git rev-parse --is-bare-repository  # true/false
git rev-parse main~3                # 3rd ancestor of main
git rev-parse main^2                # 2nd parent of main
git rev-parse --verify <ref>        # Verify ref exists
git rev-parse --quiet <ref>         # Suppress errors
```

### git rev-list

List commit OIDs in a range:

```bash
git rev-list HEAD                   # All commits reachable from HEAD
git rev-list main..feature          # Commits in feature not in main
git rev-list --count HEAD           # Count commits
git rev-list --objects HEAD         # Include tree and blob OIDs
git rev-list --all                  # All reachable commits
git rev-list --since=2024-01-01 HEAD  # Filter by date
git rev-list --author=Alice HEAD    # Filter by author
git rev-list --merges HEAD          # Only merge commits
git rev-list --no-merges HEAD       # No merge commits
git rev-list --first-parent HEAD    # Follow only first parent
```

## Pack Operations

### git pack-objects

Create a packfile from a list of objects:

```bash
git rev-list --objects HEAD | git pack-objects --stdout > pack.pack
git rev-list --objects HEAD | git pack-objects pack-name  # Creates pack-name.pack + .idx
```

Options:
- `--delta-base-offset` — Use OFS_DELTA (default in v2)
- `--thin` — Create thin pack (for push)
- `--depth=<n>` — Max delta chain depth (default 50)
- `--window=<n>` — Window size for delta search (default 10)
- `--max-pack-size=<n>` — Split packs above this size

### git index-pack

Build an index from a packfile:

```bash
git index-pack <pack-file>          # Build .idx
git index-pack --fix-thin <pack>    # Resolve thin pack (for fetch)
git index-pack --verify <pack>      # Verify pack integrity
git index-pack -v <pack>            # Verbose
```

### git unpack-objects

Unpack a packfile into loose objects:

```bash
git unpack-objects < <pack-file>
```

### git verify-pack

Verify and inspect a packfile:

```bash
git verify-pack -v <pack-file>      # Verbose: show all objects
git verify-pack -s <pack-file>      # Show only objects with delta chains
```

## Merge Base and Ancestry

### git merge-base

```bash
git merge-base A B                  # Best common ancestor
git merge-base --all A B            # All common ancestors
git merge-base --is-ancestor A B    # Exit 0 if A is ancestor of B
git merge-base --independent A B    # Show independent refs
git merge-base --octopus A B C      # Best common ancestor for octopus merge
```

### git name-rev

Find symbolic names for commits:

```bash
git name-rev <OID>                  # e.g., "main~3"
git name-rev --name-only <OID>      # Just the name
git name-rev --all                  # Name all commits
```

## Attribute and Ignore Inspection

### git check-attr

Check gitattributes for a path:

```bash
git check-attr merge <path>         # Check merge driver
git check-attr diff <path>          # Check diff driver
git check-attr -a <path>            # All attributes
```

### git check-ignore

Check if a path is ignored:

```bash
git check-ignore <path>             # Exit 0 if ignored
git check-ignore -v <path>          # Show which rule matches
git check-ignore -a <path>          # Show all matching rules
```

### git check-mailmap

Resolve contact info through mailmap:

```bash
git check-mailmap <email>           # Resolve canonical name/email
```

## Environment Variables

|| Variable | Description |
||----------|-------------|
|| `GIT_DIR` | Override `.git` directory path |
|| `GIT_WORK_TREE` | Override working tree path |
|| `GIT_INDEX_FILE` | Override index file path |
|| `GIT_OBJECT_DIRECTORY` | Override object directory |
|| `GIT_ALTERNATE_OBJECT_DIRECTORIES` | Additional object directories (colon-separated) |
|| `GIT_AUTHOR_NAME/EMAIL/DATE` | Override author info |
|| `GIT_COMMITTER_NAME/EMAIL/DATE` | Override committer info |
|| `GIT_CONFIG` | Single config file (disables system/global) |
|| `GIT_CONFIG_GLOBAL` | Override global config file |
|| `GIT_CONFIG_SYSTEM` | Override system config file |
|| `GIT_TRACE` | Trace Git commands (1=summary, 2=detailed, path=file) |
|| `GIT_TRACE_PACKET` | Trace pack protocol |
|| `GIT_TRACE_PERFORMANCE` | Trace performance |
|| `GIT_TRACE_SETUP` | Trace setup |
|| `GIT_LITERAL_PATHSPECS` | Disable glob in pathspecs |
|| `GIT_GLOB_PATHSPECS` | Enable glob in pathspecs |
|| `GIT_NO_REPLACE_OBJECTS` | Disable replace objects |
|| `GIT_PREFIX` | Subdirectory prefix (set by Git, not users) |
|| `GIT_CEILING_DIRECTORIES` | Don't look for .git above these dirs |
|| `GIT_DISCOVERY_ACROSS_FILESYSTEM` | Allow .git discovery across mount points |

## Alternates

Share objects between repositories without copying:

```
.git/objects/info/alternates
```

Content: one path per line to another repository's object directory:

```
/path/to/other-repo/.git/objects
```

Objects are looked up in the alternates first, then locally. Used by:
- `git clone --reference` (reference clone)
- `git clone --dissociate` (reference then copy)
- Fork networks on hosting platforms

## Replace Objects

Replace one object with another without rewriting history:

```bash
git replace <original-oid> <replacement-oid>
git replace --edit <oid>            # Edit commit message in replacement
git replace --graft <parent>...     # Create replacement with new parents
git replace -d <oid>                # Delete replacement
git --no-replace-objects log        # Disable replace objects
```

Replace refs are stored in `refs/replace/`. They are **not** pushed by default.

## Edge Cases and Known Issues

- **Plumbing commands don't run hooks**: Most plumbing commands bypass the hook system. This is by design — they are low-level tools.
- **`git update-index` and line endings**: Direct index manipulation bypasses `core.autocrlf` conversion. Use `git add` for correct line ending handling.
- **`git read-tree -m` stages**: The three-way merge mode populates stages 1, 2, and 3 in the index. This is how `git merge` sets up the conflict state.
- **Alternates and gc**: `git gc` in a repository with alternates will not prune objects that exist only in alternates. `git repack -a -d` may fail if alternates are removed.
- **`GIT_DIR` and worktrees**: Setting `GIT_DIR` without `GIT_WORK_TREE` can cause unexpected behavior. Use `git -C <path>` instead.

## References

- <https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain>
- <https://git-scm.com/docs/gitrepository-layout>

## Related Skills

- For object types, see [objects.md](objects.md)
- For reference types, see [refs.md](refs.md)
- For index format, see [index.md](index.md)
- For core concepts, see [concepts.md](concepts.md)
