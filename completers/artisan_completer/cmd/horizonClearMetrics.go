package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonClearMetricsCmd = &cobra.Command{
	Use:   "horizon:clear-metrics",
	Short: "Delete metrics for all jobs and queues",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonClearMetricsCmd).Standalone()

	rootCmd.AddCommand(horizonClearMetricsCmd)
}
