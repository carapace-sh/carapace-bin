# Plan: Add Darwin (macOS) Completers

Track progress for adding darwin completers to carapace-bin.

## Context

- Existing darwin completers: `mount`, `umount`, `skhd`, `orb`, `fileicon` (5 total)
- Existing bsd completer: `sed` (1 total, shared via bsd group)
- Resolution priority on darwin: `darwin` > `bsd` > `unix` > `common`
- Darwin completers go in `completers/darwin/<cmd>_completer/` (no build tags needed)
- Darwin support in shared actions uses `runtime.GOOS` switches (not build tags), e.g. `pkg/actions/fs/blockdevice.go`, `pkg/actions/fs/mount.go`, `pkg/actions/fs/filesystem.go`
- After adding/modifying completers, run `go generate ./cmd/...` to regenerate completer lists
- Build with `go install -v -tags force_all ./cmd/carapace` to verify all platforms compile
- Run `gofmt -d -s .` and `go run ./cmd/carapace-lint completers/*/*/cmd/*.go` to verify

---

## Phase 1: Core Public Actions — Add Darwin Support

Add `runtime.GOOS` darwin branches to shared actions in `pkg/actions/`, following the `fs.ActionBlockDevices` pattern.

| Status | File | Action | Notes |
|--------|------|--------|-------|
| ✅ done | `pkg/actions/fs/blockdevice.go` | `ActionBlockDevices` | Dispatches to `diskutil` on darwin |
| ✅ done | `pkg/actions/fs/mount.go` | `ActionMounts` | Parses `mount` output on darwin |
| ✅ done | `pkg/actions/fs/filesystem.go` | `ActionFilesystemTypes` | Darwin filesystem list (apfs, hfs, etc.) |
| ✅ done | `pkg/actions/tools/mount/source.go` | `ActionSources` | Darwin LABEL=/UUID= sources |
| ✅ done | `pkg/actions/tools/mount/option.go` | `ActionMountOptions` | Darwin mount options |
| ✅ done | `pkg/actions/os/user.go` | `ActionUsers` | Darwin: handles `/usr/bin/false` as nologin shell via `isNonLoginShell` |
| ☐ todo | `pkg/actions/os/group.go` | `ActionGroups` | Uses `/etc/group` which exists on darwin |
| ☐ todo | `pkg/actions/os/kernel.go` | `ActionKernelModules` | Darwin has no kernel modules in the Linux sense; may need `kextstat` or return empty |
| ☐ todo | `pkg/actions/os/locale.go` | `ActionLocales` | Static list — works on darwin as-is |
| ☐ todo | `pkg/actions/os/font.go` | `ActionFonts` | Uses `fc-list` — works on darwin if fontconfig installed |
| ☐ todo | `pkg/actions/os/display.go` | `ActionDisplays` | Uses `w` command — works on darwin |
| ✅ done | `pkg/actions/ps/ps.go` | `ActionKillSignals` | Darwin: EMT, INFO signals; no PWR/STKFLT |
| ✅ done | `pkg/actions/ps/ps.go` | `ActionProcessStates` | Darwin: D,I,R,S,T,U,Z states |

---

## Phase 2: BSD Coreutils Completers (Differ from GNU/Unix)

Create darwin-specific completers for commands where BSD/macOS flags differ significantly from GNU. These override the existing `unix/` completers via resolution priority.

### Priority 2a: High-impact, significant flag divergence

