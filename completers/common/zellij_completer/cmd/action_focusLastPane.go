package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_focusLastPaneCmd = &cobra.Command{
	Use:   "focus-last-pane",
	Short: "Change focus to the last focused frame",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_focusLastPaneCmd).Standalone()

	action_focusLastPaneCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_focusLastPaneCmd)
}
