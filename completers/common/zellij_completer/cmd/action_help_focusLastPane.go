package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_focusLastPaneCmd = &cobra.Command{
	Use:   "focus-last-pane",
	Short: "Change focus to the last focused frame",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_focusLastPaneCmd).Standalone()

	action_helpCmd.AddCommand(action_help_focusLastPaneCmd)
}
