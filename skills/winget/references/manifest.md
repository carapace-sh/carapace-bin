# Package Manifest YAML Schema

The complete specification for winget package manifests — the multi-file YAML format, all fields for version/installer/locale manifests, enumerated values, and validation.

> **Source of truth**: [winget-pkgs schema docs](https://github.com/microsoft/winget-pkgs/tree/master/doc/manifest), [winget-cli schemas](https://github.com/microsoft/winget-cli/tree/master/schemas/JSON/manifests). For **hash and validate commands**, see [package-management.md](package-management.md). For **installer types and scopes in the install workflow**, see [install-uninstall.md](install-uninstall.md).

## Multi-File Manifest Architecture

Every package version in the winget-pkgs community repository uses a **multi-file manifest** with three required YAML files plus optional locale files:

| File | `ManifestType` | Purpose |
|------|---------------|---------|
| `<PackageIdentifier>.yaml` | `singleton` | Single-file manifest (combines version, installer, and locale metadata) |
| `<PackageIdentifier>.yaml` | `version` | Multi-file: core version metadata, links other files |
| `<PackageIdentifier>.installer.yaml` | `installer` | Installer URLs, SHA256, architecture, type |
| `<PackageIdentifier>.locale.en-US.yaml` | `defaultLocale` | Default locale: name, publisher, description, license |
| `<PackageIdentifier>.locale.<tag>.yaml` | `locale` (optional) | Additional translations |

### Folder Structure

```
manifests/<first-letter>/<Publisher>/<Package>/<Version>/
  <PackageIdentifier>.yaml
  <PackageIdentifier>.installer.yaml
  <PackageIdentifier>.locale.en-US.yaml
  <PackageIdentifier>.locale.fr-FR.yaml   (optional)
```

Example: `manifests/m/Microsoft/VisualStudioCode/1.85.0/`

### Singleton Format

A manifest can also be a **single-file** (singleton) YAML containing all data in one file. This is useful for local manifests passed via `--manifest`. All fields from the version, installer, and defaultLocale manifests are combined. Use `ManifestType: singleton` for this format.

## ManifestVersion

The `ManifestVersion` field specifies the schema version. It only increments when the schema changes:

| ManifestVersion | Status |
|-----------------|--------|
| 1.0.0 | Deprecated |
| 1.1.0 | Deprecated |
| 1.2.0 | Deprecated |
| 1.4.0 | Deprecated |
| 1.5.0 | Deprecated |
| 1.6.0 | Deprecated |
| 1.7.0 | Deprecated |
| 1.9.0 | Deprecated |
| **1.10.0** | Accepted for community submissions |
| **1.12.0** | **Recommended** for all new PRs |
| 1.28.0 | Latest in-development schema version |

There is no schema version 1.3.0, 1.8.0, or 1.11.0 — those WinGet client releases did not introduce manifest schema changes.

## Version Manifest (`ManifestType: version`)

```yaml
PackageIdentifier:       # Required. Unique ID in "Publisher.Package" format (case-sensitive, max 128 chars)
PackageVersion:          # Required. The package version string (max 128 chars)
DefaultLocale:           # Required. BCP-47 locale for default metadata (e.g., "en-US", max 20 chars)
ManifestType: version    # Required. Must be "version"
ManifestVersion: 1.12.0  # Required. The manifest schema version
```

## Default Locale Manifest (`ManifestType: defaultLocale`)

```yaml
PackageIdentifier:            # Required - Unique ID
PackageVersion:               # Required - Package version
PackageLocale:                # Required - BCP-47 locale (max 20 chars)
Publisher:                    # Required - Publisher name
PublisherUrl:                 # Optional - Publisher homepage URL
PublisherSupportUrl:          # Optional - Publisher support page URL
PrivacyUrl:                   # Optional - Privacy page URL
Author:                       # Optional - Package author
PackageName:                  # Required - Package display name
PackageUrl:                   # Optional - Package homepage URL
License:                      # Required - License name or SPDX identifier
LicenseUrl:                   # Optional - License page URL
Copyright:                    # Optional - Copyright text
CopyrightUrl:                 # Optional - Copyright page URL
ShortDescription:             # Required - Short description (3-256 chars)
Description:                  # Optional - Full/long description
Moniker:                      # Optional - Most common search term
Tags:                         # Optional - List of search terms (max 16, lowercase with hyphens)
Agreements:                   # Optional - List (verified developers only)
  - AgreementLabel:           #   Optional - Agreement label
    Agreement:                #   Restricted - Full agreement text
    AgreementUrl:             #   Optional - Agreement URL
Documentations:               # Optional - List of documentation
  - DocumentLabel:            #   Optional - Documentation label
    DocumentUrl:              #   Optional - Documentation URL
ReleaseNotes:                 # Optional - Release notes text
ReleaseNotesUrl:              # Optional - Release notes URL
PurchaseUrl:                  # Optional - Purchase URL
InstallationNotes:            # Optional - Notes shown after installation
ManifestType: defaultLocale   # Required - Must be "defaultLocale"
ManifestVersion: 1.12.0       # Required - Schema version
```

## Locale Manifest (`ManifestType: locale`)

Same fields as defaultLocale, but all fields except `PackageIdentifier`, `PackageVersion`, `PackageLocale`, `ManifestType`, and `ManifestVersion` are **optional**. Used for additional translations.

## Installer Manifest (`ManifestType: installer`)

```yaml
PackageIdentifier:                    # Required - Unique ID
PackageVersion:                       # Required - Package version
Channel:                              # Optional - Distribution channel (e.g., "stable", "beta")
Installers:                           # Required - List of installer objects (1-1024)
  - Architecture:                     #   Required - Target architecture (see below)
    InstallerLocale:                  #   Optional - Locale of the installer
    Platform:                         #   Optional - ["Windows.Desktop", "Windows.Universal"]
    MinimumOSVersion:                 #   Optional - Min Windows version (e.g., "10.0.18362.0")
    InstallerType:                    #   Required - Installer type (see below)
    InstallerUrl:                     #   Required - Download URL (HTTPS, max 2048 chars)
    InstallerSha256:                  #   Required - SHA-256 hex hash (64 hex chars)
    SignatureSha256:                  #   Optional - SHA-256 of MSIX signature file
    NestedInstallerType:              #   Required when InstallerType is "zip"
    NestedInstallerFiles:             #   Required when InstallerType is "zip" (max 1024 items)
      - RelativeFilePath:             #     Required - Relative path inside archive
        PortableCommandAlias:         #     Optional - Command alias (portable only, max 40 chars)
    Scope:                            #   Optional - "user" or "machine"
    InstallModes:                     #   Optional - ["interactive", "silent", "silentWithProgress"]
    InstallerSwitches:                #   Optional - Installer command-line switches
      Silent:                         #     Optional - Silent install switch (max 512 chars)
      SilentWithProgress:             #     Optional - Silent with progress switch
      Interactive:                    #     Optional - Interactive install switch
      InstallLocation:                #     Optional - Install path switch (uses <INSTALLPATH> token)
      Log:                            #     Optional - Log path switch (uses <LOGPATH> token)
      Upgrade:                        #     Optional - Upgrade switch
      Custom:                         #     Optional - Custom switches (max 2048 chars)
      Repair:                         #     Optional - Repair switch
    UpgradeBehavior:                  #   Optional - "install", "uninstallPrevious", or "deny"
    Commands:                         #   Optional - List of commands/aliases (max 16, max 40 chars each)
    Protocols:                        #   Optional - List of URI schemes (max 64)
    FileExtensions:                   #   Optional - List of file extensions (max 512)
    Dependencies:                     #   Optional - Dependency requirements
      WindowsFeatures:                #     Optional - List of Windows features (max 16)
      WindowsLibraries:               #     Optional - List of Windows libraries (not implemented)
      PackageDependencies:            #     Optional - List of package deps (max 16)
        - PackageIdentifier:          #       Required - Dependent package ID
          MinimumVersion:             #       Optional - Minimum version
      ExternalDependencies:           #     Optional - List of external deps (not implemented)
    PackageFamilyName:                #   Optional - MSIX package family name (max 255 chars)
    Capabilities:                     #   Optional - MSIX capabilities (max 1000) (restricted)
    RestrictedCapabilities:           #   Optional - MSIX restricted capabilities (restricted)
    InstallerAbortsTerminal:          #   Optional - Boolean, installer aborts terminal
    InstallLocationRequired:          #   Optional - Boolean, requires install location
    RequireExplicitUpgrade:           #   Optional - Boolean, exclude from winget upgrade --all
    ElevationRequirement:             #   Optional - "elevationRequired", "elevationProhibited", "elevatesSelf"
    UnsupportedOSArchitectures:       #   Optional - ["x86", "x64", "arm", "arm64"]
    Markets:                          #   Optional - Supported markets (restricted)
    ExcludedMarkets:                  #   Optional - Excluded markets (restricted)
    InstallerSuccessCodes:            #   Optional - List of non-zero success codes (max 16)
    ExpectedReturnCodes:              #   Optional - List of expected return codes (max 128)
      - InstallerReturnCode:          #     Required - Non-zero integer
        ReturnResponse:               #     Required - Response string (enum)
        ReturnResponseUrl:            #     Optional - Response URL
    ProductCode:                      #   Optional - MSI product code (max 255 chars)
    AppsAndFeaturesEntries:           #   Optional - ARP entries (max 128)
      - DisplayName:                  #     Optional - ARP display name (max 256 chars)
        DisplayVersion:               #     Optional - ARP display version (max 128 chars)
        Publisher:                    #     Optional - ARP publisher (max 256 chars)
        ProductCode:                  #     Optional - Product code
        UpgradeCode:                  #     Optional - Upgrade code
        InstallerType:                #     Optional - Installer type for this ARP entry
    UnsupportedArguments:             #   Optional - ["log", "location"]
    DisplayInstallWarnings:           #   Optional - Boolean (not implemented)
    ReleaseDate:                      #   Optional - ISO 8601 date (YYYY-MM-DD)
    InstallationMetadata:             #   Optional - Deeper installation detection
      - DefaultInstallLocation:       #     Optional - Default install path (max 2048 chars)
        Files:                        #     Optional - List of installed files (max 2048)
          - RelativeFilePath:         #       Required - Relative path (max 2048 chars)
            FileSha256:               #       Optional - SHA-256 of file
            FileType:                 #       Optional - "launch", "uninstall", or "other"
            InvocationParameter:      #       Optional - Parameter for invocable files
            DisplayName:              #       Optional - Display name for invocable files
    DownloadCommandProhibited:        #   Optional - Boolean
    RepairBehavior:                   #   Optional - "modify", "uninstaller", or "installer"
    ArchiveBinariesDependOnPath:      #   Optional - Boolean (zip/portable packages)
    Authentication:                  #   Optional - Entra ID auth (restricted)
      AuthenticationType:             #     Required - "none", "microsoftEntraId", "microsoftEntraIdForAzureBlobStorage"
      MicrosoftEntraIdAuthenticationInfo: # Optional
        Resource:                     #       Optional - Resource string (max 512 chars)
        Scope:                        #       Optional - Scope string (max 512 chars)
    DesiredStateConfiguration:        #   Optional - DSC resource references
      PowerShell:                     #     Optional - PowerShell DSC modules (max 16)
        - RepositoryUrl:              #       Required - Module repository URL
          ModuleName:                 #       Required - Module name
          Resources:                  #       Required - List of resources
            - Name:                   #         Required - Resource name
      DSCv3:                          #     Optional - DSC v3 resources
        Resources:                    #       Required - List (max 128)
          - Type:                     #         Required - Resource type string
ManifestType: installer               # Required - Must be "installer"
ManifestVersion: 1.12.0               # Required - Schema version
```

### Root-Level Defaults

Fields specified at the root level (outside `Installers[]`) act as **defaults** inherited by all installer entries if not specified per-installer. This includes `InstallerType`, `Scope`, `InstallModes`, `InstallerSwitches`, `UpgradeBehavior`, `Commands`, `Protocols`, `FileExtensions`, `Dependencies`, etc.

## Enumerated Values

### InstallerType

| Value | Description |
|-------|-------------|
| `msix` | MSIX installer |
| `msi` | Windows Installer (MSI) |
| `appx` | APPX installer |
| `exe` | Executable installer |
| `zip` | Archive containing nested installer |
| `inno` | Inno Setup |
| `nullsoft` | Nullsoft (NSIS) |
| `wix` | WiX Toolset |
| `burn` | WiX Burn bundle |
| `pwa` | Progressive Web Application |
| `portable` | Standalone portable binary |
| `font` | Font file (.ttf, .ttc, .otf, .otc, .fnt) |

### NestedInstallerType (when InstallerType is `zip`)

All values from InstallerType **except** `zip`, `pwa`, and `font`: `msix`, `msi`, `appx`, `exe`, `inno`, `nullsoft`, `wix`, `burn`, `portable`, `font`.

### Architecture

`x86`, `x64`, `arm`, `arm64`, `neutral`

### Scope

`user`, `machine`

### InstallModes

`interactive`, `silent`, `silentWithProgress`

### UpgradeBehavior

| Value | Description |
|-------|-------------|
| `install` | Install alongside previous version (no uninstall) |
| `uninstallPrevious` | Uninstall previous version before installing new |
| `deny` | Do not allow upgrade; user must uninstall first |

### ElevationRequirement

| Value | Description |
|-------|-------------|
| `elevationRequired` | Installer requires admin elevation |
| `elevationProhibited` | Installer must run without elevation |
| `elevatesSelf` | Installer handles its own elevation (UAC prompt) |

### InstallerSwitches Tokens

| Token | Replaced with |
|-------|---------------|
| `<INSTALLPATH>` | The install location (from `--location` or default) |
| `<LOGPATH>` | The log file path (from `--log`) |

### ReturnResponse (in ExpectedReturnCodes)

`packageInUse`, `packageInUseByApplication`, `installInProgress`, `fileInUse`, `missingDependency`, `diskFull`, `insufficientMemory`, `invalidParameter`, `noNetwork`, `contactSupport`, `rebootRequiredToFinish`, `rebootRequiredForInstall`, `rebootInitiated`, `cancelledByUser`, `alreadyInstalled`, `downgrade`, `blockedByPolicy`, `systemNotSupported`, `custom`

### RepairBehavior

`modify`, `uninstaller`, `installer`

### AuthenticationType

`none`, `microsoftEntraId`, `microsoftEntraIdForAzureBlobStorage`

### Platform

`Windows.Desktop`, `Windows.Universal`

### FileType (in InstallationMetadata)

`launch`, `uninstall`, `other`

### UnsupportedArguments

`log`, `location`

## Validation and Hashing

### `winget validate`

Validates a manifest file against the schema:

```
winget validate --manifest <path-to-manifest.yaml>
```

Checks YAML syntax, schema compliance, required fields, and field constraints (max lengths, valid enum values, etc.).

### `winget hash`

Generates the `InstallerSha256` hash for an installer file:

```
winget hash --file <path-to-installer>
winget hash --file <path-to-installer> --msix
```

For MSIX files, the `--msix` (`-m`) flag also provides the `SignatureSha256` if the package is signed.

## Example Installer Manifest

```yaml
PackageIdentifier: Microsoft.VisualStudioCode
PackageVersion: 1.85.0
DefaultLocale: en-US
ManifestType: singleton
ManifestVersion: 1.12.0
Installers:
  - Architecture: x64
    InstallerType: exe
    InstallerUrl: https://update.code.visualstudio.com/1.85.0/win32-x64/stable
    InstallerSha256: ABC123...
    Scope: user
    InstallModes:
      - interactive
      - silent
      - silentWithProgress
    InstallerSwitches:
      Silent: /verysilent
      SilentWithProgress: /silent
      Interactive: /interactive
      InstallLocation: /INSTALLPATH="<INSTALLPATH>"
      Log: /LOG="<LOGPATH>"
    UpgradeBehavior: uninstallPrevious
    Commands:
      - code
    Protocols:
      - vscode
    FileExtensions:
      - .vscode
    ProductCode: '{771FD6BD-FA35-4B31-BF1C-C8AC059D2B0E}'
    AppsAndFeaturesEntries:
      - DisplayName: Visual Studio Code (User)
        DisplayVersion: 1.85.0
        Publisher: Microsoft Corporation
        ProductCode: '{771FD6BD-FA35-4B31-BF1C-C8AC059D2B0E}'
        InstallerType: exe
```

## Local Manifests

You can install from a local manifest file using `--manifest` (`-m`):

```
winget install --manifest .\path\to\manifest.yaml
```

Local manifests bypass source searching and use the manifest directly. They must still pass hash verification unless `--ignore-security-hash` is specified.

## Related

- For the install/upgrade/uninstall workflow and how installer selection works, see [install-uninstall.md](install-uninstall.md).
- For core concepts like PackageIdentifier and package correlation, see [concepts.md](concepts.md).
- For validation and hashing commands, see [package-management.md](package-management.md).
- For the community repository submission process, see the [winget-pkgs repository](https://github.com/microsoft/winget-pkgs).
