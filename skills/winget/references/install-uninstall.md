# Install, Upgrade, Uninstall, and Repair

The package lifecycle workflows in winget — how packages are installed, upgraded, uninstalled, and repaired, including installer selection, scopes, modes, dependencies, and portable packages.

> **Source of truth**: <https://learn.microsoft.com/en-us/windows/package-manager/winget/install>, `winget install --help`. For **manifest fields** that control installer behavior, see [manifest.md](manifest.md). For **settings** that affect installer selection, see [settings.md](settings.md). For **command structure**, see [cli.md](cli.md).

## `winget install` (alias: `add`)

Installs the selected package. The package can be specified by name, ID, moniker, or a free-text query. If multiple packages match, winget presents an interactive selection prompt (unless `--exact` or `--disable-interactivity` is used).

### Package Query

winget resolves packages using a combination of filters:

```
winget install <query>                    # free-text search
winget install --id Microsoft.VisualStudioCode
winget install --name "Visual Studio Code"
winget install --moniker vscode
winget install --exact --id Microsoft.VisualStudioCode
winget install --manifest .\package.yaml   # local manifest
```

When using `--exact` (`-e`), winget matches the filter exactly rather than as a substring. Without `--exact`, `--id vscode` would match `Microsoft.VisualStudioCode` by partial ID match.

### Install Flags

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--accept-package-agreements` | | bool | Accept all license agreements for packages |
| `--accept-source-agreements` | | bool | Accept all source agreements during source operations |
| `--allow-reboot` | | bool | Allows a reboot if applicable |
| `--architecture` | `-a` | string | Select the architecture |
| `--authentication-account` | | string | Specify the account for authentication |
| `--authentication-mode` | | string | Auth window preference (`silent`, `silentPreferred`, `interactive`) |
| `--custom` | | string | Additional arguments passed to the installer |
| `--dependency-source` | | string | Find package dependencies using the specified source |
| `--exact` | `-e` | bool | Find package using exact match |
| `--force` | | bool | Direct run, continue with non-security issues |
| `--header` | | string | Optional REST source HTTP header |
| `--id` | | string | Filter results by id |
| `--ignore-local-archive-malware-scan` | | bool | Ignore malware scan for local archive manifest |
| `--ignore-security-hash` | | bool | Ignore installer hash check failure |
| `--installer-type` | | string | Select the installer type |
| `--interactive` | `-i` | bool | Request interactive installation; user input may be needed |
| `--locale` | | string | Locale to use (BCP47 format) |
| `--location` | `-l` | string | Location to install to (if supported) |
| `--log` | `-o` | string | Log location (if supported) |
| `--manifest` | `-m` | string | The path to the manifest of the package |
| `--moniker` | | string | Filter results by moniker |
| `--name` | | string | Filter results by name |
| `--no-upgrade` | | bool | Skip upgrade if an installed version already exists |
| `--override` | | string | Override arguments to be passed on to the installer |
| `--query` | `-q` | string | The query used to search for a package |
| `--rename` | `-r` | string | The value to rename the executable file (portable) |
| `--scope` | | string | Select install scope (`user`, `machine`) |
| `--silent` | `-h` | bool | Request silent installation |
| `--skip-dependencies` | | bool | Skips processing package dependencies and Windows features |
| `--source` | `-s` | string | Find package using the specified source |
| `--uninstall-previous` | | bool | Uninstall the previous version of the package during upgrade |
| `--version` | `-v` | string | Use the specified version; default is the latest version |

### Installer Selection

When a manifest contains multiple installers, winget selects the best match based on:

1. **Architecture** — matches the current system (x64, arm64, etc.) or the `--architecture` flag
2. **Scope** — matches `--scope` flag, or `installBehavior.preferences.scope` setting (default: `user`)
3. **InstallerType** — matches `--installer-type` flag, or `installBehavior.preferences.installerTypes` priority
4. **Locale** — matches `--locale` flag, or `installBehavior.preferences.locale` setting

See [settings.md](settings.md) for the preferences and requirements configuration.

### Install Modes

| Mode | Description | Flag |
|------|-------------|------|
| `silent` | No UI, no user interaction | `--silent`, `-h` |
| `silentWithProgress` | Progress bar only, no user interaction | (default for most installers) |
| `interactive` | Full installer UI | `--interactive`, `-i` |

The mode must be listed in the manifest's `InstallModes` field. If only `silent` is listed, `--interactive` will not work.

### `--override` vs `--custom`

| Flag | Behavior |
|------|----------|
| `--override` | **Replaces** all installer switches from the manifest with the provided arguments |
| `--custom` | **Adds** arguments to the manifest's installer switches |

```
winget install --id Microsoft.VisualStudioCode --override "/verysilent /mergetasks=runcode"
winget install --id Git.Git --custom "--gitconfig core.autocrlf=true"
```

### `--force`

Runs the installer directly and continues with non-security-related issues. This skips some validation and is useful for troubleshooting or when an installer reports non-fatal warnings.

### `--no-upgrade`

If a version of the package is already installed, winget will normally upgrade to the new version. `--no-upgrade` skips installation if an installed version already exists, treating it as a no-op rather than an upgrade.

### `--ignore-security-hash`

Bypasses the SHA256 hash verification of the downloaded installer. This is a security risk — only use with trusted local manifests.

### Dependencies

If a manifest declares `Dependencies`, winget installs them before the main package. Use `--skip-dependencies` to skip this. Dependency types:

| Type | Description |
|------|-------------|
| `WindowsFeatures` | Windows features that must be enabled (e.g., `Microsoft-Hyper-V-All`) |
| `WindowsLibraries` | Windows libraries (not implemented) |
| `PackageDependencies` | Other winget packages that must be installed first |
| `ExternalDependencies` | External dependencies (not implemented, informational only) |

Use `--dependency-source` to specify which source to search for package dependencies.

## `winget upgrade` (alias: `update`)

Shows and performs available upgrades.

### Upgrade All

```
winget upgrade --all         # upgrade all installed packages with available updates
winget upgrade --recurse     # same as --all (alias flag)
```

### Upgrade Single Package

```
winget upgrade --id Microsoft.VisualStudioCode
winget upgrade Microsoft.VisualStudioCode --version 1.86.0
```

### Upgrade Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--all` | | Upgrade all installed packages to latest if available |
| `--recurse` | `-r` | Upgrade all installed packages to latest (alias for `--all`) |
| `--accept-package-agreements` | | Accept all license agreements |
| `--accept-source-agreements` | | Accept all source agreements |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference |
| `--force` | | Direct run, continue with non-security issues |
| `--header` | | REST source HTTP header |
| `--id` | | Filter results by id |
| `--ignore-local-archive-malware-scan` | | Ignore malware scan for local archive |
| `--include-pinned` | | Upgrade packages even if they have a non-blocking pin |
| `--pinned` | | Upgrade packages even if pinned (alias for `--include-pinned`) |
| `--include-unknown` | | Upgrade packages even if current version cannot be determined |
| `--unknown` | `-u` | Upgrade even if version unknown (alias for `--include-unknown`) |
| `--manifest` | `-m` | The path to the manifest of the package |
| `--name` | | Filter results by name |
| `--query` | `-q` | The query used to search |
| `--uninstall-previous` | | Uninstall previous version during upgrade |

