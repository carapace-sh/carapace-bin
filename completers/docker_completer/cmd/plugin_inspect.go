package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var plugin_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] PLUGIN [PLUGIN...]",
	Short: "Display detailed information on one or more plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_inspectCmd).Standalone()

	plugin_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	pluginCmd.AddCommand(plugin_inspectCmd)

	carapace.Gen(plugin_inspectCmd).PositionalAnyCompletion(docker.ActionPlugins())
}
