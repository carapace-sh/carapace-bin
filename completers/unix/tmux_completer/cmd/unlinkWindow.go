package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var unlinkWindowCmd = &cobra.Command{
	Use:     "unlink-window",
	Aliases: []string{"unlinkw"},
	Short:   "unlink a window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkWindowCmd).Standalone()

	unlinkWindowCmd.Flags().BoolS("k", "k", false, "kill the window if it is only in one session")
	unlinkWindowCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(unlinkWindowCmd)

	carapace.Gen(unlinkWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
