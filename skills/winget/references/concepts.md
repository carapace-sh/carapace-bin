# Core Concepts and Architecture

The foundational mental model for understanding winget — the package manager architecture, package identification, source system, package correlation, and the binary stack that powers the CLI.

> **Source of truth**: <https://github.com/microsoft/winget-cli>, <https://learn.microsoft.com/en-us/windows/package-manager/winget/>. For **source types and management**, see [sources.md](sources.md). For **package manifests**, see [manifest.md](manifest.md). For **command reference**, see [cli.md](cli.md).

## What winget Is

`winget` (Windows Package Manager) is a command-line tool for discovering, installing, upgrading, uninstalling, and configuring applications on Windows 10 (1809+) and Windows 11. It is part of the **App Installer** package (`Microsoft.DesktopAppInstaller`) and ships pre-installed on modern Windows installations.

```
winget.exe (thin console app)
  │  loads WindowsPackageManager.dll
  ▼
WindowsPackageManagerCLIMain()
  │  initializes logging, telemetry, COM apartment
  ▼
RootCommand → Subcommand → Workflow Chain
  │
  ├── Source resolution (CompositeSource)
  ├── Search (SQLite index or REST API)
  ├── Package correlation (match installed → available)
  └── Installer execution (MSI, MSIX, EXE, etc.)
```

## Binary Inventory

winget is not a single binary — it is a layered system of native C++ libraries with COM/WinRT APIs:

| Binary | Role |
|--------|------|
| `winget.exe` | User-facing CLI entry point (thin console app) |
| `WindowsPackageManager.dll` | CLI host; COM activation DLL |
| `WindowsPackageManagerServer.exe` | Out-of-process COM server |
| `Microsoft.Management.Deployment.dll` | WinRT COM API for package operations (in-proc and out-of-proc) |
| `Microsoft.Management.Configuration.dll` | WinRT COM server for DSC configuration API |
| `WinGetUtil.dll` | Standalone C-export utility (manifest indexing, hashing, validation) |
| `winrtact.dll` | Registration-free WinRT activation for unpackaged callers |

### C++ Library Stack (layered)

| Library | Key Responsibilities |
|---------|----------------------|
| `AppInstallerSharedLib` | Low-level types, error handling, string utilities |
| `AppInstallerCommonCore` | Telemetry, download engine, MSIX introspection, manifest parsing, user settings |
| `AppInstallerRepositoryCore` | SQLite index, source abstraction, REST client, composite source, search types |
| `AppInstallerCLICore` | Command hierarchy, workflow engine, execution context, installer flows |

## Packaged vs. Unpackaged Execution Context

winget runs in two different contexts depending on whether it is launched from the App Installer MSIX or as a standalone binary:

| Aspect | Packaged (App Installer) | Unpackaged (source build) |
|--------|--------------------------|---------------------------|
| Path resolution | Package local state | `%LOCALAPPDATA%\Microsoft\WinGet` |
| WinRT activation | System-registered MSIX classes | Requires `winrtact.dll` |
| Settings file | `%LOCALAPPDATA%\Packages\Microsoft.DesktopAppInstaller_8wekyb3d8bbwe\LocalState\settings.json` | `%LOCALAPPDATA%\Microsoft\WinGet\Settings\settings.json` |
| Execution alias | `WinGetDev.exe` | `winget.exe` |
| Source index delivery | Windows App Extensions | Manual download + extract |

## Package Identification

Every package in winget is identified by multiple identifiers, each serving a different purpose:

### PackageIdentifier

The primary identifier in the winget ecosystem. A unique, reverse-domain-notation string (e.g., `Microsoft.VisualStudioCode`, `Git.Git`, `Python.Python.3.12`). This is specified in the manifest under the `PackageIdentifier` field.

- Case-sensitive, max 128 characters
- Format: `Publisher.Package` (or `Publisher.Package.Subpackage`)
- Used in `winget install <PackageIdentifier>`, `winget show <PackageIdentifier>`, etc.
- Unique within a source, not globally (though the community repository aims for global uniqueness)

### Moniker

A short, common search term (e.g., `vscode`, `git`, `python`). Set in the manifest's `Moniker` field. Used as a convenience alias for searching — `winget install vscode` will match the moniker. Not guaranteed unique.

### ProductCode

An MSI/EXE identifier for installed packages (GUID format). Used by winget to match installed packages against available packages. Found in the `AppsAndFeaturesEntries` or root-level of an installer entry in the manifest.

### PackageFamilyName

An MSIX package family name (e.g., `Microsoft.VisualStudioCode_8wekyb3d8bbwe`). Used to match installed MSIX packages. Found in the manifest's `PackageFamilyName` field.

### UpgradeCode

