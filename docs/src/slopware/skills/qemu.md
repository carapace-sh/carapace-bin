# qemu

QEMU — machine emulator and virtualizer.

| reference                              | description                                          |
|----------------------------------------|------------------------------------------------------|
| [overview][ref-overview]               | architecture, command inventory, process model        |
| [qemu-system][ref-qemu-system]         | qemu-system-ARCH, -machine, -accel, -cpu, -smp, -m   |
| [qemu-img][ref-qemu-img]               | qemu-img, create, convert, check, snapshot, rebase   |
| [qemu-nbd][ref-qemu-nbd]               | qemu-nbd, NBD server, export, connect, TLS           |
| [qemu-storage-daemon][ref-qemu-storage-daemon] | block exports, block jobs, QMP control      |
| [qemu-pr-helper][ref-qemu-pr-helper]   | SCSI persistent reservations, daemon mode            |
| [block-drivers][ref-block-drivers]     | raw, qcow2, vmdk, rbd, iscsi, nfs, ssh, encryption  |
| [block-devices][ref-block-devices]     | -blockdev, -drive, virtio-blk, virtio-scsi, NVMe    |
| [networking][ref-networking]           | -netdev, -nic, TAP, user mode, bridge, vhost-user   |
| [display][ref-display]                 | -display, -vga, SPICE, VNC, gtk, sdl, virgl        |
| [devices][ref-devices]                 | -device, PCI, USB, virtio, audio, vfio, vhost      |
| [machine][ref-machine]                 | machine types, pc-q35, virt, microvm, migration      |
| [cpu-models][ref-cpu-models]           | CPU models, features, topology, SEV, TDX             |
| [qmp][ref-qmp]                         | QMP protocol, commands, events, introspection       |
| [guest-agent][ref-guest-agent]         | qemu-ga, QGA protocol, guest-sync, guest-exec       |
| [man-docs][ref-man-docs]               | man pages, documentation, carapace completions       |

[ref-overview]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/overview.md
[ref-qemu-system]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/qemu-system.md
[ref-qemu-img]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/qemu-img.md
[ref-qemu-nbd]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/qemu-nbd.md
[ref-qemu-storage-daemon]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/qemu-storage-daemon.md
[ref-qemu-pr-helper]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/qemu-pr-helper.md
[ref-block-drivers]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/block-drivers.md
[ref-block-devices]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/block-devices.md
[ref-networking]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/networking.md
[ref-display]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/display.md
[ref-devices]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/devices.md
[ref-machine]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/machine.md
[ref-cpu-models]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/cpu-models.md
[ref-qmp]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/qmp.md
[ref-guest-agent]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/guest-agent.md
[ref-man-docs]:https://github.com/carapace-sh/carapace-bin/blob/master/skills/qemu/references/man-docs.md