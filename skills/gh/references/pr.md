# Pull Requests

How to create, view, edit, merge, review, and manage pull requests with `gh pr`.

## Overview

`gh pr` manages GitHub pull requests. A PR can be referenced by number, URL, or head branch name. The `-R, --repo [HOST/]OWNER/REPO` flag targets a specific repository.

## Argument Format

A pull request can be supplied as:
- By number — e.g. `123`
- By URL — e.g. `https://github.com/OWNER/REPO/pull/123`
- By head branch name — e.g. `patch-1` or `OWNER:patch-1`

## Commands

### `gh pr create`

Create a pull request on GitHub. Aliased as `gh pr new`. Upon success, the PR URL is printed.

```sh
$ gh pr create --title "The bug is fixed" --body "Everything works again"
$ gh pr create --reviewer monalisa,hubot --reviewer myorg/team-name
$ gh pr create --project "Roadmap"
$ gh pr create --base develop --head monalisa:feature
$ gh pr create --fill                         # autofill from commits
$ gh pr create --template "pull_request_template.md"
$ gh pr create --dry-run                      # print without creating
```

| Flag | Description |
|------|-------------|
| `-a, --assignee login` | Assign people by login. `@me` = self-assign |
| `-B, --base branch` | Branch to merge into |
| `-b, --body string` | PR body |
| `-F, --body-file file` | Read body from file (`-` for stdin) |
| `-d, --draft` | Mark as draft |
| `--dry-run` | Print details instead of creating (may still push) |
| `-e, --editor` | Open editor for title and body |
| `-f, --fill` | Use commit info for title and body |
| `--fill-first` | Use first commit info for title and body |
| `--fill-verbose` | Use commits msg+body for description |
| `-H, --head branch` | Branch containing commits (default: current branch). Supports `<user>:<branch>` |
| `-l, --label name` | Add labels by name |
| `-m, --milestone name` | Add to milestone by name |
| `--no-maintainer-edit` | Disable maintainer's ability to modify PR |
| `-p, --project title` | Add to projects by title |
| `--recover string` | Recover input from a failed run |
| `-r, --reviewer handle` | Request reviews from people or teams |
| `-T, --template file` | Template file for starting body text |
| `-t, --title string` | PR title |
| `-w, --web` | Open browser to create PR |

**`--head` syntax**: `<user>:<branch>` selects a head repo owned by `<user>`. Using an organization as `<user>` is not currently supported.

**`--fill` vs `--fill-first` vs `--fill-verbose`**: `--fill` uses the latest commit info; `--fill-first` uses the first commit; `--fill-verbose` uses commit messages and bodies. If `--title`/`--body` are provided alongside `--fill`, the explicit values take precedence.

**Base branch**: If `--base` is not provided, gh uses the `gh-merge-base` git branch config, then the repository's default branch. Configure with: `git config branch.{current}.gh-merge-base {base}`.

**Project scope**: Adding a PR to projects requires the `project` scope. Run `gh auth refresh -s project`.

### `gh pr list`

List pull requests in a repository. Aliased as `gh pr ls`. Defaults to open PRs.

```sh
$ gh pr list --author "@me"
$ gh pr list --head "typo"
$ gh pr list --label bug --label "priority 1"
$ gh pr list --search "status:success review:required"
$ gh pr list --search "<SHA>" --state merged  # find PR that introduced a commit
```

| Flag | Description |
|------|-------------|
| `--app string` | Filter by GitHub App author |
| `-a, --assignee string` | Filter by assignee |
| `-A, --author string` | Filter by author |
| `-B, --base string` | Filter by base branch |
| `-d, --draft` | Filter by draft state |
| `-H, --head string` | Filter by head branch (`<owner>:<branch>` not supported) |
| `-l, --label strings` | Filter by label |
| `-L, --limit int` | Max items (default 30) |
| `-S, --search query` | Search with query syntax |
| `-s, --state string` | Filter by state: `open`, `closed`, `merged`, `all` (default `open`) |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `additions`, `assignees`, `author`, `autoMergeRequest`, `baseRefName`, `baseRefOid`, `body`, `changedFiles`, `closed`, `closedAt`, `closingIssuesReferences`, `comments`, `commits`, `createdAt`, `deletions`, `files`, `fullDatabaseId`, `headRefName`, `headRefOid`, `headRepository`, `headRepositoryOwner`, `id`, `isCrossRepository`, `isDraft`, `labels`, `latestReviews`, `maintainerCanModify`, `mergeCommit`, `mergeStateStatus`, `mergeable`, `mergedAt`, `mergedBy`, `milestone`, `number`, `potentialMergeCommit`, `projectCards`, `projectItems`, `reactionGroups`, `reviewDecision`, `reviewRequests`, `reviews`, `state`, `statusCheckRollup`, `title`, `updatedAt`, `url`.

### `gh pr view`

Display PR title, body, and other information. Without an argument, shows the PR for the current branch.

