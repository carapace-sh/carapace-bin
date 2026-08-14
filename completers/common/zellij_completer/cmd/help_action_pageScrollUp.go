package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_pageScrollUpCmd = &cobra.Command{
	Use:   "page-scroll-up",
	Short: "Scroll up one page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_pageScrollUpCmd).Standalone()

	help_actionCmd.AddCommand(help_action_pageScrollUpCmd)
}
