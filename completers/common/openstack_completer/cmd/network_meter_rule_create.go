package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_meter_rule_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new meter rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_meter_rule_createCmd).Standalone()

	network_meter_rule_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_meter_rule_createCmd.Flags().String("destination-ip-prefix", "", "The destination IP prefix to associate with this rule")
	network_meter_rule_createCmd.Flags().Bool("egress", false, "Apply rule to outgoing network traffic")
	network_meter_rule_createCmd.Flags().Bool("exclude", false, "Exclude remote IP prefix from traffic count")
	network_meter_rule_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_meter_rule_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_meter_rule_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_meter_rule_createCmd.Flags().Bool("include", false, "Include remote IP prefix from traffic count (default)")
	network_meter_rule_createCmd.Flags().Bool("ingress", false, "Apply rule to incoming network traffic (default)")
	network_meter_rule_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_meter_rule_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_meter_rule_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_meter_rule_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_meter_rule_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	network_meter_rule_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_meter_rule_createCmd.Flags().String("remote-ip-prefix", "", "The remote IP prefix to associate with this rule")
	network_meter_rule_createCmd.Flags().String("source-ip-prefix", "", "The source IP prefix to associate with this rule")
	network_meter_rule_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_meter_ruleCmd.AddCommand(network_meter_rule_createCmd)
}
