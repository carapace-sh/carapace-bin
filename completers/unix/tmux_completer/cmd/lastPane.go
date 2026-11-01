package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var lastPaneCmd = &cobra.Command{
	Use:     "last-pane",
	Aliases: []string{"lastp"},
	Short:   "select the previously selected pane",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lastPaneCmd).Standalone()

	lastPaneCmd.Flags().BoolS("Z", "Z", false, "keep window zoomed if it was zoomed")
	lastPaneCmd.Flags().BoolS("d", "d", false, "disable input to the pane")
	lastPaneCmd.Flags().BoolS("e", "e", false, "enable input to the pane")
	lastPaneCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(lastPaneCmd)

	carapace.Gen(lastPaneCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
