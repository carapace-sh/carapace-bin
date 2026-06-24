# Git Completer Verification Report

**Date:** 2026-06-24
**Git source:** ab776a62a7 (Git 2.55-rc2) — `/home/rsteube/Documents/development/github/git/git`
**Carapace-bin:** 417c33b75 — `/home/rsteube/Documents/development/github/carapace-sh/carapace-bin`
**System git:** v2.54.0

**Status:** ✅ = Fixed in this session

---

## Summary

The git completer is generally comprehensive, covering ~150 commands. The main issues fall into these categories:

1. **6 missing subcommands** (built-in commands not present in the completer at all)
2. **1 flag typo** (`--no-advise` should be `--no-advice`)
3. **Missing subcommands of existing commands** (e.g. `git stash export/import`, `git remote set-url --delete`)
4. **Missing flags on existing commands** (various, listed per command below)
5. **Flag name mismatches** (e.g. `--heads` vs `--branches` in `git ls-remote`)

---

## 1. Missing Subcommands

These built-in git commands have no completer at all:

| Command | Type | Notes |
|---------|------|-------|
| `git backfill` | builtin | ✅ Added with `--min-batch-size`, `--sparse` |
| `git diff-pairs` | builtin | ✅ Added with `common.AddDiffFlags` |
| `git format-rev` | builtin | ✅ Added with `--format`, `--stdin-mode`, `--notes`, `-z/--null`, `--null-input`, `--null-output` |
| `git last-modified` | builtin | ✅ Added with `--max-depth`, `--recursive/-r`, `--show-trees/-t`, `-z` |
| `git url-parse` | builtin | ✅ Added with `--component/-c` |
| `git scalar` | external (mainporcelain) | ✅ Added with subcommands: `clone`, `list`, `register`, `unregister`, `run`, `reconfigure`, `diagnose`, `delete`, `version` |

### Missing Subcommands of Existing Commands

| Parent Command | Missing Subcommand | Notes |
|----------------|-------------------|-------|
| `git stash` | `create` | ✅ Added (plumbing subcommand) |
| `git stash` | `store` | ✅ Added (plumbing subcommand) |
| `git stash` | `export` | ✅ Added with `--print`, `--to-ref` |
| `git stash` | `import` | ✅ Added (takes `<commit>` argument) |
| `git stash` | `save` | ✅ Added (deprecated alias for `push`) |

---

## 2. Flag Typo

| Command | Wrong | Correct | Source |
|---------|-------|---------|--------|
| `git` (root) | `--no-advise` | `--no-advice` | ✅ Fixed — `git.c:353` |

This is a genuine typo in `completers/common/git_completer/cmd/root.go:72`. The git source uses `--no-advice` with environment variable `GIT_ADVICE`.

---

## 3. Root-Level Flag Issues

| Issue | Details |
|-------|---------|
| `--list-cmds` incomplete values | ✅ Fixed — Added `builtins`, `config`, `deprecated` |

---

## 4. Missing Flags Per Command

### High-Priority (Common Commands)

#### `git add`
- ✅ `--auto-advance` — auto advance to next file during interactive hunk selection
- ✅ `--inter-hunk-context <n>` — show context between diff hunks (via `common.AddPatchContextFlags`)
- ✅ `--unified/-U <n>` — generate diffs with n lines context (via `common.AddPatchContextFlags`)

#### `git am`
- ✅ `--allow-empty` — allow empty commits
- ✅ `--empty` — how to handle empty patches (drop/keep/stop)
- ✅ `--retry` — retry failed patch application
- `--verify` — run pre-applypatch and post-applypatch hooks (only `--no-verify`/`-n` exists in git source)
- `-b` — short form alias (git 2.54+)

#### `git apply`
- `--add` — not in git source (only `--no-add` exists)
- ✅ `--allow-empty` — do not error on empty diff
- ✅ `--ours`/`--theirs`/`--union` — conflict resolution modes
- ✅ `--quiet/-q` — suppress report

#### `git archive`
- ✅ `--add-virtual-file` — add virtual file to archive (new in Git 2.53)
- ✅ `--mtime` — set modification time

#### `git bisect`
- `--first-parent` — follow only first parent commits
- `--term-` — custom term prefix (incomplete flag name parsing)

#### `git blame`
- `--diff-algorithm` — choose diff algorithm

#### `git cherry-pick`
- `--commit` — opposite of --no-commit
- `--empty` — how to handle empty commits (drop/keep/stop)

