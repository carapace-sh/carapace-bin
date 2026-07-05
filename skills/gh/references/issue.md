# Issues

How to create, view, edit, close, comment on, and manage GitHub issues with `gh issue`.

## Overview

`gh issue` manages GitHub issues. An issue can be referenced by number or URL. The `-R, --repo [HOST/]OWNER/REPO` flag targets a specific repository.

## Argument Format

An issue can be supplied as:
- By number — e.g. `123`
- By URL — e.g. `https://github.com/OWNER/REPO/issues/123`

## Commands

### `gh issue create`

Create an issue on GitHub. Aliased as `gh issue new`.

```sh
$ gh issue create --title "I found a bug" --body "Nothing works"
$ gh issue create --label "bug,help wanted"
$ gh issue create --label bug --label "help wanted"
$ gh issue create --assignee monalisa,hubot
$ gh issue create --assignee "@me"
$ gh issue create --assignee "@copilot"
$ gh issue create --project "Roadmap"
$ gh issue create --template "Bug Report"
$ gh issue create --type Bug
$ gh issue create --parent 100
$ gh issue create --parent https://github.com/cli/go-gh/issues/42
$ gh issue create --blocked-by 200,201 --blocking 300
```

| Flag | Description |
|------|-------------|
| `-a, --assignee login` | Assign people (`@me`, `@copilot`) |
| `--blocked-by numbers` | Mark as blocked by these issue numbers/URLs |
| `--blocking numbers` | Mark as blocking these issue numbers/URLs |
| `-b, --body string` | Issue body |
| `-F, --body-file file` | Read body from file (`-` for stdin) |
| `-e, --editor` | Open editor for title and body |
| `-l, --label name` | Add labels by name |
| `-m, --milestone name` | Add to milestone by name |
| `--parent number` | Add as sub-issue of this parent (number or URL) |
| `-p, --project title` | Add to projects by title |
| `--recover string` | Recover input from a failed run |
| `-T, --template name` | Template name for starting body text |
| `-t, --title string` | Issue title |
| `--type name` | Set issue type by name |
| `-w, --web` | Open browser to create issue |

**Special assignee values**: `@me` (self-assign), `@copilot` (assign Copilot; not supported on GHES).

**Project scope**: Adding issues to projects requires the `project` scope (`gh auth refresh -s project`).

### `gh issue list`

List issues in a repository. Aliased as `gh issue ls`. Defaults to open issues.

```sh
$ gh issue list --label "bug" --label "help wanted"
$ gh issue list --author monalisa
$ gh issue list --assignee "@me"
$ gh issue list --milestone "The big 1.0"
$ gh issue list --search "error no:assignee sort:created-asc"
$ gh issue list --state all
$ gh issue list --type Bug
```

| Flag | Description |
|------|-------------|
| `--app string` | Filter by GitHub App author |
| `-a, --assignee string` | Filter by assignee |
| `-A, --author string` | Filter by author |
| `-l, --label strings` | Filter by label |
| `-L, --limit int` | Max items (default 30) |
| `--mention string` | Filter by mention |
| `-m, --milestone string` | Filter by milestone number or title |
| `-S, --search query` | Search with query syntax |
| `-s, --state string` | Filter by state: `open`, `closed`, `all` (default `open`) |
| `--type name` | Filter by issue type name |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `assignees`, `author`, `blockedBy`, `blocking`, `body`, `closed`, `closedAt`, `closedByPullRequestsReferences`, `comments`, `createdAt`, `id`, `isPinned`, `issueType`, `labels`, `milestone`, `number`, `parent`, `projectCards`, `projectItems`, `reactionGroups`, `state`, `stateReason`, `subIssues`, `subIssuesSummary`, `title`, `updatedAt`, `url`.

### `gh issue view`

Display issue title, body, and other information.

```sh
$ gh issue view 123
$ gh issue view 123 --comments    # view comments
$ gh issue view 123 --web         # open in browser
```

| Flag | Description |
|------|-------------|
| `-c, --comments` | View issue comments |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output (same fields as `issue list`) |

### `gh issue edit`

Edit one or more issues within the same repository. Supports bulk editing (multiple numbers/URLs).

```sh
$ gh issue edit 23 --title "I found a bug" --body "Nothing works"
$ gh issue edit 23 --add-label "bug,help wanted" --remove-label "core"
$ gh issue edit 23 --add-assignee "@me" --remove-assignee monalisa,hubot
$ gh issue edit 23 --add-assignee "@copilot"
$ gh issue edit 23 --add-project "Roadmap" --remove-project v1,v2
$ gh issue edit 23 --milestone "Version 1"
$ gh issue edit 23 --remove-milestone
$ gh issue edit 23 --body-file body.txt
$ gh issue edit 23 34 --add-label "help wanted"   # bulk edit
$ gh issue edit 23 --type Bug
$ gh issue edit 23 --remove-type
$ gh issue edit 23 --parent 100
$ gh issue edit 23 --remove-parent
$ gh issue edit 100 --add-sub-issue 123,124
$ gh issue edit 123 --add-blocked-by 200 --add-blocking 300,301
```

