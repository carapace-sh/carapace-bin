package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_scrollUpCmd = &cobra.Command{
	Use:   "scroll-up",
	Short: "Scroll up in the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_scrollUpCmd).Standalone()

	action_helpCmd.AddCommand(action_help_scrollUpCmd)
}
