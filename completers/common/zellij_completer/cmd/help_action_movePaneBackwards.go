package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_movePaneBackwardsCmd = &cobra.Command{
	Use:   "move-pane-backwards",
	Short: "Rotate the location of the previous pane backwards",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_movePaneBackwardsCmd).Standalone()

	help_actionCmd.AddCommand(help_action_movePaneBackwardsCmd)
}
