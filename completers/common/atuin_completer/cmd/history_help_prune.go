package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete history entries matching the configured exclusion filters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_pruneCmd).Standalone()

	history_helpCmd.AddCommand(history_help_pruneCmd)
}
