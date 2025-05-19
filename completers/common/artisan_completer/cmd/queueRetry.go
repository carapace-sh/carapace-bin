package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueRetryCmd = &cobra.Command{
	Use:   "queue:retry [--queue [QUEUE]] [--range [RANGE]] [--] [<id>...]",
	Short: "Retry a failed queue job",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueRetryCmd).Standalone()

	queueRetryCmd.Flags().String("queue", "", "Retry all of the failed jobs for the specified queue")
	queueRetryCmd.Flags().String("range", "", "Range of job IDs (numeric) to be retried (e.g. 1-5)")
	rootCmd.AddCommand(queueRetryCmd)
}
