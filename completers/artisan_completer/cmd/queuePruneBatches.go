package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queuePruneBatchesCmd = &cobra.Command{
	Use:   "queue:prune-batches [--hours [HOURS]] [--unfinished [UNFINISHED]] [--cancelled [CANCELLED]]",
	Short: "Prune stale entries from the batches database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queuePruneBatchesCmd).Standalone()

	queuePruneBatchesCmd.Flags().String("cancelled", "", "The number of hours to retain cancelled batch data")
	queuePruneBatchesCmd.Flags().String("hours", "", "The number of hours to retain batch data")
	queuePruneBatchesCmd.Flags().String("unfinished", "", "The number of hours to retain unfinished batch data")
	rootCmd.AddCommand(queuePruneBatchesCmd)
}
