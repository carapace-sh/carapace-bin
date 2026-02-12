package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_dedupCmd = &cobra.Command{
	Use:   "dedup",
	Short: "Delete duplicate history entries (that have the same command, cwd and hostname)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_dedupCmd).Standalone()

	help_historyCmd.AddCommand(help_history_dedupCmd)
}
