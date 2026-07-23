package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_paneCmd = &cobra.Command{
	Use:   "pane",
	Short: "Manage plugin-owned panes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_paneCmd).Standalone()

	pluginCmd.AddCommand(plugin_paneCmd)
}
