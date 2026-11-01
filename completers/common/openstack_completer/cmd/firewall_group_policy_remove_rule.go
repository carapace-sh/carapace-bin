package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var firewall_group_policy_remove_ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "Remove a rule from a given firewall policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_policy_remove_ruleCmd).Standalone()

	firewall_group_policy_removeCmd.AddCommand(firewall_group_policy_remove_ruleCmd)

	carapace.Gen(firewall_group_policy_remove_ruleCmd).PositionalCompletion(
		action.ActionFirewallPolicies(firewall_group_policy_remove_ruleCmd),
		action.ActionFirewallRules(firewall_group_policy_remove_ruleCmd),
	)
}
