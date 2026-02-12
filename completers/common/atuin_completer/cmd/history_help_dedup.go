package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_dedupCmd = &cobra.Command{
	Use:   "dedup",
	Short: "Delete duplicate history entries (that have the same command, cwd and hostname)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_dedupCmd).Standalone()

	history_helpCmd.AddCommand(history_help_dedupCmd)
}
