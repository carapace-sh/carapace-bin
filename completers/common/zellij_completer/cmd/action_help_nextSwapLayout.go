package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_nextSwapLayoutCmd = &cobra.Command{
	Use:   "next-swap-layout",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_nextSwapLayoutCmd).Standalone()

	action_helpCmd.AddCommand(action_help_nextSwapLayoutCmd)
}
