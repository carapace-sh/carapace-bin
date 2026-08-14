package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_scrollToBottomCmd = &cobra.Command{
	Use:   "scroll-to-bottom",
	Short: "Scroll down to bottom in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_scrollToBottomCmd).Standalone()

	help_actionCmd.AddCommand(help_action_scrollToBottomCmd)
}
