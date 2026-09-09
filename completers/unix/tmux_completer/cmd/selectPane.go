package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var selectPaneCmd = &cobra.Command{
	Use:     "select-pane",
	Aliases: []string{"selectp"},
	Short:   "make a pane the active one in the window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectPaneCmd).Standalone()

	selectPaneCmd.Flags().BoolS("D", "D", false, "move to the pane below target")
	selectPaneCmd.Flags().BoolS("L", "L", false, "move to the pane left of target")
	selectPaneCmd.Flags().BoolS("M", "M", false, "clear marked pane")
	selectPaneCmd.Flags().BoolS("R", "R", false, "move to the pane right of target")
	selectPaneCmd.Flags().StringS("T", "T", "", "set the pane title")
	selectPaneCmd.Flags().BoolS("U", "U", false, "move to the pane above target")
	selectPaneCmd.Flags().BoolS("Z", "Z", false, "keep the window zoomed if it was zoomed")
	selectPaneCmd.Flags().BoolS("d", "d", false, "disable input to the pane")
	selectPaneCmd.Flags().BoolS("e", "e", false, "enable input to the pane")
	selectPaneCmd.Flags().BoolS("l", "l", false, "behave like last-pane")
	selectPaneCmd.Flags().BoolS("m", "m", false, "set marked pane")
	selectPaneCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(selectPaneCmd)

	carapace.Gen(selectPaneCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})
}
