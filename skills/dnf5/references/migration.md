# Migration from DNF4 to DNF5

All CLI, API, and configuration changes from DNF4 to DNF5. Covers dropped options, renamed options, new options, command changes, and configuration changes.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/changes_from_dnf4.7.html> and <https://dnf5.readthedocs.io/en/stable/migrating_to_dnf5.7.html>.

## Fedora Migration History

| Fedora Version | Change |
|----------------|--------|
| 38 | microdnf based on libdnf5; dnf5 and libdnf5 introduced alongside existing dnf |
| 39 | Modularity support discontinued in dnf5 |
| 40 | Build system (mock, koji, copr) switched to dnf5; filelists no longer downloaded by default |
| 41 | dnf5 becomes default; `/usr/bin/dnf` symlink -> dnf5; state files differ between dnf4 and dnf5; transaction history not migrated |

## CLI Changes

### Global Options Scoping

Options that cannot be applied to all commands are moved from global options to specific commands only. Example: `--best`/`--no-best` are now only relevant to several transaction commands.

### Options Renaming

Boolean options renamed to `--<option>`/`--no-<option>` or `--enable-<option>`/`--disable-<option>` format. Original names retained as compatibility aliases.

### Strict Behavior

- `--disable-repo=REPO_ID` and `--setopt=[REPO_ID.]OPTION=VALUE` now **error** on invalid `REPO_ID` (aligned with `--repo` and `--enable-repo`).
- The `strict` config option is no longer considered.

### No Value Separator After Short Options

Short options don't have a separator between option name and value. `dnf -x package` works, but `dnf -x=package` means `--exclude =package`.

### Dropped Options

| Option | Replacement |
|--------|-------------|
| `-4/-6` | Dropped. Use `ip_resolve` config option. |
| `--disableexcludes` / `--disableexcludepkgs` | Dropped. Use `disable_excludes` config option. |
| `--disable` / `--enable` (and `--set-disabled` / `--set-enabled`) | Dropped. Use `dnf5 config-manager setopt repo.enabled=0`. |
| `--downloaddir` | Dropped. Use `--destdir` for `download` command. |
| `-e, --errorlevel` | Dropped (both option and config). |
| `--help-cmd` | Dropped. Use `-h` or `--help`. |
| `--noautoremove` | Now only for `remove` command. Workaround: `--setopt=clean_requirements_on_remove=False`. |
| `--obsoletes` | Dropped. Use `obsoletes` config option. |
| `-R, --randomwait` | Dropped. |
| `--rpmverbosity` | Dropped. Use `rpmverbosity` config option. |
| `-v, --verbose` | Not implemented. May be added for specific commands in the future. |

### Renamed Options

| Old Name | New Name |
|----------|----------|
| `--sec-severity` | `--advisory-severities` |

### Changed Behavior

| Option | Change |
|--------|--------|
| `--installroot` | New behavior for where configuration and variables are loaded from. |
| `--version` | Behavior is different. |

### New Options

| Option | Description |
|--------|-------------|
| `--allow-downgrade` / `--no-allow-downgrade` | Enable/disable downgrade of dependencies. For `install`, `upgrade`, etc. |
| `--dump-main-config` | Print configuration values on stdout. |
| `--dump-repo-config=REPO_ID` | Print repo configuration values. |
| `--offline` | Store transaction for offline execution. For all transactional commands. |
| `--show-new-leaves` | Show newly installed leaf packages. |
| `--skip-unavailable` | Skip packages not available in repos. For `install`, `upgrade`, etc. |
| `--use-host-config` | Use config/variables from host rather than installroot. |

## Command Changes

### Optional Subcommands Removed

Commands cannot have optional subcommands. Subcommands are now mandatory.

- Before: `dnf history <transaction ID>` -> Now: `dnf history info <transaction ID>`
- Before: `dnf updateinfo` -> Now: `dnf updateinfo summary`

### Per-Command Changes

