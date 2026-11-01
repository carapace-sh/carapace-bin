package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_showFloatingPanesCmd = &cobra.Command{
	Use:   "show-floating-panes",
	Short: "Show all floating panes in the specified tab (or active tab if tab_id is not provided)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_showFloatingPanesCmd).Standalone()

	help_actionCmd.AddCommand(help_action_showFloatingPanesCmd)
}
