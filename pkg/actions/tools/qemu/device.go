package qemu

import "github.com/carapace-sh/carapace"

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
	).Tag("char devices").Uid("qemu", "char-device")
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
	).Tag("display types").Uid("qemu", "display-type")
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
	).Tag("vga types").Uid("qemu", "vga-type")
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
	).Tag("watchdog actions").Uid("qemu", "watchdog-action")
}
