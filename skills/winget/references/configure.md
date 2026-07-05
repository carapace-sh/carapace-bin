# WinGet Configuration and DSC v3

Declarative machine setup with WinGet Configuration — the `.winget` file format, v2 and v3 schema versions, DSC resource types, the configure test/get/set workflow, and the `winget dscv3` resource commands.

> **Source of truth**: <https://learn.microsoft.com/en-us/windows/package-manager/winget/configure>, <https://learn.microsoft.com/en-us/windows/package-manager/configuration/>. For **command reference**, see [cli.md](cli.md). For **settings** controlling configuration behavior, see [settings.md](settings.md).

## What WinGet Configuration Is

WinGet Configuration is a feature of the Windows Package Manager that enables **declarative, reproducible machine setup** using a YAML configuration file. Instead of manually installing packages and configuring settings one-by-one, you define a **desired state** in a file, and winget uses PowerShell Desired State Configuration (DSC) to bring the machine to that state automatically.

Introduced in WinGet v1.6.2631. Requires Windows 10 version 1809 (build 17763) or later, or Windows 11.

## Configuration File Format

### File Naming

- Files use the **`.winget`** extension (e.g., `configuration.winget`)
- For Git-based projects, the default location is `./.config/configuration.winget`

### Two Schema Versions

| Aspect | v2 Schema (0.2) | v3 Schema (DSC v3) |
|--------|-------------------|---------------------|
| **Schema URL** | `https://aka.ms/configuration-dsc-schema/0.2` | `https://raw.githubusercontent.com/PowerShell/DSC/main/schemas/2023/08/config/document.json` |
| **WinGet version** | v1.6.2631+ | v1.11+ |
| **Processor** | Built-in DSC v2 processor | `dscv3` processor (`Microsoft.DesiredStateConfiguration`) |
| **Structure** | Resources nested under `properties.resources` | Resources at document root level |
| **Resource types** | PowerShell DSC v2 class-based | Native DSC v3 resources + adapted v2 resources |
| **Module install** | Automatic by WinGet | Explicit via `RunCommandOnSet` resources |

## v2 Schema Structure

```yaml
# yaml-language-server: $schema=https://aka.ms/configuration-dsc-schema/0.2
properties:
  configurationVersion: 0.2.0
  assertions:
    - resource: Microsoft.Windows.Developer/OsVersion
      directives:
        description: Verify min OS version requirement
        allowPrerelease: true
      settings:
        MinVersion: '10.0.22000'
  resources:
    - resource: Microsoft.WinGet.DSC/WinGetPackage
      id: vsPackage
      directives:
        description: Install Visual Studio 2022 Community
        securityContext: elevated
      settings:
        id: Microsoft.VisualStudio.2022.Community
        source: winget
```

### Top-Level Components

| Key | Description |
|-----|-------------|
| `# yaml-language-server` | Schema comment for YAML language server support (VS Code) |
| `properties` | Root node containing everything |
| `properties.configurationVersion` | SemVer string (e.g., `0.2.0`) |
| `properties.assertions` | Array of preconditions (e.g., OS version check) |
| `properties.resources` | Array of configuration tasks (packages, settings, registry, etc.) |

## v3 Schema Structure (DSC v3)

```yaml
$schema: https://raw.githubusercontent.com/PowerShell/DSC/main/schemas/2023/08/config/document.json
metadata:
  winget:
    processor:
      identifier: dscv3
resources:
  - type: Microsoft.WinGet/Package
    name: PowerShell
    properties:
      id: Microsoft.PowerShell
      source: winget
      useLatest: true
    metadata:
      description: Install PowerShell 7
```

### Top-Level Keys

| Key | Description |
|-----|-------------|
| `$schema` | Points to the DSC v3 JSON schema |
| `metadata` | Contains `winget.processor.identifier: dscv3` to select the DSC v3 processor |
| `resources` | Top-level array of resource instances (no `properties` wrapper) |

