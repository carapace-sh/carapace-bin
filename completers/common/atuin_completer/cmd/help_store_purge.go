package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Delete all records in the store that cannot be decrypted with the current key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_purgeCmd).Standalone()

	help_storeCmd.AddCommand(help_store_purgeCmd)
}
