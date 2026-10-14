package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var implied_role_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List implied roles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(implied_role_listCmd).Standalone()

	implied_role_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	implied_role_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	implied_role_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	implied_role_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	implied_role_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	implied_role_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	implied_role_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	implied_role_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	implied_role_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	implied_role_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	implied_roleCmd.AddCommand(implied_role_listCmd)
}
