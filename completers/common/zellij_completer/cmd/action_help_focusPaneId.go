package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_focusPaneIdCmd = &cobra.Command{
	Use:   "focus-pane-id",
	Short: "Focus a specific pane by its ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_focusPaneIdCmd).Standalone()

	action_helpCmd.AddCommand(action_help_focusPaneIdCmd)
}
