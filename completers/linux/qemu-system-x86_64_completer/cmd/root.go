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

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"M":       carapace.ActionCallback(func(c carapace.Context) carapace.Action { return action.ActionMachines(rootCmd) }),
		"accel":   qemu.ActionAccels(),
		"bios":    carapace.ActionFiles(),
		"cdrom":   carapace.ActionFiles(),
		"cpu":     carapace.ActionCallback(func(c carapace.Context) carapace.Action { return action.ActionCpuModels(rootCmd) }),
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
		"hda":      carapace.ActionFiles(),
		"hdb":      carapace.ActionFiles(),
		"hdc":      carapace.ActionFiles(),
		"hdd":      carapace.ActionFiles(),
		"initrd":   carapace.ActionFiles(),
		"iscsi":    carapace.ActionValues(),
		"kernel":   carapace.ActionFiles(),
		"loadvm":   carapace.ActionValues(),
		"machine":  carapace.ActionCallback(func(c carapace.Context) carapace.Action { return action.ActionMachines(rootCmd) }),
		"mem-path": carapace.ActionDirectories(),
		"mon":      carapace.ActionValues(),
		"monitor": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return qemu.ActionCharDevices().Suffix(":").NoSpace(':')
			default:
				return carapace.ActionValues()
			}
		}),
		"mtdblock": carapace.ActionFiles(),
		"name":     carapace.ActionValues(),
		"nic":      carapace.ActionValues("user", "tap", "bridge", "socket", "vde", "none"),
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
		"sd":         carapace.ActionFiles(),
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
		"smp":             carapace.ActionValues(),
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
