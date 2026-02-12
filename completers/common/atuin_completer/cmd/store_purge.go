package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Delete all records in the store that cannot be decrypted with the current key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_purgeCmd).Standalone()

	store_purgeCmd.Flags().BoolP("help", "h", false, "Print help")
	storeCmd.AddCommand(store_purgeCmd)
}
