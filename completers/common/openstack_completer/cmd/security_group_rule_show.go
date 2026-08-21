package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_rule_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display security group rule details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_rule_showCmd).Standalone()

	security_group_rule_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	security_group_rule_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	security_group_rule_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	security_group_rule_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	security_group_rule_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	security_group_rule_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	security_group_rule_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	security_group_rule_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	security_group_ruleCmd.AddCommand(security_group_rule_showCmd)
}
