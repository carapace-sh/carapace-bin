package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_scrollDownCmd = &cobra.Command{
	Use:   "scroll-down",
	Short: "Scroll down in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_scrollDownCmd).Standalone()

	action_helpCmd.AddCommand(action_help_scrollDownCmd)
}
