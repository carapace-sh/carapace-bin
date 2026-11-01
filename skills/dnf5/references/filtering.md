# DNF5 Package Filtering

Package filtering removes packages from the available package set, making them invisible to most DNF5 commands. Filtered packages cannot be included in any transaction.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/misc/filtering.7.html>. For modularity, see [modularity.md](modularity.md).

## Four Ways Packages Can Be Filtered

### 1. Global Excludes Filtering

Adjusted using `includepkgs` or `excludepkgs` configuration options in the DNF5 configuration file. Applied to all repositories.

Disable with:
```bash
dnf5 --setopt=disable_excludes=* install bash
```

### 2. Repository Excludes Filtering

Similar to global excludes but repository-specific. Only affects packages within the repository where the excludes are set. Configured via per-repo `excludepkgs` and `includepkgs` options.

### 3. User Excludes Filtering

API users can fine-tune excluded packages using `*_user_excludes()` and `*_user_includes()` methods of the `libdnf5::rpm::PackageSack` object. See [libdnf5.md](libdnf5.md).

### 4. Versionlock

Implemented using excludes filtering. Applied only during transactional operations. See `dnf5-versionlock(8)`.

## Modular Filtering

Only RPM packages from **active** module streams are included in the available package set.

- Packages from **inactive** module streams are filtered out.
- Non-modular packages with the same name or provides as a package from an active module stream are filtered out.

### Not Applied To

- Packages added from the command line
- Installed packages
- Packages from repositories with `module_hotfixes=true` in their `.repo` file

### Disabling

Disabling modular filtering is **not recommended** (could cause a broken system state).

To disable for a particular repository:
```ini
module_hotfixes=true
```

Or via CLI:
```bash
dnf5 --setopt=<repo_id>.module_hotfixes=true install <package>
```

See [modularity.md](modularity.md) for full details on the modular filtering system.

## Related

- [configuration.md](configuration.md) — `excludepkgs`, `includepkgs`, `disable_excludes`
- [modularity.md](modularity.md) — Module streams, active/inactive, fail-safe
- [libdnf5.md](libdnf5.md) — `PackageSack` API for user excludes
