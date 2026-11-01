package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_renameTabCmd = &cobra.Command{
	Use:   "rename-tab",
	Short: "Renames the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_renameTabCmd).Standalone()

	action_helpCmd.AddCommand(action_help_renameTabCmd)
}
