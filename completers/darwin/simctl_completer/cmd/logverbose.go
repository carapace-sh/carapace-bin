package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var logverboseCmd = &cobra.Command{
	Use:   "logverbose",
	Short: "Enable or disable verbose logging for a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logverboseCmd).Standalone()
	rootCmd.AddCommand(logverboseCmd)
	carapace.Gen(logverboseCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("enable", "disable"),
	)
}
