# Plan: Add Windows Completers

Track progress for adding Windows completers to carapace-bin.

## Context

- Existing windows completers: `winget` (1 total)
- Existing windows completers in common: `wt` (Windows Terminal)
- Resolution priority on windows: `windows` > `common` > `bridge`
- Windows completers go in `completers/windows/<cmd>_completer/` (no build tags needed)
- Windows support in shared actions uses `runtime.GOOS` switches (not build tags), following the darwin pattern in `pkg/actions/fs/blockdevice.go`, `pkg/actions/fs/mount.go`, `pkg/actions/fs/filesystem.go`
- `net/port.go` already has Windows support (`actionUsedPortsNetstat`)
- After adding/modifying completers, run `go generate ./cmd/...` to regenerate completer lists
- Build with `go install -v -tags force_all ./cmd/carapace` to verify all platforms compile
- Run `gofmt -d -s .` and `go run ./cmd/carapace-lint completers/*/*/cmd/*.go` to verify
- Windows man pages: use Microsoft Learn URLs (e.g. `https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/<cmd>`)

---

## Phase 1: Core Public Actions — Add Windows Support

Add `runtime.GOOS` windows branches to shared actions in `pkg/actions/`, following the `fs.ActionBlockDevices` darwin pattern.

| Status | File | Action | Notes |
|--------|------|--------|-------|
| ✅ done | `pkg/actions/fs/blockdevice.go` | `ActionBlockDevices` | Windows: uses `wmic logicaldisk` via `blockdevice-wmic.go` |
| ✅ done | `pkg/actions/fs/mount.go` | `ActionMounts` | Windows: uses `wmic logicaldisk` for drive letters and providers |
| ✅ done | `pkg/actions/fs/filesystem.go` | `ActionFilesystemTypes` | Windows: NTFS, FAT32, exFAT, ReFS, CDFS, UDF |
| ✅ done | `pkg/actions/os/user.go` | `ActionUsers` | Windows: uses `net user` instead of `/etc/passwd` |
| ✅ done | `pkg/actions/os/group.go` | `ActionGroups` | Windows: uses `net localgroup` / `net group` instead of `/etc/group` |
| ✅ done | `pkg/actions/os/kernel.go` | `ActionKernelModulesLoaded` | Windows: uses `driverquery` instead of `lsmod`/`kextstat` |
| ✅ done | `pkg/actions/os/kernel.go` | `ActionKernelModules` | Windows: uses `driverquery` for installed drivers |
| ✅ done | `pkg/actions/os/font.go` | `ActionFontFamilies` | Windows: enumerates `C:\Windows\Fonts` and user font directory |
| ✅ done | `pkg/actions/os/display.go` | `ActionDisplays` | Windows: returns empty (no X display identifiers) |
| ✅ done | `pkg/actions/os/sound.go` | `ActionSoundCards` | Windows: returns empty (no standard sound card enumeration) |
| ☐ todo | `pkg/actions/os/locale.go` | `ActionLocales` | Verify Windows locale format compatibility (Windows uses `en-US` not `en_US`) |
| ✅ done | `pkg/actions/ps/ps.go` | `ActionKillSignals` | Windows: CTRL_C_EVENT, CTRL_BREAK_EVENT, etc. |
| ✅ done | `pkg/actions/ps/ps.go` | `ActionProcessStates` | Windows: Running, Ready, Waiting, Terminated |
| ✅ done | `pkg/actions/tools/mount/source.go` | `ActionSources` | Windows: VolumeGUID key-value source |
| ✅ done | `pkg/actions/tools/mount/option.go` | `ActionMountOptions` | Windows: ro, rw, noexec, nosuid, nodev, noatime, sync, async |

---

## Phase 2: Windows Core Utils (cmd.exe built-ins and System32)

