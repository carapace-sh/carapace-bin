package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_closePaneCmd = &cobra.Command{
	Use:   "close-pane",
	Short: "Close the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_closePaneCmd).Standalone()

	action_closePaneCmd.Flags().BoolP("help", "h", false, "Print help")
	action_closePaneCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_closePaneCmd)

	carapace.Gen(action_closePaneCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
