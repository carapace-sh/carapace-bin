package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowcloseCmd = &cobra.Command{
	Use:   "windowclose",
	Short: "Close a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowcloseCmd).Standalone()

	rootCmd.AddCommand(windowcloseCmd)

	carapace.Gen(windowcloseCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
