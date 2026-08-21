package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catalog_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List services in the service catalog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catalog_listCmd).Standalone()

	catalog_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	catalog_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	catalog_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	catalog_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	catalog_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	catalog_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	catalog_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	catalog_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	catalog_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	catalog_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	catalogCmd.AddCommand(catalog_listCmd)
}
