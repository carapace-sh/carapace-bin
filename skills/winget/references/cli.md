# Command Structure and Global Flags

The complete command tree for winget — all commands, subcommands, aliases, and the global flags that apply across the CLI.

> **Source of truth**: `winget --help`, `winget <command> --help`, <https://learn.microsoft.com/en-us/windows/package-manager/winget/>. For **core concepts**, see [concepts.md](concepts.md). For **install/upgrade/uninstall flags**, see [install-uninstall.md](install-uninstall.md). For **source commands**, see [sources.md](sources.md).

## Global Flags (Persistent)

These flags are available on all commands via persistent flag inheritance:

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--disable-interactivity` | | bool | Disable interactive prompts |
| `--help` | `-?` | bool | Shows help about the selected command |
| `--ignore-warnings` | | bool | Suppresses warning outputs |
| `--logs` | | bool | Open the default logs location |
| `--no-proxy` | | bool | Disable the use of proxy for this execution |
| `--nowarn` | | bool | Suppresses warning outputs (alias for `--ignore-warnings`) |
| `--open-logs` | | bool | Open the default logs location (alias for `--logs`) |
| `--proxy` | | string | Set a proxy to use for this execution |
| `--verbose` | | bool | Enables verbose logging for winget |
| `--verbose-logs` | | bool | Enables verbose logging for winget |
| `--wait` | | bool | Prompts the user to press any key before exiting |

## Root-Only Flags

These flags are only available on the root `winget` command (not inherited by subcommands):

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--info` | | bool | Display general info of the tool |
| `--version` | `-v` | bool | Display the version of the tool |

## Command Tree

### Package Operations

#### `winget install` (alias: `add`)

Installs the selected package. Accepts a package query as a positional argument or via filter flags.

Key flags (full list in [install-uninstall.md](install-uninstall.md)):

| Flag | Short | Description |
|------|-------|-------------|
| `--id` | | Filter results by id |
| `--name` | | Filter results by name |
| `--moniker` | | Filter results by moniker |
| `--query` | `-q` | The query used to search for a package |
| `--source` | `-s` | Find package using the specified source |
| `--exact` | `-e` | Find package using exact match |
| `--version` | `-v` | Use the specified version; default is the latest |
| `--manifest` | `-m` | The path to the manifest of the package |
| `--silent` | `-h` | Request silent installation |
| `--interactive` | `-i` | Request interactive installation |
| `--scope` | | Select install scope (`user`, `machine`) |
| `--architecture` | `-a` | Select the architecture |
| `--location` | `-l` | Location to install to (if supported) |
| `--log` | `-o` | Log location (if supported) |
| `--force` | | Direct run, continue with non-security issues |
| `--no-upgrade` | | Skip upgrade if installed version exists |
| `--override` | | Override arguments passed to the installer |
| `--custom` | | Additional arguments passed to the installer |
| `--accept-package-agreements` | | Accept all license agreements |
| `--accept-source-agreements` | | Accept all source agreements |

#### `winget upgrade` (alias: `update`)

Shows and performs available upgrades.

| Flag | Short | Description |
|------|-------|-------------|
| `--all` | | Upgrade all installed packages (also `--recurse`) |
| `--recurse` | `-r` | Upgrade all installed packages to latest |
| `--include-unknown` | | Upgrade even if current version unknown (also `--unknown`) |
| `--unknown` | `-u` | Upgrade even if current version unknown |
| `--include-pinned` | | Upgrade even if pinned (also `--pinned`) |
| `--pinned` | | Upgrade even if pinned |
| `--uninstall-previous` | | Uninstall previous version during upgrade |

See [install-uninstall.md](install-uninstall.md) for the full flag list.

#### `winget uninstall`

Uninstalls the given package.

| Flag | Short | Description |
|------|-------|-------------|
| `--all` | | Uninstall all versions |
| `--all-versions` | | Uninstall all versions |
| `--purge` | | Delete all files/dirs in package directory (portable) |
| `--preserve` | | Retain all files/dirs created by package (portable) |
| `--product-code` | | Filter using the product code |
| `--version` | `-v` | The version to act upon |

See [install-uninstall.md](install-uninstall.md) for the full flag list.

