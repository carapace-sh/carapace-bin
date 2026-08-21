package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set BGP VPN properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_setCmd).Standalone()

	bgpvpn_setCmd.Flags().String("export-target", "", "Add Route Target to export list (repeat option for multiple Route Targets)")
	bgpvpn_setCmd.Flags().String("import-target", "", "Add Route Target to import list (repeat option for multiple Route Targets)")
	bgpvpn_setCmd.Flags().String("local-pref", "", "Default BGP LOCAL_PREF to use in route advertisementstowards this BGPVPN.")
	bgpvpn_setCmd.Flags().String("name", "", "Name of the BGP VPN")
	bgpvpn_setCmd.Flags().Bool("no-export-target", false, "Empty export route target list")
	bgpvpn_setCmd.Flags().Bool("no-import-target", false, "Empty import route target list")
	bgpvpn_setCmd.Flags().Bool("no-route-distinguisher", false, "Empty route distinguisher list")
	bgpvpn_setCmd.Flags().Bool("no-route-target", false, "Empty route target list")
	bgpvpn_setCmd.Flags().String("route-distinguisher", "", "Add Route Distinguisher to the list of Route Distinguishers from which a Route Distinguishers will be picked from to advertise a VPN route (repeat option for multiple Route Distinguishers)")
	bgpvpn_setCmd.Flags().String("route-target", "", "Add Route Target to import/export list (repeat option for multiple Route Targets)")
	bgpvpn_setCmd.Flags().String("vni", "", "VXLAN Network Identifier to be used for this BGPVPN when a VXLAN encapsulation is used")
	bgpvpnCmd.AddCommand(bgpvpn_setCmd)
}
