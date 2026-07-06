# QEMU Overview

QEMU is a generic and open source machine emulator and virtualizer. This document covers the high-level architecture, the command inventory, and how the components relate.

## Architecture

QEMU runs as a single process that models a complete computer system. There are two modes of operation:

**Full-system emulation** (`qemu-system-*`): Emulates an entire machine — CPU, memory, devices, buses — as a user-space process. Can use an accelerator (KVM, HVF, WHPX, TCG) to speed up execution.

**Userspace emulation** (`qemu-*`): Runs a single Linux binary compiled for a different architecture. Not covered in this skill.

## Component Inventory

The QEMU project ships these commands:

| Command | Section | Role |
|---------|---------|------|
| `qemu-system-ARCH` | Emulator | Full-system virtual machine (one binary per arch, e.g. `qemu-system-x86_64`, `qemu-system-aarch64`, `qemu-system-riscv64`) |
| `qemu-img` | Disk tool | Create, convert, modify, inspect disk images offline |
| `qemu-nbd` | Network block | Export disk images as NBD server; connect to NBD devices |
| `qemu-storage-daemon` | Storage daemon | Long-lived process for block operations via QMP, no VM needed |
| `qemu-pr-helper` | SCSI helper | SCSI persistent reservation helper daemon |
| `qemu-ga` | Guest agent | In-guest agent for host↔guest communication (separate binary) |

And these reference man pages:

| Man page | Content |
|----------|---------|
| `qemu(1)` | `qemu-system-x86_64` documentation (canonical for all archs) |
| `qemu-block-drivers(7)` | All supported disk image formats and network storage protocols |
| `qemu-cpu-models(7)` | CPU model recommendations, features, and ABI compatibility |
| `qemu-ga-ref(7)` | Guest Agent Protocol command reference |
| `qemu-qmp-ref(7)` | QEMU Monitor Protocol command/event reference |
| `qemu-storage-daemon-qmp-ref(7)` | Storage daemon QMP command reference |

## QEMU vs KVM vs Accelerators

QEMU emulates CPUs in software using **TCG** (Tiny Code Generator). This is slow but runs any guest on any host.

**KVM** (Kernel-based Virtual Machine) is a Linux kernel module that allows QEMU to run guest code directly on the host CPU. QEMU becomes a VMM (Virtual Machine Monitor), handling device emulation while guest code runs natively.

Other accelerators:
- **HVF** — macOS Hypervisor.framework (Intel Macs)
- **WHV/WHVP** — Windows Hypervisor Platform
- **NVMM** — NetBSD virtual machine monitor
- **Nitro** — Apple Virtualization.framework (Apple Silicon)
- **MSHV** — Microsoft Hyper-V

Enable with `-accel kvm` or `-enable-kvm` (legacy). Multiple accelerators can be specified as fallbacks: `-accel kvm:tcg`.

## Process Model

A running QEMU process:
1. Parses the command line, building the VM configuration (machine, CPUs, RAM, devices, block backends, network)
2. Initializes the machine and all devices
3. Connects to the accelerator (if requested)
4. Opens a QMP/HMP monitor socket (if requested)
5. Loads BIOS/firmware or kernel/initrd
6. Starts the guest CPUs
7. Enters the main event loop — dispatches timers, I/O, monitor commands, and VCPU execution

## ELF Binary Variants

Each target architecture has its own `qemu-system-ARCH` binary. On Debian-based systems these are in separate packages (`qemu-system-x86`, `qemu-system-arm`, etc.) or the metapackage `qemu-system-all`.

Common variants:

| Binary | Architecture |
|--------|-------------|
| `qemu-system-x86_64` | x86-64 |
| `qemu-system-i386` | x86 (32-bit) |
| `qemu-system-aarch64` | AArch64 / ARM64 |
| `qemu-system-arm` | ARM (32-bit) |
| `qemu-system-riscv64` | RISC-V 64-bit |
| `qemu-system-ppc64` | PowerPC 64-bit |
| `qemu-system-s390x` | IBM S/390x |

All `qemu-system-ARCH` binaries share the same option syntax. The machine-specific options differ (`-machine` types, `-cpu` models).
