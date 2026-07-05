# Taps

Taps are Git repositories containing formulae, casks, and/or custom brew commands. This document covers tap structure, the `brew tap` command, API mode vs git clone mode, and formula resolution.

## What Is a Tap?

A tap is a Git repository that Homebrew clones into `$(brew --repository)/Library/Taps/`. Each tap provides additional formulae and/or casks that become available to `brew install`, `brew search`, etc.

The naming convention is `user/repository`, where the actual GitHub repository must be named `homebrew-<repository>`. On the command line, the `homebrew-` prefix is omitted.

| Command | Behavior |
|---------|---------|
| `brew tap` | List all currently tapped repositories |
| `brew tap <user>/<repo>` | Clones `https://github.com/<user>/homebrew-<repo>` into `Library/Taps` |
| `brew tap <user>/<repo> <URL>` | Clones from an arbitrary URL (any Git protocol) |
| `brew untap <tap>` [...] | Removes given taps |
| `brew tap-info` [--installed] [--json] [`<tap>` ...] | Show detailed information about taps |
| `brew tap-new` <user>/<repo> | Generate template files for a new tap |

## Core Taps

| Tap | Class | Contents |
|-----|-------|----------|
| `homebrew/core` | `CoreTap` | The primary formula repository |
| `homebrew/cask` | `CoreCaskTap` | The primary cask repository |

Both are automatically available and updated as part of a standard Homebrew installation. In API mode (default since 4.0), their formula/cask data is fetched as JSON from `https://formulae.brew.sh/api` rather than requiring a full local Git clone.

## Tap Directory Structure

```
$(brew --repository)/Library/Taps/
  <user>/
    homebrew-<repo>/
      Formula/
        f/
          foo.rb
        b/
          bar.rb
      Casks/
        b/
          baz.rb
      cmd/
        brew-custom-command    # external commands
      Aliases/
        foo-alias -> ../Formula/f/foo.rb
      README.md
```

- **`Formula/`** — formula `.rb` files (often organized in subdirectories by first letter: `Formula/f/foo.rb`)
- **`Casks/`** — cask `.rb` files
- **`cmd/`** — custom brew commands (become `brew <command>` when tapped)
- **`Aliases/`** — symlinks providing formula aliases

## `brew tap` Flags

| Flag | Purpose |
|------|---------|
| `--custom-remote` | Install with a custom remote (useful for mirrors) |
| `--repair` | Fix missing symlinks and git remote refs |
| `-f, --force` | Force install core taps even in API mode |

## `brew untap` Flags

| Flag | Purpose |
|------|---------|
| `-f, --force` | Untap even if formulae from this tap are installed |

## `brew tap-new` Flags

| Flag | Purpose |
|------|---------|
| `--no-git` | Don't initialize a Git repository |
| `--branch <name>` | Set the default branch name |
| `--github-packages` | Set up GitHub Packages for bottle hosting |

## API Mode vs Git Clone Mode

| Aspect | API Mode (default, 4.0+) | Git Clone Mode |
|--------|--------------------------|----------------|
| **Formula/cask data** | Fetched as JSON from `https://formulae.brew.sh/api` | Full Git clone of `homebrew/core` and `homebrew/cask` |
| **Speed** | Fast `brew update`, low disk usage | Slow `brew update`, large disk usage |
| **Build from source** | Still fetches the Ruby formula file on demand | Already has all formula files locally |
| **Enable** | Default | Set `HOMEBREW_NO_INSTALL_FROM_API=1` |
| **API auto-update** | Every `HOMEBREW_API_AUTO_UPDATE_SECS` (default 450s) | N/A |

In API mode, taps may still be needed for third-party repositories not served by the Homebrew API.

## Formula Resolution Order

When `brew install foo` is run, Homebrew searches in this order:

1. **Core formulae** (`homebrew/core`) — first priority
2. **Other taps** — secondary priority, in tap order

To install from a specific tap, use fully qualified names:

```bash
brew install username/repo/vim    # installs from custom tap
brew install vim                  # installs from homebrew/core
```

There is **intentionally no way** to replace dependencies of core formulae with those from other taps.

## Tap Trust

Modern Homebrew requires non-official tap formulae, casks, and commands to be trusted before they are loaded. This is a security measure to prevent arbitrary code execution from untrusted taps.

| Command | Purpose |
|---------|---------|
| `brew trust <target>` | Trust a tap, formula, cask, or command |
| `brew untrust <target>` | Revoke trust |
| `brew trust --tap <user/repo>` | Trust an entire tap |
| `brew trust --json=v1` | Print trusted entries as JSON (requires version argument) |

Related environment variables:

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_REQUIRE_TAP_TRUST` | `true` | Require trust before loading non-official tap content |
| `HOMEBREW_NO_REQUIRE_TAP_TRUST` | unset | Do not require trust (not recommended) |

In `brew bundle` (Brewfiles), taps can be marked trusted:

```ruby
tap "user/repo", trusted: true
tap "user/repo", trusted: { formula: "foo", cask: "bar" }
```

## References

- [Homebrew Taps documentation](https://docs.brew.sh/Taps)
- For formulae: [formula.md](formula.md)
- For casks: [cask.md](cask.md)
- For environment variables: [environment.md](environment.md)
- For bundle: [bundle.md](bundle.md)
