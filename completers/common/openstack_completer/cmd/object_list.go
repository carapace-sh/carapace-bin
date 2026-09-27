package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var object_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_listCmd).Standalone()

	object_listCmd.Flags().Bool("all", false, "List all objects in container (default is 10000)")
	object_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	object_listCmd.Flags().String("delimiter", "", "Roll up items with <delimiter>")
	object_listCmd.Flags().String("end-marker", "", "End anchor for paging")
	object_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	object_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	object_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	object_listCmd.Flags().Bool("long", false, "List additional fields in output")
	object_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	object_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	object_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	object_listCmd.Flags().String("prefix", "", "Filter list using <prefix>")
	object_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	object_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	object_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	object_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	object_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	objectCmd.AddCommand(object_listCmd)
}
