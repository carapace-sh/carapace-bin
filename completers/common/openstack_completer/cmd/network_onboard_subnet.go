package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var network_onboard_subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "DEPRECATED: Use 'subnet pool add network' instead.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_onboard_subnetCmd).Standalone()

	network_onboardCmd.AddCommand(network_onboard_subnetCmd)

	carapace.Gen(network_onboard_subnetCmd).PositionalCompletion(
		action.ActionNetworks(),
		action.ActionSubnetPools(),
	)
}
