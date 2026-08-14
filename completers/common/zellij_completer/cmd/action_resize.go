package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "[increase|decrease] the focused panes area at the [left|down|up|right] border",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_resizeCmd).Standalone()

	action_resizeCmd.Flags().BoolP("help", "h", false, "Print help")
	action_resizeCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_resizeCmd)

	carapace.Gen(action_resizeCmd).PositionalCompletion(
		actionResizes(),
		actionDirections(),
	)
}
