package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update one or more Helm plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_updateCmd).Standalone()
	pluginCmd.AddCommand(plugin_updateCmd)

	carapace.Gen(plugin_updateCmd).PositionalAnyCompletion(
		action.ActionPlugins().FilterArgs(),
	)
}
