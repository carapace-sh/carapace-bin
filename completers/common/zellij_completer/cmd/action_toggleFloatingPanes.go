package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_toggleFloatingPanesCmd = &cobra.Command{
	Use:   "toggle-floating-panes",
	Short: "Toggle the visibility of all floating panes in the current Tab, open one if none exist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_toggleFloatingPanesCmd).Standalone()

	action_toggleFloatingPanesCmd.Flags().BoolP("help", "h", false, "Print help")
	action_toggleFloatingPanesCmd.Flags().StringP("tab-id", "t", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_toggleFloatingPanesCmd)

	carapace.Gen(action_toggleFloatingPanesCmd).FlagCompletion(carapace.ActionMap{
		"tab-id": zellij.ActionTabs(),
	})
}
