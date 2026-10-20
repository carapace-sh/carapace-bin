package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var plugins_uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "Remove an installed plugin",
	Aliases: []string{"remove", "rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_uninstallCmd).Standalone()

	plugins_uninstallCmd.Flags().BoolP("all", "a", false, "Remove all plugins")
	plugins_uninstallCmd.Flags().BoolP("help", "h", false, "Print help")
	plugins_uninstallCmd.Flags().BoolP("purge", "p", false, "Also remove a tool plugin's installs, downloads, and cache")
	pluginsCmd.AddCommand(plugins_uninstallCmd)

	carapace.Gen(plugins_uninstallCmd).PositionalAnyCompletion(
		mise.ActionPlugins(),
	)
}
