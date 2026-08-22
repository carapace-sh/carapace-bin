package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_hideFloatingPanesCmd = &cobra.Command{
	Use:   "hide-floating-panes",
	Short: "Hide all floating panes in the specified tab (or active tab if tab_id is not provided)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_hideFloatingPanesCmd).Standalone()

	action_hideFloatingPanesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	action_hideFloatingPanesCmd.Flags().StringP("tab-id", "t", "", "")
	actionCmd.AddCommand(action_hideFloatingPanesCmd)

	carapace.Gen(action_hideFloatingPanesCmd).FlagCompletion(carapace.ActionMap{
		"tab-id": zellij.ActionTabs(),
	})
}
