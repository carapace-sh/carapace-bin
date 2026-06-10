# Git Hooks

The hook system — client-side and server-side hooks that run automatically on Git events. Covers hook types, execution model, environment, and integration with external tools.

> **Source of truth**: <https://git-scm.com/docs/githooks>, <https://git-scm.com/book/en/v2/Git-Customizing-Git-Git-Hooks>. For **configuration**, see [config.md](config.md). For **bash execution model**, see the **bash** skill → `references/execution.md`.

## Hook Overview

Hooks are executable scripts that Git runs at specific points in its workflow. They can validate, transform, or react to Git operations.

### Hook Locations

|| Location | Description |
||----------|-------------|
|| `.git/hooks/` | Default hook directory |
|| `core.hooksPath` | Global override (e.g., `~/.config/git/hooks`) |
|| `$(git --exec-path)/hooks/` | Template hooks (copied on `git init`) |

When `core.hooksPath` is set, `.git/hooks/` is **completely ignored**.

### Hook Discovery

1. If `core.hooksPath` is set, use that directory
2. Otherwise, use `.git/hooks/`
3. Git looks for `<hook-name>` (no extension) — the file must be executable

### Template Hooks

On `git init` or `git clone`, Git copies template hooks from `$(git --exec-path)/templates/hooks/`. These are sample files with `.sample` extension — they must be renamed (removing `.sample`) and made executable to activate.

## Client-Side Hooks

### commit-msg

**When**: After the user edits the commit message, before the commit is finalized.

**Purpose**: Validate or modify the commit message.

**Input**: Path to a temporary file containing the commit message.

**Exit code**: Non-zero aborts the commit.

```bash
#!/bin/sh
# Ensure commit message starts with a ticket number
if ! grep -qE '^[A-Z]+-[0-9]+' "$1"; then
    echo "Commit message must start with a ticket number (e.g., PROJ-123)"
    exit 1
fi
```

The hook can **modify** the message file in-place.

### pre-commit

**When**: Before the commit is created (after `git commit`, before any other hooks).

**Purpose**: Validate the working tree state, run linters, prevent bad commits.

**Input**: None (no arguments, no stdin).

**Exit code**: Non-zero aborts the commit.

```bash
#!/bin/sh
# Run linter before commit
npm run lint || exit 1
```

Common uses:
- Run linters and formatters
- Check for debug statements
- Validate file sizes
- Prevent commits to protected branches
- Run tests

### prepare-commit-msg

**When**: After the default message is prepared, before the editor is opened.

**Purpose**: Modify the default commit message (e.g., add ticket number from branch name).

**Input**: `<file> <source> <commit>`

|| Source | Description |
||--------|-------------|
|| `message` | `-m` or `-F` option |
|| `template` | `template` config or `--template` |
|| `merge` | Merge commit (`.git/MERGE_MSG`) |
|| `squash` | Squash commit (`.git/SQUASH_MSG`) |
|| `commit` | `--amend` or `--fixup` |

```bash
#!/bin/sh
# Add branch name to commit message
BRANCH=$(git branch --show-current)
if [ -n "$BRANCH" ]; then
    TICKET=$(echo "$BRANCH" | grep -oE '[A-Z]+-[0-9]+')
    if [ -n "$TICKET" ]; then
        sed -i "1s/^/$TICKET /" "$1"
    fi
fi
```

### pre-push

**When**: Before `git push` sends data to the remote.

**Purpose**: Validate the push (check tests, prevent pushing to protected branches).

**Input**: `<remote> <url>` on stdin, one line per ref being pushed: `<local-oid> <remote-oid> <ref>`.

**Exit code**: Non-zero aborts the push.

```bash
#!/bin/sh
while read local_oid remote_oid ref; do
    # Prevent force push to main
    if [ "$ref" = "refs/heads/main" ] && [ "$local_oid" != "0000000000000000000000000000000000000000" ]; then
        if [ "$(git merge-base "$remote_oid" "$local_oid")" != "$remote_oid" ]; then
            echo "Force push to main is not allowed"
            exit 1
        fi
    fi
done
```

The zero OID (`000...000`) indicates a delete (local) or new ref (remote).

### pre-rebase

**When**: Before `git rebase` starts.

**Purpose**: Prevent rebasing of certain branches.

