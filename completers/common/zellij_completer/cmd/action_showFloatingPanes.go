package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_showFloatingPanesCmd = &cobra.Command{
	Use:   "show-floating-panes",
	Short: "Show all floating panes in the specified tab (or active tab if tab_id is not provided)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_showFloatingPanesCmd).Standalone()

	action_showFloatingPanesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	action_showFloatingPanesCmd.Flags().StringP("tab-id", "t", "", "")
	actionCmd.AddCommand(action_showFloatingPanesCmd)

	carapace.Gen(action_showFloatingPanesCmd).FlagCompletion(carapace.ActionMap{
		"tab-id": zellij.ActionTabs(),
	})
}
