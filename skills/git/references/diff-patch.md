# Git Diffs, Patches, and Conflict Resolution

How Git computes differences, generates patches, applies them, and resolves merge conflicts — including diff algorithms, patch formats, custom merge drivers, and rerere.

> **Source of truth**: <https://git-scm.com/docs/git-diff>, <https://git-scm.com/docs/git-apply>, <https://git-scm.com/docs/git-merge-file>. For **merging**, see [branching-merging.md](branching-merging.md). For **rebase conflicts**, see [rebase.md](rebase.md).

## Diff Algorithms

Git supports multiple diff algorithms:

|| Algorithm | Flag | Description |
||-----------|------|-------------|
|| **Myers** | Default | The original diff algorithm — minimal edit script |
|| **Patience** | `--patience` | Matches unique lines first, then diffs between them |
|| **Histogram** | `--histogram` | Extension of patience — handles repeated lines better |

### Algorithm Comparison

- **Myers**: Fast, produces minimal diffs, but can produce unintuitive results when lines repeat
- **Patience**: Better for code — aligns unique function signatures, class declarations
- **Histogram**: Default for `--histogram` — best for code with repeated patterns

```bash
git diff --patience
git diff --histogram
git config diff.algorithm histogram   # Set default
```

## Diff Commands

|| Command | Compares |
||---------|---------|
|| `git diff` | Index vs working tree |
|| `git diff --cached` / `--staged` | HEAD vs index |
|| `git diff HEAD` | HEAD vs working tree |
|| `git diff <a> <b>` | Two arbitrary tree-ishes |
|| `git diff <a>..<b>` | Same as `git diff <a> <b>` |
|| `git diff <a>...<b>` | Changes on b since merge base of a and b |
|| `git diff-index <tree>` | Tree vs index or working tree |
|| `git diff-tree <a> <b>` | Two tree objects |
|| `git diff-files` | Index vs working tree (plumbing) |

### Diff Output Formats

|| Format | Flag | Description |
||--------|------|-------------|
|| Default | (none) | Unified diff with 3 lines of context |
|| Stat | `--stat` | File-level summary with insertions/deletions |
|| Numstat | `--numstat` | Machine-readable: added deleted filename |
|| Shortstat | `--shortstat` | Only the summary line |
|| Name-only | `--name-only` | Only changed filenames |
|| Name-status | `--name-status` | Filenames with change type (A/M/D/R/C) |
|| Word diff | `--word-diff` | Inline word-level changes |
|| Word diff regex | `--word-diff-regex=<re>` | Custom word boundary |
|| Raw | `--raw` | Machine-readable raw format |

### Change Type Indicators

|| Letter | Meaning |
||--------|---------|
|| A | Added |
|| C | Copied |
|| D | Deleted |
|| M | Modified |
|| R | Renamed |
|| T | Type change |
|| U | Unmerged |
|| X | Unknown |

### Rename Detection

```bash
git diff -M              # Detect renames (default threshold 50%)
git diff -M80%           # Higher threshold (stricter)
git diff -C              # Detect copies too
git diff -C -C           # Detect copies from unmodified files
git diff --no-renames    # Disable rename detection
```

### Diff Context

```bash
git diff -U5             # 5 lines of context (default 3)
git diff --inter-hunk-context=5  # Context between hunks
git diff --function-context      # Show whole function as context
```

## Patch Format

Git uses the unified diff format for patches:

```diff
diff --git a/file.txt b/file.txt
index abc1234..def5678 100644
--- a/file.txt
+++ b/file.txt
@@ -1,3 +1,4 @@
 line 1
 line 2
+new line
 line 3
```

### Hunk Header

```
@@ -<old-start>,<old-count> +<new-start>,<new-count> @@ <function-name>
```

- `-` prefix: lines from the old file
- `+` prefix: lines from the new file
- Space prefix: context lines (in both)
- `\ No newline at end of file`: file doesn't end with newline

### Binary Patches

Binary diffs use a special format:

```diff
diff --git a/image.png b/image.png
Binary files a/image.png and b/image.png differ
```

Or with `--binary`:

```diff
GIT binary patch
delta 12345
base64-encoded binary data...
```

## Applying Patches

### git apply

Applies a patch to the working tree and/or index without creating commits:

```bash
git apply <patch>              # Apply to working tree
git apply --index <patch>      # Apply to working tree + index
git apply --cached <patch>     # Apply to index only
git apply --check <patch>      # Dry run (check if patch applies cleanly)
git apply --3way <patch>       # Fall back to 3-way merge on conflict
git apply --reject <patch>     # Apply what can, leave .rej files for rejects
git apply -R <patch>           # Reverse (un-apply)
git apply --stat <patch>       # Show summary without applying
```

### git am

Applies patches from a mailbox-format file (e.g., from `git format-patch`), creating commits:

