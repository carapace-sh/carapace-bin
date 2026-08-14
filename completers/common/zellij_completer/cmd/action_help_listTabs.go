package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_listTabsCmd = &cobra.Command{
	Use:   "list-tabs",
	Short: "List all tabs with their information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_listTabsCmd).Standalone()

	action_helpCmd.AddCommand(action_help_listTabsCmd)
}
