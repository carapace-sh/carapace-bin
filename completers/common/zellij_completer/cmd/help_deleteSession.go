package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_deleteSessionCmd = &cobra.Command{
	Use:   "delete-session",
	Short: "Delete a specific session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_deleteSessionCmd).Standalone()

	helpCmd.AddCommand(help_deleteSessionCmd)
}
