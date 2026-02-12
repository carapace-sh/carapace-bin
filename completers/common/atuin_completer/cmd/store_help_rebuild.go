package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_help_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a store (eg atuin store rebuild history)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_help_rebuildCmd).Standalone()

	store_helpCmd.AddCommand(store_help_rebuildCmd)
}