| Status | Command | Key differences from GNU | Man page |
|--------|---------|--------------------------|----------|
| ✅ done | `ls` | BSD: `-G` color, `-e` ACL, `-O` flags, `-@` xattr, `-W` whiteouts; no `--author`, `--block-size`, `--dired`, `--time-style`, `--hyperlink`, `--quoting-style` | https://keith.github.io/xcode-manpages/ls.1.html |
| ✅ done | `stat` | BSD: `-f` format string (completely different specifiers: `%N`, `%z`, `%u`, `%g`), `-r` raw, `-L` follow; no `--format`, `--printf`, `--file-system` | https://keith.github.io/xcode-manpages/stat.1.html |
| ✅ done | `date` | BSD: `-j` don't set, `-f` input format, `-v` adjust value, `-r` seconds; no `--date`, `--iso-8601`, `--rfc-3339`, `--debug` | https://keith.github.io/xcode-manpages/date.1.html |
| ✅ done | `df` | BSD: `-g`, `-m`, `-h`, `-H`, `-i`; no `--output`, `--total`, `--block-size`, `--exclude-type` | https://keith.github.io/xcode-manpages/df.1.html |
| ✅ done | `du` | BSD: `-d` depth, `-t` threshold, `-h`, `-H`, `-c`; no `--apparent-size`, `--files0-from`, `--inodes`, `--time`, `--exclude` | https://keith.github.io/xcode-manpages/du.1.html |
| ✅ done | `cp` | BSD: `-c` clipboard, `-n` no-clobber, `-p` preserve, `-R`; no `--archive`, `--reflink`, `--sparse`, `--update`, `--parents` | https://keith.github.io/xcode-manpages/cp.1.html |
| ✅ done | `mv` | BSD: `-n` no-clobber; no `--backup`, `--exchange`, `--target-directory`, `--update`, `--debug` | https://keith.github.io/xcode-manpages/mv.1.html |
| ✅ done | `rm` | BSD: `-P` overwrite; no `--no-preserve-root`, `--one-file-system`, `--interactive=WHEN` | https://keith.github.io/xcode-manpages/rm.1.html |
| ✅ done | `dd` | BSD: no `iflag`/`oflag`; different `conv` options (`swab`, `block`, `unblock`); `iseq`, `oseq` | https://keith.github.io/xcode-manpages/dd.1.html |

### Priority 2b: Moderate flag divergence

| Status | Command | Key differences from GNU | Man page |
|--------|---------|--------------------------|----------|
| ✅ done | `touch` | BSD: `-t` time, `-r` reference, `-f` force; no `--date`, `--time=WORD` | https://keith.github.io/xcode-manpages/touch.1.html |
| ✅ done | `chmod` | BSD: `-H`/`-L`/`-P` symlink traversal, `-N` ACL, `-v` verbose; no `--reference` | https://keith.github.io/xcode-manpages/chmod.1.html |
| ✅ done | `chown` | BSD: `-H`/`-L`/`-P`, `-v` verbose, `-n` no-recurse; no `--from`, `--reference` | https://keith.github.io/xcode-manpages/chown.8.html |
| ✅ done | `chgrp` | BSD: `-H`/`-L`/`-P`, `-v` verbose, `-h`; no `--reference` | https://keith.github.io/xcode-manpages/chgrp.1.html |
| ✅ done | `ln` | BSD: `-h` no-follow, `-F` force hard-to-dirs, `-v` verbose; no `-T`, `--backup`, `--relative` | https://keith.github.io/xcode-manpages/ln.1.html |
| ✅ done | `mkdir` | BSD: `-v` verbose, `-m` mode, `-p` parents; no `-Z` (SELinux) | https://keith.github.io/xcode-manpages/mkdir.1.html |
| ✅ done | `rmdir` | BSD: `-v` verbose, `-p` parents; no `--ignore-fail-on-non-empty` | https://keith.github.io/xcode-manpages/rmdir.1.html |
| ✅ done | `head` | BSD: `-N` shorthand, `-c` bytes, `-n` lines; no `--zero-terminated`, `--verbose` | https://keith.github.io/xcode-manpages/head.1.html |
| ✅ done | `tail` | BSD: `-r` reverse, `-q`, `-v`; no `--follow=name`, `--retry`, `--max-unchanged-stats`, `--pid` | https://keith.github.io/xcode-manpages/tail.1.html |
| ✅ done | `cat` | BSD: `-e`, `-t`, `-v`, `-n`, `-s`, `-l`; no `-A`, `-E`, `-T`, `--show-all` | https://keith.github.io/xcode-manpages/cat.1.html |
| ✅ done | `wc` | BSD: `-l`, `-w`, `-c`, `-m`, `-L`; no `--files0-from`, `--total`, `--debug` | https://keith.github.io/xcode-manpages/wc.1.html |

### Priority 2c: Minor divergence (assess whether needed)

