package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_pasteCmd = &cobra.Command{
	Use:   "paste",
	Short: "Paste text to the terminal (using bracketed paste mode)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_pasteCmd).Standalone()

	action_pasteCmd.Flags().BoolP("help", "h", false, "Print help")
	action_pasteCmd.Flags().StringP("pane-id", "p", "", "The pane_id of the pane, eg. terminal_1, plugin_2 or 3 (equivalent to terminal_3)")
	actionCmd.AddCommand(action_pasteCmd)

	carapace.Gen(action_pasteCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
	})
}
