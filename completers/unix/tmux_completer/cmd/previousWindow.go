package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var previousWindowCmd = &cobra.Command{
	Use:     "previous-window",
	Aliases: []string{"prev"},
	Short:   "move to the previous window in a session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(previousWindowCmd).Standalone()

	previousWindowCmd.Flags().BoolS("a", "a", false, "move to the previous window with an alert")
	previousWindowCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(previousWindowCmd)

	carapace.Gen(previousWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