| Status | Command | Key differences from GNU | Man page |
|--------|---------|--------------------------|----------|
| ✅ done | `cut` | BSD: mostly compatible; no `--output-delimiter`, `--complement` | https://keith.github.io/xcode-manpages/cut.1.html |
| ✅ done | `sort` | BSD: limited GNU extensions; no `--files0-from`, `--debug`, `--parallel` | https://keith.github.io/xcode-manpages/sort.1.html |
| ✅ done | `uniq` | BSD: mostly compatible; no `--all-repeated`, `--zero-terminated` | https://keith.github.io/xcode-manpages/uniq.1.html |
| ✅ done | `tr` | BSD: mostly compatible; no `--truncate-set1` | https://keith.github.io/xcode-manpages/tr.1.html |
| ✅ done | `tee` | BSD: `-a` append, `-i` ignore; no `--output-error` | https://keith.github.io/xcode-manpages/tee.1.html |
| ✅ done | `basename` | BSD: `-a` multiple, `-s` suffix; mostly compatible | https://keith.github.io/xcode-manpages/basename.1.html |
| ✅ done | `dirname` | BSD: `-a` multiple; mostly compatible | https://keith.github.io/xcode-manpages/dirname.1.html |
| ✅ done | `env` | BSD: `-i`, `-S`, `-P`; no `--chdir`, `--debug`, `--ignore-environment` long opts | https://keith.github.io/xcode-manpages/env.1.html |
| ✅ done | `id` | BSD: `-P`, `-A` audit; no `--zero`, `--user`, `--group` long opts | https://keith.github.io/xcode-manpages/id.1.html |
| ✅ done | `who` | BSD: `-m`, `-q`, `-u`, `-H`; no `--login`, `--runlevel` | https://keith.github.io/xcode-manpages/who.1.html |
| ✅ done | `ps` | BSD: completely different flag syntax (`-aux` vs `aux`), `-e`, `-o` format; no `--sort`, `--info` | https://keith.github.io/xcode-manpages/ps.1.html |
| ✅ done | `kill` | BSD: `-s` signal, `-l` list; no `--table`, `--timeout`, `--queue` | https://keith.github.io/xcode-manpages/kill.1.html |
| ✅ done | `find` | BSD: `-E` extregex, `-X`, `-x`, `-d` depth; different `-print0`, `-printf` | https://keith.github.io/xcode-manpages/find.1.html |
| ✅ done | `xargs` | BSD: `-E` eofstr, `-I` replstr, `-R`; no `--null`, `--verbose`, `--delimiter` long opts | https://keith.github.io/xcode-manpages/xargs.1.html |
| ✅ done | `tar` | BSD: `bsdtar`, different option style; `-z`, `-j`, `-J`; no `--acls`, `--selinux` | https://keith.github.io/xcode-manpages/tar.1.html |
| ✅ done | `base64` | BSD: `-D` decode, `-i`, `-b`, `-o`; different from GNU | https://keith.github.io/xcode-manpages/base64.1.html |
| ❌ skip | `base32` | BSD: `-D` decode, `-i`, `-b`, `-o`; different from GNU | https://keith.github.io/xcode-manpages/base32.1.html |
| ✅ done | `cksum` | BSD: `-o` old format; no `--algorithm`, `--untagged` | https://keith.github.io/xcode-manpages/cksum.1.html |
| ✅ done | `od` | BSD: `-A`, `-j`, `-N`, `-t`, `-v`; no `--read-bytes`, `--format` long opt | https://keith.github.io/xcode-manpages/od.1.html |
| ✅ done | `file` | BSD: `-b`, `-k`, `-z`, `-m`; no `--mime-type`, `--mime-encoding` | https://keith.github.io/xcode-manpages/file.1.html |
| ✅ done | `diff` | BSD: `-u`, `-c`, `-y`; no `--color`, `--no-dereference` | https://keith.github.io/xcode-manpages/diff.1.html |
| ✅ done | `patch` | BSD: `-p`, `-R`, `-i`, `-o`, `-N`; no `--posix`, `--no-backup-if-mismatch` | https://keith.github.io/xcode-manpages/patch.1.html |
| ✅ done | `hostname` | BSD: `-s`; no `--fqdn`, `--all-fqdns`, `--ip-address` | https://keith.github.io/xcode-manpages/hostname.1.html |
| ✅ done | `uname` | BSD: `-a`, `-m`, `-n`, `-p`, `-r`, `-s`, `-v`; no long opts | https://keith.github.io/xcode-manpages/uname.1.html |

### Priority 2d: macOS-specific checksum commands

