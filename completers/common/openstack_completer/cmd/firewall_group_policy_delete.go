package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var firewall_group_policy_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete firewall policy(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_policy_deleteCmd).Standalone()

	firewall_group_policyCmd.AddCommand(firewall_group_policy_deleteCmd)

	carapace.Gen(firewall_group_policy_deleteCmd).PositionalAnyCompletion(action.ActionFirewallPolicies(firewall_group_policy_deleteCmd))
}
