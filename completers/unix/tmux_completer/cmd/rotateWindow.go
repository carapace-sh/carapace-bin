package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var rotateWindowCmd = &cobra.Command{
	Use:     "rotate-window",
	Aliases: []string{"rotatew"},
	Short:   "rotate positions of panes in a window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rotateWindowCmd).Standalone()

	rotateWindowCmd.Flags().BoolS("D", "D", false, "rotate downward")
	rotateWindowCmd.Flags().BoolS("U", "U", false, "rotate upward")
	rotateWindowCmd.Flags().BoolS("Z", "Z", false, "keep the window zoomed if it was zoomed")
	rotateWindowCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(rotateWindowCmd)

	carapace.Gen(rotateWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
