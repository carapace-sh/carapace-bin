package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Change settings for a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_setCmd).Standalone()

	pluginCmd.AddCommand(plugin_setCmd)

	carapace.Gen(plugin_setCmd).PositionalCompletion(action.ActionPlugins())
}
