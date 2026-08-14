package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_previousSwapLayoutCmd = &cobra.Command{
	Use:   "previous-swap-layout",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_previousSwapLayoutCmd).Standalone()

	action_previousSwapLayoutCmd.Flags().BoolP("help", "h", false, "Print help")
	action_previousSwapLayoutCmd.Flags().StringP("tab-id", "t", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_previousSwapLayoutCmd)
}
