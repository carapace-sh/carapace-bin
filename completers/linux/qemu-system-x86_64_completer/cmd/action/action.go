package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// ActionMachines completes available machine types by querying the qemu-system binary.
func ActionMachines(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		binaryName := cmd.Name()
		return carapace.ActionExecCommand(binaryName, "-machine", "help")(func(output []byte) carapace.Action {
			lines := string(output)
			values := make([]string, 0)
			for _, line := range splitLines(lines) {
				if name, desc, ok := parseMachineHelpLine(line); ok {
					values = append(values, name, desc)
				}
			}
			if len(values) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(values...).Tag("machine types")
		})
	})
}

// ActionCpuModels completes available CPU models by querying the qemu-system binary.
func ActionCpuModels(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		binaryName := cmd.Name()
		return carapace.ActionExecCommand(binaryName, "-cpu", "help")(func(output []byte) carapace.Action {
			lines := string(output)
			values := make([]string, 0)
			for _, line := range splitLines(lines) {
				if name, desc, ok := parseCpuHelpLine(line); ok {
					values = append(values, name, desc)
				}
			}
			if len(values) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionValuesDescribed(values...).Tag("cpu models")
		})
	})
}

func splitLines(s string) []string {
	var lines []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			lines = append(lines, s[start:i])
			start = i + 1
		}
	}
	if start < len(s) {
		lines = append(lines, s[start:])
	}
	return lines
}

// parseMachineHelpLine parses a line like "pc-i440fx-9.2        Standard PC (i440FX + PIIX, 1996)"
func parseMachineHelpLine(line string) (name, desc string, ok bool) {
	trimmed := trimSpaces(line)
	if trimmed == "" || trimmed == "Supported machines are:" {
		return "", "", false
	}
	// Find the first run of multiple spaces separating name from description
	for i := 0; i < len(trimmed)-1; i++ {
		if i > 0 && trimmed[i] == ' ' && trimmed[i+1] == ' ' {
			name = trimSpaces(trimmed[:i])
			desc = trimSpaces(trimmed[i:])
			if name != "" {
				return name, desc, true
			}
		}
	}
	// Single word (no description)
	if trimmed != "" {
		return trimmed, "", true
	}
	return "", "", false
}

// parseCpuHelpLine parses a line like "  486                   (alias configured by machine type)"
func parseCpuHelpLine(line string) (name, desc string, ok bool) {
	trimmed := trimSpaces(line)
	if trimmed == "" || trimmed == "Available CPUs:" {
		return "", "", false
	}
	for i := 0; i < len(trimmed)-1; i++ {
		if i > 0 && trimmed[i] == ' ' && trimmed[i+1] == ' ' {
			name = trimSpaces(trimmed[:i])
			desc = trimSpaces(trimmed[i:])
			if name != "" {
				return name, desc, true
			}
		}
	}
	if trimmed != "" {
		return trimmed, "", true
	}
	return "", "", false
}

func trimSpaces(s string) string {
	start := 0
	end := len(s)
	for start < end && s[start] == ' ' {
		start++
	}
	for end > start && s[end-1] == ' ' {
		end--
	}
	return s[start:end]
}

// ActionImageFormats completes supported QEMU image formats.
func ActionImageFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"blkdebug", "block debug",
		"blklogwrites", "block log writes",
		"blkverify", "block verify",
		"bochs", "Bochs disk image",
		"cloop", "compressed loop",
		"compress", "compressed image",
		"copy-before-write", "copy before write filter",
		"copy-on-read", "copy on read filter",
		"dmg", "Apple Disk Management image",
		"file", "raw file",
		"ftp", "FTP",
		"ftps", "FTP over TLS",
		"gluster", "GlusterFS",
		"host_cdrom", "host CD-ROM",
		"host_device", "host device",
		"http", "HTTP",
		"https", "HTTPS",
		"iscsi", "iSCSI",
		"iser", "iSER",
		"luks", "LUKS encrypted",
		"nbd", "Network Block Device",
		"nfs", "NFS",
		"null-aio", "null AIO",
		"null-co", "null",
		"nvme", "NVMe",
		"parallels", "Parallels disk image",
		"preallocate", "preallocate filter",
		"qcow", "old QEMU Copy-On-Write",
		"qcow2", "QEMU Copy-On-Write v2",
		"qed", "QEMU Enhanced Disk",
		"quorum", "quorum filter",
		"raw", "raw disk image",
		"replication", "replication filter",
		"snapshot-access", "snapshot access filter",
		"ssh", "SSH",
		"throttle", "throttle filter",
		"vdi", "VirtualBox Disk Image",
		"vhdx", "VHDX disk image",
		"vmdk", "VMware disk image",
		"vpc", "VPC (VHD) disk image",
		"vvfat", "Virtual VFAT",
	).Tag("image formats")
}

