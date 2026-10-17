package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var sessionRenameCmd = &cobra.Command{
	Use:   "rename [SESSION] LABEL",
	Short: "Rename a session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sessionRenameCmd).Standalone()

	sessionCmd.AddCommand(sessionRenameCmd)

	carapace.Gen(sessionRenameCmd).PositionalCompletion(
		action.ActionSessions(sessionRenameCmd),
	)
}
