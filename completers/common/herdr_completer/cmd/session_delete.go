package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
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

	carapace.Gen(session_deleteCmd).PositionalCompletion(herdr.ActionSessions())
}
