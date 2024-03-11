package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowkillCmd = &cobra.Command{
	Use:   "windowkill",
	Short: "Kill a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowkillCmd).Standalone()

	rootCmd.AddCommand(windowkillCmd)

	carapace.Gen(windowkillCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
