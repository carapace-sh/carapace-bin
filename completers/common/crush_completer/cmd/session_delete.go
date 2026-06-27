package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/crush"
	"github.com/spf13/cobra"
)

var session_deleteCmd = &cobra.Command{
	Use:     "delete <id>",
	Short:   "Delete a session",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(session_deleteCmd).Standalone()

	session_deleteCmd.Flags().Bool("json", false, "output in JSON format")
	sessionCmd.AddCommand(session_deleteCmd)

	carapace.Gen(session_deleteCmd).PositionalCompletion(
		crush.ActionSessions(),
	)
}
