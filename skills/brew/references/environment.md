# Environment Variables and Configuration

Homebrew configures itself primarily through `HOMEBREW_*` environment variables and `brew.env` configuration files. This document covers all environment variables, configuration files, and diagnostic commands.

## Configuration Files

`HOMEBREW_*` environment variables can be set in environment files, loaded in this precedence order (lowest to highest):

| File | Scope |
|------|-------|
| `/etc/homebrew/brew.env` | System-wide |
| `${HOMEBREW_PREFIX}/etc/homebrew/brew.env` | Prefix-specific |
| `$XDG_CONFIG_HOME/homebrew/brew.env` or `~/.homebrew/brew.env` | User-specific |

Unless `HOMEBREW_SYSTEM_ENV_TAKES_PRIORITY` is set, user-specific files override prefix-specific files, which override system-wide files. These files do **not** support shell variable expansion (`$HOME`) or command execution (`$(cat file)`).

## Core Path Variables

Set by `brew shellenv` and exported in the shell profile:

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_PREFIX` | `/opt/homebrew` (ARM), `/usr/local` (Intel), `/home/linuxbrew/.linuxbrew` (Linux) | Homebrew installation path |
| `HOMEBREW_CELLAR` | `$(brew --prefix)/Cellar` | Cellar directory |
| `HOMEBREW_REPOSITORY` | Typically same as `HOMEBREW_PREFIX` | brew Git repository path |

## API and Update

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_NO_INSTALL_FROM_API` | — | Use local Git checkouts instead of JSON API |
| `HOMEBREW_NO_AUTO_UPDATE` | — | Don't auto-update before `install`, `upgrade`, `tap` |
| `HOMEBREW_AUTO_UPDATE_SECS` | `86400` (24h) | Auto-update interval; `3600` for dev commands, `300` if no API |
| `HOMEBREW_API_AUTO_UPDATE_SECS` | `450` | API data refresh interval |
| `HOMEBREW_API_DOMAIN` | `https://formulae.brew.sh/api` | JSON API mirror URL |
| `HOMEBREW_FORCE_API_AUTO_UPDATE` | — | Update API even if `HOMEBREW_NO_AUTO_UPDATE` is set |
| `HOMEBREW_BREW_GIT_REMOTE` | `https://github.com/Homebrew/brew` | brew Git remote URL |
| `HOMEBREW_CORE_GIT_REMOTE` | `https://github.com/Homebrew/homebrew-core` | homebrew-core Git remote URL |
| `HOMEBREW_UPDATE_TO_TAG` | — | Always use latest stable tag |

## Bottle and Artifact Domains

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_BOTTLE_DOMAIN` | `https://ghcr.io/v2/homebrew/core` | Bottle download mirror |
| `HOMEBREW_ARTIFACT_DOMAIN` | — | Prefix all download URLs (mirror) |
| `HOMEBREW_ARTIFACT_DOMAIN_NO_FALLBACK` | — | Error instead of fallback when artifact domain fails |

## Download and Build

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_CACHE` | macOS: `~/Library/Caches/Homebrew`, Linux: `$XDG_CACHE_HOME/Homebrew` or `~/.cache/Homebrew` | Download cache directory |
| `HOMEBREW_TEMP` | macOS: `/private/tmp`, Linux: `/var/tmp` | Temporary build directory |
| `HOMEBREW_LOGS` | macOS: `~/Library/Logs/Homebrew`, Linux: `${XDG_CACHE_HOME}/Homebrew/Logs` or `~/.cache/Homebrew/Logs` | Log file directory |
| `HOMEBREW_DOWNLOAD_CONCURRENCY` | `auto` (2x CPU cores) | Concurrent download connections |
| `HOMEBREW_CURL_RETRIES` | `3` | curl retry count |
| `HOMEBREW_CURL_VERBOSE` | — | Pass `--verbose` to curl |
| `HOMEBREW_CURLRC` | — | If set to an absolute path, pass it with `--config` to curl; if set but not a valid path, don't pass `--disable` (enables `.curlrc` loading, which is disabled by default) |
| `HOMEBREW_CURL_PATH` | `curl` | (Linux) curl executable path |
| `HOMEBREW_MAKE_JOBS` | CPU core count | Parallel make jobs |
| `HOMEBREW_ARCH` | `native` | (Linux) compiler `-march` value |
| `HOMEBREW_PIP_INDEX_URL` | `https://pypi.org/simple` | PyPI URL for formula resources |

