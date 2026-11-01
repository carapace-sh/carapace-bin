# DNF5 Execution Workflow

The 24-step lifecycle of a `dnf5` command execution, including all plugin hooks and command-specific lifecycle methods.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/dnf5_workflow.html>. For plugin interfaces, see [plugins.md](plugins.md).

## Overview

Every `dnf5` invocation follows a strict sequence of 24 steps. DNF5 plugins (active) and libdnf5 plugins (passive) hook into this lifecycle at specific points. Command implementations also have lifecycle methods (`pre_configure`, `configure`, `run`, etc.) that are called at specific stages.

## The 24-Step Lifecycle

### 1. Set up loggers

Initializes the logging system in the `Base` object.

### 2. Create base

Creates the `libdnf5::Base` object, the central context for all operations.

### 3. Add dnf5 commands

Registers all built-in command implementations with the argument parser. Each command goes through:
- `set_parent_command` — attaches to parent command
- `set_argument_parser` — defines its arguments
- `register_subcommands` — registers any subcommands

### 4. Load dnf5 plugins

Loads `.so` files from the dnf5 plugins directory (overridable via `DNF5_PLUGINS_DIR`). For each plugin:

1. **`dnf5 plugin init` hook** — called after the `Base` object is created
2. **`dnf5 plugin create_commands` hook** — plugin registers its commands
3. For plugin commands: `set_parent_command`, `set_argument_parser`, `register_subcommands`

For native dnf5 commands, the three sub-steps above were done as part of step 3.

### 5. Load aliases

Loads command-line aliases from TOML files in the alias drop-in directories. See [aliases.md](aliases.md).

### 6. Parse command line arguments

The argument parser processes `argv`, resolving commands, options, and positional arguments.

### 7. Run command-specific `pre_configure` step

Each command's `pre_configure()` method is called. This runs before the main configuration is loaded, allowing commands to adjust how configuration is loaded.

### 8. Load main configuration

Reads `/etc/dnf/dnf.conf` and drop-in configuration files. See [configuration.md](configuration.md) for the drop-in directory processing order.

### 9. Enable/disable libdnf5 plugins

Processes `--enable-plugin` and `--disable-plugin` CLI options. The `plugins` config option and `--no-plugins` flag are also applied here.

### 10. Base setup

Initializes the `Base` object with loaded configuration:

1. Load libdnf5 plugins from `pluginpath` directories
2. **`libdnf5 plugin init` hook**
3. **`libdnf5 plugin pre_base_setup` hook**
4. Lock installroot
5. Load Vars and lock varsdir
6. **`libdnf5 plugin post_base_setup` hook**

### 11. Create repo sack

Creates the `libdnf5::repo::RepoSack` object that manages all repositories.

### 12. Create repos from system configuration

Loads `.repo` files from `reposdir` directories. Performs Vars substitution (`$arch`, `$basearch`, `$releasever`, user variables) on repository id and all configuration values.

### 13. Create repos from paths

Processes `--repofrompath=REPO_ID,REPO_PATH` arguments. Performs Vars substitution on the specified id and path.

### 14. Apply repository setopts

Processes `--setopt=REPO_ID.OPTION=VALUE` arguments to override per-repo configuration.

### 15. Run command-specific `configure` step

Each command's `configure()` method is called. Commands can adjust repository loading behavior, set config options, etc.

### 16. `libdnf5 plugin repos_configured` hook

All repositories have been configured (from files, setopts, and command configuration). Plugins can inspect and modify the final repo configuration.

### 17. Check for privileges

If the command requires privileges (most transactional commands do), DNF5 checks for them. This is where `sudo` or polkit authorization would be relevant.

### 18. Load repositories

1. If required, load system repository (installed packages database / RPMDB)
2. If required, load enabled repositories:
   1. Load metadata from cache if valid
   2. Try to reuse root's cache (copy to user cache if running as non-root)
   3. Metadata download (if cache expired or `--refresh`)
   4. Metadata OpenPGP check (if `repo_gpgcheck` enabled)
   5. If required, import repository OpenPGP keys and try again

