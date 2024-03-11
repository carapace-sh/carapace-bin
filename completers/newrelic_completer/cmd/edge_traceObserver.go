package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var edge_traceObserverCmd = &cobra.Command{
	Use:   "trace-observer",
	Short: "Interact with New Relic Edge trace observers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edge_traceObserverCmd).Standalone()
	edgeCmd.AddCommand(edge_traceObserverCmd)
}