## Cleanup

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_CLEANUP_MAX_AGE_DAYS` | `120` | Max age for cached files |
| `HOMEBREW_CLEANUP_PERIODIC_FULL_DAYS` | `30` | Full cleanup interval |
| `HOMEBREW_NO_INSTALL_CLEANUP` | — | Never auto-cleanup on install/upgrade/reinstall |
| `HOMEBREW_NO_CLEANUP_FORMULAE` | — | Comma-separated formulae exempt from cleanup |
| `HOMEBREW_NO_AUTOREMOVE` | — | Don't auto-remove unused dependents |

## Cask Options

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_CASK_OPTS` | — | Default cask options (`--*dir`, `--language`, `--require-sha`, `--no-binaries`) |
| `HOMEBREW_UPGRADE_GREEDY` | — | Pass `--greedy` to all cask upgrade commands |
| `HOMEBREW_UPGRADE_GREEDY_CASKS` | — | Space-separated casks to always upgrade greedily |
| `HOMEBREW_NO_UPGRADE_AUTO_UPDATES_CASKS` | — | Don't auto-upgrade `auto_updates true` casks |
| `HOMEBREW_NO_UPGRADE_QUIT_CASKS` | — | Don't quit running apps during cask upgrades |
| `HOMEBREW_FORBID_CASKS` | — | Refuse to install any casks |
| `HOMEBREW_FORBIDDEN_CASKS` | — | Space-separated casks to refuse |
| `HOMEBREW_FORBIDDEN_CASK_ARTIFACTS` | — | Space-separated artifact types to forbid (e.g. `pkg installer`) |

## Security and Restrictions

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_ALLOWED_TAPS` | — | Space-separated taps allowed for installation |
| `HOMEBREW_FORBIDDEN_TAPS` | — | Space-separated taps to refuse |
| `HOMEBREW_FORBIDDEN_FORMULAE` | — | Space-separated formulae to refuse |
| `HOMEBREW_FORBIDDEN_LICENSES` | — | Space-separated SPDX licenses to refuse |
| `HOMEBREW_FORBIDDEN_OWNER` | `you` | Who set the forbidden variables |
| `HOMEBREW_FORBIDDEN_OWNER_CONTACT` | — | Contact info for forbidden owner |
| `HOMEBREW_FORBID_PACKAGES_FROM_PATHS` | `true` (unless developer) | Refuse formulae/casks from file paths |
| `HOMEBREW_NO_REQUIRE_TAP_TRUST` | — | Don't require tap trust (not recommended) |
| `HOMEBREW_REQUIRE_TAP_TRUST` | `true` | Require tap trust |
| `HOMEBREW_NO_INSECURE_REDIRECT` | — | Forbid HTTPS→HTTP redirects |
| `HOMEBREW_NO_VERIFY_ATTESTATIONS` | — | Skip bottle attestation verification |
| `HOMEBREW_VERIFY_ATTESTATIONS` | — | Enable bottle attestation verification |

## Networking and Sandbox

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_FORMULA_BUILD_NETWORK` | `allow` | Sandbox network for builds: `allow` or `deny` |
| `HOMEBREW_FORMULA_POSTINSTALL_NETWORK` | `allow` | Sandbox network for postinstall |
| `HOMEBREW_FORMULA_TEST_NETWORK` | `allow` | Sandbox network for tests |
| `HOMEBREW_NO_SANDBOX_LINUX` | — | Disable Linux sandbox |
| `HOMEBREW_AVOID_NESTED_SANDBOXING` | — | Skip sandbox when inside another sandbox |

