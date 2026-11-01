package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a plugin (supports plugin@marketplace)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginInstallCmd).Standalone()
	pluginCmd.AddCommand(pluginInstallCmd)

	carapace.Gen(pluginInstallCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
