package cmd

import (
	"github.com/carapace-sh/carapace"
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
}
