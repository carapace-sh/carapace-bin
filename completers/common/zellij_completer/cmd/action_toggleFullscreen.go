package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_toggleFullscreenCmd = &cobra.Command{
	Use:   "toggle-fullscreen",
	Short: "Toggle between fullscreen focus pane and normal layout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_toggleFullscreenCmd).Standalone()

	action_toggleFullscreenCmd.Flags().BoolP("help", "h", false, "Print help")
	action_toggleFullscreenCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_toggleFullscreenCmd)
}
