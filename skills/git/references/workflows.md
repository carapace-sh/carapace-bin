# Git Workflows and Patterns

Common Git workflows, branching strategies, commit conventions, .gitignore patterns, Git LFS, and .gitattributes.

> **Source of truth**: <https://git-scm.com/book/en/v2/Distributed-Git-Distributed-Workflows>, <https://www.atlassian.com/git/tutorials/comparing-workflows>. For **branching and merging**, see [branching-merging.md](branching-merging.md). For **rebase**, see [rebase.md](rebase.md). For **hooks**, see [hooks.md](hooks.md).

## Branching Strategies

### Git Flow

A structured branching model with long-lived branches:

```
main ──────────────────────────────────────
  ↘  ↗  release/1.0 ──────────
   ↘ ↗  develop ────────────────────────
     ↘  ↗  feature/login ────
```

|| Branch | Purpose | Merges into |
||--------|---------|-------------|
|| `main` | Production releases | — |
|| `develop` | Integration branch | `main` (via release) |
|| `feature/*` | New features | `develop` |
|| `release/*` | Release preparation | `main` + `develop` |
|| `hotfix/*` | Emergency fixes | `main` + `develop` |

```bash
# Start a feature
git checkout -b feature/login develop
# Finish a feature
git checkout develop && git merge --no-ff feature/login
git branch -d feature/login
```

### GitHub Flow

A simpler model optimized for continuous deployment:

```
main ────────────────────────────────
  ↘  ↗  feature/login ────
```

1. Create a branch from `main`
2. Make commits and push
3. Open a pull request
4. Review and discuss
5. Merge into `main`
6. Delete the branch
7. Deploy from `main`

### Trunk-Based Development

All developers commit to a single branch (`main`/`trunk`):

```
main ────────────────────────────────────
     ↗↘  short-lived feature branches (max 1-2 days)
```

- Feature flags control incomplete features
- Branches are very short-lived (< 1 day)
- Continuous integration is mandatory
- Pre-commit hooks enforce quality

### Comparison

|| Aspect | Git Flow | GitHub Flow | Trunk-Based |
||--------|----------|-------------|-------------|
|| Complexity | High | Low | Low |
|| Release cadence | Scheduled | Continuous | Continuous |
|| Long-lived branches | 2+ | 1 | 0 |
|| Merge strategy | `--no-ff` | Squash or merge | Fast-forward |
|| Release branches | Yes | No | No (feature flags) |
|| Hotfix branches | Yes | Fix on main | Fix on main |

## Commit Conventions

### Conventional Commits

A structured commit message format:

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

Types:

|| Type | Description |
||------|-------------|
|| `feat` | New feature |
|| `fix` | Bug fix |
|| `docs` | Documentation |
|| `style` | Formatting (no code change) |
|| `refactor` | Code restructuring |
|| `perf` | Performance improvement |
|| `test` | Adding tests |
|| `build` | Build system |
|| `ci` | CI configuration |
|| `chore` | Maintenance |
|| `revert` | Revert a commit |

Breaking changes: `feat!:` or `BREAKING CHANGE:` in footer.

### Semantic Versioning Integration

Conventional commits enable automatic semantic versioning:
- `feat:` → minor version bump
- `fix:` → patch version bump
- `feat!:` or `BREAKING CHANGE:` → major version bump

### 7 Rules of Great Commit Messages

1. Separate subject from body with a blank line
2. Limit the subject line to 50 characters
3. Capitalize the subject line
4. Do not end the subject line with a period
5. Use the imperative mood ("Add feature" not "Added feature")
6. Wrap the body at 72 characters
7. Use the body to explain *what* and *why*, not *how*

## .gitignore Patterns

### Pattern Syntax

|| Pattern | Matches |
||---------|---------|
|| `*.log` | All `.log` files in any directory |
|| `build/` | `build` directory and its contents |
|| `!important.log` | Negation — don't ignore this file |
|| `doc/*.txt` | `.txt` files directly in `doc/` |
|| `doc/**/*.pdf` | `.pdf` files anywhere under `doc/` |
|| `/TODO` | Only `TODO` at the repository root |
|| `foo/` | Directory named `foo` |
|| `foo` | File or directory named `foo` |

