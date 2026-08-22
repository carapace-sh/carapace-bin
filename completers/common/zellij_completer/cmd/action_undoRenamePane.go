package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_undoRenamePaneCmd = &cobra.Command{
	Use:   "undo-rename-pane",
	Short: "Remove a previously set pane name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_undoRenamePaneCmd).Standalone()

	action_undoRenamePaneCmd.Flags().BoolP("help", "h", false, "Print help")
	action_undoRenamePaneCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_undoRenamePaneCmd)

	carapace.Gen(action_undoRenamePaneCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
