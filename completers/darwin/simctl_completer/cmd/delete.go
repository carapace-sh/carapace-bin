package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specified devices, unavailable devices, or all devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()
	rootCmd.AddCommand(deleteCmd)
	carapace.Gen(deleteCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionValues("all", "unavailable"),
			simctl.ActionDevices(),
		).ToA(),
	)
}
