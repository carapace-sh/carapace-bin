# Secrets and Variables

How to manage GitHub Actions secrets and variables at the repository, environment, organization, and user levels.

## Overview

`gh secret` and `gh variable` manage secrets and variables for GitHub Actions, Dependabot, Agents, and Codespaces. Secrets are encrypted before being sent to GitHub. Variables are stored as plaintext. Both can be set at multiple levels with different visibility scopes.

## Levels and Scopes

| Level | Secret | Variable | Available To |
|-------|--------|----------|-------------|
| Repository | ✓ (default) | ✓ (default) | Actions, Agents, Dependabot in a repository |
| Environment | ✓ | ✓ | Actions for a deployment environment in a repository |
| Organization | ✓ | ✓ | Actions, Agents, Dependabot, Codespaces within an org |
| User | ✓ (Codespaces) | — | Codespaces for your user |

## Secret Commands

### `gh secret set`

Set a value for a secret. Secret values are locally encrypted before being sent to GitHub.

```sh
$ gh secret set MYSECRET                          # interactive prompt
$ gh secret set MYSECRET --body "$ENV_VALUE"      # from env var
$ gh secret set MYSECRET < myfile.txt             # from file via stdin
$ gh secret set MYSECRET --env myenvironment      # environment-level
$ gh secret set MYSECRET --org myOrg --visibility all           # org, all repos
$ gh secret set MYSECRET --org myOrg --repos repo1,repo2,repo3  # org, selected
$ gh secret set MYSECRET --org myOrg --no-repos-selected         # org, none
$ gh secret set MYSECRET --user                    # user-level (Codespaces)
$ gh secret set MYSECRET --app dependabot          # for Dependabot
$ gh secret set -f .env                           # multiple from dotenv file
$ gh secret set -f - < myfile.txt                 # multiple from stdin
```

| Flag | Description |
|------|-------------|
| `-a, --app string` | Application: `actions`, `agents`, `codespaces`, `dependabot` |
| `-b, --body string` | Secret value (reads from stdin if not specified) |
| `-e, --env environment` | Set deployment environment secret |
| `-f, --env-file file` | Load names and values from a dotenv-formatted file |
| `--no-repos-selected` | No repositories can access the org secret |
| `--no-store` | Print encrypted, base64-encoded value instead of storing |
| `-o, --org organization` | Set organization secret |
| `-r, --repos repositories` | Repositories that can access an org/user secret |
| `-u, --user` | Set a secret for your user (Codespaces) |
| `-v, --visibility string` | Org secret visibility: `all`, `private`, `selected` (default `private`) |

### `gh secret list`

List secrets. Aliased as `gh secret ls`.

```sh
$ gh secret list
$ gh secret list --env myenvironment
$ gh secret list --org myOrg
$ gh secret list --user
$ gh secret list --app dependabot
```

| Flag | Description |
|------|-------------|
| `-a, --app string` | List for specific app: `actions`, `agents`, `codespaces`, `dependabot` |
| `-e, --env string` | List for an environment |
| `-o, --org string` | List for an organization |
| `-u, --user` | List for your user |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `name`, `numSelectedRepos`, `selectedReposURL`, `updatedAt`, `visibility`.

### `gh secret delete`

Delete a secret. Aliased as `gh secret remove`.

```sh
$ gh secret delete MYSECRET
$ gh secret delete MYSECRET --env myenvironment
$ gh secret delete MYSECRET --org myOrg
$ gh secret delete MYSECRET --user
$ gh secret delete MYSECRET --app dependabot
```

| Flag | Description |
|------|-------------|
| `-a, --app string` | Delete for specific app |
| `-e, --env string` | Delete for an environment |
| `-o, --org string` | Delete for an organization |
| `-u, --user` | Delete for your user |

## Variable Commands

### `gh variable set`

Set a value for a variable.

```sh
$ gh variable set MYVARIABLE                       # interactive
$ gh variable set MYVARIABLE --body "$ENV_VALUE"
$ gh variable set MYVARIABLE < myfile.txt
$ gh variable set MYVARIABLE --env myenvironment
$ gh variable set MYVARIABLE --org myOrg --visibility all
$ gh variable set MYVARIABLE --org myOrg --repos repo1,repo2,repo3
$ gh variable set -f .env                          # multiple from dotenv
```

| Flag | Description |
|------|-------------|
| `-b, --body string` | Variable value (reads from stdin if not specified) |
| `-e, --env environment` | Set deployment environment variable |
| `-f, --env-file file` | Load names and values from a dotenv-formatted file |
| `-o, --org organization` | Set organization variable |
| `-r, --repos repositories` | Repositories that can access an org variable |
| `-v, --visibility string` | Org variable visibility: `all`, `private`, `selected` (default `private`) |

### `gh variable get`

Get the value of a variable.

```sh
$ gh variable get MYVARIABLE
$ gh variable get MYVARIABLE --env myenvironment
$ gh variable get MYVARIABLE --org myOrg
```

| Flag | Description |
|------|-------------|
| `-e, --env string` | Get for an environment |
| `-o, --org string` | Get for an organization |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `createdAt`, `name`, `numSelectedRepos`, `selectedReposURL`, `updatedAt`, `value`, `visibility`.

### `gh variable list`

List variables. Aliased as `gh variable ls`.

```sh
$ gh variable list
$ gh variable list --env myenvironment
$ gh variable list --org myOrg
```

| Flag | Description |
|------|-------------|
| `-e, --env string` | List for an environment |
| `-o, --org string` | List for an organization |
| `--json` / `--jq` / `--template` | JSON output |

### `gh variable delete`

Delete a variable. Aliased as `gh variable remove`.

```sh
$ gh variable delete MYVARIABLE
$ gh variable delete MYVARIABLE --env myenvironment
$ gh variable delete MYVARIABLE --org myOrg
```

| Flag | Description |
|------|-------------|
| `-e, --env string` | Delete for an environment |
| `-o, --org string` | Delete for an organization |

## Dotenv File Format

The `--env-file` flag reads names and values from a dotenv-formatted file:

```
SECRET1=value1
SECRET2=value2
VARIABLE1=varvalue1
```

Use `-` to read from stdin: `gh secret set -f - < myfile.txt`.

## App Scope

The `--app` flag for secrets specifies which GitHub application the secret is available to:

| App | Description |
|-----|-------------|
| `actions` | GitHub Actions |
| `agents` | GitHub Agents |
| `codespaces` | GitHub Codespaces (user/org level) |
| `dependabot` | Dependabot |

Variables do not have an `--app` flag — they are available to Actions and Dependabot.

## Organization Visibility

For organization-level secrets and variables, `--visibility` controls access:

| Value | Meaning |
|-------|---------|
| `all` | All repositories in the org |
| `private` | Private repositories only (default) |
| `selected` | Only specific repositories (use `--repos` or `--no-repos-selected`) |

## Related

- For repository targeting (`--repo`), see [repo.md](repo.md).
- For JSON output formatting, see [output-formatting.md](output-formatting.md).
- For authentication scopes, see [auth.md](auth.md).
