# Overview

What Lore is, its design goals, architecture, and how it compares to other version control systems.

> **Source of truth**: <https://epicgames.github.io/lore/explanation/system-design/>

## What Lore Is

Lore is a **centralized, binary-first, MIT-licensed version control system** from Epic Games, designed for projects that combine code with large binary assets (games, entertainment, media). It treats all content as opaque byte streams — text and binary flow through the same primitives.

The project is organized as a **Rust library, server, and CLI**, with publicly specified and versioned data formats and wire protocols. Pre-built binaries ship for Windows, macOS (ARM64), and Linux (x86-64, ARM64).

## Two Subsystems

Lore is structured as two independently usable subsystems:

- **Storage Subsystem**: A partition-based, content-addressed store that deduplicates all content while enforcing strict per-partition access boundaries. Usable on its own for caches, asset pipelines, backup targets.
- **Version Control Subsystem**: Revisions, branches, merges, staging, sync, push, diff, query — built on top of the storage subsystem, exposed through its own public API. Version control is one consumer of the storage API, not a privileged layer.

## Design Goals

- **Binary-first**: All content is opaque byte streams on the hot path. No line-ending translation, no text encoding inference, no clean/smudge filters.
- **Centralized but offline-capable**: Staging, committing, branching, switching, diffing never require a network round trip.
- **Sparse by construction**: Instances materialize only the subset of files asked for, via view filters and lazy fragment fetching.
- **Atomic state**: Every operation completes fully or leaves the repository unchanged.
- **Cryptographically verifiable**: BLAKE3 hashes, hash-chained revision graph.
- **API-first**: A C header is the canonical interface. Language bindings (Rust, Python, JS, C#, Go) wrap the same operations.
- **Multi-tenant safe**: Partition-level isolation with cross-partition dedup but no cross-partition access.
- **Backend-scalable**: Horizontal scaling, hot/warm/cold tiering, edge caching, read replicas.
- **Replaceable backends**: Thin `ImmutableStore` and `MutableStore` trait interfaces. Shipped backends include local file, S3, and DynamoDB.
- **Lifecycle-complete**: Content can be added, deduplicated, and removed (obliterated).
- **Open and free**: MIT license, publicly versioned specs.

## Explicit Non-Goals

- **Peer-to-peer decentralization** (centralized by design)
- **Adversarial-server threat model** (clients hash-validate fragments; trust in the server is bounded to pointers, names, and access decisions)

## Architecture

```
┌──────────────────────────────────────────────────┐
│                   Lore CLI                        │
│  (thin layer over the C API, all commands)        │
└──────────────┬───────────────────────────────────┘
               │
┌──────────────▼───────────────────────────────────┐
│            Version Control Subsystem              │
│  revisions, branches, merges, staging, sync,      │
│  push, diff, query, metadata, locks, links, layers│
└──────────────┬───────────────────────────────────┘
               │
┌──────────────▼───────────────────────────────────┐
│              Storage Subsystem                    │
│  Immutable Store (content-addressed, BLAKE3)      │
│  Mutable Store (key-value, branch pointers, cas)  │
└──────────────┬───────────────────────────────────┘
               │
┌──────────────▼───────────────────────────────────┐
│           Replaceable Backends                    │
│  local (packfiles), S3, DynamoDB, composite,      │
│  replicated, remote, plugins                      │
└──────────────────────────────────────────────────┘
```

## Storage Protocol

One logical command set served on two transports:

- **QUIC** (ALPN `lore-storage/0.4`): Binary, low-overhead, up to 8 bidirectional streams per connection, command pipelining.
- **gRPC** (HTTP/2 + protobuf): Where QUIC is unavailable or existing HTTP infrastructure is preferred.

Commands: Authorize, Get, Put, Query, Verify, Copy (immutable store) + MutableLoad, MutableStore, MutableCas (mutable store). One fragment per command, max 256 KiB.

## How Lore Differs from Other VCS

| Aspect | Git | Perforce | Lore |
|--------|-----|----------|------|
| **Binary assets** | LFS bolted on, second-class | First-class but proprietary | First-class, chunked, deduplicated |
| **Offline** | Fully offline | Server round-trip required | Staging/committing/branching offline |
| **Sparse checkouts** | Sharp edges, offline issues | Server-dependent | Lazy by default, view filters |
| **Deduplication** | File-level only | Limited | Fragment-level, content-defined |
| **Multi-tenant** | Per-repo, no native isolation | Per-depot, limited | Partition-level, access-controlled |
| **License** | GPL | Proprietary | MIT |
| **Protocol** | Custom | Proprietary | Open specification, QUIC/gRPC |