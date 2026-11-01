package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_halfPageScrollUpCmd = &cobra.Command{
	Use:   "half-page-scroll-up",
	Short: "Scroll up half page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_halfPageScrollUpCmd).Standalone()

	action_helpCmd.AddCommand(action_help_halfPageScrollUpCmd)
}
