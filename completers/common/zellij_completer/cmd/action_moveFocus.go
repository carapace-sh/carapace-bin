package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zellij_completer/cmd/action"
	"github.com/spf13/cobra"
)

var action_moveFocusCmd = &cobra.Command{
	Use:   "move-focus",
	Short: "Move the focused pane in the specified direction. [right|left|up|down]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_moveFocusCmd).Standalone()

	action_moveFocusCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_moveFocusCmd)

	carapace.Gen(action_moveFocusCmd).PositionalAnyCompletion(action.ActionDirections())
}
