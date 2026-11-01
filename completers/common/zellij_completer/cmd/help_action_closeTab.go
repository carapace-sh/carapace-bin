package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_closeTabCmd = &cobra.Command{
	Use:   "close-tab",
	Short: "Close the current tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_closeTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_closeTabCmd)
}
