package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_undoRenameTabCmd = &cobra.Command{
	Use:   "undo-rename-tab",
	Short: "Remove a previously set tab name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_undoRenameTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_undoRenameTabCmd)
}