| Flag | Description |
|------|-------------|
| `--add-assignee login` | Add assignees (`@me`, `@copilot`) |
| `--add-blocked-by number` | Add 'blocked by' relationships |
| `--add-blocking number` | Add 'blocking' relationships |
| `--add-label name` | Add labels |
| `--add-project title` | Add to projects |
| `--add-sub-issue number` | Add sub-issues |
| `-b, --body string` | Set new body |
| `-F, --body-file file` | Read body from file |
| `-m, --milestone name` | Set milestone |
| `--parent number` | Set parent issue |
| `--remove-assignee login` | Remove assignees |
| `--remove-blocked-by number` | Remove 'blocked by' relationships |
| `--remove-blocking number` | Remove 'blocking' relationships |
| `--remove-label name` | Remove labels |
| `--remove-milestone` | Remove milestone |
| `--remove-parent` | Remove parent issue |
| `--remove-project title` | Remove from projects |
| `--remove-sub-issue number` | Remove sub-issues |
| `--remove-type` | Remove issue type |
| `-t, --title string` | Set new title |
| `--type name` | Set issue type |

### `gh issue close`

Close an issue.

```sh
$ gh issue close 123
$ gh issue close 123 --comment "Closing this issue"
$ gh issue close 123 --duplicate-of 456
$ gh issue close 123 --reason "not planned"
```

| Flag | Description |
|------|-------------|
| `-c, --comment string` | Leave a closing comment |
| `--duplicate-of string` | Mark as duplicate of another issue (number or URL) |
| `-r, --reason string` | Reason: `completed`, `not planned`, `duplicate` |

### `gh issue reopen`

Reopen an issue.

```sh
$ gh issue reopen 123
$ gh issue reopen 123 --comment "Reopening this"
```

| Flag | Description |
|------|-------------|
| `-c, --comment string` | Add a reopening comment |

### `gh issue comment`

Add a comment to an issue.

```sh
$ gh issue comment 12 --body "Hi from GitHub CLI"
$ gh issue comment 12 --body-file comment.md
$ gh issue comment 12 --edit-last
$ gh issue comment 12 --edit-last --create-if-none
$ gh issue comment 12 --delete-last
$ gh issue comment 12 --delete-last --yes
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

### `gh issue develop`

Manage linked branches for an issue — create a development branch from an issue.

```sh
$ gh issue develop --list 123                       # list linked branches
$ gh issue develop --list --repo cli/cli 123
$ gh issue develop 123 --base my-feature            # create from branch
$ gh issue develop 123 --checkout                   # create and checkout
$ gh issue develop 123 --repo cli/cli --branch-repo monalisa/cli  # cross-repo
```

| Flag | Description |
|------|-------------|
| `-b, --base string` | Remote branch to create from |
| `--branch-repo string` | Repository to create branch in |
| `-c, --checkout` | Checkout the branch after creating |
| `-l, --list` | List linked branches |
| `-n, --name string` | Name of branch to create |

When using `--base`, the new branch is configured as the base branch for PRs created with `gh pr create`.

### `gh issue status`

Show status of relevant issues (assigned to you, mentioning you, etc.).

### `gh issue transfer`

Transfer an issue to another repository.

```sh
$ gh issue transfer 123 owner/new-repo
```

### `gh issue pin` / `gh issue unpin`

Pin/unpin an issue to a repository.

```sh
$ gh issue pin 23
$ gh issue pin https://github.com/owner/repo/issues/23
$ gh issue pin 23 --repo owner/repo
$ gh issue unpin 23
```

### `gh issue lock` / `gh issue unlock`

Lock/unlock issue conversation.

```sh
$ gh issue lock 123
$ gh issue lock 123 --reason spam
$ gh issue unlock 123
```

| Flag (`lock`) | Description |
|--------|-------------|
| `-r, --reason string` | Reason: `off_topic`, `resolved`, `spam`, `too_heated` |

### `gh issue delete`

Delete an issue. Requires confirmation unless `--yes`.

```sh
$ gh issue delete 123
$ gh issue delete 123 --yes
```

| Flag | Description |
|------|-------------|
| `--yes` | Confirm deletion without prompting |

## Sub-Issue and Dependency Relationships

Issues can have parent/sub-issue relationships and blocking dependencies:

- **Parent/sub-issue**: Set with `--parent` (create) or `--add-sub-issue`/`--remove-sub-issue` (edit)
- **Blocked by**: `--blocked-by` (create), `--add-blocked-by`/`--remove-blocked-by` (edit)
- **Blocking**: `--blocking` (create), `--add-blocking`/`--remove-blocking` (edit)

All accept issue numbers or URLs.

## Search Syntax

The `--search` flag in `gh issue list` uses GitHub's search syntax. See <https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests>.

## Related

- For pull requests (similar lifecycle), see [pr.md](pr.md).
- For repository targeting (`--repo`), see [repo.md](repo.md).
- For JSON output formatting, see [output-formatting.md](output-formatting.md).
- For searching across all issues, see [search.md](search.md).
