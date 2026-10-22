package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var firewall_group_policy_add_ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "Insert a rule into a given firewall policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_policy_add_ruleCmd).Standalone()

	firewall_group_policy_add_ruleCmd.Flags().String("insert-after", "", "Insert the new rule after this existing rule  (name or ID)")
	firewall_group_policy_add_ruleCmd.Flags().String("insert-before", "", "Insert the new rule before this existing rule  (name or ID)")
	firewall_group_policy_addCmd.AddCommand(firewall_group_policy_add_ruleCmd)

	carapace.Gen(firewall_group_policy_add_ruleCmd).PositionalCompletion(
		action.ActionFirewallPolicies(firewall_group_policy_add_ruleCmd),
		action.ActionFirewallRules(firewall_group_policy_add_ruleCmd),
	)
}
