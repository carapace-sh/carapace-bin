package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_closeTabCmd = &cobra.Command{
	Use:   "close-tab",
	Short: "Close the current tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_closeTabCmd).Standalone()

	action_closeTabCmd.Flags().BoolP("help", "h", false, "Print help")
	action_closeTabCmd.Flags().StringP("tab-id", "t", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_closeTabCmd)
}
