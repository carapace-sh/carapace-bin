package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_agent_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List network agents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_agent_listCmd).Standalone()

	network_agent_listCmd.Flags().String("agent-type", "", "List only agents with the specified agent type.")
	network_agent_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_agent_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_agent_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_agent_listCmd.Flags().String("host", "", "List only agents running on the specified host")
	network_agent_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	network_agent_listCmd.Flags().Bool("long", false, "List additional fields in output")
	network_agent_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	network_agent_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	network_agent_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_agent_listCmd.Flags().String("network", "", "List agents hosting the specified network (name or ID)")
	network_agent_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_agent_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_agent_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_agent_listCmd.Flags().String("router", "", "List agents hosting the specified router (name or ID)")
	network_agent_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_agent_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_agent_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_agentCmd.AddCommand(network_agent_listCmd)
}
