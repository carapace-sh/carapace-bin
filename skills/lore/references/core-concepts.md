# Core Concepts

The Lore VCS conceptual model — revisions, branches, the two stores, content addressing, Merkle tree, staging, metadata, links, layers, locks, and obliteration.

> **Source of truth**: <https://epicgames.github.io/lore/explanation/system-design/>

## The Two Stores

### Immutable Store (Content-Addressed)

- Holds every byte ever written — file payloads, fragmented file pieces, serialized revision states, tree nodes, metadata blobs.
- Entries are addressed by the **BLAKE3 hash** (32 bytes) of their content.
- **Append-only**: entries can be added or obliterated, never modified.
- On disk: stored in **packfiles** (large append-only files with mmappable indexes).
- On the wire: fragments travel individually (one per command), for parallelism and resumability.

### Mutable Store (Key-Value)

- Holds branch latest pointers, name-to-ID mappings, repository catalog entries.
- Narrow API: `load`, `store`, `cas`, `list`.
- Small in volume but the only place where consistency, serialization, and contention live.
- `cas` (compare-and-swap) is the atomicity primitive — the consistency hot spot.

## Content Addressing

| Concept | Size | Description |
|---------|------|-------------|
| **Hash** | 32 bytes | BLAKE3 hash of the content payload. The address function throughout. |
| **Context** | 16 bytes | Opaque tag carried alongside a hash. Tracks identity (file ID for move/copy, obliteration scope). Not an access boundary — it's a deduplication construct. |
| **Address** | 48 bytes | The pair `(hash, context)` — the unit by which fragments are stored, looked up, transferred, and obliterated. |

**Properties**: automatic deduplication (same bytes → same hash), automatic integrity (re-hash to verify), idempotent transfer, forced immutability.

## Partitions

A **partition** is a 16-byte opaque identifier that scopes content in the storage layer:

- Authorization binds a session to a partition.
- Content lookups are partition-scoped — a session bound to partition A cannot look up fragments in partition B.
- Underneath, the storage subsystem is free to deduplicate identical bytes across partitions.
- The partition is **derived from the authenticated session** (JWT), not asserted by the client.

## Revisions and the Merkle Tree

### Revision

A frozen snapshot of the entire repository tree, identified by the hash of its serialized state.

**Revision State** (320-byte fragment in the immutable store):

| Field | Description |
|-------|-------------|
| Magic number + format version | Identifies the fragment type |
| Revision number | Integer, increments by one along a branch's first-parent chain |
| Parent hashes | One (ordinary) or two (merge) hash signatures |
| Merkle tree hash | Hash of the serialized Merkle tree |
| Metadata hash | Commit message, author, timestamps, key-value metadata |
| Link list hash | Serialized list of sub-repository links |
| Repository ID | For second parent in cross-repository merges |
| Reserved | Forward compatibility |

### Merkle Tree

Stored as a sequence of **node blocks**, each holding 512 fixed-size **node** records (96 bytes each) + 128-byte header = 49,280 bytes. Each node carries:

- Flags, file mode, tree indexes (parent, child, sibling — 32-bit)
- Name reference (offset + length into name table + 64-bit lowercase hash for fast lookup)
- File size, content address

**Structural dedup**: Two revisions differing in one file share every node block not on the path from root to that file. Storage cost grows with what changed, not with repository size.

### Revision Graph

Revisions form a **directed acyclic graph** where every edge is a cryptographic link. Each revision has one parent (ordinary) or two (merge). The entire history of a branch is verifiable from any current revision hash.

## Fragments and Chunking

### Content-Defined Chunking (FastCDC)

- Rolling hash, boundaries determined by content.
- Targets 64 KiB average, 32 KiB floor, 256 KiB ceiling.
- Good for sparse writes inside large files.
- Uses temporal-coherence strategy (reuse prior boundaries where bytes unchanged) to preserve dedup.

### Fixed-Size Chunking

- Boundaries at fixed offsets.
- Canonical addressing — same content always produces the same chunks.
- Good when canonical addressing matters more than dedup robustness.

### Recursive Fragmentation

If a fragment list exceeds the chunking threshold, the list is itself chunked and stored as a fragment with a list flag, forming a tree of fragment lists.

### Fragment References

Each records the hash of the child fragment and the byte offset it represents. Strictly ordered by offset, so binary search can find the fragment covering any offset in O(log n) time.

### Compression

Per-fragment, orthogonal to addressing. Zstd today, open-ended codec list. Address is hash of *uncompressed* content — changing compression doesn't change the address.

## Branches

- A **branch** is a named, mutable pointer in the mutable store to a **latest revision** (Lore's analogue of Git's HEAD).
- No separate "branch object" in the immutable store.
- Branch creation is just a mutable-store insertion — new branch ID, name-to-ID mapping, latest pointer.
- Each branch has a stable opaque ID (UUIDv7) and a human-readable name. The name-to-ID mapping is itself a mutable-store entry. A branch ID never changes; a name can be archived, restored, or reused.
- **Merge**: A revision with two parents. Both parent state hashes recorded in the 320-byte state. Tree is the combination of the two parent trees, with conflicts resolved at commit time.
- **Rebase, cherry-pick, squash**: Each produces new revisions from existing ones. No operation rewrites history in place — each "rewrite" is a fresh sequence of revisions plus a re-pointing of the branch's latest pointer. Original revisions still exist in the immutable store but become unreachable from the branch pointer.

