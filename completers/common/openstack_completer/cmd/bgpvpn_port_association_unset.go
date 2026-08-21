package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_port_association_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset BGP VPN port association properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_port_association_unsetCmd).Standalone()

	bgpvpn_port_association_unsetCmd.Flags().Bool("advertise-fixed-ips", false, "Fixed IPs of the port will not be advertised to the BGP VPN")
	bgpvpn_port_association_unsetCmd.Flags().Bool("all-bgpvpn-routes", false, "Empty BGP VPN route list")
	bgpvpn_port_association_unsetCmd.Flags().Bool("all-prefix-routes", false, "Empty prefix route list")
	bgpvpn_port_association_unsetCmd.Flags().String("bgpvpn-route", "", "Remove BGP VPN route (repeat option for multiple BGP VPN routes)")
	bgpvpn_port_association_unsetCmd.Flags().Bool("no-advertise-fixed-ips", false, "Fixed IPs of the port will be advertised to the BGP VPN")
	bgpvpn_port_association_unsetCmd.Flags().String("prefix-route", "", "Remove prefix route in CIDR notation (repeat option for multiple prefix routes)")
	bgpvpn_port_associationCmd.AddCommand(bgpvpn_port_association_unsetCmd)
}
