package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_l3_conntrack_helper_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List L3 conntrack helpers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_l3_conntrack_helper_listCmd).Standalone()

	network_l3_conntrack_helper_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_l3_conntrack_helper_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_l3_conntrack_helper_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_l3_conntrack_helper_listCmd.Flags().String("helper", "", "List only helpers using the specified netfilter conntrack helper module")
	network_l3_conntrack_helper_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	network_l3_conntrack_helper_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	network_l3_conntrack_helper_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	network_l3_conntrack_helper_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_l3_conntrack_helper_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_l3_conntrack_helper_listCmd.Flags().String("port", "", "List only helpers with the specified network port for the netfilter conntrack target rule (name or ID)")
	network_l3_conntrack_helper_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_l3_conntrack_helper_listCmd.Flags().String("protocol", "", "List only helpers with the specified network protocol for the netfilter conntrack target rule")
	network_l3_conntrack_helper_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_l3_conntrack_helper_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_l3_conntrack_helper_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_l3_conntrack_helper_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_l3_conntrack_helperCmd.AddCommand(network_l3_conntrack_helper_listCmd)
}
