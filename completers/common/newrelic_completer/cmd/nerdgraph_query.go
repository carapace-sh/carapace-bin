package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdgraph_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Execute a raw GraphQL query request to the NerdGraph API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdgraph_queryCmd).Standalone()
	nerdgraph_queryCmd.Flags().String("queryFile", "", "the query to use, represented as a JSON file")
	nerdgraph_queryCmd.Flags().String("variables", "{}", "the variables to pass to the GraphQL query, represented as a JSON string (this will overwrite duplicate variables that are set through variablesFile")
	nerdgraph_queryCmd.Flags().String("variablesFile", "", "the variables to pass to GraphQL query, represented as a JSON file")
	nerdgraphCmd.AddCommand(nerdgraph_queryCmd)

	carapace.Gen(nerdgraph_queryCmd).FlagCompletion(carapace.ActionMap{
		"queryFile":     carapace.ActionFiles(),
		"variablesFile": carapace.ActionFiles(),
	})
}
