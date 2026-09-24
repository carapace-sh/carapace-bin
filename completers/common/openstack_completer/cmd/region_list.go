package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var region_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List regions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(region_listCmd).Standalone()

	region_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	region_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	region_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	region_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	region_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	region_listCmd.Flags().String("parent-region", "", "Filter by parent region ID")
	region_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	region_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	region_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	region_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	region_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	regionCmd.AddCommand(region_listCmd)
}
