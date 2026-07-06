# Guest Agent (qemu-ga)

In-guest daemon that provides host↔guest communication. Installed inside the VM.

## Architecture

```
Host (QEMU)  ── virtio-serial / ISA serial / viosock ──►  Guest Agent
```

The host sends QGA (QEMU Guest Agent Protocol) commands via a virtio-serial channel. The guest agent executes them and returns results.

## Enabling the Channel

1. Add a virtio-serial device to the VM:
```
-chardev socket,path=/tmp/qga.sock,server=on,wait=off,id=qga0
-device virtio-serial-pci
-device virtserialport,chardev=qga0,name=org.qemu.guest_agent.0
```

2. Install and start `qemu-guest-agent` inside the guest:
```bash
# apt install qemu-guest-agent   # Debian/Ubuntu
# systemctl enable --now qemu-guest-agent
```

3. On the host, connect via `socat` or `qmp-shell`:
```
socat unix-connect:/tmp/qga.sock -
```

## Key QGA Commands

### Initialization

| Command | Description |
|---------|-------------|
| `guest-sync-delimited` | Sync stream with sentinel byte (0xFF) |
| `guest-sync` | Simple sync, returns echoed integer |

### System Information

| Command | Description |
|---------|-------------|
| `guest-info` | List supported commands |
| `guest-get-time` | Guest time (nanoseconds) |
| `guest-set-time` | Sync guest time (requires `-rtc base=utc`) |
| `guest-get-osinfo` | OS name, version, kernel |
| `guest-get-host-name` | Hostname |
| `guest-get-users` | Active user sessions |
| `guest-get-timezone` | Timezone info |
| `guest-get-vcpus` | VCPU information |

### Network

| Command | Description |
|---------|-------------|
| `guest-network-get-interfaces` | Interface name, MAC, IP addresses |

### Files and Execution

| Command | Description |
|---------|-------------|
| `guest-exec` | Execute command |
| `guest-exec-status` | Check exit status |
| `guest-file-open` | Open file (returns handle) |
| `guest-file-read` | Read from file handle |
| `guest-file-write` | Write to file handle |
| `guest-file-close` | Close file handle |
| `guest-file-seek` | Seek in file |

### FSFREEZE

| Command | Description |
|---------|-------------|
| `guest-fsfreeze-freeze` | Freeze filesystems |
| `guest-fsfreeze-thaw` | Thaw filesystems |
| `guest-fsfreeze-status` | Query freeze state |

### Shutdown and Reboot

| Command | Description |
|---------|-------------|
| `guest-shutdown` | `mode=powerdown|halt|reboot` |
| `guest-reboot` | Reboot guest |

### Disk and FS

| Command | Description |
|---------|-------------|
| `guest-get-disks` | List block devices |
| `guest-get-fsinfo` | Filesystem info |
| `guest-get-memory-block-info` | Memory block info |
| `guest-set-memory-blocks` | Online/offline memory blocks |

### SUSPEND

| Command | Description |
|---------|-------------|
| `guest-suspend-ram` | Suspend to RAM |
| `guest-suspend-disk` | Hibernate |
| `guest-suspend-hybrid` | Hybrid suspend |

## QGA Commands in QMP

Guest agent commands are accessed through the `guest-*` prefix via QMP, *not* directly:

```
{ "execute": "guest-exec", "arguments": { "path": "/bin/echo", "arg": ["hello"] } }
```

Or via `guest-agent-*` (depending on QEMU version):

```
{ "execute": "guest-agent-exec", "arguments": { "path": "/bin/echo", "arg": ["hello"] } }
```

## Security

- Guest agent runs as root inside the VM
- The host controls what commands are sent
- Use agent with `-device virtserialport,chardev=qga0,name=org.qemu.guest_agent.0` — the `name` must match what the agent expects
- Some commands can be restricted (e.g., `guest-exec`, `guest-file-open`)
- QEMU can be configured to only allow specific commands

## Use Cases

- **Backup integration**: `guest-fsfreeze-freeze` ensures consistent disk state before snapshot
- **Guest information**: Query IP address, OS version, uptime without guest network
- **Graceful shutdown**: `guest-shutdown` for clean VM shutdown without ACPI
- **File injection**: Transfer small files with `guest-file-open/write/close`
- **Orchestration**: Ansible, Terraform, libvirt use guest agent for provisioning

## See Also

- `qemu-ga-ref(7)` — complete QGA protocol reference
- `qemu-qmp-ref(7)` — QMP command reference (host side)
