package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var network_flavor_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network flavors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_flavor_deleteCmd).Standalone()

	network_flavorCmd.AddCommand(network_flavor_deleteCmd)

	carapace.Gen(network_flavor_deleteCmd).PositionalAnyCompletion(action.ActionNetworkFlavors())
}
