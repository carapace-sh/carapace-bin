package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var access_rule_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List access rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_rule_listCmd).Standalone()

	access_rule_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	access_rule_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	access_rule_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	access_rule_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	access_rule_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	access_rule_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	access_rule_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	access_rule_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	access_rule_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	access_rule_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	access_rule_listCmd.Flags().String("user", "", "User whose access rules to list (name or ID)")
	access_rule_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	access_ruleCmd.AddCommand(access_rule_listCmd)
}
