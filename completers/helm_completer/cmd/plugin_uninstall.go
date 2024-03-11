package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/helm"
	"github.com/spf13/cobra"
)

var plugin_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall one or more Helm plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_uninstallCmd).Standalone()
	pluginCmd.AddCommand(plugin_uninstallCmd)

	carapace.Gen(plugin_uninstallCmd).PositionalCompletion(
		helm.ActionPlugins(),
	)
}
