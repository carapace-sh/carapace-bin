package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_rule_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List firewall rules that belong to a given tenant",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_rule_listCmd).Standalone()

	firewall_group_rule_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	firewall_group_rule_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	firewall_group_rule_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	firewall_group_rule_listCmd.Flags().Bool("long", false, "List additional fields in output")
	firewall_group_rule_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	firewall_group_rule_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	firewall_group_rule_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	firewall_group_rule_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	firewall_group_rule_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	firewall_group_rule_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	firewall_group_rule_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	firewall_group_ruleCmd.AddCommand(firewall_group_rule_listCmd)
}
