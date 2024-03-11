package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_rmCmd = &cobra.Command{
	Use:     "rm [OPTIONS] PLUGIN [PLUGIN...]",
	Short:   "Remove one or more plugins",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_rmCmd).Standalone()

	plugin_rmCmd.Flags().BoolP("force", "f", false, "Force the removal of an active plugin")
	pluginCmd.AddCommand(plugin_rmCmd)

	carapace.Gen(plugin_rmCmd).PositionalAnyCompletion(docker.ActionPlugins())
}
