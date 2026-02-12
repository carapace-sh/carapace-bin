package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_account_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your account, and all synced data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_account_deleteCmd).Standalone()

	help_accountCmd.AddCommand(help_account_deleteCmd)
}
