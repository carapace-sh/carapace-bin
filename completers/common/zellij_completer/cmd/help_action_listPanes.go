package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_listPanesCmd = &cobra.Command{
	Use:   "list-panes",
	Short: "List all panes in the current session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_listPanesCmd).Standalone()

	help_actionCmd.AddCommand(help_action_listPanesCmd)
}
