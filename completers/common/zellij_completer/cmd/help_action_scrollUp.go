package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_scrollUpCmd = &cobra.Command{
	Use:   "scroll-up",
	Short: "Scroll up in the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_scrollUpCmd).Standalone()

	help_actionCmd.AddCommand(help_action_scrollUpCmd)
}
