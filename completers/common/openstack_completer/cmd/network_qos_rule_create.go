package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_qos_rule_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new Network QoS rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qos_rule_createCmd).Standalone()

	network_qos_rule_createCmd.Flags().Bool("any", false, "Any traffic direction from the project point of view.")
	network_qos_rule_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_qos_rule_createCmd.Flags().String("dscp-mark", "", "DSCP mark: value can be 0, even numbers from 8-56, excluding 42, 44, 50, 52, and 54")
	network_qos_rule_createCmd.Flags().Bool("egress", false, "Egress traffic direction from the project point of view")
	network_qos_rule_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_qos_rule_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_qos_rule_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_qos_rule_createCmd.Flags().Bool("ingress", false, "Ingress traffic direction from the project point of view")
	network_qos_rule_createCmd.Flags().String("max-burst-kbits", "", "Maximum burst in kilobits, 0 or not specified means automatic, which is 80%% of the bandwidth limit, which works for typical TCP traffic.")
	network_qos_rule_createCmd.Flags().String("max-kbps", "", "Maximum bandwidth in kbps")
	network_qos_rule_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_qos_rule_createCmd.Flags().String("min-kbps", "", "Minimum guaranteed bandwidth in kbps")
	network_qos_rule_createCmd.Flags().String("min-kpps", "", "Minimum guaranteed packet rate in kpps")
	network_qos_rule_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_qos_rule_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_qos_rule_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_qos_rule_createCmd.Flags().String("type", "", "QoS rule type (minimum-bandwidth, minimum-packet-rate, dscp-marking, bandwidth-limit)")
	network_qos_rule_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	network_qos_rule_createCmd.MarkFlagRequired("type")
	network_qos_ruleCmd.AddCommand(network_qos_rule_createCmd)
}
