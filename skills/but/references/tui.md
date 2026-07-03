# TUI — Interactive Terminal UI

The GitButler TUI provides a keyboard-first, interactive terminal interface for the full workspace.

> **Source of truth**: <https://docs.gitbutler.com/gitbutler-tui>

## `but tui` — Launch the TUI

```bash
but tui
```

Opens an interactive, full-terminal workspace view showing unassigned changes, branch lanes, commits, and file-level changes.

## Capabilities

The TUI provides access to most GitButler operations without typing commands:

- Inspect unassigned changes, branch lanes, commits, file-level changes
- Open details pane for branches, commits, files, or hunks
- Create commits and empty commits
- Create branches
- Move commits or branches (including multiple commits at once since 0.20.4)
- Apply branches from the TUI (since 0.20.4)
- Reorder stacks (since 0.20.4)
- Rub changes into a branch or commit
- Reword commits or branch names
- Discard local work
- Undo and redo operations
- Commit above or below the selected commit (since 0.20.4)

## Keybindings

### Navigation

| Key | Action |
|-----|--------|
| `j` / `k` / arrows | Move by line |
| `ctrl+d` / `ctrl+u` | Jump by 10 lines |
| `shift+j` / `shift+k` | Jump by section |
| `g` | Go to top |
| `shift+g` | Go to bottom |

### Actions

| Key | Action |
|-----|--------|
| `d` | Open details pane |
| `c` | Create commit |
| `r` | Rub (combine entities) |
| `u` | Undo |
| `shift+u` | Redo |
| `:` | Command mode |
| `!` | External command |
| `shift+y` | Copy (commit details, diff hunks, branch names, etc.) |

### Confirmation Behavior (since 0.20.4)

Unapplying a branch and undoing no longer prompt for confirmation — both can be undone via the oplog.

## Diff TUI

```bash
but diff --tui
```

Opens a focused diff viewer for inspecting changes interactively.

## Stage Hunk Picker

```bash
but stage
```

Running `but stage` without arguments opens the interactive hunk picker, where you can select specific hunks to stage to a branch.

## Related

- [commits.md](commits.md) — Commit commands available from the TUI
- [rubbing.md](rubbing.md) — Rub operations available from the TUI
- [inspection.md](inspection.md) — Status and diff commands
