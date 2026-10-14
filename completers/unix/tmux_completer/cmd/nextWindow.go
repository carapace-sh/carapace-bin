package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var nextWindowCmd = &cobra.Command{
	Use:     "next-window",
	Aliases: []string{"next"},
	Short:   "move to the next window in a session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nextWindowCmd).Standalone()

	nextWindowCmd.Flags().BoolS("a", "a", false, "move to the next window with an alert")
	nextWindowCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(nextWindowCmd)

	carapace.Gen(nextWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
