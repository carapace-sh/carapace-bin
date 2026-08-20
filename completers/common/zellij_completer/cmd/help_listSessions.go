package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_listSessionsCmd = &cobra.Command{
	Use:   "list-sessions",
	Short: "List active sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_listSessionsCmd).Standalone()

	helpCmd.AddCommand(help_listSessionsCmd)
}
