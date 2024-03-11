package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowmoveCmd = &cobra.Command{
	Use:   "windowmove",
	Short: "Move the window to the given position",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowmoveCmd).Standalone()

	windowmoveCmd.Flags().Bool("relative", false, "Make movement relative to the current window position")
	windowmoveCmd.Flags().Bool("sync", false, "After sending the window move request, wait until the window is actually moved")
	rootCmd.AddCommand(windowmoveCmd)

	carapace.Gen(windowmoveCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
