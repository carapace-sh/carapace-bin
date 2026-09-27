package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_store_accountCmd = &cobra.Command{
	Use:   "account",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_store_accountCmd).Standalone()

	object_storeCmd.AddCommand(object_store_accountCmd)
}
