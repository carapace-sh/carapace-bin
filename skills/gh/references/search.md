# Search

How to search across GitHub for repositories, issues, pull requests, code, and commits with `gh search`.

## Overview

`gh search` provides five subcommands for searching across GitHub. Each supports constructing queries via GitHub search syntax, parameter/qualifier flags, or a combination of both.

GitHub search syntax documentation:
- Issues and PRs: <https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests>
- Code: <https://docs.github.com/search-github/searching-on-github/searching-code>
- Commits: <https://docs.github.com/search-github/searching-on-github/searching-commits>

> **Hyphen handling**: For information on handling search queries containing a hyphen, run `gh search --help`.

## Commands

### `gh search repos`

Search for repositories on GitHub.

```sh
$ gh search repos cli
$ gh search repos --language Go --stars ">1000"
$ gh search repos --topic api --owner cli
$ gh search repos --visibility public --archived=false
$ gh search repos --sort stars --order desc --limit 50
$ gh search repos --match name --match description
```

| Flag | Description |
|------|-------------|
| `--archived` | Filter by archived state: `true`, `false` |
| `--created date` | Filter by creation date |
| `--followers number` | Filter by followers |
| `--forks number` | Filter by forks |
| `--good-first-issues number` | Filter by good first issues count |
| `--help-wanted-issues number` | Filter by help wanted issues count |
| `--include-forks string` | Include forks: `false`, `true`, `only` |
| `--language string` | Filter by language |
| `--license strings` | Filter by license |
| `-L, --limit int` | Max results (default 30) |
| `--match strings` | Restrict search to: `name`, `description`, `readme` |
| `--number-topics number` | Filter by number of topics |
| `--order string` | Order: `asc`, `desc` (default `desc`; requires `--sort`) |
| `--owner strings` | Filter on owner |
| `--size string` | Filter by size (KB) |
| `--sort string` | Sort: `forks`, `help-wanted-issues`, `stars`, `updated` (default `best-match`) |
| `--stars number` | Filter by stars |
| `--topic strings` | Filter by topic |
| `--updated date` | Filter by last updated date |
| `--visibility strings` | Filter by visibility: `public`, `private`, `internal` |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

### `gh search issues`

Search for issues on GitHub. Does not include pull requests unless `--include-prs` is used.

```sh
$ gh search issues readme typo
$ gh search issues "broken feature"
$ gh search issues --include-prs --owner=cli
$ gh search issues --assignee=@me --state=open
$ gh search issues --comments=">100"
$ gh search issues -- -label:bug          # exclude label
$ gh search issues --owner github --archived=false
```

| Flag | Description |
|------|-------------|
| `--app string` | Filter by GitHub App author |
| `--archived` | Filter by repo archived state: `true`, `false` |
| `--assignee string` | Filter by assignee |
| `--author string` | Filter by author |
| `--closed date` | Filter on closed at date |
| `--commenter user` | Filter by comments by user |
| `--comments number` | Filter on comment count |
| `--created date` | Filter by creation date |
| `--include-prs` | Include pull requests in results |
| `--interactions number` | Filter by reactions + comments |
| `--involves user` | Filter by involvement of user |
| `--label strings` | Filter on label |
| `--language string` | Filter by repo language |
| `-L, --limit int` | Max results (default 30) |
| `--locked` | Filter on locked conversation status |
| `--match strings` | Restrict to: `title`, `body`, `comments` |
| `--mentions user` | Filter by user mentions |
| `--milestone title` | Filter by milestone title |
| `--no-assignee` | Filter on missing assignee |
| `--no-label` | Filter on missing label |
| `--no-milestone` | Filter on missing milestone |
| `--no-project` | Filter on missing project |
| `--order string` | Order: `asc`, `desc` (default `desc`) |
| `--owner strings` | Filter on repo owner |
| `--project owner/number` | Filter on project board |
| `--reactions number` | Filter by reactions |
| `-R, --repo strings` | Filter on repository |
| `--sort string` | Sort: `comments`, `created`, `interactions`, `reactions`, `reactions-+1`, `reactions--1`, `reactions-heart`, `reactions-smile`, `reactions-tada`, `reactions-thinking_face`, `updated` (default `best-match`) |
| `--state string` | Filter by state: `open`, `closed` |
| `--team-mentions string` | Filter by team mentions |
| `--updated date` | Filter by last updated date |
| `--visibility strings` | Filter by repo visibility |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `assignees`, `author`, `authorAssociation`, `body`, `closedAt`, `commentsCount`, `createdAt`, `id`, `isLocked`, `isPullRequest`, `labels`, `number`, `repository`, `state`, `title`, `updatedAt`, `url`.

### `gh search prs`

Search for pull requests on GitHub.

```sh
$ gh search prs fix bug
$ gh search prs --repo=cli/cli --draft
$ gh search prs --review-requested=@me --state=open
$ gh search prs --assignee=@me --merged
$ gh search prs --reactions=">100"
$ gh search prs --checks failure
$ gh search prs --base main --head feature-branch
```

