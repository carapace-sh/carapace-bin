package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_listCmd).Standalone()

	plugin_listCmd.Flags().Bool("available", false, "Include available plugins from marketplaces (requires --json)")
	plugin_listCmd.Flags().String("data-size", "", "Measure each installed plugin's saved data directory, or only the named plugin's (requires --json)")
	plugin_listCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_listCmd.Flags().Bool("json", false, "Output as JSON")
	plugin_listCmd.Flag("data-size").NoOptDefVal = " "
	pluginCmd.AddCommand(plugin_listCmd)
}
