# Git Object Model

The four object types that form Git's content-addressable storage — blobs, trees, commits, and tags — plus packfiles, deltas, and garbage collection.

> **Source of truth**: <https://git-scm.com/book/en/v2/Git-Internals-Git-Objects>, <https://git-scm.com/docs/gitformat-pack>. For **core concepts**, see [concepts.md](concepts.md). For **references that point to objects**, see [refs.md](refs.md). For **plumbing commands**, see [internals.md](internals.md).

## Object Types

Every Git object shares a common format:

```
type content-size\0content
```

The SHA-1 (or SHA-256) hash of this entire byte sequence is the **object ID** (OID).

|| Type | Purpose | Content Format |
||------|---------|----------------|
|| **blob** | File content | Raw file data (no metadata — no name, no permissions) |
|| **tree** | Directory listing | List of `(mode, name, OID)` entries |
|| **commit** | Snapshot metadata | Tree OID, parent OIDs, author, committer, message |
|| **tag** | Annotated tag | Object OID, type, tag name, tagger, message |

## Blob Objects

A blob stores **raw file content** with no metadata:

```
$ git hash-object -w <<< "hello"
ce013625030ba8dba906f756967f9e9ca394464a
```

Key properties:

- **No filename**: The name is stored in the parent tree, not the blob
- **No permissions**: Mode is stored in the tree entry
- **Deduplication**: Two identical files (any name) share one blob
- **Binary-safe**: Blobs store raw bytes — text and binary are treated identically

### Empty Blob

The empty blob has a well-known hash:

```
SHA-1:   e69de29bb2d1d6434b8b29ae775ad8c2e48c5391
SHA-256: 473a0f4bc3b3949c6c3c5cb3b5445309f1e4a9e3d8f8f6a1e1c1b1a0f1e2d3c4
```

## Tree Objects

A tree is a sorted list of entries, each pointing to a blob or subtree:

```
100644 blob e69de29...  README.md
040000 tree 0ac5bfc...  src
100755 blob 8ba561e...  install.sh
```

### Entry Format

Each entry in the raw tree object:

```
mode name\0OID-as-20-bytes
```

### Mode Values

|| Mode | Type | Description |
||------|------|-------------|
|| `040000` | tree | Subdirectory |
|| `100644` | blob | Regular file (not executable) |
|| `100755` | blob | Executable file |
|| `120000` | blob | Symbolic link (content = link target) |
|| `160000` | commit | Gitlink (submodule pointer) |

### Tree Sorting

Entries are sorted by name using a special comparison: **tree entries sort as if they have a trailing `/`**. This means:

```
foo       (blob)
foo.c     (blob)
foo/      (tree)  ← sorts after foo.c because of the implicit /
```

This matches how directory listings work — a directory `foo/` sorts after files `foo` and `foo.c`.

### Tree Inspection

```bash
git ls-tree HEAD              # list top-level tree
git ls-tree -r HEAD           # recursive listing
git ls-tree -r -t HEAD        # recursive, show tree entries too
git cat-file -p HEAD^{tree}   # raw tree content
```

## Commit Objects

A commit records a snapshot of the project tree plus metadata:

```
tree 0ac5bfc5e3c7...
parent d7439b0612e4...
author Alice <alice@example.com> 1700000000 +0000
committer Alice <alice@example.com> 1700000000 +0000

Commit message here
```

### Fields

|| Field | Required | Description |
||-------|----------|-------------|
|| `tree` | Yes | OID of the root tree for this commit |
|| `parent` | No (1+ for merges) | OID(s) of parent commit(s) |
|| `author` | Yes | Name, email, timestamp (who wrote the change) |
|| `committer` | Yes | Name, email, timestamp (who created the commit object) |
|| Message | Yes | Free-form text after a blank line |

### Author vs Committer

- **Author**: The person who originally wrote the change
- **Committer**: The person who created this commit object

These differ when:
- Rebasing (committer changes, author stays)
- Cherry-picking (committer changes, author stays)
- Amending (both change)
- `git commit --author=` (overrides author)

### Timestamp Format

```
name <email> UNIX_TIMESTAMP TIMEZONE_OFFSET
```

Example: `Alice <alice@example.com> 1700000000 +0100`

The timezone is always in `+HHMM` / `-HHMM` format.

### Root Commit

A commit with **no parent** is a root commit. Every repository has at least one. Orphan branches (`git checkout --orphan`) create new root commits.

### Merge Commits

Commits with **two or more parents** are merge commits. The first parent is conventionally the branch you were on when you merged; subsequent parents are the merged branches.

## Tag Objects (Annotated Tags)

An annotated tag is a signed object pointing to another object:

```
object d7439b0612e4...
type commit
tag v1.0.0
tagger Alice <alice@example.com> 1700000000 +0000

Tag message here
```

### Annotated vs Lightweight

|| Property | Annotated Tag | Lightweight Tag |
||----------|--------------|-----------------|
|| Object type | `tag` | `commit` (points directly) |
|| Stored in | Object database | `.git/refs/tags/` |
|| Message | Yes | No |
|| Tagger info | Yes | No |
|| Can be signed | Yes (GPG/SSH) | No |
|| Can tag non-commits | Yes (any object) | No (always commit) |

