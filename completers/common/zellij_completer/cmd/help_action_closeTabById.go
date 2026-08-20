package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_closeTabByIdCmd = &cobra.Command{
	Use:   "close-tab-by-id",
	Short: "Close tab with stable ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_closeTabByIdCmd).Standalone()

	help_actionCmd.AddCommand(help_action_closeTabByIdCmd)
}
