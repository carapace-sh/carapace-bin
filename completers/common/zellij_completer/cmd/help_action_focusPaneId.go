package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_focusPaneIdCmd = &cobra.Command{
	Use:   "focus-pane-id",
	Short: "Focus a specific pane by its ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_focusPaneIdCmd).Standalone()

	help_actionCmd.AddCommand(help_action_focusPaneIdCmd)
}
