package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_pageScrollDownCmd = &cobra.Command{
	Use:   "page-scroll-down",
	Short: "Scroll down one page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_pageScrollDownCmd).Standalone()

	help_actionCmd.AddCommand(help_action_pageScrollDownCmd)
}
