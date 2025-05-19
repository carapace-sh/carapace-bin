package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var edge_traceObserver_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a New Relic Edge trace observer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edge_traceObserver_createCmd).Standalone()
	edge_traceObserver_createCmd.Flags().StringP("name", "n", "", "the name of the trace observer")
	edge_traceObserver_createCmd.Flags().StringP("providerRegion", "r", "", "the provider region in which to create the trace observer")
	edge_traceObserverCmd.AddCommand(edge_traceObserver_createCmd)

	carapace.Gen(edge_traceObserver_createCmd).FlagCompletion(carapace.ActionMap{
		"providerRegion": carapace.ActionValues("AWS_US_EAST_1", "AWS_US_EAST_2."),
	})
}
