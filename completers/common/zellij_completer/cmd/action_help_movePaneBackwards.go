package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_movePaneBackwardsCmd = &cobra.Command{
	Use:   "move-pane-backwards",
	Short: "Rotate the location of the previous pane backwards",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_movePaneBackwardsCmd).Standalone()

	action_helpCmd.AddCommand(action_help_movePaneBackwardsCmd)
}
