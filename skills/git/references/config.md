# Git Configuration System

How Git's configuration works — config levels, precedence, includes, conditional includes, key format, and environment variable overrides.

> **Source of truth**: <https://git-scm.com/docs/git-config>, <https://git-scm.com/book/en/v2/Git-Customizing-Git-Git-Configuration>. For **hooks configuration**, see [hooks.md](hooks.md). For **worktree-specific config**, see [worktree.md](worktree.md).

## Config Levels and Precedence

Git reads configuration from multiple files, with later entries **overriding** earlier ones:

|| Level | File | Scope | Priority |
||-------|------|-------|----------|
|| **system** | `/etc/gitconfig` | All users on the system | 1 (lowest) |
|| **global** | `~/.gitconfig` or `$XDG_CONFIG_HOME/git/config` | Current user | 2 |
|| **local** | `.git/config` | Current repository | 3 |
|| **worktree** | `.git/config.worktree` | Current worktree (with `extensions.worktreeConfig`) | 4 |
|| **command** | `-c key=value` | Single command | 5 (highest) |
|| **env** | `GIT_CONFIG_*` | Process lifetime | 6 (highest) |

### Reading Order

```
system → global → local → worktree → -c flag → env var
```

Each level can add new keys or override values from lower levels.

## Config Commands

```bash
git config --list                    # List all config (all levels)
git config --list --show-origin      # Show which file each entry comes from
git config --list --show-scope       # Show scope (system/global/local)
git config <key>                     # Get value
git config --get <key>               # Get value (same)
git config --get-all <key>           # Get all values (for multi-valued keys)
git config --get-regexp <pattern>    # Get all matching keys
git config <key> <value>             # Set value (default: local)
git config --global <key> <value>    # Set in global config
git config --system <key> <value>    # Set in system config
git config --local <key> <value>     # Set in local config
git config --unset <key>             # Remove key
git config --unset-all <key>         # Remove all values for key
git config --rename-section <old> <new>  # Rename section
git config --remove-section <name>   # Remove section
git config --edit                    # Open config in editor
git config --global --edit           # Edit global config
```

## Config File Format

Git config uses INI-like format:

```ini
[user]
    name = Alice
    email = alice@example.com

[core]
    editor = vim
    autocrlf = input

[remote "origin"]
    url = https://github.com/user/repo.git
    fetch = +refs/heads/*:refs/remotes/origin/*

[branch "main"]
    remote = origin
    merge = refs/heads/main

[alias]
    co = checkout
    st = status --short
    lg = log --oneline --graph --all
```

### Key Format

Config keys are `section.subsection.name`:

- `section` = top-level group (e.g., `user`, `core`, `remote`)
- `subsection` = optional quoted group (e.g., `"origin"` in `remote "origin"`)
- `name` = the specific key

In commands, use dot notation:

```bash
git config remote.origin.url
git config branch.main.remote
```

### Value Types

|| Type | Example | Parsing |
||------|---------|---------|
|| Boolean true | `true`, `yes`, `on`, `1` | Case-insensitive |
|| Boolean false | `false`, `no`, `off`, `0`, empty string | Case-insensitive |
|| Integer | `42` | Parsed as number |
|| Path | `~/bin/git-diff-wrapper` | `~` expanded for some keys |
|| Color | `red bold`, `green underline` | See color format |
|| String | Any value | Taken as-is |

### Multi-Valued Keys

Some keys can have multiple values:

```ini
[remote "origin"]
    fetch = +refs/heads/*:refs/remotes/origin/*
    fetch = +refs/tags/*:refs/tags/*
```

```bash
git config --get-all remote.origin.fetch
```

## Includes

### Unconditional Include

```ini
[include]
    path = ~/.gitconfig-work
    path = /etc/gitconfig-local
```

### Conditional Include

Include a config file only when certain conditions are met:

```ini
[includeIf "gitdir:~/projects/work/"]
    path = ~/.gitconfig-work

[includeIf "gitdir/i:~/Projects/Work/"]
    path = ~/.gitconfig-work

[includeIf "hasconfig:remote.*.url:github.com/work/**"]
    path = ~/.gitconfig-work

[includeIf "onbranch:main"]
    path = ~/.gitconfig-main
```

### Condition Types

|| Condition | Syntax | Matches |
||-----------|--------|---------|
|| `gitdir` | `gitdir:<pattern>` | Working directory path (glob) |
|| `gitdir/i` | `gitdir/i:<pattern>` | Case-insensitive gitdir |
|| `hasconfig` | `hasconfig:remote.*.url:<pattern>` | Config value matches pattern |
|| `onbranch` | `onbranch:<pattern>` | Current branch name matches |

## Key Sections Reference

### user

```ini
[user]
    name = Alice                    # Commit author name
    email = alice@example.com        # Commit author email
    signingKey = ABC12345            # GPG/SSH key for signing
    useConfigOnly = true            # Don't guess name/email from system
```

### core

```ini
[core]
    editor = vim                     # Default editor (VISUAL, EDITOR, then vi)
    autocrlf = input                 # Line ending conversion (input, true, false)
    safecrlf = true                  # Warn on irreversible CRLF conversion
    fileMode = false                 # Ignore file mode changes
    symlinks = true                  # Enable symlink support
    ignoreCase = false               # Case-insensitive path matching
    precomposeunicode = true         # Normalize Unicode (macOS)
    quotePath = true                 # Quote unusual chars in path output
    bigFileThreshold = 512m          # Threshold for "big file" handling
    compression = 6                  # Zlib compression level (0-9)
    packedGitLimit = 256m            # Max memory for packfile mmap
    packedGitWindowSize = 1m         # Packfile window size
    deltaCacheSize = 256m            # Delta cache size
    packSizeLimit = 2g               # Max pack size
    whitespace = fix                 # Default whitespace handling
    bare = false                     # Repository is bare
    worktree = /path/to/worktree     # Working tree path
    logAllRefUpdates = true          # Enable reflogs
    repositoryFormatVersion = 0     # Repo format version
    sharedRepository = group         # Shared repo permissions
    warnAmbiguousRefs = true         # Warn on ambiguous ref names
```

