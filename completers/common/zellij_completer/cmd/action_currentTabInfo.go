package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_currentTabInfoCmd = &cobra.Command{
	Use:   "current-tab-info",
	Short: "Get information about the currently active tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_currentTabInfoCmd).Standalone()

	action_currentTabInfoCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	action_currentTabInfoCmd.Flags().BoolP("json", "j", false, "Output as JSON with full TabInfo")
	actionCmd.AddCommand(action_currentTabInfoCmd)
}
