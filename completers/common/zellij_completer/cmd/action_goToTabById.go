package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_goToTabByIdCmd = &cobra.Command{
	Use:   "go-to-tab-by-id",
	Short: "Go to tab with stable ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_goToTabByIdCmd).Standalone()

	action_goToTabByIdCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_goToTabByIdCmd)

	carapace.Gen(action_goToTabByIdCmd).PositionalAnyCompletion(zellij.ActionTabs())
}
