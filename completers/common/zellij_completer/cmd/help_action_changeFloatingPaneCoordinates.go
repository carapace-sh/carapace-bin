package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_changeFloatingPaneCoordinatesCmd = &cobra.Command{
	Use:   "change-floating-pane-coordinates",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_changeFloatingPaneCoordinatesCmd).Standalone()

	help_actionCmd.AddCommand(help_action_changeFloatingPaneCoordinatesCmd)
}
