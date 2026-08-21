package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mapping_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List mappings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mapping_listCmd).Standalone()

	mapping_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	mapping_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	mapping_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	mapping_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	mapping_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	mapping_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	mapping_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	mapping_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	mapping_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	mapping_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	mappingCmd.AddCommand(mapping_listCmd)
}