#### `git clone`
- `--also-filter-submodules` — also apply partial clone filter to submodules
- `--bundle-uri` — fetch from bundle URI before remote
- `--checkout` — opposite of --no-checkout
- `--hardlinks` — use hardlinks when possible
- ✅ `--ref-format` — specify reference format (new in Git 2.55, with completion values: files/flat/tree)
- ✅ `--revision` — specify revision to clone (new in Git 2.55)
- `--tags` — opposite of --no-tags

#### `git commit`
- ✅ `--inter-hunk-context <n>` — show context between diff hunks (via `common.AddPatchContextFlags`)
- `--post-rewrite` — insert note about rewrite in commit log (used by git-filter-repo)
- ✅ `--unified/-U <n>` — generate diffs with n lines context (via `common.AddPatchContextFlags`)
- `--verify` — opposite of --no-verify

#### `git fetch`
- `--auto` — not in git source (has `--auto-gc`/`--auto-maintenance` instead)
- ✅ `--porcelain` — machine-readable output (new in Git 2.55)
- ✅ `--refetch` — re-fetch all objects

#### `git format-patch`
- `--binary` — include binary diffs
- `--commit-list-format` — specify format for commit list (new)
- `--description-file` — description file for cover letter
- `--force-in-body-from` — force In-Reply-To in body

#### `git gc`
- `--cruft` — pack cruft objects separately
- `--detach` — detach from terminal
- `--expire-to` — expire to specific pack
- `--max-cruft-size` — max size for cruft pack

#### `git grep`
- ✅ `--function-context/-W` — show surrounding function context
- `--index` — not in git source (has `--cached` and `--no-index` instead)
- ✅ `--max-count/-m` — limit number of matches per file

#### `git history`
- `--dry-run` — preview history rewrite
- `--update-refs` — update branches pointing to rewritten commits

#### `git hook`
- `--allow-unknown-hook-name` — allow unknown hooks
- `--ignore-missing` — ignore missing hooks
- `--show-scope` — show hook scope
- `--to-stdin` — provide stdin to hook

#### `git init`
- ✅ `--ref-format` — specify reference format (new in Git 2.55, with completion values: files/flat/tree)

#### `git log`
- `--clear-decorations` — clear all decorations
- `--quiet/-q` — suppress diff output

#### `git merge`
- ✅ `--compact-summary` — show compact summary of merge results
- `--verify` — not in git source (has `--no-verify` and `--verify-signatures`)

#### `git mv`
- ✅ `--sparse` — allow moving into sparse-checkout cone

#### `git notes`
- ✅ `--abort` — already in notes_merge.go
- ✅ `--allow-empty` — already in notes_add.go/notes_append.go/notes_edit.go
- ✅ `--commit` — already in notes_merge.go
- ✅ `--separator` — added to add/append/edit subcommands
- ✅ `--stripspace` — added to add/append/edit subcommands
- ✅ `-F`/`--file` — already in notes_add.go/notes_append.go; added to notes_edit.go
- ✅ `-q` — already in notes_merge.go

#### `git pull`
- ✅ `--compact-summary` — show compact summary
- ✅ `--verify` — run pre-merge/pre-rebase hooks

#### `git push`
- ✅ `--branches` — push all branches (with --mirror)
- `--verify` — not in git source (has `--no-verify`)

#### `git range-diff`
Missing many diff-related flags because range-diff uses the same diff option set as `git diff` but the completer doesn't share those flags. Missing include all the common diff flags like `--abbrev`, `--anchored`, `--binary`, `--stat`, `--patch`, `--unified`, etc.

#### `git rebase`
- ✅ `--ff` — opposite of --no-ff
- ✅ `--reset-author-date` — ignore author date
- ✅ `--stat` — opposite of --no-stat
- ✅ `--trailer` — add custom trailer
- ✅ `--update-refs` — update branches pointing to rebased commits
- ✅ `--verify` — opposite of --no-verify

#### `git reflog`
- `--all` — show reflogs from all refs
- `--dry-run` — dry run for expire
- `--expire` — expire time
- `--expire-unreachable` — expire unreachable entries
- `--rewrite` — rewrite entries
- `--single-worktree` — only show current worktree
- `--stale-fix` — fix stale reflog entries
- `--updateref` — update ref with final value
- `--verbose` — verbose output
- `-n` — limit count

#### `git refs`
Many flags missing: `--all`, `--auto`, `--contains`, `--count`, `--dry-run`, `--exclude`, `--format`, `--include`, `--include-root-refs`, `--merged`, `--perl`, `--points-at`, `--python`, `--ref-format`, `--shell`, `--sort`, `--start-after`, `--stdin`, `--strict`, `--tcl`, `--verbose`

