package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored authentication identities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_help_listCmd).Standalone()

	auth_helpCmd.AddCommand(auth_help_listCmd)
}
