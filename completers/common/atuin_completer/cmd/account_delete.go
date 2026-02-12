package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var account_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your account, and all synced data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(account_deleteCmd).Standalone()

	account_deleteCmd.Flags().BoolP("help", "h", false, "Print help")
	accountCmd.AddCommand(account_deleteCmd)
}
