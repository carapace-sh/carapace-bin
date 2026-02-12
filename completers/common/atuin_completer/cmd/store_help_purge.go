package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_help_purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Delete all records in the store that cannot be decrypted with the current key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_help_purgeCmd).Standalone()

	store_helpCmd.AddCommand(store_help_purgeCmd)
}
