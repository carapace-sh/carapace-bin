package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var selectWindowCmd = &cobra.Command{
	Use:     "select-window",
	Aliases: []string{"selectw"},
	Short:   "select a window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectWindowCmd).Standalone()

	selectWindowCmd.Flags().BoolS("T", "T", false, "if selected window is the current behave like last-window")
	selectWindowCmd.Flags().BoolS("l", "l", false, "behave like last-window")
	selectWindowCmd.Flags().BoolS("n", "n", false, "behave like next-window")
	selectWindowCmd.Flags().BoolS("p", "p", false, "behave like previous-window")
	selectWindowCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(selectWindowCmd)

	carapace.Gen(selectWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
