package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qemu-vmsr-helper",
	Short: "Virtual MSR helper program for QEMU",
	Long:  "https://www.qemu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("daemon", "d", false, "run in the background")
	rootCmd.Flags().StringP("group", "g", "", "group to drop privileges to")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().StringP("pidfile", "f", "", "PID file when running as a daemon")
	rootCmd.Flags().StringP("socket", "k", "", "path to the unix socket")
	rootCmd.Flags().StringP("trace", "T", "", "specify tracing options")
	rootCmd.Flags().StringP("user", "u", "", "user to drop privileges to")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"group":   os.ActionGroups(),
		"pidfile": carapace.ActionFiles(),
		"socket":  carapace.ActionFiles(),
		"user":    os.ActionUsers(),
	})
}
