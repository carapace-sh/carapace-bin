# DNF5 Configuration Reference

All configuration options, option types, drop-in directories, repo override directories, and file paths for DNF5.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/dnf5.conf.5.html>. For CLI options that override config, see [cli.md](cli.md).

## Main Configuration Options (`[main]`)

| Option | Type | Description | Default |
|--------|------|-------------|---------|
| `allow_downgrade` | boolean | Allow downgrading packages while resolving dependencies | `True` |
| `allow_vendor_change` | boolean | If disabled, sticks to original vendor during upgrades/downgrades | `True` |
| `assumeno` | boolean | Assume No for all questions | `False` |
| `assumeyes` | boolean | Assume Yes for all questions | `False` |
| `best` | boolean | Use highest available version or fail | `True` |
| `cachedir` | string | Path to cache directory (metadata + packages). Regular user: `~/.cache/libdnf5`. Superuser: overwritten by `system_cachedir`. | `~/.cache/libdnf5` |
| `cacheonly` | string | `all` = entirely from cache; `metadata` = cache metadata only; `none` = normal | `none` |
| `check_config_file_age` | boolean | Auto-expire metadata of repos older than their config file | `True` |
| `clean_requirements_on_remove` | boolean | Remove unneeded dependencies during `remove` | `True` |
| `debugdir` | string | Location for libsolv debug files | `./debugdata` |
| `debug_solver` | boolean | Create libsolv debug files when solving | `False` |
| `defaultyes` | boolean | Default answer to confirmation prompts is Yes (still prompts) | `False` |
| `destdir` | string | Redirect downloaded packages to provided directory | `<repo cachedir>/packages` |
| `exclude_from_weak` | list | Prevent installing packages as weak dependencies (name or glob) | `[]` |
| `exclude_from_weak_autodetect` | boolean | Autodetect unmet weak dependencies; skip their providers | `true` |
| `excludeenvs` | list | Exclude environments by id or glob | `[]` |
| `excludegroups` | list | Exclude groups by id or glob | `[]` |
| `group_package_types` | list | Types of group packages installed | `default,mandatory,conditional` |
| `ignorearch` | boolean | Allow installing packages incompatible with CPU architecture | `False` |
| `installonlypkgs` | list | Packages that should only be installed, never upgraded | `dnf5,glob:/etc/dnf/protected.d/*.conf` |
| `installonly_limit` | integer | Max concurrent installonly packages. `0` = unlimited | `3` |
| `installroot` | string | Root filesystem for all packaging operations. Absolute path required. | `/` |
| `install_weak_deps` | boolean | Pull in weak dependency packages (Recommends/Supplements) | `True` |
| `keepcache` | boolean | Keep downloaded packages in cache | `False` |
| `logdir` | string | Directory for log files. User: `~/.local/state`. Root: `/var/log`. | `~/.local/state` |
| `log_rotate` | integer | Number of log rotations before removal. `0` = no rotation. | `4` |
| `log_size` | storage size | Rotate logs when bigger than this. `0` = no rotation. | `1M` |
| `module_platform_id` | string | Override `PLATFORM_ID` from `/etc/os-release`. Format: `$name:$stream`. | empty |
| `module_stream_switch` | boolean | Allow switching enabled streams of a module | `False` |
| `multilib_policy` | string | `best` = prefer best arch match; `all` = install all arches | `best` |
| `obsoletes` | boolean | Use obsoletes processing logic | `True` |
| `optional_metadata_types` | list | Types of metadata to load beyond `primary`/`modules` | `comps,updateinfo` |
| `persistdir` | string | Directory for persistent data between runs | `/var/lib/dnf` |
| `pluginconfpath` | list | Directories for libdnf5 plugin configurations | `/etc/dnf/libdnf5-plugins` |
| `pluginpath` | list | Directories for libdnf5 plugins | `/usr/lib64/libdnf5/plugins/` |
| `plugins` | boolean | Enable libdnf5 plugins | `True` |
| `protected_packages` | list | Packages DNF5 should never completely remove | `dnf5,glob:/etc/dnf/protected.d/*.conf` |
| `protect_running_kernel` | boolean | Protect the running kernel package from removal | `True` |
| `recent` | integer | Time period in days for `--recent` option | `7` |
| `reposdir` | list | Repository config file locations | `[/etc/yum.repos.d, /etc/distro.repos.d, /usr/share/dnf5/repos.d]` |
| `skip_broken` | boolean | Skip uninstallable packages during dependency resolution | `False` |
| `skip_system_repo_lock` | boolean | Skip acquiring lock on system repository (RPM database) | — |
| `skip_unavailable` | boolean | Skip unavailable packages during transaction preparation | `False` |
| `system_cachedir` | string | Superuser cache dir, overwrites `cachedir` | `/var/cache/libdnf5` |
| `system_state_dir` | string | System state files location | `/usr/lib/sysimage/libdnf5` |
| `transaction_history_dir` | string | History database location | Same as `system_state_dir` |
| `tsflags` | list | Extra flags for RPM transaction | empty |
| `use_host_config` | boolean | Use config/variables from host rather than installroot | `False` |
| `varsdir` | list | Directories for variable definition files | `/etc/dnf/vars` |
| `zchunk` | boolean | Use zchunk compression for repo metadata (if available) | `True` |

