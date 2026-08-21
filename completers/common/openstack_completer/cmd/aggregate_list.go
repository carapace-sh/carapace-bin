package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aggregate_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all aggregates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aggregate_listCmd).Standalone()

	aggregate_listCmd.Flags().String("availability-zone", "", "Filter by availability zone name")
	aggregate_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	aggregate_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	aggregate_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	aggregate_listCmd.Flags().String("host", "", "Filter by aggregates containing this host")
	aggregate_listCmd.Flags().Bool("long", false, "List additional fields in output")
	aggregate_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	aggregate_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	aggregate_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	aggregate_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	aggregate_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	aggregate_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	aggregate_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	aggregateCmd.AddCommand(aggregate_listCmd)
}
