package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shut down a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shutdownCmd).Standalone()
	rootCmd.AddCommand(shutdownCmd)
	carapace.Gen(shutdownCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionValues("all"),
			simctl.ActionDevicesByState("Booted"),
		).ToA(),
	)
}
