package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var clearHistoryCmd = &cobra.Command{
	Use:     "clear-history",
	Aliases: []string{"clearhist"},
	Short:   "remove and clear history for a pane",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearHistoryCmd).Standalone()

	clearHistoryCmd.Flags().BoolS("H", "H", false, "also remove all hyperlinks")
	clearHistoryCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(clearHistoryCmd)

	carapace.Gen(clearHistoryCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})
}
