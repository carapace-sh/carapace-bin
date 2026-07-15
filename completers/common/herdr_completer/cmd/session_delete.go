package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var session_deleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete a stopped session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(session_deleteCmd).Standalone()

	session_deleteCmd.Flags().Bool("json", false, "")
	sessionCmd.AddCommand(session_deleteCmd)

	carapace.Gen(session_deleteCmd).PositionalCompletion(action.ActionSessions())
}
