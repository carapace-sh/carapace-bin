# qemu-system-ARCH

The full-system emulators. One binary per target architecture, same option interface across all.

## Synopsis

```
qemu-system-x86_64 [options] [disk_image]
```

## Minimal VM

```
qemu-system-x86_64 \
  -accel kvm \
  -machine q35 \
  -cpu host \
  -smp 4 \
  -m 4096 \
  -drive file=vmdisk.qcow2,if=virtio \
  -nic user
```

## Option Categories

### Machine and Accelerator

| Option | Description |
|--------|-------------|
| `-machine [type=]NAME[,prop=val[,...]]` | Machine type (e.g. `pc-q35-9.0`, `virt`, `microvm`). Use `-machine help` to list. Properties: `accel=KVM:TCG`, `vmport=on|off|auto`, `dump-guest-core`, `mem-merge`, `nvdimm`, `hmat`, `memory-encryption`, `cxl-fmw.X`, etc. |
| `-accel NAME[,prop=val[,...]]` | Accelerator: `kvm`, `tcg`, `hvf`, `whpx`, `nvmm`, `nitro`. Colon-separated for fallback: `kvm:tcg`. Properties: `igd-passthru=on`, `split-wx=on`, `one-insn-per-tb=on`, `tb-size=N`, `thread=single|multi`, `dirty-ring-size=N`. |
| `-enable-kvm` | Legacy shorthand for `-accel kvm` |
| `-M` | Short form of `-machine` |
| `-smp [[cpus=]N][,maxcpus=N][,sockets=N][,dies=N][,clusters=N][,cores=N][,threads=N]` | SMP topology. Required: at least `cpus=N`. |
| `-m [size=]N[M|G|T]` | RAM size. Suffix defaults to MiB. |
| `-cpu MODEL[,FEATURE[=on|off][,...]]` | CPU model and feature flags. Use `-cpu help` to list models. |

### Memory and NUMA

| Option | Description |
|--------|-------------|
| `-m 8G` | 8 GiB RAM |
| `-m 8G,slots=4,maxmem=32G` | With hotplug slots |
| `-numa node[,memdev=ID][,cpus=RANGE]` | NUMA topology |
| `-object memory-backend-ram,id=ID,size=N` | Explicit memory backend |
| `-object memory-backend-file,id=ID,size=N,mem-path=PATH` | File-backed / hugetlbfs |

### Firmware and Booting

| Option | Description |
|--------|-------------|
| `-bios FILE` | Firmware/BIOS image |
| `-kernel FILE` | Linux kernel image (direct boot) |
| `-initrd FILE` | Initrd for direct kernel boot |
| `-append STRING` | Kernel command line |
| `-dtb FILE` | Device tree blob |
| `-fw_cfg [name=]NAME,file=FILE\|string=VAL` | QEMU firmware config injection |
| `-smbios file=BINARY` | SMBIOS entry injection |

### Snapshots and State

| Option | Description |
|--------|-------------|
| `-snapshot` | Write to temporary files; discard on exit |
| `-S` | Freeze CPU at startup (wait for `continue` via QMP/HMP or `-gdb`) |
| `-realtime mlock=on\|off` | Lock memory (legacy, use `-overcommit mem-lock=on`) |
| `-overcommit mem-lock=on\|off,cpu-pm=on\|off` | Memory locking and CPU power management hints |
| `-no-reboot` | Exit on guest shutdown instead of rebooting |
| `-no-shutdown` | Don't exit on guest shutdown (wait for QMP) |
| `-daemonize` | Daemonize after startup |
| `-name STRING` | VM name (visible via QMP) |
| `-uuid UUID` | VM UUID |

### RTC, Timers, Watchdog

| Option | Description |
|--------|-------------|
| `-rtc [base=utc|localtime|DATE][,clock=host|rt|vm][,driftfix=none|slew]` | Real-time clock configuration |
| `-icount [shift=N][,align=on|off][,sleep=on|off][,rr=record|replay][,rrfile=FILE][,rrsnapshot=NAME]` | Instruction counting for record/replay |
| `-watchdog MODEL[,action=reset|shutdown|poweroff|pause|debug|none|nmi]` | Watchdog device and action |

### Debug and Expert

| Option | Description |
|--------|-------------|
| `-gdb dev` | Accept GDB connection on `dev` (e.g. `tcp::1234`) |
| `-d ITEMS` | Enable logging of specific items (`cpu_errors`, `in_asm`, `int`, `exec`, etc.) |
| `-D FILE` | Log to file |
| `-L PATH` | BIOS/VGA BIOS directory |
| `-msg [timestamp=on|off][,guest-name=on|off]` | Control message format |
| `-trace [[enable=]PATTERN][,events=FILE][,file=FILE]` | Trace events |
| `-readconfig FILE` | Read config from file |
| `-writeconfig FILE` | Write config to file |
| `-no-acpi` | Disable ACPI |
| `-no-hpet` | Disable HPET |
| `-acpitable [sig=...][,rev=N][,oem_id=...][,oem_table_id=...][,oem_rev=N][,asl_compiler_id=...][,asl_compiler_rev=N][,data=FILE[:N]]` | ACPI table injection |
| `-semihosting` | Enable semihosting |
| `-semihosting-config [enable=on|off][,target=native|gdb|auto][,arg=STR[,...]]` | Semihosting configuration |

### Direct Boot (Linux)

For booting a Linux kernel directly without firmware:

```
qemu-system-x86_64 \
  -kernel vmlinuz \
  -initrd initramfs.img \
  -append "root=/dev/vda1 console=ttyS0" \
  -drive file=rootfs.qcow2,if=virtio \
  -nographic
```

This bypasses the BIOS/firmware entirely — `-bios` is not needed.

### Machine Properties Reference

Common machine properties (set via `-machine NAME,PROP=VAL`):

| Property | Applies to | Values |
|----------|-----------|--------|
| `accel` | All | Colon-separated accelerator list |
| `vmport` | x86 | `on`, `off`, `auto` |
| `nvdimm` | x86 | `on`, `off` |
| `hmat` | x86 | `on`, `off` (ACPI HMAT) |
| `memory-encryption` | x86 | ID of `sev-guest` or `tdx-guest` object |
| `memory-backend` | All | ID of memory backend object |
| `cxl-fmw.X.targets.Y` | x86 | CXL fixed memory window targets |
| `cxl-fmw.X.size` | x86 | CXL FMW size |
| `cxl-fmw.X.interleave-granularity` | x86 | CXL interleave granularity (512|256|1k|2k|4k|8k|16k|32k|64k) |
| `ctl-version` | s390x | `1` or `2` |
| `x-south-bridge` | x86 | `PIIX3`, `piix4-isa` |
| `aux-ram-share` | All | `on`, `off` (for cpr-transfer migration) |
