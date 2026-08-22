package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_switchModeCmd = &cobra.Command{
	Use:   "switch-mode",
	Short: "Switch input mode of all connected clients [locked|pane|tab|resize|move|search|session]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_switchModeCmd).Standalone()

	help_actionCmd.AddCommand(help_action_switchModeCmd)
}
