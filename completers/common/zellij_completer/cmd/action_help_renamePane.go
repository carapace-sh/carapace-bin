package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_renamePaneCmd = &cobra.Command{
	Use:   "rename-pane",
	Short: "Renames the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_renamePaneCmd).Standalone()

	action_helpCmd.AddCommand(action_help_renamePaneCmd)
}
