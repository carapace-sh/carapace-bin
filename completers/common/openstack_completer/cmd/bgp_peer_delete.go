package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bgp_peer_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a BGP peer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_peer_deleteCmd).Standalone()

	bgp_peerCmd.AddCommand(bgp_peer_deleteCmd)

	carapace.Gen(bgp_peer_deleteCmd).PositionalAnyCompletion(action.ActionBgpPeers(bgp_peer_deleteCmd))
}
