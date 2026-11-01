# DNF5 Plugin System

DNF5 has a two-tier plugin architecture: **active plugins** (dnf5 command plugins that add new CLI commands) and **passive plugins** (libdnf5 hook plugins that hook into the library lifecycle).

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/tutorial/plugins/index.html>. For the execution lifecycle where hooks fire, see [workflow.md](workflow.md).

## Two-Tier Architecture

```
┌─────────────────────────────────────────────────┐
│                 dnf5 CLI (Application)          │
│  ┌───────────────────────────────────────────┐  │
│  │  DNF5 Plugins (Active / Command Plugins)  │  │
│  │  Interface: dnf5::IPlugin                 │  │
│  │  Adds: new commands (dnf5 <command>)      │  │
│  │  Loaded from: dnf5/plugins/ (via DNF5_PLUGINS_DIR)│
│  └───────────────────────────────────────────┘  │
│                      │                           │
│                      ▼                           │
│  ┌───────────────────────────────────────────┐  │
│  │  libdnf5 (Core Library)                   │  │
│  │  ┌─────────────────────────────────────┐  │  │
│  │  │  libdnf5 Plugins (Passive / Hook)   │  │  │
│  │  │  Interface: libdnf5::plugin::IPlugin│  │  │
│  │  │  Adds: hooks into lifecycle events  │  │  │
│  │  │  Loaded from: libdnf5/plugins/      │  │  │
│  │  └─────────────────────────────────────┘  │  │
│  └───────────────────────────────────────────┘  │
└─────────────────────────────────────────────────┘
```

## DNF5 Active Plugins (Command Plugins)

### Overview

DNF5 active plugins are shared libraries (`.so`) that add new commands to the `dnf5` CLI. Users type `dnf5 <plugin-command>` to access them. Core commands live in `dnf5/commands/`; plugin commands live in `dnf5-plugins/`.

### The `dnf5::IPlugin` Interface

Defined in `<dnf5/iplugin.hpp>`. A plugin implements:

**Mandatory methods:**
- `get_api_version()` — Returns the plugin API version required.
- `get_name()` — Plugin name (used in log messages).
- `get_version()` — Plugin's own version.
- `create_commands()` — Returns `std::vector<std::unique_ptr<Command>>` of all commands.

**Optional methods:**
- `get_attributes()` / `get_attribute()` — Custom attributes (author, description).
- `init()` — Called after `Base` creation, before argument parsing.
- `finish()` — Called when plugin objects are garbage collected.

### C Linkage Entry Points (Mandatory)

Every DNF5 command plugin must export these C symbols:

| Function | Description |
|----------|-------------|
| `dnf5_plugin_get_api_version()` | Returns API version implemented by plugin |
| `dnf5_plugin_get_name()` | Internal name of the plugin |
| `dnf5_plugin_get_version()` | Plugin's own version |
| `dnf5_plugin_new_instance()` | Factory: creates `IPlugin` object |
| `dnf5_plugin_delete_instance()` | Cleanup: destroys plugin instance |
| `dnf5_plugin_get_last_exception()` | Returns pointer to last caught exception |

### Code Example

```cpp
#include "template_cmd.hpp"
#include <dnf5/iplugin.hpp>

using namespace dnf5;

namespace {

constexpr const char * PLUGIN_NAME{"dnf5_template_plugin"};
constexpr PluginVersion PLUGIN_VERSION{.major = 1, .minor = 0, .micro = 0};
constexpr PluginAPIVersion REQUIRED_PLUGIN_API_VERSION{.major = 2, .minor = 0};

class TemplateCmdPlugin : public IPlugin {
public:
    using IPlugin::IPlugin;

    PluginAPIVersion get_api_version() const noexcept override { return REQUIRED_PLUGIN_API_VERSION; }
    const char * get_name() const noexcept override { return PLUGIN_NAME; }
    PluginVersion get_version() const noexcept override { return PLUGIN_VERSION; }

    std::vector<std::unique_ptr<Command>> create_commands() override;

    void init() override {}
    void finish() noexcept override {}
};

std::vector<std::unique_ptr<Command>> TemplateCmdPlugin::create_commands() {
    std::vector<std::unique_ptr<Command>> commands;
    commands.push_back(std::make_unique<TemplateCommand>(get_context()));
    return commands;
}

}  // namespace

// C linkage - mandatory exports
PluginAPIVersion dnf5_plugin_get_api_version(void) { return REQUIRED_PLUGIN_API_VERSION; }
const char * dnf5_plugin_get_name(void) { return PLUGIN_NAME; }
PluginVersion dnf5_plugin_get_version(void) { return PLUGIN_VERSION; }
IPlugin * dnf5_plugin_new_instance(ApplicationVersion, Context & context) try {
    return new TemplateCmdPlugin(context);
} catch (...) { return nullptr; }
void dnf5_plugin_delete_instance(IPlugin * plugin_object) { delete plugin_object; }
```

