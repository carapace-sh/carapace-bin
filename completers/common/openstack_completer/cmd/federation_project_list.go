package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federation_project_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List accessible projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federation_project_listCmd).Standalone()

	federation_project_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	federation_project_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	federation_project_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	federation_project_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	federation_project_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	federation_project_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	federation_project_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	federation_project_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	federation_project_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	federation_project_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	federation_projectCmd.AddCommand(federation_project_listCmd)
}
