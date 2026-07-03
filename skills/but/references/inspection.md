# Inspection — Status, Diff, and CLI IDs

Commands for inspecting workspace state and the CLI ID system that identifies changes, commits, and branches.

> **Source of truth**: <https://docs.gitbutler.com/cli-guides/cli-tutorial/inspecting>

## `but status` — Workspace Overview

Show the project workspace state: applied stacks, branches, commits, and unassigned changes.

```bash
but status                # Basic overview
but status -f             # With file IDs (CLI IDs) in commits
but status -u             # With upstream info
but status -v             # Verbose (author and timestamp)
but status -r             # Refresh PRs from forge first
```

| Flag | Short | Description |
|------|-------|-------------|
| `--files` | | Show committed files (hidden, use `-f` instead) |
| `-f` | `-f` | Show committed files with CLI IDs |
| `--no-hint` | | Disable command hints at end of output |
| `--refresh-prs` | `-r` | Sync pull requests from forge before showing status |
| `--upstream` | `-u` | Show upstream commits not yet integrated |
| `--verbose` | `-v` | Show commit author and timestamp |

### Status Output Structure

The status output shows:
- **Stacks** — Applied stacks with their branches and commits
- **Unassigned changes** — Changes not assigned to any branch (labeled "uncommitted" since 0.20.4)
- **Upstream state** — How far behind the merge base is (with `-u`)

### JSON Output

```bash
but --format json status
but --format json status -f    # With committed file details
```

The JSON structure contains:
- `unassignedChanges` — changes not assigned to a branch
- `stacks` — applied stacks, each containing branches and their commits
- `mergeBase` — the target branch commit
- `upstreamState` — how far behind the merge base is (`behind`, `latestCommit`, `lastFetched`)

Each stack contains branches, each branch contains commits, and each commit contains changes.

## `but diff` — Show Diffs

Display the diff of changes in the repo. More readable than `git diff` for GitButler's multi-branch workspace.

```bash
but diff                       # All uncommitted changes
but diff <id>                  # Diff a specific file/commit by CLI ID
but diff --tui                 # Open diff in TUI viewer
```

Positional arguments are CLI IDs or targets — can be changes, commits, or branches.

## `but show` — Show Details

Show detailed information about a commit or branch.

```bash
but show <commit-or-branch>
but show -v <commit>           # Full commit messages and files changed
```

| Flag | Short | Description |
|------|-------|-------------|
| `--verbose` | `-v` | Show full commit messages and files changed for each commit |

Positional arguments are targets (commits or branches by CLI ID or name).

## CLI ID System

CLI IDs are short, stable identifiers used throughout the `but` CLI to refer to changes, commits, branches, and stacks. They appear in `but status -f` output and JSON as the `cliId` field.

### ID Format

CLI IDs are multi-part identifiers separated by `:`. Examples:

| ID Pattern | Refers To |
|------------|-----------|
| `h0` | A hunk (change) |
| `i0` | Another change |
| `c3` | A commit |
| `b1` | A branch |
| `zz` | The unassigned area (special) |

### ID Types

CLI IDs identify different entity types:

| Type | Description |
|------|-------------|
| Changes | Uncommitted file/hunk changes (assigned and unassigned) |
| Branches | Branch CLI IDs |
| Commits | Commit CLI IDs |
| Stacks | Stack CLI IDs (including `zz` for unassigned) |

### Change Types

Each change has a `changeType` field:

| Type | Abbreviation | Style |
|------|-------------|-------|
| `added` | `A` | green (positive) |
| `removed` | `D` | red (negative) |
| `modified` | `M` | yellow (ambiguous) |
| `renamed` | `R` | dim |

### Hunk IDs

Hunk IDs are stable across commits and discards (since 0.20.4). This means the same hunk retains its ID even after operations that move it between commits or discard it.

## Output Formats

The `--format` global flag controls output format:

| Format | Description |
|--------|-------------|
| `human` | Default human-readable output |
| `agent` | Formatted for AI agent consumption |
| `shell` | Shell-friendly output |
| `json` | Machine-parseable JSON |
| `none` | No output |

```bash
but --format json status     # JSON status
but --format agent status    # Agent-optimized status
```

## Related

- [workspace.md](workspace.md) — Workspace model and concepts
- [commits.md](commits.md) — Commit commands that use CLI IDs
- [rubbing.md](rubbing.md) — Rub command using CLI IDs
