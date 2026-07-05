# Authentication

How `gh` authenticates with GitHub hosts, manages credentials, and handles multi-account setups.

## Overview

`gh auth` manages authentication for GitHub.com and GitHub Enterprise Server (GHES) hosts. Credentials are stored in the system credential store (with a plain-text fallback), or supplied via environment variables. The `gh auth` command group handles login, logout, token retrieval, scope management, and git credential helper setup.

## Authentication Methods

### Web-Based Browser Flow (default)

The default `gh auth login` flow opens a browser for OAuth authentication. After completion, a token is stored securely in the system credential store. If no credential store is found, the token is written to a plain-text file (see `gh auth status` for its location).

```sh
$ gh auth login                      # interactive, web flow
$ gh auth login --web --clipboard    # open browser, copy OAuth code to clipboard
$ gh auth login --hostname enterprise.internal
```

### Token via Standard Input

Pass a personal access token (classic) via stdin. Minimum required scopes: `repo`, `read:org`, and `gist`.

```sh
$ gh auth login --with-token < mytoken.txt
$ echo "$TOKEN" | gh auth login --with-token
```

> **Caution with fine-grained PATs**: The inherent scoping to certain resources may cause confusing behaviour when interacting with other resources. Favour setting `GH_TOKEN` for fine-grained PAT usage.

### Environment Variables (headless / automation)

gh uses authentication tokens found in environment variables. This is most suitable for "headless" use such as in automation.

| Variable | Scope |
|----------|-------|
| `GH_TOKEN` | `github.com` and subdomains of `ghe.com` (highest precedence) |
| `GITHUB_TOKEN` | `github.com` and subdomains of `ghe.com` (second precedence) |
| `GH_ENTERPRISE_TOKEN` | GitHub Enterprise Server hosts (highest precedence) |
| `GITHUB_ENTERPRISE_TOKEN` | GitHub Enterprise Server hosts (second precedence) |

For GitHub Actions, add `GH_TOKEN: ${{ github.token }}` to `env`.

## Commands

### `gh auth login`

Authenticate with a GitHub host.

| Flag | Description |
|------|-------------|
| `-c, --clipboard` | Copy one-time OAuth device code to clipboard |
| `-p, --git-protocol string` | Protocol for git operations: `ssh` or `https` |
| `-h, --hostname string` | GitHub instance hostname (default: `github.com`) |
| `--insecure-storage` | Save credentials in plain text instead of credential store |
| `-s, --scopes strings` | Additional authentication scopes to request |
| `--skip-ssh-key` | Skip generate/upload SSH key prompt |
| `-w, --web` | Open a browser to authenticate |
| `--with-token` | Read token from standard input |

When `--git-protocol ssh` is specified, gh detects existing SSH keys and prompts to create/upload a new key if none is found. Use `--skip-ssh-key` to skip this.

The `--git-protocol` setting takes effect for **all users** on the host (not just the current account), since login is per-account but the git protocol is per-host.

### `gh auth status`

Display active account and authentication state on each known GitHub host. For each host, each known account's authentication is tested and any issues are included.

| Flag | Description |
|------|-------------|
| `-a, --active` | Display the active account only |
| `-h, --hostname string` | Check only a specific hostname |
| `-t, --show-token` | Display the auth token |
| `--json fields` | Output JSON (field: `hosts`) |
| `--jq expression` | Filter JSON output |
| `--template string` | Format JSON output using a Go template |

If any account has authentication issues, the command exits with 1 and outputs to stderr. With `--json`, the command always exits 0 regardless of auth issues (unless a fatal error occurs).

### `gh auth token`

Print the authentication token for an account on a given host. Without `--hostname`, the default host is chosen. Without `--user`, the active account is chosen.

| Flag | Description |
|------|-------------|
| `-h, --hostname string` | GitHub hostname |
| `-u, --user string` | Account to output the token for |

### `gh auth switch`

Switch the active account for a GitHub host. If the host has exactly two accounts, the switch is automatic. With more than two, disambiguation is needed via `--user` or interactive prompt.

| Flag | Description |
|------|-------------|
| `-h, --hostname string` | GitHub hostname |
| `-u, --user string` | Account to switch to |

### `gh auth logout`

Remove authentication for a GitHub account. This only removes the **local** configuration — it does **not** revoke tokens. To revoke all gh-generated tokens, visit <https://github.com/settings/applications>, select "GitHub CLI", and "Revoke Access".

| Flag | Description |
|------|-------------|
| `-h, --hostname string` | GitHub hostname |
| `-u, --user string` | Account to log out of |

### `gh auth refresh`

Expand or fix permission scopes for stored credentials for the active account.

| Flag | Description |
|------|-------------|
| `-c, --clipboard` | Copy one-time OAuth device code to clipboard |
| `-h, --hostname string` | GitHub host |
| `--insecure-storage` | Save credentials in plain text |
| `-r, --remove-scopes strings` | Scopes to remove (idempotent; minimum scopes `repo`, `read:org`, `gist` cannot be removed) |
| `--reset-scopes` | Reset to the default minimum set of scopes |
| `-s, --scopes strings` | Additional scopes to add |

To refresh credentials for an **inactive** account, first `gh auth switch` to it, refresh, then switch back.

### `gh auth setup-git`

Configure `git` to use gh as a credential helper. By default, configures all authenticated hosts. Use `--hostname` for a single host.

| Flag | Description |
|------|-------------|
| `-f, --force` | Force setup even if the host is not known (requires `--hostname`) |
| `-h, --hostname string` | Hostname to configure git for |

## OAuth Scopes

The minimum required scopes for gh are: `repo`, `read:org`, and `gist`. Additional scopes can be requested at login or via `gh auth refresh`. Common additional scopes:

| Scope | Purpose |
|-------|---------|
| `project` | Adding issues/PRs to GitHub Projects |
| `write:org` | Organization write operations |
| `read:public_key` | Public key reading |
| `delete_repo` | Repository deletion |
| `admin:org` | Organization ruleset management |

For OAuth scopes documentation, see <https://docs.github.com/en/developers/apps/building-oauth-apps/scopes-for-oauth-apps/>.

## Multi-Account Support

gh supports multiple authenticated accounts per host. One account per host is "active" — it's the one used when targeting that host. Use `gh auth switch` to change the active account. `gh auth status` shows all accounts and which is active.

```
$ gh auth status
github.com
  ✓ Logged in to github.com as user1 (active account)
  ✓ Logged in to github.com as user2
enterprise.internal
  ✓ Logged in to enterprise.internal as corpuser (active account)
```

## Credential Storage

| Method | When Used |
|--------|-----------|
| System credential store | Default (keychain on macOS, libsecret on Linux, Credential Manager on Windows) |
| Plain-text file | Fallback when no credential store is available, or `--insecure-storage` |
| Environment variables | `GH_TOKEN`/`GITHUB_TOKEN` — takes precedence over stored credentials |

## Related

- For git credential helper internals, see the **git** skill.
- For environment variables, see [config.md](config.md).
- For making raw API requests with the token, see [api.md](api.md).
