package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_listCmd).Standalone()

	container_listCmd.Flags().Bool("all", false, "List all containers (default is 10000)")
	container_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	container_listCmd.Flags().String("end-marker", "", "End anchor for paging")
	container_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	container_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	container_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	container_listCmd.Flags().Bool("long", false, "List additional fields in output")
	container_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	container_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	container_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	container_listCmd.Flags().String("prefix", "", "Filter list using <prefix>")
	container_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	container_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	container_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	container_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	container_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	containerCmd.AddCommand(container_listCmd)
}
