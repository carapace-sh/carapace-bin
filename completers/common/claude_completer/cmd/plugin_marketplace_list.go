package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_marketplace_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured marketplaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_marketplace_listCmd).Standalone()

	plugin_marketplace_listCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_marketplace_listCmd.Flags().Bool("json", false, "Output as JSON")
	plugin_marketplaceCmd.AddCommand(plugin_marketplace_listCmd)
}
