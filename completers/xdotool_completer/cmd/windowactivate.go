package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowactivateCmd = &cobra.Command{
	Use:   "windowactivate",
	Short: "Activate the window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowactivateCmd).Standalone()

	windowactivateCmd.Flags().Bool("sync", false, "After sending the window activation, wait until the window is actually activated")
	rootCmd.AddCommand(windowactivateCmd)

	carapace.Gen(windowactivateCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
