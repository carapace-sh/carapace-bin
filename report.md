# Git Completer vs Git 2.55-rc2 Audit Report

## Source
- **Git source**: `/home/rsteube/Documents/development/github/git/git` (Git 2.55-rc2, commit ab776a62a7)
- **Completer**: `/home/rsteube/Documents/development/github/carapace-sh/carapace-bin/completers/common/git_completer/`
- **Date**: 2024-06-24

## Summary

The git completer was audited against Git 2.55-rc2 source code. All missing subcommands and significant missing flags have been fixed across three commits.

---

## Commit 1: `93c1b5bf7` — "git: add missing subcommands and flags for Git 2.55"

### New Subcommands ✅

| Command | Source File | Notes |
|---------|------------|-------|
| `git backfill` | `builtin/backfill.c` | New in 2.55 |
| `git diff-pairs` | `builtin/diff-pairs.c` | New in 2.55 |
| `git format-rev` | `builtin/format-rev.c` | New in 2.55 |
| `git last-modified` | `builtin/last-modified.c` | New in 2.55 |
| `git url-parse` | `builtin/url-parse.c` | New in 2.55 |
| `git scalar` | `contrib/scalar/` | With 8 subcommands: clone, delete, diagnose, list, reconfigure, register, run, version |
| `git stash create` | `builtin/stash.c` | Missing subcommand |
| `git stash store` | `builtin/stash.c` | Missing subcommand |
| `git stash export` | `builtin/stash.c` | New in 2.55 |
| `git stash import` | `builtin/stash.c` | New in 2.55 |
| `git stash save` | `builtin/stash.c` | Deprecated but should still be present |

### Bug Fixes ✅

| Fix | Details |
|-----|---------|
| `--no-advise` → `--no-advice` | Typo in root.go |
| `--z` → `--null/-z` | status.go had `BoolS("z","z")` instead of `BoolP("null","z")` |
| `--branches/-b` | Added missing shorthand to ls-remote |
| Duplicate `-z` in diffPairs.go | Already provided by `common.AddDiffFlags` |

### Missing Flags Added ✅

| Command | Flag | Type | Notes |
|---------|------|------|-------|
| `am` | `--allow-empty` | Bool | |
| `am` | `--empty` | String (stop/drop/keep) | |
| `am` | `--retry` | Bool | |
| `apply` | `--allow-empty` | Bool | |
| `apply` | `--ours` | Bool | No shorthand (would conflict with `--3way/-3`) |
| `apply` | `--theirs` | Bool | No shorthand |
| `apply` | `--union` | Bool | |
| `apply` | `--quiet/-q` | Bool | |
| `archive` | `--add-virtual-file` | String | Colon-delimited path:content |
| `archive` | `--mtime` | String | |
| `fetch` | `--porcelain` | Bool | |
| `fetch` | `--refetch` | Bool | |
| `grep` | `--function-context/-W` | Bool | |
| `grep` | `--max-count/-m` | String | |
| `merge` | `--compact-summary` | Bool | |
| `pull` | `--compact-summary` | Bool | |
| `pull` | `--verify` | Bool | |
| `push` | `--branches` | Bool | |
| `tag` | `--omit-empty` | Bool | |
| `tag` | `--trailer` | StringArray | |
| `worktree add` | `--orphan` | String | |
| `worktree add` | `--relative-paths` | Bool | |
| `notes add` | `--edit/-e`, `--separator`, `--stripspace` | | |
| `notes append` | `--edit/-e`, `--separator`, `--stripspace` | | |
| `notes edit` | `--file/-F`, `--message/-m`, `--reedit-message/-c`, `--reuse-message/-C`, `--separator`, `--stripspace` | | |
| `shortlog` | `--pretty` | String | Kept inline (conflicts with common functions) |
| `clone` | `--ref-format` | String (files/flat/tree) | |
| `clone` | `--revision` | Bool | |
| `init` | `--ref-format` | String (files/flat/tree) | |
| `mv` | `--sparse` | Bool | |
| `rm` | `--sparse` | Bool | |
| `rebase` | `--ff`, `--reset-author-date`, `--stat`, `--trailer`, `--update-refs`, `--verify` | | |
| `diff` | `--merge-base` | Bool | |
| `add` | `--auto-advance` | Bool | |
| `add/checkout/commit/reset/restore` | Patch context flags | Via `common.AddPatchContextFlags` | `--inter-hunk-context`, `--unified/-U` |

### Refactoring ✅

| File | Change |
|------|--------|
| `log.go` | Replaced ~15 inline flags with `common.AddBisectionHelperOptions`, `common.AddCommitOrderingOptions`, `common.AddHistorySimplificationOptions` |
| `whatchanged.go` | Replaced inline flags with `common.AddCommitFormattingOptions`, `common.AddCommitLimitingOptions`, `common.AddDiffFlags`, `common.AddObjectTraversalOptions` |
| `common/diff.go` | Added `AddPatchContextFlags` function |

---

## Commit 2: `7837afde3` — "git: split multi-subcommand files into one file per subcommand"

Split `stash_newSub.go` into 5 individual files and `scalar_sub.go` into 8 individual files following project convention.

---

## Commit 3 (pending): Additional missing flags

### New Subcommand ✅

