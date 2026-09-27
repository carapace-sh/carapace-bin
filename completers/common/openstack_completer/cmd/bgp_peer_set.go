package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_peer_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update a BGP peer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_peer_setCmd).Standalone()

	bgp_peer_setCmd.Flags().String("name", "", "Updated name of the BGP peer")
	bgp_peer_setCmd.Flags().String("password", "", "Updated authentication password")
	bgp_peerCmd.AddCommand(bgp_peer_setCmd)
}
