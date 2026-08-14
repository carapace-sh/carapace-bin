package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_listTabsCmd = &cobra.Command{
	Use:   "list-tabs",
	Short: "List all tabs with their information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_listTabsCmd).Standalone()

	help_actionCmd.AddCommand(help_action_listTabsCmd)
}
