package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var plugin_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_removeCmd).Standalone()

	pluginCmd.AddCommand(plugin_removeCmd)

	carapace.Gen(plugin_removeCmd).PositionalCompletion(
		yarn.ActionPlugins(),
	)
}
