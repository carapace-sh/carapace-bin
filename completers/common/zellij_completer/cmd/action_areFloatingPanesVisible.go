package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_areFloatingPanesVisibleCmd = &cobra.Command{
	Use:   "are-floating-panes-visible",
	Short: "Check if floating panes are visible in the specified tab (or active tab)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_areFloatingPanesVisibleCmd).Standalone()

	action_areFloatingPanesVisibleCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	action_areFloatingPanesVisibleCmd.Flags().StringP("tab-id", "t", "", "")
	actionCmd.AddCommand(action_areFloatingPanesVisibleCmd)

	carapace.Gen(action_areFloatingPanesVisibleCmd).FlagCompletion(carapace.ActionMap{
		"tab-id": zellij.ActionTabs(),
	})
}