**Input**: `<upstream> <branch>` (branch is optional).

**Exit code**: Non-zero aborts the rebase.

### post-checkout

**When**: After `git checkout` or `git switch` completes.

**Purpose**: Set up the working tree, update dependencies, notify.

**Input**: `<prev-head> <new-head> <flag>`

|| Flag | Meaning |
||------|---------|
|| 1 | Branch checkout |
|| 0 | File checkout (e.g., `git checkout -- <path>`) |

```bash
#!/bin/sh
# Reinstall dependencies after branch switch
if [ "$3" = "1" ]; then
    npm install
fi
```

### post-commit

**When**: After a commit is created.

**Purpose**: Notification, logging, CI triggers.

**Input**: None.

**Exit code**: Ignored (commit already happened).

### post-merge

**When**: After `git merge` completes (not for fast-forward).

**Purpose**: Update dependencies after merge.

**Input**: Whether the merge was a squash merge (`1` for squash, `0` otherwise).

### post-rewrite

**When**: After a command rewrites commits (rebase, amend, filter-branch).

**Purpose**: Update references, notify.

**Input**: `<rewrite-type>` as argument, then `<old-oid> <new-oid> <extra-info>` on stdin.

|| Rewrite type | Command |
||-------------|---------|
|| `rebase` | `git rebase` |
|| `amend` | `git commit --amend` |

### pre-applypatch / applypatch-msg

**When**: Before/after `git am` applies a patch.

**Purpose**: Validate patch messages.

### post-applypatch

**When**: After `git am` applies a patch.

**Purpose**: Notification after patch application.

## Server-Side Hooks

### pre-receive

**When**: Before accepting a push (once per push, not per ref).

**Purpose**: Validate the entire push — reject if any ref is invalid.

**Input**: `<old-oid> <new-oid> <ref>` on stdin, one line per ref being updated.

**Exit code**: Non-zero rejects the **entire** push.

```bash
#!/bin/sh
while read old_oid new_oid ref; do
    # Prevent force push to main
    if [ "$ref" = "refs/heads/main" ]; then
        if [ "$old_oid" != "0000000000000000000000000000000000000000" ]; then
            if [ "$(git merge-base "$old_oid" "$new_oid")" != "$old_oid" ]; then
                echo "Force push to main is not allowed"
                exit 1
            fi
        fi
    fi
    # Prevent deleting main
    if [ "$ref" = "refs/heads/main" ] && [ "$new_oid" = "0000000000000000000000000000000000000000" ]; then
        echo "Deleting main is not allowed"
        exit 1
    fi
done
```

### update

**When**: Before accepting a push (once per ref being updated).

**Purpose**: Validate individual ref updates.

**Input**: `<ref> <old-oid> <new-oid>` as arguments.

**Exit code**: Non-zero rejects that specific ref (other refs in the push can still succeed).

```bash
#!/bin/sh
ref="$1"
old_oid="$2"
new_oid="$3"

if [ "$ref" = "refs/heads/main" ]; then
    # Only allow fast-forward to main
    if [ "$old_oid" != "0000000000000000000000000000000000000000" ]; then
        merge_base=$(git merge-base "$old_oid" "$new_oid")
        if [ "$merge_base" != "$old_oid" ]; then
            echo "Only fast-forward pushes to main are allowed"
            exit 1
        fi
    fi
fi
```

### post-receive

**When**: After a push is accepted and all refs are updated.

**Purpose**: Notification, CI triggers, deployment.

**Input**: `<old-oid> <new-oid> <ref>` on stdin.

**Exit code**: Ignored (push already accepted).

```bash
#!/bin/sh
while read old_oid new_oid ref; do
    if [ "$ref" = "refs/heads/main" ]; then
        # Trigger deployment
        /usr/local/bin/deploy.sh
    fi
done
```

### post-update

**When**: After a push is accepted (similar to post-receive but for each ref).

**Input**: List of ref names as arguments.

**Purpose**: Update server info for dumb HTTP protocol.

### push-to-checkout

**When**: When a push updates the branch that HEAD points to (on a non-bare repo).

**Purpose**: Update the working tree to match the pushed ref.

**Input**: `<ref> <old-oid> <new-oid>` on stdin.

### reference-transaction

**When**: When a batch of refs is about to be committed (e.g., during push).

