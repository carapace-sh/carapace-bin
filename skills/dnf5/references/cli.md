# DNF5 CLI Reference

The `dnf5` command-line interface: synopsis, global options, built-in commands, plugin commands, exit codes, and environment variables.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/dnf5.8.html>. For configuration options, see [configuration.md](configuration.md).

## Synopsis

```
dnf5 <command> [options] [<args>...]
```

## Global Options

These options are available for all (or most) commands. Some options only apply to transactional commands.

| Option | Description |
|--------|-------------|
| `--assumeno` | Automatically answer no for all questions. |
| `--best` | Try the best available package versions in transactions. Forces DNF5 to only consider the latest packages. |
| `-C, --cacheonly` | Use only cached data for packages and repo metadata. Cache won't be updated even if expired. Sets `cacheonly=all`. |
| `--comment=COMMENT` | Add a comment to the transaction history. |
| `--config=CONFIG_FILE_PATH` | Define configuration file location. |
| `--debugsolver` | Dump additional data from solver for debugging. Saved in `./debugdata`. |
| `--disable-plugin=PLUGIN_NAME,...` | Disable specified libdnf5 library plugins. Accepts names or globs. List option (can be specified multiple times). |
| `--disable-repo=REPO_ID,...` | Temporarily disable active repositories. Accepts ids or globs. List option. Errors on invalid REPO_ID. |
| `--dump-main-config` | Print main configuration values to stdout. |
| `--dump-repo-config=REPO_ID,...` | Print repository configuration values to stdout. List option. |
| `--dump-variables` | Print variable values to stdout. |
| `--enable-plugin=PLUGIN_NAME,...` | Enable specified libdnf5 library plugins. List option. |
| `--enable-repo=REPO_ID,...` | Temporarily enable additional repositories. List option. |
| `--forcearch=ARCH` | Force the use of a specific architecture. See [forcearch.md](forcearch.md). |
| `-h, --help` | Show the help. |
| `--installroot=ABSOLUTE_PATH` | Setup installroot path. See [installroot.md](installroot.md). |
| `--no-best` | Do not limit the transaction to the best candidates only. |
| `--no-docs` | Do not install documentation files (man pages, texinfo). Sets `RPMTRANS_FLAG_NODOCS`. |
| `--no-gpgchecks` | Skip checking OpenPGP signatures on packages (if RPM policy allows). |
| `--no-plugins` | Disable all libdnf5 plugins. |
| `-q, --quiet` | Show just relevant content. Suppresses state/action messages. |
| `--refresh` | Force refreshing metadata before running the command. |
| `--color=<color>` | Control color output. Values: `always`, `never`, `auto` (default). |
| `--repo=REPO_ID,...` | Enable just specified repositories. List option. |
| `--repofrompath=REPO_ID,REPO_PATH` | Add a repository for this run only. Can be used multiple times. Variables in both values are substituted. |
| `--releasever=RELEASEVER` | Override the distribution release version. Also sets `releasever_major` and `releasever_minor` by splitting on first `.`. |
| `--releasever-major=RELEASEVER_MAJOR` | Override the `releasever_major` variable. Must be specified after `--releasever`. |
| `--releasever-minor=RELEASEVER_MINOR` | Override the `releasever_minor` variable. Must be specified after `--releasever`. |
| `--setopt=[REPO_ID.]OPTION=VALUE` | Override a configuration option. List-type options (`excludepkgs`, `includepkgs`, `installonlypkgs`, `tsflags`) are appended to, not replaced. Empty value clears. |
| `--setvar=VAR_NAME=VALUE` | Override a DNF5 variable value. |
| `--skip-file-locks` | Skip acquiring file locks (e.g., system repository lock). |
| `--show-new-leaves` | Show newly installed leaf packages and packages that became leaves after a transaction. |
| `--use-host-config` | Use config files and variable definitions from host system rather than installroot. |
| `--version` | Display version of dnf5 and libdnf5, plus loaded plugins. |
| `-y, --assumeyes` | Automatically answer yes for all questions. |
| `-x PACKAGE-SPEC-N,..., --exclude=PACKAGE-SPEC-N,...` | Exclude packages from the transaction. List option. |

