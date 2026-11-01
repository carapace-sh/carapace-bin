package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_hideFloatingPanesCmd = &cobra.Command{
	Use:   "hide-floating-panes",
	Short: "Hide all floating panes in the specified tab (or active tab if tab_id is not provided)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_hideFloatingPanesCmd).Standalone()

	action_helpCmd.AddCommand(action_help_hideFloatingPanesCmd)
}
