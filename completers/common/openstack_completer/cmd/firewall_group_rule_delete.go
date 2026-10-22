package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var firewall_group_rule_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete firewall rule(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_rule_deleteCmd).Standalone()

	firewall_group_ruleCmd.AddCommand(firewall_group_rule_deleteCmd)

	carapace.Gen(firewall_group_rule_deleteCmd).PositionalAnyCompletion(action.ActionFirewallRules(firewall_group_rule_deleteCmd))
}
