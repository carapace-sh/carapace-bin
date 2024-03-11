package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list installed Helm plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_listCmd).Standalone()
	pluginCmd.AddCommand(plugin_listCmd)
}
