package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_dedupCmd = &cobra.Command{
	Use:   "dedup",
	Short: "Delete duplicate history entries (that have the same command, cwd and hostname)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_dedupCmd).Standalone()

	history_dedupCmd.Flags().StringP("before", "b", "", "Only delete results added before this date")
	history_dedupCmd.Flags().BoolP("dry-run", "n", false, "List matching history lines without performing the actual deletion")
	history_dedupCmd.Flags().String("dupkeep", "", "How many recent duplicates to keep")
	history_dedupCmd.Flags().BoolP("help", "h", false, "Print help")
	history_dedupCmd.MarkFlagRequired("before")
	history_dedupCmd.MarkFlagRequired("dupkeep")
	historyCmd.AddCommand(history_dedupCmd)
}
