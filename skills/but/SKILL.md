---
name: but
description: >
  Use when working with the GitButler CLI (`but`) — virtual branches, stacked branches,
  workspace model, commit editing (rub, amend, squash, reword), CLI IDs, forge integration,
  PRs, operations log, conflict resolution, TUI, AI agent setup, skills, aliases, configuration.
  Triggers on: "but", "gitbutler", "but rub", "but commit", "but status", "but branch",
  "but apply", "but unapply", "but push", "but pull", "but pr", "but oplog", "but undo",
  "but redo", "but resolve", "but stage", "but squash", "but amend", "but reword",
  "but config", "but alias", "but setup", "but teardown", "but tui", "but diff",
  "but show", "but agent", "but skill", "but clean", "but move", "but pick", "but absorb",
  "but update", "but completions", "virtual branch", "stacked branch", "integration branch",
  "gitbutler workspace", "but oplog restore", "but oplog snapshot", "CLI ID", "hunk ID".
user-invocable: true
---

# GitButler CLI (`but`) In-Depth Reference

Reference for the [GitButler CLI](https://docs.gitbutler.com) — a stacked-branch Git workflow tool that enables virtual branches, parallel branch work, commit editing, and AI agent integration from the terminal.

## Data Flow

```
Working directory changes
  → but stage (assign changes to branches)
    → but commit (create commits on virtual branches)
      → but push (push branches to remote)
        → but pr (create/update pull requests on forge)
```

```
but status
  → but diff (inspect changes)
    → but show (inspect commits/branches)
      → but rub (edit: amend, squash, move, split)
        → but oplog (review history) / but undo (revert)
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| workspace, integration branch, virtual branches, parallel branches, gitbutler/workspace, setup, teardown, merge base, upstream, target branch | [references/workspace.md](references/workspace.md) |
| branch, branch new, branch delete, branch list, branch show, branch update, apply, unapply, clean, anchor, stacked branch, stack, strategy | [references/branches.md](references/branches.md) |
| commit, commit empty, stage, amend, squash, reword, move, pick, absorb, discard, commit message, --ai, --changes, --only, hunk, --after, --before | [references/commits.md](references/commits.md) |
| rub, rubbing, amend via rub, squash via rub, move via rub, split commits, uncommit via rub, stage via rub, assign changes, zz, unassigned | [references/rubbing.md](references/rubbing.md) |
| status, diff, show, CLI ID, cliId, hunk ID, change ID, change type, but status --json, but diff --json, format, --format, files, verbose, upstream | [references/inspection.md](references/inspection.md) |
| push, pull, pull --check, pr, pr new, pr setDraft, pr setReady, pr autoMerge, pr template, forge, GitHub, authentication, auth, list-users, remote, review | [references/forge.md](references/forge.md) |
| oplog, oplog list, oplog restore, oplog snapshot, undo, redo, snapshot, operation history, CreateCommit, AmendCommit, SquashCommit, RestoreFromSnapshot | [references/oplog.md](references/oplog.md) |
| resolve, resolve finish, resolve cancel, resolve status, conflict, conflict resolution, zdiff3, conflicted commit, rebase conflict | [references/conflicts.md](references/conflicts.md) |
| config, config user, config target, config forge, config metrics, config ui, config ai, anthropic, openai, ollama, lmstudio, openrouter, alias, alias add, alias list, alias remove, editor, name, email | [references/config.md](references/config.md) |
| tui, interactive TUI, keybindings, navigation, diff TUI, stage hunk picker, reorder stacks, apply from TUI, move commits, copy, scroll | [references/tui.md](references/tui.md) |
| agent, agent setup, skill, skill install, skill check, AI, --ai, coding agent, steering instructions, parallel agents, tuning agent behavior, checkpoint commits, stacked PRs, ship it | [references/ai-agents.md](references/ai-agents.md) |
| root command, global flags, --current-dir, --format, --version, completions, update, update check, update install, update suppress, command groups | [references/misc.md](references/misc.md) |

## Quick Guide

- **What is the GitButler workspace model?** → [references/workspace.md](references/workspace.md)
- **How do virtual branches work?** → [references/workspace.md](references/workspace.md)
- **How do I create a stacked branch?** → [references/branches.md](references/branches.md)
- **How do I commit changes to a specific branch?** → [references/commits.md](references/commits.md)
- **What is `but rub` and how does it work?** → [references/rubbing.md](references/rubbing.md)
- **How do CLI IDs work?** → [references/inspection.md](references/inspection.md)
- **How do I push and create PRs?** → [references/forge.md](references/forge.md)
- **How do I undo an operation?** → [references/oplog.md](references/oplog.md)
- **How do I resolve conflicts?** → [references/conflicts.md](references/conflicts.md)
- **How do I configure GitButler?** → [references/config.md](references/config.md)
- **What can the TUI do?** → [references/tui.md](references/tui.md)
- **How do I set up GitButler for AI coding agents?** → [references/ai-agents.md](references/ai-agents.md)
- **What are the global flags?** → [references/misc.md](references/misc.md)

## Cross-Project References

- For **Git internals** (branches, refs, remotes, changes, tags), use the **git** skill.
- For **carapace** completion integration (specs, actions, macros), use the **carapace** skill.
