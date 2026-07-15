package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_enableCmd = &cobra.Command{
	Use:   "enable <plugin_id>",
	Short: "Enable a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_enableCmd).Standalone()

	pluginCmd.AddCommand(plugin_enableCmd)

	carapace.Gen(plugin_enableCmd).PositionalCompletion(action.ActionPlugins())
}
