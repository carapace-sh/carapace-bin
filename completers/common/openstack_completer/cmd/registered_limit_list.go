package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registered_limit_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List registered limits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registered_limit_listCmd).Standalone()

	registered_limit_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	registered_limit_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	registered_limit_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	registered_limit_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	registered_limit_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	registered_limit_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	registered_limit_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	registered_limit_listCmd.Flags().String("region", "", "Region for the limit to affect.")
	registered_limit_listCmd.Flags().String("resource-name", "", "The name of the resource to limit")
	registered_limit_listCmd.Flags().String("service", "", "Service responsible for the resource to limit (name or ID)")
	registered_limit_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	registered_limit_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	registered_limit_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	registered_limitCmd.AddCommand(registered_limit_listCmd)
}