```bash
git am <patch>                 # Apply patch as commit
git am --continue              # Continue after conflict resolution
git am --abort                 # Abort
git am --skip                  # Skip current patch
git am --3way                  # Fall back to 3-way merge
git am --signoff               # Add Signed-off-by
git am -s                      # Same
```

### git format-patch

Generates patch files in mailbox format:

```bash
git format-patch -1             # Patch for last commit
git format-patch -3             # Patches for last 3 commits
git format-patch main..feature  # Patches for commits in feature not in main
git format-patch -o /tmp/       # Output directory
git format-patch --cover-letter  # Add cover letter (thread summary)
```

### Patch-id

`git patch-id` computes a hash of a diff's content (ignoring whitespace and line numbers), used to identify equivalent changes:

```bash
git patch-id < diff-output
```

Used by `git cherry` and `git rebase` to detect cherry-picked commits.

## Conflict Resolution

### Conflict Markers

Standard markers (2-way):

```
<<<<<<< HEAD
our version
=======
their version
>>>>>>> feature
```

Diff3 markers (with `merge.conflictstyle = diff3`):

```
<<<<<<< HEAD
our version
||||||| base
original version
=======
their version
>>>>>>> feature
```

The `|||||||` section shows the common ancestor, making it easier to understand both changes.

### Resolution Workflow

1. Identify conflicted files: `git status`
2. Open each file and resolve markers
3. Stage resolved files: `git add <path>`
4. Continue: `git merge --continue` / `git rebase --continue` / `git cherry-pick --continue`

### Resolution Tools

```bash
git mergetool                    # Launch configured merge tool
git mergetool --tool=vimdiff     # Use specific tool
git config merge.tool meld       # Set default tool
```

### Common Merge Tools

|| Tool | Platform |
||------|----------|
|| `meld` | Linux |
|| `kdiff3` | Cross-platform |
|| `vimdiff` | Terminal |
|| `opendiff` | macOS (FileMerge) |
|| `tortoisemerge` | Windows |
|| `vscode` | Cross-platform |

### Custom Merge Drivers

Define in `.gitattributes`:

```
*.json merge=json-merge
```

Configure in `.git/config`:

```ini
[merge "json-merge"]
    name = JSON merge driver
    driver = json-merge %O %A %B %L
```

Parameters:
- `%O` — base (common ancestor) file path
- `%A` — ours file path (result is written here)
- `%B` — theirs file path
- `%L` — conflict marker size

Exit codes:
- 0 = merge succeeded
- non-zero = merge failed (fall back to conflict markers)

### Built-in Merge Drivers

|| Driver | Behavior |
||--------|----------|
|| `text` | Default three-way text merge |
|| `binary` | Keep ours, no merge attempt |
|| `union` | Concatenate both sides (no conflict markers) |

## Rerere (Reuse Recorded Resolution)

`git rerere` records how you resolved conflicts and automatically applies the same resolution next time:

```bash
git config rerere.enabled true   # Enable
git rerere status                 # Show files with recorded resolutions
git rerere diff                   # Show what rerere would apply
git rerere forget <path>          # Delete recorded resolution for path
```

### How Rerere Works

1. During a conflict, rerere records the conflict markers
2. When you resolve and commit, rerere records the resolution
3. Next time the same conflict appears, rerere auto-applies the resolution
4. You can review with `git rerere diff` before committing

Use case: long-lived feature branches that are frequently rebased — the same conflicts recur.

## Merge Conf

```
.git/MERGE_MSG          # Merge commit message (during merge)
.git/MERGE_HEAD         # The commit being merged
.git/MERGE_MODE         # Merge options
```

## Edge Cases and Known Issues

- **Whitespace conflicts**: Changes that only differ in whitespace can cause unexpected conflicts. Use `--ignore-space-change` or `--ignore-all-space` with diff/merge.
- **Encoding conflicts**: Files with different encodings may produce garbled conflict markers. Ensure `core.precomposeunicode` and `i18n.commitEncoding` are set correctly.
- **Large file conflicts**: Binary files always conflict (no three-way merge). Use `.gitattributes` merge drivers or LFS.
- **Submodule conflicts**: Submodule conflicts show two commit OIDs. You must choose one and check it out.
- **Rename + modify conflicts**: If one side renames a file and the other modifies it, the ort strategy handles this, but older strategies may produce delete/modify conflicts.
- **Criss-cross merge conflicts**: Multiple merge bases can cause unexpected conflict regions. The ort strategy handles this better than recursive.

## References

- <https://git-scm.com/docs/git-diff>
- <https://git-scm.com/docs/git-apply>
- <https://git-scm.com/docs/git-rerere>
- <https://git-scm.com/docs/git-merge-file>

## Related Skills

- For merging strategies, see [branching-merging.md](branching-merging.md)
- For rebase conflicts, see [rebase.md](rebase.md)
- For index internals (stage numbers during conflicts), see [index.md](index.md)
