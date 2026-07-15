package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_configDirCmd = &cobra.Command{
	Use:   "config-dir <plugin_id>",
	Short: "Print a plugin config directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_configDirCmd).Standalone()

	pluginCmd.AddCommand(plugin_configDirCmd)

	carapace.Gen(plugin_configDirCmd).PositionalCompletion(action.ActionPlugins())
}
