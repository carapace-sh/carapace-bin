package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_event_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List recent events of a server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_event_listCmd).Standalone()

	server_event_listCmd.Flags().String("changes-before", "", "List only server events changed earlier or equal to a certain point of time.")
	server_event_listCmd.Flags().String("changes-since", "", "List only server events changed later or equal to a certain point of time.")
	server_event_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	server_event_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	server_event_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	server_event_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	server_event_listCmd.Flags().Bool("long", false, "List additional fields in output")
	server_event_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	server_event_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	server_event_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	server_event_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	server_event_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	server_event_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	server_event_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	server_event_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	server_event_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	server_eventCmd.AddCommand(server_event_listCmd)
}