### `tsflags` Values

| Value | Description |
|-------|-------------|
| `noscripts` | Don't execute RPM scriptlets |
| `test` | Don't actually perform the transaction |
| `notriggers` | Don't run trigger scriptlets |
| `nodocs` | Don't install documentation files |
| `justdb` | Only update database, don't modify files |
| `nocontexts` | Don't set SELinux contexts (related to SELinux) |
| `nocaps` | Don't set file capabilities |
| `nocrypto` | Don't use crypto (useful for testing) |
| `deploops` | Allow dependency loops |
| `noplugins` | Don't run RPM plugins |

### `optional_metadata_types` Values

| Value | Description |
|-------|-------------|
| `comps` | Package groups and environment descriptions |
| `filelists` | All files provided by packages |
| `updateinfo` | Security-related updates and advisories |
| `presto` | Delta RPM information |
| `other` | Additional metadata (e.g., changelogs) |
| `all` | Load all metadata types |

## Color Options

| Option | Type | Description | Default |
|--------|------|-------------|---------|
| `color_list_available_upgrade` | color | Available packages newer than installed | `bold,blue` |
| `color_list_available_downgrade` | color | Available packages older than installed | `dim,magenta` |
| `color_list_available_reinstall` | color | Available packages identical to installed | `bold,green` |
| `color_list_available_install` | color | Available for installation, none installed | `bold,cyan` |
| `color_update_installed` | color | Removed packages (transaction display) | `dim,red` |
| `color_update_local` | color | Local packages from @commandline repo | `dim,green` |
| `color_update_remote` | color | Packages from remote repositories | `bold,green` |
| `color_search_match` | color | Patterns matched in search output | `bold,magenta` |

## Repo Options

| Option | Type | Description | Default |
|--------|------|-------------|---------|
| `baseurl` | list | List of URLs for the repository (tried in order) | `[]` |
| `cost` | integer | Relative cost of accessing repository. Lower = preferred. | `1000` |
| `enabled` | boolean | Include this repository as a package source | `True` |
| `gpgkey` | list | URLs of OpenPGP key files for signing | `[]` |
| `metalink` | string | URL of a metalink for the repository | `None` |
| `mirrorlist` | string | URL of a mirrorlist for the repository | `None` |
| `name` | string | Human-readable name | Defaults to repo ID |
| `priority` | integer | Priority. Lower = picked first (even if lower version). | `99` |
| `type` | string | Type of repo metadata. Supported: `rpm-md` (and aliases). | empty |

### Source/Debuginfo Naming Convention

For `<ID>-rpms`: source = `<ID>-source-rpms`, debuginfo = `<ID>-debug-rpms`.
For `<ID>` (without `-rpms`): source = `<ID>-source`, debuginfo = `<ID>-debuginfo`.

## Options for Both `[main]` and Repo Sections

These can be set globally in `[main]` (as default for all repos) and overridden per-repository:

