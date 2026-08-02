# DNF5 Build System

The CMake-based build system for DNF5. Covers build options, dependencies, build order, and compiler flags.

> **Source of truth**: <https://github.com/rpm-software-management/dnf5/blob/main/CMakeLists.txt>.

## Overview

- **Build system**: CMake (minimum version 3.21)
- **Language**: C++20 (`set(CMAKE_CXX_STANDARD 20)`)
- **Compiler flags**: `-Wall -Wextra -Werror`, strict warning flags
- **Formatting**: `.clang-format`, `.clang-tidy`
- **Version**: Defined in `VERSION.cmake` (prime=5, major=4, minor=4, micro=0)

## Build Options

### Components (sub-packages)

| Option | Default | Description |
|--------|---------|-------------|
| `WITH_DNF5` | ON | Build dnf5 CLI |
| `WITH_DNF5DAEMON_CLIENT` | ON | Build dnf5daemon-client |
| `WITH_DNF5DAEMON_SERVER` | ON | Build dnf5daemon-server |
| `WITH_LIBDNF5_CLI` | ON | Build libdnf5-cli |
| `WITH_DNF5_PLUGINS` | ON | Build dnf5 plugins |
| `WITH_DNF5_OBSOLETES_DNF` | ON | Build with DNF5 providing files conflicting with DNF4 |

### Features

| Option | Default | Description |
|--------|---------|-------------|
| `WITH_ACL` | ON | POSIX ACL support |
| `WITH_COMPS` | ON | Comps groups support |
| `WITH_MODULEMD` | ON | Modulemd support |
| `WITH_SYSTEMD` | ON | systemd & D-Bus features |
| `WITH_TESTS` | ON | Build tests |
| `WITH_PERFORMANCE_TESTS` | OFF | Performance tests (DEBUG only) |
| `WITH_DNF5DAEMON_TESTS` | OFF | Daemon tests (requires D-Bus) |
| `WITH_SANITIZERS` | OFF | Address, leak, undefined sanitizers |

### Bindings

| Option | Default | Description |
|--------|---------|-------------|
| `WITH_GO` | OFF | Go bindings |
| `WITH_PERL5` | ON | Perl 5 bindings |
| `WITH_PYTHON3` | ON | Python 3 bindings |
| `WITH_RUBY` | ON | Ruby bindings |

### Plugins

| Option | Default | Description |
|--------|---------|-------------|
| `WITH_PLUGIN_ACTIONS` | ON | Actions plugin |
| `WITH_PLUGIN_APPSTREAM` | ON | AppStream plugin |
| `WITH_PLUGIN_EXPIRED_PGP_KEYS` | ON | Expired PGP keys plugin |
| `WITH_PLUGIN_RHSM` | OFF | Red Hat Subscription Manager plugin |
| `WITH_PLUGIN_MANIFEST` | ON | Manifest plugin |
| `WITH_PYTHON_PLUGINS_LOADER` | ON | Python plugin loader |
| `WITH_PLUGIN_LOCAL` | ON | Local plugin |

## Dependencies

| Dependency | Purpose |
|------------|---------|
| `libsolv` | Dependency solver |
| `librepo` | Repository downloading |
| `sdbus-cpp` | D-Bus bindings (for daemon) |
| `fmt` | String formatting |
| `librpm` | RPM package library |
| `modulemd` | Modulemd support (optional, via `WITH_MODULEMD`) |
| `libacl` | POSIX ACL support (optional, via `WITH_ACL`) |

## Build Order

From the root `CMakeLists.txt`:

1. `common/` — Shared utilities across all components
2. `etc/` — systemd unit files, tmpfiles.d config
3. `include/` — Public headers for libdnf5 and libdnf5-cli
4. `libdnf5/` — Core library
5. `libdnf5-cli/` — CLI helper library
6. `libdnf5-plugins/` — Library-level plugins
7. `doc/` — Documentation (Sphinx + Doxygen)
8. `bindings/` — Language bindings (Python3, Perl5, Ruby, Go)
9. `dnf5daemon-server/` — D-Bus daemon server
10. `dnf5daemon-client/` — D-Bus daemon client CLI
11. `dnf5/` — Main CLI application
12. `dnf5-plugins/` — CLI plugins
13. `test/` — Test suite

## Building

```bash
# Clone
git clone https://github.com/rpm-software-management/dnf5.git
cd dnf5

# Configure
cmake -B build -DCMAKE_BUILD_TYPE=Debug

# Build
cmake --build build -j$(nproc)

# Install (optional)
cmake --install build
```

### Debug Build

```bash
cmake -B build -DCMAKE_BUILD_TYPE=Debug -DWITH_SANITIZERS=ON
cmake --build build -j$(nproc)
```

### Build Only Specific Components

```bash
cmake -B build -DWITH_DNF5DAEMON_SERVER=OFF -DWITH_DNF5DAEMON_CLIENT=OFF -DWITH_TESTS=OFF
cmake --build build -j$(nproc)
```

## Testing

Tests are in `test/` and mirror the source structure:

| Directory | Purpose |
|-----------|---------|
| `test/libdnf5/` | Unit tests for libdnf5 |
| `test/libdnf5-cli/` | Tests for CLI library |
| `test/dnf5-plugins/` | Tests for dnf5 plugins |
| `test/dnf5daemon-server/` | Daemon server tests |
| `test/shared/` | Shared test utilities |
| `test/data/` | Test data (repo metadata, packages) |
| `test/tutorial/` | Tutorial tests |
| `test/go/`, `test/perl5/`, `test/python3/`, `test/ruby/` | Binding language tests |

## Coding Style

### C++

- `.clang-format` defines formatting rules
- `.clang-tidy` defines linting rules
- Scripts: `clang-format` and `clang-tidy-changed` for applying/checking

### Python

- Follows PEP 8 with project-specific additions

### Best Practices (from docs)

- Use unit tests
- Avoid `shared_ptr` in SWIG bindings
- Use templates/lambdas sparingly
- Document code with Doxygen-style comments (headers are the primary documentation)

## Project Layout

```
dnf5/
├── .github/              # GitHub workflows, CI
├── bindings/             # Language bindings
├── cmake/                # CMake helper modules
├── common/               # Shared utility code
├── dnf5/                 # Main CLI binary
│   ├── commands/         # 30+ command implementations
│   ├── config/           # Default config files
│   ├── include/          # dnf5 public headers (IPlugin interface)
├── dnf5-plugins/         # CLI plugins (automatic, builddep, copr, etc.)
├── dnf5daemon-client/    # D-Bus daemon client
├── dnf5daemon-server/    # D-Bus daemon server
├── doc/                  # Sphinx + Doxygen documentation
├── etc/                  # systemd unit files, tmpfiles.d
├── include/              # Public headers for libdnf5/libdnf5-cli
├── libdnf5/              # Core library implementation
├── libdnf5-cli/          # CLI helper library
├── libdnf5-plugins/      # Library-level plugins
├── test/                 # Test suite
├── CMakeLists.txt        # Root build file
├── VERSION.cmake         # Version definitions
└── dnf5.spec             # RPM spec file
```

## Licensing

| Component | License |
|-----------|---------|
| libdnf5, libdnf5-cli, libdnf5-plugins | LGPLv2.1+ |
| dnf5, dnf5daemon-server, dnf5daemon-client, dnf5-plugins | GPLv2+ |
| Language bindings | Follow library's LGPL |