#### `winget repair`

Repairs the selected package. Uses the same authentication/filter flags as install.

| Flag | Short | Description |
|------|-------|-------------|
| `--exact` | `-e` | Find package using exact match |
| `--force` | | Direct run, continue with non-security issues |
| `--ignore-security-hash` | | Ignore installer hash check failure |
| `--ignore-local-archive-malware-scan` | | Ignore malware scan for local archive |
| `--locale` | | Locale to use (BCP47 format) |
| `--log` | `-o` | Log location (if supported) |

### Discovery Commands

#### `winget search` (alias: `find`)

Find and show basic info of packages.

| Flag | Short | Description |
|------|-------|-------------|
| `--id` | | Filter results by id |
| `--name` | | Filter results by name |
| `--moniker` | | Filter results by moniker |
| `--query` | `-q` | The query used to search |
| `--source` | `-s` | Find package using the specified source |
| `--exact` | `-e` | Find package using exact match |
| `--count` | `-n` | Show no more than specified number of results |
| `--tag` | | Filter results by tag |
| `--cmd` | | Filter results by command |
| `--command` | | Filter results by command |
| `--versions` | | Show available versions of the package |

#### `winget list` (alias: `ls`)

Display installed packages.

| Flag | Short | Description |
|------|-------|-------------|
| `--upgrade-available` | | List only packages with an upgrade available |
| `--include-pinned` | | List packages even if pinned (with `--upgrade-available`) |
| `--pinned` | | List packages even if pinned (with `--upgrade-available`) |
| `--include-unknown` | | List packages even if version unknown (with `--upgrade-available`) |
| `--unknown` | `-u` | List packages even if version unknown |
| `--count` | `-n` | Show no more than specified number of results (1–1000) |
| `--tag` | | Filter results by tag |
| `--cmd` | | Filter results by command |
| `--command` | | Filter results by command |
| `--scope` | | Select installed package scope filter |

See [package-management.md](package-management.md) for full details.

#### `winget show` (alias: `view`)

Displays details for a specified package — full manifest information including installer URLs, SHA256 hashes, description, license, and available versions.

| Flag | Short | Description |
|------|-------|-------------|
| `--query` | `-q` | The query used to search for a package |
| `--id` | | Filter results by id |
| `--name` | | Filter results by name |
| `--moniker` | | Filter results by moniker |
| `--source` | `-s` | Find package using the specified source |
| `--exact` | `-e` | Find package using exact match |
| `--version` | `-v` | Use the specified version; default is the latest |
| `--scope` | | Select install scope (`user`, `machine`) |
| `--architecture` | `-a` | Select the architecture to show |
| `--installer-type` | | Select the installer type to show |
| `--locale` | | Locale to use (BCP47 format) |
| `--manifest` | `-m` | The path to the manifest of the package |
| `--versions` | | Show available versions of the package |
| `--header` | | Optional REST source HTTP header |
| `--authentication-mode` | | Auth window preference (`silent`, `silentPreferred`, `interactive`) |
| `--authentication-account` | | Account for authentication |
| `--accept-source-agreements` | | Accept all source agreements |

See [package-management.md](package-management.md) for full details.

### Source Management

#### `winget source`

Manage sources of packages. Subcommands:

| Subcommand | Aliases | Description |
|------------|---------|-------------|
| `add` | | Add a new source |
| `list` | `ls` | List current sources |
| `remove` | `rm` | Remove current sources |
| `update` | `refresh` | Update current sources |
| `export` | | Export current sources |
| `reset` | | Reset sources |

See [sources.md](sources.md) for full details and flags.

### Configuration

#### `winget configure` (aliases: `configuration`, `dsc`)

Configures the system into a desired state using DSC.

| Flag | Short | Description |
|------|-------|-------------|
| `--file` | `-f` | The path to the configuration file |
| `--module-path` | | Location on local computer to store modules |
| `--processor-path` | | Path to the configuration processor |

Subcommands:

| Subcommand | Description |
|------------|-------------|
| `export` | Exports configuration resources to a configuration file |
| `list` (alias `ls`) | Shows configuration history |
| `show` (alias `view`) | Shows details of a configuration |
| `test` | Checks the system against a desired state |
| `validate` | Validates a configuration file |