## Object Storage

### Loose Objects

New objects are stored as individually compressed files:

```
.git/objects/ce/013625030ba8dba906f756967f9e9ca394464a
         ^^  ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
         |   remaining 38 hex chars
         first 2 hex chars = directory
```

The content is zlib-compressed. To read a loose object:

```bash
git cat-file -p <OID>       # pretty-print
git cat-file -t <OID>       # show type
git cat-file -s <OID>       # show size
git cat-file blob <OID>     # raw blob content
```

### Packfiles

When the object database grows, Git packs objects into a single file with delta compression:

```
.git/objects/pack/
├── pack-abc123.idx      # Index for random access
└── pack-abc123.pack     # Packed objects
```

#### Pack Format

A packfile contains:

1. **Header**: `PACK`, version number (2 or 3), number of objects
2. **Objects**: Each object is either:
   - **Full object**: type + size + uncompressed content
   - **OFS_DELTA**: offset + delta against a base object at a known offset in the same pack
   - **REF_DELTA**: base OID + delta against a base object (by OID, used in thin packs)
3. **Trailer**: SHA-1 checksum of the pack content + SHA-1 checksum of the entire pack

#### Delta Compression

Deltas store only the differences from a base object:

```
base object (full content)
  └── delta object (instructions: copy from base, insert new data)
       └── delta of delta (chained deltas, max depth 50)
```

Delta instructions:
- **Copy**: offset + length from base object
- **Insert**: literal new bytes

Delta chains are limited to depth 50 (`pack.depth`) and window size 16 (`pack.window`).

#### Pack Index

The `.idx` file maps OIDs to offsets within the pack for O(1) lookup:

```
Header (magic + version)
Fanout table (256 entries — cumulative count of objects with first byte ≤ N)
SHA-1 table (sorted OIDs)
CRC32 table (per-object checksums)
Offset table (32-bit offsets into pack)
Large offset table (64-bit offsets, if any)
Pack checksum
Index checksum
```

#### Thin Packs

A **thin pack** is created during `git push` — it contains `REF_DELTA` objects whose bases are expected to exist on the receiving side. The receiver resolves them using `git index-pack --fix-thin`.

### Repacking

```bash
git repack -a -d    # Repack all objects into a single pack, delete loose objects
git repack -a -d -l # Same, but keep unreachable objects (for safety)
git gc              # Repack + prune unreachable + other cleanup
git gc --aggressive # Slower but better delta compression
```

## Object Reachability and Garbage Collection

An object is **reachable** if it can be traced from any ref, reflog entry, or the index. Unreachable objects are eventually pruned:

```
git fsck              # Check for dangling/unreachable objects
git fsck --unreachable # List all unreachable objects
git prune             # Remove unreachable loose objects
git prune --expire=2.weeks.ago  # Only prune objects older than 2 weeks
git gc                # Repack + prune (respects gc.pruneExpire)
```

### Protection Mechanisms

- **Reflog**: Keeps unreachable commits accessible for `gc.reflogExpire` (default 90 days)
- **gc.pruneExpire**: Default `2.weeks.ago` — unreachable objects younger than this are preserved
- **`--no-prune`**: `git gc --no-prune` never prunes unreachable objects

## Object Inspection Commands

|| Command | Purpose |
||---------|---------|
|| `git cat-file -t <OID>` | Show object type |
|| `git cat-file -s <OID>` | Show object size |
|| `git cat-file -p <OID>` | Pretty-print object content |
|| `git cat-file blob <OID>` | Raw blob content |
|| `git hash-object` | Compute OID of content |
|| `git hash-object -w` | Compute OID and write to object store |
|| `git verify-pack -v <pack>` | Inspect packfile contents |
|| `git index-pack <file>` | Build index from packfile |

## Edge Cases and Known Issues

- **Corrupted objects**: If a single bit flips in a loose object, `git fsck` will detect the hash mismatch. Recovery requires finding the object from a remote or backup.
- **Packfile corruption**: The pack checksum catches corruption, but individual object CRC32s in the index provide finer-grained detection.
- **Delta chain depth**: Very deep delta chains (approaching 50) slow down object access because each delta must be resolved. `git repack` re-optimizes chains.
- **Large files**: Git stores full copies of each version of a large file. Git LFS or git-fat externalize large blobs.
- **Empty trees**: The empty tree has a well-known hash `4b825dc642cb6eb9a060e54bf899d15363d7aa92` (SHA-1). It can be referenced without creating an object.
- **Submodule entries**: Mode `160000` in a tree points to a commit in another repository, not an object in the current one. `git cat-file` on such an OID fails if the submodule isn't initialized.

## References

- <https://git-scm.com/book/en/v2/Git-Internals-Git-Objects>
- <https://git-scm.com/docs/gitformat-pack>
- <https://git-scm.com/docs/gitformat-index>

## Related Skills

- For core concepts (three states, DAG), see [concepts.md](concepts.md)
- For references that point to objects, see [refs.md](refs.md)
- For index format, see [index.md](index.md)
- For plumbing commands, see [internals.md](internals.md)
