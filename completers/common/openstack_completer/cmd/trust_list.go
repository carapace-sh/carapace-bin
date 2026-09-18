package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trust_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List trusts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_listCmd).Standalone()

	trust_listCmd.Flags().Bool("auth-user", false, "Only list trusts related to the authenticated user")
	trust_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	trust_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	trust_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	trust_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	trust_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	trust_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	trust_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	trust_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	trust_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	trust_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	trust_listCmd.Flags().String("trustee", "", "Trustee user to filter (name or ID)")
	trust_listCmd.Flags().String("trustee-domain", "", "Domain that contains <trustee> (name or ID)")
	trust_listCmd.Flags().String("trustor", "", "Trustor user to filter (name or ID)")
	trust_listCmd.Flags().String("trustor-domain", "", "Domain that contains <trustor> (name or ID)")
	trustCmd.AddCommand(trust_listCmd)
}
