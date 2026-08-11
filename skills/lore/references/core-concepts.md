# Core Concepts

The Lore version control system's conceptual model — revisions, branches, the immutable store, layers, links, and identity.

> **Source of truth**: <https://epicgames.github.io/lore/reference/lore-cli-commands/>

## Repository

A Lore repository stores versioned content as an **immutable directed acyclic graph (DAG) of revisions**. Each revision is a snapshot of the repository's file tree at a point in time, identified by a content-addressed hash signature.

## Revisions and Branches

- **Revision**: A snapshot of the entire repository state, identified by a hash signature. Revisions form a parent-child DAG.
- **Branch**: A named pointer to a revision (the "latest" on that branch). Branches track the tip of a line of development.
- **Revisions are immutable** — once committed, they never change. Branches move to point at new revisions.
- **Metadata**: Key-value pairs attachable to repositories, branches, revisions, and files. Values can be strings, binary files, or numeric (u64).

## Immutable Store

The local store is content-addressed — every blob and tree fragment is stored by its hash. This enables:

- **Deduplication**: identical content stored once
- **Integrity verification**: hash verification catches corruption
- **Partial sync**: only fetch needed fragments
- **Shared store**: multiple repositories can share a single store (deduplicated across repos)

## Fragments and Blobs

- **Blob**: Raw file content, identified by its hash
- **Fragment**: A node in the Merkle tree — either a blob or a tree node listing children
- **Address**: The hash identifier used to query fragments in the store

## File System Model

Lore uses a **split-write filesystem** for virtual clones: metadata is tracked in the store while file content is written on demand. This supports:

- **Virtual clones**: clone without all files, fetch on access
- **Direct file write**: write directly to destination instead of temp+move
- **View filters**: client-side filter files that control which files are materialized

## Layers

Layers overlay one repository's content onto another at a mount point. A layer repository's content is merged at a specific path in the current repository. Layer revisions are matched via a metadata key.

## Links

Links embed one repository's content at a subpath of another, similar to git submodules. Links can be pinned to a specific branch or revision. Automatic branching in linked repos can be disabled.

## Locks

File-level locking to coordinate exclusive write access. Locks are associated with a branch and owner. Supports querying locks by branch, owner, or path.

## Dependencies

Files can declare dependency edges to other files, optionally tagged. Dependency-based selective clone/sync fetches only files reachable from root files, respecting tags and transitive depth limits.

## Identity

The CLI uses **identities** tied to authentication tokens. The `--identity` global option picks which identity to use. Auth supports multiple stored identities with token types including `api-key`, `eg1`, and `lore`.

## Global Options

| Flag | Description |
|------|-------------|
| `--repository <path>` | Use given path as repository path |
| `--log-level <level>` | Set the logging level |
| `-d`, `--debug` | Enable debug output |
| `-f`, `--force` | Force the operation if possible |
| `--dry-run` | Dry run mode, only report what would have changed |
| `-P`, `--no-pager` | Disable pagination |
| `--offline` | Force offline mode |
| `--remote` | Use remote data |
| `--local` | Use local data |
| `--identity <IDENTITY>` | Use given identity |
| `--max-connections <MAX_CONNECTIONS>` | Set max parallel connections |
| `--file-count-limit <count>` | Set max parallel files opened |
| `--file-size-limit <size>` | Set max total bytes of parallel files opened |
| `--compress-limit <count>` | Set max parallel compress operations |
| `--search-limit <SEARCH_LIMIT>` | Set max revisions to search when matching/finding |
| `--search-nearest` | Search for nearest match when matching revisions |
| `--gc` | Run automatic garbage collection on local store in background |
| `--sync-data` | Force sync data to storage media during flush |
| `--non-interactive` | Disable interactive prompts |