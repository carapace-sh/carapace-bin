package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var edge_traceObserver_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the New Relic Edge trace observers for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edge_traceObserver_listCmd).Standalone()
	edge_traceObserverCmd.AddCommand(edge_traceObserver_listCmd)
}
