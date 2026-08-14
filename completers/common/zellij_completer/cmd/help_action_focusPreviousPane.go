package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_focusPreviousPaneCmd = &cobra.Command{
	Use:   "focus-previous-pane",
	Short: "Change focus to the previous pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_focusPreviousPaneCmd).Standalone()

	help_actionCmd.AddCommand(help_action_focusPreviousPaneCmd)
}
