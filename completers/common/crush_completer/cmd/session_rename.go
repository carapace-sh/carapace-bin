package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/crush"
	"github.com/spf13/cobra"
)

var session_renameCmd = &cobra.Command{
	Use:   "rename <id> <title>",
	Short: "Rename a session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(session_renameCmd).Standalone()

	session_renameCmd.Flags().Bool("json", false, "output in JSON format")
	sessionCmd.AddCommand(session_renameCmd)

	carapace.Gen(session_renameCmd).PositionalCompletion(
		crush.ActionSessions(),
	)
}
