package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

func AddScopeFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("block", "b", "", "target a `block`")
	cmd.Flags().StringP("session", "s", "", "scope to session SESSION")
	cmd.Flags().StringP("window", "w", "", "scope to window WINDOW")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"block":   action.ActionBlocks(cmd),
		"session": action.ActionSessions(cmd),
		"window":  action.ActionWindows(cmd),
	})
}

func AddWindowScopeFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("session", "s", "", "scope to session SESSION")
	cmd.Flags().StringP("window", "w", "", "scope to window WINDOW")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(cmd),
		"window":  action.ActionWindows(cmd),
	})
}
