package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var plugin_enableCmd = &cobra.Command{
	Use:   "enable <plugin_id>",
	Short: "Enable a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_enableCmd).Standalone()

	pluginCmd.AddCommand(plugin_enableCmd)

	carapace.Gen(plugin_enableCmd).PositionalCompletion(herdr.ActionPlugins())
}
