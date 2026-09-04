package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_qos_rule_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set Network QoS rule properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qos_rule_setCmd).Standalone()

	network_qos_rule_setCmd.Flags().Bool("any", false, "Any traffic direction from the project point of view.")
	network_qos_rule_setCmd.Flags().String("dscp-mark", "", "DSCP mark: value can be 0, even numbers from 8-56, excluding 42, 44, 50, 52, and 54")
	network_qos_rule_setCmd.Flags().Bool("egress", false, "Egress traffic direction from the project point of view")
	network_qos_rule_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_qos_rule_setCmd.Flags().Bool("ingress", false, "Ingress traffic direction from the project point of view")
	network_qos_rule_setCmd.Flags().String("max-burst-kbits", "", "Maximum burst in kilobits, 0 or not specified means automatic, which is 80%% of the bandwidth limit, which works for typical TCP traffic.")
	network_qos_rule_setCmd.Flags().String("max-kbps", "", "Maximum bandwidth in kbps")
	network_qos_rule_setCmd.Flags().String("min-kbps", "", "Minimum guaranteed bandwidth in kbps")
	network_qos_rule_setCmd.Flags().String("min-kpps", "", "Minimum guaranteed packet rate in kpps")
	network_qos_ruleCmd.AddCommand(network_qos_rule_setCmd)
}