| Flag | Description |
|------|-------------|
| `--app string` | Filter by GitHub App author |
| `--archived` | Filter by repo archived state |
| `--assignee string` | Filter by assignee |
| `--author string` | Filter by author |
| `-B, --base string` | Filter on base branch name |
| `--checks string` | Filter by check status: `pending`, `success`, `failure` |
| `--closed date` | Filter on closed at date |
| `--commenter user` | Filter by comments by user |
| `--comments number` | Filter on comment count |
| `--created date` | Filter by creation date |
| `--draft` | Filter by draft state |
| `-H, --head string` | Filter on head branch name |
| `--interactions number` | Filter by reactions + comments |
| `--involves user` | Filter by involvement of user |
| `--label strings` | Filter on label |
| `--language string` | Filter by repo language |
| `-L, --limit int` | Max results (default 30) |
| `--locked` | Filter on locked status |
| `--match strings` | Restrict to: `title`, `body`, `comments` |
| `--mentions user` | Filter by user mentions |
| `--merged` | Filter by merged state |
| `--merged-at date` | Filter on merged at date |
| `--milestone title` | Filter by milestone title |
| `--no-assignee` / `--no-label` / `--no-milestone` / `--no-project` | Filter on missing attributes |
| `--order string` | Order: `asc`, `desc` (default `desc`) |
| `--owner strings` | Filter on repo owner |
| `--project owner/number` | Filter on project board |
| `--reactions number` | Filter by reactions |
| `-R, --repo strings` | Filter on repository |
| `--review string` | Filter by review status: `none`, `required`, `approved`, `changes_requested` |
| `--review-requested user` | Filter on user/team requested to review |
| `--reviewed-by user` | Filter on user who reviewed |
| `--sort string` | Sort: `comments`, `reactions`, `reactions-+1`, `reactions--1`, `reactions-smile`, `reactions-thinking_face`, `reactions-heart`, `reactions-tada`, `interactions`, `created`, `updated` (default `best-match`) |
| `--state string` | Filter by state: `open`, `closed` |
| `--team-mentions string` | Filter by team mentions |
| `--updated date` | Filter by last updated date |
| `--visibility strings` | Filter by repo visibility |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `assignees`, `author`, `authorAssociation`, `body`, `closedAt`, `commentsCount`, `createdAt`, `id`, `isDraft`, `isLocked`, `isPullRequest`, `labels`, `number`, `repository`, `state`, `title`, `updatedAt`, `url`.

### `gh search code`

Search within code in GitHub repositories.

> **Note**: This uses a legacy GitHub code search engine. Results may not match `github.com`, and regex search is not yet available via the API.

```sh
$ gh search code react lifecycle
$ gh search code "error handling"
$ gh search code deque --language=python
$ gh search code cli --owner=microsoft
$ gh search code panic --repo cli/cli
$ gh search code lint --filename package.json
$ gh search code --extension ts --match path
```

| Flag | Description |
|------|-------------|
| `--extension string` | Filter on file extension |
| `--filename string` | Filter on filename |
| `--language string` | Filter by language |
| `-L, --limit int` | Max results (default 30) |
| `--match strings` | Restrict to: `file` (contents), `path` |
| `--owner strings` | Filter on owner |
| `-R, --repo strings` | Filter on repository |
| `--size string` | Filter on size range (KB) |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `path`, `repository`, `sha`, `textMatches`, `url`.

### `gh search commits`

Search for commits on GitHub.

```sh
$ gh search commits readme typo
$ gh search commits "bug fix"
$ gh search commits --committer=monalisa
$ gh search commits --author-name="Jane Doe"
$ gh search commits --hash=8dd03144ffdc6c0d...
$ gh search commits --author-date="<2022-02-01"
$ gh search commits --merge   # only merge commits
```

| Flag | Description |
|------|-------------|
| `--author string` | Filter by author |
| `--author-date date` | Filter by authored date |
| `--author-email string` | Filter on author email |
| `--author-name string` | Filter on author name |
| `--committer string` | Filter by committer |
| `--committer-date date` | Filter by committed date |
| `--committer-email string` | Filter on committer email |
| `--committer-name string` | Filter on committer name |
| `--hash string` | Filter by commit hash |
| `-L, --limit int` | Max results (default 30) |
| `--merge` | Filter on merge commits |
| `--order string` | Order: `asc`, `desc` (default `desc`) |
| `--owner strings` | Filter on repo owner |
| `--parent string` | Filter by parent hash |
| `-R, --repo strings` | Filter on repository |
| `--sort string` | Sort: `author-date`, `committer-date` (default `best-match`) |
| `--tree string` | Filter by tree hash |
| `--visibility strings` | Filter by repo visibility |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `author`, `commit`, `committer`, `id`, `parents`, `repository`, `sha`, `url`.

## Excluding with `--`

To exclude a qualifier, use `--` before the exclusion syntax:

```sh
$ gh search issues -- -label:bug          # exclude issues with label "bug"
$ gh search prs -- -label:bug             # exclude PRs with label "bug"
```

## Related

- For listing issues/PRs within a specific repo (not search), see [issue.md](issue.md) and [pr.md](pr.md).
- For JSON output formatting, see [output-formatting.md](output-formatting.md).
