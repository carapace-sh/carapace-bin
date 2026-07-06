# Block Drivers — Disk Image Formats and Protocols

QEMU supports a wide range of disk image formats and network storage protocols. These can be used with `qemu-img`, `qemu-system-*`, `qemu-nbd`, and `qemu-storage-daemon`.

## Image File Formats

### raw

Raw disk image. Simplest format, no metadata overhead.

- Supported options: `preallocation` (`off`, `falloc`, `full`)
- Host filesystems with holes (ext4, NTFS) only allocate space for written sectors
- `ls -ls` shows actual disk usage

### qcow2

QEMU Copy-On-Write version 2. The most versatile format. Default for QEMU.

| Feature | Description |
|---------|-------------|
| Sparse allocation | Only allocate clusters on write |
| Compression | zlib-based, set `-c` on convert or `qcow2-compression-type=zlib|zstd` |
| Encryption | Native AES-CBC (deprecated) or LUKS |
| Snapshots | Internal snapshots (differ from backing files) |
| Backing files | Copy-on-write overlay on a base image |
| Cluster size | Configurable (default 64 KiB, options: 512B–2M, power of 2) |
| Lazy refcounts | Performance optimization for write-heavy workloads |
| Data file | External data file (separates metadata from data) |
| Subclusters | Allocation tracking at subcluster granularity |
| Dirty bitmaps | Persistent bitmaps for incremental backup |

Options:
- `compat=0.10|1.1` — qcow2 version (1.1 enables zero clusters, subclusters, data files)
- `backing_file=FMT` — base image filename
- `backing_fmt=FMT` — base image format
- `encryption=on|off` — deprecated, use `encrypt.format=aes|luks`
- `cluster_size=N` — cluster size in bytes
- `preallocation=off|metadata|falloc|full` — preallocation mode
- `lazy_refcounts=on|off` — delayed refcount updates
- `refcount_bits=N` — refcount entry width (16 default)
- `data_file=FMT` — external data file path
- `data_file_raw=on|off` — data file is raw (no qcow2 metadata in it)
- `compression_type=zlib|zstd` — compression algorithm
- `extended_l2=on|off` — enable subclusters (64-bit L2 entries)

### qed

QEMU Enhanced Disk format. Legacy, no longer recommended. Replaced by qcow2 with compat=1.1.

### vmdk

VMware Virtual Disk format. Supported for compatibility with VMware products.
- Options: `compat6`, `subformat`, `adapter_type`, `hwversion`, `backing_file`, `zeroed_grain`

### vhdx

Microsoft Hyper-V Virtual Hard Disk v2 format. Supports dynamic and differencing disks.
- Options: `subformat=dynamic|fixed`, `block_size=N`, `log_size=N`

### vdi

Oracle VirtualBox Virtual Disk Image format.

### qcow

Legacy QEMU Copy-On-Write version 1. Replaced by qcow2.

### vpc

Microsoft Virtual PC / Hyper-V VHD format.

### luks

LUKS encrypted volume (not a disk image format per se, but a driver wrapping raw/block storage).
- Options: `key-secret=ID`

### Parallels

Parallels Desktop disk image format (`.hdd`, `.hds`).

### bochs, cloop, dmg

Read-only formats for Bochs, compressed loop, and Apple DMG images.

## Network Storage Backends

These protocols can be used as block drivers, both for images and direct storage access:

| Driver | Protocol | Typical URI |
|--------|----------|-------------|
| `nbd` | Network Block Device | `nbd:host:port:exportname=NAME` |
| `iscsi` | iSCSI | `iscsi://[user:pass@]host[:port]/target/lun` |
| `iser` | iSCSI over RDMA | Same as iSCSI |
| `nfs` | Network File System | `nfs://host/path` |
| `ssh` | SSH/SFTP | `ssh://[user@]host[:port]/path` |
| `gluster` | GlusterFS | `gluster://host/vol/file` |
| `http` / `https` | HTTP(S) | `http://host/path` (read-only) |
| `ftp` / `ftps` | FTP(S) | `ftp://host/path` (read-only) |
| `nvme` | NVMe over PCIe (local) | Direct NVMe namespace access |

## Filter Drivers

These wrap another block node to add functionality:

| Driver | Purpose |
|--------|---------|
| `blkdebug` | Inject I/O errors for testing |
| `blkverify` | Verify read data matches a reference driver |
| `blklogwrites` | Log writes to a separate device |
| `copy-on-read` | Copy data to backing file on read |
| `copy-before-write` | Copy old data before overwrite (for backup/cbs) |
| `preallocate` | Pre-allocate space on demand |
| `throttle` | Rate-limit I/O |
| `replication` | Active-active/active-passive block replication |
| `compress` | Compress/decompress on the fly |
| `quorum` | Vote across multiple child nodes for read reliability |
| `snapshot-access` | Access internal snapshots as separate block nodes |

## Null Drivers

| Driver | Purpose |
|--------|---------|
| `null-aio` | Discard all I/O, use AIO internally |
| `null-co` | Discard all I/O, use coroutines |

## See Also

- `qemu-block-drivers(7)` — canonical reference
- `qemu-img --help` — format list at runtime
