# DNF5 Caching

How DNF5 caches metadata and packages, the cache directory structure, metadata types, and cache management.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/misc/caching.7.html>. For configuration options, see [configuration.md](configuration.md).

## Cache Directory Location

| User | Default Path | Config Option |
|------|-------------|----------------|
| Regular user | `~/.cache/libdnf5/` | `cachedir` |
| Superuser (root) | `/var/cache/libdnf5/` | `system_cachedir` (overwrites `cachedir`) |

## Cache Directory Structure

```
/var/cache/libdnf5/
├── fedora-*
│   ├── metalink.xml
│   ├── repodata/
│   │   ├── *-comps-Everything.x86_64.xml
│   │   ├── *-primary.xml.zck
│   │   └── repomd.xml
│   └── solv/
│       ├── fedora-group.solvx
│       └── fedora.solv
├── temporary_files.toml
└── updates-*
    ├── metalink.xml
    ├── packages/
    │   └── bash-5.2.21-1.fc38.x86_64.rpm
    ├── repodata/
    │   ├── *-primary.xml.zck
    │   ├── *-updateinfo.xml.zck
    │   ├── *-comps-Everything.x86_64.xml.gz
    │   └── repomd.xml
    └── solv/
        ├── updates-group.solvx
        ├── updates-updateinfo.solvx
        └── updates.solv
```

### Subdirectory Contents

| Subdirectory | Description |
|--------------|-------------|
| `repodata/` | Metadata files downloaded from the repository |
| `solv/` | Solver-generated cached files (libsolv format, for performance). Controlled by `build_cache` config option. |
| `packages/` | Downloaded RPM packages (if `keepcache=True`) |
| `metalink.xml` / `mirrorlist` | Information on remote locations of repository data |
| `temporary_files.toml` | Tracking file for temporarily stored packages |

## Metadata Types

### Mandatory Metadata (always loaded)

| Type | Description |
|------|-------------|
| `repomd.xml` | Info about specific metadata type files (checksums, file sizes, locations) |
| `primary` | Detailed package info (names, versions, dependencies, etc.) |
| `modules` | Module metadata (if compiled with `WITH_MODULEMD`) |

### Optional Metadata Types

Loaded via `optional_metadata_types` config option, automatically during runtime, or on explicit user request:

| Type | Description |
|------|-------------|
| `comps` | Package groups and environment descriptions |
| `filelists` | All files provided by packages (loaded when user passes a filepath argument) |
| `updateinfo` | Security-related updates and advisories |
| `presto` | Delta RPM information |
| `other` | Additional metadata (e.g., changelogs) |

If required metadata is missing, results can include: empty query, error output for no match, or transaction error.

## Package Caching

| Setting | Behavior |
|---------|----------|
| `keepcache=False` (default) | Downloaded packages are tracked in `temporary_files.toml` and removed after the next successful transaction (only if transaction contains inbound actions). |
| `keepcache=True` | Downloaded files are not tracked and persist until manually removed or via `dnf5 clean packages`. |
| `download` command | Packages downloaded via `dnf5 download` are always retained regardless of `keepcache`. |

## Sharing Root Cache Among Users

When running as non-root, DNF5 first checks the root's cache location. If metadata is present there, it is **copied** to the user cache location. This cloning is optimized when copy-on-write functionality is present (e.g., btrfs).

## Cacheonly

| Setting | Description |
|---------|-------------|
| `cacheonly=all` | Run entirely from system cache. No downloads. |
| `cacheonly=metadata` | Use only cached metadata. Packages can still be downloaded. Useful when optional repos are temporarily unavailable or cached metadata is expired but still functional. |
| `cacheonly=none` (default) | Normal operation. |
| `--cacheonly` / `-C` | Sets `cacheonly=all` (fully cache-driven operation). |

## Cache Management Commands

| Command/Option | Description |
|---------------|-------------|
| `dnf5 clean all` | Clean all cached data |
| `dnf5 clean packages` | Remove cached packages |
| `dnf5 clean metadata` | Remove cached metadata |
| `dnf5 makecache` | Generate metadata cache |
| `--refresh` | Force metadata update before executing a command |

## Metadata Expiration

The `metadata_expire` configuration option determines whether local repo data is due for re-sync. Default: 48 hours. Values: time in seconds, `-1` or `never` for no expiration.

Commands where up-to-date metadata is not critical (e.g., `group list`) use whatever version is locally available, ignoring expiration.
