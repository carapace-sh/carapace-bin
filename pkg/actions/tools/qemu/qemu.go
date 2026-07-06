package qemu

import "github.com/carapace-sh/carapace"

// ActionAccels completes QEMU accelerators
//
//	kvm (Linux KVM)
//	tcg (Tiny Code Generator (software emulation))
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
	).Tag("accelerators").Uid("qemu", "accels")
}

// ActionAioModes completes QEMU AIO modes
//
//	threads (pthreads based AIO)
//	native (Linux native AIO)
func ActionAioModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"threads", "pthreads based AIO",
		"native", "Linux native AIO",
		"io_uring", "Linux io_uring AIO",
	).Tag("aio modes").Uid("qemu", "aio-modes")
}

// ActionCacheModes completes QEMU cache modes
//
//	none (no caching)
//	writeback (write-back caching (default))
func ActionCacheModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"none", "no caching",
		"writeback", "write-back caching (default)",
		"writethrough", "write-through caching",
		"directsync", "direct sync",
		"unsafe", "unsafe caching",
	).Tag("cache modes").Uid("qemu", "cache-modes")
}

// ActionCharDevices completes QEMU character device backend types
//
//	chardev (use a named character device defined with -chardev)
//	file (log to a file)
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
	).Tag("char devices").Uid("qemu", "char-devices")
}

// ActionDisplayTypes completes QEMU display types
//
//	gtk (GTK window)
//	sdl (SDL window)
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
	).Tag("display types").Uid("qemu", "display-types")
}

// ActionImageFormats completes supported QEMU image formats
//
//	blkdebug (block debug)
//	blklogwrites (block log writes)
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
	).Tag("image formats").Uid("qemu", "image-formats")
}

// ActionVgaTypes completes QEMU VGA types
//
//	cirrus (Cirrus Logic GD5446)
//	std (standard VGA)
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
	).Tag("vga types").Uid("qemu", "vga-types")
}

// ActionWatchdogActions completes QEMU watchdog actions
//
//	reset (reset the guest)
//	shutdown (shutdown the guest)
func ActionWatchdogActions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"reset", "reset the guest",
		"shutdown", "shutdown the guest",
		"poweroff", "power off the guest",
		"pause", "pause the guest",
		"debug", "print debug message",
		"none", "do nothing",
		"inject-nmi", "inject an NMI into the guest",
	).Tag("watchdog actions").Uid("qemu", "watchdog-actions")
}
