package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_focusNextPaneCmd = &cobra.Command{
	Use:   "focus-next-pane",
	Short: "Change focus to the next pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_focusNextPaneCmd).Standalone()

	action_focusNextPaneCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_focusNextPaneCmd)
}
