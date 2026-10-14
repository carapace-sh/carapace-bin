package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var joinPaneCmd = &cobra.Command{
	Use:     "join-pane",
	Aliases: []string{"joinp"},
	Short:   "split a pane and move an existing one into the new space",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinPaneCmd).Standalone()

	joinPaneCmd.Flags().BoolS("b", "b", false, "join source pane left of or above target pane")
	joinPaneCmd.Flags().BoolS("d", "d", false, "don't make the new window become the active one")
	joinPaneCmd.Flags().BoolS("f", "f", false, "span the full size")
	joinPaneCmd.Flags().BoolS("h", "h", false, "split horizontally")
	joinPaneCmd.Flags().StringS("l", "l", "", "define new pane's size")
	joinPaneCmd.Flags().StringS("s", "s", "", "specify source pane")
	joinPaneCmd.Flags().StringS("t", "t", "", "specify target pane")
	joinPaneCmd.Flags().BoolS("v", "v", false, "split vertically")
	rootCmd.AddCommand(joinPaneCmd)

	carapace.Gen(joinPaneCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionPanes(),
		"t": tmux.ActionPanes(),
	})
}
