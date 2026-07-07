package qemu

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
)

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

// ActionCpuFeatures completes QEMU CPUID feature flags
//
//	3dnow (3DNow!)
//	avx (Advanced Vector Extensions)
func ActionCpuFeatures() carapace.Action {
	return carapace.ActionValues(
		"3dnow",
		"3dnowext",
		"3dnowprefetch",
		"abm",
		"ace2",
		"ace2-en",
		"acpi",
		"adx",
		"aes",
		"amd-no-ssb",
		"amd-psfd",
		"amd-ssbd",
		"amd-stibp",
		"amx-avx512",
		"amx-bf16",
		"amx-bf16-alias",
		"amx-complex",
		"amx-complex-alias",
		"amx-fp16",
		"amx-fp16-alias",
		"amx-fp8",
		"amx-int8",
		"amx-int8-alias",
		"amx-movrs",
		"amx-tf32",
		"amx-tile",
		"apic",
		"apx-nci-ndd-nf",
		"apxf",
		"arat",
		"arch-capabilities",
		"arch-lbr",
		"auto-ibrs",
		"avic",
		"avx",
		"avx-ifma",
		"avx-ne-convert",
		"avx-vnni",
		"avx-vnni-int16",
		"avx-vnni-int8",
		"avx10",
		"avx10-128",
		"avx10-256",
		"avx10-512",
		"avx10-vnni-int",
		"avx2",
		"avx512-4fmaps",
		"avx512-4vnniw",
		"avx512-bf16",
		"avx512-fp16",
		"avx512-vp2intersect",
		"avx512-vpopcntdq",
		"avx512bitalg",
		"avx512bw",
		"avx512cd",
		"avx512dq",
		"avx512er",
		"avx512f",
		"avx512ifma",
		"avx512pf",
		"avx512vbmi",
		"avx512vbmi2",
		"avx512vl",
		"avx512vnni",
		"bhi-ctrl",
		"bhi-no",
		"bmi1",
		"bmi2",
		"bus-lock-detect",
		"cet-ibt",
		"cet-ss",
		"cid",
		"cldemote",
		"clflush",
		"clflushopt",
		"clwb",
		"clzero",
		"cmov",
		"cmp-legacy",
		"cmpccxadd",
		"core-capability",
		"cr8legacy",
		"cx16",
		"cx8",
		"dca",
		"ddpd-u",
		"de",
		"decodeassists",
		"ds",
		"ds-cpl",
		"dtes64",
		"eraps",
		"erms",
		"est",
		"extapic",
		"f16c",
		"fb-clear",
		"fbsdp-no",
		"fdp-excptn-only",
		"flush-l1d",
		"flushbyasid",
		"fma",
		"fma4",
		"fpu",
		"fred",
		"fs-gs-base-ns",
		"fsgsbase",
		"fsrc",
		"fsrm",
		"fsrs",
		"full-width-write",
		"fxsr",
		"fxsr-opt",
		"fzrm",
		"gds-no",
		"gfni",
		"gmet",
		"hle",
		"ht",
		"hypervisor",
		"ia64",
		"ibpb",
		"ibpb-brtype",
		"ibrs",
		"ibrs-all",
		"ibs",
		"intel-psfd",
		"intel-pt",
		"intel-pt-lip",
		"invpcid",
		"invtsc",
		"ipred-ctrl",
		"its-no",
		"kvm-asyncpf",
		"kvm-asyncpf-int",
		"kvm-asyncpf-vmexit",
		"kvm-hint-dedicated",
		"kvm-mmu",
		"kvm-msi-ext-dest-id",
		"kvm-nopiodelay",
		"kvm-poll-control",
		"kvm-pv-eoi",
		"kvm-pv-ipi",
		"kvm-pv-sched-yield",
		"kvm-pv-tlb-flush",
		"kvm-pv-unhalt",
		"kvm-steal-time",
		"kvmclock",
		"kvmclock-stable-bit",
		"la57",
		"lahf-lm",
		"lam",
		"lbrv",
		"lfence-always-serializing",
		"lkgs",
		"lm",
		"lwp",
		"mca",
		"mcdt-no",
		"mce",
		"md-clear",
		"mds-no",
		"misalignsse",
		"mmx",
		"mmxext",
		"monitor",
		"movbe",
		"movdir64b",
		"movdiri",
		"movrs",
		"mpx",
		"msr",
		"msr-imm",
		"mtrr",
		"no-nested-data-bp",
		"nodeid-msr",
		"npt",
		"nrip-save",
		"null-sel-clr-base",
		"nx",
		"osvw",
		"overflow-recov",
		"pae",
		"pat",
		"pause-filter",
		"pbe",
		"pbrsb-no",
		"pcid",
		"pclmulqdq",
		"pcommit",
		"pdcm",
		"pdpe1gb",
		"perfctr-core",
		"perfctr-nb",
		"perfmon-v2",
		"pfthreshold",
		"pge",
		"phe",
		"phe-en",
		"pks",
		"pku",
		"pmm",
		"pmm-en",
		"pn",
		"pni",
		"popcnt",
		"prefetchi",
		"prefetchiti",
		"pschange-mc-no",
		"psdp-no",
		"pse",
		"pse36",
		"rdctl-no",
		"rdpid",
		"rdrand",
		"rdseed",
		"rdtscp",
		"rfds-clear",
		"rfds-no",
		"rrsba-ctrl",
		"rsba",
		"rtm",
		"sbdr-ssdp-no",
		"sbpb",
		"sep",
		"serialize",
		"sgx",
		"sgx-aex-notify",
		"sgx-debug",
		"sgx-edeccssa",
		"sgx-exinfo",
		"sgx-kss",
		"sgx-mode64",
		"sgx-provisionkey",
		"sgx-tokenkey",
		"sgx1",
		"sgx2",
		"sgxlc",
		"sha-ni",
		"sha512",
		"skinit",
		"skip-l1dfl-vmentry",
		"sm3",
		"sm4",
		"smap",
		"smep",
		"smx",
		"spec-ctrl",
		"split-lock-detect",
		"srso-no",
		"srso-user-kernel-no",
		"ss",
		"ssb-no",
		"ssbd",
		"sse",
		"sse2",
		"sse4.1",
		"sse4.2",
		"sse4a",
		"ssse3",
		"stibp",
		"stibp-always-on",
		"succor",
		"svm",
		"svm-lock",
		"svme-addr-chk",
		"syscall",
		"taa-no",
		"tbm",
		"tce",
		"tm",
		"tm2",
		"topoext",
		"tsa-l1-no",
		"tsa-sq-no",
		"tsc",
		"tsc-adjust",
		"tsc-deadline",
		"tsc-scale",
		"tsx-ctrl",
		"tsx-ldtrk",
		"umip",
		"v-vmsave-vmload",
		"vaes",
		"verw-clear",
		"vgif",
		"virt-ssbd",
		"vmcb-clean",
		"vme",
		"vmx",
		"vnmi",
		"vpclmulqdq",
		"waitpkg",
		"wbnoinvd",
		"wdt",
		"wrmsrns",
		"x2apic",
		"xcrypt",
		"xcrypt-en",
		"xfd",
		"xgetbv1",
		"xop",
		"xsave",
		"xsavec",
		"xsaveerptr",
		"xsaveopt",
		"xsaves",
		"xstore",
		"xstore-en",
		"xtpr",
		"zero-fcs-fds",
	).Tag("cpu features").Uid("qemu", "cpu-features")
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
		"xenfb", "Xen framebuffer",
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

// ActionNetdevTypes completes QEMU network backend types
//
//	user (user-mode networking (SLIRP))
//	tap (host TAP network backend)
func ActionNetdevTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"user", "user-mode networking (SLIRP)",
		"tap", "host TAP network backend",
		"bridge", "host TAP connected to bridge",
		"passt", "passt network backend",
		"l2tpv3", "L2TPv3 pseudowire",
		"vde", "VDE switch",
		"af-xdp", "AF_XDP socket",
		"socket", "socket network backend",
		"stream", "stream socket",
		"dgram", "datagram socket",
		"vhost-user", "vhost-user network",
		"vhost-vdpa", "vhost-vDPA network",
		"hubport", "hub port",
	).Tag("netdev types").Uid("qemu", "netdev-types")
}

