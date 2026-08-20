package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zellij_completer/cmd/action"
	"github.com/spf13/cobra"
)

var action_movePaneCmd = &cobra.Command{
	Use:   "move-pane",
	Short: "Change the location of the focused pane in the specified direction or rotate forwrads [right|left|up|down]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_movePaneCmd).Standalone()

	action_movePaneCmd.Flags().BoolP("help", "h", false, "Print help")
	action_movePaneCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_movePaneCmd)

	carapace.Gen(action_movePaneCmd).PositionalAnyCompletion(action.ActionDirections())
}
