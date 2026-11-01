package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_toggleActiveSyncTabCmd = &cobra.Command{
	Use:   "toggle-active-sync-tab",
	Short: "Toggle between sending text commands to all panes on the current tab and normal mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_toggleActiveSyncTabCmd).Standalone()

	action_helpCmd.AddCommand(action_help_toggleActiveSyncTabCmd)
}
