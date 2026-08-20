package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_scrollDownCmd = &cobra.Command{
	Use:   "scroll-down",
	Short: "Scroll down in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_scrollDownCmd).Standalone()

	action_scrollDownCmd.Flags().BoolP("help", "h", false, "Print help")
	action_scrollDownCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_scrollDownCmd)
}