```sh
$ gh pr view
$ gh pr view 123
$ gh pr view --comments    # view comments
$ gh pr view --web         # open in browser
```

| Flag | Description |
|------|-------------|
| `-c, --comments` | View PR comments |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output (same fields as `pr list`) |

### `gh pr checkout`

Check out a pull request in git. Aliased as `gh pr co`.

```sh
$ gh pr checkout           # interactively select from 10 most recent
$ gh pr checkout 32
$ gh pr checkout https://github.com/OWNER/REPO/pull/32
$ gh pr checkout feature
```

| Flag | Description |
|------|-------------|
| `-b, --branch string` | Local branch name (default: head branch name) |
| `--detach` | Checkout with detached HEAD |
| `-f, --force` | Reset existing local branch to latest PR state |
| `--recurse-submodules` | Update submodules after checkout |

### `gh pr edit`

Edit a pull request. Without an argument, selects the PR for the current branch.

```sh
$ gh pr edit 23 --title "I found a bug" --body "Nothing works"
$ gh pr edit 23 --body-file body.txt
$ gh pr edit 23 --add-label "bug,help wanted" --remove-label "core"
$ gh pr edit 23 --add-reviewer monalisa,hubot --remove-reviewer myorg/team-name
$ gh pr edit 23 --add-reviewer "@copilot"     # request Copilot review
$ gh pr edit 23 --add-assignee "@me" --remove-assignee monalisa,hubot
$ gh pr edit 23 --add-project "Roadmap" --remove-project v1,v2
$ gh pr edit 23 --milestone "Version 1"
$ gh pr edit 23 --remove-milestone
```

| Flag | Description |
|------|-------------|
| `--add-assignee login` | Add assignees (`@me`, `@copilot`) |
| `--add-label name` | Add labels |
| `--add-project title` | Add to projects |
| `--add-reviewer login` | Add/re-request reviewers (`@copilot`) |
| `-B, --base branch` | Change base branch |
| `-b, --body string` | Set new body |
| `-F, --body-file file` | Read body from file |
| `-m, --milestone name` | Set milestone |
| `--remove-assignee login` | Remove assignees |
| `--remove-label name` | Remove labels |
| `--remove-milestone` | Remove milestone |
| `--remove-project title` | Remove from projects |
| `--remove-reviewer login` | Remove reviewers |
| `-t, --title string` | Set new title |

**Special values for `--add-assignee`/`--remove-assignee`**: `@me` (self), `@copilot` (Copilot; not supported on GHES).

**Special values for `--add-reviewer`/`--remove-reviewer`**: `@copilot` (Copilot; not supported on GHES).

**Re-requesting review**: Use `--add-reviewer` with an already-requested reviewer to re-request.

**Project scope**: Editing PR projects requires the `project` scope (`gh auth refresh -s project`).

### `gh pr merge`

Merge a pull request. Without an argument, selects the PR for the current branch.

```sh
$ gh pr merge --squash --delete-branch
$ gh pr merge --merge
$ gh pr merge --rebase
$ gh pr merge --auto              # auto-merge after requirements met
$ gh pr merge --admin             # bypass requirements with admin privileges
```

| Flag | Description |
|------|-------------|
| `--admin` | Use admin privileges to merge without meeting requirements |
| `-A, --author-email text` | Email for merge commit author |
| `--auto` | Auto-merge after necessary requirements are met |
| `-b, --body text` | Body for merge commit |
| `-F, --body-file file` | Read body from file |
| `-d, --delete-branch` | Delete local and remote branch after merge |
| `--disable-auto` | Disable auto-merge for this PR |
| `--match-head-commit SHA` | PR head must match this SHA to allow merge |
| `-m, --merge` | Merge commits with base branch |
| `-r, --rebase` | Rebase commits onto base branch |
| `-s, --squash` | Squash commits into one and merge |
| `-t, --subject text` | Subject for merge commit |

**Merge queue**: When targeting a branch that requires a merge queue, no merge strategy is required. If checks haven't passed, auto-merge is enabled. If checks have passed, the PR is added to the queue. Use `--admin` to bypass the queue and merge directly.

### `gh pr close`

Close a pull request.

```sh
$ gh pr close 123
$ gh pr close 123 --comment "Closing this PR"
$ gh pr close 123 --delete-branch
```

| Flag | Description |
|------|-------------|
| `-c, --comment string` | Leave a closing comment |
| `-d, --delete-branch` | Delete local and remote branch after close |

### `gh pr reopen`

Reopen a pull request.

```sh
$ gh pr reopen 123
$ gh pr reopen 123 --comment "Reopening this"
```

| Flag | Description |
|------|-------------|
| `-c, --comment string` | Add a reopening comment |

### `gh pr review`

Add a review to a pull request. Without an argument, reviews the PR for the current branch.

```sh
$ gh pr review --approve
$ gh pr review --comment -b "interesting"
$ gh pr review 123 -r -b "needs more ASCII art"
```

