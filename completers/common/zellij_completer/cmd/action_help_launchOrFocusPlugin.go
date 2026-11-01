package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_launchOrFocusPluginCmd = &cobra.Command{
	Use:   "launch-or-focus-plugin",
	Short: "Returns: Plugin pane ID (format: plugin_<id>) when creating or focusing plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_launchOrFocusPluginCmd).Standalone()

	action_helpCmd.AddCommand(action_help_launchOrFocusPluginCmd)
}
