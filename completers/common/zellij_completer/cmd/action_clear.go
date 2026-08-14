package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all buffers for a focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_clearCmd).Standalone()

	action_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	action_clearCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_clearCmd)
}
