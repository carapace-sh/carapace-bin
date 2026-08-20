package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_toggleNoUiFullscreenCmd = &cobra.Command{
	Use:   "toggle-no-ui-fullscreen",
	Short: "Toggle between fullscreen over the entire display (including the UI bars) and normal layout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_toggleNoUiFullscreenCmd).Standalone()

	action_toggleNoUiFullscreenCmd.Flags().BoolP("help", "h", false, "Print help")
	action_toggleNoUiFullscreenCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_toggleNoUiFullscreenCmd)
}
