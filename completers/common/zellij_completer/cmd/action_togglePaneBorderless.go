package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_togglePaneBorderlessCmd = &cobra.Command{
	Use:   "toggle-pane-borderless",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_togglePaneBorderlessCmd).Standalone()

	action_togglePaneBorderlessCmd.Flags().BoolP("help", "h", false, "Print help")
	action_togglePaneBorderlessCmd.Flags().StringP("pane-id", "p", "", "The pane_id of the pane, eg. terminal_1, plugin_2 or 3 (equivalent to terminal_3)")
	action_togglePaneBorderlessCmd.MarkFlagRequired("pane-id")
	actionCmd.AddCommand(action_togglePaneBorderlessCmd)

	carapace.Gen(action_togglePaneBorderlessCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
