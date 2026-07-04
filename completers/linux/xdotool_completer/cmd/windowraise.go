package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowraiseCmd = &cobra.Command{
	Use:   "windowraise",
	Short: "Raise the window to the top of the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowraiseCmd).Standalone()

	rootCmd.AddCommand(windowraiseCmd)

	carapace.Gen(windowraiseCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
