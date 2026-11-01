package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "[increase|decrease] the focused panes area at the [left|down|up|right] border",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_resizeCmd).Standalone()

	help_actionCmd.AddCommand(help_action_resizeCmd)
}
