package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_network_association_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a BGP VPN network association(s) for a given BGP VPN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_network_association_deleteCmd).Standalone()

	bgpvpn_network_associationCmd.AddCommand(bgpvpn_network_association_deleteCmd)
}
