package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset BGP VPN properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_unsetCmd).Standalone()

	bgpvpn_unsetCmd.Flags().Bool("all-export-target", false, "Empty export route target list")
	bgpvpn_unsetCmd.Flags().Bool("all-import-target", false, "Empty import route target list")
	bgpvpn_unsetCmd.Flags().Bool("all-route-distinguisher", false, "Empty route distinguisher list")
	bgpvpn_unsetCmd.Flags().Bool("all-route-target", false, "Empty route target list")
	bgpvpn_unsetCmd.Flags().String("export-target", "", "Remove Route Target from export list (repeat option for multiple Route Targets)")
	bgpvpn_unsetCmd.Flags().String("import-target", "", "Remove Route Target from import list (repeat option for multiple Route Targets)")
	bgpvpn_unsetCmd.Flags().String("local-pref", "", "Default BGP LOCAL_PREF to use in route advertisementstowards this BGPVPN.")
	bgpvpn_unsetCmd.Flags().String("route-distinguisher", "", "Remove Route Distinguisher from the list of Route Distinguishers from which a Route Distinguishers will be picked from to advertise a VPN route (repeat option for multiple Route Distinguishers)")
	bgpvpn_unsetCmd.Flags().String("route-target", "", "Remove Route Target from import/export list (repeat option for multiple Route Targets)")
	bgpvpn_unsetCmd.Flags().String("vni", "", "VXLAN Network Identifier to be used for this BGPVPN when a VXLAN encapsulation is used")
	bgpvpnCmd.AddCommand(bgpvpn_unsetCmd)
}
