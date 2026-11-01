package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_scrollDownCmd = &cobra.Command{
	Use:   "scroll-down",
	Short: "Scroll down in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_scrollDownCmd).Standalone()

	help_actionCmd.AddCommand(help_action_scrollDownCmd)
}
