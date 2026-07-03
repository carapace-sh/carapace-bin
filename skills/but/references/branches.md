# Branch Management

Commands for managing virtual branches and stacked branches in the GitButler workspace.

> **Source of truth**: <https://docs.gitbutler.com/cli-guides/cli-tutorial/branching-and-commiting> and <https://docs.gitbutler.com/features/branch-management/virtual-branches>

## Command Groups

Branch commands are in the `branching and committing` group. The `branch` command is a parent with subcommands.

## `but branch` — Branch Subcommands

Parent command for branch management operations.

```bash
but branch <subcommand>
```

### `but branch list`

List branches in the repository. Shows applied and unapplied branches.

| Flag | Short | Description |
|------|-------|-------------|
| `--all` | `-a` | Show all branches (not just active + 20 most recent) |
| `--empty` | | Include branches with no commits (hidden by default) |
| `--local` | `-l` | Show only local branches |
| `--no-ahead` | | Don't calculate commits ahead of base (faster) |
| `--no-check` | | Don't check if each branch merges cleanly into upstream |
| `--remote` | `-r` | Show only remote branches |
| `--review` | | Fetch and display review information (PRs, MRs, etc.) |

### `but branch new`

Create a new branch in the workspace.

```bash
but branch new <name>                    # Create a parallel branch
but branch new --anchor <base> <name>    # Create a stacked branch
```

| Flag | Short | Description |
|------|-------|-------------|
| `--anchor` | `-a` | Anchor point — commit ID or branch name to create from |

The `--anchor` flag creates **stacked branches** — dependent branches where each is based on the previous one. See [Stacked Branches](#stacked-branches) below.

### `but branch show`

Show commits ahead of base for a specific branch.

```bash
but branch show <branch>
```

| Flag | Short | Description |
|------|-------|-------------|
| `--ai` | | Generate AI summary of the branch changes |
| `--check` | | Check if branch merges cleanly into upstream, identify conflicting commits |
| `--files` | `-f` | Show files modified in each commit with line counts |
| `--review` | `-r` | Fetch and display review information |

### `but branch update`

Update a local branch with the content of its remote counterpart.

```bash
but branch update <branch>
```

| Flag | Short | Description |
|------|-------|-------------|
| `--dry-run` | | Preview resulting branch state without persisting |
| `--interactive` | `-i` | Open the generated integration script in an editor |
| `--strategy` | `-s` | Strategy for integration (see below) |
| `--verbose` | `-v` | Show additional dry-run details like current divergence |

#### Update Strategies

| Strategy | Description |
|----------|-------------|
| `pull-rebase` | Rebuilds the branch picking first remote commits, then local commits (default) |
| `smart-squash` | Folds matching remote work into related local commits via Change IDs, falls back to pull-rebase |
| `merge` | Keeps local history and merges the remote tip into it |
| `pick-remote` | Rebuilds the branch picking only the commits on the remote |

### `but branch delete`

Delete a branch from the workspace.

```bash
but branch delete <branch>
```

## `but apply` — Apply a Branch

Apply a branch to the workspace. Adds the branch's changes to your working directory.

```bash
but apply <branch>
```

Accepts local or remote branch names. The branch becomes "applied" and its changes are visible in the working directory.

## `but unapply` — Unapply a Branch

Remove a branch from the workspace. Only removes that branch's changes, leaving other applied branches intact.

```bash
but unapply <branch>
```

## `but clean` — Remove Empty Branches

Remove empty branches from the workspace.

```bash
but clean
but clean --dry-run              # Preview which branches would be removed
but clean --include-upstream      # Also remove branches with only upstream commits
but clean --pull                  # Pull latest changes before cleaning
```

| Flag | Description |
|------|-------------|
| `--dry-run` | Preview which branches would be removed |
| `--include-upstream` | Also remove branches with upstream-only commits but no local commits |
| `--pull` | Pull latest changes from the remote before cleaning |

## Stacked Branches

Stacked branches are ordered, **dependent** branch sequences where each branch is based on the previous one. They allow reviewing and merging changes in sequence.

### Creating Stacks

```bash
but branch new -a <base-branch> <new-branch>   # Create a stacked branch
```

The `--anchor` (`-a`) flag specifies the base. Each stacked branch depends on the one below it.

### Stack Structure

```
main (merge base)
  └── feature-part1 (bottom of stack, PR targets main)
        └── feature-part2 (PR targets feature-part1)
              └── feature-part3 (PR targets feature-part2)
```

### Stacked PRs

When you run `but pr` on a stacked branch, GitButler creates a PR that targets the **parent branch** (not `main`). The bottom branch targets the default branch. GitHub automatic branch deletion is highly recommended for stacked PR workflow.

### Parallel vs Stacked

- **Parallel branches** are independent — no dependency between them
- **Stacked branches** depend on the ones before them
- They are complementary, not mutually exclusive — a workspace can have both

## Related

- [workspace.md](workspace.md) — Workspace model and virtual branch concepts
- [commits.md](commits.md) — Committing changes to branches
- [forge.md](forge.md) — Push and PR creation
