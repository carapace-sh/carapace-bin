# Man Pages and Documentation

Reference for QEMU man pages and how they relate to carapace completions.

## Man Pages

All QEMU man pages are available on the system (`man -k qemu`):

| Man page | Section | Content |
|----------|---------|---------|
| `qemu` | 1 | `qemu-system-x86_64` — canonical system emulator documentation |
| `qemu-img` | 1 | Disk image utility |
| `qemu-nbd` | 8 | Network Block Device server |
| `qemu-pr-helper` | 8 | SCSI persistent reservation helper |
| `qemu-storage-daemon` | 1 | Storage daemon |
| `qemu-block-drivers` | 7 | Block drivers and image formats |
| `qemu-cpu-models` | 7 | CPU model configuration |
| `qemu-ga-ref` | 7 | Guest Agent Protocol reference |
| `qemu-qmp-ref` | 7 | QMP command/event reference |
| `qemu-storage-daemon-qmp-ref` | 7 | Storage daemon QMP reference |
| `qemu-system-ARCH` | 1 | Per-architecture system emulators (32 distinct archs) |

## Man Page Topics by Section

### Section 1 (Commands)
- `qemu` — system emulator (x86_64, the canonical page)
- `qemu-img` — disk image utility
- `qemu-storage-daemon` — standalone storage daemon
- `qemu-system-aarch64`, `qemu-system-alpha`, ..., `qemu-system-xtensaeb` — each emulator binary

### Section 7 (References)
- `qemu-block-drivers` — disk image formats, network protocols, filter drivers
- `qemu-cpu-models` — CPU model recommendations, features, ABI levels
- `qemu-ga-ref` — QEMU Guest Agent Protocol (QGA) commands and events
- `qemu-qmp-ref` — QEMU Monitor Protocol (QMP) commands, events, and data types
- `qemu-storage-daemon-qmp-ref` — storage daemon QMP subset

### Section 8 (System Administration)
- `qemu-nbd` — NBD server/utility
- `qemu-pr-helper` — persistent reservation helper daemon

## Carapace Completions

In carapace-bin, QEMU completers live under `completers/linux/` (they are Linux-specific tools). As of this skill's creation, no QEMU completers exist in the carapace-bin repository.

To create completions for QEMU commands, follow the patterns in `skills/carapace/references/spec.md` for YAML specs or `skills/carapace/references/scrape.md` for automated spec generation.

Key QEMU commands that benefit from completion:
- `qemu-system-x86_64` — machine type names (`-machine help`), CPU models (`-cpu help`), device names, block device options
- `qemu-img` — command names, format types, image filenames
- `qemu-nbd` — format, cache modes, AIO modes, TLS objects
- `qemu-storage-daemon` — export types, blockdev options, chardev options

## Documentation Best Practices

When documenting QEMU completers (see `skills/man-docs/`):

- Use the canonical man page names for `documentation.command` values
- For `documentation.flag`, reference the flag name as documented in `qemu(1)` (e.g. `-accel`, `-machine`, `-blockdev`)
- Architecture-specific notes should reference the specific arch man page if behavior differs
- Flag descriptions should match the option descriptions in `qemu(1)` where possible
