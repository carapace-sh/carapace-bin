package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_port_association_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set BGP VPN port association properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_port_association_setCmd).Standalone()

	bgpvpn_port_association_setCmd.Flags().Bool("advertise-fixed-ips", false, "Fixed IPs of the port will be advertised to the BGP VPN")
	bgpvpn_port_association_setCmd.Flags().String("bgpvpn-route", "", "Add BGP VPN route for route leaking.")
	bgpvpn_port_association_setCmd.Flags().Bool("no-advertise-fixed-ips", false, "Fixed IPs of the port will not be advertised to the BGP VPN")
	bgpvpn_port_association_setCmd.Flags().Bool("no-bgpvpn-route", false, "Empty BGP VPN route list")
	bgpvpn_port_association_setCmd.Flags().Bool("no-prefix-route", false, "Empty prefix route list")
	bgpvpn_port_association_setCmd.Flags().String("prefix-route", "", "Add prefix route in CIDR notation.")
	bgpvpn_port_associationCmd.AddCommand(bgpvpn_port_association_setCmd)
}
