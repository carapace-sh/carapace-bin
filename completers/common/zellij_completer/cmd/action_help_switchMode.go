package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_switchModeCmd = &cobra.Command{
	Use:   "switch-mode",
	Short: "Switch input mode of all connected clients [locked|pane|tab|resize|move|search|session]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_switchModeCmd).Standalone()

	action_helpCmd.AddCommand(action_help_switchModeCmd)
}
