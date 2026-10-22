package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var network_qos_rule_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Network QoS rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qos_rule_deleteCmd).Standalone()

	network_qos_ruleCmd.AddCommand(network_qos_rule_deleteCmd)

	carapace.Gen(network_qos_rule_deleteCmd).PositionalCompletion(
		action.ActionNetworkQosPolicies(),
		action.ActionNetworkQosRules(),
	)
}
