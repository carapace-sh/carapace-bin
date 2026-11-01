# DNF5 Installroot

The `--installroot` parameter specifies an alternative root filesystem for all packaging operations.

> **Source of truth**: <https://dnf5.readthedocs.io/en/stable/misc/installroot.7.html>.

## Overview

It is like doing `chroot <root> dnf`, but `--installroot` allows DNF5 to work before the chroot is created.

```bash
dnf5 --installroot=/mnt/root install system-release
```

## Path Resolution Rules

| Item | Source | Notes |
|------|--------|-------|
| `cachedir`, `system_cachedir` | Installroot | Taken from or stored in installroot |
| Log files | Installroot | |
| `releasever` | Installroot | Detected from RPMDB inside installroot |
| `gpgkey` | Installroot | OpenPGP keys imported into installroot from path relative to host |
| Configuration file | Installroot | Unless `--use-host-config` is passed |
| `reposdir` | Installroot | Unless `--use-host-config` is passed |
| `vars` | Installroot | Unless `--use-host-config` is passed |
| Command-line paths (`--config`, `--setopt=reposdir=`, `--setopt=cachedir=`, `--setopt=system_cachedir=`, `--setopt=logdir=`, `--setopt=varsdir=`) | **Host** | Always relative to host, no exceptions |
| `pluginpath`, `pluginconfpath` | **Host** | Always relative to host |

## `--use-host-config`

When `--use-host-config` is passed:
- Configuration files and variable definitions are taken from the **host system** (`/`)
- `releasever` is detected from the host (`/`)
- Useful when the installroot doesn't have its own configuration yet

## Important Notes

### Creating a New Installroot

When creating an installroot, use `--releasever=RELEASEVER`:

```bash
dnf5 --installroot=/mnt/root --releasever=40 install system-release
```

Otherwise `$releasever` is taken from the RPMDB inside the installroot (empty at creation time, causing transaction failure).

### Modular Systems

Use `--setopt=module_platform_id=<module_platform_name:stream>` when creating the installroot:

```bash
dnf5 --installroot=/mnt/root --setopt=module_platform_id=platform:el9 install system-release
```

Otherwise `module_platform_id` is taken from `/etc/os-release` inside the installroot (empty at creation, modular dependency could be unsatisfied, modules content could be excluded).

## Examples

```bash
# Permanently set releasever in the installroot
dnf5 --installroot=/mnt/root --releasever=40 install system-release

# Use config from a specific path
dnf5 --installroot=/mnt/root --setopt=reposdir=/path/to/repodir --config /path/dnf.conf upgrade

# Use host configuration
dnf5 --installroot=/mnt/root --use-host-config --releasever=40 install @core
```
