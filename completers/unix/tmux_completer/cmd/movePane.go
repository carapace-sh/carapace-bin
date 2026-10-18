package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var movePaneCmd = &cobra.Command{
	Use:     "move-pane",
	Aliases: []string{"movep"},
	Short:   "split a pane and move an existing one into the new space",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(movePaneCmd).Standalone()

	movePaneCmd.Flags().StringS("D", "D", "", "move floating pane down by lines")
	movePaneCmd.Flags().StringS("L", "L", "", "move floating pane left by columns")
	movePaneCmd.Flags().BoolS("M", "M", false, "begin a mouse drag")
	movePaneCmd.Flags().StringS("P", "P", "", "move floating pane to position")
	movePaneCmd.Flags().StringS("R", "R", "", "move floating pane right by columns")
	movePaneCmd.Flags().StringS("U", "U", "", "move floating pane up by lines")
	movePaneCmd.Flags().StringS("X", "X", "", "move to absolute X position")
	movePaneCmd.Flags().StringS("Y", "Y", "", "move to absolute Y position")
	movePaneCmd.Flags().BoolS("b", "b", false, "move source pane left of or above target pane")
	movePaneCmd.Flags().BoolS("d", "d", false, "don't change the active pane")
	movePaneCmd.Flags().BoolS("f", "f", false, "span the full size")
	movePaneCmd.Flags().BoolS("h", "h", false, "split horizontally")
	movePaneCmd.Flags().StringS("l", "l", "", "define new pane's size")
	movePaneCmd.Flags().StringS("s", "s", "", "specify source pane")
	movePaneCmd.Flags().StringS("t", "t", "", "specify target pane")
	movePaneCmd.Flags().BoolS("v", "v", false, "split vertically")
	movePaneCmd.Flags().StringS("z", "z", "", "move to z-index")
	rootCmd.AddCommand(movePaneCmd)

	carapace.Gen(movePaneCmd).FlagCompletion(carapace.ActionMap{
		"P": carapace.ActionValues("top-left", "top-centre", "top-right", "centre-left", "centre", "centre-right", "bottom-left", "bottom-centre", "bottom-right", "front", "back", "forward", "backward", "forward-loop", "backward-loop"),
		"s": tmux.ActionPanes(),
		"t": tmux.ActionPanes(),
	})
}
