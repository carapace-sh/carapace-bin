package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var swapPaneCmd = &cobra.Command{
	Use:     "swap-pane",
	Aliases: []string{"swapp"},
	Short:   "swap two panes",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swapPaneCmd).Standalone()

	swapPaneCmd.Flags().BoolS("D", "D", false, "move pane down")
	swapPaneCmd.Flags().BoolS("U", "U", false, "move pane up")
	swapPaneCmd.Flags().BoolS("Z", "Z", false, "keep the window zoomed if it was zoomed")
	swapPaneCmd.Flags().BoolS("d", "d", false, "don't change the active pane")
	swapPaneCmd.Flags().StringS("s", "s", "", "specify source pane")
	swapPaneCmd.Flags().StringS("t", "t", "", "specify destination pane")
	rootCmd.AddCommand(swapPaneCmd)

	carapace.Gen(swapPaneCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionPanes(),
		"t": tmux.ActionPanes(),
	})
}
