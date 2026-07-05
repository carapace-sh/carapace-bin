package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var pairCmd = &cobra.Command{
	Use:   "pair",
	Short: "Create a new watch and phone pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pairCmd).Standalone()
	rootCmd.AddCommand(pairCmd)
	carapace.Gen(pairCmd).PositionalCompletion(
		simctl.ActionDevices(),
		simctl.ActionDevices(),
	)
}
