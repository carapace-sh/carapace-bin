package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a store (eg atuin store rebuild history)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_rebuildCmd).Standalone()

	help_storeCmd.AddCommand(help_store_rebuildCmd)
}
