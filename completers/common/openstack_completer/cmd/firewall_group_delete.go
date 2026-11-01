package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var firewall_group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete firewall group(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(firewall_group_deleteCmd).Standalone()

	firewall_groupCmd.AddCommand(firewall_group_deleteCmd)

	carapace.Gen(firewall_group_deleteCmd).PositionalAnyCompletion(action.ActionFirewallGroups(firewall_group_deleteCmd))
}
