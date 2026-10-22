package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var subnet_pool_add_networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Onboard network subnets into a subnet pool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_pool_add_networkCmd).Standalone()

	subnet_pool_addCmd.AddCommand(subnet_pool_add_networkCmd)

	carapace.Gen(subnet_pool_add_networkCmd).PositionalCompletion(
		action.ActionNetworks(subnet_pool_add_networkCmd),
		action.ActionSubnetPools(subnet_pool_add_networkCmd),
	)
}
