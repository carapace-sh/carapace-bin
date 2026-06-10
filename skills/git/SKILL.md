---
name: git
description: >
  Use when working with Git — the distributed version control system. Covers the object model,
  references, the index/staging area, branching and merging, rebasing and history rewriting,
  remotes and transfer protocols, diffs and conflict resolution, configuration, hooks,
  worktrees and submodules, plumbing commands, packfiles, and common workflows.
  Triggers on: "git", "git object", "git blob", "git tree", "git commit object", "git tag",
  "git branch", "git ref", "git HEAD", "git index", "git staging area", "git staging",
  "git merge", "git rebase", "git cherry-pick", "git remote", "git fetch", "git push",
  "git pull", "git diff", "git conflict", "git config", "git hook", "git worktree",
  "git submodule", "git packfile", "git plumbing", "git porcelain", "git reflog",
  "git stash", "git bisect", "git blame", "git reset", "git checkout", "git switch",
  "git restore", "git sparse-checkout", "git transfer protocol", "git pack protocol",
  "git shallow clone", "git bare repo", "git detached HEAD", "git fast-forward",
  "git merge commit", "git squash merge", "git orphan branch", "git refspec",
  "git upstream", "git tracking branch", "git OID", "git hash", "git SHA-1",
  "git object store", "git loose object", "git delta", "git thin pack",
  "git alternates", "git replace", "git graft", "git notes", "git attribute",
  "gitignore", "git LFS", "git filter-branch", "git filter-repo".
user-invocable: true
---

# Git — Distributed Version Control In-Depth Reference

Comprehensive reference for Git internals, concepts, and workflows. Covers the object model, reference system, index, branching/merging, remotes, configuration, hooks, and plumbing commands.

## Data Flow

