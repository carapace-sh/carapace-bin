package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_scrollToTopCmd = &cobra.Command{
	Use:   "scroll-to-top",
	Short: "Scroll up to top in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_scrollToTopCmd).Standalone()

	help_actionCmd.AddCommand(help_action_scrollToTopCmd)
}
