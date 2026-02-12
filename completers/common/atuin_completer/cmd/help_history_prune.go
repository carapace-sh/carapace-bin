package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete history entries matching the configured exclusion filters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_pruneCmd).Standalone()

	help_historyCmd.AddCommand(help_history_pruneCmd)
}
