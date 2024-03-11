package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var edge_traceObserver_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a New Relic Edge trace observer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edge_traceObserver_deleteCmd).Standalone()
	edge_traceObserver_deleteCmd.Flags().IntP("id", "i", 0, "the ID of the trace observer to delete")
	edge_traceObserverCmd.AddCommand(edge_traceObserver_deleteCmd)
}
