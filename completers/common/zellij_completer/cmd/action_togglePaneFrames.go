package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_togglePaneFramesCmd = &cobra.Command{
	Use:   "toggle-pane-frames",
	Short: "Toggle frames around panes in the UI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_togglePaneFramesCmd).Standalone()

	action_togglePaneFramesCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_togglePaneFramesCmd)
}
