package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_listPanesCmd = &cobra.Command{
	Use:   "list-panes",
	Short: "List all panes in the current session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_listPanesCmd).Standalone()

	action_helpCmd.AddCommand(action_help_listPanesCmd)
}
