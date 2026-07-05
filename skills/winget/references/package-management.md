# Package Management — Search, List, Pin, Export, Import, Download, Hash, Validate

The discovery and management commands in winget — searching for packages, listing installed packages, managing pins, exporting/importing package lists, downloading installers, and hashing/validating manifests.

> **Source of truth**: <https://learn.microsoft.com/en-us/windows/package-manager/winget/>, `winget <command> --help`. For **install/upgrade/uninstall**, see [install-uninstall.md](install-uninstall.md). For **manifest schema**, see [manifest.md](manifest.md). For **command structure**, see [cli.md](cli.md).

## `winget search` (alias: `find`)

Find and show basic info of packages from configured sources.

### Usage

```
winget search <query>
winget search --id Microsoft.VisualStudioCode
winget search --name "Visual Studio Code"
winget search --moniker vscode
winget search --tag editor
winget search --command code
winget search --source winget
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-source-agreements` | | Accept all source agreements |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference (`silent`, `silentPreferred`, `interactive`) |
| `--cmd` | | Filter results by command |
| `--command` | | Filter results by command |
| `--count` | `-n` | Show no more than specified number of results |
| `--exact` | `-e` | Find package using exact match |
| `--header` | | REST source HTTP header |
| `--id` | | Filter results by id |
| `--moniker` | | Filter results by moniker |
| `--name` | | Filter results by name |
| `--query` | `-q` | The query used to search for a package |
| `--source` | `-s` | Find package using the specified source |
| `--tag` | | Filter results by tag |
| `--versions` | | Show available versions of the package |

### Search Behavior

- Without `--exact`, filters match by substring (case-insensitive)
- With `--exact`, filters must match exactly
- `--source` restricts search to a specific source (e.g., `winget`, `msstore`)
- `--count` limits results (useful for scripting)
- `--versions` shows all available versions for matching packages
- `--tag` and `--command`/`--cmd` filter by manifest fields

### Authentication Mode Values

| Mode | Description |
|------|-------------|
| `silent` | Attempt silent authentication; fail if not possible |
| `silentPreferred` | Attempt silent; fall back to interactive (default) |
| `interactive` | Always use interactive authentication window |

## `winget show` (alias: `view`)

Displays detailed information for a specified package — full manifest data including installer URLs, SHA256 hashes, description, license, and available versions.

### Usage

```
winget show <query>
winget show --id Microsoft.VisualStudioCode
winget show --name "Visual Studio Code"
winget show --moniker vscode --exact
winget show --manifest .\package.yaml
winget show --id Microsoft.VisualStudioCode --versions
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-source-agreements` | | Accept all source agreements |
| `--architecture` | `-a` | Select the architecture to show |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference (`silent`, `silentPreferred`, `interactive`) |
| `--exact` | `-e` | Find package using exact match |
| `--header` | | REST source HTTP header |
| `--id` | | Filter results by id |
| `--installer-type` | | Select the installer type to show |
| `--locale` | | Locale to use (BCP47 format) |
| `--manifest` | `-m` | The path to the manifest of the package |
| `--moniker` | | Filter results by moniker |
| `--name` | | Filter results by name |
| `--query` | `-q` | The query used to search for a package |
| `--scope` | | Select install scope (`user`, `machine`) |
| `--source` | `-s` | Find package using the specified source |
| `--version` | `-v` | Use the specified version; default is the latest |
| `--versions` | | Show available versions of the package |

### `--versions`

Shows all available versions of the package from the configured sources, not just the latest. Without this flag, `winget show` displays only the latest version's manifest details.

## `winget list` (alias: `ls`)

Display installed packages from the `Installed` predefined source (ARP + MSIX).

### Usage

```
winget list                           # all installed packages
winget list --upgrade-available       # only packages with upgrades
winget list --id Microsoft.VisualStudioCode
winget list --name "Visual Studio"
winget list --source winget           # only packages from winget source
winget list --scope user              # only user-scope packages
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-source-agreements` | | Accept all source agreements |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference |
| `--cmd` | | Filter results by command |
| `--command` | | Filter results by command |
| `--count` | `-n` | Show no more than specified number of results (1–1000) |
| `--exact` | `-e` | Find package using exact match |
| `--header` | | REST source HTTP header |
| `--id` | | Filter results by id |
| `--include-pinned` | | List packages even if pinned (with `--upgrade-available`) |
| `--pinned` | | List packages even if pinned (alias for `--include-pinned`) |
| `--include-unknown` | | List packages even if version unknown (with `--upgrade-available`) |
| `--unknown` | `-u` | List packages even if version unknown |
| `--moniker` | | Filter results by moniker |
| `--name` | | Filter results by name |
| `--query` | `-q` | The query used to search |
| `--scope` | | Select installed package scope filter (`user`, `machine`) |
| `--source` | `-s` | Find package using the specified source |
| `--tag` | | Filter results by tag |
| `--upgrade-available` | | Lists only packages which have an upgrade available |