## Built-in Commands

| Command | Description |
|---------|-------------|
| `advisory` | Manage advisories (security updates). |
| `autoremove` | Remove unneeded packages. |
| `check` | Check for problems in the package database. |
| `check-upgrade` | Check for available package upgrades. |
| `clean` | Remove or invalidate cached data. |
| `distro-sync` | Upgrade or downgrade installed packages to the latest available version. |
| `downgrade` | Downgrade packages. |
| `download` | Download packages. |
| `environment` | Manage comps environments. |
| `group` | Manage comps groups. |
| `history` | Manage transaction history. |
| `info` | Provide detailed information about installed or available packages. |
| `install` | Install packages. |
| `leaves` | List groups of leaf packages. |
| `list` | List installed or available packages. |
| `makecache` | Generate the metadata cache. |
| `mark` | Change the reason of an installed package. |
| `module` | Manage modules. |
| `offline` | Manage offline transactions. |
| `provides` | Find what package provides the given value. |
| `reinstall` | Reinstall packages. |
| `remove` | Remove packages. |
| `replay` | Replay stored transactions. |
| `repo` | Manage repositories. |
| `repoquery` | Search for packages in repositories. |
| `search` | Search for packages using keywords. |
| `swap` | Remove software and install another in a single transaction. |
| `system-upgrade` | Upgrade the system to a new major release. |
| `upgrade` | Upgrade packages. |
| `versionlock` | Protect packages from updates to newer versions. |

## Plugin Commands

Available after installing the `dnf5-plugins` package:

| Plugin Command | Description |
|----------------|-------------|
| `automatic` | Alternative CLI for automated execution (systemd timers, cron). |
| `builddep` | Install missing build dependencies for an RPM package. |
| `changelog` | Show package changelogs. |
| `config-manager` | Manage main config, repository config, and variables. |
| `copr` | Manage Copr (community/third-party) repositories. |
| `manifest` | Generate package manifest. |
| `needs-restarting` | Determine whether the system should be rebooted. |
| `repoclosure` | Display unresolved dependencies for repositories. |
| `repomanage` | Manage older/newer packages in a repository directory. |
| `reposync` | Synchronize remote repository to a local directory. |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Operation was successful. |
| 1 | An error occurred during processing of the command. |
| 2 | An error occurred during parsing the arguments. |

Individual commands may define additional exit codes.

## Environment Variables

| Variable | Description |
|----------|-------------|
| `DNF5_FORCE_INTERACTIVE` | `0` = pretend I/O is non-interactive (no questions). `1` = force interactive questions even on non-interactive terminal. |
| `DNF5_PLUGINS_DIR` | Override directory for dnf5 application plugins. Empty string disables plugin loading. |
| `LIBDNF_PLUGINS_CONFIG_DIR` | Override directory for libdnf5 plugin configuration files. |
| `SOURCE_DATE_EPOCH` | Timestamp (seconds since UNIX epoch). Transaction timestamps set to this instead of current time. Packages ordered for consistency. |
| `TEMP`, `TEMPDIR`, `TMP`, `TMPDIR` | Override path to temp files directory. Must exist. Default: `/tmp`. |

## Files

| Category | Path |
|----------|------|
| Cache Files | `/var/cache/libdnf5/` |
| Main Configuration | `/etc/dnf/dnf.conf` |
| Repository Configuration | `/etc/yum.repos.d/` |
| Repository Persistence | `/var/lib/dnf/` |
| System State | `/usr/lib/sysimage/libdnf5/` |

## Metadata Synchronization

DNF5 uses `metadata_expire` to determine whether local repo data is due for re-sync. Commands where up-to-date metadata is not critical (e.g., `group list`) use whatever version is locally available, ignoring expiration. See [caching.md](caching.md) for details.

## Configuration File Replacement Policy

Updated packages may replace old modified config files or keep them. Neither file is actually replaced — RPM gives the conflicting file an additional suffix. The true name is determined by the package itself following packaging guidelines.
