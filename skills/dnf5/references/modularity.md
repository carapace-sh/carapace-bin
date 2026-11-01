# DNF5 Modularity

Modularity is an alternative way of building, organizing, and delivering packages in RPM-based distributions. DNF5 supports modulemd metadata and modular package filtering.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/misc/modularity.7.html>. For package filtering, see [filtering.md](filtering.md).

## Definitions

| Term | Description |
|------|-------------|
| **modulemd** | Metadata documents in repository `modules` metadata. Hold module info: `Name`, `Stream`, list of packages. |
| **(non-modular) package** | Package that doesn't belong to a module. |
| **modular package** | Package belonging to a module. Listed in modulemd under `artifacts`. Has `%{modularitylabel}` RPM header set. |
| **(module) stream** | A collection of packages, a virtual repository. Identified by `Name:Stream` (e.g., `postgresql:9.6`). |

## Stream States

| State | Description |
|-------|-------------|
| **active** | RPM packages from this stream are included in the available package set. |
| **inactive** | Packages from this stream are filtered out. |

A stream is **active** if:
- Marked as **default**, OR
- Explicitly **enabled** by a user action, OR
- Satisfies dependencies of default or enabled streams.

Only **one** stream of a particular module can be active at a given point in time.

## Package Filtering

Without modules, packages with the highest version are used by default. Module streams can distribute packages with lower versions than available in OS repositories.

### How It Works

Non-modular packages are filtered out when their **name or provide** matches against a modular package name from any enabled, default, or dependent stream.

**Note**: Modular source packages will not cause non-modular binary packages to be filtered out.

## Demodularized RPMs

Contains names of RPMs excluded from package filtering for a particular module stream. When defined in the latest active module, non-modular RPMs with the same name or provide that were previously filtered out will **reappear** in the available package set.

## Hotfix Repositories

In special cases, users want to cherry-pick individual packages provided outside module streams. To make packages from a repository available regardless of modularity:

```ini
# In .repo file
module_hotfixes=true
```

Or via CLI:

```bash
dnf5 --setopt=<repo_id>.module_hotfixes=true install <package>
```

**Important**: Hotfix packages do **not override** module packages. They only become part of the available package set. It is the package `Epoch`, `Version`, and `Release` that determines if the package is the latest.

## Fail-Safe Mechanisms

### When Repositories with Module Metadata Are Unavailable

Package filtering must keep working even when module metadata is unavailable. This happens when:
- User disables a repository via `--disable-repo` or uses `--repo`
- User removes a `.repo` file from disk
- Repository is not available and has `skip_if_unavailable=true`

DNF5 keeps copies of the **latest modulemd for every active stream** and uses them if no modulemd is available for the stream.

- Copies are made any time a transaction is resolved and started (includes RPM transactions and `dnf5 module <enable|disable|reset>` operations).
- When fail-safe data is used, DNF5 shows such modules as part of `@modulefailsafe` repository.

### Orphaned Modular Packages

All packages built as part of a module have `%{modularitylabel}` RPM header set. If such a package becomes part of an RPM transaction and **cannot be associated with any available modulemd**, DNF5 prevents it from getting on the system (package is available but cannot be installed, upgraded, etc.).

Packages from **Hotfix repositories** or **Commandline repository** are not affected by fail-safe mechanisms.

## Configuration

| Option | Description |
|--------|-------------|
| `module_platform_id` | Override `PLATFORM_ID` from `/etc/os-release`. Format: `$name:$stream`. |
| `module_stream_switch` | Allow switching enabled streams of a module. Default: `False`. |

## Note on Fedora 39+

Modularity support was discontinued in dnf5 starting with Fedora 39. The `module` command is conditionally compiled with `WITH_MODULEMD`.
