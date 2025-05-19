package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var plugin_licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "install license",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_licenseCmd).Standalone()

	pluginCmd.AddCommand(plugin_licenseCmd)

	carapace.Gen(plugin_licenseCmd).PositionalCompletion(
		vagrant.ActionPlugins(vagrant.PluginOpts{Global: true}),
	)
}
