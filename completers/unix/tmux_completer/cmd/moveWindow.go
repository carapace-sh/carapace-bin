package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var moveWindowCmd = &cobra.Command{
	Use:     "move-window",
	Aliases: []string{"movew"},
	Short:   "move a window to another",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moveWindowCmd).Standalone()

	moveWindowCmd.Flags().BoolS("a", "a", false, "move window to next index after destination window")
	moveWindowCmd.Flags().BoolS("b", "b", false, "move window to next index before destination window")
	moveWindowCmd.Flags().BoolS("d", "d", false, "don't make the new window become the active one")
	moveWindowCmd.Flags().BoolS("k", "k", false, "kill the target window if it exists")
	moveWindowCmd.Flags().BoolS("r", "r", false, "renumber windows in session in sequential order")
	moveWindowCmd.Flags().StringS("s", "s", "", "specify source window")
	moveWindowCmd.Flags().StringS("t", "t", "", "specify destination window")
	rootCmd.AddCommand(moveWindowCmd)

	carapace.Gen(moveWindowCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionWindows(),
		"t": tmux.ActionWindows(),
	})
}
