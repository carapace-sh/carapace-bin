package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_disableCmd = &cobra.Command{
	Use:   "disable <plugin_id>",
	Short: "Disable a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_disableCmd).Standalone()

	pluginCmd.AddCommand(plugin_disableCmd)

	carapace.Gen(plugin_disableCmd).PositionalCompletion(action.ActionPlugins())
}