### Upgrade Behavior

The manifest's `UpgradeBehavior` field controls how upgrades are performed:

| Value | Behavior |
|-------|----------|
| `install` | Install new version alongside old (no automatic uninstall) |
| `uninstallPrevious` | Uninstall old version before installing new (default for most) |
| `deny` | Do not allow upgrade; user must uninstall first, then install |

### `--include-unknown` and `--include-pinned`

By default, `winget upgrade --all` skips:
- Packages where the installed version cannot be determined (unknown) — use `--include-unknown` to include them
- Packages with a non-blocking pin — use `--include-pinned` to include them

See [package-management.md](package-management.md) for details on pinning.

## `winget uninstall`

Uninstalls the given package. The package is matched against the `Installed` predefined source (ARP + MSIX).

### Uninstall Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-source-agreements` | | Accept all source agreements |
| `--all` | | Uninstall all versions |
| `--all-versions` | | Uninstall all versions (alias for `--all`) |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference |
| `--exact` | `-e` | Find package using exact match |
| `--force` | | Direct run, continue with non-security issues |
| `--header` | | REST source HTTP header |
| `--interactive` | `-i` | Request interactive uninstall; user input may be needed |
| `--log` | `-o` | Log location (if supported) |
| `--moniker` | | Filter results by moniker |
| `--preserve` | | Retain all files/dirs created by the package (portable) |
| `--product-code` | | Filter using the product code |
| `--purge` | | Delete all files/dirs in the package directory (portable) |
| `--scope` | | Select installed package scope filter (`user`, `machine`) |
| `--silent` | `-h` | Request silent uninstall |
| `--source` | `-s` | Find package using the specified source |
| `--version` | `-v` | The version to act upon |

### Portable Package Cleanup

For `portable` installer type packages, the `--purge` and `--preserve` flags control what happens to files:

| Flag | Behavior |
|------|----------|
| `--purge` | Deletes all files and directories in the package directory |
| `--preserve` | Retains all files and directories created by the package |
| (default) | Follows `uninstallBehavior.purgePortablePackage` setting (default: `false`, preserves) |

## `winget repair`

Repairs the selected package. Uses the installer's repair functionality if available (controlled by the manifest's `RepairBehavior` field: `modify`, `uninstaller`, or `installer`).

### Repair Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-package-agreements` | | Accept all license agreements |
| `--accept-source-agreements` | | Accept all source agreements |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference (`silent`, `silentPreferred`, `interactive`) |
| `--exact` | `-e` | Find package using exact match |
| `--force` | | Direct run, continue with non-security issues |
| `--header` | | REST source HTTP header |
| `--ignore-local-archive-malware-scan` | | Ignore malware scan for local archive |
| `--ignore-security-hash` | | Ignore installer hash check failure |
| `--locale` | | Locale to use (BCP47 format) |
| `--log` | `-o` | Log location (if supported) |

## Portable Packages

`portable` is a special installer type for standalone binaries. Key differences:

- Installed to a portable package root directory (configurable via `installBehavior.portablePackageUserRoot` and `portablePackageMachineRoot` settings)
- No traditional installer — winget extracts the binary and creates a command alias
- `--rename` (`-r`) renames the executable file
- `--purge` / `--preserve` control file cleanup on uninstall
- Supports `ArchiveBinariesDependOnPath` manifest field for PATH-dependent behavior

## Authentication

REST sources may require authentication. The authentication flags control how winget handles this:

| Flag | Description |
|------|-------------|
| `--authentication-account` | Specify the account to use for authentication |
| `--authentication-mode` | Authentication window preference |

### Authentication Modes

| Mode | Description |
|------|-------------|
| `silent` | Attempt silent authentication; fail if silent auth is not possible |
| `silentPreferred` | Attempt silent authentication; fall back to interactive if needed (default) |
| `interactive` | Always use interactive authentication window |

## Related

- For the full manifest schema (installer types, scopes, switches, dependencies), see [manifest.md](manifest.md).
- For settings that control install behavior (preferences, requirements, portable roots), see [settings.md](settings.md).
- For search and list commands to find packages before installing, see [package-management.md](package-management.md).
- For command structure and global flags, see [cli.md](cli.md).
