package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_halfPageScrollDownCmd = &cobra.Command{
	Use:   "half-page-scroll-down",
	Short: "Scroll down half page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_halfPageScrollDownCmd).Standalone()

	action_helpCmd.AddCommand(action_help_halfPageScrollDownCmd)
}
