package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_renameTabCmd = &cobra.Command{
	Use:   "rename-tab",
	Short: "Renames the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_renameTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_renameTabCmd)
}
