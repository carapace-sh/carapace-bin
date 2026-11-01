package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_moveFocusCmd = &cobra.Command{
	Use:   "move-focus",
	Short: "Move the focused pane in the specified direction. [right|left|up|down]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_moveFocusCmd).Standalone()

	help_actionCmd.AddCommand(help_action_moveFocusCmd)
}
