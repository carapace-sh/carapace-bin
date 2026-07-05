# Settings and Experimental Features

The winget settings.json file — all configuration sections, every setting key with types and defaults, experimental feature flags, and the `winget settings` commands.

> **Source of truth**: <https://github.com/microsoft/winget-cli/blob/master/doc/Settings.md>, <https://aka.ms/winget-settings.schema.json>. For **command reference**, see [cli.md](cli.md). For **install behavior** settings interaction, see [install-uninstall.md](install-uninstall.md).

## Settings File Location

| Context | Path |
|---------|------|
| **Packaged** (App Installer) | `%LOCALAPPDATA%\Packages\Microsoft.DesktopAppInstaller_8wekyb3d8bbwe\LocalState\settings.json` |
| **Unpackaged** (source build) | `%LOCALAPPDATA%\Microsoft\WinGet\Settings\settings.json` |

The settings file is a JSON file with a schema reference:

```json
{
    "$schema": "https://aka.ms/winget-settings.schema.json",
    "source": { ... },
    "visual": { ... },
    "installBehavior": { ... }
}
```

## `winget settings` Commands

| Subcommand | Description |
|------------|-------------|
| (no subcommand) | Opens settings.json in the default editor |
| `export` | Export settings |
| `set` | Sets the value of an admin setting |
| `reset` | Resets an admin setting to its default value |

### `winget settings set`

| Flag | Description |
|------|-------------|
| `--setting` | Name of the setting to modify |
| `--value` | Value to set for the setting |

### `winget settings reset`

| Flag | Description |
|------|-------------|
| `--setting` | Name of the setting to modify (reset to default) |

## Settings Sections

### `source`

Controls source auto-update behavior.

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `autoUpdateIntervalInMinutes` | integer | `15` | Positive integer = update interval in minutes. `0` disables checks. |

```json
"source": { "autoUpdateIntervalInMinutes": 3 }
```

### `visual`

Controls visual output.

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `progressBar` | string enum | `"accent"` | Progress spinner/bar style: `accent` (Windows Accent), `retro` (foreground), `rainbow`, `sixel` (requires sixel terminal), `disabled`. Overridable via `--no-progress`. |
| `anonymizeDisplayedPaths` | boolean | `true` | Replaces known folder paths with env variable equivalents. |
| `enableSixels` | boolean | `false` | Enables sixel image output in certain contexts. |

```json
"visual": {
    "progressBar": "accent",
    "anonymizeDisplayedPaths": true,
    "enableSixels": true
}
```

### `installBehavior`

Controls package installation behavior. Only applies to `winget install`, not `winget configure`.

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `disableInstallNotes` | boolean | `false` | Whether installation notes are shown after successful install. |
| `portablePackageUserRoot` | string (path) | `%LOCALAPPDATA%/Microsoft/WinGet/Packages/` | Root directory for user-scope portable package installs. |
| `portablePackageMachineRoot` | string (path) | `%PROGRAMFILES%/WinGet/Packages/` | Root directory for machine-scope portable package installs. |
| `skipDependencies` | boolean | `false` | Whether dependencies are installed for a given package. |
| `archiveExtractionMethod` | string enum | `"shellApi"` | How installer archives are extracted: `"tar"` (via tar.exe) or `"shellApi"` (Windows Shell API). |
| `defaultInstallRoot` | string (path) | — | Install location when a manifest includes `InstallLocationRequired`. Package ID is appended. Overridable by `--location`. |
| `maxResumes` | integer | `3` | Maximum times a command may be resumed automatically (requires experimental `resume`). |

#### `installBehavior.preferences`

Preferences for installer selection (acts as ordered preference, not a filter):

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `scope` | string enum | `"user"` | Preferred install scope: `"user"` or `"machine"`. |
| `locale` | array of strings (BCP47) | — | Preferred installer locales (ordered). |
| `architectures` | array of strings | — | Preferred architectures (e.g. `["x64", "arm64"]`). |
| `installerTypes` | array of strings | MSIX > MSI/Wix/Burn > Nullsoft/Inno/EXE > Portable | Preferred installer type priority order. |

#### `installBehavior.requirements`

Same fields as `preferences`, but acts as a **filter** rather than a preference. When multiple values are listed, they are treated as an ordered preference in addition to a filter. Command-line arguments override requirements for the duration of the command.

#### Allowed `installerTypes` Values

`appx`, `burn`, `exe`, `font`, `inno`, `msi`, `msix`, `msstore`, `nullsoft`, `portable`, `wix`, `zip`

```json
"installBehavior": {
    "preferences": {
        "scope": "user",
        "locale": ["en-US", "fr-FR"],
        "architectures": ["x64", "arm64"],
        "installerTypes": ["msi", "msix"]
    },
    "requirements": {
        "installerTypes": ["msix", "msi"]
    }
}
```

### `uninstallBehavior`

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `purgePortablePackage` | boolean | `false` | If `true`, uninstall removes all files/directories for `portable` packages. |

```json
"uninstallBehavior": { "purgePortablePackage": true }
```

### `configureBehavior`

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `defaultModuleRoot` | string (path) | `%LOCALAPPDATA%/Microsoft/WinGet/Configuration/Modules` | Default root directory where configuration modules are installed. |

```json
"configureBehavior": { "defaultModuleRoot": "C:/Program Files/Modules/" }
```

