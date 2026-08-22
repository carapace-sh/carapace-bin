package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_scrollUpCmd = &cobra.Command{
	Use:   "scroll-up",
	Short: "Scroll up in the focused pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_scrollUpCmd).Standalone()

	action_scrollUpCmd.Flags().BoolP("help", "h", false, "Print help")
	action_scrollUpCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_scrollUpCmd)

	carapace.Gen(action_scrollUpCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
