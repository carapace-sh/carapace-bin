package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var swapWindowCmd = &cobra.Command{
	Use:     "swap-window",
	Aliases: []string{"swapw"},
	Short:   "swap two windows",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swapWindowCmd).Standalone()

	swapWindowCmd.Flags().BoolS("d", "d", false, "don't make the new window become the active one")
	swapWindowCmd.Flags().StringS("s", "s", "", "specify source window")
	swapWindowCmd.Flags().StringS("t", "t", "", "specify destination window")
	rootCmd.AddCommand(swapWindowCmd)

	carapace.Gen(swapWindowCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionWindows(),
		"t": tmux.ActionWindows(),
	})
}
