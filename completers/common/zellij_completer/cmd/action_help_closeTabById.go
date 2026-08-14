package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_closeTabByIdCmd = &cobra.Command{
	Use:   "close-tab-by-id",
	Short: "Close tab with stable ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_closeTabByIdCmd).Standalone()

	action_helpCmd.AddCommand(action_help_closeTabByIdCmd)
}
