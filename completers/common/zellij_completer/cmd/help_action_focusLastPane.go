package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_focusLastPaneCmd = &cobra.Command{
	Use:   "focus-last-pane",
	Short: "Change focus to the last focused frame",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_focusLastPaneCmd).Standalone()

	help_actionCmd.AddCommand(help_action_focusLastPaneCmd)
}
