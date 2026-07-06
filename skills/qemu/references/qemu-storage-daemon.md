# qemu-storage-daemon — Standalone Block Storage Daemon

Long-lived process that provides disk image functionality without running a VM. Controlled via QMP commands.

## Synopsis

```
qemu-storage-daemon [options]
```

## Overview

The daemon provides a subset of QEMU block features:

- Block nodes (`--blockdev`)
- Block jobs (commit, mirror, backup)
- Block exports (NBD, vhost-user-blk, FUSE, vduse-blk)
- Throttle groups
- Character devices (`--chardev`)
- Crypto and secrets (`--object`)
- QMP monitor
- IOThreads

Controlled via `--monitor ... mode=control` (there is no `--qmp` flag). Shutdown with `quit` QMP command or SIGINT/SIGHUP/SIGTERM.

## Command-Line Options

| Option | Description |
|--------|-------------|
| `-h, --help` | Display help |
| `-V, --version` | Display version |
| `-T, --trace [...]` | Tracing options |
| `--blockdev BLOCKDEVDEF` | Block node definition (same syntax as qemu-system `-blockdev`) |
| `--chardev CHARDEVDEF` | Character device definition |
| `--export EXPORTDEF` | Block export definition (NBD, vhost-user-blk, FUSE, vduse-blk) |
| `--monitor MONITORDEF` | QMP monitor configuration (use `mode=control`; e.g. `chardev=char0,mode=control`) |
| `--nbd-server addr.type=unix,addr.path=PATH` or `addr.type=inet,...` | Configure NBD server for dynamic exports |
| `--object OBJECTDEF` | QEMU user-creatable object (secrets, crypto, throttlegroups, iothreads) |
| `--pidfile PATH` | Write PID to file |
| `--daemonize` | Fork to background |

## Export Types

### NBD Export

```
--export [type=]nbd,id=ID,node-name=NODE[,name=EXPORT_NAME][,writable=on|off][,bitmap=BITMAP]
```

Requires `--nbd-server` to be set first.

### vhost-user-blk Export

```
--export [type=]vhost-user-blk,id=ID,node-name=NODE,addr.type=unix,addr.path=SOCKET[,writable=on|off][,logical-block-size=BYTES][,num-queues=N]
```

Accepts connections over a vhost-user UNIX socket.

### FUSE Export

```
--export [type=]fuse,id=ID,node-name=NODE,mountpoint=FILE[,growable=on|off][,writable=on|off][,allow-other=on|off|auto]
```

Exports a block node as a regular file via FUSE.

### vduse-blk Export

```
--export [type=]vduse-blk,id=ID,node-name=NODE,name=VDUSE_NAME[,writable=on|off][,num-queues=N][,queue-size=QS][,logical-block-size=BYTES][,serial=SERIAL]
```

Creates a VDUSE device; must be added to vDPA bus with `vdpa(8)`.

## Examples

Simple NBD export:
```
qemu-storage-daemon \
  --blockdev file,node-name=file,filename=vmdisk.qcow2 \
  --blockdev qcow2,node-name=disk,file=file \
  --nbd-server addr.type=unix,addr.path=/tmp/nbd.sock \
  --export nbd,id=export0,node-name=disk,writable=on \
  --monitor chardev=mon0,mode=control,pretty=on
```

(Note: there is no `--qmp` flag; use `--monitor ... mode=control`)`

FUSE export (mount disk image as a file):
```
qemu-storage-daemon \
  --blockdev file,node-name=file,filename=vmdisk.qcow2 \
  --blockdev qcow2,node-name=disk,file=file \
  --export fuse,id=export0,node-name=disk,mountpoint=/mnt/disk.raw,writable=on
```

With IOThread:
```
qemu-storage-daemon \
  --object iothread,id=iothread0 \
  --blockdev file,node-name=file,filename=vmdisk.qcow2,aio=io_uring,iothread=iothread0 \
  --blockdev qcow2,node-name=disk,file=file \
  --export nbd,id=export0,node-name=disk \
  --nbd-server addr.type=unix,addr.path=/tmp/nbd.sock
```

## QMP Commands

The daemon supports these QMP command categories:

- `block-*` — block job management (commit, mirror, backup, create, resize)
- `blockdev-*` — dynamic block graph changes (add/remove nodes, create/destroy)
- `query-*` — information queries
- `nbd-server-*` — dynamic NBD server/add/remove exports
- `qmp_capabilities` — QMP handshake
- `quit` — shutdown

See `qemu-storage-daemon-qmp-ref(7)` for the full list.