## Resource Fields

| Field | v2 Name | v3 Name | Required | Description |
|-------|---------|---------|----------|-------------|
| Resource identifier | `resource` | `type` | Yes | Format: `{Provider}/{Resource}` |
| Identifier | `id` | `name` | v2: No, v3: Yes | Unique name for dependency references |
| Settings/Properties | `settings` | `properties` | Yes | Key-value pairs passed to the DSC resource |
| Directives/Metadata | `directives` | `metadata` | Yes | Info about the resource (description, securityContext) |
| Dependencies | `dependsOn` | `dependsOn` | No | Array of resource names that must complete first |

### Resource Type Mapping (v2 → v3)

| v2 Resource Type | v3 Resource Type | Notes |
|------------------|------------------|-------|
| `Microsoft.WinGet.DSC/WinGetPackage` | `Microsoft.WinGet/Package` | Add `useLatest: true` in v3 |
| `PSDscResources/Registry` | `Microsoft.Windows/Registry` | Native v3 — dramatically faster |
| `PSDscResources/Script` | `Microsoft.DSC.Transitional/RunCommandOnSet` | Or `PowerShellScript` / `WindowsPowerShellScript` |
| `Microsoft.Windows.Developer/*` | `Microsoft.Windows.Developer/*` | Unchanged type, needs explicit module install in v3 |
| `Microsoft.Windows.Settings/*` | `Microsoft.Windows.Settings/*` | Unchanged type, needs explicit module install in v3 |

## Directives and Metadata

### v2 (`directives`)

```yaml
directives:
  description: Enable Developer Mode
  allowPrerelease: true
  securityContext: elevated
```

### v3 (`metadata`)

```yaml
metadata:
  winget:
    securityContext: elevated
  description: Enable Developer Mode
```

### Security Context Values

| Value | Description |
|-------|-------------|
| `current` | Run with current user's privileges |
| `restricted` | Run in a restricted context |
| `elevated` | Run with administrator privileges (triggers one UAC prompt for the whole config) |

### When to Use `elevated`

| Scenario | Elevated? |
|----------|-----------|
| Package with `ElevationRequirement: elevationRequired` | Yes |
| Package that installs a service | Yes |
| Registry key under `HKLM\` | Yes |
| Registry key under `HKCU\` | No |
| `Microsoft.Windows.Developer/EnableLongPathSupport` | Yes |
| `Microsoft.Windows.Settings/WindowsSettings` (DeveloperMode) | Yes |
| `Microsoft.Windows.Developer/WindowsExplorer` | No |
| `Microsoft.Windows.Developer/Taskbar` | No |

## Dependencies (`dependsOn`)

The `dependsOn` field specifies that a resource must wait for another resource to complete. If the dependency fails, the dependent resource also fails.

- In v2, `dependsOn` references the `id` value
- In v3, `dependsOn` references the `name` value

```yaml
# v3 example
- type: Microsoft.WinGet/Package
  name: Ubuntu
  dependsOn:
    - WSL
  properties:
    id: Canonical.Ubuntu
    source: winget
    useLatest: true
```

## Assertions

Assertions are preconditions that must be met before resources run. Only in v2 schema.

```yaml
assertions:
  - resource: Microsoft.Windows.Developer/OsVersion
    directives:
      description: Verify min OS version requirement
      allowPrerelease: true
    settings:
      MinVersion: '10.0.22000'