An MSI upgrade code (GUID). Used to determine if an installed MSI package can be upgraded. Found in `AppsAndFeaturesEntries` in the manifest.

## The Source System

winget discovers packages through **sources** — sets of packages with metadata. At runtime, the **CompositeSource** aggregates all configured sources into a unified view for search and install operations.

### Source Types

| Type | Identifier | Description |
|------|------------|-------------|
| PreIndexed Package | `Microsoft.PreIndexed.Package` | SQLite index inside a signed MSIX file (default type) |
| REST API | `Microsoft.Rest` | HTTP-based source implementing the WinGet REST API spec |

### Predefined (Internal) Sources

Not persisted — represent system-level package information:

| Source | Description |
|--------|-------------|
| `Installed` | All ARP + MSIX packages for current user and machine |
| `InstalledUser` | Packages in user scope only |
| `InstalledMachine` | Packages in machine scope only |
| `ARP` | Add/Remove Programs registry entries only |
| `MSIX` | MSIX packages only |
| `Installing` | Writable source tracking in-progress installs |
| `InstalledForceCacheUpdate` | Like `Installed` but forces cache refresh |

### Default Well-Known Sources

| Name | Type | Description |
|------|------|-------------|
| `winget` | PreIndexed | WinGet Community Repository (winget-pkgs) |
| `msstore` | PreIndexed | Microsoft Store catalog |
| `winget-font` | PreIndexed | Font packages (explicit — requires `--source winget-font`) |

For full source management details, see [sources.md](sources.md).

## Package Correlation

When winget needs to determine if an installed package matches an available package (for `upgrade`, `list`, or `uninstall`), it uses **correlation** — matching installed packages to available manifests.

### Strong Match Fields

These fields provide high-confidence matches without heuristics:

| Field | Package Type | Description |
|-------|-------------|-------------|
| `PackageFamilyName` | MSIX | Unique to an MSIX package family |
| `ProductCode` | MSI/EXE | GUID identifying the installed product |
| `UpgradeCode` | MSI | GUID used by MSI for upgrade relationships |

### Version Mapping

When an installed package is matched to an available package, winget maps the ARP (Add/Remove Programs) version string to the manifest version using `ArpMinVersion` and `ArpMaxVersion` properties. This handles cases where the ARP version string differs from the manifest `PackageVersion` (e.g., `15.0.1.0` in ARP vs `15.0.1` in the manifest).

### Search Request Structure

The `SearchRequest` structure controls how searches work:

| Field | Purpose |
|-------|---------|
| `Query` | A `RequestMatch` for general search terms |
| `Inclusions` | Filters that expand the result set |
| `Filters` | Filters that restrict the result set |
| `Purpose` | `Default`, `CorrelationToInstalled`, or `CorrelationToAvailable` |

### Installed Package Discovery

The predefined `Installed` source aggregates packages from:

1. **ARP (Add/Remove Programs)** — Scans registry keys in `HKLM` and `HKCU` for both 32-bit and 64-bit scopes
2. **MSIX** — Uses `PackageManager` to enumerate provisioned and user-installed packages (filtering out system packages)
3. **Fonts** — Scans Windows font registry and filesystem

## How winget Downloads and Installs

### Installer Flow

```
winget install <package>
  → search sources for matching package
    → select best installer (architecture, scope, installer type)
      → download installer from InstallerUrl
        → verify SHA256 hash (InstallerSha256)
          → execute installer with appropriate switches
            → record in Installed source
```

### Installer Selection Priority

When multiple installers exist in a manifest, winget selects based on:

1. **Architecture** — matches current system (or preferences in [settings.md](settings.md))
2. **Scope** — matches `installBehavior.preferences.scope` setting (default: `user`)
3. **InstallerType** — matches `installBehavior.preferences.installerTypes` priority
4. **Locale** — matches `installBehavior.preferences.locale` setting

### Install Modes

| Mode | Description | Flag |
|------|-------------|------|
| `silent` | No UI, no user interaction | `--silent`, `-h` |
| `silentWithProgress` | Progress bar only, no user interaction | (default for most installers) |
| `interactive` | Full installer UI | `--interactive`, `-i` |

## Group Policy

Organizations can control winget behavior via Group Policy:

- **Source management** — Add/remove/restrict configured sources
- **Configuration** — `EnableWindowsPackageManagerConfiguration` and `EnableWindowsPackageManagerConfigurationExplanation` can disable WinGet Configuration
- **Other policies** — Various restrictions on package installation

Run `winget --info` to see any configured policies.

## Related

- For source management commands and REST API details, see [sources.md](sources.md).
- For manifest YAML structure and all fields, see [manifest.md](manifest.md).
- For install/upgrade/uninstall command flags and workflows, see [install-uninstall.md](install-uninstall.md).
- For settings that control installer selection and behavior, see [settings.md](settings.md).
