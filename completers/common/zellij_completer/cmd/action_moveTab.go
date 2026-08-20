package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zellij_completer/cmd/action"
	"github.com/spf13/cobra"
)

var action_moveTabCmd = &cobra.Command{
	Use:   "move-tab",
	Short: "Move the focused tab in the specified direction. [right|left]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_moveTabCmd).Standalone()

	action_moveTabCmd.Flags().BoolP("help", "h", false, "Print help")
	action_moveTabCmd.Flags().StringP("tab-id", "t", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_moveTabCmd)

	carapace.Gen(action_moveTabCmd).PositionalAnyCompletion(action.ActionDirections())
}