```
Working Tree
  ↕ (git add / git checkout)
Index (Staging Area)
  ↕ (git commit / git reset)
Object Store (.git/objects/)
  ↕ (git pack / git gc)
Packfiles (.git/objects/pack/)
  ↕ (git fetch / git push)
Remote Object Store
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

|| Keywords | Reference |
||----------|----------|
|| concept, model, DAG, directed acyclic graph, snapshot, content-addressable, integrity, hash, SHA-1, SHA-256, OID, object ID, three states, working tree, index, repository, bare repo, .git, git init, git clone, detached HEAD, HEAD, symbolic ref, plumbing vs porcelain, high-level vs low-level | [references/concepts.md](references/concepts.md) |
|| object, blob, tree, commit, tag object, annotated tag, lightweight tag, object store, loose object, packfile, pack index, delta, OFS_DELTA, REF_DELTA, thin pack, object type, object format, hash algorithm, object ID, content-addressable storage, git hash-object, git cat-file, git mktree, git mktag, object graph, reachability, garbage collection, git gc, git prune, git fsck | [references/objects.md](references/objects.md) |
|| ref, reference, branch, HEAD, symbolic ref, reflog, remote-tracking branch, tracking branch, refspec, tag, annotated tag, lightweight tag, namespace, refs/heads, refs/tags, refs/remotes, refs/replace, refs/notes, packed-refs, git update-ref, git symbolic-ref, git for-each-ref, git show-ref, git branch, git tag, git remote, upstream, tracking, fast-forward, ff-only | [references/refs.md](references/refs.md) |
|| index, staging area, cache, .git/index, git add, git reset, git rm, git status, staged, unstaged, untracked, index format, extensions, skip-worktree, assume-unchanged, intent-to-add, sparse-checkout, git ls-files, git update-index, git diff-index, git read-tree, git write-tree, DIRC, index version | [references/index.md](references/index.md) |
|| cli, command, subcommand, flag, porcelain, git log, git show, git diff, git status, git add, git commit, git reset, git checkout, git switch, git restore, git stash, git bisect, git blame, git grep, git shortlog, git cherry, git describe, git rerere, git notes, git clean, git mv, git rm, git range-diff, git bugreport, git diagnose, git maintenance | [references/cli.md](references/cli.md) |
|| branch, merge, fast-forward, merge commit, squash merge, octopus merge, recursive merge, ort strategy, resolve strategy, ours strategy, merge base, git merge-base, git merge, git branch, conflict, merge driver, .gitattributes merge, merge conflict markers, three-way merge, base commit, first-parent merge, no-ff, ff-only, orphan branch, git merge-file | [references/branching-merging.md](references/branching-merging.md) |
|| remote, fetch, push, pull, refspec, remote-tracking branch, tracking branch, upstream, git remote, git fetch, git push, git pull, transfer protocol, pack protocol, dumb protocol, smart protocol, SSH, HTTPS, git://, file://, git daemon, git http-backend, credential helper, git credential, proxy, mirror, bare push, force push, lease, push options, fetch options, prune, git ls-remote, git archive | [references/remotes.md](references/remotes.md) |
|| rebase, cherry-pick, git rebase, git cherry-pick, interactive rebase, rebase -i, fixup, squash, reword, edit, drop, rebase onto, rebase abort, rebase continue, rebase skip, rebase --autosquash, rebase --update-refs, rebase --rebase-merges, history rewriting, filter-branch, filter-repo, git revert, git reset, git reflog, recover lost commit, git commit --amend, force push, golden rule, rebase vs merge | [references/rebase.md](references/rebase.md) |
|| diff, patch, conflict, conflict resolution, conflict marker, git diff, git diff-index, git diff-tree, git diff-files, git apply, git am, git format-patch, git rerere, merge driver, custom merge driver, .gitattributes, diff algorithm, histogram, patience, myers, word diff, stat, numstat, patch-id, git patch-id, binary diff, git apply --3way, reject file, hunk, context line | [references/diff-patch.md](references/diff-patch.md) |
|| config, configuration, .gitconfig, .git/config, git config, config level, system, global, local, worktree, include, includeIf, conditional include, config key, config value, config section, config subsection, config list, config edit, config get, config unset, environment variable, GIT_CONFIG, GIT_CONFIG_GLOBAL, GIT_CONFIG_SYSTEM, credential, signing, merge tool, diff tool, core, user, init, alias, multiple values, config precedence | [references/config.md](references/config.md) |
|| internals, plumbing, git cat-file, git hash-object, git mktree, git mktag, git update-ref, git write-tree, git commit-tree, git read-tree, git update-index, git unpack-objects, git pack-objects, git index-pack, git verify-pack, git rev-parse, git ls-tree, git ls-files, git for-each-ref, git for-each-repo, git name-rev, git interpret-triangular, git check-attr, git check-ignore, git check-mailmap, git merge-base, git rev-list, git show-index, .git directory layout, gitfile, alternates, replace objects, grafts, env variables, GIT_DIR, GIT_WORK_TREE, GIT_INDEX_FILE, GIT_OBJECT_DIRECTORY, ALTERNATE_DB | [references/internals.md](references/internals.md) |
|| worktree, git worktree, linked worktree, pruning worktrees, submodule, git submodule, git submodule update, git submodule init, git submodule sync, git submodule absorbgitdirs, .gitmodules, subtree merge, subtree, git subtree, monorepo, sparse-checkout, git sparse-checkout, cone mode, non-cone mode, partial clone, filter, blobless clone, treeless clone, shallow clone, depth, git clone --depth, single-branch, promisor, lazy fetch | [references/worktree.md](references/worktree.md) |
|| hook, git hook, pre-commit, pre-push, pre-rebase, prepare-commit-msg, commit-msg, post-commit, post-merge, post-checkout, post-rewrite, applypatch-msg, pre-applypatch, update, pre-receive, update, post-receive, post-update, push-to-checkout, reference-transaction, fsmonitor-watchman, p4-changelist, p4-prepare-changelist, p4-post-changelist, sendemail-validate, hook directory, core.hooksPath, hook ordering, hook exit code, hook environment, hook stdin, hook chaining, client-side hook, server-side hook, husky, lefthook | [references/hooks.md](references/hooks.md) |
|| workflow, git flow, github flow, trunk-based development, feature branch, release branch, hotfix, pull request, code review, conventional commits, semantic versioning, git bisect workflow, git stash workflow, gitignore patterns, .gitignore, LFS, git lfs, lock, git lfs lock, git lfs track, git lfs fetch, git lfs pull, git lfs push, git lfs prune, smudge, clean, attribute, .gitattributes, signing, GPG, SSH sign, commit signing, tag signing | [references/workflows.md](references/workflows.md) |

## Quick Guide

- **How do Git objects work?** → [references/objects.md](references/objects.md)
- **How do branches and tags work internally?** → [references/refs.md](references/refs.md)
- **What is the index/staging area?** → [references/index.md](references/index.md)
- **What are Git's core concepts?** → [references/concepts.md](references/concepts.md)
- **How do I use git commands?** → [references/cli.md](references/cli.md)
- **How does merging work?** → [references/branching-merging.md](references/branching-merging.md)
- **How do remotes and fetch/push work?** → [references/remotes.md](references/remotes.md)
- **How does rebase work?** → [references/rebase.md](references/rebase.md)
- **How do diffs and conflicts work?** → [references/diff-patch.md](references/diff-patch.md)
- **How does git config work?** → [references/config.md](references/config.md)
- **What are the plumbing commands?** → [references/internals.md](references/internals.md)
- **How do worktrees and submodules work?** → [references/worktree.md](references/worktree.md)
- **How do git hooks work?** → [references/hooks.md](references/hooks.md)
- **What are common Git workflows?** → [references/workflows.md](references/workflows.md)
- **How do I recover a lost commit?** → [references/rebase.md](references/rebase.md) and [references/refs.md](references/refs.md)
- **How does a shallow clone work?** → [references/worktree.md](references/worktree.md)
- **How does the pack protocol work?** → [references/remotes.md](references/remotes.md) and [references/objects.md](references/objects.md)
- **What is a detached HEAD?** → [references/concepts.md](references/concepts.md) and [references/refs.md](references/refs.md)
- **How do I write a custom merge driver?** → [references/diff-patch.md](references/diff-patch.md)
- **How do I configure git hooks?** → [references/hooks.md](references/hooks.md)

## Cross-Project References

- For jj (Jujutsu VCS) concepts and commands, use the **jj** skill (in the carapace-jjlex repo).
- For carapace-specific git integration (completion actions, macros like `tools.git.Changes`), see the **carapace-dev** skill → `references/action.md`.
- For bash shell internals that git hooks run in (traps, signals, execution model), see the **bash** skill → `references/execution.md`.
