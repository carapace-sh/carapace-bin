package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_currentTabInfoCmd = &cobra.Command{
	Use:   "current-tab-info",
	Short: "Get information about the currently active tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_currentTabInfoCmd).Standalone()

	help_actionCmd.AddCommand(help_action_currentTabInfoCmd)
}
