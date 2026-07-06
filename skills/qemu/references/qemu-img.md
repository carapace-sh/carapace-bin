# qemu-img — Disk Image Utility

Offline tool for creating, converting, inspecting, and modifying disk images. Never use on images attached to a running VM.

## Synopsis

```
qemu-img [standard options] COMMAND [command options]
```

## Standard Options

| Option | Description |
|--------|-------------|
| `-h, --help` | Display help |
| `-V, --version` | Display version |
| `-T, --trace [...]` | Tracing options |

## Commands

### `create` — Create a new image

```
qemu-img create [-f FMT] [-b BACKING_FILE [-F BACKING_FMT]] [-u] [-o OPTIONS] FILENAME [SIZE]
```

| Flag | Description |
|------|-------------|
| `-f FMT` | Image format (default: `qcow2`) |
| `-b FILE` | Backing file (for overlay images) |
| `-F FMT` | Backing file format |
| `-o OPTIONS` | Format-specific options (comma-separated `key=val`) |
| `-u` | Unsafe: allow unsafe backing file (don't verify) |
| `-q` | Quiet mode |
| `SIZE` | Size with suffix (optional for `qcow2` with backing file) |

Examples:
```
qemu-img create -f qcow2 vmdisk.qcow2 50G
qemu-img create -f qcow2 -b base.qcow2 -F qcow2 overlay.qcow2
qemu-img create -f raw -o preallocation=falloc bigdisk.raw 100G
```

### `convert` — Convert image format

```
qemu-img convert [-c] [-p] [-q] [-n] [-f FMT] [-O OUTPUT_FMT] [-B BACKING_FILE] [-o OPTIONS] [-S SPARSE_SIZE] [-m COROUTINES] [-W] FILENAME [FILENAME2 [...]] OUTPUT_FILENAME
```

| Flag | Description |
|------|-------------|
| `-c` | Compress (qcow2/zlib) |
| `-O FMT` | Output format |
| `-B FILE` | Output backing file |
| `-S SIZE` | Sparse size (sectors smaller than this are written as holes) |
| `-m N` | Number of parallel coroutines |
| `-W` | Copy during write |
| `-p` | Show progress |
| `-n` | Skip creation of target (target must already exist) |
| `--target-is-zero` | Assume target is already zero-filled |
| `--bitmaps` | Copy persistent bitmaps |
| `--salvage` | Try to continue after I/O errors |

Examples:
```
qemu-img convert -f qcow2 -O raw vmdisk.qcow2 vmdisk.raw
qemu-img convert -c -O qcow2 -p large.qcow2 compressed.qcow2
```

### `check` — Verify image integrity

```
qemu-img check [-f FMT] [--output=OFMT] [-r [leaks | all]] [-T CACHE] [-U] FILENAME
```

| Flag | Description |
|------|-------------|
| `-r MODE` | Repair mode: `leaks` (fix leaks only) or `all` |
| `-U` | Force-share (allow opening in use by running VM) |
| `--output=json` | Machine-readable JSON output |

### `info` — Show image information

```
qemu-img info [-f FMT] [--backing-chain] [--limits] [-U] FILENAME
```

| Flag | Description |
|------|-------------|
| `--backing-chain` | Show info for entire backing chain |
| `--limits` | Show format limits |
| `-U` | Force-share |
| `--output=json` | JSON output |

### `snapshot` — List/manage snapshots

```
qemu-img snapshot [-l | -a SNAPSHOT | -c SNAPSHOT | -d SNAPSHOT] FILENAME
```

| Flag | Description |
|------|-------------|
| `-l` | List snapshots |
| `-c NAME` | Create snapshot |
| `-a NAME` | Apply/restore snapshot |
| `-d NAME` | Delete snapshot |

### `commit` — Commit overlays to backing file

```
qemu-img commit [-b BASE] [-d] [-p] [-r RATE_LIMIT] FILENAME
```

| Flag | Description |
|------|-------------|
| `-b BASE` | Commit only to this backing file level |
| `-d` | Delete the overlay after committing |
| `-p` | Show progress |
| `-r BYTES/s` | Rate limit |

### `rebase` — Change backing file

```
qemu-img rebase [-u] [-c] [-b BACKING_FILE] [-F BACKING_FMT] [-p] FILENAME
```

| Flag | Description |
|------|-------------|
| `-b FILE` | New backing file |
| `-F FMT` | New backing file format |
| `-u` | Unsafe: update backing file reference only (no data rewrite) |
| `-c` | Compress (must also use `-b`) |

### `resize` — Resize an image

```
qemu-img resize [--preallocation=PREALLOC] [--shrink] FILENAME [+|-]SIZE
```

| Flag | Description |
|------|-------------|
| `--preallocation` | `off`, `falloc`, `full` |
| `--shrink` | Allow shrinking (data loss risk) |
| `+SIZE` | Grow by SIZE |
| `-SIZE` | Shrink by SIZE |

### `amend` — Modify format-specific options

```
qemu-img amend [-o OPTIONS] [-f FMT] [-p] [-q] FILENAME
```

Used to change image metadata (e.g. qcow2 compat level, encryption, cluster size).

### `bench` — Benchmark image performance

```
qemu-img bench [-c COUNT] [-d DEPTH] [-f FMT] [-i AIO] [-n] [-o OFFSET] [-s SIZE] [-S STEP] [-t CACHE] [-w] FILENAME
```

| Flag | Description |
|------|-------------|
| `-c N` | I/O count per iteration |
| `-d N` | AIO queue depth |
| `-i MODE` | AIO mode: `threads`, `native`, `io_uring` |
| `-s SIZE` | I/O buffer size |
| `-w` | Write test (default: read) |

### `dd` — Raw copy with format conversion

```
qemu-img dd [-f FMT] [-O FMT] [bs=BLOCK_SIZE] [count=BLOCKS] [skip=BLOCKS] if=INPUT of=OUTPUT
```

Same semantics as `dd(1)` but with QEMU I/O path and format support.

### `compare` — Compare two images

```
qemu-img compare [-f FMT] [-F FMT] [-s] [-p] [-q] [-U] FILENAME1 FILENAME2
```

| Flag | Description |
|------|-------------|
| `-s` | Strict (compare all data including unallocated) |
| `-p` | Show progress |

### `map` — Dump image metadata

```
qemu-img map [-f FMT] [--start-offset=OFFSET] [--max-length=LEN] [--output=OFMT] [-U] FILENAME
```

Shows which sectors are allocated. Useful for understanding image sparseness.

### `measure` — Calculate required size

```
qemu-img measure [-O OUTPUT_FMT] [-o OPTIONS] [--size N | [-f FMT] [-l SNAPSHOT] FILENAME]
```

Estimates the file size needed for a new image or conversion.

### `bitmap` — Manipulate persistent bitmaps

```
qemu-img bitmap (--merge SOURCE | --add | --remove | --clear | --enable | --disable)... [-g GRANULARITY] [-b SOURCE_FILE] FILENAME BITMAP
```

| Flag | Description |
|------|-------------|
| `--add` | Create a new persistent dirty bitmap |
| `--remove` | Remove a bitmap |
| `--clear` | Clear a bitmap |
| `--enable` / `--disable` | Enable/disable tracking |
| `--merge SOURCE` | Merge another bitmap into this one |
| `-g N` | Granularity in bytes (power of 2, default 65536) |
| `-b FILE` | Source file for merge |

## Image Formats

List formats with `qemu-img --help`:

`blkdebug`, `blklogwrites`, `blkverify`, `bochs`, `cloop`, `compress`, `copy-before-write`, `copy-on-read`, `dmg`, `file`, `ftp`, `ftps`, `gluster`, `host_cdrom`, `host_device`, `http`, `https`, `iscsi`, `iser`, `luks`, `nbd`, `nfs`, `null-aio`, `null-co`, `nvme`, `parallels`, `preallocate`, `qcow`, `qcow2`, `qed`, `quorum`, `raw`, `replication`, `snapshot-access`, `ssh`, `throttle`, `vdi`, `vhdx`, `vmdk`, `vpc`, `vvfat`

See [block-drivers.md](block-drivers.md) for detailed format descriptions.
