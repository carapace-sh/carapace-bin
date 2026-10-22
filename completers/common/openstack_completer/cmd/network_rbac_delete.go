package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var network_rbac_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network RBAC policy(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_rbac_deleteCmd).Standalone()

	network_rbacCmd.AddCommand(network_rbac_deleteCmd)

	carapace.Gen(network_rbac_deleteCmd).PositionalAnyCompletion(action.ActionNetworkRbacPolicies(network_rbac_deleteCmd))
}
