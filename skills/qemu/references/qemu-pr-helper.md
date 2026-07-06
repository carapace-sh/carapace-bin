# qemu-pr-helper — SCSI Persistent Reservation Helper

SCSI persistent reservations restrict access to block devices to specific initiators in shared storage setups. This helper daemon handles PR commands from QEMU's SCSI passthrough devices (`scsi-block`, `scsi-generic`).

## Synopsis

```
qemu-pr-helper [OPTION]
```

## How It Works

1. QEMU's `scsi-block` or `scsi-generic` device receives a PR command from the guest
2. Instead of executing it directly (which requires privileges), QEMU forwards it to `qemu-pr-helper` over a UNIX socket
3. `qemu-pr-helper` executes the PR command on behalf of QEMU
4. Returns the result to QEMU, which passes it to the guest

This allows unprivileged QEMU processes to use SCSI persistent reservations while the helper runs with elevated privileges.

## Options

| Option | Description |
|--------|-------------|
| `-d, --daemon` | Run in background (creates PID file) |
| `-q, --quiet` | Decrease verbosity |
| `-v, --verbose` | Increase verbosity |
| `-f, --pidfile PATH` | PID file (default: `/var/run/qemu-pr-helper.pid`) |
| `-k, --socket PATH` | Socket path (default: `/var/run/qemu-pr-helper.sock`) |
| `-u, --user USER` | Drop privileges to this user |
| `-g, --group GROUP` | Drop privileges to this group |
| `-T, --trace [...]` | Tracing options |
| `-h, --help` | Display help |
| `-V, --version` | Display version |

## systemd Integration

Supports socket activation. Example socket unit:

```
[Socket]
ListenStream=/var/run/qemu-pr-helper.sock
```

## QEMU Configuration

In the QEMU command line, configure the PR manager on the block device:

```
-object pr-manager-helper,id=pr0,path=/var/run/qemu-pr-helper.sock \
-drive file=shared.qcow2,if=none,id=drive0,format=qcow2 \
-device scsi-block,drive=drive0,pr-manager=pr0
```

## Use Case

Shared SCSI storage in a cluster of virtual machines. Each VM needs to coordinate access to the same LUN. Persistent reservations provide the cluster fence mechanism — when a node fails, the cluster manager releases its PR and another node takes ownership.
