# Block Devices — Storage Controllers and Configuration

How to attach storage to a QEMU VM. There are two configuration syntaxes: modern (`-blockdev` + `-device`) and legacy (`-drive`).

## -blockdev and -device (Modern API)

Modern QEMU storage uses a node-based graph. `-blockdev` creates block driver nodes, `-device` attaches a virtual controller.

```
-blockdev file,node-name=file0,filename=vmdisk.qcow2
-blockdev qcow2,node-name=disk0,file=file0
-device virtio-blk,drive=disk0
```

The block graph can be arbitrarily complex (filters, snapshots, mirroring). Nodes are referenced by `node-name`.

## -drive (Legacy API)

Syntax:
```
-drive file=FILENAME[,if=TYPE][,bus=N][,unit=N][,index=N][,media=cdrom|disk][,format=FMT][,cache=MODE][,snapshot=on|off][,read-only=on|off][,werror=ACTION][,rerror=ACTION]
```

Shortcut that combines blockdev creation and device attachment. `if=TYPE` selects the controller.

## Storage Controllers

### virtio-blk

Para-virtualized block device. Highest performance.

```
-device virtio-blk,drive=disk0 [,num-queues=N] [,queue-size=QS] [,serial=STR]
```
- `num-queues` — multi-queue (match host CPU count)
- `queue-size` — virtqueue depth (default 256)
- Uses `PCI` or `virtio-mmio` depending on machine type
- Legacy: `-drive file=...,if=virtio`

### virtio-scsi

SCSI HBA emulation with virtio transport. Supports more devices per controller than virtio-blk.

```
-device virtio-scsi-pci,id=scsi0 [,num_queues=N]
-device scsi-hd,drive=disk0,bus=scsi0.0,channel=0,scsi-id=0,lun=0
-device scsi-cd,drive=cdrom0,bus=scsi0.0
```

Supports hotplug, multiple LUNs, SCSI command passthrough. Use `scsi-block` for real-device passthrough with full SCSI semantics.

### AHCI / SATA

```
-device ahci,id=ahci0
-drive file=...,if=none,id=drive0,format=qcow2
-device ide-hd,drive=drive0,bus=ahci0.0
```

Emulates Intel ICH9 AHCI controller. Each port is a separate bus (`ahci0.0`, `ahci0.1`, etc.)

### IDE (PIIX3 / PIIX4)

Legacy PATA interface. Default for `pc-i440fx-*` machines.

```
-device ide-hd,drive=disk0,bus=ide.0,unit=0
-device ide-cd,drive=cdrom0,bus=ide.1,unit=0
```

Or simply `-drive file=...,if=ide`.

### NVMe

```
-device nvme,serial=deadbeef,id=nvme0 [,num_queues=N] [,msix_qsize=N]
-device nvme-ns,drive=disk0,bus=nvme0.0
```

Emulates an NVMe controller. Supports multiple namespaces per controller.

### USB Storage

```
-device usb-storage,drive=disk0
```

Emulates a USB mass storage device.

### FUSE Export (no VM)

The `qemu-storage-daemon` can export a block node via FUSE:
```
--export fuse,id=export0,node-name=disk,mountpoint=/mnt/disk.raw,writable=on
```

See [qemu-storage-daemon.md](qemu-storage-daemon.md).

## Cache Modes

Set via `-drive cache=MODE` or `--blockdev ...,cache={direct,writeback,...}`:

| Mode | Host Cache | Flush | Safe for |
|------|-----------|-------|----------|
| `writeback` | Yes | Ignore guest cache flush | Guest with battery-backed or virtio |
| `writethrough` | Yes | On every write | Single VM, no risk |
| `none` | No (O_DIRECT) | Ignore guest flush | Dedicated storage, high perf |
| `directsync` | No (O_DIRECT) | On every write | Safest + direct |
| `unsafe` | Yes | Never flush | Development/testing only |

## I/O and Discard Options

```
[...],aio=threads|native|io_uring
[...],discard=ignore|unmap
[...],detect-zeroes=off|on|unmap
```

- `aio=native` — Linux AIO (requires `cache=none` or `cache.direct=on`)
- `aio=io_uring` — Linux io_uring (best performance, modern kernels)
- `discard=unmap` — forward discard/trim to storage
- `detect-zeroes=unmap` — punch holes on zero writes (thin provisioning)

## Error Handling

```
-drive ...,werror=report|ignore|enospc|stop|auto,rerror=report|ignore|stop|auto
```

- `report` — report error to guest
- `ignore` — silently ignore
- `enospc` — pause VM, wait for ENOSPC resolution
- `stop` — pause VM
- `auto` — report for reads, enospc for writes (default)

## Object Definitions

Shared by `-blockdev` and `-drive`:

```
-object secret,id=sec0,data=PASSWORD
-object secret,id=sec1,file=/path/to/secret
-object iothread,id=iothread0
-object throttle-group,id=throttle0,x-bps-read=N,x-bps-write=N
```
