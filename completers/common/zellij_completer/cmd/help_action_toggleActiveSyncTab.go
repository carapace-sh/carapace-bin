package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_toggleActiveSyncTabCmd = &cobra.Command{
	Use:   "toggle-active-sync-tab",
	Short: "Toggle between sending text commands to all panes on the current tab and normal mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_toggleActiveSyncTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_toggleActiveSyncTabCmd)
}
