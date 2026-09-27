package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new router",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_createCmd).Standalone()

	router_createCmd.Flags().Bool("auto-evpn-vni", false, "Associate the router with an EVPN using an auto-assigned VNI.")
	router_createCmd.Flags().String("availability-zone-hint", "", "Availability Zone in which to create this router (Router Availability Zone extension required, repeat option to set multiple availability zones)")
	router_createCmd.Flags().Bool("centralized", false, "Create a centralized router")
	router_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	router_createCmd.Flags().String("description", "", "Set router description")
	router_createCmd.Flags().Bool("disable", false, "Disable router")
	router_createCmd.Flags().Bool("disable-default-route-bfd", false, "Disable BFD sessions for default routes inferred from the external gateway port subnets for this router")
	router_createCmd.Flags().Bool("disable-default-route-ecmp", false, "Add default route only for first gateway port")
	router_createCmd.Flags().Bool("disable-ndp-proxy", false, "Disable IPv6 NDP proxy on external gateway")
	router_createCmd.Flags().Bool("disable-snat", false, "Disable Source NAT on external gateway")
	router_createCmd.Flags().Bool("distributed", false, "Create a distributed router")
	router_createCmd.Flags().Bool("enable", false, "Enable router (default)")
	router_createCmd.Flags().Bool("enable-default-route-bfd", false, "Enable BFD sessions for default routes inferred from the external gateway port subnets for this router")
	router_createCmd.Flags().Bool("enable-default-route-ecmp", false, "Add ECMP default routes if multiple are available via different gateway ports")
	router_createCmd.Flags().Bool("enable-ndp-proxy", false, "Enable IPv6 NDP proxy on external gateway")
	router_createCmd.Flags().Bool("enable-snat", false, "Enable Source NAT on external gateway")
	router_createCmd.Flags().String("evpn-vni", "", "Associate the router with an EVPN identified by a VNI.")
	router_createCmd.Flags().String("external-gateway", "", "External Network used as router's gateway (name or ID) (repeat option to set multiple gateways per router if the L3 service plugin in use supports it)")
	router_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	router_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	router_createCmd.Flags().String("fixed-ip", "", "Desired IP and/or subnet (name or ID) on external gateway: subnet=<subnet>,ip-address=<ip-address> (repeat option to set multiple fixed IP addresses)")
	router_createCmd.Flags().String("flavor", "", "Associate the router to a flavor (by name or ID")
	router_createCmd.Flags().String("flavor-id", "", "==SUPPRESS==")
	router_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	router_createCmd.Flags().Bool("ha", false, "Create a highly available router")
	router_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	router_createCmd.Flags().Bool("no-ha", false, "Create a legacy router")
	router_createCmd.Flags().Bool("no-tag", false, "No tags associated with the router")
	router_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	router_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	router_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	router_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	router_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	router_createCmd.Flags().String("qos-policy", "", "Attach QoS policy to router gateway IPs")
	router_createCmd.Flags().String("tag", "", "Tag to be added to the router (repeat option to set multiple tags)")
	router_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	routerCmd.AddCommand(router_createCmd)
}
