---
name: winget
description: >
  Use when working with the Windows Package Manager CLI (`winget`) — package installation,
  search, upgrade, uninstall, repair, sources, package manifests (YAML schema), WinGet
  Configuration / DSC v3, settings.json, pins, export/import, download, hash, validate,
  experimental features, and the dscv3 resource commands. Triggers on: "winget", "winget install",
  "winget search", "winget upgrade", "winget uninstall", "winget list", "winget source",
  "winget configure", "winget settings", "winget pin", "winget export", "winget import",
  "winget download", "winget hash", "winget validate", "winget features", "winget dscv3",
  "winget repair", "winget mcp", "PackageIdentifier", "winget-pkgs", "manifest",
  "PreIndexed", "Microsoft.Rest", "msstore", "settings.json", "configuration.winget",
  "DSC v3", "WinGet Configuration", "portable package", "MSIX", "ARP".
user-invocable: true
---

# winget — Windows Package Manager CLI In-Depth Reference

Comprehensive reference for [winget](https://github.com/microsoft/winget-cli), the Windows Package Manager command-line tool. Covers the command structure, source system, package manifests, install/upgrade/uninstall workflows, WinGet Configuration (DSC v3), settings, and package management operations.

## Data Flow

```
winget command
  → parse global flags (--disable-interactivity, --verbose, --proxy, --no-proxy)
    → resolve sources (winget, msstore, user-added, REST, predefined Installed)
      → search across sources (CompositeSource aggregates results)
        → match package (PackageIdentifier, ProductCode, PackageFamilyName)
          → download installer → verify SHA256 → execute installer
            → record in Installed source (ARP registry + MSIX)
```

```
WinGet Configuration:
  winget configure -f file.winget
    → parse YAML (v2 or v3 schema)
      → download DSC modules (v2) / verify dscv3 processor (v3)
        → test each resource (is system in desired state?)
          → set resources not in desired state
            → report results (idempotent)
```

## Sub-Resources

Load the reference that matches your task. When in doubt, load multiple references.

| Keywords | Reference |
|----------|----------|
| concept, model, architecture, package manager, PackageIdentifier, ProductCode, PackageFamilyName, UpgradeCode, moniker, installed packages, ARP, Add/Remove Programs, MSIX, package correlation, CompositeSource, search request, version mapping, binary inventory, packaged vs unpackaged, COM server, WinRT API, winget.exe, WindowsPackageManager.dll, group policy | [references/concepts.md](references/concepts.md) |
| cli, command, subcommand, global flags, --disable-interactivity, --verbose, --verbose-logs, --proxy, --no-proxy, --wait, --open-logs, --logs, --ignore-warnings, --nowarn, --info, --version, -v, --help, -?, command aliases, install, search, show, list, upgrade, uninstall, source, configure, settings, pin, export, import, download, hash, validate, features, dscv3, repair, mcp | [references/cli.md](references/cli.md) |
| source, source add, source list, source remove, source update, source reset, source export, PreIndexed, Microsoft.PreIndexed.Package, Microsoft.Rest, REST API source, msstore, winget source, winget-font, well-known sources, trust level, SourceTrustLevel, source agreements, accept-source-agreements, source origin, SourceOrigin, explicit source, certificate pinning, SQLite index, MSIX delivery, REST API version negotiation, --header, --type, --trust-level, --arg, --explicit | [references/sources.md](references/sources.md) |
| manifest, YAML, PackageIdentifier, PackageVersion, DefaultLocale, PackageLocale, ManifestType, ManifestVersion, version manifest, installer manifest, defaultLocale manifest, locale manifest, InstallerType, msi, msix, exe, nullsoft, inno, wix, burn, portable, zip, font, appx, pwa, Architecture, x86, x64, arm, arm64, neutral, Scope, user, machine, InstallerSwitches, Silent, SilentWithProgress, Interactive, InstallLocation, Log, Upgrade, Custom, Repair, Dependencies, WindowsFeature, WindowsLibrary, Package, External, NestedInstallerType, NestedInstallerFiles, UpgradeBehavior, ElevationRequirement, ExpectedReturnCodes, AppsAndFeaturesEntries, ProductCode, PackageFamilyName, Capabilities, Commands, Protocols, FileExtensions, Markets, ReleaseDate, InstallationMetadata, Authentication, winget-pkgs, manifest validation, winget validate, winget hash | [references/manifest.md](references/manifest.md) |
| install, uninstall, upgrade, repair, installer type, scope, architecture, silent, interactive, silentWithProgress, --exact, -e, --force, --silent, -h, --interactive, -i, --location, -l, --log, -o, --manifest, -m, --version, -v, --source, -s, --id, --name, --moniker, --query, -q, --scope, --architecture, -a, --installer-type, --override, --custom, --no-upgrade, --skip-dependencies, --accept-package-agreements, --accept-source-agreements, --dependency-source, --authentication-account, --authentication-mode, --header, --locale, --allow-reboot, --uninstall-previous, --ignore-security-hash, --ignore-local-archive-malware-scan, --rename, -r, portable package, install behavior, upgrade all, --all, --recurse, -r, --include-unknown, --include-pinned, uninstall --all, --all-versions, --preserve, --purge, --product-code | [references/install-uninstall.md](references/install-uninstall.md) |
| configure, configuration, DSC, DSC v3, DSC v2, .winget, configuration.winget, configuration file, configurationVersion, assertions, resources, dependsOn, directives, securityContext, elevated, current, restricted, WinGetPackage, Microsoft.WinGet.DSC, Microsoft.WinGet/Package, Microsoft.Windows/Registry, RunCommandOnSet, PowerShellScript, WindowsPowerShellScript, configure test, configure show, configure validate, configure list, configure export, dscv3, dscv3 package, dscv3 source, dscv3 admin-settings, dscv3 user-settings-file, processor, identifier, schema, get, set, test, export, idempotent, desired state, module-path, processor-path, --file, -f | [references/configure.md](references/configure.md) |
| settings, settings.json, config, settings export, settings set, settings reset, source.autoUpdateIntervalInMinutes, visual, progressBar, anonymizeDisplayedPaths, enableSixels, installBehavior, preferences, requirements, scope, locale, architectures, installerTypes, uninstallBehavior, purgePortablePackage, configureBehavior, downloadBehavior, telemetry, logging, network, downloader, wininet, do, Delivery Optimization, interactivity, experimentalFeatures, experimentalCmd, experimentalArg, directMSI, resume, fonts, sourcePriority, portablePackageUserRoot, portablePackageMachineRoot, defaultInstallRoot, skipDependencies, archiveExtractionMethod, fileNameStrategy, ageLimitInDays, totalSizeLimitInMB, countLimit, individualSizeLimitInMB, --setting, --value, winget features, experimental features | [references/settings.md](references/settings.md) |
| search, list, show, pin, pin add, pin list, pin remove, pin reset, export, import, download, hash, validate, features, winget show, winget view, show, view, --count, -n, --exact, -e, --versions, --tag, --cmd, --command, --upgrade-available, --include-pinned, --include-unknown, --blocking, --installed, --include-versions, --ignore-unavailable, --ignore-versions, --import-file, -i, --no-upgrade, --download-directory, -d, --skip-license, --skip-microsoft-store-package-license, --os-version, --platform, --output, -o, --blocking, pin blocking, pin installed, export import workflow, package query, filter, moniker, id, name | [references/package-management.md](references/package-management.md) |

## Quick Guide

- **What is winget and how does it work?** → [references/concepts.md](references/concepts.md)
- **What are the commands and global flags?** → [references/cli.md](references/cli.md)
- **How do sources work and how do I add one?** → [references/sources.md](references/sources.md)
- **What is the package manifest YAML format?** → [references/manifest.md](references/manifest.md)
- **How do I install, upgrade, or uninstall a package?** → [references/install-uninstall.md](references/install-uninstall.md)
- **How do I write a WinGet Configuration file?** → [references/configure.md](references/configure.md)
- **How do I configure winget settings?** → [references/settings.md](references/settings.md)
- **How do I search, list, pin, export, or import packages?** → [references/package-management.md](references/package-management.md)
- **How do I show details of a package?** → [references/package-management.md](references/package-management.md)
- **How do I validate or hash a manifest?** → [references/manifest.md](references/manifest.md) and [references/package-management.md](references/package-management.md)
- **What installer types and scopes are supported?** → [references/manifest.md](references/manifest.md) and [references/install-uninstall.md](references/install-uninstall.md)
- **How do I download an installer without installing?** → [references/package-management.md](references/package-management.md)
- **What experimental features exist?** → [references/settings.md](references/settings.md)
- **How do DSC v3 resources work with winget?** → [references/configure.md](references/configure.md)
- **How does winget match installed packages to available ones?** → [references/concepts.md](references/concepts.md)
- **What is the difference between the v2 and v3 configuration schemas?** → [references/configure.md](references/configure.md)

## Cross-Project References

- For **carapace** completion integration (the `winget` completer in carapace-bin at `completers/windows/winget_completer/`, shared actions in `pkg/actions/tools/winget/`), see the **carapace** skill.
- For **PowerShell DSC** (Desired State Configuration) as a standalone technology beyond winget's integration, see the [Microsoft DSC documentation](https://learn.microsoft.com/en-us/powershell/dsc/). This skill covers winget's orchestration of DSC v3 only.
- For **MSIX** packaging format internals beyond winget's consumption, see the [MSIX documentation](https://learn.microsoft.com/en-us/windows/msix/). This skill covers how winget uses MSIX for source delivery and package installation.
