# DNF5 Forcearch

The `--forcearch=<arch>` parameter overrides the system architecture detected by DNF5.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/misc/forcearch.7.html>.

## Overview

`--forcearch` allows querying repositories for packages not compatible with the host system and installing them.

- Any architecture can be specified.
- Using a package with an architecture not supported natively by the CPU will require emulation (e.g., `qemu-user-static`).
- The `arch` and `basearch` config options can no longer be set in config files (they were dropped from DNF4). Use `--forcearch` instead.

## Supported Commands

`distro-sync`, `download`, `group`, `info`, `install`, `list`, `makecache`, `repo`, `repoquery`, `search`, `swap`

## Examples

```bash
# Install the AArch64 version regardless of host architecture
dnf5 install --forcearch=aarch64 my-example-package

# Download the hello package for s390x
dnf5 download --forcearch=s390x hello

# Query all packages available for AArch64
dnf5 repoquery --forcearch=aarch64 --arch=aarch64
```

## Interaction with `--arch`

Both `--arch` and `--forcearch` are needed when the system has a different native architecture:

- `--arch` filters for only packages matching the specified architecture
- `--forcearch` sets the `arch` and `basearch` substitution variables for correct repository queries

## Related

- [repo-variables.md](repo-variables.md) — `$arch` and `$basearch` variables
- [migration.md](migration.md) — `arch`/`basearch` config options were dropped