### 19. `libdnf5 plugin repos_loaded` hook

All repository metadata has been loaded. Plugins can now query packages across all available repositories.

### 20. Run command-specific `load_additional_packages` step

Commands can add packages from non-repository sources (e.g., command-line RPM files).

### 21. Run command-specific `run` step

The main logic of the command executes here. For transactional commands, this creates a `libdnf5::base::Goal`, adds install/upgrade/remove/etc. requests to it, and resolves the transaction.

### 22. Transaction execution (if goal produced)

If the command produced a goal:

1. **`libdnf5 plugin pre_add_cmdline_packages` hook**
2. Add commandline packages to the goal
3. **`libdnf5 plugin post_add_cmdline_packages` hook**
4. Resolve goal (resolve dependencies via libsolv)
5. **`libdnf5 plugin goal_resolved` hook** (API 2.1+)
6. Run command-specific `goal_resolved` step
7. Print transaction table
8. Check for user approval
9. Download inbound transaction packages
10. Check OpenPGP signatures for inbound transaction packages
11. Lock transaction (`libdnf5::utils::Locker`)
12. Create RPM transaction
13. Run RPM test transaction
14. **`libdnf5 plugin pre_transaction` hook**
15. Start database transaction
16. Run RPM transaction
17. Update system_state
18. Finish database transaction
19. **`libdnf5 plugin post_transaction` hook**
20. Unlock transaction

### 23. `libdnf5 plugin finish` hook

All libdnf5 plugins receive the finish callback. Cleanup of resources.

### 24. `dnf5 plugin finish` hook

All dnf5 plugins receive the finish callback. Final cleanup before exit.

## Summary of All Hooks

| Step | Hook Name | Type |
|------|-----------|------|
| 4.1 | `dnf5 plugin init` | dnf5 plugin (active) |
| 4.2 | `dnf5 plugin create_commands` | dnf5 plugin (active) |
| 10.2 | `libdnf5 plugin init` | libdnf5 plugin (passive) |
| 10.3 | `libdnf5 plugin pre_base_setup` | libdnf5 plugin (passive) |
| 10.6 | `libdnf5 plugin post_base_setup` | libdnf5 plugin (passive) |
| 16 | `libdnf5 plugin repos_configured` | libdnf5 plugin (passive) |
| 19 | `libdnf5 plugin repos_loaded` | libdnf5 plugin (passive) |
| 22.1 | `libdnf5 plugin pre_add_cmdline_packages` | libdnf5 plugin (passive) |
| 22.3 | `libdnf5 plugin post_add_cmdline_packages` | libdnf5 plugin (passive) |
| 22.5 | `libdnf5 plugin goal_resolved` | libdnf5 plugin (passive, API 2.1+) |
| 22.14 | `libdnf5 plugin pre_transaction` | libdnf5 plugin (passive) |
| 22.19 | `libdnf5 plugin post_transaction` | libdnf5 plugin (passive) |
| 23 | `libdnf5 plugin finish` | libdnf5 plugin (passive) |
| 24 | `dnf5 plugin finish` | dnf5 plugin (active) |

## Command-Specific Lifecycle Methods

These are methods on the `dnf5::Command` base class that command implementations override:

| Method | Step | Purpose |
|--------|------|---------|
| `set_parent_command` | 3/4 | Attach to parent command in the tree |
| `set_argument_parser` | 3/4 | Define arguments, options, positional specs |
| `register_subcommands` | 3/4 | Register any subcommands |
| `pre_configure` | 7 | Adjust config loading (runs before config is loaded) |
| `configure` | 15 | Adjust repo loading, set config options |
| `load_additional_packages` | 20 | Add non-repository packages |
| `run` | 21 | Main command logic, create and resolve goal |
| `goal_resolved` | 22.6 | Post-resolution processing (e.g., filter results) |
