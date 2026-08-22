package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_scrollToTopCmd = &cobra.Command{
	Use:   "scroll-to-top",
	Short: "Scroll up to top in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_scrollToTopCmd).Standalone()

	action_helpCmd.AddCommand(action_help_scrollToTopCmd)
}
