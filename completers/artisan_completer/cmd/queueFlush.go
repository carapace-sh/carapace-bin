package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueFlushCmd = &cobra.Command{
	Use:   "queue:flush [--hours [HOURS]]",
	Short: "Flush all of the failed queue jobs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueFlushCmd).Standalone()

	queueFlushCmd.Flags().String("hours", "", "The number of hours to retain failed job data")
	rootCmd.AddCommand(queueFlushCmd)
}
