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
		"none", "no device allocated",
		"null", "discard all output",
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
		"cocoa", "macOS Cocoa window",
		"curses", "curses/ncurses interface",
		"dbus", "D-Bus display",
		"egl-headless", "EGL headless rendering",
		"gl=core", "OpenGL core profile",
		"gl=es", "OpenGL ES",
		"gl=on", "OpenGL rendering",
		"gtk", "GTK window",
		"none", "no display output",
		"sdl", "SDL window",
	).Tag("display types").Uid("qemu", "display-type")
}

// ActionVgaTypes completes QEMU VGA types
//
//	cirrus (Cirrus Logic GD5446)
//	std (standard VGA)
func ActionVgaTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"cg3", "CG3 (Sun)",
		"cirrus", "Cirrus Logic GD5446",
		"none", "no VGA",
		"qxl", "QXL (SPICE)",
		"std", "standard VGA",
		"tcx", "TCX (Sun)",
		"virtio", "VirtIO GPU",
		"vmware", "VMware SVGA-II",
		"xenfb", "Xen framebuffer",
	).Tag("vga types").Uid("qemu", "vga-type")
}

// ActionWatchdogActions completes QEMU watchdog actions
//
//	reset (reset the guest)
//	shutdown (shutdown the guest)
func ActionWatchdogActions() carapace.Action {
	return carapace.ActionValuesDescribed(
		"debug", "print debug message",
		"inject-nmi", "inject an NMI into the guest",
		"none", "do nothing",
		"pause", "pause the guest",
		"poweroff", "power off the guest",
		"reset", "reset the guest",
		"shutdown", "shutdown the guest",
	).Tag("watchdog actions").Uid("qemu", "watchdog-action")
}
