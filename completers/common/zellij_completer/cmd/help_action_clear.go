package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all buffers for a focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_clearCmd).Standalone()

	help_actionCmd.AddCommand(help_action_clearCmd)
}
