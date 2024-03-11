package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nrql_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Execute a NRQL query to New Relic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nrql_queryCmd).Standalone()
	nrql_queryCmd.Flags().StringP("query", "q", "", "the NRQL query you want to execute")
	nrqlCmd.AddCommand(nrql_queryCmd)

	// TODO query completion using shlex and api
}