| Flag | Description |
|------|-------------|
| `-a, --approve` | Approve pull request |
| `-b, --body string` | Review body |
| `-F, --body-file file` | Read body from file |
| `-c, --comment` | Comment on pull request |
| `-r, --request-changes` | Request changes |

### `gh pr ready`

Mark a PR as ready for review. Without an argument, selects the PR for the current branch. Use `--undo` to convert back to draft.

```sh
$ gh pr ready
$ gh pr ready --undo    # convert to draft
```

| Flag | Description |
|------|-------------|
| `--undo` | Convert to draft |

### `gh pr diff`

View changes in a pull request. Without an argument, selects the PR for the current branch.

```sh
$ gh pr diff
$ gh pr diff 123
$ gh pr diff --exclude '*.yml' --exclude 'generated/*'
$ gh pr diff --name-only --exclude '*.generated.*'
```

| Flag | Description |
|------|-------------|
| `--color string` | Color output: `always`, `never`, `auto` (default `auto`) |
| `-e, --exclude patterns` | Exclude files matching glob patterns |
| `--name-only` | Show only names of changed files |
| `--patch` | Display in patch format |
| `-w, --web` | Open diff in browser |

The `--exclude` pattern uses forward slashes as path separators on all platforms. Repeat the flag for multiple patterns.

### `gh pr comment`

Add a comment to a pull request.

```sh
$ gh pr comment 13 --body "Hi from GitHub CLI"
$ gh pr comment 13 --body-file comment.md
$ gh pr comment 13 --edit-last              # edit your last comment
$ gh pr comment 13 --edit-last --create-if-none
$ gh pr comment 13 --delete-last
$ gh pr comment 13 --delete-last --yes      # skip confirmation
```

| Flag | Description |
|------|-------------|
| `-b, --body text` | Comment body |
| `-F, --body-file file` | Read body from file |
| `--create-if-none` | Create new comment if none found (only with `--edit-last`) |
| `--delete-last` | Delete your last comment |
| `--edit-last` | Edit your last comment |
| `-e, --editor` | Open editor for body |
| `-w, --web` | Open browser to write comment |
| `--yes` | Skip delete confirmation (with `--delete-last`) |

### `gh pr status`

Show status of relevant pull requests (assigned to you, requesting your review, etc.).

```sh
$ gh pr status
$ gh pr status --conflict-status
```

| Flag | Description |
|------|-------------|
| `-c, --conflict-status` | Display merge conflict status |
| `--json` / `--jq` / `--template` | JSON output |

### `gh pr checks`

Show CI status for a single pull request. Without an argument, selects the PR for the current branch.

```sh
$ gh pr checks
$ gh pr checks --watch              # watch until checks finish
$ gh pr checks --required           # only show required checks
$ gh pr checks --fail-fast          # exit on first failure (watch mode)
$ gh pr checks --web                # open in browser
```

| Flag | Description |
|------|-------------|
| `--fail-fast` | Exit watch mode on first check failure |
| `-i, --interval int` | Refresh interval in seconds in watch mode (default 10) |
| `--required` | Only show checks that are required |
| `--watch` | Watch checks until they finish |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `bucket`, `completedAt`, `description`, `event`, `link`, `name`, `startedAt`, `state`, `workflow`.

> Additional exit code: `8` (checks pending).

### `gh pr update-branch`

Update a PR branch with latest changes of the base branch. Without an argument, selects the PR for the current branch.

```sh
$ gh pr update-branch 23
$ gh pr update-branch 23 --rebase    # rebase instead of merge
```

| Flag | Description |
|------|-------------|
| `--rebase` | Update PR branch by rebasing on top of latest base branch |

### `gh pr revert`

Revert a pull request.

```sh
$ gh pr revert 123
$ gh pr revert 123 --title "Revert: fix bug" --body "Reverting this PR"
$ gh pr revert 123 --draft           # mark revert PR as draft
```

| Flag | Description |
|------|-------------|
| `-b, --body string` | Body for the revert pull request |
| `-F, --body-file file` | Read body from file (`-` for stdin) |
| `-d, --draft` | Mark revert PR as a draft |
| `-t, --title string` | Title for the revert pull request |

### `gh pr lock` / `gh pr unlock`

Lock/unlock pull request conversation.

```sh
$ gh pr lock 123
$ gh pr lock 123 --reason off_topic
$ gh pr unlock 123
```

| Flag (`lock`) | Description |
|--------|-------------|
| `-r, --reason string` | Reason: `off_topic`, `resolved`, `spam`, `too_heated` |

## Search Syntax

The `--search` flag in `gh pr list` uses GitHub's search syntax. See <https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests>.

## Related

- For issues (similar lifecycle), see [issue.md](issue.md).
- For repository targeting (`--repo`), see [repo.md](repo.md).
- For JSON output formatting, see [output-formatting.md](output-formatting.md).
- For search across all PRs, see [search.md](search.md).
