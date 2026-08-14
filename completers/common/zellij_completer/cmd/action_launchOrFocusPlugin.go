package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_launchOrFocusPluginCmd = &cobra.Command{
	Use:   "launch-or-focus-plugin",
	Short: "Returns: Plugin pane ID (format: plugin_<id>) when creating or focusing plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_launchOrFocusPluginCmd).Standalone()

	action_launchOrFocusPluginCmd.Flags().Bool("close-replaced-pane", false, "Close the replaced pane instead of suspending it (only effective with --in-place)")
	action_launchOrFocusPluginCmd.Flags().StringP("configuration", "c", "", "")
	action_launchOrFocusPluginCmd.Flags().BoolP("floating", "f", false, "")
	action_launchOrFocusPluginCmd.Flags().BoolP("help", "h", false, "Print help")
	action_launchOrFocusPluginCmd.Flags().BoolP("in-place", "i", false, "")
	action_launchOrFocusPluginCmd.Flags().BoolP("move-to-focused-tab", "m", false, "")
	action_launchOrFocusPluginCmd.Flags().BoolP("skip-plugin-cache", "s", false, "")
	action_launchOrFocusPluginCmd.Flags().String("tab-id", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_launchOrFocusPluginCmd)
}
