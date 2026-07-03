# Forge Integration — Push, Pull, and PRs

Commands for interacting with remote repositories and forges (GitHub).

> **Source of truth**: <https://docs.gitbutler.com/cli-guides/cli-tutorial/forges>

## Supported Forges

Currently supports **GitHub** (GitLab Merge Requests planned). Supports GitHub Enterprise via PAT or device flow.

## `but push` — Push Branch to Remote

Push changes in a branch to the remote repository.

```bash
but push <branch>
```

Positional argument is the branch name to push.

## `but pull` — Update from Upstream

Fetch and rebase active branches on new upstream changes.

```bash
but pull                # Fetch + rebase active branches
but pull --check        # Preview what pull would do
```

`but pull` fetches the latest from the remote and updates the merge base. All applied branches are rebased onto the new upstream.

## `but pr` — Pull Request Management

Parent command for PR operations.

### `but pr new`

Create a new pull request for a branch.

```bash
but pr new <branch>
but pr new --file <template-file> <branch>
```

| Flag | Short | Description |
|------|-------|-------------|
| `--file` | | PR template file to use |

PRs are created with the branch targeting the default branch (or the parent branch for stacked PRs). Stacked PRs maintain correct base branch metadata automatically.

### `but pr setDraft`

Set a PR to draft status.

```bash
but pr setDraft <branch-or-pr>
```

Accepts CLI IDs (branches, stacks) or local branch names, comma-separated.

### `but pr setReady`

Mark a PR as ready for review (remove draft status).

```bash
but pr setReady <branch-or-pr>
```

### `but pr autoMerge`

Enable auto-merge for a PR.

```bash
but pr autoMerge <branch-or-pr>
```

Accepts CLI IDs (branches, stacks) or local branch names, comma-separated.

### `but pr template`

Manage PR templates.

```bash
but pr template
```

Uses the git worktree as the base directory for template files.

## Authentication

### `but config forge` — Forge Configuration

View and manage forge (GitHub) settings.

```bash
but config forge              # View forge config
but config forge auth         # Authenticate to forge (device flow or PAT)
but config forge forget       # Forget a stored authentication
but config forge list-users   # List authenticated forge users
```

### Authentication Methods

| Method | Description |
|--------|-------------|
| Device flow (OAuth) | Browser-based OAuth flow, no token management |
| Personal Access Token (PAT) | Manual token entry |
| GitHub Enterprise | Custom server URL with PAT or device flow |

### Forge Users

`but config forge list-users` shows authenticated users with their auth type:

```
- OAuth: user1
- PAT: user2
```

## Stacked PRs

When working with stacked branches, GitButler automatically creates stacked PRs:

- The bottom branch's PR targets the default branch (e.g. `main`)
- Each stacked branch's PR targets its **parent branch**
- GitHub automatic branch deletion is highly recommended
- The "Merge" strategy on GitHub offers the best experience for stacked PRs

```bash
but pr new feature-part1     # PR targets main
but pr new feature-part2     # PR targets feature-part1 (stacked)
```

## Review Information

Several commands support `--review` to fetch and display review information:

```bash
but branch list --review      # Show PR/MR status for branches
but branch show --review       # Show review info for a specific branch
```

## Related

- [branches.md](branches.md) — Branch management including stacked branches
- [config.md](config.md) — Forge configuration commands
- [workspace.md](workspace.md) — Upstream and merge base concepts