### Plugin Directory Structure

```
dnf5-plugins/builddep_plugin/
├── builddep_cmd_plugin.cpp   # IPlugin implementation + C exports
├── builddep.cpp              # Command implementation
├── builddep.hpp               # Command header
└── CMakeLists.txt             # Build script
```

### Packaging (CMake)

```cmake
add_library(template_dnf5_cmd_plugin MODULE template_cmd.cpp template_cmd_plugin.cpp)
set_target_properties(template_dnf5_cmd_plugin PROPERTIES PREFIX "")  # no 'lib' prefix
target_link_libraries(template_dnf5_cmd_plugin PRIVATE dnf5)
# install(TARGETS template_dnf5_cmd_plugin LIBRARY DESTINATION ${CMAKE_INSTALL_FULL_LIBDIR}/dnf5/plugins/)
```

RPM spec provides:
```
Provides: dnf5-command(template)
```

## libdnf5 Passive Plugins (Hook Plugins)

### Overview

libdnf5 plugins hook into the core library's lifecycle events. They are executed regardless of how `libdnf5` is used — through the `dnf5` CLI or via API calls. They can modify behavior, implement logic, or trigger loading of additional plugins.

### The `libdnf5::plugin::IPlugin` Interface

Defined in `<libdnf5/plugin/iplugin.hpp>`. Two base classes:

- **`IPlugin`** — Base class with hooks up to API version 2.0.
- **`IPlugin2_1`** — Extends `IPlugin` with `goal_resolved()` hook (API 2.1+).

**Mandatory methods:**
- `get_api_version()` — Returns required API version.
- `get_name()` — Plugin name.
- `get_version()` — Plugin version.

**Hook methods (override as needed):**
- `init()` — After plugin is loaded
- `pre_base_setup()` — Before Base setup
- `post_base_setup()` — After Base setup
- `repos_configured()` — After repos are configured
- `repos_loaded()` — After repos are loaded
- `pre_add_cmdline_packages()` — Before adding cmdline packages to goal
- `post_add_cmdline_packages()` — After adding cmdline packages
- `goal_resolved()` — After goal resolution (API 2.1+)
- `pre_transaction()` — Before RPM transaction
- `post_transaction()` — After RPM transaction
- `finish()` — Plugin cleanup

### C Linkage Entry Points (Mandatory)

| Function | Description |
|----------|-------------|
| `libdnf_plugin_get_api_version()` | Returns API version |
| `libdnf_plugin_get_name()` | Plugin name |
| `libdnf_plugin_get_version()` | Plugin version |
| `libdnf_plugin_new_instance()` | Factory: receives `LibraryVersion`, `IPluginData`, `ConfigParser` |
| `libdnf_plugin_delete_instance()` | Cleanup |
| `libdnf_plugin_get_last_exception()` | Returns pointer to last exception |

### Configuration File Requirement

Each libdnf5 plugin **requires** a configuration file:

```ini
[main]
name = template_plugin
enabled = yes

# Optional custom keys
custom_key = some_value
```

The `enabled` option supports:
- `no` — Plugin is disabled.
- `yes` — Plugin is enabled.
- `host-only` — Enabled only without installroot.
- `installroot-only` — Enabled only with installroot.

### Code Example

