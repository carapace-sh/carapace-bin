package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_moveTabCmd = &cobra.Command{
	Use:   "move-tab",
	Short: "Move the focused tab in the specified direction. [right|left]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_moveTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_moveTabCmd)
}
