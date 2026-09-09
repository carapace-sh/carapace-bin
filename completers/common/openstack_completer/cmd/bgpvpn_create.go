package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create BGP VPN resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_createCmd).Standalone()

	bgpvpn_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgpvpn_createCmd.Flags().String("export-target", "", "Add Route Target to export list (repeat option for multiple Route Targets)")
	bgpvpn_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgpvpn_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgpvpn_createCmd.Flags().String("import-target", "", "Add Route Target to import list (repeat option for multiple Route Targets)")
	bgpvpn_createCmd.Flags().String("local-pref", "", "Default BGP LOCAL_PREF to use in route advertisementstowards this BGPVPN.")
	bgpvpn_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgpvpn_createCmd.Flags().String("name", "", "Name of the BGP VPN")
	bgpvpn_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgpvpn_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	bgpvpn_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgpvpn_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	bgpvpn_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	bgpvpn_createCmd.Flags().String("route-distinguisher", "", "Add Route Distinguisher to the list of Route Distinguishers from which a Route Distinguishers will be picked from to advertise a VPN route (repeat option for multiple Route Distinguishers)")
	bgpvpn_createCmd.Flags().String("route-target", "", "Add Route Target to import/export list (repeat option for multiple Route Targets)")
	bgpvpn_createCmd.Flags().String("type", "", "BGP VPN type selection between IP VPN (l3) and Ethernet VPN (l2) (default: l3)")
	bgpvpn_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	bgpvpn_createCmd.Flags().String("vni", "", "VXLAN Network Identifier to be used for this BGPVPN when a VXLAN encapsulation is used")
	bgpvpnCmd.AddCommand(bgpvpn_createCmd)
}
