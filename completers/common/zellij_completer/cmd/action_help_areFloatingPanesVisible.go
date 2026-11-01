package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_areFloatingPanesVisibleCmd = &cobra.Command{
	Use:   "are-floating-panes-visible",
	Short: "Check if floating panes are visible in the specified tab (or active tab)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_areFloatingPanesVisibleCmd).Standalone()

	action_helpCmd.AddCommand(action_help_areFloatingPanesVisibleCmd)
}
