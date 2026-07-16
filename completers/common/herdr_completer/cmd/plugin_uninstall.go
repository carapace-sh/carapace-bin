package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var plugin_uninstallCmd = &cobra.Command{
	Use:   "uninstall <plugin>",
	Short: "Uninstall a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_uninstallCmd).Standalone()

	pluginCmd.AddCommand(plugin_uninstallCmd)

	carapace.Gen(plugin_uninstallCmd).PositionalCompletion(herdr.ActionPlugins())
}