| Status | Command | Description | Man page |
|--------|---------|-------------|----------|
| ✅ done | `md5` | BSD md5 checksum (not `md5sum`) | https://keith.github.io/xcode-manpages/md5.1.html |
| ✅ done | `shasum` | Perl SHA checksum script (`-a` algorithm, `-c` check) | https://keith.github.io/xcode-manpages/shasum.1.html |

---

## Phase 3: macOS System Commands (Native to macOS)

Create new darwin completers for commands that only exist on macOS.

### Priority 3a: Most commonly used system commands

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ✅ done | `launchctl` | Control launchd services and daemons | https://keith.github.io/xcode-manpages/launchctl.1.html |
| ✅ done | `defaults` | Read/write user preferences (plist) | https://keith.github.io/xcode-manpages/defaults.1.html |
| ✅ done | `diskutil` | Disk partition/format/verify utility | https://keith.github.io/xcode-manpages/diskutil.8.html |
| ✅ done | `plutil` | Property list utility (convert, lint, extract) | https://keith.github.io/xcode-manpages/plutil.1.html |
| ✅ done | `hdiutil` | Create/attach/detach disk images | https://keith.github.io/xcode-manpages/hdiutil.1.html |
| ✅ done | `sw_vers` | Show macOS version info | https://keith.github.io/xcode-manpages/sw_vers.1.html |
| ✅ done | `softwareupdate` | Install macOS software updates | https://keith.github.io/xcode-manpages/softwareupdate.8.html |
| ✅ done | `system_profiler` | Report system hardware/software info | https://keith.github.io/xcode-manpages/system_profiler.8.html |
| ✅ done | `networksetup` | Configure network settings | https://keith.github.io/xcode-manpages/networksetup.8.html |
| ✅ done | `scutil` | System configuration utility | https://keith.github.io/xcode-manpages/scutil.1.html |

### Priority 3b: Directory services and user management

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ✅ done | `dscl` | Directory Service command line | https://keith.github.io/xcode-manpages/dscl.1.html |
| ✅ done | `dscacheutil` | Directory Service cache utility | https://keith.github.io/xcode-manpages/dscacheutil.1.html |
| ☐ todo | `sysadminctl` | System admin operations (user management) | https://developer.apple.com/library/archive/documentation/Darwin/Reference/ManPages/man8/sysadminctl.8.html |
| ☐ todo | `systemsetup` | System configuration (hostname, timezone) | https://keith.github.io/xcode-manpages/systemsetup.8.html |

### Priority 3c: Power, boot, and security

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ✅ done | `pmset` | Power management settings | https://keith.github.io/xcode-manpages/pmset.1.html |
| ☐ todo | `powermetrics` | Power/performance metrics | https://developer.apple.com/library/archive/documentation/Darwin/Reference/ManPages/man1/powermetrics.1.html |
| ✅ done | `nvram` | Read/write NVRAM firmware variables | https://keith.github.io/xcode-manpages/nvram.8.html |
| ✅ done | `bless` | Set bootable disk/volume | https://keith.github.io/xcode-manpages/bless.8.html |
| ☐ todo | `fdesetup` | FileVault setup | https://keith.github.io/xcode-manpages/fdesetup.8.html |
| ☐ todo | `profiles` | Install/remove configuration profiles | https://keith.github.io/xcode-manpages/profiles.8.html |
| ☐ todo | `firmwarepasswd` | Set firmware password | https://developer.apple.com/library/archive/documentation/Darwin/Reference/ManPages/man8/firmwarepasswd.8.html |

### Priority 3d: Spotlight and search

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ✅ done | `mdfind` | Spotlight search | https://keith.github.io/xcode-manpages/mdfind.1.html |
| ✅ done | `mdls` | List Spotlight metadata attributes | https://keith.github.io/xcode-manpages/mdls.1.html |
| ✅ done | `mdimport` | Import files into Spotlight index | https://keith.github.io/xcode-manpages/mdimport.1.html |
| ✅ done | `mdutil` | Manage Spotlight indexing | https://keith.github.io/xcode-manpages/mdutil.1.html |

