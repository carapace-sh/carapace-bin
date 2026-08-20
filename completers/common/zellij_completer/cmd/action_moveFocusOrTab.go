package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zellij_completer/cmd/action"
	"github.com/spf13/cobra"
)

var action_moveFocusOrTabCmd = &cobra.Command{
	Use:   "move-focus-or-tab",
	Short: "Move focus to the pane or tab (if on screen edge) in the specified direction [right|left|up|down]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_moveFocusOrTabCmd).Standalone()

	action_moveFocusOrTabCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_moveFocusOrTabCmd)

	carapace.Gen(action_moveFocusOrTabCmd).PositionalAnyCompletion(action.ActionDirections())
}
