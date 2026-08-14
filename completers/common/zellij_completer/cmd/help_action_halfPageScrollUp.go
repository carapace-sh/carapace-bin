package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_halfPageScrollUpCmd = &cobra.Command{
	Use:   "half-page-scroll-up",
	Short: "Scroll up half page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_halfPageScrollUpCmd).Standalone()

	help_actionCmd.AddCommand(help_action_halfPageScrollUpCmd)
}
