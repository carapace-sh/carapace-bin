package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_launchPluginCmd = &cobra.Command{
	Use:   "launch-plugin",
	Short: "Returns: Plugin pane ID (format: plugin_<id>)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_launchPluginCmd).Standalone()

	action_launchPluginCmd.Flags().Bool("close-replaced-pane", false, "Close the replaced pane instead of suspending it (only effective with --in-place)")
	action_launchPluginCmd.Flags().StringP("configuration", "c", "", "")
	action_launchPluginCmd.Flags().BoolP("floating", "f", false, "")
	action_launchPluginCmd.Flags().BoolP("help", "h", false, "Print help")
	action_launchPluginCmd.Flags().BoolP("in-place", "i", false, "")
	action_launchPluginCmd.Flags().Bool("no-focus", false, "if set, will open the plugin pane without changing the focus of any client")
	action_launchPluginCmd.Flags().BoolP("skip-plugin-cache", "s", false, "")
	action_launchPluginCmd.Flags().String("tab-id", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_launchPluginCmd)
}
