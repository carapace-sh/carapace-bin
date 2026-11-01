package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "[increase|decrease] the focused panes area at the [left|down|up|right] border",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_resizeCmd).Standalone()

	action_helpCmd.AddCommand(action_help_resizeCmd)
}
