package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugins_installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a plugin",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_installCmd).Standalone()

	plugins_installCmd.Flags().BoolP("all", "a", false, "Install all missing plugins from config files")
	plugins_installCmd.Flags().BoolP("force", "f", false, "Reinstall plugin if already installed")
	pluginsCmd.AddCommand(plugins_installCmd)

	carapace.Gen(plugins_installCmd).PositionalCompletion(
		action.ActionPlugins(),
	)
}
