package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_toggleActiveSyncTabCmd = &cobra.Command{
	Use:   "toggle-active-sync-tab",
	Short: "Toggle between sending text commands to all panes on the current tab and normal mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_toggleActiveSyncTabCmd).Standalone()

	action_toggleActiveSyncTabCmd.Flags().BoolP("help", "h", false, "Print help")
	action_toggleActiveSyncTabCmd.Flags().StringP("tab-id", "t", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_toggleActiveSyncTabCmd)

	carapace.Gen(action_toggleActiveSyncTabCmd).FlagCompletion(carapace.ActionMap{
		"tab-id": zellij.ActionTabs(),
	})
}
