package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_undoRenameTabCmd = &cobra.Command{
	Use:   "undo-rename-tab",
	Short: "Remove a previously set tab name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_undoRenameTabCmd).Standalone()

	action_undoRenameTabCmd.Flags().BoolP("help", "h", false, "Print help")
	action_undoRenameTabCmd.Flags().StringP("tab-id", "t", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_undoRenameTabCmd)

	carapace.Gen(action_undoRenameTabCmd).FlagCompletion(carapace.ActionMap{
		"tab-id": zellij.ActionTabs(),
	})
}
