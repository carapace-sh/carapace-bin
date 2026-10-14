package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_segment_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List network segments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_segment_listCmd).Standalone()

	network_segment_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_segment_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_segment_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_segment_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	network_segment_listCmd.Flags().Bool("long", false, "List additional fields in output")
	network_segment_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	network_segment_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	network_segment_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_segment_listCmd.Flags().String("network", "", "List only network segments associated with the specified network (name or ID)")
	network_segment_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_segment_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_segment_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_segment_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_segment_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_segment_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_segmentCmd.AddCommand(network_segment_listCmd)
}
