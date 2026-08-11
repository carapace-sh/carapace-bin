---
name: lore
description: >
  Use when working with the Epic Games Lore CLI — a version control system for game
  development. Covers the full command reference including repository, branch, revision,
  file, auth, layer, link, lock, and operational commands. Triggers on: "lore", "lore CLI",
  "Epic Games Lore", "lore command", "lore repository", "lore branch", "lore revision",
  "lore file", "lore auth", "lore layer", "lore link", "lore lock", "lore status",
  "lore clone", "lore commit", "lore sync", "lore push", "lore stage", "lore diff",
  "lore history", "lore cherry-pick", "lore revert", "lore bisect", "lore merge",
  "lore resolve", "lore dirty", "lore unstage", "lore reset", "lore obliterate",
  "lore service", "lore notification", "lore completions", "lore shared-store",
  "lore logfile", "lore login", "lore metadata", "lore dependency", "lore gc",
  "lore verify", "lore dump", "lore store", "lore instance", "lore config",
  "lore archive", "lore protect", "lore latest", "lore amend", "lore restore",
  "lore find", "lore hash", "lore write", "lore info", "immutable store",
  "split-write filesystem", "layer", "link", "shared store", "lore VCS".
user-invocable: true
---

# Lore CLI In-Depth Reference

A version control system from Epic Games for game development. Covers the full CLI command surface. Source: <https://epicgames.github.io/lore/reference/lore-cli-commands/>

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| core concepts, VCS model, revision graph, branch model, immutable store, split-write, layer, link, shared store, fragment, blob, hash, identity | [references/core-concepts.md](references/core-concepts.md) |
| repository, create, clone, delete, status, info, list, verify, dump, gc, store, immutable query, metadata, instance, config, update-path, remote | [references/repository.md](references/repository.md) |
| branch, list, info, create, switch, push, merge, unresolve, into, start, restart, resolve, mine, theirs, abort, diff, archive, reset, protect, unprotect, latest | [references/branch.md](references/branch.md) |
| revision, history, info, commit, amend, sync, bisect, start, end, diff, find, metadata, number, restore, cherry-pick, unresolve, restart, resolve, mine, theirs, abort, revert | [references/revision.md](references/revision.md) |
| file, info, metadata, dependency, add, remove, list, tag, stage, dirty, move, copy, unstage, reset, purge, obliterate, history, diff, write, hash, address | [references/file.md](references/file.md) |
| auth, login, logout, info, list, clear, token, layer, add, remove, list, link, add, remove, update, pin, logfile, lock, acquire, status, query, release | [references/auth-layer-link.md](references/auth-layer-link.md) |
| status, clone, stage, dirty, unstage, reset, diff, history, commit, sync, push, service, run, start, stop, notification, subscribe, completions, shared-store, create, info, set-use-automatically, shell, shortcuts | [references/ops.md](references/ops.md) |

## Quick Guide

- **How do I create a new repository?** → [references/repository.md](references/repository.md)
- **How do I clone a remote repository?** → [references/repository.md](references/repository.md)
- **How do I create/switch branches?** → [references/branch.md](references/branch.md)
- **How do I merge branches and resolve conflicts?** → [references/branch.md](references/branch.md)
- **How do I commit changes?** → [references/revision.md](references/revision.md)
- **How do I sync to a specific revision?** → [references/revision.md](references/revision.md)
- **How do I stage/unstage files?** → [references/file.md](references/file.md)
- **How do I view diffs and history?** → [references/file.md](references/file.md) and [references/revision.md](references/revision.md)
- **How do I cherry-pick or revert a revision?** → [references/revision.md](references/revision.md)
- **How do I authenticate with a remote server?** → [references/auth-layer-link.md](references/auth-layer-link.md)
- **How do I add/remove layers?** → [references/auth-layer-link.md](references/auth-layer-link.md)
- **How do I link repositories?** → [references/auth-layer-link.md](references/auth-layer-link.md)
- **How do I acquire/release file locks?** → [references/auth-layer-link.md](references/auth-layer-link.md)
- **How do I generate shell completions?** → [references/ops.md](references/ops.md)
- **How do I start/stop the service process?** → [references/ops.md](references/ops.md)
- **What is the conceptual model (revisions, branches, store)?** → [references/core-concepts.md](references/core-concepts.md)
- **How do I add file dependencies?** → [references/file.md](references/file.md)
- **How do I bisect revisions to find a change?** → [references/revision.md](references/revision.md)
- **How do I manage metadata on repos/branches/revisions/files?** → [references/repository.md](references/repository.md), [references/branch.md](references/branch.md), [references/revision.md](references/revision.md), [references/file.md](references/file.md)

## Cross-Project References

- For **git** internals (comparison mental model), see the **git** skill.
- For **jj** (another VCS with similar concepts), see the **jj** skill.