## State Detection and Staging

- **Filesystem as ground truth**: When Lore needs to know what's in a file, it reads the file. No intermediary index, no reconciliation step needed.
- **Modification tracking**: Lives directly on Merkle tree nodes. Each node carries a **dirty** flag (orthogonal to **staged** flag). Dirty propagates to parent directories.
- **Staged anchor**: A per-instance pointer in the mutable store to the state tree. Both dirty and staged state are persisted here.

**Three paths into the dirty set**:
1. **Notification**: `file dirty` operation — flag update, no content read.
2. **Scanning**: Walks the working tree, compares against committed revision.
3. **Verification**: Re-examines only files already marked dirty (bounded by dirty set size, not working tree).

**Staging as recorded intent**: Staging pins the *path*, not the content. Between staging and committing, the user can keep editing. Fragments are produced at commit time, not stage time.

## Sparseness and Partial Working Copies

- **View file** (`.lore/view`): Inbound filter declaring which paths the user wants on disk. Paths outside the view are not materialized. Local to a client — doesn't travel with a clone.
- **Ignore file** (`.loreignore`): Outbound filter declaring paths excluded from staging and committing.
- **FilterMode**: Operations on committed state → consult view only. Operations on local state → consult both.
- **Lazy fetch**: Fragments are fetched on demand. Only the parts of the tree the view asks for are walked. Reading a 4 MiB range from a multi-gigabyte file fetches only the overlapping fragments.
- **Local cache**: LRU policy, user-sized budget. Multiple instances on the same machine can share a single local store.

## Metadata

A typed key-value store attached to entities. Stored as content-addressed blobs in the immutable store.

| Attachment type | Mutability | Description |
|----------------|------------|-------------|
| **Revision metadata** | Immutable | Referenced from the 320-byte revision state. Part of the committed revision. |
| **File metadata** | Immutable per revision | Stored in file metadata block stream parallel to node block stream. |
| **Branch metadata** | Mutable | Referenced by mutable pointer. Can change over time. |
| **Repository metadata** | Mutable | Same shape as branch metadata. |

**Format**: Binary typed key-value array. Value types: address, boolean, context, hash, numeric, string, or binary. Capped at 1 MiB.

**Key space**: Built-in keys (reserved by the system) + application keys (free for any tool to use).

## Sub-Repository Links

A **link** is a reference from one repository to a specific revision of another, mounted at a path in the parent's tree. The link is recorded in the parent's revision and travels with it.

- Each linked repository is its own partition with its own access control.
- Source path remapping: a linked repo's `lib/widgets` can appear as `vendor/widgets` in the parent.
- **Auto-follow**: When enabled, creating a branch in the parent creates a corresponding branch in each linked repository. Can be disabled for vendored third-party content.
- **Link-scoped operations**: `commit --link <path>` and `branch merge start <branch> --link <path>` work on a single link.

## Layers

A **layer** is a local overlay of one repository's content into another at a path. Applied locally, not stored in the parent's revision. Two machines on the same revision can have different layer configurations.

- **Key differences from links**: Links are versioned dependencies every consumer sees; layers are local decorations.
- **Revision matching**: Metadata-keyed matching (matches by metadata values) and branch auto-follow on switch.
- **Per-layer staging and commit**: Layers participate in staging/commit alongside the parent. Per-layer commit messages supported.

## File Locking

File-level locking to coordinate exclusive write access. Locks are associated with a branch and owner. Supports querying locks by branch, owner, or path.

**Current state**: Basic locking exists — users can place locks to signal non-mergeable asset editing. Enforcement is not yet in place (the current implementation informs rather than blocks). The next iteration focuses on cross-branch lock scalability for millions of files and thousands of concurrent users.

## Obliteration (Data Removal)

Removes a fragment's **payload** while keeping its **address** intact in the store's index. Readers get a typed "obliterated" response, not corrupted bytes.

- **Two-phase**: `PayloadObliterating` → `PayloadObliterated`. Crash-safe design.
- **File ID as scope**: When a file is obliterated, every fragment whose address matches the file's context is obliterated. Cost is proportional to the number of fragments belonging to the file, not the repository size.
- **What it cannot do**: Recover old hashes, un-leak data already cloned, forget a revision existed, or substitute different content.

## Shared Stores and Instances

- **Shared store**: A single on-disk store (immutable + mutable) referenced by multiple working directories.
- **Instance**: A working directory with its own state (UUIDv7 in `.lore/instance`). Per-instance state: current revision anchor, current branch, staged anchor, metadata blob.
- **No main repository**: All instances are peers. Removing any one leaves others intact. The shared store persists as long as something is configured to use it.

## Identity

The CLI uses **identities** tied to authentication tokens. The `--identity` global option picks which identity to use. Auth supports multiple stored identities with token types including `api-key`, `eg1`, and `lore`. Commit identity is resolved at create/clone time via `--identity` flag → server connection identity → unset. If unset and you try to commit, Lore fails with a clear error message.