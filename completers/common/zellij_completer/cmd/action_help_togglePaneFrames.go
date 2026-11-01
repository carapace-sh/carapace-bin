package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_togglePaneFramesCmd = &cobra.Command{
	Use:   "toggle-pane-frames",
	Short: "Toggle frames around panes in the UI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_togglePaneFramesCmd).Standalone()

	action_helpCmd.AddCommand(action_help_togglePaneFramesCmd)
}
