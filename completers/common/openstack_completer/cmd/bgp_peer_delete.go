package cmd

import (
	"github.com/carapace-sh/carapace"
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
}
