# GitHub Actions — Workflows, Runs, and Caches

How to manage GitHub Actions workflows, workflow runs, and Actions caches with `gh workflow`, `gh run`, and `gh cache`.

## Overview

gh provides three command groups for GitHub Actions:

- **`gh workflow`** — manage workflow files (list, view, enable, disable, run)
- **`gh run`** — manage workflow runs (list, view, watch, download, rerun, cancel, delete)
- **`gh cache`** — manage Actions caches (list, delete)

All three inherit the `-R, --repo [HOST/]OWNER/REPO` flag for targeting a specific repository.

## Workflow Commands

### `gh workflow list`

List workflow files, hiding disabled workflows by default. Aliased as `gh workflow ls`.

```sh
$ gh workflow list
$ gh workflow list --all          # include disabled
$ gh workflow list --limit 100
```

| Flag | Description |
|------|-------------|
| `-a, --all` | Include disabled workflows |
| `-L, --limit int` | Max workflows (default 50) |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `id`, `name`, `path`, `state`.

### `gh workflow view`

View the summary of a workflow. Accepts workflow ID, workflow name, or filename.

```sh
$ gh workflow view                    # interactive selection
$ gh workflow view 0451
$ gh workflow view triage.yml
$ gh workflow view --yaml             # view the YAML file
$ gh workflow view --web              # open in browser
$ gh workflow view --ref my-branch    # view at a specific ref
```

| Flag | Description |
|------|-------------|
| `-r, --ref string` | Branch or tag name containing the workflow version |
| `-w, --web` | Open in browser |
| `-y, --yaml` | View the workflow YAML file |

### `gh workflow run`

Trigger a `workflow_dispatch` event for a given workflow. The workflow file must support an `on.workflow_dispatch` trigger.

```sh
$ gh workflow run                     # interactive
$ gh workflow run triage.yml          # at default branch
$ gh workflow run triage.yml --ref my-branch
$ gh workflow run triage.yml -f name=scully -f greeting=hello
$ echo '{"name":"scully","greeting":"hello"}' | gh workflow run triage.yml --json
```

| Flag | Description |
|------|-------------|
| `-F, --field key=value` | Add a typed parameter (respects `@` syntax, see [api.md](api.md)) |
| `--json` | Read workflow inputs as JSON via stdin |
| `-f, --raw-field key=value` | Add a string parameter |
| `-r, --ref string` | Branch or tag containing the workflow version |

The created workflow run URL is returned if available.

### `gh workflow enable` / `gh workflow disable`

Enable or disable a workflow. Accepts workflow ID or name.

```sh
$ gh workflow enable triage.yml
$ gh workflow disable 0451
```

## Run Commands

### `gh run list`

List recent workflow runs. Aliased as `gh run ls`.

```sh
$ gh run list
$ gh run list --limit 50
$ gh run list --branch main
$ gh run list --status failure
$ gh run list --event push
$ gh run list --workflow triage.yml
$ gh run list --user octocat
$ gh run list --created ">=2024-01-01"
```

| Flag | Description |
|------|-------------|
| `-a, --all` | Include disabled workflows |
| `-b, --branch string` | Filter by branch |
| `-c, --commit SHA` | Filter by commit SHA |
| `--created date` | Filter by creation date |
| `-e, --event event` | Filter by triggering event |
| `-L, --limit int` | Max runs (default 20) |
| `-s, --status string` | Filter by status (see below) |
| `-u, --user string` | Filter by user who triggered |
| `-w, --workflow string` | Filter by workflow |
| `--json` / `--jq` / `--template` | JSON output |

**Status values**: `queued`, `completed`, `in_progress`, `requested`, `waiting`, `pending`, `action_required`, `cancelled`, `failure`, `neutral`, `skipped`, `stale`, `startup_failure`, `success`, `timed_out`.

**JSON fields**: `attempt`, `conclusion`, `createdAt`, `databaseId`, `displayTitle`, `event`, `headBranch`, `headSha`, `name`, `number`, `startedAt`, `status`, `updatedAt`, `url`, `workflowDatabaseId`, `workflowName`.

> Runs created by organization and enterprise ruleset workflows will not display a workflow name due to GitHub API limitations. To see runs associated with a PR, use `gh pr checks`.

### `gh run view`

View a summary of a workflow run.

```sh
$ gh run view                     # interactive selection
$ gh run view 12345
$ gh run view 12345 --attempt 3   # specific attempt
$ gh run view --job 456789        # specific job
$ gh run view --log --job 456789  # full log for a job
$ gh run view --log-failed        # logs for failed steps
$ gh run view --verbose           # show job steps
$ gh run view 0451 --exit-status  # non-zero if run failed
```

| Flag | Description |
|------|-------------|
| `-a, --attempt uint` | Attempt number of the run |
| `--exit-status` | Exit non-zero if run failed |
| `-j, --job string` | View a specific job ID |
| `--log` | View full log for run or specific job |
| `--log-failed` | View log for failed steps |
| `-v, --verbose` | Show job steps |
| `-w, --web` | Open in browser |
| `--json` / `--jq` / `--template` | JSON output |