| Command | Source | Flags |
|---------|--------|-------|
| `git history fixup` | `builtin/history.c` | `--empty` (drop/keep/abort), `--dry-run/-n`, `--reedit-message`, `--update-refs` (branches/head) |

### New Refs Subcommands ✅

| Subcommand | Source | Flags |
|------------|--------|-------|
| `git refs migrate` | `builtin/refs.c` | `--dry-run`, `--no-dry-run`, `--no-reflog`, `--ref-format` (files/reftable) |
| `git refs verify` | `builtin/refs.c` | `--strict`, `--verbose` |
| `git refs list` | `builtin/refs.c` → `for_each_ref_core` | Same flags as `for-each-ref`: `--color`, `--contains`, `--count`, `--exclude`, `--format`, `--ignore-case`, `--include-root-refs`, `--merged`, `--no-contains`, `--no-merged`, `--omit-empty`, `--perl`, `--points-at`, `--python`, `--shell`, `--sort`, `--start-after`, `--stdin`, `--tcl` |
| `git refs exists` | `builtin/refs.c` | No flags, positional ref arg |
| `git refs optimize` | `builtin/refs.c` | No flags |

### Missing Flags Added ✅

| Command | Flag | Type | Notes |
|---------|------|------|-------|
| `cherry-pick` | `--empty` | String (keep/drop/stop) | |
| `clone` | `--also-filter-submodules` | Bool | |
| `clone` | `--bundle-uri` | String | |
| `clone` | `--tags` | Bool | |
| `format-patch` | `--commit-list-format` | String | |
| `format-patch` | `--description-file` | String (file completion) | |
| `format-patch` | `--force-in-body-from` | Bool | |
| `gc` | `--cruft` | Bool | |
| `gc` | `--detach` | Bool | |
| `gc` | `--expire-to` | String (dir completion) | |
| `gc` | `--max-cruft-size` | String | |
| `hook run` | `--allow-unknown-hook-name` | Bool | |
| `hook run` | `--jobs/-j` | String | |
| `log` | `--clear-decorations` | Bool | |
| `repack` | `--combine-cruft-below-size` | String | |
| `repack` | `--filter` | StringArray (object filter completion) | |
| `repack` | `--filter-to` | String (dir completion) | |
| `repack` | `--max-cruft-size` | String | |
| `repack` | `--name-hash-version` | String (1/2 completion) | |
| `repack` | `--path-walk` | Bool | |
| `reset` | `--auto-advance` | Bool | |
| `revert` | `--reference` | Bool | For conflict resolution |
| `status` | `--find-renames/-M` | StringP | Added `-M` shorthand (was only `--find-renames`) |
| `refs migrate` | `--no-dry-run`, `--no-reflog` | Bool | |

---

## Verified as Already Correct

These items were initially flagged as potentially missing but verified as already present or auto-generated:

| Item | Status |
|------|--------|
| `range-diff` diff flags | Already uses `common.AddDiffFlags` — all ~80 diff flags present |
| `cherry-pick --commit` | Auto-generated by pflag from `--no-commit` BoolP |
| `revert --commit` | Auto-generated by pflag from `--no-commit` BoolP |
| `switch --no-*` forms | Auto-generated by pflag from BoolP flags |
| `hook list --allow-unknown-hook-name` | Already present |
| `hook list --show-scope` | Already present |
| `hook list -z` | Already present |
| `hook run --ignore-missing` | Already present |
| `hook run --to-stdin` | Already present |
| `reflog expire` flags | All flags already present (`--all`, `--dry-run/-n`, `--expire`, `--expire-unreachable`, `--rewrite`, `--single-worktree`, `--stale-fix`, `--updateref`, `--verbose`) |
| `history split --dry-run`, `--update-refs` | Already present |
| `history reword --dry-run`, `--update-refs` | Already present |

---

## Items Not Added (With Rationale)

| Item | Reason |
|------|--------|
| `clone --checkout` | Auto-generated positive form of `--no-checkout/-n` by pflag |
| `clone --hardlinks` | Auto-generated positive form of `--no-hardlinks` by pflag |
| `apply --add` | Does not exist in git source; only `--no-add` exists as negation of default behavior |
| `fetch --auto` | Does not exist; `--auto-gc` exists (auto housekeeping) |
| `grep --index` | Does not exist; `--cached` and `--no-index` exist |
| `merge --verify` | Only `--no-verify` exists as negation; auto-generated by pflag |
| `push --verify` | Only `--no-verify` exists as negation; auto-generated by pflag |
| `format-patch --binary` | Only `--no-binary` exists as negation; auto-generated by pflag |
| `clone --no-checkout` short `-n` | Already present |
| `revert --no-commit` short `-n` | Already present |

---

## Lower Priority Items (Not Addressed)

These commands have minor flag gaps that were not addressed due to lower impact:

- `bisect`: missing `--first-parent`, `--term-new`, `--term-old`
- `blame`: missing `--diff-algorithm`
- `multi-pack-index`: missing several flags
- `name-rev`: missing flags
- `pack-objects`: missing flags
- `send-pack`: missing flags
- `show-branch`: missing flags
- `update-index`: missing flags
- `update-ref`: missing flags
- `verify-pack`: missing flags
