package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_toggleFloatingPanesCmd = &cobra.Command{
	Use:   "toggle-floating-panes",
	Short: "Toggle the visibility of all floating panes in the current Tab, open one if none exist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_toggleFloatingPanesCmd).Standalone()

	action_helpCmd.AddCommand(action_help_toggleFloatingPanesCmd)
}
