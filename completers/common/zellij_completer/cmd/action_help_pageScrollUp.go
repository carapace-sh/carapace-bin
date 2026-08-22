package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_pageScrollUpCmd = &cobra.Command{
	Use:   "page-scroll-up",
	Short: "Scroll up one page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_pageScrollUpCmd).Standalone()

	action_helpCmd.AddCommand(action_help_pageScrollUpCmd)
}
