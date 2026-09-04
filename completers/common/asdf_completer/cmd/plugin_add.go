package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a plugin from the plugin repo OR a Git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_addCmd).Standalone()

	pluginCmd.AddCommand(plugin_addCmd)
}