| Option | Type | Description | Default |
|--------|------|-------------|---------|
| `bandwidth` | storage size | Total bandwidth available for downloading (used with `throttle`) | `0` |
| `build_cache` | boolean | Save libsolv cache from downloaded metadata | `True` |
| `countme` | boolean | HTTP "countme" flag for repository provider statistics | `False` |
| `disable_excludes` | list | Disable package/group/environment filtering. Values: repo IDs, `main`, `*` | — |
| `enablegroups` | boolean | Allow use of package groups | `True` |
| `excludepkgs` | list | Exclude packages by name or glob | `[]` |
| `fastestmirror` | boolean | Use TCP socket latency to find closest mirror | `False` |
| `pkg_gpgcheck` | boolean | OpenPGP signature check on packages in this repo (`gpgcheck` also supported) | `False` |
| `includepkgs` | list | Include only packages matching name/glob | `[]` |
| `ip_resolve` | ip address type | How host names are resolved: `4`/`IPv4`, `6`/`IPv6`, `whatever` | `whatever` |
| `localpkg_gpgcheck` | boolean | OpenPGP signature check on local packages | `False` |
| `max_parallel_downloads` | integer | Maximum simultaneous downloads (max 20) | `3` |
| `max_downloads_per_mirror` | integer | Max simultaneous downloads per mirror (max 20) | `3` |
| `metadata_expire` | time in seconds | Period after which remote repo metadata is checked for update. `-1`/`never` = never expire. | 48 hours |
| `minrate` | storage size | Low speed threshold. If slower for `timeout` seconds, aborts. | `1000` |
| `password` | string | Password for basic HTTP authentication | empty |
| `proxy` | string | URL of proxy server. Empty string or `_none_` disables. | empty |
| `proxy_username` | string | Username for proxy server | empty |
| `proxy_password` | string | Password for proxy server | empty |
| `proxy_auth_method` | string | Proxy auth method: `basic`, `digest`, `negotiate`, `ntlm`, `digest_ie`, `ntlm_wb`, `none`, `any` | `any` |
| `proxy_sslcacert` | string | CA file for proxy SSL verification | empty |
| `proxy_sslclientcert` | string | SSL client certificate for proxy | empty |
| `proxy_sslclientkey` | string | SSL client key for proxy | empty |
| `proxy_sslverify` | boolean | Verify proxy SSL certificates | `True` |
| `repo_gpgcheck` | boolean | OpenPGP signature check on repository metadata | `False` |
| `skip_if_unavailable` | boolean | Continue running and disable repo that couldn't sync | `False` |
| `sslcacert` | string | CA file for SSL verification | empty |
| `sslclientcert` | string | SSL client certificate | empty |
| `sslclientkey` | string | SSL client key | empty |
| `sslverify` | boolean | Verify remote SSL certificates | `True` |
| `throttle` | storage size | Limit download speed. Absolute or percentage of `bandwidth`. `0` = no throttling. | `0` |
| `timeout` | time in seconds | Seconds to wait for a connection before timing out | `30` |
| `username` | string | Username for basic HTTP authentication | empty |
| `user_agent` | string | User-Agent string for HTTP requests | `libdnf (NAME VERSION_ID; VARIANT_ID; OS.BASEARCH)` |

## Option Types

