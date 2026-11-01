package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_currentTabInfoCmd = &cobra.Command{
	Use:   "current-tab-info",
	Short: "Get information about the currently active tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_currentTabInfoCmd).Standalone()

	action_helpCmd.AddCommand(action_help_currentTabInfoCmd)
}
