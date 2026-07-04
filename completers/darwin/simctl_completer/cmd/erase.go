package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var eraseCmd = &cobra.Command{
	Use:   "erase",
	Short: "Erase a device's contents and settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eraseCmd).Standalone()
	rootCmd.AddCommand(eraseCmd)
	carapace.Gen(eraseCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionValues("all"),
			simctl.ActionDevices(),
		).ToA(),
	)
}
