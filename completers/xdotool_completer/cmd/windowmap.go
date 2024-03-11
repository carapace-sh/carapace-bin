package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowmapCmd = &cobra.Command{
	Use:   "windowmap",
	Short: "Map a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowmapCmd).Standalone()

	windowmapCmd.Flags().Bool("sync", false, "After requesting the window map, wait until the window is actually mapped (visible)")
	rootCmd.AddCommand(windowmapCmd)

	carapace.Gen(windowmapCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