## Output and Display

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_COLOR` | — | Force color output on non-TTY |
| `HOMEBREW_NO_COLOR` | `$NO_COLOR` | Disable color output |
| `HOMEBREW_NO_EMOJI` | — | Don't print install badge emoji |
| `HOMEBREW_INSTALL_BADGE` | 🍺 | Text printed before install summary |
| `HOMEBREW_VERBOSE` | — | Always assume `--verbose` |
| `HOMEBREW_VERBOSE_USING_DOTS` | — | Print `.` once per minute (avoid timeout kills) |
| `HOMEBREW_DEBUG` | — | Always assume `--debug` |
| `HOMEBREW_DISPLAY_INSTALL_TIMES` | — | Print install times |
| `HOMEBREW_NO_ENV_HINTS` | — | Don't print env var hints |
| `HOMEBREW_NO_PATH_SHADOW_CHECK` | — | Don't warn about shadowed executables |

## Developer Mode

| Variable | Default | Purpose |
|----------|---------|---------|
| `HOMEBREW_DEVELOPER` | — | Developer mode (warnings → errors, faster auto-update) |
| `HOMEBREW_DISABLE_DEBREW` | — | Disable interactive formula debugger |
| `HOMEBREW_DISABLE_LOAD_FORMULA` | — | Refuse to load formulae (for untrusted PRs) |
| `HOMEBREW_EVAL_ALL` | — | Evaluate all formulae and casks |
| `HOMEBREW_SORBET_RUNTIME` | — | Enable Sorbet runtime typechecking |
| `HOMEBREW_SORBET_RECURSIVE` | — | Enable recursive Sorbet typechecking |
| `HOMEBREW_SKIP_OR_LATER_BOTTLES` | — | (With developer) Don't use older macOS bottles |
| `HOMEBREW_SIMULATE_MACOS_ON_LINUX` | — | Simulate macOS code paths on Linux |
| `HOMEBREW_FAIL_LOG_LINES` | `15` | Output lines on formula failure |

## GitHub Integration

| Variable | Purpose |
|----------|---------|
| `HOMEBREW_GITHUB_API_TOKEN` | GitHub API token (for search, issue fetching) |
| `HOMEBREW_GITHUB_PACKAGES_TOKEN` | GitHub Packages Registry token |
| `HOMEBREW_GITHUB_PACKAGES_USER` | GitHub Packages username |
| `HOMEBREW_DOCKER_REGISTRY_TOKEN` | Docker registry bearer token |
| `HOMEBREW_DOCKER_REGISTRY_BASIC_AUTH_TOKEN` | Docker registry basic auth (base64) |
| `HOMEBREW_NO_GITHUB_API` | Disable GitHub API usage |

## Git Configuration

| Variable | Purpose |
|----------|---------|
| `HOMEBREW_GIT_NAME` | Git author name |
| `HOMEBREW_GIT_EMAIL` | Git author email |
| `HOMEBREW_GIT_COMMITTER_NAME` | Git committer name |
| `HOMEBREW_GIT_COMMITTER_EMAIL` | Git committer email |
| `HOMEBREW_GIT_PATH` | (Linux) git executable path |
| `HOMEBREW_SSH_CONFIG_PATH` | SSH config file path (default `~/.ssh/config`) |

## Tools and Editors

| Variable | Purpose |
|----------|---------|
| `HOMEBREW_EDITOR` | Editor for `brew edit` (default `$EDITOR` or `$VISUAL`) |
| `HOMEBREW_BROWSER` | Browser for `brew home` (default `$BROWSER` or OS default) |
| `HOMEBREW_DISPLAY` | X11 display for Linux |
| `HOMEBREW_SVN` | svn binary path |
| `HOMEBREW_BAT` | Use `bat` for `brew cat` |
| `HOMEBREW_BAT_CONFIG_PATH` | bat config file |
| `HOMEBREW_BAT_THEME` | bat theme |
| `HOMEBREW_BOOTSNAP` | Use Bootsnap for faster repeated calls |
| `HOMEBREW_NO_BOOTSNAP` | Don't use Bootsnap |
| `HOMEBREW_FORCE_VENDOR_RUBY` | Use vendored Ruby even if system Ruby is new enough |
| `HOMEBREW_FORCE_BREWED_CURL` | Always use Homebrew curl |
| `HOMEBREW_FORCE_BREWED_GIT` | Always use Homebrew git |
| `HOMEBREW_FORCE_BREWED_CA_CERTIFICATES` | Always use Homebrew ca-certificates |
| `HOMEBREW_PRY` | Use Pry for `brew irb` |

## Install Behavior

| Variable | Purpose |
|----------|---------|
| `HOMEBREW_NO_INSTALL_UPGRADE` | Don't upgrade if installed but outdated |
| `HOMEBREW_NO_INSTALLED_DEPENDENTS_CHECK` | Don't check dependents after install/upgrade |
| `HOMEBREW_NO_ANALYTICS` | Don't send analytics |
| `HOMEBREW_NO_ASK` | Disable default ask mode (does not disable explicit `--ask`) |
| `HOMEBREW_ENV_SYNC_STRICT` | Only sync exact installed versions |
| `HOMEBREW_SBOM` | Write SBOM files for source installs (hidden variable, not in public manpage; default `true`) |

## Proxy Variables

Standard proxy variables used by `curl`, `git`, and `svn` within Homebrew:

| Variable | Purpose |
|----------|---------|
| `all_proxy` | SOCKS5 proxy |
| `http_proxy` / `https_proxy` / `ftp_proxy` | Protocol-specific proxies |
| `no_proxy` | Comma-separated hosts excluded from proxying |
| `SUDO_ASKPASS` | Pass `-A` to sudo |

## Diagnostic Commands

### `brew config` / `brew --config`

Displays current Homebrew and system configuration — useful for bug reports.

### `brew doctor`

Checks system for potential problems. Exits non-zero if problems found.

| Flag | Purpose |
|------|---------|
| `--list-checks` | List all audit methods |
| `-D, --audit-debug` | Enable debugging and profiling |

### `brew missing` [--hide=`<hidden>`] [`<formula|cask>` ...]

Checks for missing dependencies.

## How Environment Variables Are Filtered

Homebrew has **multiple levels of filtering**:

1. **Overall filtering** — the environment is filtered to avoid contamination breaking from-source builds. Only a select list of variables is kept, plus any prefixed with `HOMEBREW_`.
2. **Sensitive variable removal** — credentials, keys, passwords, and tokens are stripped to prevent malicious subprocesses from obtaining them.

Environment variables must have a value set to be detected: `export HOMEBREW_NO_ANALYTICS=1` (not just `export HOMEBREW_NO_ANALYTICS`).

## References

- [Homebrew Manpage — Environment](https://docs.brew.sh/Manpage#environment)
- For maintenance: [maintenance.md](maintenance.md)
- For bottles: [bottle.md](bottle.md)
