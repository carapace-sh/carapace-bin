package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_focusPreviousPaneCmd = &cobra.Command{
	Use:   "focus-previous-pane",
	Short: "Change focus to the previous pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_focusPreviousPaneCmd).Standalone()

	action_helpCmd.AddCommand(action_help_focusPreviousPaneCmd)
}
