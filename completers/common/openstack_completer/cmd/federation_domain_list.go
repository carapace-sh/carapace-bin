package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federation_domain_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List accessible domains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federation_domain_listCmd).Standalone()

	federation_domain_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	federation_domain_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	federation_domain_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	federation_domain_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	federation_domain_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	federation_domain_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	federation_domain_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	federation_domain_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	federation_domain_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	federation_domain_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	federation_domainCmd.AddCommand(federation_domain_listCmd)
}
