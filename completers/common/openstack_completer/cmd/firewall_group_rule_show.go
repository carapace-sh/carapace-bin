package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var firewall_group_rule_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display firewall rule details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_rule_showCmd).Standalone()

	firewall_group_rule_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	firewall_group_rule_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	firewall_group_rule_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	firewall_group_rule_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	firewall_group_rule_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	firewall_group_rule_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	firewall_group_rule_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	firewall_group_rule_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	firewall_group_ruleCmd.AddCommand(firewall_group_rule_showCmd)
}
