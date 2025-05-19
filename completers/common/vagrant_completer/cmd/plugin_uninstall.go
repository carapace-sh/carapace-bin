package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var plugin_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_uninstallCmd).Standalone()

	plugin_uninstallCmd.Flags().Bool("local", false, "Remove plugin from local project")
	pluginCmd.AddCommand(plugin_uninstallCmd)

	carapace.Gen(plugin_uninstallCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return vagrant.ActionPlugins(vagrant.PluginOpts{
				Local:  plugin_updateCmd.Flag("local").Changed,
				Global: !plugin_updateCmd.Flag("local").Changed,
			}).Invoke(c).Filter(c.Args...).ToA()
		}),
	)
}
