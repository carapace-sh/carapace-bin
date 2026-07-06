# QMP — QEMU Monitor Protocol

JSON-based protocol for controlling QEMU programmatically. The primary management interface for QEMU.

## Enabling QMP

```
-qmp tcp:localhost:4444,server,nowait    # TCP
-qmp unix:/tmp/qmp.sock,server,nowait    # UNIX socket
```

Also via `-chardev` + `-mon`:

```
-chardev socket,id=qmp0,path=/tmp/qmp.sock,server=on,wait=off
-mon chardev=qmp0,mode=control
```

## HMP (Human Monitor Protocol)

Legacy text-based monitor:

```
-hmp tcp:localhost:4445,server,nowait
-hmp unix:/tmp/hmp.sock,server,nowait
```

Can also be accessed from QMP with the `human-monitor-command` command.

## Protocol Basics

QMP is JSON-over-TCP. Messages are newline-delimited JSON objects.

### Handshake

```
← {"QMP": {"version": {"qemu": {"major": 11, "minor": 0, "micro": 1}, "package": ""}, "capabilities": []}}
→ { "execute": "qmp_capabilities" }
← { "return": {} }
```

### Command Structure

```
{ "execute": "COMMAND", "arguments": { ... }, "id": "TAG" }
```

- `execute` — command name (required)
- `arguments` — command arguments (optional)
- `id` — request identifier, echoed in response (optional)

### Response Structure

Success:
```
{ "return": VALUE, "id": "TAG" }
```

Error:
```
{ "error": { "class": "ERROR_CLASS", "desc": "Human-readable description" }, "id": "TAG" }
```

Events (asynchronous):
```
{ "event": "EVENT_NAME", "data": { ... }, "timestamp": { "seconds": N, "microseconds": N } }
```

## Key Commands

### VM Lifecycle

| Command | Description |
|---------|-------------|
| `stop` | Pause VM |
| `cont` | Resume VM |
| `system_powerdown` | ACPI power button |
| `system_reset` | Reset VM |
| `system_wakeup` | Wake from S3 |
| `quit` | Shut down QEMU |
| `human-monitor-command` | Execute HMP command |

### State Queries

| Command | Description |
|---------|-------------|
| `query-kvm` | KVM status |
| `query-status` | VM run state |
| `query-name` | VM name |
| `query-uuid` | VM UUID |
| `query-machines` | List machine types |
| `query-cpus-fast` | CPU information |
| `query-block` | Block device info |
| `query-blockstats` | Block I/O statistics |
| `query-netdev` | Netdev info |
| `query-migrate` | Migration status |
| `query-commands` | All available commands |
| `qom-list` | QOM object tree |
| `qom-get` | QOM property value |

### Block Operations

| Command | Description |
|---------|-------------|
| `blockdev-snapshot-sync` | Create snapshot (legacy) |
| `blockdev-snapshot` | Take snapshot (new API) |
| `blockdev-backup` | Backup job |
| `blockdev-mirror` | Mirror job |
| `block-commit` | Commit overlay chain |
| `block-stream` | Stream data from backing file |
| `block-job-cancel` | Cancel active job |
| `block-job-pause` / `block-job-resume` | Job control |
| `blockdev-add` / `blockdev-del` | Dynamic block node add/remove |
| `blockdev-create` | Create image via QMP |
| `transaction` | Atomic batch of block operations |
| `x-blockdev-amend` | Modify image format options |
| `query-jobs` | List active jobs |
| `nbd-server-start` | Start NBD server |
| `nbd-server-add` / `nbd-server-remove` | Dynamic NBD export management |
| `nbd-server-stop` | Stop NBD server |

### Migration

| Command | Description |
|---------|-------------|
| `migrate` | Start migration (`uri` argument) |
| `migrate_cancel` | Cancel migration |
| `migrate-incoming` | Set migration target for incoming (`-incoming defer`) |
| `migrate-set-parameters` | Set migration tuning |
| `migrate-start-postcopy` | Switch to post-copy |
| `migrate-continue` | Continue after pause |
| `query-migrate` | Migration status |
| `query-migrate-parameters` | Migration parameters |

### Device Hotplug

| Command | Description |
|---------|-------------|
| `device_add` | Add device (QOM path, id, properties) |
| `device_del` | Remove device by ID |
| `netdev_add` / `netdev_del` | Network backend hotplug |
| `chardev-add` / `chardev-remove` | Character device hotplug |
| `object-add` / `object-del` | QOM object lifecycle |

### QOM (QEMU Object Model)

```
qom-list path=/machine/peripheral
qom-get path=/machine/peripheral/virtio0,property=drive
qom-set path=/machine/peripheral/virtio0,property=drive,value=disk0
```

### Dump and Snapshot

```
dump-guest-memory paging=false protocol=file:/tmp/dump
pmemsave val=... size=... filename=...
```

### Set Password

```
{ "execute": "set_password", "arguments": { "protocol": "vnc", "password": "mypass" } }
{ "execute": "expire_password", "arguments": { "protocol": "vnc", "time": "+10m" } }
```

## Key Events

QMP events the management layer subscribes to:

| Event | Trigger |
|-------|---------|
| `SHUTDOWN` | Guest shutdown |
| `RESET` | Guest reset |
| `STOP` | VM paused |
| `RESUME` | VM resumed |
| `SUSPEND` | Guest suspended |
| `RTC_CHANGE` | RTC time changed |
| `WATCHDOG` | Watchdog fired |
| `BLOCK_IO_ERROR` | Block I/O error |
| `BLOCK_JOB_COMPLETED` | Block job finished |
| `BLOCK_JOB_READY` | Block job ready (ready bit) |
| `BLOCK_JOB_PENDING` | Block job pending |
| `DEVICE_DELETED` | Device removed |
| `NIC_RX_FILTER_CHANGED` | Guest changed MAC |
| `MIGRATION` | Migration state change |
| `MIGRATION_PASS` | Migration pass |
| `VSERPORT_CHANGE` | virtio-serial port change |
| `NVDIMM_REPLACEMENT` | NVDIMM replaced |
| `MEMORY_DEVICE_SIZE_CHANGE` | Memory device size changed |
| `SUSPEND_DISK` | Guest hibernate |

## Tools

**qmp-shell** (in QEMU source tree):
```
python3 scripts/qmp/qmp-shell /tmp/qmp.sock
```

**qmp-proxy** and **qemu-vm** for programmatic Python use via `qemu.qmp` package:
```python
from qemu.qmp import QMPClient
qmp = QMPClient()
await qmp.connect("/tmp/qmp.sock")
await qmp.execute("query-status")
await qmp.disconnect()
```

## JSON Formatting

For constructing QMP commands and parsing responses programmatically, use `jq`. See the `jq` skill (`skills/jq/`).

## See Also

- `qemu-qmp-ref(7)` — complete command and event reference
- `qemu-storage-daemon-qmp-ref(7)` — storage daemon QMP subset
