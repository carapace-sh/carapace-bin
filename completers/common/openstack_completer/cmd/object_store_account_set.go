package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_store_account_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set account properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_store_account_setCmd).Standalone()

	object_store_account_setCmd.Flags().String("property", "", "Set a property on this account (repeat option to set multiple properties)")
	object_store_account_setCmd.MarkFlagRequired("property")
	object_store_accountCmd.AddCommand(object_store_account_setCmd)
}
