package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all items in history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_listCmd).Standalone()

	help_historyCmd.AddCommand(help_history_listCmd)
}
