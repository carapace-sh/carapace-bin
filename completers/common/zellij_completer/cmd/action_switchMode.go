package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zellij_completer/cmd/action"
	"github.com/spf13/cobra"
)

var action_switchModeCmd = &cobra.Command{
	Use:   "switch-mode",
	Short: "Switch input mode of all connected clients [locked|pane|tab|resize|move|search|session]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_switchModeCmd).Standalone()

	action_switchModeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	actionCmd.AddCommand(action_switchModeCmd)

	carapace.Gen(action_switchModeCmd).PositionalAnyCompletion(action.ActionModes())
}
