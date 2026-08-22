package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_pageScrollDownCmd = &cobra.Command{
	Use:   "page-scroll-down",
	Short: "Scroll down one page in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_pageScrollDownCmd).Standalone()

	action_pageScrollDownCmd.Flags().BoolP("help", "h", false, "Print help")
	action_pageScrollDownCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_pageScrollDownCmd)

	carapace.Gen(action_pageScrollDownCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