```cpp
#include <libdnf5/base/base.hpp>
#include <libdnf5/common/exception.hpp>
#include <libdnf5/plugin/iplugin.hpp>

using namespace libdnf5;

namespace {

constexpr const char * PLUGIN_NAME{"libdnf5_template_plugin"};
constexpr plugin::Version PLUGIN_VERSION{.major = 1, .minor = 1, .micro = 0};
constexpr PluginAPIVersion REQUIRED_PLUGIN_API_VERSION{.major = 2, .minor = 0};

class TemplatePlugin final : public plugin::IPlugin2_1 {
public:
    TemplatePlugin(libdnf5::plugin::IPluginData & data, libdnf5::ConfigParser & parser)
        : IPlugin2_1(data) {}

    PluginAPIVersion get_api_version() const noexcept override { return REQUIRED_PLUGIN_API_VERSION; }
    const char * get_name() const noexcept override { return PLUGIN_NAME; }
    plugin::Version get_version() const noexcept override { return PLUGIN_VERSION; }

    void init() override {}
    void finish() noexcept override {}

    void post_base_setup() override { /* hook logic */ }
    void pre_transaction(const libdnf5::base::Transaction & transaction) override {
        auto & base = get_base();
        auto & logger = *base.get_logger();
        logger.info("{} packages in transaction", transaction.get_transaction_packages_count());
    }
};

}  // namespace

PluginAPIVersion libdnf_plugin_get_api_version(void) { return REQUIRED_PLUGIN_API_VERSION; }
const char * libdnf_plugin_get_name(void) { return PLUGIN_NAME; }
plugin::Version libdnf_plugin_get_version(void) { return PLUGIN_VERSION; }
plugin::IPlugin * libdnf_plugin_new_instance(LibraryVersion, plugin::IPluginData & data, ConfigParser & parser) try {
    return new TemplatePlugin(data, parser);
} catch (...) { return nullptr; }
void libdnf_plugin_delete_instance(plugin::IPlugin * plugin_object) { delete plugin_object; }
```

## Plugin API Versioning

### Version Structures

```cpp
struct PluginAPIVersion {
    std::uint16_t major;  // must exactly match between plugin and host
    std::uint16_t minor;  // plugin's minor must be <= host's minor
};

struct PluginVersion {
    std::uint16_t major;
    std::uint16_t minor;
    std::uint16_t micro;
};

struct LibraryVersion {
    std::uint16_t prime;   // stays as 5
    std::uint16_t major;   // incompatible API changes
    std::uint16_t minor;   // backward-compatible additions
    std::uint16_t micro;   // bug fixes
};
```

### Compatibility Rules

1. **Major version match**: Plugin's `major` must exactly match the host's `major`.
2. **Minor version**: Plugin's required `minor` must be <= the version provided by the host. A plugin compiled against API 2.0 works with library at 2.2.

### API Version History

| API Version | Changes |
|-------------|---------|
| 2.0 | Base `IPlugin` with hooks: `init`, `pre_base_setup`, `post_base_setup`, `repos_configured`, `repos_loaded`, `pre_add_cmdline_packages`, `post_add_cmdline_packages`, `pre_transaction`, `post_transaction`, `finish` |
| 2.1 | Added `IPlugin2_1` with `goal_resolved()` hook |
| 2.2 | Current library version |

## Plugin Loading Process

### DNF5 Command Plugins

1. `dnf5::Plugins` scans the plugins directory for `.so` files (alphabetical order).
2. Loads each shared library and calls `dnf5_plugin_get_api_version()` for version check.
3. If compatible, calls `dnf5_plugin_new_instance()` to create the plugin instance.
4. Calls `create_commands()` to register commands with the argument parser.

### libdnf5 Hook Plugins

1. `libdnf5::plugin::Plugins` searches for `.so` files in `pluginpath` directories.
2. Reads configuration file from `pluginconfpath` for each plugin.
3. Checks if plugin is enabled via the `enabled` option.
4. Loads the shared library and calls `libdnf_plugin_get_api_version()`.
5. Calls `libdnf_plugin_new_instance()` passing `LibraryVersion`, `IPluginData`, and `ConfigParser`.
6. Registers the plugin with the `libdnf5::Base` object.

## Built-in libdnf5 Plugins

| Plugin | Hook(s) | Purpose |
|--------|---------|---------|
| **actions** | All hooks | Execute external commands/shell scripts at lifecycle hook points. Configured via `.actions` files. Supports variable substitution (`${pkg.nevra}`, `${conf.installroot}`). Communication modes: plain (stdout) and JSON (bidirectional). |
| **python_plugins_loader** | `load_plugins()` | Loads Python-based plugins using SWIG bindings. Maintains a ref counter for Python interpreter lifecycle. |
| **expired-pgp-keys** | `goal_resolved()` | Detects and removes expired PGP signing keys from RPMDB. Runs `gpg --show-keys --with-colon` to parse expiration. |
| **rhsm** | `post_base_setup()` | Synchronizes Red Hat Subscription Manager repo configs. Generates `redhat.repo` from enrollment context. |
| **appstream** | `repos_loaded()` | Installs AppStream metadata to system cache. |
| **local** | `post_base_setup()`, `post_transaction()` | Maintains a local repository of downloaded packages. Creates `_dnf_local` and `_dnf_local_nogpgcheck` repos. Copies RPM files post-transaction. |

