package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_undoRenameTabCmd = &cobra.Command{
	Use:   "undo-rename-tab",
	Short: "Remove a previously set tab name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_undoRenameTabCmd).Standalone()

	action_helpCmd.AddCommand(action_help_undoRenameTabCmd)
}
