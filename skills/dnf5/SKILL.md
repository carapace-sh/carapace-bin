---
name: dnf5
description: >
  Use when working with dnf5 — the next-generation RPM package manager for Fedora/RHEL.
  Covers CLI commands and options, configuration, plugin system (dnf5 active plugins and
  libdnf5 passive plugins), dnf5daemon D-Bus API, execution workflow, package specs (NEVRA),
  modularity, caching, migration from DNF4, libdnf5 library architecture, build system, and
  command-line aliases. Triggers on: "dnf5", "dnf", "libdnf5", "libdnf", "dnf5daemon",
  "dnf5-plugins", "libdnf5-plugins", "dnf5.conf", "repoquery", "versionlock", "distro-sync",
  "offline upgrade", "system-upgrade", "advisory", "comps", "modulemd", "NEVRA",
  "installroot", "forcearch", "repo variables", "metadata_expire", "setopt", "setvar",
  "config-manager", "repofrompath", "releasever", "basearch", "pluginpath".
user-invocable: true
---

# DNF5 In-Depth Reference

DNF5 is the next-generation RPM package manager for Fedora/RHEL, rewritten in C++ (C++20) for better performance and fewer external dependencies. It consists of the `dnf5` CLI, the `libdnf5` core library, the `libdnf5-cli` CLI helper, a two-tier plugin system, and `dnf5daemon` (D-Bus daemon).

## Data Flow

