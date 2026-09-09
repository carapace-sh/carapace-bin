package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_router_association_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset BGP VPN router association properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_router_association_unsetCmd).Standalone()

	bgpvpn_router_association_unsetCmd.Flags().Bool("advertise_extra_routes", false, "Routes from the router will not be advertised to the BGP VPN")
	bgpvpn_router_association_unsetCmd.Flags().Bool("no-advertise_extra_routes", false, "Routes will be advertised to the BGP VPN")
	bgpvpn_router_associationCmd.AddCommand(bgpvpn_router_association_unsetCmd)
}
