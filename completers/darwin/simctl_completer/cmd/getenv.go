package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var getenvCmd = &cobra.Command{
	Use:   "getenv",
	Short: "Print an environment variable from a running device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getenvCmd).Standalone()
	rootCmd.AddCommand(getenvCmd)
	carapace.Gen(getenvCmd).PositionalCompletion(
		simctl.ActionDevicesByState("Booted"),
	)
}
