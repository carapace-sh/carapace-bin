# Machine Types

The machine type defines the motherboard/chipset, devices, and firmware interface for the VM. Use `-machine help` to list available machines for your architecture.

## x86 Machine Types

### pc-i440fx

Standard PC with Intel 440FX chipset. The original QEMU PC machine.

```
-machine pc-i440fx-11.0   # current version (alias: pc)
-machine pc-i440fx-9.0    # versioned (migration-compatible)
```

- i440FX host bridge + PIIX3 southbridge
- ISA bus (keyboard, serial, parallel, floppy, RTC via PIIX3)
- PCI bus with PIRQ routing
- Firmware: SeaBIOS (legacy BIOS) or OVMF (UEFI)
- Legacy: `-machine pc`

### pc-q35

Standard PC with Intel Q35 chipset. Modern replacement for i440FX.

```
-machine q35
-machine pc-q35-11.0
```

- Q35 MCH + ICH9 southbridge
- PCIe root bus (native PCIe)
- AHCI (SATA) on ICH9
- ICH9 integrated Intel HD Audio and GbE
- Better for modern operating systems
- Supports PCIe native hotplug
- Firmware: SeaBIOS or OVMF

### microvm

Minimal machine type. Ultra-small, no PCI, no ACPI.

```
-machine microvm
```

- No PCI bus at all
- Only virtio-mmio devices
- Device tree (DT) instead of ACPI
- Minimal firmware overhead
- Use with `-kernel`, x86 needs `-machine microvm,acpi=off` for some configs
- Good for fast boot, tiny VMs, and kata containers

### nitro-enclave

```
-machine nitro-enclave
```

AWS Nitro Enclave compatible machine. Minimal, no persistent storage, no network (except vsock). For building confidential computing enclaves.

## ARM / AArch64 Machine Types

### virt

```
-machine virt
```

- Generic ARM virtual platform
- No physical hardware correspondence
- VirtIO devices (PCIe or MMIO)
- ACPI or device tree
- UEFI boot via `-bios QEMU_EFI.fd`

### Other ARM machines

```
-machine raspi4b      # Raspberry Pi 4 Model B
-machine virt-11.0    # versioned virt
-machine sbsa-ref     # ARM SBSA reference (server-class)
```

## RISC-V Machine Types

### virt

```
-machine virt
```

- Generic RISC-V virtual platform
- VirtIO MMIO and PCIe
- OpenSBI firmware

## s390x Machine Types

### s390-ccw-virtio

```
-machine s390-ccw-virtio
```

- IBM S/390 channel I/O
- VirtIO via CCW (channel command word)
- No PCI (zPCI adds PCI passthrough)
- IPL from DASD or SCSI

## Versioned Machine Types

Machine types are versioned to support live migration compatibility across QEMU releases:

```
pc-i440fx-9.0       # migratable to pc-i440fx-9.x
pc-i440fx-9.2       # migratable to pc-i440fx-9.x
pc-q35-9.0          # migratable to pc-q35-9.x
```

Each release adds a new aliased version. The `pc` alias always points to the latest `pc-i440fx` version; `q35` to the latest `pc-q35`.

The `latest` (unversioned) machine type enables new features but breaks migration to older QEMU. Versioned types preserve migration compatibility within the major version range.

## Migration Compatibility

To enable live migration:
1. Both hosts must use compatible QEMU versions
2. Must use the same versioned machine type
3. CPU model must be stable across hosts (use a named model, not `host`)
4. Device configuration must match
5. Check with `qemu-system-x86_64 -machine pc-q35-9.0,usb=on` — enable needed features explicitly

## See Also

- `qemu-system-ARCH -machine help` for available machines
- [qemu-system.md](qemu-system.md) for machine properties
