---
name: qemu
description: >
  Use when working with QEMU — qemu-system-*, qemu-img, qemu-nbd,
  qemu-storage-daemon, qemu-pr-helper, block drivers, machine types, CPU models,
  device emulation, networking, display options, QMP protocol, and guest agent.
  Triggers on: "qemu", "qemu-system", "qemu-img", "qemu-nbd", "qemu-kvm",
  "QEMU", "qcow2", "qcow", "qed", "virtio", "virgl", "vhost-user", "VirtIO",
  "block driver", "machine type", "SEV", "SEV-SNP", "TDX", "snapshot",
  "backing file", "NBD", "Network Block Device", "QMP",
  "QEMU Monitor Protocol", "guest agent", "qemu-ga", "cpu model",
  "host passthrough", "SPICE", "VNC".
user-invocable: true
---

# QEMU Skill

Comprehensive reference for the QEMU command suite — full-system emulators, disk image tools, block storage infrastructure, and management protocols.

## Sub-Resources

| Keywords | Reference |
|----------|----------|
| architecture, command inventory, component relationships, process model, QEMU vs KVM vs accelerators | [references/overview.md](references/overview.md) |
| qemu-system-ARCH, -machine, -accel, -smp, -m, -cpu, -M, -name, -uuid, -rtc, -icount, -S, -gdb, -no-reboot, -no-shutdown, -daemonize, snapshot mode, -overcommit, -msg, -fw_cfg, -smbios, -bios, -kernel, -initrd, -append, -dtb, -initiator, -semihosting, -semihosting-config | [references/qemu-system.md](references/qemu-system.md) |
| qemu-img, create, convert, check, info, snapshot, commit, rebase, resize, amend, map, measure, bench, dd, compare, bitmap, image formats | [references/qemu-img.md](references/qemu-img.md) |
| qemu-nbd, NBD server, export, connect, disconnect, TLS, persistent, bitmap, list, offset, snapshot | [references/qemu-nbd.md](references/qemu-nbd.md) |
| qemu-storage-daemon, block exports, block jobs, QMP control, --blockdev, --export, --nbd-server | [references/qemu-storage-daemon.md](references/qemu-storage-daemon.md) |
| qemu-pr-helper, SCSI persistent reservations, helper socket, daemon mode, privilege drop | [references/qemu-pr-helper.md](references/qemu-pr-helper.md) |
| qemu-block-drivers, raw, qcow2, qed, vmdk, vhdx, vdi, luks, rbd, iscsi, nfs, nbd, http, https, ssh, blkdebug, blkverify, copy-on-read, copy-before-write, throttle, replication, compression, encryption, backing files, snapshot, preallocation, discard, detect-zeroes | [references/block-drivers.md](references/block-drivers.md) |
| -blockdev, -drive, -device storage controllers, virtio-blk, virtio-scsi, NVMe, AHCI, IDE, FUSE exports, vduse-blk, vhost-user-blk | [references/block-devices.md](references/block-devices.md) |
| -netdev, -nic, -net, network backends, TAP, user mode, bridge, L2TPv3, socket, VDE, vhost-user, vhost-vdpa, virtio-net, e1000, rtl8139, i8255x, vmxnet3 | [references/networking.md](references/networking.md) |
| -display, -vga, -full-screen, -spice, -vnc, gtk, sdl, egl-headless, dbus, curses, cocoa, virgl, egl, opengl, window options | [references/display.md](references/display.md) |
| -device, PCI, USB, virtio, audio, virtio-gpu, virtio-balloon, virtio-rng, virtio-fs, vfio, vhost, ivshmem, audio devices, USB devices | [references/devices.md](references/devices.md) |
| machine types, pc-i440fx, pc-q35, microvm, virt, versioned machine types, migration compatibility | [references/machine.md](references/machine.md) |
| CPU models, features, topology, host passthrough, named models, host model, -cpu, -smp topology, -numa, ABI levels, AMD SEV, SEV-ES, SEV-SNP, Intel TDX, x86, ARM, RISC-V, s390x, PPC | [references/cpu-models.md](references/cpu-models.md) |
| QMP protocol, qemu-qmp-ref, commands, events, introspection, qmp-shell, -qmp, -hmp, HMP monitor | [references/qmp.md](references/qmp.md) |
| guest agent, qemu-ga, QGA protocol, guest-sync, guest-info, guest-exec, guest-fsfreeze, guest-network-get-interfaces, qemu-ga-ref | [references/guest-agent.md](references/guest-agent.md) |
| man pages, documentation, carapace completions, existing completers | [references/man-docs.md](references/man-docs.md) |

## Quick Guide

| When you need to... | Start here |
|---------------------|------------|
| Understand the overall architecture and run your first VM | [references/overview.md](references/overview.md) |
| Configure a VM's CPU, memory, machine type, and accelerators | [references/qemu-system.md](references/qemu-system.md) |
| Create or manipulate a disk image | [references/qemu-img.md](references/qemu-img.md) |
| Export a disk image over the network | [references/qemu-nbd.md](references/qemu-nbd.md) |
| Run a standalone storage daemon | [references/qemu-storage-daemon.md](references/qemu-storage-daemon.md) |
| Set up SCSI persistent reservations | [references/qemu-pr-helper.md](references/qemu-pr-helper.md) |
| Understand image formats (qcow2, raw, LUKS, RBD) | [references/block-drivers.md](references/block-drivers.md) |
| Configure block devices (virtio-blk, virtio-scsi, NVMe) | [references/block-devices.md](references/block-devices.md) |
| Set up networking (TAP, user-mode, virtio-net) | [references/networking.md](references/networking.md) |
| Configure display output and remote access | [references/display.md](references/display.md) |
| Add devices (PCI, USB, audio, vfio) | [references/devices.md](references/devices.md) |
| Choose the right machine type | [references/machine.md](references/machine.md) |
| Configure CPU models and features | [references/cpu-models.md](references/cpu-models.md) |
| Control QEMU via QMP | [references/qmp.md](references/qmp.md) |
| Install and use the guest agent | [references/guest-agent.md](references/guest-agent.md) |