### Priority 3e: Filesystem and disk

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ☐ todo | `fsck` | File system consistency check (macOS) | https://keith.github.io/xcode-manpages/fsck.8.html |
| ☐ todo | `fsck_apfs` | APFS filesystem consistency check | https://keith.github.io/xcode-manpages/fsck_apfs.8.html |
| ☐ todo | `fsck_hfs` | HFS+ filesystem consistency check | https://keith.github.io/xcode-manpages/fsck_hfs.8.html |
| ☐ todo | `fstyp` | Report filesystem type of a device | https://keith.github.io/xcode-manpages/fstyp.8.html |
| ☐ todo | `newfs_apfs` | Create APFS filesystem | https://keith.github.io/xcode-manpages/newfs_apfs.8.html |
| ☐ todo | `newfs_hfs` | Create HFS+ filesystem | https://keith.github.io/xcode-manpages/newfs_hfs.8.html |
| ☐ todo | `asr` | Apple Software Restore | https://keith.github.io/xcode-manpages/asr.8.html |
| ☐ todo | `apfsctl` | APFS filesystem control | — |

### Priority 3f: Network

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ☐ todo | `airport` | Airport wireless utility | https://keith.github.io/xcode-manpages/airport.1.html |
| ☐ todo | `wdutil` | Wi-Fi diagnostics utility | — |
| ☐ todo | `ipconfig` | IP configuration (macOS variant) | https://keith.github.io/xcode-manpages/ipconfig.8.html |
| ☐ todo | `ifconfig` | Interface config (BSD variant) | https://keith.github.io/xcode-manpages/ifconfig.8.html |
| ☐ todo | `netstat` | Network stats (BSD variant) | https://keith.github.io/xcode-manpages/netstat.1.html |
| ☐ todo | `route` | Routing table (BSD variant) | https://keith.github.io/xcode-manpages/route.8.html |
| ☐ todo | `arp` | Address resolution protocol (BSD) | https://keith.github.io/xcode-manpages/arp.8.html |
| ☐ todo | `traceroute` | Trace route (BSD variant) | https://keith.github.io/xcode-manpages/traceroute.8.html |
| ☐ todo | `ping` | Ping (BSD variant) | https://keith.github.io/xcode-manpages/ping.8.html |

### Priority 3g: Logging and diagnostics

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ☐ todo | `log` | Unified logging system (macOS) | https://keith.github.io/xcode-manpages/log.1.html |
| ☐ todo | `dtrace` | Dynamic tracing framework | https://keith.github.io/xcode-manpages/dtrace.1.html |
| ☐ todo | `dtruss` | Trace system calls (DTrace wrapper) | — |
| ☐ todo | `opensnoop` | Trace file open operations (DTrace) | — |
| ☐ todo | `powermetrics` | Power/performance metrics | — |

### Priority 3h: Kernel extensions

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ☐ todo | `kextstat` | List loaded kernel extensions | https://keith.github.io/xcode-manpages/kextstat.8.html |
| ☐ todo | `kextload` | Load kernel extensions | https://keith.github.io/xcode-manpages/kextload.8.html |
| ☐ todo | `kextunload` | Unload kernel extensions | https://keith.github.io/xcode-manpages/kextunload.8.html |
| ☐ todo | `kmutil` | Kernel manager utility | https://keith.github.io/xcode-manpages/kmutil.8.html |
| ☐ todo | `systemextensionsctl` | Manage system extensions | — |

---

## Phase 4: macOS Development / Xcode Commands

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ✅ done | `xcrun` | Run tool from Xcode active developer directory | https://keith.github.io/xcode-manpages/xcrun.1.html |
| ✅ done | `xcodebuild` | Build Xcode projects and workspaces | https://keith.github.io/xcode-manpages/xcodebuild.1.html |
| ✅ done | `xcode-select` | Switch active Xcode developer directory | https://keith.github.io/xcode-manpages/xcode-select.1.html |
| ✅ done | `simctl` | iOS Simulator control | https://keith.github.io/xcode-manpages/simctl.1.html |
| ✅ done | `xcode-select` | Switch Xcode path | https://keith.github.io/xcode-manpages/xcode-select.1.html |
| ☐ todo | `swift` | Swift REPL / run Swift scripts | https://keith.github.io/xcode-manpages/swift.1.html |
| ☐ todo | `swiftc` | Swift compiler | https://keith.github.io/xcode-manpages/swiftc.1.html |
| ☐ todo | `lldb` | LLVM debugger | https://keith.github.io/xcode-manpages/lldb.1.html |
| ✅ done | `codesign` | Code signing tool | https://keith.github.io/xcode-manpages/codesign.1.html |
| ✅ done | `security` | Keychain and certificate management | https://keith.github.io/xcode-manpages/security.1.html |
| ☐ todo | `otool` | Object file display tool (Mach-O) | https://keith.github.io/xcode-manpages/otool.1.html |
| ☐ todo | `install_name_tool` | Change dynamic shared library install names | https://keith.github.io/xcode-manpages/install_name_tool.1.html |
| ☐ todo | `dyld_info` | Dyld linker info | https://keith.github.io/xcode-manpages/dyld_info.1.html |
| ☐ todo | `agvtool` | Apple Generic Versioning tool | https://keith.github.io/xcode-manpages/agvtool.1.html |
| ☐ todo | `ibtool` | Interface Builder compile tool | https://keith.github.io/xcode-manpages/ibtool.1.html |
| ☐ todo | `actool` | Asset catalog compiler | https://keith.github.io/xcode-manpages/actool.1.html |

