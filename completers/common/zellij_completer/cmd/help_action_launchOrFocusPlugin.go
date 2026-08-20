package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_launchOrFocusPluginCmd = &cobra.Command{
	Use:   "launch-or-focus-plugin",
	Short: "Returns: Plugin pane ID (format: plugin_<id>) when creating or focusing plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_launchOrFocusPluginCmd).Standalone()

	help_actionCmd.AddCommand(help_action_launchOrFocusPluginCmd)
}
