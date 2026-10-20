package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Show a plugin's component inventory and projected token cost",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_detailsCmd).Standalone()

	plugin_detailsCmd.Flags().BoolP("help", "h", false, "Display help for command")
	pluginCmd.AddCommand(plugin_detailsCmd)
}
