package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_nextSwapLayoutCmd = &cobra.Command{
	Use:   "next-swap-layout",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_nextSwapLayoutCmd).Standalone()

	action_nextSwapLayoutCmd.Flags().BoolP("help", "h", false, "Print help")
	action_nextSwapLayoutCmd.Flags().StringP("tab-id", "t", "", "Target a specific tab by ID")
	actionCmd.AddCommand(action_nextSwapLayoutCmd)

	carapace.Gen(action_nextSwapLayoutCmd).FlagCompletion(carapace.ActionMap{
		"tab-id": zellij.ActionTabs(),
	})
}
