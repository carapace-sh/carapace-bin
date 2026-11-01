package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_moveTabCmd = &cobra.Command{
	Use:   "move-tab",
	Short: "Move the focused tab in the specified direction. [right|left]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_moveTabCmd).Standalone()

	action_helpCmd.AddCommand(action_help_moveTabCmd)
}
