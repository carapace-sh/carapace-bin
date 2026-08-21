package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List consistency groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_listCmd).Standalone()

	consistency_group_listCmd.Flags().Bool("all-projects", false, "Show details for all projects.")
	consistency_group_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	consistency_group_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	consistency_group_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	consistency_group_listCmd.Flags().Bool("long", false, "List additional fields in output")
	consistency_group_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	consistency_group_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	consistency_group_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	consistency_group_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	consistency_group_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	consistency_group_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	consistency_group_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	consistency_groupCmd.AddCommand(consistency_group_listCmd)
}
