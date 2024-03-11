package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/helm"
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
		helm.ActionPlugins().FilterArgs(),
	)
}
