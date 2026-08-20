package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_launchPluginCmd = &cobra.Command{
	Use:   "launch-plugin",
	Short: "Returns: Plugin pane ID (format: plugin_<id>)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_launchPluginCmd).Standalone()

	help_actionCmd.AddCommand(help_action_launchPluginCmd)
}