## Built-in DNF5 Command Plugins

Provided by the `dnf5-plugins` package:

| Plugin | Purpose |
|--------|---------|
| **automatic** | Automated updates with notification emitters (email, motd, command) |
| **builddep** | Install build requirements for spec files or SRPMs |
| **changelog** | Display package changelogs |
| **config-manager** | Manage repo config (addrepo, setopt, setvar, unsetopt, unsetvar) |
| **copr** | Manage COPR repositories (enable, disable, list, remove) |
| **manifest** | Manage pinned package sets using `libpkgmanifest` |
| **needs-restarting** | Detect if system/services need restart after updates |
| **repoclosure** | Check for unresolved dependencies in repos |
| **repomanage** | Manage older/newer packages in a repo directory |
| **reposync** | Synchronize remote repos to a local directory |

## Environment Variables

| Variable | Applies To | Description |
|----------|------------|-------------|
| `DNF5_PLUGINS_DIR` | DNF5 command plugins | Override directory for loading command plugins. Empty string disables. |
| `LIBDNF_PLUGINS_CONFIG_DIR` | libdnf5 plugins | Override directory for plugin config files. |
| `LIBDNF_PYTHON_PLUGIN_DIR` | Python plugin loader | Override directory for Python plugin configs. |

### Configuration Options Affecting Plugins

| Option | Default | Description |
|--------|---------|-------------|
| `pluginpath` | `/usr/lib64/libdnf5/plugins/` | Directories for libdnf5 plugin `.so` files |
| `pluginconfpath` | `/etc/dnf/libdnf5-plugins` | Directories for plugin config files |
| `plugins` | `True` | Master switch for libdnf5 plugins |

## Debugging Tips

### DNF5 Command Plugins

```bash
# Load plugin from build directory
DNF5_PLUGINS_DIR=~/dnf5/build/dnf5-plugins/template_plugin dnf5 <command>
```

### libdnf5 Plugins

```bash
# Override plugin config directory
LIBDNF_PLUGINS_CONFIG_DIR=libdnf5_baz_plugin dnf5 ...

# Override plugin binary path
LIBDNF_PLUGINS_CONFIG_DIR=libdnf5_baz_plugin dnf5 reinstall libsolv \
  --setopt=pluginpath=build/libdnf5_baz_plugin --assumeno
```

### General

- Build with `-DCMAKE_BUILD_TYPE=Debug` for debugging symbols.
- Enable `debug_solver` in config for libsolv debug files.
- Plugin name and version are used in informative log messages.

## Packaging Steps

### For a DNF5 Command Plugin

1. Write source files (`*_plugin.cpp` with IPlugin + C exports, `*.cpp`/`*.hpp` for command)
2. Write `CMakeLists.txt` with `add_library(... MODULE ...)`, set `PREFIX ""`, link `dnf5`
3. Add `add_subdirectory("plugin_name")` to parent `dnf5-plugins/CMakeLists.txt`
4. Add RPM spec section with `Provides: dnf5-command(name)`
5. Add documentation: `doc/dnf5_plugins/name.8.rst`, update `doc/dnf5_plugins/index.rst`, `doc/dnf5.8.rst`

### For a libdnf5 Plugin

1. Write source file (`*.cpp` with IPlugin + C exports)
2. Write config file (`name.conf` with `[main]` section, `name` and `enabled`)
3. Write `CMakeLists.txt` with `add_library(... MODULE ...)`, set `PREFIX ""`, link `libdnf5`
4. Add `option(WITH_PLUGIN_NAME "Build plugin" ON)` to `dnf5/CMakeLists.txt`
5. Add `add_subdirectory("name")` to `libdnf5-plugins/CMakeLists.txt`
6. Add RPM spec section: `libdnf5-plugin-name` package, install `.so` and `.conf`
7. Add documentation: `doc/libdnf5_plugins/name.8.rst`, update index files
