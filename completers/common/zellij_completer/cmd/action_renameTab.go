package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_renameTabCmd = &cobra.Command{
	Use:   "rename-tab",
	Short: "Renames the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_renameTabCmd).Standalone()

	action_renameTabCmd.Flags().BoolP("help", "h", false, "Print help")
	action_renameTabCmd.Flags().StringP("tab-id", "t", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_renameTabCmd)

	carapace.Gen(action_renameTabCmd).FlagCompletion(carapace.ActionMap{
		"tab-id": zellij.ActionTabs(),
	})
}
