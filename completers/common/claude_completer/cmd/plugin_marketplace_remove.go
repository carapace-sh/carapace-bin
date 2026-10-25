package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_marketplace_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a configured marketplace",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_marketplace_removeCmd).Standalone()

	plugin_marketplace_removeCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_marketplace_removeCmd.Flags().Bool("json", false, "Print one machine-readable result line as the last line on stdout (same exit codes)")
	plugin_marketplace_removeCmd.Flags().String("scope", "", "Remove the marketplace declaration from a specific settings scope: user, project, or local")
	plugin_marketplaceCmd.AddCommand(plugin_marketplace_removeCmd)

	carapace.Gen(plugin_marketplace_removeCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("local", "project", "user"),
	})
}
