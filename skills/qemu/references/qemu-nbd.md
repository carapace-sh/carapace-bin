# qemu-nbd — Network Block Device Server

Export disk images over the network using the NBD protocol, or connect to an existing NBD server.

## Synopsis

```
qemu-nbd [OPTIONS] FILE
qemu-nbd -L [OPTIONS]
```

## Modes

**Server mode**: Export a local disk image via NBD. Default port is 10809.

**List mode** (`-L`): Query an existing NBD server and list its exports.

**Kernel client mode** (`-c`/`-d`): Connect/disconnect a kernel NBD device.

## Connection Options

| Option | Description |
|--------|-------------|
| `-p, --port PORT` | Listen port (default `10809`) |
| `-b, --bind IFACE` | Interface to bind to (default `0.0.0.0`) |
| `-k, --socket PATH` | UNIX socket path (default `/var/lock/qemu-nbd-DEVICE`) |
| `-e, --shared NUM` | Number of clients (default `1`) |
| `-t, --persistent` | Don't exit after last connection |
| `-v, --verbose` | Debug output |
| `-x, --export-name NAME` | Export by name (default empty string) |
| `-D, --description TEXT` | Human-readable export description |
| `--handshake-limit N` | Client handshake timeout in seconds (default 10) |
| `--fork` | Fork server to background |
| `--pid-file PATH` | Write PID to file |

## Image Exposure

| Option | Description |
|--------|-------------|
| `-o, --offset OFFSET` | Expose image starting at byte offset |
| `-A, --allocation-depth` | Expose allocation depth (thin provisioning info) |
| `-B, --bitmap NAME` | Expose a persistent dirty bitmap |

## TLS and Authorization

| Option | Description |
|--------|-------------|
| `--object type,id=ID,...` | Define objects (secrets, TLS creds) |
| `--tls-creds ID` | TLS server credentials (use `--object tls-creds-x509`) |
| `--tls-authz ID` | Authorization object (use `--object authz-*`) |
| `--tls-hostname HOSTNAME` | Override hostname for x509 certificate check (client mode) |

## Kernel NBD Client

| Option | Description |
|--------|-------------|
| `-c, --connect DEV` | Connect FILE to kernel NBD device `DEV` (e.g. `/dev/nbd0`) |
| `-d, --disconnect` | Disconnect the specified device |

Kernel module must be loaded: `modprobe nbd`.

## Block Device Options

| Option | Description |
|--------|-------------|
| `-f, --format FMT` | Image format (`raw`, `qcow2`, etc.) |
| `-r, --read-only` | Export read-only |
| `-s, --snapshot` | Use FILE as backing for a temporary overlay |
| `-l, --load-snapshot PARAM` | Load internal snapshot (format: `snapshot.id=ID,snapshot.name=NAME` or `ID_OR_NAME`) |
| `-n, --nocache` | Disable host cache |
| `--cache MODE` | Cache mode: `none`, `writeback` (default), `writethrough`, `directsync`, `unsafe` |
| `--aio MODE` | AIO mode: `native`, `io_uring`, `threads` |
| `--discard MODE` | Discard: `ignore`, `unmap` |
| `--detect-zeroes MODE` | `off`, `on`, `unmap` |
| `--image-opts` | Treat FILE as full option string |

## Examples

Export a qcow2 image read-only:
```
qemu-nbd -f qcow2 -r vmdisk.qcow2
```

Export with persistent mode and multiple clients:
```
qemu-nbd -t -e 4 -f qcow2 vmdisk.qcow2
```

Connect to kernel NBD device:
```
sudo modprobe nbd
sudo qemu-nbd -c /dev/nbd0 vmdisk.qcow2
sudo mount /dev/nbd0 /mnt
sudo qemu-nbd -d /dev/nbd0
```

List exports on an NBD server:
```
qemu-nbd -L server.example.com
```

UNIX socket export:
```
qemu-nbd -k /tmp/nbd.sock -t vmdisk.qcow2
```

TLS-encrypted export:
```
qemu-nbd --object tls-creds-x509,id=tls0,endpoint=server,dir=/path/to/certs \
  --tls-creds tls0 -t vmdisk.qcow2
```