### `downloadBehavior`

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `defaultDownloadDirectory` | string (path) | `%USERPROFILE%/Downloads/` | Default directory where installers are downloaded to. |

```json
"downloadBehavior": { "defaultDownloadDirectory": "C:/Users/FooBar/Downloads" }
```

### `telemetry`

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `disable` | boolean | `false` | If `true`, prevents any ETW events from being written. |

```json
"telemetry": { "disable": true }
```

### `logging`

Controls log file behavior.

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `level` | string enum | `"info"` | Log detail level: `verbose`, `info`, `warning`, `error`, `critical`. `--verbose-logs` overrides to always create verbose. |
| `channels` | array of strings | `["default"]` | Channel identifiers (e.g. `CORE`). Special: `default` (default set), `all` (all channels). |
| `fileNameStrategy` | string enum | `"manifest"` | Log file naming: `manifest` (manifest name), `timestamp`, `guid`, `shortguid` (first 8 chars of GUID). |

#### `logging.file` (auto-cleanup)

Runs at start of each WinGet process; current process logs excluded from limits.

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `ageLimitInDays` | integer | `7` | Max age in days; older files deleted. `0` disables. |
| `totalSizeLimitInMB` | integer | `128` | Max total size in MB; oldest deleted first. `0` disables. |
| `countLimit` | integer | `0` | Max number of log files; oldest deleted first. `0` (default) disables. |
| `individualSizeLimitInMB` | integer | `16` | Max size of an individual log file; wraps when exceeded. `0` disables. |

```json
"logging": {
    "level": "info",
    "channels": ["default"],
    "fileNameStrategy": "manifest",
    "file": {
        "ageLimitInDays": 7,
        "totalSizeLimitInMB": 128,
        "countLimit": 0,
        "individualSizeLimitInMB": 16
    }
}
```

### `network`

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `downloader` | string enum | `"default"` | Download code: `default` (auto-selected), `wininet` (WinINet APIs), `do` (Delivery Optimization). |
| `doProgressTimeoutInSeconds` | integer | `60` | Seconds to wait without progress before fallback. Min: `1`, Max: `600`. |

```json
"network": {
    "downloader": "do",
    "doProgressTimeoutInSeconds": 60
}
```

### `interactivity`

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `disable` | boolean | `false` | If `true`, prevents any interactive prompts from winget itself (not installer prompts). Equivalent to `--disable-interactivity`. |

```json
"interactivity": { "disable": true }
```

## Experimental Features

Boolean flags in the `experimentalFeatures` section — each enables/disables an experimental feature. All default to `false`.

| Feature | Description |
|---------|-------------|
| `experimentalCmd` | Example experimental command feature. |
| `experimentalArg` | Example experimental argument feature. |
| `directMSI` | Installs MSI packages directly via MSI APIs instead of msiexec. Already in effect during silent installs for MSI packages requiring elevation. |
| `resume` | Enables resume support for some commands. |
| `fonts` | Enables font support via `winget font list` (lists installed font families and face counts). |
| `sourcePriority` | Enables priority values on sources. Higher priority sources appear earlier in search results and are preferred for new package installs. |

```json
"experimentalFeatures": {
    "experimentalCmd": true,
    "experimentalArg": false,
    "directMSI": true,
    "resume": true,
    "fonts": true,
    "sourcePriority": true
}
```

### `winget features`

The `winget features` command shows the current status of all experimental features — whether each is enabled or disabled.

## Complete Settings Example

```json
{
    "$schema": "https://aka.ms/winget-settings.schema.json",
    "source": {
        "autoUpdateIntervalInMinutes": 15
    },
    "visual": {
        "progressBar": "accent",
        "anonymizeDisplayedPaths": true
    },
    "installBehavior": {
        "disableInstallNotes": false,
        "portablePackageUserRoot": "C:/Users/Me/PortableApps/",
        "preferences": {
            "scope": "user",
            "architectures": ["x64"],
            "installerTypes": ["msix", "msi", "exe"]
        },
        "requirements": {
            "scope": "user"
        }
    },
    "uninstallBehavior": {
        "purgePortablePackage": false
    },
    "telemetry": {
        "disable": false
    },
    "logging": {
        "level": "info",
        "fileNameStrategy": "manifest"
    },
    "network": {
        "downloader": "default"
    },
    "interactivity": {
        "disable": false
    },
    "experimentalFeatures": {
        "directMSI": true,
        "sourcePriority": true
    }
}
```

## Summary of All Sections

| Section | Purpose |
|---------|---------|
| `source` | Source update interval configuration |
| `visual` | Progress bar style, path anonymization, sixel output |
| `installBehavior` | Install defaults, preferences, requirements, portable roots |
| `uninstallBehavior` | Portable package purge behavior |
| `configureBehavior` | Configuration module root directory |
| `downloadBehavior` | Default download directory |
| `telemetry` | ETW event emission (privacy control) |
| `logging` | Log level, channels, file naming, log rotation/limits |
| `network` | Downloader selection (WinINet vs Delivery Optimization) |
| `interactivity` | Interactive prompt suppression |
| `experimentalFeatures` | Feature flags for in-development functionality |

## Related

- For install behavior settings interaction (how preferences affect installer selection), see [install-uninstall.md](install-uninstall.md).
- For source auto-update settings, see [sources.md](sources.md).
- For command reference, see [cli.md](cli.md).
