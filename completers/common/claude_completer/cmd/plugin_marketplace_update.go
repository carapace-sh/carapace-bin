package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_marketplace_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update marketplace(s) from their source - updates all if no name specified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_marketplace_updateCmd).Standalone()

	plugin_marketplace_updateCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_marketplace_updateCmd.Flags().Bool("json", false, "With a marketplace name: print one machine-readable result line as the last line on stdout…")
	plugin_marketplaceCmd.AddCommand(plugin_marketplace_updateCmd)
}