**Purpose**: Hook into the ref transaction system.

**Input**: `<state> <ref-count>` as arguments, then `<ref> <old-oid> <new-oid>` on stdin.

|| State | Meaning |
||-------|---------|
|| `prepared` | Refs are about to be committed |
|| `committed` | Refs have been committed |
|| `aborted` | Transaction was aborted |

## Hook Environment

### Environment Variables

Hooks receive these environment variables:

|| Variable | Available in | Description |
||----------|-------------|-------------|
|| `GIT_DIR` | All hooks | Path to .git directory |
|| `GIT_QUARANTINE_PATH` | pre-receive | Path to quarantined objects (not yet in store) |
|| `GIT_PUSH_OPTION_COUNT` | pre-receive, update | Number of push options |
|| `GIT_PUSH_OPTION_<n>` | pre-receive, update | Push option value |
|| `GIT_REF_PATH` | reference-transaction | Ref being updated |
|| `GIT_AUTHOR_NAME/EMAIL/DATE` | commit-msg, pre-commit | Commit author info |
|| `GIT_COMMITTER_NAME/EMAIL/DATE` | commit-msg, pre-commit | Committer info |

### Quarantine Environment

In `pre-receive`, pushed objects are in a **quarantine directory** — they are not yet in the main object store. This means:

- `git cat-file` can access quarantined objects
- `git fsck` can verify quarantined objects
- Objects are only moved to the main store after all hooks pass

### Hook Execution

- Hooks run in the **repository root** (working tree for client hooks, `.git/` for server hooks on bare repos)
- Hooks inherit the **user's environment** (client) or **server's environment** (server)
- Hooks are run by the **shell** (`/bin/sh`) if they have a shebang, or directly if executable
- Hooks run **synchronously** — Git waits for them to complete
- There is **no timeout** — a hung hook blocks the operation indefinitely

### Hook Chaining

Git runs hooks in a fixed order. If any hook exits non-zero, the operation is aborted:

```
pre-commit → prepare-commit-msg → commit-msg → (commit) → post-commit
```

```
pre-rebase → (rebase starts) → post-rewrite
```

```
pre-push → (push starts) → pre-receive → update → (refs updated) → post-receive → post-update
```

## Hook Management Tools

### core.hooksPath

Set a global hooks directory:

```bash
git config --global core.hooksPath ~/.config/git/hooks
```

This replaces `.git/hooks/` entirely. All repositories use the same hooks.

### Husky

A Node.js-based hook manager:

```bash
npm install husky --save-dev
npx husky init
```

Creates hooks in `.husky/` and sets `core.hooksPath` to `.husky`.

### Lefthook

A Go-based hook manager with parallel execution:

```bash
# lefthook.yml
pre-commit:
  parallel: true
  commands:
    lint:
      run: npm run lint
    test:
      run: npm test
```

### pre-commit Framework

A Python-based framework for managing pre-commit hooks:

```bash
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.5.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
```

## Edge Cases and Known Issues

- **Hooks are not cloned**: `.git/hooks/` is not transferred by `git clone`. Use a hook manager (husky, lefthook) or `core.hooksPath`.
- **Shebang matters**: Hooks without a shebang are run by `/bin/sh`. Always include a shebang for other interpreters.
- **Executable bit**: Hooks must be executable (`chmod +x`). Git silently skips non-executable hooks.
- **Windows**: On Windows, hooks with `.sh` extension are run by Git Bash. Without extension, they must be executable.
- **Quarantine and pre-receive**: Objects in quarantine are not accessible by all Git commands. Use `GIT_QUARANTINE_PATH` to locate them.
- **Hook output**: stdout/stderr from hooks is shown to the user. Use this for error messages.
- **No hook for `git stash`**: There is no pre-stash or post-stash hook.
- **`--no-verify`**: Most client-side hooks can be bypassed with `git commit --no-verify` or `git push --no-verify`. Server-side hooks cannot be bypassed by the client.

## References

- <https://git-scm.com/docs/githooks>
- <https://git-scm.com/book/en/v2/Git-Customizing-Git-Git-Hooks>

## Related Skills

- For configuration, see [config.md](config.md)
- For bash execution model (subshells, signals, traps), see the **bash** skill → `references/execution.md`
- For rebase hooks (post-rewrite), see [rebase.md](rebase.md)
