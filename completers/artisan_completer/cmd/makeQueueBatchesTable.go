package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeQueueBatchesTableCmd = &cobra.Command{
	Use:     "make:queue-batches-table",
	Short:   "Create a migration for the batches database table",
	Aliases: []string{"queue:batches-table"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeQueueBatchesTableCmd).Standalone()

	rootCmd.AddCommand(makeQueueBatchesTableCmd)
}
