package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify that all records in the store can be decrypted with the current key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_verifyCmd).Standalone()

	help_storeCmd.AddCommand(help_store_verifyCmd)
}
