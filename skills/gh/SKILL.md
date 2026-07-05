---
name: gh
description: >
  Use when working with the GitHub CLI (`gh`) — authentication, repositories, pull requests,
  issues, releases, GitHub Actions (workflows, runs, caches), the gh API command, extensions,
  secrets and variables, search, gists, SSH/GPG keys, labels, rulesets, codespaces, browsing,
  configuration, aliases, JSON output formatting, and environment variables. Triggers on:
  "gh", "gh auth", "gh repo", "gh pr", "gh issue", "gh release", "gh run", "gh workflow",
  "gh api", "gh extension", "gh secret", "gh variable", "gh search", "gh gist", "gh label",
  "gh codespace", "gh browse", "gh status", "gh config", "gh alias", "gh completion",
  "gh ssh-key", "gh gpg-key", "gh ruleset", "gh cache", "gh org", "GH_TOKEN", "GH_REPO",
  "GH_HOST", "gh pr create", "gh pr checkout", "gh pr merge", "gh repo clone", "gh repo create",
  "gh issue create", "gh release create", "gh workflow run", "gh run list", "gh api graphql",
  "gh extension install", "gh secret set", "gh search repos".
user-invocable: true
---

# GitHub CLI (`gh`) In-Depth Reference

Comprehensive reference for the [GitHub CLI](https://cli.github.com) — the official command-line tool for working with GitHub from the terminal. Covers all command groups, flags, argument formats, environment variables, output formatting, and the extension system.

## Data Flow

```
Local git repo
  → gh repo clone / gh repo create
    → gh pr create / gh issue create
      → gh pr checkout / gh pr merge / gh pr review
        → gh run watch / gh release create
          → gh browse (open in browser)
```

```
Authentication:
  gh auth login → credential store (or GH_TOKEN env)
    → gh auth status / gh auth switch (multi-account)
      → gh auth refresh (add/remove scopes)
        → gh auth setup-git (credential helper)
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| auth, login, logout, refresh, setup-git, status, switch, token, authentication, credential store, OAuth, scopes, GH_TOKEN, GITHUB_TOKEN, GH_ENTERPRISE_TOKEN, hostname, multi-account, web flow, with-token, insecure-storage, git-protocol, ssh key upload, GitHub Enterprise, GHES | [references/auth.md](references/auth.md) |
| config, configuration, config get, config set, config list, clear-cache, git_protocol, editor, prompt, pager, browser, http_unix_socket, color_labels, accessible_colors, accessible_prompter, spinner, telemetry, per-host config, environment variables, GH_HOST, GH_REPO, GH_EDITOR, GH_BROWSER, GH_PAGER, GH_DEBUG, GH_FORCE_TTY, GH_CONFIG_DIR, GH_NO_UPDATE_NOTIFIER, NO_COLOR, GLAMOUR_STYLE, exit codes, completion, shell completion | [references/config.md](references/config.md) |
| repo, repository, repo create, repo clone, repo fork, repo view, repo edit, repo list, repo sync, repo delete, repo archive, repo unarchive, repo rename, repo set-default, gitignore, license, deploy-key, autolink, OWNER/REPO, template repo, --source, --push, --private, --public, --internal, visibility, --add-readme | [references/repo.md](references/repo.md) |
| pr, pull request, pr create, pr list, pr view, pr checkout, pr edit, pr merge, pr close, pr reopen, pr review, pr diff, pr comment, pr ready, pr lock, pr unlock, pr status, pr checks, pr update-branch, pr revert, draft, --fill, --fill-first, --fill-verbose, --base, --head, --reviewer, --assignee, --label, --milestone, --project, --template, --dry-run, --no-maintainer-edit, @me, @copilot, merge strategy, squash, rebase, auto-merge, delete-branch | [references/pr.md](references/pr.md) |
| issue, issue create, issue list, issue view, issue edit, issue close, issue reopen, issue comment, issue delete, issue develop, issue transfer, issue pin, issue unpin, issue lock, issue unlock, issue status, --assignee, --label, --milestone, --project, --template, --type, --parent, sub-issue, blocked-by, blocking, duplicate-of, --reason | [references/issue.md](references/issue.md) |
| api, GraphQL, REST API, endpoint, placeholder, {owner}, {repo}, {branch}, field, raw-field, --method, --paginate, --slurp, --header, --input, --cache, --preview, --silent, --verbose, --include, jq, template, workflow_dispatch, nested parameters, key[subkey], key[] | [references/api.md](references/api.md) |
| actions, workflow, run, cache, workflow list, workflow run, workflow view, workflow enable, workflow disable, run list, run view, run watch, run download, run rerun, run cancel, run delete, cache list, cache delete, --ref, --field, --raw-field, --json, --status, --event, --branch, --commit, --job, --log, --log-failed, --exit-status, --attempt, artifacts, workflow_dispatch | [references/actions.md](references/actions.md) |
| release, release create, release list, release view, release download, release upload, release delete, release edit, release verify, release verify-asset, release delete-asset, --generate-notes, --notes-from-tag, --notes-start-tag, --target, --title, --draft, --prerelease, --latest, --verify-tag, --discussion-category, --fail-on-no-commits, --cleanup-tag, immutable release, asset label, --archive, --pattern, --clobber | [references/release.md](references/release.md) |
| extension, gh extension, ext, extension install, extension list, extension create, extension remove, extension upgrade, extension browse, extension search, extension exec, gh- prefix, binary extension, script extension, --pin, --force, --precompiled, local extension, upgrade notice, GH_NO_EXTENSION_UPDATE_NOTIFIER | [references/extension.md](references/extension.md) |
| secret, variable, secret set, secret list, secret delete, variable set, variable get, variable list, variable delete, --app, --org, --env, --env-file, --visibility, --repos, --no-repos-selected, --no-store, --user, actions, dependabot, codespaces, agents, repository secret, organization secret, environment secret, user secret, dotenv | [references/secret-variable.md](references/secret-variable.md) |
| search, search repos, search issues, search prs, search code, search commits, --owner, --language, --topic, --stars, --forks, --visibility, --archived, --created, --updated, --sort, --order, --limit, --match, --assignee, --author, --label, --state, --draft, --merged, --review, --reviewed-by, --review-requested, --checks, --base, --head, --include-prs, GitHub search syntax, --web | [references/search.md](references/search.md) |
| gist, gpg-key, ssh-key, label, org, ruleset, gist create, gist list, gist view, gist clone, gist edit, gist delete, gist rename, --public, --desc, ssh-key add, ssh-key list, ssh-key delete, authentication key, signing key, gpg-key add, gpg-key list, gpg-key delete, label create, label list, label edit, label delete, label clone, org list, ruleset list, ruleset view, ruleset check, --parents, --org, --default | [references/gist-keys.md](references/gist-keys.md) |
| browse, status, codespace, cs, browse --actions, browse --settings, browse --wiki, browse --releases, browse --projects, browse --blame, browse --branch, browse --commit, --no-browser, codespace create, codespace list, codespace code, codespace ssh, codespace delete, codespace edit, codespace logs, codespace ports, codespace rebuild, codespace stop, codespace view, codespace cp, codespace jupyter, --machine, --location, --idle-timeout, --retention-period, --devcontainer-path | [references/browse-status.md](references/browse-status.md) |
| JSON, --json, --jq, --template, Go template, formatting, tablerow, tablerender, color, autocolor, timeago, timefmt, truncate, hyperlink, pluck, join, Sprig, contains, hasPrefix, hasSuffix, regexMatch, pretty-print, JSON fields | [references/output-formatting.md](references/output-formatting.md) |

## Quick Guide

- **How do I authenticate with GitHub?** → [references/auth.md](references/auth.md)
- **How do I configure gh settings and environment variables?** → [references/config.md](references/config.md)
- **How do I clone, create, or fork a repository?** → [references/repo.md](references/repo.md)
- **How do I create and manage pull requests?** → [references/pr.md](references/pr.md)
- **How do I create and manage issues?** → [references/issue.md](references/issue.md)
- **How do I make raw API requests or GraphQL queries?** → [references/api.md](references/api.md)
- **How do I work with GitHub Actions workflows and runs?** → [references/actions.md](references/actions.md)
- **How do I create and manage releases?** → [references/release.md](references/release.md)
- **How do I install or create gh extensions?** → [references/extension.md](references/extension.md)
- **How do I manage secrets and variables?** → [references/secret-variable.md](references/secret-variable.md)
- **How do I search across GitHub?** → [references/search.md](references/search.md)
- **How do I manage gists, keys, labels, orgs, or rulesets?** → [references/gist-keys.md](references/gist-keys.md)
- **How do I open things in the browser or manage codespaces?** → [references/browse-status.md](references/browse-status.md)
- **How do I format JSON output with jq or Go templates?** → [references/output-formatting.md](references/output-formatting.md)
- **How do I switch between multiple GitHub accounts?** → [references/auth.md](references/auth.md)
- **How do I use `--json` and `--jq` flags?** → [references/output-formatting.md](references/output-formatting.md)
- **How do I set up shell completion?** → [references/config.md](references/config.md)
- **What environment variables does gh recognize?** → [references/config.md](references/config.md)

## Cross-Project References

- For **Git internals** (branches, refs, remotes, config, hooks), use the **git** skill.
- For **carapace** completion integration (the `gh` completer in carapace-bin, specs, actions, macros), see the **carapace** skill.
- For **cobra** command structure (gh is built with cobra), see the **cobra** skill.
