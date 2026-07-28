package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var metrics_printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print current metric values in Prometheus text format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metrics_printCmd).Standalone()

	metricsCmd.AddCommand(metrics_printCmd)
}
