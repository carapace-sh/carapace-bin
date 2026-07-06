# Networking — Backends, Frontends, and Configuration

QEMU networking consists of a **backend** (how the host connects to the network) and a **frontend** (the virtual device the guest sees).

## Configuration Syntax

Three ways to configure:

- `-nic` — modern shortcut, combines backend + frontend in one option
- `-netdev` + `-device` — explicit backend and device
- `-net` — legacy syntax (avoid)

## NIC Models (Frontends)

Common virtual NICs:

| Model | Use Case |
|-------|----------|
| `virtio-net-pci` | Para-virtualized, best performance |
| `virtio-net-ccw` | s390x virtio |
| `virtio-net-device` | RISC-V / non-PCI virtio |
| `e1000` | Intel PRO/1000, broad guest support |
| `e1000e` | Intel PRO/1000 (Gen 2), Windows 8+ |
| `rtl8139` | Realtek, legacy |
| `i82551` / `i82557b` / `i82559er` | Intel PRO, legacy |
| `vmxnet3` | VMware VMXNET3, paravirtualized |

Specify with `-nic model=MODEL` or `-device MODEL,...`.

## Network Backends

### User Mode (SLIRP)

```
-nic user
-netdev user,id=net0
```

- No root/setup needed
- NAT-based: guest gets 10.0.2.0/24
- Built-in DHCP, DNS proxy, TFTP, SMB server
- Port forwarding: `hostfwd=tcp::2222-:22`
- Limited performance, no ICMP/ping

### TAP

```
-netdev tap,id=net0 [,ifname=NAME] [,script=FILE] [,downscript=FILE] [,vhost=on|off] [,vnet_hdr=on|off]
```

- High performance, bridged networking
- Requires `tap` kernel module and `/dev/net/tun` access
- Default scripts: `/etc/qemu-ifup` and `/etc/qemu-ifdown`
- Alternative: `-netdev tap,id=net0,fd=FD` (pass pre-opened fd)

### Bridge

```
-netdev bridge,id=net0 [,br=BRIDGE] [,helper=HELPER]
```

- Uses `qemu-bridge-helper` setuid binary
- Default bridge: `virbr0` (libvirt default)
- No TAP management needed

### passt

```
-netdev passt,id=net0 [,hostfwd=...]
```

- User-plane NAT via passt daemon
- Modern replacement for SLIRP
- Better performance than user mode
- Supports IPv6, ICMP/ping
- Requires `passt` installed on host

### stream

```
-netdev stream,id=net0,addr.type=tcp,addr.host=HOST,addr.port=PORT[,server=on|off]
-netdev stream,id=net0,addr.type=unix,addr.path=PATH[,server=on|off]
```

- TCP or UNIX stream socket connection
- Replaces legacy `socket` for new development
- Supports `server=on|off` for listening vs connecting

### dgram

```
-netdev dgram,id=net0,local.addr=...,remote.addr=...
-netdev dgram,id=net0,remote.type=udp,remote.host=HOST,remote.port=PORT,mcast=GROUP:PORT
```

- UDP/datagram transport
- Supports local/remote addressing and multicast

### af-xdp

```
-netdev af-xdp,id=net0,ifname=eth0 [,start-queue=N] [,queues=N]
```

- AF_XDP socket (high-performance kernel bypass)
- Zero-copy packet processing
- Requires `CONFIG_XDP_SOCKETS` kernel option

### L2TPv3

```
-netdev l2tpv3,id=net0,src=SRC,dst=DST [,srcport=N] [,dstport=N] [,udp=on] [,cookie64=...] [,counter=on]
```

- Transport Ethernet over L2TPv3 (tunnel)
- No extra privileges needed (UDP socket)
- Useful for connecting VMs across network boundaries

### Socket

```
-netdev socket,id=net0,listen=:PORT    # server
-netdev socket,id=net0,connect=HOST:PORT  # client
-netdev socket,id=net0,fd=FD             # multicast: mcast=GROUP:PORT
```

### VDE

```
-netdev vde,id=net0 [,sock=PATH] [,port=N] [,group=GROUP] [,mode=MODE]
```

- Virtual Distributed Ethernet
- Requires `vde_switch` running

### vhost-user

```
-chardev socket,id=char0,path=/tmp/vhost.sock,server=on,wait=off
-netdev vhost-user,id=net0,chardev=char0 [,queues=N]
```

- Userspace virtio data path
- DPDK / SPDK / VPP integration
- Each queue is a separate vring

### vhost-vdpa

```
-netdev vhost-vdpa,id=net0,vhostdev=/dev/vhost-vdpa-0
```

- Hardware vDPA device passthrough
- Minimal host CPU involvement
- Requires vDPA-capable hardware

## MAC Addresses

```
-nic ...,mac=52:54:00:XX:XX:XX
```

QEMU default uses the OUI `52:54:00` (reserved for QEMU/KVM). Set explicitly for predictable addressing or guest-visible MAC.

## Multi-Queue

```
-netdev tap,id=net0,queues=N
-device virtio-net-pci,netdev=net0,mq=on,vectors=N
```

- Set `queues=N` on `-netdev` and `mq=on` on `virtio-net-pci`
- `vectors = 2*N + 1` (each TX/RX pair + control)
- Required for RSS/finite queues in guest

## Offloads (virtio-net)

```
-device virtio-net-pci,...,csum=on,gso=on,guest_csum=on,guest_tso4=on,guest_tso6=on,guest_ecn=on,guest_ufo=on,host_tso4=on,host_tso6=on,host_ecn=on,host_ufo=on
```

Control checksum offloading, TSO/GSO, and UFO for virtio-net.

## Filter Drivers

```
-netdev tap,id=net0
-filter-rewriter id=f0,netdev=net0,...   # COLIN filter
-filter-mirror id=f0,netdev=net0,...     # port mirroring
-filter-dump id=f0,netdev=net0,file=dump.pcap  # pcap dump
-filter-redirector id=f0,netdev=net0,...  # redirect traffic
```

## Examples

User-mode with SSH forwarding:
```
qemu-system-x86_64 -nic user,hostfwd=tcp::2222-:22 vmdisk.qcow2
```

TAP bridge (requires setup):
```
qemu-system-x86_64 -netdev tap,id=net0 -device virtio-net-pci,netdev=net0 vmdisk.qcow2
```

Multi-queue with vhost:
```
qemu-system-x86_64 -netdev tap,id=net0,queues=4,vhost=on \
  -device virtio-net-pci,netdev=net0,mq=on,vectors=9 vmdisk.qcow2
```

L2TPv3 tunnel:
```
qemu-system-x86_64 -netdev l2tpv3,id=net0,src=10.0.0.1,dst=10.0.0.2,udp=on \
  -device virtio-net-pci,netdev=net0 vmdisk.qcow2
```
