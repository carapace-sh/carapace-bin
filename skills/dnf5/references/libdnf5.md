# libdnf5 Library Architecture

The core library of DNF5, written in C++20. Provides the package management engine used by the `dnf5` CLI, `dnf5daemon`, and language bindings.

> **Source of truth**: <https://github.com/rpm-software-management/dnf5/tree/main/libdnf5> and <https://github.com/rpm-software-management/dnf5/tree/main/include/libdnf5>.

## Overview

`libdnf5` is the core library implementing all package management logic: configuration, repository management, RPM package handling, dependency resolution (via libsolv), transaction execution, history, advisories, comps, and modularity. Public headers are in `include/libdnf5/`, implementation in `libdnf5/`.

## Core Subsystems

| Directory | Purpose | Key Files |
|-----------|---------|-----------|
| `base/` | Core Base class, Goals, Transactions | `base.cpp`, `goal.cpp`, `transaction.cpp`, `transaction_package.cpp`, `solver_problems.cpp`, `log_event.cpp` |
| `conf/` | Configuration system | `config.cpp`, `config_main.cpp`, `config_parser.cpp`, `option.cpp`, `option_bool.cpp`, `option_enum.cpp`, `option_number.cpp`, `option_path.cpp`, `option_seconds.cpp`, `option_string.cpp`, `option_string_list.cpp`, `option_binds.cpp`, `vars.cpp` |
| `repo/` | Repository management | `repo.cpp`, `repo_sack.cpp`, `repo_cache.cpp`, `repo_downloader.cpp`, `repo_query.cpp`, `config_repo.cpp`, `repo_pgp.cpp`, `file_downloader.cpp`, `package_downloader.cpp`, `librepo.cpp`, `solv_repo.cpp` |
| `rpm/` | RPM package handling | `package.cpp`, `package_query.cpp`, `package_sack.cpp`, `package_set.cpp`, `nevra.cpp`, `reldep.cpp`, `arch.cpp`, `changelog.cpp`, `checksum.cpp`, `rpm_signature.cpp`, `transaction.cpp`, `versionlock_config.cpp` |
| `transaction/` | Transaction history & persistence | `transaction.cpp`, `transaction_history.cpp`, `transaction_item.cpp`, `transaction_sr.cpp`, `offline.cpp` |
| `advisory/` | Advisory (errata/updateinfo) handling | Advisory query and management |
| `comps/` | Comps groups & environments | Group and environment management |
| `module/` | Modulemd module support | Module stream management |
| `plugin/` | libdnf5 plugin infrastructure | `iplugin.cpp`, `plugins.cpp` |
| `logger/` | Logging system | Multi-logger support |
| `system/` | System-level operations | OS state operations |
| `solv/` | libsolv wrapper utilities | Solver integration |
| `utils/` | General utilities | Filesystem, string, etc. |

## The Base Object

`libdnf5::Base` is the central context object. It owns:

- Configuration (`ConfigMain`)
- Logger(s)
- Repo sack (`RepoSack`)
- Package sack (`PackageSack`)
- Plugin manager
- Vars (variable substitution)

All operations start by creating and configuring a `Base` object. See [workflow.md](workflow.md) for the setup sequence.

## Goal and Transaction

### Goal

`libdnf5::base::Goal` accepts operations (install, upgrade, remove, downgrade, reinstall, distro-sync) and resolves them through libsolv. The resolution produces a `Transaction` object.

### Transaction

`libdnf5::base::Transaction` represents the resolved set of operations. It provides:
- List of transaction packages (`TransactionPackage`)
- List of transaction groups (`TransactionGroup`)
- List of transaction environments (`TransactionEnvironment`)
- List of transaction modules (`TransactionModule`)
- Solver problems (`SolverProblems`)
- Log events (`LogEvent`)

### TransactionItem Actions

