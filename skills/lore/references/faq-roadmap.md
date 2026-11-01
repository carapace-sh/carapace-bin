# FAQ and Roadmap

Frequently asked questions about Lore and the planned feature roadmap.

> **Source of truth**: <https://epicgames.github.io/lore/faq/> and <https://epicgames.github.io/lore/roadmap/>

## FAQ

### What is Lore?

A next-generation open source version control system from Epic Games, designed for unprecedented scalability — both in data size and the size/distribution of teams. Optimized for projects combining code with large binary assets (games, entertainment). MIT licensed.

### How does Lore differ from Git?

- **Git**: Content-addressed revision graph is excellent, but binary files are second-class — large files require LFS, sparse checkouts have sharp edges in offline use, no native multi-tenant isolation.
- **Lore**: Binary-first (chunked, deduplicated), lazy sparse working trees, fragment-level dedup effective on both multi-gigabyte binaries and kilobyte text files, free branching, full offline operation for staging/committing/branching, partition-level multi-tenant isolation.

### How does Lore differ from Perforce?

- **Perforce**: Handles large binary assets but requires server round trips for everyday operations, proprietary wire protocols, limited binary-level deduplication.
- **Lore**: Centralized server-of-record but offline-capable, content-addressed with fragment-level dedup, open specification and MIT license, QUIC/gRPC protocols.

### What platforms are supported?

Windows, macOS (ARM64), Linux (x86-64, ARM64). Pre-built binaries for all three. Docker image for server deployment. Language bindings: C/C++, Rust, JavaScript, C#, Python, Go.

### Is Lore production-ready?

Pre-stable 0.x release. APIs and protocols may evolve before 1.0 (when strict backward compatibility applies). Data written now is designed to remain readable by every future release.

### Does Lore support file locking?

Basic locking exists — users can place locks on non-mergeable assets to signal editing intent. **Enforcement is not yet in place** (informs rather than blocks). Next iteration focuses on cross-branch lock scalability for millions of files and thousands of concurrent users.

### How does Lore handle merge conflicts?

- **Text files**: Standard three-way merge.
- **Binary assets**: Conflicts surface as explicit divergence — user chooses which version to carry forward. File locking is the recommended workflow for binaries.

### How is Lore licensed?

MIT license. Fully open source, no restrictions or license fees.

### Is there a desktop client?

An early desktop client binary is available on the releases page. It is **not yet open source** (depends on some proprietary components including Epic's internal design system). Open-sourcing is committed for the future.

### Where is Lore already in use at Epic?

Lore (formerly **Unreal Revision Control**) is the built-in VCS for UEFN (Unreal Editor for Fortnite). Also seeing progressive adoption by internal Epic teams and being implemented as the backing store for UEFN's cook pipeline.

### How can I contribute?

All contributions welcome via GitHub. Uses **DCO (Developer Certificate of Origin)** rather than a CLA. Join the [Discord](https://discord.gg/E4SFJKRPbg). For significant changes (wire protocol, on-disk format, public APIs), use a **Lore Enhancement Proposal (LEP)**.

## Roadmap

### 2026 — Foundations

| Feature | Description |
|---------|-------------|
| **Lore OSS and UEFN compatibility** | Converge OSS and UEFN implementations. Key gap: UEFN uses Oodle compression, Lore uses Zstandard — actively moving UEFN to Zstandard. |
| **VS Code plugin** | Graphical interface in VS Code. Both plugin and source will be released as a complementary repository. |
| **Scalable file locking** | Scale locking to enforce single-editor access across millions of files and thousands of concurrent users. |
| **Virtual file system (VFS)** | Load files lazily as you open them, serving from Lore's shared store with fragment-level dedup. Eliminates up-front full clones. |
| **Links and layers** | Unlock multi-repository composition into a single working tree with per-repository access control. Data model already supports this. |

### 2027 — Scale and Collaboration

| Feature | Description |
|---------|-------------|
| **Desktop client** | Open-source the desktop client. |
| **Unreal Editor plugin** | Visual version control inside the Unreal Editor, built on the plugin that already ships with UEFN. |
| **Web client & code review tools** | Open-source browser-based home for code review, discussion, and repository management. |
| **Edge instances & advanced server topologies** | Server replication and caching for extreme distributed-team use cases. |

### Later — Exploring

| Feature | Description |
|---------|-------------|
| **Forks and isolated partitions** | Fork projects as independent, access-controlled copies that can merge back. Uses copy-on-write. |

### Governance

Evolving toward a **technical steering group** drawn from internal and external contributors, operating through public roadmaps, RFCs, and open meetings.