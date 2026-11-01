package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_halfPageScrollDownCmd = &cobra.Command{
	Use:   "half-page-scroll-down",
	Short: "Scroll down half page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_halfPageScrollDownCmd).Standalone()

	help_actionCmd.AddCommand(help_action_halfPageScrollDownCmd)
}
