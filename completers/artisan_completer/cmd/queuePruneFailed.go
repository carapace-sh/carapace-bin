package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queuePruneFailedCmd = &cobra.Command{
	Use:   "queue:prune-failed [--hours [HOURS]]",
	Short: "Prune stale entries from the failed jobs table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queuePruneFailedCmd).Standalone()

	queuePruneFailedCmd.Flags().String("hours", "", "The number of hours to retain failed jobs data")
	rootCmd.AddCommand(queuePruneFailedCmd)
}
