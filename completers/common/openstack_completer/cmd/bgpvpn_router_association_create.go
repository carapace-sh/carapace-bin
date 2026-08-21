package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_router_association_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a BGP VPN router association",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_router_association_createCmd).Standalone()

	bgpvpn_router_association_createCmd.Flags().Bool("advertise_extra_routes", false, "Routes will be advertised to the BGP VPN (default)")
	bgpvpn_router_association_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgpvpn_router_association_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgpvpn_router_association_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgpvpn_router_association_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgpvpn_router_association_createCmd.Flags().Bool("no-advertise_extra_routes", false, "Routes from the router will not be advertised to the BGP VPN")
	bgpvpn_router_association_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgpvpn_router_association_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	bgpvpn_router_association_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgpvpn_router_association_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	bgpvpn_router_association_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	bgpvpn_router_association_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	bgpvpn_router_associationCmd.AddCommand(bgpvpn_router_association_createCmd)
}
