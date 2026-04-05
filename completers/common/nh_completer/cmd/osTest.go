package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_testCmd = &cobra.Command{
	Use:   "test",
	Short: "Build and activate the new configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_testCmd).Standalone()

	addOsRebuildFlags(os_testCmd)
	os_testCmd.Flags().Bool("show-activation-logs", false, "Show activation logs")

	osCmd.AddCommand(os_testCmd)
}