### Pattern Rules

- A blank line matches nothing (separator)
- A line starting with `#` is a comment
- Trailing spaces are ignored (unless quoted)
- A `!` prefix negates the pattern
- A `/` prefix anchors to the repository root
- A trailing `/` matches directories only
- A `**/` prefix matches in all directories
- A `/**` suffix matches everything inside
- A `/**/` matches zero or more directories

### .gitignore Locations

|| File | Scope | Tracked |
||------|-------|---------|
|| `.gitignore` (repo root) | Entire repository | Yes |
|| `src/.gitignore` | `src/` and below | Yes |
|| `.git/info/exclude` | Entire repository | No (local only) |
|| `core.excludesFile` | All repos (user-level) | No |

### Precedence

Patterns are checked in order; later patterns override earlier ones. Within a single file, a negation pattern (`!`) after an ignore pattern works. Across files, deeper `.gitignore` files take precedence.

### Common .gitignore

```
# Build output
/build/
/dist/
*.o
*.so
*.exe

# Dependencies
/node_modules/
/vendor/

# IDE
.idea/
.vscode/
*.swp

# OS
.DS_Store
Thumbs.db

# Environment
.env
.env.local
```

## .gitattributes

Assign attributes to paths for diff, merge, and other operations:

```
# Text files — normalize line endings
*.txt text
*.md text

# Auto-detect text and normalize
* text=auto

# Binary files — no diff, no merge
*.png binary
*.jpg binary
*.zip binary

# Custom diff driver
*.py diff=python

# Custom merge driver
*.json merge=json-merge

# LFS tracking
*.psd filter=lfs diff=lfs merge=lfs -text
*.mp4 filter=lfs diff=lfs merge=lfs -text

# Export-ignore (not in archive)
.gitattributes export-ignore
.gitignore export-ignore

# Linguist overrides (GitHub)
vendor/*.rb linguist-vendored
*.js linguist-language=JavaScript
```

### Attribute Assignment

|| Attribute | Values | Description |
||-----------|--------|-------------|
|| `text` | `auto`, `true`, `false` | Line ending normalization |
|| `eol` | `lf`, `crlf`, `native` | Force line ending |
|| `binary` | `true`, `false` | Treat as binary |
|| `diff` | driver name, `false` | Diff driver |
|| `merge` | driver name, `false`, `union` | Merge driver |
|| `filter` | driver name | Smudge/clean filter |
|| `export-subst` | `true` | Expand gitattributes in archive |
|| `export-ignore` | `true` | Exclude from archive |
|| `linguist-*` | Various | GitHub language detection |

## Git LFS (Large File Storage)

Git LFS replaces large files with pointer files, storing the actual content on a separate server.

### Setup

```bash
git lfs install                        # Install LFS hooks
git lfs track "*.psd"                  # Track pattern
git lfs track "assets/**"             # Track directory
git lfs track --filename "*.bin"      # Track by filename
```

This updates `.gitattributes`:

```
*.psd filter=lfs diff=lfs merge=lfs -text
```

### Pointer Files

LFS replaces tracked files with pointer files:

```
version https://git-lfs.github.com/spec/v1
oid sha256:4d7a214614ab2935c943f9e0ff69d22eadbb8f32b1258daa4e2e6a0e0e6d2b2c
size 12345
```

### LFS Commands

```bash
git lfs fetch                          # Download LFS objects
git lfs pull                           # Fetch + checkout
git lfs push --all origin              # Push all LFS objects
git lfs prune                          # Remove old LFS objects
git lfs ls-files                       # List LFS-tracked files
git lfs status                         # Show LFS status
git lfs migrate import --include="*.psd"  # Convert existing files to LFS
git lfs migrate export --include="*.psd"  # Convert LFS back to regular
```

### LFS Locking

Prevent concurrent editing of binary files:

```bash
git lfs lock images/logo.psd           # Lock file
git lfs unlock images/logo.psd         # Unlock
git lfs locks                          # List locks
```

Requires server support (GitHub, GitLab, Bitbucket).

### Smudge/Clean Filters

LFS uses Git's filter mechanism:

- **Smudge** (`git lfs smudge`): On checkout, replace pointer with actual file
- **Clean** (`git lfs clean`): On `git add`, replace file with pointer

