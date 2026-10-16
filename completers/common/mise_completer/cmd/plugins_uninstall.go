package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugins_uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "Removes a plugin",
	Aliases: []string{"remove", "rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_uninstallCmd).Standalone()

	plugins_uninstallCmd.Flags().BoolP("all", "a", false, "Remove all plugins")
	plugins_uninstallCmd.Flags().BoolP("purge", "p", false, "Also remove all associated tool versions")
	pluginsCmd.AddCommand(plugins_uninstallCmd)

	carapace.Gen(plugins_uninstallCmd).PositionalAnyCompletion(
		action.ActionPlugins(),
	)
}
