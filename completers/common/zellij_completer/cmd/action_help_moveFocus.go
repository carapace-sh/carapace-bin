package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_moveFocusCmd = &cobra.Command{
	Use:   "move-focus",
	Short: "Move the focused pane in the specified direction. [right|left|up|down]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_moveFocusCmd).Standalone()

	action_helpCmd.AddCommand(action_help_moveFocusCmd)
}
