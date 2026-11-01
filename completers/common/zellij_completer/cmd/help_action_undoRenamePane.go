package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_undoRenamePaneCmd = &cobra.Command{
	Use:   "undo-rename-pane",
	Short: "Remove a previously set pane name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_undoRenamePaneCmd).Standalone()

	help_actionCmd.AddCommand(help_action_undoRenamePaneCmd)
}
