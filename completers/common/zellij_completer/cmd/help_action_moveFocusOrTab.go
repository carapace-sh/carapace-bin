package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_moveFocusOrTabCmd = &cobra.Command{
	Use:   "move-focus-or-tab",
	Short: "Move focus to the pane or tab (if on screen edge) in the specified direction [right|left|up|down]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_moveFocusOrTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_moveFocusOrTabCmd)
}
