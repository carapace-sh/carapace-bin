# Workspace Model

The GitButler workspace model — how virtual branches, the integration branch, and the merge base enable parallel branch work in a single working directory.

## Overview

GitButler replaces Git's single-HEAD checkout model with a **workspace** where multiple branches are applied simultaneously. The `but` CLI operates on this workspace, managing branches, commits, and changes without requiring `git checkout` or worktrees.

> **Source of truth**: <https://docs.gitbutler.com/overview> and <https://docs.gitbutler.com/features/branch-management/integration-branch>

## The Integration Branch

GitButler maintains a special internal branch: `gitbutler/workspace`. This is a **merge commit** that GitButler rebuilds as you modify branches. It exists because stock Git has no concept of having several branches applied at once.

### How It Works

- When other tools run `git status` / `git diff` / `git log`, GitButler modifies your index to look like the **union of all committed states** of all applied parallel branches
- This makes normal Git commands behave mostly as expected
- The merge base (target branch, e.g. `origin/main`) is the common ancestor all virtual branches are based on

### Commit Protection

Git pre-commit and post-checkout hooks prevent committing directly to `gitbutler/workspace` — they tell you to use `but commit` instead. This prevents corrupting the workspace state.

### Setup and Teardown

```bash
but setup     # Initialize GitButler workspace in a git repository
but teardown  # Exit GitButler mode, return to vanilla Git workflow
```

`but setup` makes the necessary changes for GitButler to manage your data. `but teardown` undoes those changes and checks out a normal branch. Use `--checkout-to` to specify which branch to land on:

```bash
but teardown --checkout-to main
```

`but setup` also supports `--init` to initialize a new git repository with an empty commit if one doesn't exist.

## Virtual Branches (Parallel Branches)

Virtual branches are GitButler's core innovation — the ability to work on **multiple branches simultaneously** in a single working directory.

### Key Concepts

- Each virtual branch is kept in a vertical lane (like a Kanban board)
- Files and changes are like cards you can drag between lanes
- You don't "switch" branches (which implies replacement) — you **apply** or **unapply** branches
- Applying adds changes to your working directory; unapplying removes only those changes
- Each branch has its own staging area
- Since all branches start from changes in one working directory, they merge cleanly

### How Commits Work

When you commit, GitButler calculates what the branch would look like if only those changes existed, and commits that file tree. Pushed commits look like single-branch changes on GitHub — the virtual branch abstraction is invisible to remotes.

### Creating Virtual Branches

```bash
but branch new <name>              # Create a parallel branch
but apply <branch>                  # Apply a branch to the workspace
but unapply <branch>                # Remove a branch from the workspace
```

For stacked branches (dependent branches), see [branches.md](branches.md).

## Merge Base and Upstream

The **merge base** is the target branch (typically `origin/main`) that all virtual branches are based on. The workspace tracks how far behind the upstream the merge base is.

```bash
but status -u              # Show upstream info
but pull                   # Fetch + rebase active branches on new upstream
but pull --check           # Preview what pull would do
```

View or change the target branch:

```bash
but config target                   # Show current target branch
but config target origin/develop    # Change target branch
```

## Unassigned Changes

Changes in the working directory that haven't been assigned to a branch are called **unassigned changes** (labeled "uncommitted" in output since 0.20.4). The special ID `zz` refers to the unassigned area.

To assign changes to a branch, use `but stage` or `but rub`:

```bash
but stage <file> <branch>     # Assign file changes to a branch
but rub <file> <branch>       # Alias for staging via rub
but rub <file> zz             # Move changes back to unassigned
```

See [rubbing.md](rubbing.md) for the full `but rub` reference.

## Status Output

`but status` shows the workspace state: applied stacks, branches, commits, and unassigned changes. See [inspection.md](inspection.md) for details.

## Related

- [branches.md](branches.md) — Branch management commands
- [commits.md](commits.md) — Commit creation and editing
- [inspection.md](inspection.md) — Status, diff, and CLI IDs
