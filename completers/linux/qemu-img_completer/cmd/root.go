package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qemu-img",
	Short: "QEMU disk image utility",
	Long:  "https://www.qemu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "display help and exit")
	rootCmd.Flags().StringP("trace", "T", "", "specify tracing options")
	rootCmd.Flags().BoolP("version", "V", false, "display version info and exit")
}
