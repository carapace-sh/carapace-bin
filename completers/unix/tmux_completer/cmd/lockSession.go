package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var lockSessionCmd = &cobra.Command{
	Use:     "lock-session",
	Aliases: []string{"locks"},
	Short:   "lock all clients attached to a session",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockSessionCmd).Standalone()

	lockSessionCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(lockSessionCmd)

	carapace.Gen(lockSessionCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
