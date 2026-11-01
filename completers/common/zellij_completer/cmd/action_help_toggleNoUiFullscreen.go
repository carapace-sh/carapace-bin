package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_toggleNoUiFullscreenCmd = &cobra.Command{
	Use:   "toggle-no-ui-fullscreen",
	Short: "Toggle between fullscreen over the entire display (including the UI bars) and normal layout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_toggleNoUiFullscreenCmd).Standalone()

	action_helpCmd.AddCommand(action_help_toggleNoUiFullscreenCmd)
}
