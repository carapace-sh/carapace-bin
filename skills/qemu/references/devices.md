# Devices — PCI, USB, VirtIO, and More

Attach emulated devices to a VM using `-device`. Devices are instantiated on the system bus (typically PCI in x86 machines; the bus topology depends on `-machine`).

## PCI Device Addressing

```
-device e1000,bus=pci.0,addr=0x6
-device virtio-blk-pci,bus=pci.0,addr=0x7
-device vfio-pci,host=0000:01:00.0
```

PCIe devices use root ports: `-device pcie-root-port,slot=X,id=rpiX` then `-device ...,bus=rpiX`.

## VirtIO Family

The virtio family provides para-virtualized devices. They require virtio drivers in the guest.

| Device | Description |
|--------|-------------|
| `virtio-blk-pci` | Block device (storage) |
| `virtio-scsi-pci` | SCSI controller |
| `virtio-net-pci` | Network adapter |
| `virtio-gpu-pci` | GPU (with virgl 3D support) |
| `virtio-vga` | VGA-compatible GPU (combines VGA + virtio-gpu) |
| `virtio-balloon-pci` | Memory balloon for dynamic memory management |
| `virtio-rng-pci` | Random number generator |
| `virtio-serial-pci` | Serial ports (for guest agent) |
| `virtio-9p-pci` | Plan 9 shared filesystem (9p/virtio-fs predecessor) |
| `virtio-fs` | Shared filesystem via FUSE (requires vhost-user-fs daemon) |
| `virtio-input-*` | Input devices (keyboard, tablet, mouse) |
| `virtio-pmem-pci` | Persistent memory device |
| `virtio-mem-pci` | Memory hotplug device with flexible sizing |
| `virtio-iommu-pci` | IOMMU for virtio devices |
| `virtio-crypto-pci` | Crypto accelerator |

### VirtIO Common Properties

```
-device virtio-blk-pci,drive=disk0,num-queues=4,queue-size=256,iothread=iothread0
```

- `num-queues` — multi-queue count
- `iothread` — IOThread assignment
- `ioeventfd=on|off` — eventfd-based notification
- `vectors=N` — MSI-X vector count

## PCI Storage Controllers

| Device | Description |
|--------|-------------|
| `ahci` | AHCI/SATA controller (Intel ICH9) |
| `nvme` | NVMe controller (multiple namespaces) |
| `vhost-user-blk-pci` | vhost-user block device |
| `vhost-scsi-pci` | vhost SCSI target |
| `vhost-vsock-pci` | vsock (host-guest socket) |

## USB

USB is provided by a USB controller plus devices on its bus.

Controllers:
```
-device usb-ehci,id=ehci      # EHCI (USB 2.0)
-device nec-usb-xhci,id=xhci  # xHCI (USB 3.0)
-device piix3-usb-uhci,id=uhci # UHCI (USB 1.1, part of PIIX3)
-device qemu-xhci,id=xhci     # QEMU xHCI
```

Devices:
```
-device usb-tablet            # Absolute pointer (no grab needed)
-device usb-mouse             # Relative mouse
-device usb-kbd               # Keyboard
-device usb-storage,drive=disk0
-device usb-net,netdev=net0
-device usb-audio             # USB audio class device
-device usb-braille           # Brltty braille device
-device usb-ccid              # Smart card reader
-device usb-host,vendorid=N,productid=M  # Host USB passthrough
-device usb-redir,chardev=char0          # Usbredir (remote USB)
-device usb-wacom-tablet      # Wacom tablet
```

## Audio Devices

| Device | Description |
|--------|-------------|
| `intel-hda` + `hda-duplex`/`hda-micro`/`hda-output` | Intel HD Audio |
| `AC97` | Intel 82801AA AC97 |
| `cs4231a` | Windows Sound System |
| `sb16` | Sound Blaster 16 |
| `gus` | Gravis Ultrasound GF1 |
| `adlib` / `opl2` | Yamaha YM3812 |

Audio devices take an `audiodev=ID` parameter pointing to an `-audiodev` backend.

## VFIO / Device Passthrough

Pass a physical PCI device to the guest:

```
-device vfio-pci,host=0000:01:00.0
-device vfio-pci,host=0000:01:00.0,x-vga=on   # for GPU passthrough
```

Requires:
- `intel_iommu=on` or `amd_iommu=on` kernel parameter
- VFIO drivers loaded: `modprobe vfio-pci`
- Device detached from host driver
- `-machine ...,kernel_irqchip=on`

## vhost-user

Userspace virtio data path. Backend processes serve virtqueues over UNIX sockets.

```
-chardev socket,id=char0,path=/tmp/vhost-blk.sock,server=on,wait=off
-device vhost-user-blk-pci,chardev=char0,num-queues=4
```

Common backends:
- **SPDK** — vhost-blk/vhost-scsi
- **DPDK** — vhost-net (via testpmd or vswitch)
- **virtiofsd** — vhost-user-fs
- **QEMU storage daemon** — vhost-user-blk export

## ivshmem

Inter-VM shared memory:

```
-device ivshmem-plain,memdev=shmem0
-object memory-backend-file,id=shmem0,size=1M,mem-path=/dev/shm/ivshmem,share=on
```

For interrupt-capable shared memory, use the `ivshmem-doorbell` variant with a server.

## Input Devices

```
-device virtio-keyboard-pci   # virtio keyboard
-device virtio-tablet-pci     # virtio absolute pointer (no grab)
-device virtio-mouse-pci      # virtio mouse

-device usb-tablet            # USB absolute pointer (common, no grab)
-device usb-kbd               # USB keyboard
-device usb-mouse             # USB mouse
```

`usb-tablet` and `virtio-tablet` send absolute coordinates — the pointer doesn't need to be grabbed/released, which is smoother for desktop use.

## Misc Devices

| Device | Description |
|--------|-------------|
| `pvpanic` | Panic notification (ISA or PCI) |
| `vmware-svga` | VMware SVGA II graphics |
| `bochs-display` | Simple display with Bochs VBE |
| `ramfb` | Simple framebuffer on RAM (no PCI BAR) |
| `edu` | Educational PCI device for testing |
| `isadma` | ISA DMA controller |
| `i8259` | PIC (interrupt controller) |
| `mc146818rtc` | Real-time clock |
