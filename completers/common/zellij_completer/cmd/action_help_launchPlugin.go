package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_launchPluginCmd = &cobra.Command{
	Use:   "launch-plugin",
	Short: "Returns: Plugin pane ID (format: plugin_<id>)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_launchPluginCmd).Standalone()

	action_helpCmd.AddCommand(action_help_launchPluginCmd)
}
