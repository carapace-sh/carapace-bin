package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "Build the new configuration and make it the boot default",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_bootCmd).Standalone()

	addOsRebuildFlags(os_bootCmd)
	os_bootCmd.Flags().Bool("show-activation-logs", false, "Show activation logs")

	osCmd.AddCommand(os_bootCmd)
}