---

## Phase 5: Other macOS-Specific Utilities

| Status | Command | Description | Man page / URL |
|--------|---------|-------------|----------------|
| ✅ done | `open` | Open files/apps/URLs with default handler | https://keith.github.io/xcode-manpages/open.1.html |
| ✅ done | `pbcopy` | Copy to clipboard (pasteboard) | https://keith.github.io/xcode-manpages/pbcopy.1.html |
| ✅ done | `pbpaste` | Paste from clipboard | https://keith.github.io/xcode-manpages/pbpaste.1.html |
| ✅ done | `caffeinate` | Prevent system sleep | https://keith.github.io/xcode-manpages/caffeinate.1.html |
| ✅ done | `say` | Text-to-speech | https://keith.github.io/xcode-manpages/say.1.html |
| ✅ done | `osascript` | Execute AppleScript / OSA scripts | https://keith.github.io/xcode-manpages/osascript.1.html |
| ✅ done | `screencapture` | Screen capture utility | https://keith.github.io/xcode-manpages/screencapture.1.html |
| ✅ done | `shortcuts` | Run/manage macOS Shortcuts | https://keith.github.io/xcode-manpages/shortcuts.1.html |
| ✅ done | `sips` | Scriptable image processing system | https://keith.github.io/xcode-manpages/sips.1.html |
| ✅ done | `textutil` | Convert text file formats (rtf, txt, html, doc) | https://keith.github.io/xcode-manpages/textutil.1.html |
| ☐ todo | `sharing` | Manage file/web/screen sharing | https://keith.github.io/xcode-manpages/sharing.1.html |
| ☐ todo | `blueutil` | Bluetooth control (third-party) | https://github.com/toy/blueutil |
| ☐ todo | `automount` | Automount control (NFS) | https://keith.github.io/xcode-manpages/automount.8.html |
| ☐ todo | `createinstallmedia` | Create macOS install media | — |
| ☐ todo | `startosinstall` | Start macOS OS install | — |
| ☐ todo | `repair_packages` | Repair/reverify packages | — |

---

## Phase 6: Third-Party Package Managers (macOS)

| Status | Command | Description | Notes |
|--------|---------|-------------|-------|
| ✅ done | `brew` | Homebrew package manager | Already in `common/brew_completer` |
| ☐ todo | `port` | MacPorts package manager | https://guide.macports.org/ |
| ☐ todo | `mas` | Mac App Store CLI | https://github.com/mas-cli/mas |

---

## Workflow per completer

1. Create `completers/darwin/<cmd>_completer/` with `main.go` and `cmd/root.go`
2. Use BSD/macOS man page URL in `Long` field (e.g. `https://keith.github.io/xcode-manpages/<cmd>.1.html`)
3. Define flags matching the BSD/macOS version (alphabetically ordered within contiguous blocks)
4. Add flag completions and positional completions as appropriate
5. Run `go generate ./cmd/...` to regenerate completer lists
6. Build: `go install -v -tags force_all ./cmd/carapace`
7. Test: `go test -v ./cmd/...`
8. Lint: `go run ./cmd/carapace-lint completers/darwin/<cmd>_completer/cmd/*.go`
9. Format: `gofmt -d -s .`

## Workflow per shared action (darwin support)

1. Locate the action in `pkg/actions/<category>/<file>.go`
2. Add `runtime.GOOS` switch with `case "darwin":` branch
3. Implement darwin-specific logic (e.g. using `diskutil`, `dscl`, `defaults` commands)
4. Run `go generate ./cmd/...` to regenerate macro map if public action was modified
5. Build and test as above
