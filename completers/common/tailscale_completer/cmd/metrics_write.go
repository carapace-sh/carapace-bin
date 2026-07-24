package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var metrics_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write metric values to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metrics_writeCmd).Standalone()

	metricsCmd.AddCommand(metrics_writeCmd)

	carapace.Gen(metrics_writeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
