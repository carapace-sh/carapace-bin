package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var policy_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List policies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_listCmd).Standalone()

	policy_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	policy_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	policy_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	policy_listCmd.Flags().Bool("long", false, "List additional fields in output")
	policy_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	policy_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	policy_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	policy_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	policy_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	policy_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	policy_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	policyCmd.AddCommand(policy_listCmd)
}
