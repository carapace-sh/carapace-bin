package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sessionMoveCmd = &cobra.Command{
	Use:   "move [SESSION] (--before S | --after S | --index N)",
	Short: "Move a session to a new place in the session order",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sessionMoveCmd).Standalone()

	sessionMoveCmd.Flags().String("after", "", "move immediately after session S")
	sessionMoveCmd.Flags().String("before", "", "move immediately before session S")
	sessionMoveCmd.Flags().String("index", "", "move to one-based position N")

	sessionCmd.AddCommand(sessionMoveCmd)

	carapace.Gen(sessionMoveCmd).FlagCompletion(carapace.ActionMap{
		"after":  action.ActionSessions(sessionMoveCmd),
		"before": action.ActionSessions(sessionMoveCmd),
	})

	carapace.Gen(sessionMoveCmd).PositionalAnyCompletion(
		action.ActionSessions(sessionMoveCmd),
	)
}
