# The Source System

How winget discovers packages through sources — source types, default sources, trust levels, source agreements, REST API sources, and the PreIndexed package delivery mechanism.

> **Source of truth**: <https://github.com/microsoft/winget-cli>, <https://learn.microsoft.com/en-us/windows/package-manager/winget/source>. For **core concepts**, see [concepts.md](concepts.md). For **source command reference**, see [cli.md](cli.md).

## What Sources Are

A **source** is a set of packages with associated metadata that winget can search and retrieve. Sources provide the ability to discover and retrieve metadata about packages so the client can act on it. At runtime, the **CompositeSource** aggregates all configured sources into a unified view, enabling searches across both installed and available packages.

## Source Types

### Configurable Source Types

These are the types users can add via `winget source add -t <type>`:

| Type | Identifier | Description |
|------|------------|-------------|
| **PreIndexed Package** | `Microsoft.PreIndexed.Package` | The default source type. SQLite-based index packaged inside a signed MSIX file. Used by the `winget` and `msstore` default sources. |
| **REST API** | `Microsoft.Rest` | A Microsoft REST API-based source. Organizations can host their own private REST source using the [winget-cli-restsource](https://github.com/microsoft/winget-cli-restsource) reference implementation. |

### Predefined (Internal) Sources

Not persisted — represent system-level package information. These sources are always available and cannot be removed:

| Source | Description |
|--------|-------------|
| `Installed` | All ARP + MSIX packages for current user and machine |
| `InstalledUser` | Packages installed in user scope only |
| `InstalledMachine` | Packages installed in machine scope only |
| `ARP` | Add/Remove Programs registry entries only |
| `MSIX` | MSIX packages only |
| `Installing` | Writable source tracking in-progress installs |
| `InstalledForceCacheUpdate` | Like `Installed` but forces cache refresh |

## Default Well-Known Sources

winget ships with default sources that are available out of the box:

| Name | Type | URL | Explicit? | Description |
|------|------|-----|-----------|-------------|
| **winget** | `Microsoft.PreIndexed.Package` | `https://cdn.winget.microsoft.com/cache` | No | WinGet Community Repository ([winget-pkgs](https://github.com/microsoft/winget-pkgs)) |
| **msstore** | `Microsoft.PreIndexed.Package` | `https://storeedgefd.dsx.mp.microsoft.com/v9.0` | No | Microsoft Store catalog (free apps rated "e") |
| **winget-font** | `Microsoft.PreIndexed.Package` | `https://cdn.winget.microsoft.com/fonts` | Yes | WinGet Community Repository for fonts |

A fourth well-known source, `microsoft.builtin.desktop.frameworks` (`https://cdn.winget.microsoft.com/platform`), exists for desktop framework packages.

### Explicit Sources

When a source is **explicit** (like `winget-font`), winget commands must directly target it using `--source <name>`; it is not included in automatic searches.

```
winget search --source winget-font
winget install --source winget-font <font-package>
```

## Source Agreements

Sources may require the user to agree to presented agreements before adding or using the source:

- **Source agreements** must be accepted before use
- If a source updates its agreement terms, or if a source is removed and re-added (e.g., `winget source reset --force`), agreements are presented again
- The `--accept-source-agreements` flag accepts agreements non-interactively
- There are also **package-level agreements** (accepted via `--accept-package-agreements`), separate from source-level agreements

### Agreement Structure

The `SourceInformation` structure includes:
- `SourceAgreementsIdentifier` — version stamp for agreement text
- `SourceAgreements` — list of `SourceAgreement` entries with `label`, `text`, and `url`

## Trust Levels (`SourceTrustLevel`)

Trust level is a bitmask enum stored on each source:

| Value | Meaning |
|-------|---------|
| `None` | No trust elevation |
| `Trusted` | Source packages are treated as trusted |
| `StoreOrigin` | Source is expected to carry Store-origin certificates |

When adding a source, the `--trust-level` option can specify `none` or `trusted`:

```
winget source add --name mysource --type Microsoft.Rest \
  --arg "https://my-rest-source.example.com/api" \
  --trust-level trusted
```

The default `winget` source has trust level `Trusted|StoreOrigin` (both flags set).

Sources also support **certificate pinning** via `CertificatePinningConfiguration` in the `SourceDetails` struct, ensuring that source connections only succeed with the expected server certificate.

## Source Origin (`SourceOrigin`)

Tracks where a source came from:

| Origin | Description |
|---------|-------------|
| `Default` | Ships with winget; can be tombstoned but not freely removed |
| `User` | Added by the user via `winget source add` |
| `Predefined` | Internal system source (Installed, Installing, etc.) |
| `GroupPolicy` | Injected by administrative group policy |
| `Metadata` | Represents metadata-only tracking records |
| `PackageTracking` | Used by the package tracking catalog |

Group policy may be configured to modify or restrict configured sources. Run `winget --info` to see any configured policies.

## How PreIndexed Sources Work

**PreIndexed Package** sources (`Microsoft.PreIndexed.Package`) package a complete SQLite database index inside an MSIX file. This provides:

1. **Pre-built indexes** — The SQLite database is constructed server-side, eliminating client-side indexing
2. **Signed delivery** — MSIX signature validation ensures package integrity and authenticity
3. **Version management** — Package version numbers enable efficient update detection
4. **Compressed distribution** — MSIX compression reduces download sizes
5. **Platform integration** — Leverages Windows deployment APIs and app extensions

### Execution Contexts

| Context | How Index is Accessed |
|---------|----------------------|
| **Packaged** (inside App Installer MSIX) | Uses Windows App Extensions system. The MSIX containing the index is deployed through `PackageManager` APIs and registered as an app extension, accessed via `ExtensionCatalog` API. |
| **Desktop** (unpackaged) | Downloads the MSIX to a local state directory and manually extracts the `index.db` file from it. |

### Update Mechanism

1. First checks the `x-ms-meta-sourceversion` HTTP header for version
2. If header missing, opens the remote MSIX and reads version from `AppxManifest.xml`
3. If remote version > local version, an update is triggered
4. Retry logic: on `ServiceUnavailableException`, backs off with `DoNotUpdateBefore` timestamps

### Trust Validation

For core sources, winget validates that the package is signed by a certificate chaining to the Microsoft Root Certificate. The `WriteLockedMsixFile` class ensures packages cannot be modified on disk during signature validation.

## How REST API Sources Work

REST sources use HTTP-based communication with remote package repositories implementing the WinGet REST API specification.

### API Contract Version Negotiation

The client supports multiple API schema versions and negotiates the highest common version with the server by matching major and minor versions:

| Version | Features Added |
|---------|----------------|
| 1.0.0 | Base manifest search, package retrieval |
| 1.1.0 | Additional search filters |
| 1.4.0 | Extended installer metadata |
| 1.5.0 | Dependencies support |
| 1.6.0 | Agreements support |
| 1.7.0 | Authentication support |
| 1.9.0 | Repair support |
| 1.10.0 | Archive installers |
| 1.12.0 | Enhanced metadata |
| 1.28.0 | Extended manifest fields |

### Core REST Endpoints

| Endpoint | Method | Purpose |
|----------|--------|---------|
| `/information` | GET | Retrieve server capabilities and supported versions |
| `/manifestSearch` | POST | Search for packages matching criteria |
| `/packageManifests/` | GET | Retrieve manifest for specific package version |

### HTTP Features

- Supports custom `Windows-Package-Manager` header (max 1024 chars), passed via `--header`
- Custom `User-Agent` string with caller info (anonymized via `%USERPROFILE%` replacement)
- `RestInformationCache` caches `/information` endpoint responses to reduce network overhead
- Optional `--header` parameter for custom headers on REST sources

## Source Commands

### `winget source add`

Add a new source.

| Flag | Short | Description |
|------|-------|-------------|
| `--name` | `-n` | Name of the source |
| `--type` | `-t` | Type of the source (`Microsoft.PreIndexed.Package`, `Microsoft.Rest`) |
| `--arg` | `-a` | Argument given to the source (typically the URL) |
| `--trust-level` | | Trust level of the source (`none`, `trusted`) |
| `--explicit` | | Excludes a source from discovery unless specified |
| `--header` | | Optional REST source HTTP header |
| `--accept-source-agreements` | | Accept all source agreements |

```
winget source add --name mysource --type Microsoft.Rest \
  --arg "https://my-rest-source.example.com/api" \
  --trust-level trusted
```

### `winget source list` (alias: `ls`)

List current sources.

| Flag | Short | Description |
|------|-------|-------------|
| `--name` | `-n` | Name of the source to show details for |

### `winget source remove` (alias: `rm`)

Remove current sources.

| Flag | Short | Description |
|------|-------|-------------|
| `--name` | `-n` | Name of the source to remove |

### `winget source update` (alias: `refresh`)

Update current sources (re-download index).

| Flag | Short | Description |
|------|-------|-------------|
| `--name` | `-n` | Name of the source to update |

### `winget source export`

Export current sources.

| Flag | Short | Description |
|------|-------|-------------|
| `--name` | `-n` | Name of the source to export |

### `winget source reset`

Reset sources (removes and re-adds default sources, re-presents agreements).

| Flag | Short | Description |
|------|-------|-------------|
| `--name` | `-n` | Name of the source to reset |

## Source Auto-Update

Sources automatically update based on the `source.autoUpdateIntervalInMinutes` setting in settings.json (default: 15 minutes). Setting this to `0` disables automatic updates. See [settings.md](settings.md) for details.

## Related

- For the overall winget architecture and package correlation, see [concepts.md](concepts.md).
- For source command flags and aliases, see [cli.md](cli.md).
- For settings that control source auto-update, see [settings.md](settings.md).
- For the REST API source reference implementation, see [winget-cli-restsource](https://github.com/microsoft/winget-cli-restsource).
