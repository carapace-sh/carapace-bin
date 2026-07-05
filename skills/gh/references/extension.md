# Extensions

How to install, create, manage, search for, and browse gh extensions.

## Overview

gh extensions are repositories that provide additional gh commands. An extension repository must:
- Start with `gh-` in its name
- Contain an executable of the same name

All arguments passed to `gh <extname>` are forwarded to the `gh-<extname>` executable.

**Key behaviors**:
- An extension **cannot override** core gh commands. If a name conflicts, use `gh extension exec <extname>`.
- When executed, gh checks for new versions once every 24 hours and displays an upgrade notice.
- Extensions are **not verified, signed, or endorsed** by GitHub. Review source before installing.

For available extensions, see <https://github.com/topics/gh-extension>.

## Commands

### `gh extension install`

Install a gh extension from a GitHub or local repository. Aliased via `gh ext install`.

```sh
$ gh extension install owner/gh-extension
$ gh extension install https://my.ghes.com/owner/gh-extension
$ gh extension install .                    # local repo (development)
```

| Flag | Description |
|------|-------------|
| `--force` | Force upgrade, or ignore if latest already installed |
| `--pin string` | Pin to a release tag (binary) or commit ref (script) |

**Binary vs script extensions**:
- **Binary**: gh looks for release artifacts (prebuilt binaries). Use `--pin` with a tag.
- **Script**: If no release, the repository is cloned. Use `--pin` with a commit ref.

**Local installation** (`.`): gh manages the extension as a symbolic link to the executable in the repository root. For precompiled extensions, the executable must be built and placed in the root manually.

### `gh extension list`

List installed extension commands. Aliased as `gh ext ls`, `gh extension ls`, `gh extensions ls`.

### `gh extension create`

Create a new extension.

```sh
$ gh extension create                  # interactive
$ gh extension create foobar           # script-based
$ gh extension create --precompiled=go foobar     # Go precompiled
$ gh extension create --precompiled=other foobar  # non-Go precompiled
```

| Flag | Description |
|------|-------------|
| `--precompiled string` | Create a precompiled extension: `go` or `other` |

### `gh extension remove`

Remove an installed extension. Aliased as `gh ext uninstall`, `gh extension uninstall`, `gh extensions uninstall`.

```sh
$ gh extension remove foobar
```

### `gh extension upgrade`

Upgrade installed extensions.

```sh
$ gh extension upgrade foobar
$ gh extension upgrade --all
$ gh extension upgrade --all --dry-run     # show what would upgrade
$ gh extension upgrade --all --force
```

| Flag | Description |
|------|-------------|
| `--all` | Upgrade all extensions |
| `--dry-run` | Only display upgrades |
| `--force` | Force upgrade |

### `gh extension browse`

Interactive TUI for browsing, adding, and removing extensions. Terminal width >100 columns recommended.

```sh
$ gh extension browse
$ gh extension browse --single-column    # for screen readers / high zoom
```

| Flag | Description |
|------|-------------|
| `--debug` | Log to `/tmp/extBrowse-*` |
| `-s, --single-column` | Render with one column (accessible) |

### `gh extension search`

Search for gh extensions. Similar to `gh search repos` but focused on extensions.

```sh
$ gh extension search                    # first 30 by stars
$ gh extension search --limit 300
$ gh extension search branch             # matching "branch"
$ gh extension search --owner github
$ gh extension search --sort updated --order asc
$ gh extension search --license MIT
$ gh extension search -w                 # open in browser
```

| Flag | Description |
|------|-------------|
| `--license strings` | Filter by license type |
| `-L, --limit int` | Max results (default 30) |
| `--order string` | Order: `asc`, `desc` (default `desc`; requires `--sort`) |
| `--owner strings` | Filter on owner |
| `--sort string` | Sort: `forks`, `help-wanted-issues`, `stars`, `updated` (default `best-match`) |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `createdAt`, `defaultBranch`, `description`, `forksCount`, `fullName`, `hasDownloads`, `hasIssues`, `hasPages`, `hasProjects`, `hasWiki`, `homepage`, `id`, `isArchived`, `isDisabled`, `isFork`, `isPrivate`, `language`, `license`, `name`, `openIssuesCount`, `owner`, `pushedAt`, `size`, `stargazersCount`, `updatedAt`, `url`, `visibility`, `watchersCount`.

For finer-grained search, use `gh search repos --topic "gh-extension"` with additional qualifiers.

### `gh extension exec`

Execute an installed extension, bypassing name conflicts with core commands.

```sh
$ gh extension exec <extname> [args...]
```

## Aliases

The `gh extension` command has several aliases:

| Alias | Full form |
|-------|-----------|
| `gh ext` | `gh extension` |
| `gh extensions` | `gh extension` |
| `gh ext ls` | `gh extension list` |
| `gh ext uninstall` | `gh extension remove` |
| `gh extension uninstall` | `gh extension remove` |

## Update Notifications

gh checks for extension updates once every 24 hours when an extension is executed. Disable with:

```sh
$ export GH_NO_EXTENSION_UPDATE_NOTIFIER=1
```

## Related

- For environment variables (`GH_NO_EXTENSION_UPDATE_NOTIFIER`), see [config.md](config.md).
- For searching repos (used to find extensions), see [search.md](search.md).
- For JSON output formatting, see [output-formatting.md](output-formatting.md).
