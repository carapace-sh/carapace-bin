package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_hideFloatingPanesCmd = &cobra.Command{
	Use:   "hide-floating-panes",
	Short: "Hide all floating panes in the specified tab (or active tab if tab_id is not provided)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_hideFloatingPanesCmd).Standalone()

	help_actionCmd.AddCommand(help_action_hideFloatingPanesCmd)
}