See [configure.md](configure.md) for full details.

#### `winget dscv3`

DSC v3 resource commands. Exposes DSC v3 resources for direct management.

| Flag | Short | Description |
|------|-------|-------------|
| `--manifest` | | Get the resource manifest |
| `--output` | `-o` | Directory where results are written |

Subcommands (all share the same operation flags):

| Subcommand | Description |
|------------|-------------|
| `admin-settings` | Manage administrator settings |
| `package` | Manage package state |
| `source` | Manage source configuration |
| `user-settings-file` | Manage user settings file |

Each subcommand supports: `--export`, `--get`, `--schema`, `--set`, `--test`.

See [configure.md](configure.md) for full details.

### Settings

#### `winget settings` (alias: `config`)

Open settings or set administrator settings.

| Subcommand | Description |
|------------|-------------|
| `export` | Export settings |
| `set` | Sets the value of an admin setting |
| `reset` | Resets an admin setting to its default value |

See [settings.md](settings.md) for full details.

### Package Management Utilities

#### `winget pin`

Manage package pins. Subcommands:

| Subcommand | Description |
|------------|-------------|
| `add` | Add a new pin |
| `list` | List current pins |
| `remove` | Remove a package pin |
| `reset` | Reset pins |

See [package-management.md](package-management.md) for full details.

#### `winget export`

Exports a list of the installed packages.

| Flag | Short | Description |
|------|-------|-------------|
| `--output` | `-o` | File where the result is to be written |
| `--source` | `-s` | Export packages from the specified source |
| `--include-versions` | | Include package versions in export file |

#### `winget import`

Installs all the packages in a file.

| Flag | Short | Description |
|------|-------|-------------|
| `--import-file` | `-i` | File describing the packages to install |
| `--ignore-unavailable` | | Ignore unavailable packages |
| `--ignore-versions` | | Ignore package versions in import file |
| `--no-upgrade` | | Skip upgrade if installed version exists |

#### `winget download`

Downloads the installer from a given package.

| Flag | Short | Description |
|------|-------|-------------|
| `--download-directory` | `-d` | Directory where installers are downloaded |
| `--skip-license` | | Skip retrieving MS Store offline license |
| `--skip-microsoft-store-package-license` | | Skip retrieving MS Store offline license |
| `--os-version` | | Target OS version |
| `--platform` | | Select the target platform |

See [package-management.md](package-management.md) for the full flag list.

### Developer Tools

#### `winget hash`

Helper to hash installer files.

| Flag | Short | Description |
|------|-------|-------------|
| `--file` | `-f` | File to be hashed |
| `--msix` | `-m` | Treat input as MSIX; signature hash provided if signed |

#### `winget validate`

Validates a manifest file.

| Flag | Description |
|------|-------------|
| `--manifest` | The path to the manifest to be validated |

#### `winget features`

Shows the status of experimental features. No flags.

#### `winget mcp`

MCP information. Manages extended MCP features.

| Flag | Description |
|------|-------------|
| `--enable` | Enable extended features. Requires store access. |
| `--disable` | Disable extended features. Requires store access. |

## Common Filter Flags

Several commands share these package filter flags:

| Flag | Short | Description |
|------|-------|-------------|
| `--id` | | Filter results by id |
| `--name` | | Filter results by name |
| `--moniker` | | Filter results by moniker |
| `--query` | `-q` | The query used to search for a package |
| `--source` | `-s` | Find package using the specified source |
| `--exact` | `-e` | Find package using exact match |
| `--tag` | | Filter results by tag |
| `--cmd` | | Filter results by command |
| `--command` | | Filter results by command |
| `--header` | | Optional REST source HTTP header |
| `--authentication-account` | | Account for authentication |
| `--authentication-mode` | | Auth window preference (`silent`, `silentPreferred`, `interactive`) |
| `--accept-source-agreements` | | Accept all source agreements |

## Related

- For install/upgrade/uninstall/repair workflows and all flags, see [install-uninstall.md](install-uninstall.md).
- For source commands and source types, see [sources.md](sources.md).
- For configuration and DSC v3, see [configure.md](configure.md).
- For settings.json structure, see [settings.md](settings.md).
