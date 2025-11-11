package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_pushCmd = &cobra.Command{
	Use:   "push [OPTIONS] PLUGIN[:TAG]",
	Short: "Push a plugin to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_pushCmd).Standalone()

	plugin_pushCmd.Flags().Bool("disable-content-trust", false, "Skip image verification (deprecated)")
	plugin_pushCmd.Flag("disable-content-trust").Hidden = true
	pluginCmd.AddCommand(plugin_pushCmd)

	carapace.Gen(plugin_pushCmd).PositionalCompletion(
		docker.ActionPlugins(),
	)
}
