package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/qemu"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qemu-storage-daemon",
	Short: "QEMU storage daemon",
	Long:  "https://www.qemu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringArray("blockdev", nil, "configure a block backend")
	rootCmd.Flags().StringArray("chardev", nil, "configure a character device backend")
	rootCmd.Flags().Bool("daemonize", false, "daemonize the process, and have the parent exit once startup is complete")
	rootCmd.Flags().StringArray("export", nil, "export the specified block node")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().StringArray("monitor", nil, "configure a QMP monitor")
	rootCmd.Flags().String("nbd-server", "", "start an NBD server for exporting block nodes")
	rootCmd.Flags().StringArray("object", nil, "create a new object of type <type>")
	rootCmd.Flags().String("pidfile", "", "write process ID to a file after startup")
	rootCmd.Flags().StringP("trace", "T", "", "specify tracing options")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"blockdev": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("driver=").NoSpace()
			default:
				if c.Parts[0] == "driver=" || c.Parts[len(c.Parts)-1] == "" {
					return qemu.ActionImageFormats().Suffix(",")
				}
				return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
					switch len(c.Parts) {
					case 0:
						return carapace.ActionValues(
							"auto-read-only=",
							"cache.direct=",
							"cache.no-flush=",
							"detect-zeroes=",
							"discard=",
							"driver=",
							"file=",
							"force-share=",
							"node-name=",
							"read-only=",
						).NoSpace()
					case 1:
						switch c.Parts[0] {
						case "driver":
							return qemu.ActionImageFormats().Suffix(",")
						case "discard":
							return carapace.ActionValues("ignore", "unmap").Suffix(",")
						case "detect-zeroes":
							return carapace.ActionValues("on", "off", "unmap").Suffix(",")
						case "cache.direct", "cache.no-flush", "force-share", "read-only", "auto-read-only":
							return carapace.ActionValues("on", "off").Suffix(",")
						default:
							return carapace.ActionValues()
						}
					default:
						return carapace.ActionValues()
					}
				})
			}
		}),
		"chardev": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return qemu.ActionCharDevices().NoSpace()
				case 1:
					return carapace.ActionValues()
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"export": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValues(
						"addr.path=",
						"addr.type=",
						"allow-other=",
						"bitmap=",
						"growable=",
						"id=",
						"logical-block-size=",
						"mountpoint=",
						"name=",
						"node-name=",
						"num-queues=",
						"queue-size=",
						"serial=",
						"type=",
						"writable=",
					).NoSpace()
				case 1:
					switch c.Parts[0] {
					case "type":
						return qemu.ActionExportTypes().Suffix(",")
					case "addr.type":
						return carapace.ActionValues("unix", "inet", "fd").Suffix(",")
					case "writable", "growable", "allow-other":
						return carapace.ActionValues("on", "off", "auto").Suffix(",")
					case "mountpoint":
						return carapace.ActionFiles().Suffix(",")
					case "addr.path":
						return carapace.ActionFiles().Suffix(",")
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"monitor": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValues(
						"chardev=",
						"mode=",
						"pretty=",
					).NoSpace()
				case 1:
					switch c.Parts[0] {
					case "mode":
						return carapace.ActionValues("control", "readline").Suffix(",")
					case "pretty":
						return carapace.ActionValues("on", "off").Suffix(",")
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"nbd-server": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValues(
						"addr.host=",
						"addr.path=",
						"addr.port=",
						"addr.type=",
						"max-connections=",
						"tls-authz=",
						"tls-creds=",
					).NoSpace()
				case 1:
					switch c.Parts[0] {
					case "addr.type":
						return carapace.ActionValues("inet", "unix").Suffix(",")
					case "addr.path":
						return carapace.ActionFiles().Suffix(",")
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"object": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValues(
						"id=",
						"qom-type=",
					).NoSpace()
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"pidfile": carapace.ActionFiles(),
		"trace": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValues(
						"enable=",
						"events=",
						"file=",
					).NoSpace()
				case 1:
					switch c.Parts[0] {
					case "events", "file":
						return carapace.ActionFiles().Suffix(",")
					default:
						return carapace.ActionValues()
					}
				default:
					return carapace.ActionValues()
				}
			})
		}),
	})
}