### init

```ini
[init]
    defaultBranch = main             # Default branch name
    templateDir = /path/to/template  # Template directory
    defaultHash = sha256              # Hash algorithm (experimental)
```

### alias

```ini
[alias]
    co = checkout
    br = branch
    st = status --short
    lg = log --oneline --graph --all --decorate
    unstage = reset HEAD --
    last = log -1 HEAD
    visual = log --graph --oneline --all
    amend = commit --amend --no-edit
    pushf = push --force-with-lease
```

Shell command aliases (prefixed with `!`):

```ini
[alias]
    incoming = !git fetch && git log ..origin/$(git branch --show-current)
    outgoing = !git log origin/$(git branch --show-current)..
    find-file = "!f() { git ls-files | grep -i \"$1\"; }; f"
```

### diff

```ini
[diff]
    algorithm = histogram            # Default diff algorithm
    renames = true                   # Enable rename detection
    indentHeuristic = true           # Better diff hunk splitting
    context = 5                      # Default context lines
    statMaxWidth = 120               # Max width for --stat
    colorMoved = default             # Color moved lines
    wsErrorHighlight = all           # Highlight whitespace errors
```

### merge

```ini
[merge]
    tool = meld                      # Default merge tool
    conflictstyle = diff3           # Conflict marker style (merge, diff3)
    ff = true                        # Default fast-forward behavior
    verifySignatures = false         # Verify GPG signatures on merge
    log = true                       # Include merge summary in commit
    renameLimit = 1000               # Rename detection limit
    renames = true                   # Enable rename detection
    autoSquash = true                # Auto-squash in rebase
```

### push / pull / fetch

```ini
[push]
    default = simple                 # Push default mode
    autoSetupRemote = true           # Auto set upstream on first push
    followTags = true                # Push annotated tags
    recurseSubmodules = check        # Submodule handling

[pull]
    rebase = true                    # Rebase instead of merge
    ff = only                        # Only fast-forward

[fetch]
    prune = true                     # Prune stale remote-tracking refs
    recurseSubmodules = on-demand    # Submodule fetching
```

### rerere

```ini
[rerere]
    enabled = true                   # Enable rerere
    autoupdate = true                # Auto-stage resolved files
```

### credential

```ini
[credential]
    helper = cache --timeout=3600    # Credential helper
    useHttpPath = true               # Include path in credential key
```

### signing

```ini
[gpg]
    format = ssh                     # Signing format (openpgp, ssh, x509)
    program = gpg                    # GPG program

[commit]
    gpgSign = true                   # Sign all commits

[tag]
    gpgSign = true                   # Sign all tags
```

## Environment Variables

Environment variables override config values:

|| Variable | Overrides | Description |
||----------|----------|-------------|
|| `GIT_AUTHOR_NAME` | `user.name` | Author name for commits |
|| `GIT_AUTHOR_EMAIL` | `user.email` | Author email |
|| `GIT_AUTHOR_DATE` | Author date | Format: `@<timestamp> <tz>` |
|| `GIT_COMMITTER_NAME` | `user.name` | Committer name |
|| `GIT_COMMITTER_EMAIL` | `user.email` | Committer email |
|| `GIT_COMMITTER_DATE` | Committer date | Same format |
|| `GIT_CONFIG` | All config files | Single config file (disables system/global) |
|| `GIT_CONFIG_GLOBAL` | Global config | Override global config file |
|| `GIT_CONFIG_SYSTEM` | System config | Override system config file |
|| `GIT_EDITOR` | `core.editor` | Editor for commits |
|| `GIT_DIFF_PATH_COUNTER` | — | Diff counter (in external diff) |
|| `GIT_DIFF_PATH_TOTAL` | — | Total diff count |
|| `GIT_EXTERNAL_DIFF` | `diff.external` | External diff command |
|| `GIT_MERGE_VERBOSITY` | `merge.verbosity` | Merge verbosity |
|| `GIT_TRACE` | — | Trace Git commands (1, 2, or path) |
|| `GIT_TRACE_PACKET` | — | Trace pack protocol |
|| `GIT_LITERAL_PATHSPECS` | — | Disable glob pathspecs |

## Edge Cases and Known Issues

- **Config precedence is last-wins**: Within a single file, later entries override earlier ones. Across files, higher-priority files override lower.
- **Subsection quoting**: Subsections with special characters must be quoted: `[remote "my-repo"]` not `[remote my-repo]`.
- **Boolean empty string**: An empty value for a boolean key is `false`. `git config core.bare ""` sets it to false.
- **Include cycles**: Git detects include cycles and stops after a depth limit.
- **Conditional include path**: `gitdir` patterns use glob matching. `~` is expanded. Trailing `/` means "directory and its children".
- **Worktree config**: Requires `extensions.worktreeConfig = true` in the local config. Without it, `.git/config.worktree` is ignored.
- **`-c` flag quoting**: `git -c key="value with spaces"` — the shell handles quoting, not Git.

## References

- <https://git-scm.com/docs/git-config>
- <https://git-scm.com/book/en/v2/Git-Customizing-Git-Git-Configuration>

## Related Skills

- For hooks configuration, see [hooks.md](hooks.md)
- For worktree-specific config, see [worktree.md](worktree.md)
- For credential helpers, see [remotes.md](remotes.md)
