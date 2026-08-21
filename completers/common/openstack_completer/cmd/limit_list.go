package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var limit_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List limits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(limit_listCmd).Standalone()

	limit_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	limit_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	limit_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	limit_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	limit_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	limit_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	limit_listCmd.Flags().String("project", "", "List resource limits associated with project")
	limit_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	limit_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	limit_listCmd.Flags().String("region", "", "Region for the registered limit to affect.")
	limit_listCmd.Flags().String("resource-name", "", "The name of the resource to limit")
	limit_listCmd.Flags().String("service", "", "Service responsible for the resource to limit")
	limit_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	limit_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	limit_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	limitCmd.AddCommand(limit_listCmd)
}
