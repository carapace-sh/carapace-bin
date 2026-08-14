package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_halfPageScrollUpCmd = &cobra.Command{
	Use:   "half-page-scroll-up",
	Short: "Scroll up half page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_halfPageScrollUpCmd).Standalone()

	action_halfPageScrollUpCmd.Flags().BoolP("help", "h", false, "Print help")
	action_halfPageScrollUpCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_halfPageScrollUpCmd)
}
