package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_areFloatingPanesVisibleCmd = &cobra.Command{
	Use:   "are-floating-panes-visible",
	Short: "Check if floating panes are visible in the specified tab (or active tab)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_areFloatingPanesVisibleCmd).Standalone()

	help_actionCmd.AddCommand(help_action_areFloatingPanesVisibleCmd)
}
