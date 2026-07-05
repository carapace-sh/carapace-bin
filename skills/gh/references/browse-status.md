# Browse, Status, and Codespaces

How to open things in the browser, view work status, and manage codespaces with `gh`.

## Browse

`gh browse` opens repositories, issues, pull requests, files, commits, and more in the web browser.

### Argument Format

A browser location can be specified as:
- By number (issue or PR) — e.g. `123`
- By path (folders and files) — e.g. `cmd/gh/main.go`
- By commit SHA
- With line number — `path:line` (e.g. `main.go:312`)

### Usage

```sh
$ gh browse                    # open repo home page
$ gh browse script/            # open a directory
$ gh browse 217                # open issue or PR 217
$ gh browse 77507cd94ccafcf... # open commit page
$ gh browse --settings         # open settings
$ gh browse main.go:312        # open file at line
$ gh browse main.go:312 --blame   # blame view
$ gh browse main.go --branch bug-fix
$ gh browse main.go --commit=77507cd
$ gh browse --no-browser       # print URL only
```

| Flag | Description |
|------|-------------|
| `-a, --actions` | Open repository actions |
| `--blame` | Open blame view for a file |
| `-b, --branch string` | Select another branch |
| `-c, --commit string[="last"]` | Select a commit by SHA (default: last commit) |
| `-n, --no-browser` | Print URL instead of opening browser |
| `-p, --projects` | Open repository projects |
| `-r, --releases` | Open repository releases |
| `-R, --repo [HOST/]OWNER/REPO` | Select another repository |
| `-s, --settings` | Open repository settings |
| `-w, --wiki` | Open repository wiki |

The `BROWSER` or `GH_BROWSER` environment variable configures the web browser.

## Status

`gh status` prints information about your work on GitHub across all subscribed repositories:

- Assigned Issues
- Assigned Pull Requests
- Review Requests
- Mentions
- Repository Activity (new issues/PRs, comments)

```sh
$ gh status
$ gh status -e cli/cli -e cli/go-gh    # exclude repos
$ gh status -o cli                      # limit to an organization
```

| Flag | Description |
|------|-------------|
| `-e, --exclude strings` | Comma-separated repos to exclude (owner/name format) |
| `-o, --org string` | Report status within an organization |

## Codespaces

`gh codespace` (aliased `gh cs`) connects to and manages codespaces.

### Subcommands

| Subcommand | Description |
|------------|-------------|
| `code` | Open a codespace in VS Code |
| `cp` | Copy files between local and remote codespace |
| `create` | Create a codespace |
| `delete` | Delete a codespace |
| `edit` | Edit a codespace |
| `jupyter` | Open a codespace in JupyterLab |
| `list` | List codespaces |
| `logs` | Access codespace logs |
| `ports` | List and manage forwarded ports |
| `rebuild` | Rebuild a codespace |
| `ssh` | SSH into a codespace |
| `stop` | Stop a codespace |
| `view` | View details about a codespace |

### `gh codespace create`

```sh
$ gh codespace create                          # interactive
$ gh codespace create --repo cli/cli
$ gh codespace create --repo cli/cli --branch develop
$ gh codespace create --machine basicLinux32gb
$ gh codespace create --display-name "My Codespace"
$ gh codespace create --idle-timeout 30m
$ gh codespace create --retention-period 720h
$ gh codespace create --location WestUs2
$ gh codespace create --web
```

Key flags for `create`:

| Flag | Description |
|------|-------------|
| `-b, --branch string` | Repository branch |
| `--default-permissions` | Do not prompt to accept additional permissions requested by the codespace |
| `--devcontainer-path string` | Path to devcontainer.json |
| `-d, --display-name string` | Display name (48 characters or less) |
| `--idle-timeout duration` | Allowed inactivity before stopping, e.g. `10m`, `1h` |
| `-l, --location string` | Location: `EastUs`, `SouthEastAsia`, `WestEurope`, `WestUs2` |
| `-m, --machine string` | Hardware specifications for the VM |
| `-R, --repo string` | Repository name with owner: `user/repo` |
| `--retention-period duration` | Allowed time after shutdown before auto-deletion (max 30 days), e.g. `1h`, `72h` |
| `-s, --status` | Show status of post-create command and dotfiles |
| `-w, --web` | Create codespace from browser |

### `gh codespace list`

List codespaces.

```sh
$ gh codespace list
$ gh codespace list --limit 50
```

### `gh codespace code`

Open a codespace in VS Code.

```sh
$ gh codespace code                    # interactive
$ gh codespace code -c <codespace-name>
$ gh codespace code --web              # open in VS Code Web
```

### `gh codespace ssh`

SSH into a codespace.

```sh
$ gh codespace ssh
$ gh codespace ssh -c <codespace-name>
$ gh codespace ssh --profile <profile-name>
```

### `gh codespace logs`

Access codespace logs.

### `gh codespace ports`

List and manage forwarded ports.

```sh
$ gh codespace ports
$ gh codespace ports visibility 8000:public   # change port visibility
$ gh codespace ports forward 8000:8000        # forward a port
```

### `gh codespace cp`

Copy files between local and remote codespace using `scp` syntax.

```sh
$ gh codespace cp -c <name> remote.txt local.txt
$ gh codespace cp -c <name> ./localdir remote:/workspace
```

### `gh codespace view`

View details about a codespace.

```sh
$ gh codespace view
$ gh codespace view -c <codespace-name>
```

### `gh codespace delete`

Delete a codespace.

### `gh codespace stop`

Stop a running codespace.

### `gh codespace rebuild`

Rebuild a codespace.

### `gh codespace edit`

Edit a codespace's settings (display name, machine type, idle timeout).

### `gh codespace jupyter`

Open a codespace in JupyterLab.

## Related

- For environment variables (`BROWSER`, `GH_BROWSER`), see [config.md](config.md).
- For repository targeting (`--repo`), see [repo.md](repo.md).
