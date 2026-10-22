package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var network_trunk_subport_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a subport to a given network trunk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_trunk_subport_addCmd).Standalone()

	network_trunk_subport_addCmd.Flags().String("segmentation-id", "", "Segmentation ID of the subport")
	network_trunk_subport_addCmd.Flags().String("segmentation-type", "", "Segmentation type of the subport, e.g. 'vlan'")
	network_trunk_subportCmd.AddCommand(network_trunk_subport_addCmd)

	carapace.Gen(network_trunk_subport_addCmd).PositionalCompletion(
		action.ActionNetworkTrunks(),
		action.ActionPorts(),
	)
}
