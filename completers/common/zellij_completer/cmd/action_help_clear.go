package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all buffers for a focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_clearCmd).Standalone()

	action_helpCmd.AddCommand(action_help_clearCmd)
}
