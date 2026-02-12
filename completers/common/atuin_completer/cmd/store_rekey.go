package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_rekeyCmd = &cobra.Command{
	Use:   "rekey",
	Short: "Re-encrypt the store with a new key (potential for data loss!)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_rekeyCmd).Standalone()

	store_rekeyCmd.Flags().BoolP("help", "h", false, "Print help")
	storeCmd.AddCommand(store_rekeyCmd)
}