// ActionBootDrives completes QEMU boot drive letters
//
//	a (floppy)
//	c (hard disk)
func ActionBootDrives() carapace.Action {
	return carapace.ActionValuesDescribed(
		"a", "floppy",
		"b", "floppy",
		"c", "hard disk",
		"d", "CD-ROM",
		"n", "network",
	).Tag("boot drives").Uid("qemu", "boot-drives")
}

// ActionDebugItems completes QEMU debug log items
//
//	cpu (show CPU registers before entering a TB)
//	exec (show trace before each executed TB)
func ActionDebugItems() carapace.Action {
	return carapace.ActionValuesDescribed(
		"out_asm", "show generated host assembly code for each compiled TB",
		"in_asm", "show target assembly code for each compiled TB",
		"op", "show micro ops for each compiled TB",
		"op_opt", "show micro ops after optimization",
		"op_ind", "show micro ops before indirect lowering",
		"op_plugin", "show micro ops before plugin injection",
		"int", "show interrupts/exceptions in short format",
		"exec", "show trace before each executed TB",
		"cpu", "show CPU registers before entering a TB",
		"fpu", "include FPU registers in the 'cpu' logging",
		"mmu", "log MMU-related activities",
		"pcall", "x86 only: show protected mode far calls/returns/exceptions",
		"cpu_reset", "show CPU state before CPU resets",
		"unimp", "log unimplemented functionality",
		"guest_errors", "log when the guest OS does something invalid",
		"page", "dump pages at beginning of user mode emulation",
		"nochain", "do not chain compiled TBs",
		"plugin", "output from TCG plugins",
		"strace", "log every user-mode syscall",
		"tid", "open a separate log file per thread",
		"vpu", "include VPU registers in the 'cpu' logging",
		"invalid_mem", "log invalid memory accesses",
	).Tag("debug items").Uid("qemu", "debug-items")
}