#### `git remote`
Flags that are actually subcommand-specific but appear in the top-level completer: `--add`, `--all`, `--auto`, `--delete/-d`, `--dry-run`, `--mirror`, `--progress`, `--prune`, `--push`, `--tags`

#### `git repack`
- `--combine-cruft-below-size` — combine cruft packs below threshold
- `--filter` — partial clone filter
- `--filter-to` — filter destination
- `--max-cruft-size` — max size for cruft pack
- `--name-hash-version` — name hash version (new)
- `--path-walk` — walk paths for repacking (new)

#### `git reset`
- `--auto-advance` — auto advance in interactive mode
- `--inter-hunk-context <n>` — show context between diff hunks
- `--refresh` — refresh index
- `--unified/-U <n>` — generate diffs with n lines context

#### `git restore`
- `--inter-hunk-context <n>` — show context between diff hunks
- `--unified/-U <n>` — generate diffs with n lines context

#### `git revert`
- `--commit` — opposite of --no-commit
- `--reference` — use reference to resolve conflicts

#### `git rm`
- ✅ `--sparse` — allow removing files outside sparse-checkout cone

#### `git stash push`
- `--all/-a` — include ignored files
- `--include-untracked/-u` — include untracked files
- `--index` — try to restore index
- `--keep-index/-k` — keep index changes
- `--message/-m` — stash message
- `--only-untracked` — only untracked files
- `--patch/-p` — interactive mode
- `--pathspec-file-nul` — NUL-separated pathspec
- `--pathspec-from-file` — read pathspec from file
- `--print` — print stash object ID
- `--quiet/-q` — quiet mode
- `--staged/-S` — stash only staged changes
- `--to-ref` — save to specific ref

#### `git status`
- `--null` — NUL-terminated output (completer has `--z` instead; git uses `--null` with `-z` as shorthand)
- `-M` — show move/copy detection info

#### `git submodule`
The completer's `submodule` command has a single flat structure, but git's submodule uses per-subcommand flags. Key missing flags across subcommands: `--all`, `--branch/-b`, `--cached`, `--checkout`, `--default`, `--files`, `--filter`, `--force/-f`, `--init`, `--merge`, `--name`, `--quiet`, `--rebase`, `--recommend-shallow`, `--recursive`, `--reference`, `--remote`, `--single-branch`, `--summary-limit`

#### `git switch`
Missing `--no-*` toggle forms that have positive forms in completer: `--no-conflict`, `--no-create`, `--no-detach`, `--no-discard-changes`, `--no-force`, `--no-force-create`, `--no-ignore-other-worktrees`, `--no-merge`, `--no-orphan`, `--no-overwrite-ignore`, `--no-quiet`, `--no-recurse-submodules`

#### `git shortlog`
- ✅ `--pretty` — added pretty-print format option (was missing entirely)
- Note: Could not use `common.AddCommitLimitingOptions`/`AddHistorySimplificationOptions` because shortlog's `-c`/`--committer` and `-n`/`--numbered` have different semantics than rev-list's versions

#### `git worktree add`
- `--checkout` — already present in completer
- `--detach/-d` — already present in completer
- `--lock` — already present in completer
- ✅ `--orphan` — create unborn branch
- `--reason` — already present in completer
- `-B` — already present in completer
- ✅ `--relative-paths` — use relative paths for worktrees

### Lower-Priority (Plumbing/Uncommon Commands)

#### `git multi-pack-index`
- `--batch-size`, `--bitmap`, `--incremental`, `--preferred-pack`, `--refs-snapshot`, `--stdin-packs`

#### `git name-rev`
- `--undefined` — show undefined commits

#### `git pack-objects`
Many flags missing: `--exclude-promisor-objects-best-effort`, `--indexed-objects`, `--keep-true-parents`, `--name-hash-version`, `--path-walk`, `--quiet`, `--reflog`, `--reuse-delta`, `--reuse-object`, `--uri-protocol`, `--use-bitmap-index`, `--write-bitmap-index`

#### `git send-pack`
Many flags missing: `--force-if-includes`, `--force-with-lease`, `--helper-status`, `--mirror`, `--progress`, `--quiet`, `--remote`, `--stateless-rpc`, `-f/-n/-q`

#### `git show-branch`
- `--name` — show branch names only

#### `git update-index`
- `--clear-resolve-undo`, `--force-write-index`

#### `git update-ref`
- `--batch-updates` — batch update from stdin
- `-0` — NUL-delimited input

