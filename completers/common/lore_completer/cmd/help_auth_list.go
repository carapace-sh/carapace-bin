package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_auth_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all stored authentication identities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_auth_listCmd).Standalone()

	help_authCmd.AddCommand(help_auth_listCmd)
}
