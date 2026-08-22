package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_previousSwapLayoutCmd = &cobra.Command{
	Use:   "previous-swap-layout",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_previousSwapLayoutCmd).Standalone()

	help_actionCmd.AddCommand(help_action_previousSwapLayoutCmd)
}
