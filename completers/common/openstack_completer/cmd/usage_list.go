package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var usage_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List resource usage per project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(usage_listCmd).Standalone()

	usage_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	usage_listCmd.Flags().String("end", "", "Usage range end date, ex 2012-01-20 (default: tomorrow)")
	usage_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	usage_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	usage_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	usage_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	usage_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	usage_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	usage_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	usage_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	usage_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	usage_listCmd.Flags().String("start", "", "Usage range start date, ex 2012-01-20 (default: 4 weeks ago)")
	usageCmd.AddCommand(usage_listCmd)
}
