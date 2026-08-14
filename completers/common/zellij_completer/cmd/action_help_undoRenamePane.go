package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_undoRenamePaneCmd = &cobra.Command{
	Use:   "undo-rename-pane",
	Short: "Remove a previously set pane name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_undoRenamePaneCmd).Standalone()

	action_helpCmd.AddCommand(action_help_undoRenamePaneCmd)
}
