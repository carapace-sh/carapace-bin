package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_goToTabByIdCmd = &cobra.Command{
	Use:   "go-to-tab-by-id",
	Short: "Go to tab with stable ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_goToTabByIdCmd).Standalone()

	help_actionCmd.AddCommand(help_action_goToTabByIdCmd)
}
