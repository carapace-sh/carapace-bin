package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sessionInspectCmd = &cobra.Command{
	Use:   "inspect [SESSION]",
	Short: "Inspect a session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sessionInspectCmd).Standalone()

	sessionInspectCmd.Flags().Bool("json", false, "output session information as JSON")

	sessionCmd.AddCommand(sessionInspectCmd)

	carapace.Gen(sessionInspectCmd).PositionalCompletion(
		action.ActionSessions(sessionInspectCmd),
	)
}
