package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_help_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify that all records in the store can be decrypted with the current key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_help_verifyCmd).Standalone()

	store_helpCmd.AddCommand(store_help_verifyCmd)
}
