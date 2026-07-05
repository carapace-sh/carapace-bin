# Repository Management

How to create, clone, fork, view, edit, sync, and manage GitHub repositories with `gh repo`.

## Overview

`gh repo` manages GitHub repositories. Repository arguments use the `OWNER/REPO` format or full URLs. When the `OWNER/` portion is omitted, it defaults to the authenticating user.

## Argument Format

A repository can be supplied as:
- `"OWNER/REPO"` — e.g. `cli/cli`
- By URL — e.g. `https://github.com/OWNER/REPO`

## Commands

### `gh repo create`

Create a new GitHub repository. Aliased as `gh repo new`.

```sh
$ gh repo create                              # interactive
$ gh repo create my-project --public --clone  # remote + local clone
$ gh repo create my-org/my-project --public   # in an organization
$ gh repo create my-project --private --source=. --remote=upstream  # from existing dir
```

| Flag | Description |
|------|-------------|
| `--add-readme` | Add a README file |
| `-c, --clone` | Clone the new repository to current directory |
| `-d, --description string` | Repository description |
| `--disable-issues` | Disable issues |
| `--disable-wiki` | Disable wiki |
| `-g, --gitignore string` | Gitignore template (see [github/gitignore](https://github.com/github/gitignore)) |
| `-h, --homepage URL` | Repository homepage URL |
| `--include-all-branches` | Include all branches from template repository |
| `--internal` | Make repository internal |
| `-l, --license string` | Open source license (run `gh repo license list` or see [choosealicense.com](https://choosealicense.com)) |
| `--private` | Make repository private |
| `--public` | Make repository public |
| `--push` | Push local commits to the new repository |
| `-r, --remote string` | Remote name for the new repository |
| `-s, --source string` | Path to local repository to use as source |
| `-t, --team name` | Organization team to grant access |
| `-p, --template repository` | Base on a template repository |

When creating from a local source (`--source`), the remote repo name defaults to the source directory name. If the repo is bare, `--push` mirrors all refs.

### `gh repo clone`

Clone a GitHub repository locally. Pass additional `git clone` flags after `--`.

```sh
$ gh repo clone cli/cli
$ gh repo clone myrepo                     # defaults to authenticating user
$ gh repo clone https://github.com/cli/cli # override git protocol
$ gh repo clone cli/cli -- --depth=1      # pass git clone flags
$ gh repo clone myfork --no-upstream       # fork without upstream remote
```

| Flag | Description |
|------|-------------|
| `--no-upstream` | Do not add upstream remote when cloning a fork |
| `-u, --upstream-remote-name string` | Upstream remote name (default `upstream`; supports `@owner` for owner-based naming) |

If the repo is a fork, its parent is added as an `upstream` remote and set as the default remote repository. Use `--no-upstream` to skip this.

### `gh repo fork`

Create a fork of a repository. With no argument, forks the current repository.

```sh
$ gh repo fork                           # fork current repo
$ gh repo fork cli/cli --clone           # fork and clone
$ gh repo fork cli/cli --org myOrg       # fork into an organization
$ gh repo fork cli/cli --remote=false    # don't add a remote
```

| Flag | Description |
|------|-------------|
| `--clone` | Clone the fork |
| `--default-branch-only` | Only include the default branch |
| `--fork-name string` | Rename the forked repository |
| `--org string` | Create the fork in an organization |
| `--remote` | Add a git remote for the fork |
| `--remote-name string` | Name for the new remote (default `origin`) |

By default, the new fork is set as `origin` and any existing origin is renamed to `upstream`. The upstream remote becomes the default remote repository.

### `gh repo view`

Display the description and README of a repository. With no argument, shows the current directory's repository.

```sh
$ gh repo view
$ gh repo view --web              # open in browser
$ gh repo view --branch develop   # view a specific branch
```

| Flag | Description |
|------|-------------|
| `-b, --branch string` | View a specific branch |
| `-w, --web` | Open in browser |
| `--json fields` | Output JSON |
| `--jq expression` | Filter JSON |
| `--template string` | Format with Go template |

**JSON fields**: `archivedAt`, `assignableUsers`, `codeOfConduct`, `contactLinks`, `createdAt`, `defaultBranchRef`, `deleteBranchOnMerge`, `description`, `diskUsage`, `forkCount`, `fundingLinks`, `hasDiscussionsEnabled`, `hasIssuesEnabled`, `hasProjectsEnabled`, `hasWikiEnabled`, `homepageUrl`, `id`, `isArchived`, `isBlankIssuesEnabled`, `isEmpty`, `isFork`, `isInOrganization`, `isMirror`, `isPrivate`, `isSecurityPolicyEnabled`, `isTemplate`, `isUserConfigurationRepository`, `issueTemplates`, `issues`, `labels`, `languages`, `latestRelease`, `licenseInfo`, `mentionableUsers`, `mergeCommitAllowed`, `milestones`, `mirrorUrl`, `name`, `nameWithOwner`, `openGraphImageUrl`, `owner`, `parent`, `primaryLanguage`, `projects`, `projectsV2`, `pullRequestTemplates`, `pullRequests`, `pushedAt`, `rebaseMergeAllowed`, `repositoryTopics`, `securityPolicyUrl`, `squashMergeAllowed`, `sshUrl`, `stargazerCount`, `templateRepository`, `updatedAt`, `url`, `usesCustomOpenGraphImage`, `viewerCanAdminister`, `viewerDefaultCommitEmail`, `viewerDefaultMergeMethod`, `viewerHasStarred`, `viewerPermission`, `viewerPossibleCommitEmails`, `viewerSubscription`, `visibility`, `watchers`.

### `gh repo edit`

Edit repository settings. Toggle settings off with `--<flag>=false`.

```sh
$ gh repo edit --enable-issues --enable-wiki
$ gh repo edit --enable-projects=false
$ gh repo edit --visibility private --accept-visibility-change-consequences
```

| Flag | Description |
|------|-------------|
| `--accept-visibility-change-consequences` | Required when using `--visibility` |
| `--add-topic strings` | Add repository topic |
| `--allow-forking` | Allow forking of an organization repository |
| `--allow-update-branch` | Allow updating a behind base branch |
| `--default-branch name` | Set default branch name |
| `--delete-branch-on-merge` | Delete head branch when PRs are merged |
| `-d, --description string` | Description |
| `--enable-advanced-security` | Enable advanced security |
| `--enable-auto-merge` | Enable auto-merge |
| `--enable-discussions` | Enable discussions |
| `--enable-issues` | Enable issues |
| `--enable-merge-commit` | Enable merge commit |
| `--enable-projects` | Enable projects |
| `--enable-rebase-merge` | Enable rebase merge |
| `--enable-secret-scanning` | Enable secret scanning |
| `--enable-secret-scanning-push-protection` | Enable secret scanning push protection (requires secret scanning enabled first) |
| `--enable-squash-merge` | Enable squash merge |
| `--enable-wiki` | Enable wiki |
| `-h, --homepage URL` | Homepage URL |
| `--remove-topic strings` | Remove repository topic |
| `--squash-merge-commit-message string` | Default squash merge commit message: `default`, `pr-title`, `pr-title-commits`, `pr-title-description` |
| `--template` | Make available as a template repository |
| `--visibility string` | Change visibility: `public`, `private`, `internal` |

> **Visibility change consequences**: Losing stars/watchers, detaching public forks, disabling push rulesets, allowing access to Actions history. See <https://gh.io/setting-repository-visibility>.

### `gh repo list`

List repositories owned by a user or organization. Aliased as `gh repo ls`.

```sh
$ gh repo list
$ gh repo list cli --language Go --limit 50
$ gh repo list --fork       # only forks
$ gh repo list --source     # only non-forks
$ gh repo list --topic api  # filter by topic
```

| Flag | Description |
|------|-------------|
| `--archived` | Show only archived repositories |
| `--fork` | Show only forks |
| `-l, --language string` | Filter by primary language |
| `-L, --limit int` | Max repositories (default 30) |
| `--no-archived` | Omit archived repositories |
| `--source` | Show only non-forks |
| `--topic strings` | Filter by topic |
| `--visibility string` | Filter by visibility: `public`, `private`, `internal` |
| `--json fields` / `--jq` / `--template` | JSON output (same fields as `repo view`) |

> The list only includes repositories owned by the provided argument. `--fork` or `--source` do not traverse ownership boundaries.

### `gh repo sync`

Sync destination repository from source repository. Uses the default branch. Fast-forward by default; `--force` does a hard reset.

```sh
$ gh repo sync                              # local from remote parent
$ gh repo sync --branch v1                  # specific branch
$ gh repo sync owner/cli-fork               # remote fork from parent
$ gh repo sync owner/repo --source owner2/repo2  # custom source
```

| Flag | Description |
|------|-------------|
| `-b, --branch string` | Branch to sync (default: default branch) |
| `--force` | Hard reset to match source |
| `-s, --source string` | Source repository (default: parent) |

Without an argument, the local repository is the destination. The source defaults to the parent of the destination.

### `gh repo delete`

Delete a repository. Requires confirmation.

### `gh repo archive`

Archive a repository (makes it read-only).

### `gh repo unarchive`

Unarchive a repository.

### `gh repo rename`

Rename a repository.

### `gh repo set-default`

Configure the default repository for the current directory. This is the repository used when no `--repo` flag is provided.

### `gh repo gitignore`

List and view available repository gitignore templates.

```sh
$ gh repo gitignore list
$ gh repo gitignore view Go
```

### `gh repo license`

Explore repository licenses.

```sh
$ gh repo license list
$ gh repo license view MIT
```

### `gh repo deploy-key`

Manage deploy keys in a repository. Subcommands: `list`, `add`, `delete`.

### `gh repo autolink`

Manage autolink references. Subcommands: `list`, `create`, `delete`, `view`.

## The `--repo` Flag

Many commands accept `-R, --repo [HOST/]OWNER/REPO` to target a different repository than the current directory's. This is inherited by most `pr`, `issue`, `release`, `workflow`, `run`, `cache`, `label`, `ruleset`, `secret`, and `variable` subcommands.

```sh
$ gh pr list --repo cli/cli
$ gh issue list --repo enterprise.internal/owner/repo
```

## Related

- For pull requests, see [pr.md](pr.md).
- For issues, see [issue.md](issue.md).
- For releases, see [release.md](release.md).
- For authentication and `GH_REPO` env var, see [auth.md](auth.md) and [config.md](config.md).
