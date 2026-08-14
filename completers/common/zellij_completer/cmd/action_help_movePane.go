package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_movePaneCmd = &cobra.Command{
	Use:   "move-pane",
	Short: "Change the location of the focused pane in the specified direction or rotate forwrads [right|left|up|down]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_movePaneCmd).Standalone()

	action_helpCmd.AddCommand(action_help_movePaneCmd)
}
