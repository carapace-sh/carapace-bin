package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_togglePanePinnedCmd = &cobra.Command{
	Use:   "toggle-pane-pinned",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_togglePanePinnedCmd).Standalone()

	action_togglePanePinnedCmd.Flags().BoolP("help", "h", false, "Print help")
	action_togglePanePinnedCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_togglePanePinnedCmd)

	carapace.Gen(action_togglePanePinnedCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
