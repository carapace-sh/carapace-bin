package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugin_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a plugin to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_pushCmd).Standalone()

	plugin_pushCmd.Flags().Bool("disable-content-trust", false, "Skip image signing (default true)")
	pluginCmd.AddCommand(plugin_pushCmd)

	carapace.Gen(plugin_pushCmd).PositionalCompletion(
		action.ActionPlugins(),
	)
}
