package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/qemu-system-x86_64_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qemu-system-x86_64",
	Short: "QEMU system emulator",
	Long:  "https://www.qemu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("S", "S", false, "freeze CPU at startup")
	rootCmd.Flags().BoolS("h", "h", false, "display help and exit")
	rootCmd.Flags().BoolS("s", "s", false, "shorthand for -gdb tcp::1234")

	rootCmd.Flags().StringS("D", "D", "", "log to file")
	rootCmd.Flags().StringS("L", "L", "", "BIOS/VGA BIOS directory")
	rootCmd.Flags().StringS("M", "M", "", "short form of -machine")
	rootCmd.Flags().StringS("accel", "accel", "", "select accelerator")
	rootCmd.Flags().StringS("acpitable", "acpitable", "", "add ACPI table")
	rootCmd.Flags().StringS("action", "action", "", "set actions for events")
	rootCmd.Flags().StringS("add-fd", "add-fd", "", "add a file descriptor to an fd set")
	rootCmd.Flags().StringS("append", "append", "", "use cmdline as kernel command line")
	rootCmd.Flags().StringS("audio", "audio", "", "configure audio backend and hardware")
	rootCmd.Flags().StringArrayS("audiodev", "audiodev", nil, "add a new audio backend driver")
	rootCmd.Flags().StringS("bios", "bios", "", "set the filename for the BIOS")
	rootCmd.Flags().StringArrayS("blockdev", "blockdev", nil, "define a new block driver node")
	rootCmd.Flags().StringS("boot", "boot", "", "specify boot order")
	rootCmd.Flags().StringS("cdrom", "cdrom", "", "use file as CD-ROM image")
	rootCmd.Flags().StringArrayS("chardev", "chardev", nil, "add a character device backend")
	rootCmd.Flags().StringS("compat", "compat", "", "set policy for handling unstable management interfaces")
	rootCmd.Flags().StringS("cpu", "cpu", "", "select CPU model")
	rootCmd.Flags().StringS("d", "d", "", "enable logging of specified items")
	rootCmd.Flags().BoolS("daemonize", "daemonize", false, "daemonize the QEMU process after initialization")
	rootCmd.Flags().StringS("debugcon", "debugcon", "", "redirect the debug console to host device")
	rootCmd.Flags().StringArrayS("device", "device", nil, "add device driver")
	rootCmd.Flags().StringS("dfilter", "dfilter", "", "limit logging to specified address range")
	rootCmd.Flags().StringS("display", "display", "", "select display type")
	rootCmd.Flags().StringArrayS("drive", "drive", nil, "use file as a drive image")
	rootCmd.Flags().StringS("dtb", "dtb", "", "use file as device tree blob")
	rootCmd.Flags().BoolS("dump-vmstate", "dump-vmstate", false, "dump vmstate to file")
	rootCmd.Flags().StringS("echr", "echr", "", "set the escape character for the monitor")
	rootCmd.Flags().BoolS("enable-kvm", "enable-kvm", false, "enable KVM full virtualization support")
	rootCmd.Flags().BoolS("enable-sync-profile", "enable-sync-profile", false, "enable synchronization profiling")
	rootCmd.Flags().StringS("fda", "fda", "", "use file as floppy disk 0 image")
	rootCmd.Flags().StringS("fdb", "fdb", "", "use file as floppy disk 1 image")
	rootCmd.Flags().StringArrayS("fsdev", "fsdev", nil, "define a new filesystem device")
	rootCmd.Flags().BoolS("full-screen", "full-screen", false, "start in full screen")
	rootCmd.Flags().StringS("fw_cfg", "fw_cfg", "", "add a fw_cfg item")
	rootCmd.Flags().StringS("g", "g", "", "wait for gdb connection on dev")
	rootCmd.Flags().StringS("gdb", "gdb", "", "wait for gdb connection on dev")
	rootCmd.Flags().StringS("global", "global", "", "set a global default for a driver property")
	rootCmd.Flags().StringS("hda", "hda", "", "use file as hard disk 0 image")
	rootCmd.Flags().StringS("hdb", "hdb", "", "use file as hard disk 1 image")
	rootCmd.Flags().StringS("hdc", "hdc", "", "use file as hard disk 2 image")
	rootCmd.Flags().StringS("hdd", "hdd", "", "use file as hard disk 3 image")
	rootCmd.Flags().StringS("icount", "icount", "", "enable instruction counting for record/replay")
	rootCmd.Flags().StringS("incoming", "incoming", "", "configure incoming migration")
	rootCmd.Flags().StringS("initrd", "initrd", "", "use file as initial ramdisk")
	rootCmd.Flags().StringS("iscsi", "iscsi", "", "use iSCSI as backing store")
	rootCmd.Flags().BoolS("jitdump", "jitdump", false, "output jitdump data for perf")
	rootCmd.Flags().StringS("k", "k", "", "use keyboard layout language")
	rootCmd.Flags().StringS("kernel", "kernel", "", "use file as kernel image")
	rootCmd.Flags().StringS("loadvm", "loadvm", "", "restore from named snapshot")
	rootCmd.Flags().StringS("m", "m", "", "configure guest RAM")
	rootCmd.Flags().StringS("machine", "machine", "", "select emulated machine")
	rootCmd.Flags().StringS("mem-path", "mem-path", "", "set the path for backing RAM file")
	rootCmd.Flags().BoolS("mem-prealloc", "mem-prealloc", false, "preallocate guest memory")
	rootCmd.Flags().StringS("mon", "mon", "", "configure a monitor")
	rootCmd.Flags().StringS("monitor", "monitor", "", "redirect the monitor to host device")
	rootCmd.Flags().StringS("msg", "msg", "", "control message format")
	rootCmd.Flags().StringS("mtdblock", "mtdblock", "", "use file as flash memory image")
	rootCmd.Flags().StringS("name", "name", "", "set the name of the guest")
	rootCmd.Flags().StringS("net", "net", "", "configure network device")
	rootCmd.Flags().StringArrayS("netdev", "netdev", nil, "configure network backend")
	rootCmd.Flags().StringS("nic", "nic", "", "configure default NIC")
	rootCmd.Flags().BoolS("no-fd-bootchk", "no-fd-bootchk", false, "disable boot signature checking for floppies")
	rootCmd.Flags().BoolS("no-reboot", "no-reboot", false, "exit instead of rebooting")
	rootCmd.Flags().BoolS("no-shutdown", "no-shutdown", false, "don't exit on guest shutdown")
	rootCmd.Flags().BoolS("no-user-config", "no-user-config", false, "don't load user-provided config files")
	rootCmd.Flags().BoolS("nodefaults", "nodefaults", false, "don't create default devices")
	rootCmd.Flags().BoolS("nographic", "nographic", false, "disable graphical output")
	rootCmd.Flags().StringArrayS("numa", "numa", nil, "configure NUMA topology")
	rootCmd.Flags().StringArrayS("object", "object", nil, "create a new object of type <type>")
	rootCmd.Flags().BoolS("only-migratable", "only-migratable", false, "only allow migratable devices")
	rootCmd.Flags().StringS("option-rom", "option-rom", "", "add an option ROM")
	rootCmd.Flags().StringS("overcommit", "overcommit", "", "set overcommit policy")
	rootCmd.Flags().StringS("parallel", "parallel", "", "redirect the parallel port to host device")
	rootCmd.Flags().BoolS("perfmap", "perfmap", false, "output perfmap data for perf")
	rootCmd.Flags().StringArrayS("pflash", "pflash", nil, "use file as parallel flash image")
	rootCmd.Flags().StringS("pidfile", "pidfile", "", "write process ID to a file")
	rootCmd.Flags().StringS("plugin", "plugin", "", "load plugin")
	rootCmd.Flags().Bool("preconfig", false, "pause QEMU before machine is initialized (experimental)")
	rootCmd.Flags().StringS("prom-env", "prom-env", "", "set OpenBIOS PROM variables")
	rootCmd.Flags().StringS("qmp", "qmp", "", "redirect QMP to host device")
	rootCmd.Flags().StringS("qmp-pretty", "qmp-pretty", "", "redirect QMP with pretty-printing to host device")
	rootCmd.Flags().StringS("readconfig", "readconfig", "", "read config from file")
	rootCmd.Flags().StringS("rtc", "rtc", "", "configure real-time clock")
	rootCmd.Flags().StringS("run-with", "run-with", "", "set isolation features")
	rootCmd.Flags().StringS("sandbox", "sandbox", "", "enable seccomp sandbox")
	rootCmd.Flags().StringS("sd", "sd", "", "use file as Secure Digital card image")
	rootCmd.Flags().StringS("seed", "seed", "", "seed for random number generator")
	rootCmd.Flags().BoolS("semihosting", "semihosting", false, "enable semihosting")
	rootCmd.Flags().StringS("semihosting-config", "semihosting-config", "", "configure semihosting")
	rootCmd.Flags().StringArrayS("serial", "serial", nil, "redirect the serial port to host device")
	rootCmd.Flags().StringS("set", "set", "", "set a driver property")
	rootCmd.Flags().StringS("shim", "shim", "", "load a UEFI shim")
	rootCmd.Flags().StringS("smbios", "smbios", "", "add SMBIOS entry")
	rootCmd.Flags().StringS("smp", "smp", "", "set the number of CPUs and CPU topology")
	rootCmd.Flags().BoolS("snapshot", "snapshot", false, "write to temporary files")
	rootCmd.Flags().StringS("spice", "spice", "", "configure SPICE server")
	rootCmd.Flags().StringS("tpmdev", "tpmdev", "", "configure TPM backend")
	rootCmd.Flags().StringS("trace", "trace", "", "specify tracing options")
	rootCmd.Flags().BoolS("usb", "usb", false, "enable USB emulation")
	rootCmd.Flags().StringS("usbdevice", "usbdevice", "", "add USB device")
	rootCmd.Flags().StringS("uuid", "uuid", "", "set the UUID of the guest")
	rootCmd.Flags().BoolS("version", "version", false, "display version information and exit")
	rootCmd.Flags().StringS("vga", "vga", "", "select VGA type")
	rootCmd.Flags().StringS("virtfs", "virtfs", "", "configure VirtFS export")
	rootCmd.Flags().StringS("vnc", "vnc", "", "configure VNC server")
	rootCmd.Flags().StringS("watchdog-action", "watchdog-action", "", "set watchdog action")
	rootCmd.Flags().BoolS("win2k-hack", "win2k-hack", false, "use Windows 2000 boot hack")
	rootCmd.Flags().BoolS("xen-attach", "xen-attach", false, "attach to Xen domain")
	rootCmd.Flags().StringS("xen-domid", "xen-domid", "", "specify Xen domain ID")
	rootCmd.Flags().BoolS("xen-domid-restrict", "xen-domid-restrict", false, "restrict set of available xen operations")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"M":     carapace.ActionCallback(func(c carapace.Context) carapace.Action { return action.ActionMachines(rootCmd) }),
		"accel": qemu.ActionAccels(),
		"action": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"reboot", "action when guest reboots",
						"shutdown", "action when guest shuts down",
						"panic", "action when guest panics",
						"watchdog", "action when watchdog fires",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "reboot":
						return carapace.ActionValues("reset", "shutdown")
					case "shutdown":
						return carapace.ActionValues("poweroff", "pause")
					case "panic":
						return carapace.ActionValues("pause", "shutdown", "exit-failure", "none")
					case "watchdog":
						return qemu.ActionWatchdogActions()
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"bios": carapace.ActionFiles(),
		"boot": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"order", "boot drive order",
						"once", "boot drive order for first boot",
						"menu", "boot menu",
						"splash", "splash image file",
						"splash-time", "splash time in ms",
						"reboot-timeout", "reboot timeout in ms",
						"strict", "strict boot",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "order", "once":
						return qemu.ActionBootDrives()
					case "menu", "strict":
						return carapace.ActionValues("on", "off")
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"cdrom": carapace.ActionFiles(),
		"compat": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"deprecated-input", "policy for deprecated input",
						"deprecated-output", "policy for deprecated output",
						"unstable-input", "policy for unstable input",
						"unstable-output", "policy for unstable output",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "deprecated-input":
						return carapace.ActionValues("accept", "reject", "crash")
					case "deprecated-output":
						return carapace.ActionValues("accept", "hide")
					case "unstable-input":
						return carapace.ActionValues("accept", "reject", "crash")
					case "unstable-output":
						return carapace.ActionValues("accept", "hide")
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"cpu": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					return action.ActionCpuModels(rootCmd).Suffix(",").NoSpace(',')
				})
			}
			return qemu.ActionCpuFeatures().Prefix("+").Suffix(",").NoSpace(',')
		}),
		"d": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action { return qemu.ActionDebugItems() }),
		"debugcon": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return qemu.ActionCharDevices().Suffix(":").NoSpace(':')
			default:
				return carapace.ActionValues()
			}
		}),
		"display": qemu.ActionDisplayTypes(),
		"dtb":     carapace.ActionFiles(),
		"fda":     carapace.ActionFiles(),
		"fdb":     carapace.ActionFiles(),
		"fw_cfg":  carapace.ActionFiles(),
		"gdb": carapace.ActionMultiPartsN(":", 3, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return qemu.ActionCharDevices().Suffix(":").NoSpace(':')
			case 1:
				switch c.Parts[0] {
				case "tcp":
					return carapace.ActionValues("localhost", "0.0.0.0", "::").NoSpace(':')
				case "unix":
					return carapace.ActionFiles()
				case "chardev":
					return carapace.ActionValues()
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"hda": carapace.ActionFiles(),
		"hdb": carapace.ActionFiles(),
		"hdc": carapace.ActionFiles(),
		"hdd": carapace.ActionFiles(),
		"icount": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"shift", "number of clock ticks per instruction (2^N)",
						"align", "align host and virtual clocks",
						"sleep", "real time CPU sleeping",
						"rr", "record/replay mode",
						"rrfile", "record/replay file",
						"rrsnapshot", "record/replay snapshot name",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "shift":
						return carapace.ActionValues("auto")
					case "align", "sleep":
						return carapace.ActionValues("on", "off")
					case "rr":
						return carapace.ActionValues("record", "replay")
					case "rrfile":
						return carapace.ActionFiles()
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"incoming": carapace.ActionValues("tcp:", "rdma:", "unix:", "fd:", "file:", "exec:", "defer").NoSpace(':'),
		"initrd":   carapace.ActionFiles(),
		"iscsi":    carapace.ActionValues(),
		"kernel":   carapace.ActionFiles(),
		"loadvm":   carapace.ActionValues(),
		"m": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"size", "initial amount of guest memory",
						"slots", "number of hotplug slots",
						"maxmem", "maximum amount of guest memory",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "size", "maxmem":
						return carapace.ActionValues("128M", "256M", "512M", "1G", "2G", "4G", "8G", "16G", "32G", "64G", "128G")
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"machine":  carapace.ActionCallback(func(c carapace.Context) carapace.Action { return action.ActionMachines(rootCmd) }),
		"mem-path": carapace.ActionDirectories(),
		"mon": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"chardev", "character device name",
						"mode", "monitor mode",
						"pretty", "pretty-print JSON",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "mode":
						return carapace.ActionValues("readline", "control")
					case "pretty":
						return carapace.ActionValues("on", "off")
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"monitor": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return qemu.ActionCharDevices().Suffix(":").NoSpace(':')
			default:
				return carapace.ActionValues()
			}
		}),
		"msg": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"timestamp", "enable timestamps",
						"guest-name", "enable guest name prefix",
					).Suffix("=").NoSpace('=')
				default:
					return carapace.ActionValues("on", "off")
				}
			})
		}),
		"mtdblock": carapace.ActionFiles(),
		"name": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"process", "process name",
						"debug-threads", "name individual threads",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "debug-threads":
						return carapace.ActionValues("on", "off")
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"nic": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				return qemu.ActionNetdevTypes().Suffix(",").NoSpace(',')
			}
			return carapace.ActionValues()
		}),
		"overcommit": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"mem-lock", "memory lock support",
						"cpu-pm", "CPU power management",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "mem-lock":
						return carapace.ActionValues("on", "off", "on-fault")
					case "cpu-pm":
						return carapace.ActionValues("on", "off")
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"parallel": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return qemu.ActionCharDevices().Suffix(":").NoSpace(':')
			case 1:
				switch c.Parts[0] {
				case "file":
					return carapace.ActionFiles()
				case "chardev":
					return carapace.ActionValues()
				case "unix":
					return carapace.ActionFiles()
				case "parport":
					return carapace.ActionFiles("/dev/parport*")
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"pflash":  carapace.ActionFiles(),
		"pidfile": carapace.ActionFiles(),
		"plugin":  carapace.ActionFiles(".so"),
		"qmp": carapace.ActionMultiPartsN(":", 3, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return qemu.ActionCharDevices().Suffix(":").NoSpace(':')
			case 1:
				switch c.Parts[0] {
				case "tcp":
					return carapace.ActionValues("localhost", "0.0.0.0", "::").NoSpace(':')
				case "unix":
					return carapace.ActionFiles()
				case "chardev":
					return carapace.ActionValues()
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"qmp-pretty": carapace.ActionMultiPartsN(":", 3, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return qemu.ActionCharDevices().Suffix(":").NoSpace(':')
			case 1:
				switch c.Parts[0] {
				case "tcp":
					return carapace.ActionValues("localhost", "0.0.0.0", "::").NoSpace(':')
				case "unix":
					return carapace.ActionFiles()
				case "chardev":
					return carapace.ActionValues()
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"readconfig": carapace.ActionFiles(),
		"rtc": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"base", "RTC base time",
						"clock", "RTC clock source",
						"driftfix", "drift fix for clock ticks",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "base":
						return carapace.ActionValues("utc", "localtime")
					case "clock":
						return carapace.ActionValues("host", "rt", "vm")
					case "driftfix":
						return carapace.ActionValues("none", "slew")
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"run-with": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"async-teardown", "enable asynchronous teardown",
						"chroot", "chroot to directory before starting VM",
						"exit-with-parent", "exit when parent process exits",
						"user", "switch to user before starting VM",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "async-teardown", "exit-with-parent":
						return carapace.ActionValues("on", "off")
					case "chroot":
						return carapace.ActionDirectories()
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"sandbox": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					if c.Value == "" {
						return carapace.ActionValues("on", "off").Suffix(",").NoSpace(',')
					}
					return carapace.ActionValuesDescribed(
						"obsolete", "allow obsolete system calls",
						"elevateprivileges", "allow or deny privilege elevation",
						"spawn", "allow or deny spawning new threads/processes",
						"resourcecontrol", "allow or deny resource control",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "obsolete":
						return carapace.ActionValues("allow", "deny")
					case "elevateprivileges":
						return carapace.ActionValues("allow", "deny", "children")
					case "spawn":
						return carapace.ActionValues("allow", "deny")
					case "resourcecontrol":
						return carapace.ActionValues("allow", "deny")
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"sd": carapace.ActionFiles(),
		"semihosting-config": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"enable", "enable semihosting",
						"target", "target for semihosting",
						"chardev", "use a named character device",
						"userspace", "enable userspace semihosting",
						"arg", "argument string",
					).Suffix("=").NoSpace('=')
				default:
					switch c.Parts[0] {
					case "enable":
						return carapace.ActionValues("on", "off")
					case "target":
						return carapace.ActionValues("native", "gdb", "auto")
					case "userspace":
						return carapace.ActionValues("on", "off")
					case "chardev":
						return carapace.ActionValues()
					default:
						return carapace.ActionValues()
					}
				}
			})
		}),
		"serial": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return qemu.ActionCharDevices().Suffix(":").NoSpace(':')
			case 1:
				switch c.Parts[0] {
				case "file":
					return carapace.ActionFiles()
				case "chardev":
					return carapace.ActionValues()
				case "unix":
					return carapace.ActionFiles()
				case "serial":
					return carapace.ActionFiles("/dev/ttyS*")
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"smp": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValuesDescribed(
						"cpus", "total number of guest CPUs",
						"maxcpus", "maximum CPUs including hotplug",
						"sockets", "physical sockets",
						"dies", "dies per socket",
						"clusters", "clusters per die",
						"modules", "modules per cluster",
						"cores", "cores per module",
						"threads", "threads per core",
						"books", "books per drawer",
						"drawers", "drawers",
					).Suffix("=").NoSpace('=')
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"spice":           carapace.ActionValues(),
		"tpmdev":          carapace.ActionValues(),
		"trace":           carapace.ActionValues(),
		"usbdevice":       carapace.ActionValues("tablet", "mouse", "keyboard", "host:", "bt:"),
		"uuid":            carapace.ActionValues(),
		"vga":             qemu.ActionVgaTypes(),
		"vnc":             carapace.ActionValues(),
		"watchdog-action": qemu.ActionWatchdogActions(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
