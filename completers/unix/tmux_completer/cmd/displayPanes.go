package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var displayPanesCmd = &cobra.Command{
	Use:     "display-panes",
	Aliases: []string{"displayp"},
	Short:   "display an indicator for each visible pane",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayPanesCmd).Standalone()

	displayPanesCmd.Flags().BoolS("N", "N", false, "do not close when a key is pressed")
	displayPanesCmd.Flags().BoolS("Z", "Z", false, "start unzoomed")
	displayPanesCmd.Flags().StringS("d", "d", "", "time to show indicator for")
	displayPanesCmd.Flags().BoolS("k", "k", false, "kill the pane when the mode is exited")
	displayPanesCmd.Flags().StringS("s", "s", "", "display panes from source-window")
	displayPanesCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(displayPanesCmd)

	carapace.Gen(displayPanesCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionWindows(),
		"t": tmux.ActionPanes(),
	})
}
