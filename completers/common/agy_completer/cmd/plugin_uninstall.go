package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/agy"
	"github.com/spf13/cobra"
)

var pluginUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginUninstallCmd).Standalone()
	pluginCmd.AddCommand(pluginUninstallCmd)

	carapace.Gen(pluginUninstallCmd).PositionalCompletion(
		agy.ActionPlugins(),
	)
}
