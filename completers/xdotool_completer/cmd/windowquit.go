package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowquitCmd = &cobra.Command{
	Use:   "windowquit",
	Short: "Close a window gracefully",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowquitCmd).Standalone()

	rootCmd.AddCommand(windowquitCmd)

	carapace.Gen(windowquitCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