| Command | Changes |
|---------|---------|
| `alias` | Dropped. Replaced by TOML-based aliases. |
| `automatic` | Now a DNF5 plugin. Only one `dnf5-automatic` timer shipped. |
| `autoremove` | Dropped `<spec>` positional arg. Dropped variants `autoremove-n/na/nevra`. |
| `builddep` | `--spec` and `--srpm` now only apply to following arguments. |
| `config-manager` | Parameters replaced by subcommands (`addrepo`, `setopt`). Drop-in override files instead of modifying existing. Dropped `--dump`. |
| `debuginfo-install` | No longer supports autoupdate. Use `config-manager` to enable debug repos. |
| `distro-sync` | Fails on unmatched args (use `--skip-unavailable`). Dropped aliases. |
| `downgrade` | Fails on unmatched args (use `--skip-unavailable`). |
| `download` | `--source` renamed to `--srpm`. Dropped `--downloaddir`. |
| `group` | New `--contains-pkgs`. Dropped `--ids`. Dropped `group mark install/remove` (use `--no-packages`). Dropped aliases (`groupinstall`, etc.). |
| `help` | Dropped. Use `--help`. |
| `history` | Subcommands mandatory. `userinstalled` dropped (use `repoquery --userinstalled`). `store` creates a directory. `replay` is now standalone. |
| `info` | Dropped `--all` and `--updates`. |
| `install` | Dropped variants `install-n/na/nevra`. |
| `list` | Dropped `--all`. `--available` behavior changed. Prints repo for installed packages. |
| `makecache` | Different cache dirs. Dropped `--timer`. Avoids downloading when possible. |
| `mark` | Renamed: `install`->`user`, `remove`->`dependency`. New `weak` subcommand. |
| `module` | Dropped `--all`. |
| `needs-restarting` | No longer scans open files. Default = old `--reboothint`. Dropped `-r` and `-u`. |
| `offline-distrosync` | Alias of `dnf5 distro-sync --offline`. |
| `offline-upgrade` | Alias of `dnf5 upgrade --offline`. |
| `remove` | No longer removes by provides (NEVRA/file provide only). Dropped variants and aliases. |
| `repoclosure` | Dropped `--pkg`. Use positional args. |
| `reposync` | Dropped `--downloadcomps`. Use `--download-metadata`. |
| `repolist`/`repoinfo` | Now subcommands of `repo`: `repo list` and `repo info`. |
| `repoquery` | Dropped `-a/--all`, `--alldeps`, `--nevra`, `--envra`, `--nvr`, `--unsatisfied`. Dropped `--archlist` alias. Dropped `-f` alias. `--resolve` -> `--providers-of=ATTRIBUTE`. `--queryformat` no longer adds trailing newline. Dropped `size` tag. `--source` -> `--sourcerpm`. |
| `shell` | Superseded by `do` command. No interactive shell or file redirect. |
| `system-upgrade` | Moved from plugin to built-in command. |
| `upgrade` | New `--minimal`. Fails on unmatched args. Dropped aliases. |
| `updateinfo` | Renamed to `advisory`. Subcommands mandatory. `--summary`/`--list`/`--info` -> subcommands. |
| `versionlock` | New format of configuration file. |

## API Changes

### Notable Changes

| Change | Details |
|--------|---------|
| `PackageSet::operator[]` | **Removed** (O(n^2) performance). Use `PackageSet` iterator. |
| `Package::get_epoch()` | Return type changed from `unsigned long` to `std::string`. |
| `Package.size` / `dnf_package_get_size()` | Ambiguous. Use `get_download_size()` and `get_install_size()`. |
| `dnf_sack_set_installonly` etc. | **Dropped**. Installonly from main Conf in Base. |
| `HY_PKG_UPGRADES_BY_PRIORITY` etc. | Priority filter separated. Use `filter_priority()` + `filter_latest_evr()`. |
| `HY_PKG_LATEST` | Replaced with `filter_latest_evr()`. |
| `proxy_auth_method()` | Return type changed from `OptionEnum<string>` to `OptionStringSet`. Multiple auth methods can be combined. |

## Configuration Changes

### Deprecated Options

| Option | Replacement |
|--------|-------------|
| `strict` | Split into `skip_broken` + `skip_unavailable`. Plus CLI `--skip-broken` and `--skip-unavailable`. |
| `metadata_timer_sync` | Obsoleted by `dnf5-makecache.timer` systemd timer. |
| `retries` | Deprecated. Was for deltarpm error retries. |
| `deltarpm` | Deprecated. Delta RPM support not planned. |
| `deltarpm_percentage` | Deprecated. Delta RPM support not planned. |

### Dropped Options

| Option | Details |
|--------|---------|
| `arch` / `basearch` | Use `--forcearch` instead. |
| `errorlevel` | Deprecated in DNF < 5, now dropped. |

### Changed Options

| Option | Change |
|--------|--------|
| `best` | Default changed to **`true`**. |
| `cachedir` | Paths changed. User: `~/.cache/libdnf5`. Root: `/var/cache/libdnf5`. |
| `cacheonly` | Changed from `bool` to **enum**: `all`, `metadata`, `none`. |
| `disable_excludes` | Use `*` instead of `all` to disable all excludes. |
| `keepcache` | Behavior slightly modified. |
| `optional_metadata_types` | Default changed to `comps,updateinfo`. |

### New Options

| Option | Description |
|--------|-------------|
| `allow_downgrade` | Enable/disable downgrade of dependencies. |
| `skip_broken` | Skip uninstallable packages. |
| `skip_unavailable` | Skip unavailable packages. |