```

- Assertions run in parallel
- If an assertion returns `false`, any resource that depends on it (via `dependsOn`) is skipped
- The configuration continues running other resources even if some assertions fail

## DSC v3 Resource Types

DSC v3 supports three kinds of resources:

| Type | Examples | Description |
|------|----------|-------------|
| **Native v3** | `Microsoft.Windows/Registry`, `Microsoft.WinGet/Package` | Fast, compiled, no PowerShell needed |
| **Adapted v2** | `Microsoft.Windows.Developer/*`, `PSDscResources/*` | PowerShell class-based DSC v2 resources through a compatibility adapter |
| **Transitional** | `Microsoft.DSC.Transitional/RunCommandOnSet`, `PowerShellScript`, `WindowsPowerShellScript` | Bridge resources for running arbitrary commands/scripts |

### Performance Characteristics

| Resource Type | Approx. Eval Cost | Scaling |
|---------------|-------------------|---------|
| `Microsoft.Windows/Registry` (native v3) | ~3s each | Linear |
| `Microsoft.DSC.Transitional/RunCommandOnSet` | ~5s each | Linear |
| `Microsoft.WinGet/Package` | ~8–16s each | Linear |
| `Microsoft.Windows.Settings/*` (adapted) | ~95s each | Variable |
| `Microsoft.Windows.Developer/*` (adapted) | ~30–122s each | Super-linear (worse when elevated) |
| `PSDscResources/Registry` (adapted) | ~89–228s each | Super-linear — avoid |

Fixed overhead for dscv3 processor startup: ~23 seconds.

## `winget configure` Commands

### `winget configure` (default — apply)

Applies the configuration — installs packages, configures settings, sets registry values.

```
winget configure -f <file.winget>
```

| Flag | Short | Description |
|------|-------|-------------|
| `--file` | `-f` | The path to the configuration file |
| `--module-path` | | Location on local computer to store modules |
| `--processor-path` | | Path to the configuration processor |

### `winget configure test`

Checks the system against the desired state and reports whether the current state conforms. No changes are applied.

```
winget configure test -f <file.winget>
```

### `winget configure show` (alias: `view`)

Displays the details of a configuration file without applying it.

```
winget configure show -f <file.winget>
```

### `winget configure validate`

Validates the syntax and structure of a configuration file.

```
winget configure validate -f <file.winget>
```

### `winget configure list` (alias: `ls`)

Shows the history of configurations that have been applied to the system.

| Flag | Short | Description |
|------|-------|-------------|
| `--history` | `-h` | Select items from history |
| `--output` | `-o` | File where the result is to be written |
| `--remove` | | Remove the item from history |

### `winget configure export`

Exports configuration resources to a configuration file.

| Flag | Short | Description |
|------|-------|-------------|
| `--all` | | Exports all package configurations |
| `--include-versions` | | Include package versions in export file |
| `--module` | | The module of the resource to export |
| `--output` | `-o` | File where the result is to be written |
| `--package-id` | | The package identifier to export |
| `--recurse` | `-r` | Exports all package configurations |
| `--resource` | | The configuration resource to export |
| `--source` | `-s` | Export packages from the specified source |

## `winget dscv3` Commands

The `winget dscv3` command exposes DSC v3 resources for direct management, outside of a configuration file. Each subcommand supports the standard DSC operations.

| Persistent Flag | Short | Description |
|-----------------|-------|-------------|
| `--manifest` | | Get the resource manifest |
| `--output` | `-o` | Directory where results are written |

### Subcommands

| Subcommand | Description |
|------------|-------------|
| `admin-settings` | Manage administrator settings |
| `package` | Manage package state |
| `source` | Manage source configuration |
| `user-settings-file` | Manage user settings file |

### Operation Flags (all subcommands)

| Flag | Description |
|------|-------------|
| `--export` | Get all state instances |
| `--get` | Get the resource state |
| `--schema` | Get the resource schema |
| `--set` | Set the resource state |
| `--test` | Test the resource state |

These map to the standard DSC v3 operations: **Get** (retrieve current state), **Set** (apply desired state), **Test** (check if in desired state), **Export** (list all instances), **Schema** (retrieve JSON schema for the resource).

## Configuration Workflow

### How It Works Internally

1. **Parse** — WinGet parses the YAML file and validates it
2. **Download modules** — All required PowerShell DSC modules are downloaded (v2 auto-installs; v3 requires explicit `RunCommandOnSet`)
3. **Test** — Each resource is tested to see if the system is already in the desired state
4. **Apply** — Resources not in the desired state are applied (set)
5. **Report** — Results are displayed, showing successes and failures

The configuration is **idempotent** — running it multiple times produces the same result. If the system is already in the desired state, no changes are made.

## Security Considerations

Before running any WinGet Configuration file:

1. **Always review the file contents** — It defines what will be installed/changed
2. **Check PowerShell module sources** — Use `Get-PSRepository` to verify where modules come from
3. **Test in an isolated environment** — Use a VM, Azure VM, or Windows Sandbox
4. **DSC resources can run arbitrary code** — Including pulling additional binaries from the internet
5. **One UAC prompt for the entire configuration** — Can be suppressed with `--accept-configuration-agreements`
6. **Group Policy controls** — `EnableWindowsPackageManagerConfiguration` can disable the feature organization-wide

## Example: v2 Configuration

```yaml
# yaml-language-server: $schema=https://aka.ms/configuration-dsc-schema/0.2
properties:
  configurationVersion: 0.2.0
  resources:
    - resource: Microsoft.Windows.Developer/DeveloperMode
      id: enable-devmode
      directives:
        description: Enable Developer Mode
        allowPrerelease: true
        securityContext: elevated
      settings:
        Ensure: Present
    - resource: Microsoft.WinGet.DSC/WinGetPackage
      id: Mozilla.Firefox
      directives:
        description: Installing Mozilla Firefox
        securityContext: current
      settings:
        id: "Mozilla.Firefox"
        source: winget
    - resource: Microsoft.WinGet.DSC/WinGetPackage
      id: Remove7Zip
      directives:
        description: Remove 7Zip
        securityContext: elevated
      settings:
        id: 7zip.7zip
        ensure: absent
```

## Example: v3 Configuration

```yaml
$schema: https://raw.githubusercontent.com/PowerShell/DSC/main/schemas/2023/08/config/document.json
metadata:
  winget:
    processor:
      identifier: dscv3
resources:
  - type: Microsoft.WinGet/Package
    name: PowerShell
    properties:
      id: Microsoft.PowerShell
      source: winget
      useLatest: true
    metadata:
      description: Install PowerShell 7

  - type: Microsoft.WinGet/Package
    name: VSCode
    properties:
      id: Microsoft.VisualStudioCode
      source: winget
      useLatest: true
    metadata:
      description: Install Visual Studio Code

  - type: Microsoft.DSC.Transitional/RunCommandOnSet
    name: InstallDeveloperModule
    dependsOn:
      - PowerShell
    properties:
      executable: C:\Program Files\PowerShell\7\pwsh.exe
      arguments:
        "0": -NoProfile
        "1": -NoLogo
        "2": -Command
        "3": >-
          if (-not (Get-Module -ListAvailable -Name Microsoft.Windows.Developer))
          { Install-PSResource -Name Microsoft.Windows.Developer -Prerelease
          -TrustRepository -AcceptLicense }
      treatAsArray: true
    metadata:
      description: Ensure Microsoft.Windows.Developer module is installed

  - type: Microsoft.Windows/Registry
    name: EnableLongPaths
    properties:
      keyPath: HKLM\SYSTEM\CurrentControlSet\Control\FileSystem
      valueName: LongPathsEnabled
      valueData:
        DWord: 1
      _exist: true
    metadata:
      winget:
        securityContext: elevated
      description: Enable Win32 long path support
```

## Related

- For command structure and global flags, see [cli.md](cli.md).
- For settings controlling configuration module root, see [settings.md](settings.md).
- For the DSC v3 specification, see the [PowerShell DSC repository](https://github.com/PowerShell/DSC).
- For sample configurations, see the [winget-dsc samples](https://github.com/microsoft/winget-dsc/tree/main/samples).
