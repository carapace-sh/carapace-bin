package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_meter_rule_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List meter rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_meter_rule_listCmd).Standalone()

	network_meter_rule_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_meter_rule_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_meter_rule_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_meter_rule_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	network_meter_rule_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	network_meter_rule_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	network_meter_rule_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_meter_rule_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_meter_rule_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_meter_rule_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_meter_rule_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_meter_rule_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_meter_rule_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_meter_ruleCmd.AddCommand(network_meter_rule_listCmd)
}
