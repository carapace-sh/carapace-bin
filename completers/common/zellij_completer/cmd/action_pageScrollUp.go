package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_pageScrollUpCmd = &cobra.Command{
	Use:   "page-scroll-up",
	Short: "Scroll up one page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_pageScrollUpCmd).Standalone()

	action_pageScrollUpCmd.Flags().BoolP("help", "h", false, "Print help")
	action_pageScrollUpCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_pageScrollUpCmd)

	carapace.Gen(action_pageScrollUpCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
