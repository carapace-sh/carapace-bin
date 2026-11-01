package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_togglePaneFramesCmd = &cobra.Command{
	Use:   "toggle-pane-frames",
	Short: "Toggle frames around panes in the UI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_togglePaneFramesCmd).Standalone()

	help_actionCmd.AddCommand(help_action_togglePaneFramesCmd)
}
