package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var resizeWindowCmd = &cobra.Command{
	Use:     "resize-window",
	Aliases: []string{"resizew"},
	Short:   "resize a window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resizeWindowCmd).Standalone()

	resizeWindowCmd.Flags().BoolS("A", "A", false, "set size of largest session containing the window")
	resizeWindowCmd.Flags().BoolS("D", "D", false, "resize downward")
	resizeWindowCmd.Flags().BoolS("L", "L", false, "resize to the left")
	resizeWindowCmd.Flags().BoolS("R", "R", false, "resize to the right")
	resizeWindowCmd.Flags().BoolS("U", "U", false, "resize upward")
	resizeWindowCmd.Flags().BoolS("a", "a", false, "set size of smallest session containing the window")
	resizeWindowCmd.Flags().StringS("t", "t", "", "specify target window")
	resizeWindowCmd.Flags().StringS("x", "x", "", "specify width")
	resizeWindowCmd.Flags().StringS("y", "y", "", "specify height")
	rootCmd.AddCommand(resizeWindowCmd)

	carapace.Gen(resizeWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