### `--upgrade-available`

Filters the list to only show packages that have an available upgrade. The `--include-pinned` and `--include-unknown` flags modify this filter:

- `--include-pinned` / `--pinned`: Include packages with non-blocking pins
- `--include-unknown` / `--unknown`: Include packages where the current version cannot be determined

## `winget pin`

Manage package pins. Pins prevent or warn about upgrades to specific packages.

### `winget pin add`

Add a new pin to a package.

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-source-agreements` | | Accept all source agreements |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference |
| `--blocking` | | Block from upgrading until pin is removed, preventing override arguments |
| `--cmd` | | Filter results by command |
| `--command` | | Filter results by command |
| `--exact` | `-e` | Find package using exact match |
| `--force` | | Direct run, continue with non-security issues |
| `--header` | | REST source HTTP header |
| `--id` | | Filter results by id |
| `--installed` | | Pin a specific installed version |
| `--moniker` | | Filter results by moniker |
| `--name` | | Filter results by name |
| `--query` | `-q` | The query used to search |
| `--source` | `-s` | Find package using the specified source |
| `--tag` | | Filter results by tag |
| `--version` | `-v` | Version to which to pin the package |

### Pin Types

| Type | Behavior |
|------|----------|
| **Non-blocking** (default) | Warns about the upgrade but can be overridden with `--include-pinned` |
| **Blocking** (`--blocking`) | Prevents upgrade entirely until the pin is removed; cannot be overridden |

### `--installed`

Pins the currently installed version of the package, preventing upgrades from that specific version.

### `winget pin list`

List current pins.

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-source-agreements` | | Accept all source agreements |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference |
| `--cmd` | | Filter results by command |
| `--command` | | Filter results by command |
| `--exact` | `-e` | Find package using exact match |
| `--header` | | REST source HTTP header |
| `--id` | | Filter results by id |
| `--moniker` | | Filter results by moniker |
| `--name` | | Filter results by name |
| `--query` | `-q` | The query used to search |
| `--source` | `-s` | Find package using the specified source |
| `--tag` | | Filter results by tag |

### `winget pin remove`

Remove a package pin.

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-source-agreements` | | Accept all source agreements |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference |
| `--exact` | `-e` | Find package using exact match |
| `--header` | | REST source HTTP header |
| `--id` | | Filter results by id |
| `--installed` | | Pin a specific installed version |
| `--moniker` | | Filter results by moniker |
| `--name` | | Filter results by name |
| `--query` | `-q` | The query used to search |
| `--source` | `-s` | Find package using the specified source |
| `--tag` | | Filter results by tag |

### `winget pin reset`

Reset pins.

| Flag | Short | Description |
|------|-------|-------------|
| `--force` | | Direct run, continue with non-security issues |
| `--source` | `-s` | Find package using the specified source |

## `winget export`

Exports a list of the installed packages to a JSON file. Useful for backing up installed packages or replicating an environment.

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-source-agreements` | | Accept all source agreements |
| `--include-versions` | | Include package versions in export file |
| `--output` | `-o` | File where the result is to be written |
| `--source` | `-s` | Export packages from the specified source |

### Export File Format

The export produces a JSON file containing package identifiers and (optionally) versions:

```json
{
    "$schema": "https://aka.ms/winget-packages.schema.2.0.json",
    "CreationDate": "2024-01-15T12:00:00.000Z",
    "Sources": [
        {
            "Packages": [
                { "Id": "Microsoft.VisualStudioCode" },
                { "Id": "Git.Git", "Version": "2.43.0" }
            ],
            "SourceDetails": {
                "Name": "winget",
                "Type": "Microsoft.PreIndexed.Package",
                "Argument": "https://cdn.winget.microsoft.com/cache"
            }
        }
    ],
    "IncludeVersions": true
}
```

## `winget import`