```
dnf5 <command> [options] [args]
  → logger setup, Base creation, command registration
    → plugin loading (dnf5 active + libdnf5 passive)
      → argument parsing, config loading
        → repo sack creation, repo loading, Vars substitution
          → goal resolution, transaction table
            → download, GPG check, RPM transaction
              → system state update
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| CLI, commands, options, flags, synopsis, exit codes, environment variables, assumeyes, assumeno, quiet, exclude, repo, setopt, setvar, repofrompath, releasever, installroot, forcearch, refresh, cacheonly, color, version, comment, dump-main-config | [references/cli.md](references/cli.md) |
| configuration, dnf.conf, config options, main section, repo options, color options, option types, boolean, integer, storage size, drop-in directories, repos.override.d, varsdir, reposdir, pluginconfpath, pluginpath, cachedir, system_cachedir, persistdir, protected_packages, installonlypkgs, tsflags, metadata_expire, zchunk, multilib_policy, obsoletes, best, skip_broken, skip_unavailable, allow_downgrade, allow_vendor_change, vendor change policy, group_package_types, optional_metadata_types | [references/configuration.md](references/configuration.md) |
| workflow, lifecycle, execution flow, hooks, pre_base_setup, post_base_setup, repos_configured, repos_loaded, pre_add_cmdline_packages, post_add_cmdline_packages, goal_resolved, pre_transaction, post_transaction, finish, init, create_commands, pre_configure, configure, load_additional_packages, run, goal_resolved step | [references/workflow.md](references/workflow.md) |
| plugins, dnf5 plugins, libdnf5 plugins, active plugins, passive plugins, IPlugin, IPlugin2_1, plugin API versioning, PluginAPIVersion, PluginVersion, LibraryVersion, DNF5_PLUGINS_DIR, LIBDNF_PLUGINS_CONFIG_DIR, actions plugin, appstream plugin, expired-pgp-keys, rhsm, local plugin, python_plugins_loader, CMakeLists.txt, dnf5_plugin_get_api_version, libdnf_plugin_get_api_version | [references/plugins.md](references/plugins.md) |
| daemon, dnf5daemon, D-Bus, dbus, SessionManager, open_session, close_session, polkit, PolicyKit, org.rpm.dnf.v0, Rpm interface, Repo interface, Goal interface, Base interface, Offline interface, Advisory interface, History interface, comps.Group interface, transaction signals, download signals, dnf5daemon-server, dnf5daemon-client | [references/daemon.md](references/daemon.md) |
| specs, NEVRA, NEVR, provides, file provides, binaries, globs, glob patterns, package-spec, group-spec, environment-spec, module-spec, NSVCA, transaction-spec, partial matching, epoch, version, release, arch, noarch, src, rich dependencies | [references/specs.md](references/specs.md) |
| modularity, modulemd, module stream, active stream, inactive stream, default stream, enabled stream, module_hotfixes, modular filtering, demodularized, orphaned modular packages, fail-safe, modulefailsafe, module_platform_id, module_stream_switch | [references/modularity.md](references/modularity.md) |
| caching, cache, metadata, cacheonly, keepcache, build_cache, solv, repodata, temporary_files.toml, metadata_expire, optional metadata, filelists, comps, updateinfo, presto, other, metadata types, cache sharing, copy-on-write | [references/caching.md](references/caching.md) |
| migration, dnf4, DNF4, changes from dnf, dropped options, renamed options, new options, command changes, API changes, configuration changes, strict deprecated, deltarpm deprecated, best default true, cacheonly enum, migrating to dnf5, Fedora migration | [references/migration.md](references/migration.md) |
| libdnf5, library architecture, Base, Goal, Transaction, PackageQuery, PackageSet, Package, Repo, RepoSack, RepoQuery, ConfigMain, ConfigRepo, ConfigParser, Vars, advisory, comps, rpm, module, transaction history, system state, NEVRA, reldep, arch, changelog, checksum, rpm_signature, versionlock_config, solv, libsolv | [references/libdnf5.md](references/libdnf5.md) |
| build, CMake, CMakeLists.txt, build options, WITH_DNF5, WITH_DNF5DAEMON_SERVER, WITH_MODULEMD, WITH_TESTS, WITH_PYTHON3, WITH_PERL5, WITH_RUBY, WITH_GO, WITH_PLUGIN_ACTIONS, WITH_PLUGIN_RHSM, sanitizers, clang-format, clang-tidy, dependencies, libsolv, librepo, sdbus-cpp, fmt, librpm | [references/build.md](references/build.md) |
| aliases, TOML, command alias, cloned_named_arg, named_arg, group, aliases.d, value_help, attached_command, attached_named_args, required_values, complete, group_id, descr | [references/aliases.md](references/aliases.md) |
| comps, groups, environments, mandatory, default, optional, conditional, group_package_types, with-optional, no-packages, group install, group remove, environment install | [references/comps.md](references/comps.md) |
| filtering, excludepkgs, includepkgs, disable_excludes, user excludes, versionlock filtering, modular filtering, module_hotfixes | [references/filtering.md](references/filtering.md) |
| system-state, system_state_dir, TOML, package reasons, install reason, user, dependency, weak dependency, group, external, corrupted state, regeneration | [references/system-state.md](references/system-state.md) |
| installroot, --installroot, use-host-config, path resolution, chroot, releasever detection, module_platform_id | [references/installroot.md](references/installroot.md) |
| forcearch, --forcearch, architecture override, arch, basearch, qemu-user-static | [references/forcearch.md](references/forcearch.md) |
| repo variables, $arch, $basearch, $releasever, $releasever_major, $releasever_minor, DNF_VAR_, varsdir, variable files, DNF0-DNF9, vars substitution | [references/repo-variables.md](references/repo-variables.md) |

## Quick Guide

- **What CLI commands and options does dnf5 have?** → [references/cli.md](references/cli.md)
- **How do I configure dnf5?** → [references/configuration.md](references/configuration.md)
- **What is the execution lifecycle of a dnf5 command?** → [references/workflow.md](references/workflow.md)
- **How do I write a dnf5 or libdnf5 plugin?** → [references/plugins.md](references/plugins.md)
- **How does the dnf5daemon D-Bus API work?** → [references/daemon.md](references/daemon.md)
- **What package spec formats does dnf5 accept?** → [references/specs.md](references/specs.md)
- **How does modularity work in dnf5?** → [references/modularity.md](references/modularity.md)
- **How does caching work?** → [references/caching.md](references/caching.md)
- **What changed from DNF4 to DNF5?** → [references/migration.md](references/migration.md)
- **What is the libdnf5 library architecture?** → [references/libdnf5.md](references/libdnf5.md)
- **How do I build dnf5 from source?** → [references/build.md](references/build.md)
- **How do command-line aliases work?** → [references/aliases.md](references/aliases.md)
- **How do comps groups and environments work?** → [references/comps.md](references/comps.md)
- **How does package filtering work?** → [references/filtering.md](references/filtering.md)
- **What is the dnf5 system state?** → [references/system-state.md](references/system-state.md)
- **How does --installroot work?** → [references/installroot.md](references/installroot.md)
- **How does --forcearch work?** → [references/forcearch.md](references/forcearch.md)
- **What repo variables are available?** → [references/repo-variables.md](references/repo-variables.md)

## Cross-Project References

- For **RPM** package format internals (spec files, macros, headers), see the RPM project documentation at <https://rpm.org/>.
- For **libsolv** (the dependency solver used by dnf5), see <https://github.com/openSUSE/libsolver>.
- For **librepo** (repository downloading library), see <https://github.com/rpm-software-management/librepo>.
- For **modulemd** specification, see <https://github.com/fedora-modularity/libmodulemd>.
- For the existing **dnf completer** in carapace-bin, see `completers/linux/dnf_completer/`.
