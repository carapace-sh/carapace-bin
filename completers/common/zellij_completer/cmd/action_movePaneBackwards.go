package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_movePaneBackwardsCmd = &cobra.Command{
	Use:   "move-pane-backwards",
	Short: "Rotate the location of the previous pane backwards",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_movePaneBackwardsCmd).Standalone()

	action_movePaneBackwardsCmd.Flags().BoolP("help", "h", false, "Print help")
	action_movePaneBackwardsCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_movePaneBackwardsCmd)
}
