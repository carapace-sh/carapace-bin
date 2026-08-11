---
name: lore
description: >
  Use when working with the Epic Games Lore version control system — a centralized VCS
  designed for game development with binary-first content-addressed storage. Covers
  architecture, core concepts, server deployment, CLI and server configuration, workflows,
  and the full command reference. Triggers on: "lore", "Lore", "Epic Games Lore",
  "loreserver", "version control system", "VCS", "content-addressed", "immutable store",
  "mutable store", "BLAKE3", "Merkle tree", "FastCDC", "fragment", "partition",
  "content-defined chunking", "split-write", "shared store", "revision", "branch",
  "merge", "cherry-pick", "revert", "bisect", "layer", "link", "lock", "staging",
  "dirty", "obliterate", "packfile", "LoreError", "config.toml", "cli.toml",
  "lore://", "Lore server", "Lore CLI", "Lore configuration", "Lore FAQ",
  "Lore roadmap", "Lore system design".
user-invocable: true
---

# Lore In-Depth Reference

Lore is a centralized, binary-first, MIT-licensed version control system from Epic Games, designed for projects that combine code with large binary assets. Source: <https://epicgames.github.io/lore/>

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| overview, what is Lore, design goals, non-goals, architecture, two subsystems, storage, version control, API-first, binary-first, centralized, offline-capable, replaceable backends, specification, open source, MIT | [references/overview.md](references/overview.md) |
| core concepts, revision, branch, immutable store, mutable store, content addressing, BLAKE3, hash, context, address, partition, Merkle tree, revision state, node block, fragment, blob, chunking, FastCDC, fixed-size, compression, packfile, revision graph, parent, DAG, staging, dirty, staged anchor, sparse, view, .loreignore, lazy fetch, metadata, link, layer, lock, obliteration, shared store, instance | [references/core-concepts.md](references/core-concepts.md) |
| setup, install, Lore CLI, prebuilt binary, build from source, loreserver, Docker, deploy, certificate, ports, durable storage, health check, demo mode, local server, quickstart, prerequisites | [references/setup.md](references/setup.md) |
| configuration, config.toml, cli.toml, remote_url, identity, store, max_capacity, eviction, compaction, direct_write, shared_store_to_use, pager, loreserver config, --config, --env, config layering, default.toml, local.toml, environment, LORE__ env vars, server settings, QUIC, HTTP, gRPC, auth, JWT, JWKS, topology, plugins, hooks, telemetry, notification | [references/config.md](references/config.md) |
| CLI reference, command tree, lore, repository, create, clone, status, info, list, delete, verify, dump, gc, store, metadata, instance, config, update-path, branch, list, create, switch, push, merge, unresolve, into, start, restart, resolve, mine, theirs, abort, diff, archive, reset, protect, unprotect, latest, revision, history, commit, amend, sync, bisect, diff, find, restore, cherry-pick, revert, file, info, stage, dirty, unstage, reset, obliterate, history, diff, write, hash, dependency, add, remove, list, auth, login, logout, clear, layer, add, remove, link, add, remove, update, list, lock, acquire, status, query, release, logfile, login, status, clone, stage, dirty, unstage, reset, diff, history, commit, sync, push, service, run, start, stop, notification, subscribe, completions, shared-store, create, info, set-use-automatically, global options | [references/cli.md](references/cli.md) |
| workflow, quickstart, tutorial, create repository, stage, commit, push, clone, shared store, branch, switch, merge, resolve conflicts, sync, amend, cherry-pick, revert, bisect, restore, file history, diff, obliterate, view, sparse checkout, .loreignore, file locking, dependency, blame, find revision | [references/workflow.md](references/workflow.md) |
| FAQ, roadmap, how is Lore different, Git vs Perforce, platforms, production readiness, locking, merge conflicts, security, contribute, license, desktop client, UEFN, 2026 roadmap, 2027 roadmap, VFS, VS Code plugin, edge instances, forks, governance | [references/faq-roadmap.md](references/faq-roadmap.md) |

## Quick Guide

- **What is Lore and what problem does it solve?** → [references/overview.md](references/overview.md)
- **How does the revision graph and content addressing work?** → [references/core-concepts.md](references/core-concepts.md)
- **How do I install the CLI and start a server?** → [references/setup.md](references/setup.md)
- **How do I configure the CLI or server?** → [references/config.md](references/config.md)
- **How do I create a repo, commit, branch, merge, and sync?** → [references/workflow.md](references/workflow.md)
- **What CLI commands are available?** → [references/cli.md](references/cli.md)
- **How is Lore different from Git/Perforce?** → [references/faq-roadmap.md](references/faq-roadmap.md)
- **What is on the roadmap?** → [references/faq-roadmap.md](references/faq-roadmap.md)
- **How do I deploy a persistent server with certificates?** → [references/setup.md](references/setup.md)
- **How do I configure server settings, auth, topology, or telemetry?** → [references/config.md](references/config.md)
- **How do I install shell completions?** → [references/setup.md](references/setup.md)
- **How do I resolve merge conflicts?** → [references/workflow.md](references/workflow.md)
- **How do I manage file locks?** → [references/workflow.md](references/workflow.md)
- **How do I use links and layers?** → [references/core-concepts.md](references/core-concepts.md)
- **How do I obliterate content?** → [references/core-concepts.md](references/core-concepts.md)
- **How do I contribute to Lore?** → [references/faq-roadmap.md](references/faq-roadmap.md)

## Cross-Project References

- For **git** internals (comparison mental model), see the **git** skill.
- For **jj** (another VCS with similar concepts), see the **jj** skill.