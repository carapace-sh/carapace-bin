package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_toggleFloatingPanesCmd = &cobra.Command{
	Use:   "toggle-floating-panes",
	Short: "Toggle the visibility of all floating panes in the current Tab, open one if none exist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_toggleFloatingPanesCmd).Standalone()

	help_actionCmd.AddCommand(help_action_toggleFloatingPanesCmd)
}
