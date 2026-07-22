package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_swapCmd = &cobra.Command{
	Use:   "swap",
	Short: "Swap panes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_swapCmd).Standalone()

	pane_swapCmd.Flags().Bool("current", false, "")
	pane_swapCmd.Flags().String("direction", "", "")
	pane_swapCmd.Flags().String("pane", "", "")
	pane_swapCmd.Flags().String("source-pane", "", "")
	pane_swapCmd.Flags().String("target-pane", "", "")
	paneCmd.AddCommand(pane_swapCmd)

	carapace.Gen(pane_swapCmd).FlagCompletion(carapace.ActionMap{
		"direction":   carapace.ActionValues("left", "right", "up", "down"),
		"pane":        herdr.ActionPanes(herdr.PaneOpts{}),
		"source-pane": herdr.ActionPanes(herdr.PaneOpts{}),
		"target-pane": herdr.ActionPanes(herdr.PaneOpts{}),
	})
}
