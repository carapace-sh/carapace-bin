package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_deleteAllSessionsCmd = &cobra.Command{
	Use:   "delete-all-sessions",
	Short: "Delete all sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_deleteAllSessionsCmd).Standalone()

	helpCmd.AddCommand(help_deleteAllSessionsCmd)
}
