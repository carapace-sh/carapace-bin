package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_qos_rule_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Network QoS rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qos_rule_listCmd).Standalone()

	network_qos_rule_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_qos_rule_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_qos_rule_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_qos_rule_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_qos_rule_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_qos_rule_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_qos_rule_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	network_qos_rule_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	network_qos_rule_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	network_qos_rule_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	network_qos_ruleCmd.AddCommand(network_qos_rule_listCmd)
}