Create Windows-specific completers for commands that ship with Windows. These are the equivalent of Unix coreutils — the everyday commands available in `cmd.exe` and `C:\Windows\System32\`.

### Priority 2a: File operations (Windows equivalents of coreutils)

| Status | Command | Description | Reference URL |
|--------|---------|-------------|---------------|
| ✅ done | `type` | Display file contents (Windows `cat`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/type |
| ✅ done | `copy` | Copy files (Windows `cp`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/copy |
| ✅ done | `move` | Move files (Windows `mv`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/move |
| ✅ done | `del` | Delete files (Windows `rm`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/del |
| ✅ done | `ren` | Rename files (Windows `mv` for rename) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/ren |
| ✅ done | `mkdir` | Create directory (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/mkdir |
| ✅ done | `rmdir` | Remove directory (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/rmdir |
| ✅ done | `dir` | List directory (Windows `ls`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/dir |
| ✅ done | `cd` | Change directory (Windows `cd`/`chdir`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/cd |
| ✅ done | `tree` | Display directory tree (Windows variant, overrides common) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tree |
| ✅ done | `attrib` | Display/set file attributes (Windows `chmod`-adjacent) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/attrib |
| ✅ done | `icacls` | Display/modify ACLs (Windows `chown`/`chmod`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/icacls |
| ✅ done | `where` | Locate commands (Windows `which`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/where |
| ✅ done | `more` | Page output (Windows variant, overrides common) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/more |
| ✅ done | `sort` | Sort lines (Windows variant, differs from Unix sort) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/sort |
| ✅ done | `find` | Search for text (Windows variant, NOT Unix find) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/find |
| ✅ done | `findstr` | Search for text with regex (Windows `grep`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/findstr |
| ✅ done | `fc` | Compare files (Windows `diff`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/fc |
| ✅ done | `comp` | Compare files byte-by-byte | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/comp |
| ✅ done | `clip` | Redirect to clipboard (Windows `pbcopy`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/clip |
| ✅ done | `prompt` | Set command prompt string | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/prompt |
| ✅ done | `title` | Set console window title | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/title |
| ✅ done | `color` | Set console text color | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/color |
| ✅ done | `chcp` | Change active code page | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/chcp |
| ✅ done | `mode` | Configure console/serial/parallel ports | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/mode |
| ✅ done | `subst` | Associate path with drive letter | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/subst |
| ✅ done | `set` | Display/set environment variables | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/set |
| ✅ done | `ver` | Display Windows version | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/ver |
| ✅ done | `date` | Display/set date (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/date |
| ✅ done | `time` | Display/set time (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/time |
| ✅ done | `choice` | Prompt user for choice | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/choice |
| ✅ done | `timeout` | Wait for timeout/keypress | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/timeout |
| ✅ done | `pause` | Suspend processing | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/pause |
| ✅ done | `start` | Start program in new window | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/start |
| ✅ done | `call` | Call batch program | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/call |
| ☐ todo | `exit` | Exit cmd.exe | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/exit |
| ✅ done | `echo` | Display messages (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/echo |
| ☐ todo | `rem` | Record comments in batch files | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/rem |
| ✅ done | `path` | Display/set executable search path | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/path |
| ✅ done | `assoc` | Display/modify file associations | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/assoc |
| ☐ todo | `ftype` | Display/modify file type associations | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/ftype |
| ✅ done | `doskey` | Edit command lines, recall commands, create macros | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/doskey |
| ✅ done | `pushd` | Push current directory and change | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/pushd |
| ✅ done | `popd` | Pop directory from stack | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/popd |
| ✅ done | `setlocal` | Start localization of environment | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/setlocal |
| ✅ done | `endlocal` | End localization of environment | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/endlocal |

### Priority 2b: File copy/archive commands (System32)

| Status | Command | Description | Reference URL |
|--------|---------|-------------|---------------|
| ✅ done | `xcopy` | Extended copy with directory support | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/xcopy |
| ✅ done | `robocopy` | Robust file copy (successor to xcopy) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/robocopy |
| ☐ todo | `replace` | Replace files | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/replace |
| ☐ todo | `expand` | Extract files from .cab archives | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/expand |
| ☐ todo | `makecab` | Create .cab archive | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/makecab |
| ☐ todo | `extract` | Extract from cabinet (older) | — |
| ☐ todo | `compact` | Compress/uncompress files (NTFS) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/compact |
| ☐ todo | `cipher` | Encrypt/decrypt files (EFS) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/cipher |
| ☐ todo | `print` | Print a text file | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/print |

---

## Phase 3: Windows Disk and Filesystem Commands

| Status | Command | Description | Reference URL |
|--------|---------|-------------|---------------|
| ☐ todo | `diskpart` | Disk partitioning/scripting tool | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/diskpart |
| ☐ todo | `format` | Format a volume | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/format |
| ☐ todo | `chkdsk` | Check disk for errors | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/chkdsk |
| ☐ todo | `label` | Create/change/delete volume label | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/label |
| ☐ todo | `vol` | Display volume label and serial number | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/vol |
| ☐ todo | `fsutil` | Filesystem utility (many subcommands) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/fsutil |
| ☐ todo | `mountvol` | Create/delete volume mount points | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/mountvol |
| ☐ todo | `convert` | Convert FAT to NTFS | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/convert |
| ☐ todo | `defrag` | Defragment a volume | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/defrag |
| ☐ todo | `recover` | Recover readable info from bad disk | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/recover |
| ☐ todo | `vssadmin` | Volume Shadow Copy administration | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/vssadmin |
| ☐ todo | `bcdedit` | Boot configuration data editor | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/bcdedit |
| ☐ todo | `bootcfg` | Configure boot entries (legacy) | — |
| ☐ todo | `diskraid` | RAID configuration tool | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/diskraid |
| ☐ todo | `fltMC` | Filter manager control (minifilters) | — |
| ☐ todo | `chkntfs` | Check NTFS at boot | — |
| ☐ todo | `pagefileconfig` | Page file configuration | — |
| ☐ todo | `recover` | Recover readable data | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/recover |

---

## Phase 4: Windows System and Process Management

| Status | Command | Description | Reference URL |
|--------|---------|-------------|---------------|
| ✅ done | `tasklist` | List running processes (Windows `ps`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tasklist |
| ✅ done | `taskkill` | End processes (Windows `kill`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/taskkill |
| ✅ done | `sc` | Service control manager (query/start/stop) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/sc |
| ✅ done | `net` | Network/service commands (net start/stop/user/share) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/net |
| ✅ done | `schtasks` | Scheduled tasks (Windows `cron`/`at`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/schtasks |
| ✅ done | `systeminfo` | Display system information | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/systeminfo |
| ✅ done | `whoami` | Display current user (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/whoami |
| ✅ done | `shutdown` | Shutdown/restart computer | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/shutdown |
| ✅ done | `powercfg` | Power configuration settings | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/powercfg |
| ✅ done | `driverquery` | List installed drivers | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/driverquery |
| ☐ todo | `wmic` | WMI command-line interface | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/wmic |
| ✅ done | `reg` | Registry editor command-line | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/reg |
| ✅ done | `runas` | Run program as different user | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/runas |
| ☐ todo | `openfiles` | Query/display open files | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/openfiles |
| ☐ todo | `qprocess` | Query processes | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/qprocess |
| ☐ todo | `qwinsta` | Query sessions (RDP) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/qwinsta |
| ☐ todo | `rwinsta` | Reset sessions (RDP) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/rwinsta |
| ☐ todo | `quser` | Query users | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/quser |
| ☐ todo | `msg` | Send message to user/session | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/msg |
| ☐ todo | `logoff` | Log off user session | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/logoff |
| ☐ todo | `change` | Change RDP settings (logon, port, user) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/change |
| ☐ todo | `shadow` | Shadow (monitor) another RDP session | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/shadow |
| ☐ todo | `tsdiscon` | Disconnect terminal session | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tsdiscon |
| ☐ todo | `tskill` | End process in session | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tskill |
| ☐ todo | `tsprof` | Copy user config | — |
| ☐ todo | `flattemp` | Enable/disable flat temp dirs | — |
| ☐ todo | `bootcfg` | Configure boot entries | — |

---

## Phase 5: Windows Network Commands

| Status | Command | Description | Reference URL |
|--------|---------|-------------|---------------|
| ✅ done | `ipconfig` | IP configuration display/release/renew | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/ipconfig |
| ☐ todo | `netsh` | Network shell (extensive subcommands) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/netsh |
| ☐ todo | `ping` | Ping (Windows variant, overrides common) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/ping |
| ✅ done | `tracert` | Trace route (Windows `traceroute`) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tracert |
| ✅ done | `pathping` | Path ping (ping + traceroute hybrid) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/pathping |
| ✅ done | `nslookup` | DNS lookup (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/nslookup |
| ✅ done | `arp` | ARP table (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/arp |
| ✅ done | `route` | Routing table (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/route |
| ✅ done | `netstat` | Network statistics (Windows variant) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/netstat |
| ☐ todo | `hostname` | Display hostname (Windows variant, overrides common) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/hostname |
| ☐ todo | `getmac` | Get MAC address | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/getmac |
| ✅ done | `nbtstat` | NetBIOS statistics | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/nbtstat |
| ☐ todo | `netdiag` | Network diagnostics | — |
| ☐ todo | `portqry` | Port query tool | — |
| ☐ todo | `dnscmd` | DNS server administration | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/dnscmd |
| ☐ todo | `dsacls` | DS ACL management | — |
| ☐ todo | `dcdiag` | Domain controller diagnostics | — |
| ☐ todo | `repadmin` | AD replication administration | — |
| ☐ todo | `nltest` | NLB/AD testing | — |
| ☐ todo | `dism` | Deployment Image Servicing and Management | https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/what-is-dism |

---

## Phase 6: Windows Security and Certificate Commands

| Status | Command | Description | Reference URL |
|--------|---------|-------------|---------------|
| ☐ todo | `certutil` | Certificate utility (many subcommands) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/certutil |
| ☐ todo | `cipher` | Encrypt/decrypt (EFS) — also in Phase 2 | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/cipher |
| ☐ todo | `manage-bde` | Manage BitLocker Drive Encryption | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/manage-bde |
| ☐ todo | `repair-bde` | Repair BitLocker-encrypted drive | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/repair-bde |
| ☐ todo | `bdehdcfg` | BitLocker drive preparation | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/bdehdcfg |
| ☐ todo | `ksetup` | Kerberos setup | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/ksetup |
| ☐ todo | `ktsetup` | Kerberos setup (alt) | — |
| ☐ todo | `setspn` | Set service principal names | — |
| ☐ todo | `gpupdate` | Update group policy | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/gpupdate |
| ☐ todo | `gpresult` | Display group policy results | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/gpresult |
| ☐ todo | `auditpol` | Audit policy management | — |
| ☐ todo | `lodctr` | Load performance counters | — |
| ☐ todo | `unlodctr` | Unload performance counters | — |
| ☐ todo | `typeperf` | Performance counter data | — |
| ☐ todo | `logman` | Performance log management | — |
| ☐ todo | `relog` | Relog performance counter data | — |

---

## Phase 7: Modern Windows Tools and Package Managers

| Status | Command | Description | Reference URL |
|--------|---------|-------------|---------------|
| ✅ done | `winget` | Windows Package Manager | Already in `completers/windows/winget_completer` |
| ✅ done | `wt` | Windows Terminal | Already in `completers/common/wt_completer` |
| ☐ todo | `pwsh` | PowerShell 7+ (cross-platform) | https://learn.microsoft.com/en-us/powershell/scripting/developer/cmdlet-aliases |
| ☐ todo | `powershell` | Windows PowerShell 5.1 | https://learn.microsoft.com/en-us/powershell/scripting/developer/cmdlet-aliases |
| ☐ todo | `dotnet` | .NET CLI | Already in common? Verify |
| ☐ todo | `azcopy` | Azure storage copy tool | https://learn.microsoft.com/en-us/azure/storage/common/storage-use-azcopy-v10 |
| ☐ todo | `winsw` | Windows Service Wrapper | https://github.com/winsw/winsw |
| ☐ todo | `nssm` | Non-Sucking Service Manager | https://nssm.cc/ |
| ☐ todo | `choco` | Chocolatey package manager | https://docs.chocolatey.org/en-us/choco/ |
| ☐ todo | `scoop` | Scoop package manager | https://scoop.sh/ |
| ☐ todo | `sfc` | System File Checker | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/sfc |
| ☐ todo | `dism` | DISM (also in Phase 5) | https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/what-is-dism |
| ☐ todo | `bcdboot` | Boot file creation/repair | https://learn.microsoft.com/en-us/windows-hardware/manufacture/desktop/bcdboot-command-line-options |
| ☐ todo | `sysprep` | System preparation tool | — |

---

## Phase 8: Windows Event Log and Diagnostics

| Status | Command | Description | Reference URL |
|--------|---------|-------------|---------------|
| ☐ todo | `wevtutil` | Windows event log utility | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/wevtutil |
| ☐ todo | `eventvwr` | Event Viewer (MMC snap-in launcher) | — |
| ☐ todo | `perfmon` | Performance Monitor (MMC snap-in) | — |
| ☐ todo | `wmic` | WMI (also in Phase 4) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/wmic |
| ☐ todo | `dispdiag` | Display diagnostics | — |
| ☐ todo | `dxdiag` | DirectX diagnostics | — |
| ☐ todo | `mdt` | Microsoft Deployment Toolkit | — |
| ☐ todo | `logman` | Performance log (also in Phase 6) | — |
| ☐ todo | `tracerpt` | Convert trace logs to reports | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tracerpt |
| ☐ todo | `getmac` | MAC address (also in Phase 5) | — |

---

## Phase 9: Miscellaneous Windows Commands

| Status | Command | Description | Reference URL |
|--------|---------|-------------|---------------|
| ☐ todo | `cmd` | Windows command interpreter | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/cmd |
| ☐ todo | `cscript` | Windows script host (CLI) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/cscript |
| ☐ todo | `wscript` | Windows script host (GUI) | — |
| ☐ todo | `mshta` | HTML application host | — |
| ☐ todo | `forfiles` | Select files by date/mask for batch processing | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/forfiles |
| ☐ todo | `if` | Conditional processing (cmd builtin) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/if |
| ☐ todo | `for` | Loop command (cmd builtin) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/for |
| ☐ todo | `goto` | Jump to label (cmd builtin) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/goto |
| ☐ todo | `shift` | Shift batch parameters | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/shift |
| ☐ todo | `mklink` | Create symbolic/hard link (cmd builtin) | https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/mklink |
| ☐ todo | `debug` | Debug tool (legacy, 16-bit era) | — |
| ☐ todo | `edit` | MS-DOS Editor (legacy) | — |
| ☐ todo | `edlin` | Line editor (legacy) | — |
| ☐ todo | `qbasic` | QuickBASIC (legacy) | — |
| ☐ todo | `backup` | Backup (legacy) | — |
| ☐ todo | `restore` | Restore (legacy) | — |
| ☐ todo | `lpq` | Print queue status | — |
| ☐ todo | `lpr` | Print file to network printer | — |
| ☐ todo | `ftype` | File type (also in Phase 2) | — |

---

## Phase 10: Completer Quality Improvements (Follow-up)

| Status | Task | Notes |
|--------|------|-------|
| ☐ todo | Subcommand-level flags for complex commands | `diskpart`, `netsh`, `sc`, `certutil`, `wmic`, `fsutil`, `schtasks`, `reg` have many subcommands |
| ☐ todo | Dynamic completions | `sc query` service names, `reg` registry keys, `schtasks` task names, `net share` share names |
| ☐ todo | Flag accuracy review | Verify all flags match actual Windows 10/11/Server 2019+ behavior |
| ☐ todo | Override common completers where flags differ | `ping`, `hostname`, `tree`, `more`, `sort` have Windows-specific flag sets |
| ☐ todo | PowerShell-aware completions | Some commands interact with PowerShell (pwsh, powershell) |

---

## Workflow per completer

1. Create `completers/windows/<cmd>_completer/` with `main.go` and `cmd/root.go`
2. Use Microsoft Learn URL in `Long` field (e.g. `https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/<cmd>`)
3. Define flags matching the Windows version (alphabetically ordered within contiguous blocks)
4. Add flag completions and positional completions as appropriate
5. Run `go generate ./cmd/...` to regenerate completer lists
6. Build: `go install -v -tags force_all ./cmd/carapace`
7. Test: `go test -v ./cmd/...`
8. Lint: `go run ./cmd/carapace-lint completers/windows/<cmd>_completer/cmd/*.go`
9. Format: `gofmt -d -s .`

## Workflow per shared action (Windows support)

1. Locate the action in `pkg/actions/<category>/<file>.go`
2. Add `runtime.GOOS` switch with `case "windows":` branch
3. Implement Windows-specific logic (e.g. using `wmic`, `net`, `driverquery`, PowerShell commands)
4. Run `go generate ./cmd/...` to regenerate macro map if public action was modified
5. Build and test as above

## Notes

- Windows commands natively use `/` for flags (e.g. `/v`, `/q`), but carapace completers use dash-style (`-v`) for consistency with the rest of the project
- Windows paths use `/` (forward slash) — carapace uses slash for file paths everywhere; modern PowerShell supports this
- Many Windows commands are cmd.exe built-ins (not standalone executables) — completers still work as they describe the command interface
- `wmic` is deprecated in Windows 11 — prefer PowerShell `Get-CimInstance` / `Get-WmiObject` for dynamic actions
- Commands marked as appearing in multiple phases should be implemented once and cross-referenced