| Type | Description | Valid Values |
|------|-------------|--------------|
| **boolean** | Two possible values | `1`, `0`, `True`, `False`, `yes`, `no` |
| **color** | Color and modifiers, comma-separated | Colors: `black`, `blue`, `cyan`, `green`, `magenta`, `red`, `white`, `yellow`. Modifiers: `bold`, `blink`, `dim`, `normal`, `reverse`, `underline` |
| **integer** | Whole number | Any integer |
| **ip address type** | IP address resolution type | `4`, `IPv4`, `6`, `IPv6` |
| **list** | One or more strings separated by space or comma. Backslash `\` escapes. | Strings with escaping |
| **storage size** | Integer with a unit | `k`, `M`, `G` |
| **string** | Sequence of symbols/digits | Any string |
| **time in seconds** | Time in seconds; can be `-1` or `never` | Integer, `-1`, `never` |

## Drop-in Configuration Directories

DNF5 loads configuration from three sources, processed in this order:

1. **Distribution configs**: `/usr/share/dnf5/libdnf.conf.d/`
2. **User configs**: `/etc/dnf/libdnf5.conf.d/`
3. **Main config file**: `/etc/dnf/dnf.conf`

### Processing Rules

- Files from both directories are sorted alphabetically.
- If a file with the same name exists in both, only the user file is used (distribution file is **masked**).
- Options are applied in order — **last option wins**.
- The main config file `/etc/dnf/dnf.conf` is loaded last.

### Example Load Order

```
1. /etc/dnf/libdnf5.conf.d/20-user-settings.conf
2. /usr/share/dnf5/libdnf.conf.d/50-something.conf
3. /etc/dnf/libdnf5.conf.d/60-something.conf  (masks /usr/share version)
4. /etc/dnf/libdnf5.conf.d/80-user-settings.conf
5. /usr/share/dnf5/libdnf.conf.d/90-something.conf
6. /etc/dnf/dnf.conf
```

## Drop-in Repo Override Directories

After repo configurations are loaded from `.repo` files, overrides can be applied from:

| Directory | Description |
|-----------|-------------|
| `/etc/dnf/repos.override.d/` | User repos override directory |
| `/usr/share/dnf5/repos.override.d/` | Distribution repos override directory |

### Rules

- Same format as repo config files (INI with `[repo_id]` sections).
- **Cannot create new repositories** — only modify existing ones.
- Support **globs** in repository ID for bulk modifications.
- Files with same name: `/etc/` version overrides `/usr/share/` version (masking).
- Applied in alphabetical order — **last option wins**.

### Example

```ini
# Enable skip_if_unavailable for all repositories
[*]
skip_if_unavailable = true

# Disable skip_if_unavailable for repos with id prefix "fedora"
[fedora*]
skip_if_unavailable = false
```

## Vendor Change Policy

Fine-tune vendor change behavior beyond the boolean `allow_vendor_change`. Supported versions: v1.0 and v1.1. Policy files define which vendors are allowed to replace which, using structured configuration. See <https://dnf5.readthedocs.io/en/stable/dnf5.conf-vendorpolicy.5.html>.

## All File and Directory Paths

### Configuration

| Path | Description |
|------|-------------|
| `/etc/dnf/dnf.conf` | Main configuration file |
| `/etc/dnf/libdnf5.conf.d/` | User drop-in config directory |
| `/usr/share/dnf5/libdnf.conf.d/` | Distribution drop-in config directory |
| `/etc/dnf/libdnf5-plugins` | Plugin configuration directory |

### Repositories

| Path | Description |
|------|-------------|
| `/etc/yum.repos.d/` | YUM repositories directory (primary) |
| `/etc/distro.repos.d/` | Distribution repositories directory |
| `/usr/share/dnf5/repos.d/` | Distribution repositories directory |
| `/etc/dnf/repos.override.d/` | User repos override directory |
| `/usr/share/dnf5/repos.override.d/` | Distribution repos override directory |

### Variables

| Path | Description |
|------|-------------|
| `/etc/dnf/vars/` | Variables directory |
| `/etc/yum/vars/` | YUM compatibility variables directory |
| `/usr/share/dnf5/vars.d/` | Distribution variables directory |

### Cache and Data

| Path | Description |
|------|-------------|
| `~/.cache/libdnf5/` | Cache directory (regular user) |
| `/var/cache/libdnf5` | Cache directory (superuser, via `system_cachedir`) |
| `/var/lib/dnf` | Persistent data directory (`persistdir`) |
| `/usr/lib/sysimage/libdnf5` | System state directory |
| `./debugdata` | Debug solver output (default) |

### Logs

| Path | Description |
|------|-------------|
| `~/.local/state` | Log directory (regular user) |
| `/var/log` | Log directory (superuser) |

### Plugins

| Path | Description |
|------|-------------|
| `/usr/lib64/libdnf5/plugins/` | Plugin loading directory |
| `/etc/dnf/libdnf5-plugins` | Plugin configuration directory |

### Protected Packages

| Path | Description |
|------|-------------|
| `/etc/dnf/protected.d/` | Directory for protected package config files (glob: `*.conf`) |
