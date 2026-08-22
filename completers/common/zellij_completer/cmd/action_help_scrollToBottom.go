package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_scrollToBottomCmd = &cobra.Command{
	Use:   "scroll-to-bottom",
	Short: "Scroll down to bottom in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_scrollToBottomCmd).Standalone()

	action_helpCmd.AddCommand(action_help_scrollToBottomCmd)
}
