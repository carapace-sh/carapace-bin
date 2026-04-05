package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Build and activate the new configuration, and make it the boot default",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_switchCmd).Standalone()

	addOsRebuildFlags(os_switchCmd)
	os_switchCmd.Flags().Bool("show-activation-logs", false, "Show activation logs")

	osCmd.AddCommand(os_switchCmd)
}
