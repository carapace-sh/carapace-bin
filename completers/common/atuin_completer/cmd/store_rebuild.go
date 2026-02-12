package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a store (eg atuin store rebuild history)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_rebuildCmd).Standalone()

	store_rebuildCmd.Flags().BoolP("help", "h", false, "Print help")
	storeCmd.AddCommand(store_rebuildCmd)
}
