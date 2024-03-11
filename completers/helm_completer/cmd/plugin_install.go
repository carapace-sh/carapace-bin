package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var plugin_installCmd = &cobra.Command{
	Use:   "install",
	Short: "install one or more Helm plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_installCmd).Standalone()
	plugin_installCmd.Flags().String("version", "", "specify a version constraint. If this is not specified, the latest version is installed")
	pluginCmd.AddCommand(plugin_installCmd)

	carapace.Gen(plugin_installCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionRepositorySearch(git.SearchOpts{}.Default()),
			carapace.ActionDirectories(),
		).ToA(),
	)
}
