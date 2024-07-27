package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueMonitorCmd = &cobra.Command{
	Use:   "queue:monitor [--max [MAX]] [--] <queues>",
	Short: "Monitor the size of the specified queues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueMonitorCmd).Standalone()

	queueMonitorCmd.Flags().String("max", "", "The maximum number of jobs that can be on the queue before an event is dispatched")
	rootCmd.AddCommand(queueMonitorCmd)
}
