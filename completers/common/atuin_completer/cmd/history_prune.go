package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete history entries matching the configured exclusion filters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_pruneCmd).Standalone()

	history_pruneCmd.Flags().BoolP("dry-run", "n", false, "List matching history lines without performing the actual deletion")
	history_pruneCmd.Flags().BoolP("help", "h", false, "Print help")
	historyCmd.AddCommand(history_pruneCmd)
}
