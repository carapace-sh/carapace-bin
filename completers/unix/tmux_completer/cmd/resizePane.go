package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var resizePaneCmd = &cobra.Command{
	Use:     "resize-pane",
	Aliases: []string{"resizep"},
	Short:   "resize a pane",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resizePaneCmd).Standalone()

	resizePaneCmd.Flags().BoolS("D", "D", false, "resize downward")
	resizePaneCmd.Flags().BoolS("L", "L", false, "resize to the left")
	resizePaneCmd.Flags().BoolS("M", "M", false, "begin mouse resizing")
	resizePaneCmd.Flags().BoolS("R", "R", false, "resize to the right")
	resizePaneCmd.Flags().BoolS("T", "T", false, "trim lines below the cursor position")
	resizePaneCmd.Flags().BoolS("U", "U", false, "resize upward")
	resizePaneCmd.Flags().BoolS("Z", "Z", false, "toggle zoom of pane")
	resizePaneCmd.Flags().StringS("t", "t", "", "specify target pane")
	resizePaneCmd.Flags().StringS("x", "x", "", "specify width")
	resizePaneCmd.Flags().StringS("y", "y", "", "specify height")
	rootCmd.AddCommand(resizePaneCmd)

	carapace.Gen(resizePaneCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})
}
