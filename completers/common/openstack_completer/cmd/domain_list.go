package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var domain_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List domains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(domain_listCmd).Standalone()

	domain_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	domain_listCmd.Flags().Bool("enabled", false, "The domains that are enabled will be returned")
	domain_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	domain_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	domain_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	domain_listCmd.Flags().String("name", "", "The domain name")
	domain_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	domain_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	domain_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	domain_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	domain_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	domain_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	domainCmd.AddCommand(domain_listCmd)
}