// ActionExportTypes completes QEMU storage daemon export types
//
//	nbd (Network Block Device)
//	fuse (Filesystem in Userspace)
func ActionExportTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"nbd", "Network Block Device",
		"fuse", "Filesystem in Userspace",
		"vhost-user-blk", "vhost-user block device",
		"vduse-blk", "VDUSE block device",
	).Tag("export types").Uid("qemu", "export-types")
}

// ActionKeymapLayouts completes XKB keyboard layouts
//
//	us (English (US))
//	de (German)
func ActionKeymapLayouts() carapace.Action {
	return carapace.ActionExecCommand("awk", "/^! layout/{found=1; next} /^!/{found=0} found && NF{sub(/^[ \\t]+/, \"\"); print}", "/usr/share/X11/xkb/rules/base.lst")(func(output []byte) carapace.Action {
		return parseXkbEntries(string(output)).Tag("layouts").Uid("qemu", "keymap-layouts")
	})
}

// ActionKeymapModels completes XKB keyboard models
//
//	pc105 (Generic 105-key PC)
//	pc104 (Generic 104-key PC)
func ActionKeymapModels() carapace.Action {
	return carapace.ActionExecCommand("awk", "/^! model/{found=1; next} /^!/{found=0} found && NF{sub(/^[ \\t]+/, \"\"); print}", "/usr/share/X11/xkb/rules/base.lst")(func(output []byte) carapace.Action {
		return parseXkbEntries(string(output)).Tag("models").Uid("qemu", "keymap-models")
	})
}

// ActionKeymapVariants completes XKB keyboard variants for a given layout
//
//	dvorak (Dvorak)
//	colemak (Colemak)
func ActionKeymapVariants() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		layout := c.Value
		if layout == "" {
			layout = "us"
		}
		return carapace.ActionExecCommand("awk", fmt.Sprintf("/^! variant/{found=1; next} /^!/{found=0} found && NF && $2 ~ /^%s:/{sub(/^[ \\t]+/, \"\"); gsub(/ %s:/, \" \", $0); print}", layout, layout), "/usr/share/X11/xkb/rules/base.lst")(func(output []byte) carapace.Action {
			return parseXkbEntries(string(output)).Tag("variants").Uid("qemu", "keymap-variants")
		})
	})
}

// ActionKeymapOptions completes XKB keyboard options
//
//	grp:alt_shift_toggle (Alt+Shift toggle)
//	caps:escape (Caps Lock acts as Escape)
func ActionKeymapOptions() carapace.Action {
	return carapace.ActionExecCommand("awk", "/^! option/{found=1; next} /^!/{found=0} found && NF{sub(/^[ \\t]+/, \"\"); print}", "/usr/share/X11/xkb/rules/base.lst")(func(output []byte) carapace.Action {
		return parseXkbEntries(string(output)).Tag("options").Uid("qemu", "keymap-options")
	})
}

func parseXkbEntries(output string) carapace.Action {
	lines := strings.Split(output, "\n")
	vals := make([]string, 0, len(lines)*2)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 2)
		vals = append(vals, parts[0])
		if len(parts) > 1 {
			vals = append(vals, strings.TrimSpace(parts[1]))
		} else {
			vals = append(vals, parts[0])
		}
	}
	return carapace.ActionValuesDescribed(vals...)
}