Actions include: `install`, `upgrade`, `downgrade`, `reinstall`, `remove`, `reason change`, `reason change group`, `installable`, `removed`, `upgraded`, `downgraded`, `reinstalled`, `replaced`.

## Package Query and PackageSet

### PackageQuery

`libdnf5::rpm::PackageQuery` provides filtering methods:
- `filter_name()`, `filter_arch()`, `filter_evr()`
- `filter_provides()`, `filter_requires()`, `filter_conflicts()`
- `filter_latest_evr()` (replaces `HY_PKG_LATEST`)
- `filter_priority()` (separated from combined filter)
- `filter_upgrades()`, `filter_downgrades()`, `filter_obsoletes()`
- `filter_installed()`, `filter_available()`
- Chainable: queries can be combined and filtered sequentially

### PackageSet

`libdnf5::rpm::PackageSet` is a set of packages. `operator[]` was removed (O(n^2) performance). Use iterators instead.

## Repository Management

### RepoSack

`libdnf5::repo::RepoSack` manages all repositories. Creates repos from:
- System configuration files (`.repo` files in `reposdir`)
- `--repofrompath` arguments
- Command-line packages

### Repo

`libdnf5::repo::Repo` represents a single repository. Key operations:
- Metadata download (via librepo)
- Metadata OpenPGP check
- Cache management (`RepoCache`)
- Package downloading (`PackageDownloader`, `FileDownloader`)

### RepoQuery

`libdnf5::repo::RepoQuery` filters repositories by enabled/disabled, name patterns, etc.

## Configuration System

### ConfigMain and ConfigRepo

`ConfigMain` holds all `[main]` configuration options. `ConfigRepo` holds per-repo options. Both use typed `Option*` classes:

| Option Class | Type |
|--------------|------|
| `OptionBool` | boolean |
| `OptionString` | string |
| `OptionStringList` | list |
| `OptionEnum<T>` | enum |
| `OptionNumber<T>` | number |
| `OptionPath` | path |
| `OptionSeconds` | time in seconds |

### ConfigParser

`libdnf5::ConfigParser` reads INI-style configuration files. Used for main config, repo configs, and plugin configs.

### Vars

`libdnf5::Vars` provides variable substitution (`$arch`, `$basearch`, `$releasever`, user-defined variables). See [repo-variables.md](repo-variables.md).

## Transaction History

`libdnf5::transaction::TransactionHistory` provides access to past transactions stored in a SQLite database (`transaction_history_dir`). Operations include listing, getting info, undoing, and storing transactions.

Transaction serialization/replay is handled by `transaction_sr.cpp` (store/replay).

## Language Bindings

Built using SWIG. Available for:

| Language | Status |
|----------|--------|
| Python 3 | Fully supported |
| Perl 5 | Best effort |
| Ruby | Best effort |
| Go | Not working, seeking contributors |

Python bindings are loaded via the `python_plugins_loader` libdnf5 plugin. Bindings mirror the C++ API structure.

## libdnf5-cli

The CLI helper library (`libdnf5-cli/`) provides:
- `argument_parser.cpp` — CLI argument parsing (used by dnf5 and dnf5daemon-client)
- `session.cpp` — CLI session management
- `output/` — Output formatting (transaction tables, adapters)
- `progressbar/` — Progress bar display
- `tty.cpp` — Terminal handling
- `utils/` — CLI utilities (units, user confirmation)

## Versioning

### Library Version

```cpp
struct LibraryVersion {
    std::uint16_t prime;   // stays as 5
    std::uint16_t major;   // incompatible API changes
    std::uint16_t minor;   // backward-compatible additions
    std::uint16_t micro;   // bug fixes
};
```

Current version: 5.4.4.0 (prime=5, major=4, minor=4, micro=0).

### ABI Stability

The `prime` version stays as 5 across all DNF5 releases. ABI changes are signaled by `major` version increments. The C++ API is documented via Doxygen from the public headers in `include/libdnf5/`.