> Due to platform limitations, `gh` may not always be able to associate jobs with their logs when using the primary zip-based method. It falls back to fetching logs per job via the API (slower). If more than 25 job logs are missing, the operation fails. Some log lines may show as `UNKNOWN STEP` when they can't be associated with a specific step.

### `gh run watch`

Watch a run until it completes, showing progress.

```sh
$ gh run watch
$ gh run watch --compact          # only relevant/failed steps
$ gh run watch --exit-status      # non-zero if run fails
$ gh run watch --interval 5       # refresh every 5 seconds
$ gh run watch && notify-send 'run is done!'
```

| Flag | Description |
|------|-------------|
| `--compact` | Show only relevant/failed steps |
| `--exit-status` | Exit non-zero if run fails |
| `-i, --interval int` | Refresh interval in seconds (default 3) |

> Does not support fine-grained PATs (no `checks:read` permission available for PATs).

### `gh run download`

Download artifacts generated by a workflow run. Each artifact's contents are extracted into separate directories based on artifact name. A single artifact is extracted into the current directory.

```sh
$ gh run download <run-id>
$ gh run download <run-id> -n <name>
$ gh run download -n <name1> -n <name2>    # across all runs
$ gh run download                            # interactive
$ gh run download <run-id> -D ./artifacts   # custom directory
$ gh run download <run-id> -p '*.zip'       # glob pattern
```

| Flag | Description |
|------|-------------|
| `-D, --dir string` | Download directory (default `.`) |
| `-n, --name stringArray` | Download artifacts matching given names |
| `-p, --pattern stringArray` | Download artifacts matching glob patterns |

### `gh run rerun`

Rerun an entire run, only failed jobs, or a specific job.

```sh
$ gh run rerun <run-id>
$ gh run rerun <run-id> --failed     # only failed jobs + dependencies
$ gh run rerun <run-id> --job 456789 # specific job + dependencies
$ gh run rerun <run-id> --debug      # with debug logging
```

| Flag | Description |
|------|-------------|
| `-d, --debug` | Rerun with debug logging |
| `--failed` | Rerun only failed jobs (including dependencies) |
| `-j, --job string` | Rerun a specific job ID (including dependencies) |

> **Job ID gotcha**: The `--job` flag uses the `databaseId`, not the number shown in the browser URL. Get correct job IDs with:
> ```sh
> gh run view <run-id> --json jobs --jq '.jobs[] | {name, databaseId}'
> ```

### `gh run cancel`

Cancel a workflow run.

```sh
$ gh run cancel <run-id>
$ gh run cancel <run-id> --force
```

| Flag | Description |
|------|-------------|
| `--force` | Force cancel |

### `gh run delete`

Delete a workflow run.

```sh
$ gh run delete
$ gh run delete 12345
```

## Cache Commands

### `gh cache list`

List GitHub Actions caches. Aliased as `gh cache ls`.

```sh
$ gh cache list
$ gh cache list --repo cli/cli
$ gh cache list --sort last_accessed_at --order asc
$ gh cache list --key key-prefix
$ gh cache list --ref refs/heads/<branch-name>
$ gh cache list --ref refs/pull/<pr-number>/merge
```

| Flag | Description |
|------|-------------|
| `-k, --key string` | Filter by cache key prefix |
| `-L, --limit int` | Max caches (default 30) |
| `-O, --order string` | Order: `asc`, `desc` (default `desc`) |
| `-r, --ref string` | Filter by ref (`refs/heads/<branch>` or `refs/pull/<number>/merge`) |
| `-S, --sort string` | Sort: `created_at`, `last_accessed_at`, `size_in_bytes` (default `last_accessed_at`) |
| `--json` / `--jq` / `--template` | JSON output |

**JSON fields**: `createdAt`, `id`, `key`, `lastAccessedAt`, `ref`, `sizeInBytes`, `version`.

### `gh cache delete`

Delete GitHub Actions caches. Requires `repo` scope.

```sh
$ gh cache delete 1234                    # by id
$ gh cache delete cache-key               # by key
$ gh cache delete cache-key --ref refs/heads/feature-branch
$ gh cache delete --all                   # all caches
$ gh cache delete --all --ref refs/pull/<PR>/merge  # all for a ref
$ gh cache delete --all --succeed-on-no-caches      # exit 0 if none
```

| Flag | Description |
|------|-------------|
| `-a, --all` | Delete all caches (can combine with `--ref`) |
| `-r, --ref string` | Delete by key and ref |
| `--succeed-on-no-caches` | Exit 0 if no caches found (requires `--all`) |

## Help Topic: `gh actions`

The `gh actions` help topic provides an overview of Actions-related commands:

```
gh run list      → List recent workflow runs
gh run view      → View details for a run or job
gh run watch     → Watch a run while it executes
gh run rerun     → Rerun a failed run
gh run download  → Download artifacts from runs

gh workflow list    → List workflow files
gh workflow view    → View details for a workflow
gh workflow enable  → Enable a workflow
gh workflow disable → Disable a workflow
gh workflow run     → Trigger a workflow_dispatch run

gh cache list   → List Actions caches
gh cache delete → Delete Actions caches
```

## Related

- For the `gh api` command (used for advanced Actions API queries), see [api.md](api.md).
- For JSON output formatting, see [output-formatting.md](output-formatting.md).
- For repository targeting (`--repo`), see [repo.md](repo.md).