```
working tree → clean filter → index (pointer)
index (pointer) → smudge filter → working tree (actual file)
```

## Git Bisect Workflow

Binary search to find the commit that introduced a bug:

```bash
git bisect start
git bisect bad                          # Current is bad
git bisect good v1.0                    # v1.0 was good
# Git checks out middle commit
# Test it
git bisect good  # or git bisect bad
# Repeat...
git bisect reset                        # Return to original branch
```

### Automated Bisect

```bash
git bisect run ./test-script.sh
```

The script should:
- Exit 0 if the commit is good
- Exit 1 if the commit is bad
- Exit 125 if the commit cannot be tested (skip)

### Bisect Log

```bash
git bisect log                          # Show bisect progress
git bisect replay bisect-log.txt        # Replay a bisect
```

## Git Stash Workflow

Temporarily save uncommitted changes:

```bash
git stash                               # Stash all changes
git stash -u                            # Include untracked files
git stash -m "work in progress"        # With message
git stash list                          # List stashes
git stash pop                           # Apply and drop latest
git stash apply stash@{2}              # Apply specific stash
git stash drop stash@{0}               # Drop stash
git stash branch feature stash@{0}     # Create branch from stash
```

### Partial Stash

```bash
git stash -p                            # Interactive stash
git stash push -m "msg" -- <path>      # Stash specific paths
```

## Code Review Workflows

### Pull Request (GitHub)

```bash
# Create feature branch
git checkout -b feature/login
# Make commits
git push -u origin feature/login
# Open PR on GitHub
# After review, merge via GitHub UI
# Clean up
git checkout main && git pull
git branch -d feature/login
git push origin --delete feature/login
```

### Merge Request (GitLab)

Same flow as PR, but using GitLab's merge request interface.

### Gerrit

```bash
# Push for review (not to a branch, but to refs/for/)
git push origin HEAD:refs/for/main
# Review happens on Gerrit
# After approval, Gerrit merges automatically
```

## Signing Commits and Tags

### GPG Signing

```bash
git config commit.gpgSign true          # Sign all commits
git config tag.gpgSign true             # Sign all tags
git commit -S -m "Signed commit"        # Sign this commit
git tag -s v1.0 -m "Signed tag"         # Sign this tag
git verify-commit <oid>                 # Verify commit signature
git tag -v v1.0                         # Verify tag signature
```

### SSH Signing

```bash
git config gpg.format ssh
git config user.signingKey ~/.ssh/id_ed25519.pub
git commit -S -m "SSH-signed commit"
```

### Allowed Signers

For SSH verification, configure allowed signers:

```bash
git config gpg.ssh.allowedSignersFile ~/.ssh/allowed_signers
```

Format of `allowed_signers`:

```
alice@example.com ssh-ed25519 AAAAC3NzaC1lZDI1NTE5...
```

## Edge Cases and Known Issues

- **LFS and partial clone**: LFS objects are not fetched by `--filter=blob:none`. Use `git lfs fetch` separately.
- **.gitignore negation order**: A negation pattern only works if the file is ignored by an earlier pattern in the same file or a parent `.gitignore`.
- **.gitattributes and merge conflicts**: Custom merge drivers must be defined in config. If the driver is missing, Git falls back to the text driver.
- **Conventional commits and rebase**: Interactive rebase with `reword` preserves the conventional commit format if the editor doesn't strip the type prefix.
- **LFS migration**: `git lfs migrate import` rewrites history. All collaborators must re-clone.
- **Signing and CI**: CI systems need access to the signing key. Use environment variables or secret management.

## References

- <https://git-scm.com/book/en/v2/Distributed-Git-Distributed-Workflows>
- <https://www.conventionalcommits.org/>
- <https://git-lfs.github.com/>
- <https://git-scm.com/docs/gitattributes>

## Related Skills

- For branching and merging, see [branching-merging.md](branching-merging.md)
- For rebase, see [rebase.md](rebase.md)
- For hooks (pre-commit, commit-msg), see [hooks.md](hooks.md)
- For configuration, see [config.md](config.md)
- For diff and merge drivers, see [diff-patch.md](diff-patch.md)
