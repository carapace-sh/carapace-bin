package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_goToTabByIdCmd = &cobra.Command{
	Use:   "go-to-tab-by-id",
	Short: "Go to tab with stable ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_goToTabByIdCmd).Standalone()

	action_helpCmd.AddCommand(action_help_goToTabByIdCmd)
}
