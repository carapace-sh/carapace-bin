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

	plugin_listCmd.Flags().Bool("refs", false, "Show Refs")
	plugin_listCmd.Flags().Bool("urls", false, "Show URLs")
	pluginCmd.AddCommand(plugin_listCmd)
}