Installs all the packages in a file (typically produced by `winget export`).

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-package-agreements` | | Accept all license agreements for packages |
| `--accept-source-agreements` | | Accept all source agreements |
| `--ignore-unavailable` | | Ignore unavailable packages |
| `--ignore-versions` | | Ignore package versions in import file |
| `--import-file` | `-i` | File describing the packages to install |
| `--no-upgrade` | | Skip upgrade if an installed version already exists |

### Import Behavior

- Without `--ignore-versions`: installs the exact versions specified in the file
- With `--ignore-versions`: installs the latest available version of each package
- With `--ignore-unavailable`: skips packages that are not found in any source
- With `--no-upgrade`: skips packages that are already installed (regardless of version)

## `winget download`

Downloads the installer from a given package without installing it. Useful for offline installation or inspection.

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--accept-package-agreements` | | Accept all license agreements |
| `--accept-source-agreements` | | Accept all source agreements |
| `--architecture` | `-a` | Select the architecture |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference |
| `--download-directory` | `-d` | Directory where the installers are downloaded to |
| `--exact` | `-e` | Find package using exact match |
| `--header` | | REST source HTTP header |
| `--id` | | Filter results by id |
| `--ignore-security-hash` | | Ignore installer hash check failure |
| `--installer-type` | | Select the installer type |
| `--locale` | | Locale to use (BCP47 format) |
| `--manifest` | `-m` | The path to the manifest of the package |
| `--moniker` | | Filter results by moniker |
| `--name` | | Filter results by name |
| `--os-version` | | Target OS version |
| `--platform` | | Select the target platform |
| `--query` | `-q` | The query used to search |
| `--scope` | | Select install scope (`user`, `machine`) |
| `--skip-dependencies` | | Skip processing package dependencies |
| `--skip-license` | | Skip retrieving Microsoft Store offline license |
| `--skip-microsoft-store-package-license` | | Skip retrieving MS Store offline license |
| `--source` | `-s` | Find package using the specified source |
| `--version` | `-v` | Use the specified version; default is the latest |

### Download Directory

The `--download-directory` (`-d`) flag specifies where installers are downloaded. If not specified, the `downloadBehavior.defaultDownloadDirectory` setting is used (default: `%USERPROFILE%/Downloads/`). See [settings.md](settings.md).

## `winget hash`

Helper to hash installer files. Generates the SHA256 hash needed for manifest `InstallerSha256` fields.

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--file` | `-f` | File to be hashed |
| `--msix` | `-m` | Input file will be treated as MSIX; signature hash will be provided if signed |

### Usage

```
winget hash --file .\installer.exe
winget hash --file .\package.msix --msix
```

For MSIX files, the `--msix` flag outputs both the package hash (`InstallerSha256`) and the signature hash (`SignatureSha256`), which is needed if the manifest includes MSIX signature validation.

## `winget validate`

Validates a manifest file against the winget manifest schema.

### Flags

| Flag | Description |
|------|-------------|
| `--manifest` | The path to the manifest to be validated |

### Usage

```
winget validate --manifest .\path\to\manifest.yaml
```

Checks YAML syntax, schema compliance, required fields, and field constraints (max lengths, valid enum values, etc.). See [manifest.md](manifest.md) for the full schema specification.

## `winget features`

Shows the status of experimental features. No flags.

```
winget features
```

Displays a table of all experimental features and whether each is enabled or disabled. Features are controlled via the `experimentalFeatures` section in settings.json. See [settings.md](settings.md).

## `winget mcp`

MCP (Model Context Protocol) information. Manages extended MCP features.

| Flag | Description |
|------|-------------|
| `--enable` | Enable extended features. Requires store access. |
| `--disable` | Disable extended features. Requires store access. |

## Export/Import Workflow

The export/import workflow is useful for:

1. **Environment replication** — Export from one machine, import on another
2. **Backup** — Keep a record of installed packages
3. **CI/CD** — Ensure consistent package sets across build agents

```
# On source machine
winget export -o packages.json --include-versions

# On target machine
winget import -i packages.json --accept-package-agreements --accept-source-agreements
```

## Related

- For install/upgrade/uninstall/repair commands, see [install-uninstall.md](install-uninstall.md).
- For the manifest schema (what `validate` checks against), see [manifest.md](manifest.md).
- For settings that control download directory and other behaviors, see [settings.md](settings.md).
- For command structure and global flags, see [cli.md](cli.md).
