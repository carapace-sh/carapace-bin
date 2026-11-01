package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_toggleFullscreenCmd = &cobra.Command{
	Use:   "toggle-fullscreen",
	Short: "Toggle between fullscreen focus pane and normal layout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_toggleFullscreenCmd).Standalone()

	help_actionCmd.AddCommand(help_action_toggleFullscreenCmd)
}
