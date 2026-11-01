package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write bytes to the terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_writeCmd).Standalone()

	action_writeCmd.Flags().BoolP("help", "h", false, "Print help")
	action_writeCmd.Flags().StringP("pane-id", "p", "", "The pane_id of the pane, eg. terminal_1, plugin_2 or 3 (equivalent to terminal_3)")
	actionCmd.AddCommand(action_writeCmd)

	carapace.Gen(action_writeCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
