package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_focusNextPaneCmd = &cobra.Command{
	Use:   "focus-next-pane",
	Short: "Change focus to the next pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_focusNextPaneCmd).Standalone()

	action_helpCmd.AddCommand(action_help_focusNextPaneCmd)
}
