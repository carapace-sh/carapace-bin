package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_closePaneCmd = &cobra.Command{
	Use:   "close-pane",
	Short: "Close the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_closePaneCmd).Standalone()

	help_actionCmd.AddCommand(help_action_closePaneCmd)
}
