package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_marketplaceCmd = &cobra.Command{
	Use:   "marketplace",
	Short: "Manage Claude Code marketplaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_marketplaceCmd).Standalone()

	plugin_marketplaceCmd.Flags().BoolP("help", "h", false, "Display help for command")
	pluginCmd.AddCommand(plugin_marketplaceCmd)
}
