package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowfocusCmd = &cobra.Command{
	Use:   "windowfocus",
	Short: "Focus a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowfocusCmd).Standalone()

	windowfocusCmd.Flags().Bool("sync", false, "After sending the window focus request, wait until the window is actually focused")
	rootCmd.AddCommand(windowfocusCmd)

	carapace.Gen(windowfocusCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
