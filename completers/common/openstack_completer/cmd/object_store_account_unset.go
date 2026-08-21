package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_store_account_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset account properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_store_account_unsetCmd).Standalone()

	object_store_account_unsetCmd.Flags().String("property", "", "Property to remove from account (repeat option to remove multiple properties)")
	object_store_account_unsetCmd.MarkFlagRequired("property")
	object_store_accountCmd.AddCommand(object_store_account_unsetCmd)
}