// ActionCacheModes completes QEMU cache modes.
func ActionCacheModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"none", "no caching",
		"writeback", "write-back caching (default)",
		"writethrough", "write-through caching",
		"directsync", "direct sync",
		"unsafe", "unsafe caching",
	).Tag("cache modes")
}

// ActionAioModes completes QEMU AIO modes.
func ActionAioModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"threads", "pthreads based AIO",
		"native", "Linux native AIO",
		"io_uring", "Linux io_uring AIO",
	).Tag("aio modes")
}

// ActionAccels completes QEMU accelerators.
func ActionAccels() carapace.Action {
	return carapace.ActionValuesDescribed(
		"kvm", "Linux KVM",
		"tcg", "Tiny Code Generator (software emulation)",
		"xen", "Xen",
		"hvf", "macOS Hypervisor.framework",
		"whpx", "Windows Hypervisor Platform",
		"nvmm", "NetBSD virtual machine monitor",
		"nitro", "AWS Nitro Enclaves",
		"mshv", "Microsoft Hyper-V",
	).Tag("accelerators")
}

// ActionDisplayTypes completes QEMU display types.
func ActionDisplayTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"gtk", "GTK window",
		"sdl", "SDL window",
		"egl-headless", "EGL headless rendering",
		"gl=on", "OpenGL rendering",
		"gl=core", "OpenGL core profile",
		"gl=es", "OpenGL ES",
		"curses", "curses/ncurses interface",
		"none", "no display output",
		"dbus", "D-Bus display",
		"cocoa", "macOS Cocoa window",
	).Tag("display types")
}

// ActionVgaTypes completes QEMU VGA types.
func ActionVgaTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"cirrus", "Cirrus Logic GD5446",
		"std", "standard VGA",
		"vmware", "VMware SVGA-II",
		"virtio", "VirtIO GPU",
		"qxl", "QXL (SPICE)",
		"tcx", "TCX (Sun)",
		"cg3", "CG3 (Sun)",
		"none", "no VGA",
	).Tag("vga types")
}

// ActionCharDevices completes QEMU character device backend types for
// -serial, -parallel, -monitor, -qmp, -qmp-pretty, and -gdb.
// These accept values like "chardev:id", "file:path", "tcp:host:port",
// "unix:path", "pty", "vc", "stdio", "null", "none", etc.
// Returns the device prefix with NoSpace(":") so the user can continue
// typing the colon-separated argument.
func ActionCharDevices() carapace.Action {
	return carapace.ActionValuesDescribed(
		"chardev", "use a named character device defined with -chardev",
		"file", "log to a file",
		"mon", "multiplexed character backend (monitor + device)",
		"null", "discard all output",
		"none", "no device allocated",
		"parport", "use host parallel port (Linux only)",
		"pipe", "named pipe (Windows only)",
		"pty", "pseudo TTY (Linux only)",
		"serial", "use host serial port (Linux only)",
		"spicevmc", "SPICE virtual machine channel",
		"stdio", "standard input/output (Unix only)",
		"tcp", "TCP socket",
		"udp", "UDP socket",
		"unix", "Unix domain socket",
		"vc", "virtual console",
	).Suffix(":").NoSpace(':').Tag("char devices")
}

// ActionWatchdogActions completes QEMU watchdog actions.
func ActionWatchdogActions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"reset", "reset the guest",
		"shutdown", "shutdown the guest",
		"poweroff", "power off the guest",
		"pause", "pause the guest",
		"debug", "print debug message",
		"none", "do nothing",
		"inject-nmi", "inject an NMI into the guest",
	).Tag("watchdog actions")
}
