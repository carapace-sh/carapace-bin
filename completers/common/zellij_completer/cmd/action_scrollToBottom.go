package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_scrollToBottomCmd = &cobra.Command{
	Use:   "scroll-to-bottom",
	Short: "Scroll down to bottom in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_scrollToBottomCmd).Standalone()

	action_scrollToBottomCmd.Flags().BoolP("help", "h", false, "Print help")
	action_scrollToBottomCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_scrollToBottomCmd)

	carapace.Gen(action_scrollToBottomCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