#### `git verify-pack`
- `--object-format` — verify object format

---

## 5. Flag Name Mismatches

| Command | Completer | Git | Notes |
|---------|-----------|-----|-------|
| `git ls-remote` | `--heads` | ✅ Fixed: `--branches/-b` | Git uses `--branches` as the primary name (since 2.54); `--heads` is an older alias |
| `git status` | `--z` | ✅ Fixed: `-z` / `--null` | Changed `BoolS("z", "z")` to `BoolP("null", "z")` |

---

## 6. Structural Issues

### `git archive` compression level flags
The completer defines 10 separate boolean flags `--0` through `--9` for compression levels. Git's actual syntax is a non-standard `-<n>` where n is 0-9. The boolean flag approach is incorrect — these should ideally be handled differently (or at least as a single integer-like flag).

### `git submodule` flat vs hierarchical flags
The completer defines all submodule flags at the top level, but git's submodule command has per-subcommand flags (e.g., `--force` means different things for `add` vs `update`). The completer should ideally use the hierarchical subcommand structure it already has (`submodule_add.go`, `submodule_update.go`, etc.) with flags on each subcommand.

### `git stash` missing subcommands
`git stash` is missing `export`, `import`, `create`, `store`, and `save` subcommands. The `export` and `import` subcommands are new in Git 2.55.

### `git config` flag handling
`git config` uses `NO_PARSEOPT` in the git source, meaning `git config -h` doesn't show parse-options output. The completer has many more flags than appear in `git config -h`, which is correct since it includes all the flags from the actual option definitions. However, the completer is missing `--all`, `--regexp`, `--url`, and `--value` which are present in git's source code (`builtin/config.c`).

### ✅ `git diff` family and `--merge-base`
✅ Added `--merge-base` to `diff.go`.

---

## 7. Verification Notes

- Comparison was done against **Git 2.55-rc2 source** (`git.c` commands[] array, `command-list.txt`, and `git <cmd> -h` output from system git v2.54)
- Flags prefixed with `--[no-]` in git's help output are boolean toggles; pflag auto-generates `--no-*` forms for BoolVarP flags, so missing `--no-*` forms are typically lower priority unless the positive form is present
- Some "missing" flags may be intentionally omitted (e.g., plumbing-only flags, deprecated flags, or flags that conflict with existing completion behavior)
- The `git diff` command and its relatives (range-diff, log, show, etc.) share common diff flags via `common/diff.go`, which means my initial automated comparison undercounted their flags

---

## 8. Common Flag Grouping Refactoring

### ✅ `log.go` — Refactored
- Removed inline flags that duplicate `AddCommitLimitingOptions` (`--ancestry-path`, `--dense`, `--full-history`, `--show-pulls`, `--simplify-by-decoration`, `--simplify-merges`, `--sparse`), `AddCommitOrderingOptions` (`--author-date-order`, `--date-order`, `--reverse`, `--topo-order`), and `AddBisectionHelperOptions` (`--bisect`)
- Now uses `common.AddBisectionHelperOptions`, `common.AddCommitOrderingOptions`, `common.AddHistorySimplificationOptions` in addition to the previously-used `AddCommitFormattingOptions`, `AddCommitLimitingOptions`, `AddDiffFlags`, `AddObjectTraversalOptions`

### ✅ `whatchanged.go` — Refactored
- Was only using inline flags with no common function calls
- Now uses `common.AddCommitFormattingOptions`, `common.AddCommitLimitingOptions`, `common.AddDiffFlags`, `common.AddObjectTraversalOptions`
- Removed `--quiet/-q` inline (now comes from `AddCommitLimitingOptions`)

### ⚠️ `shortlog.go` — Could NOT use common functions
- Shortlog has its own semantics for `-c`/`--committer` (BoolP, alias for `--group=committer`) vs rev-list's `--committer` (StringArray, pattern match)
- Shortlog's `-n`/`--numbered` conflicts with `AddCommitLimitingOptions`' `-n`/`--max-count`
- Added missing `--pretty` flag directly instead of via `AddCommitFormattingOptions`
- Kept all flags inline to avoid shorthand/type conflicts

### ✅ New `AddPatchContextFlags` in `common/diff.go`
- Added shared function providing `--inter-hunk-context` and `--unified/-U`
- Used by: `add.go`, `checkout.go`, `commit.go`, `reset.go`, `restore.go`
- Commands like `add`, `commit`, `checkout` only need these two context flags for `--patch` mode, not the full `AddDiffFlags` set
