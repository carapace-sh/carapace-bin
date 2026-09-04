package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var router_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set router properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(router_setCmd).Standalone()

	router_setCmd.Flags().Bool("centralized", false, "Set router to centralized mode (disabled router only)")
	router_setCmd.Flags().String("description", "", "Set router description")
	router_setCmd.Flags().Bool("disable", false, "Disable router")
	router_setCmd.Flags().Bool("disable-default-route-bfd", false, "Disable BFD sessions for default routes inferred from the external gateway port subnets for this router")
	router_setCmd.Flags().Bool("disable-default-route-ecmp", false, "Add default route only for first gateway port")
	router_setCmd.Flags().Bool("disable-ndp-proxy", false, "Disable IPv6 NDP proxy on external gateway")
	router_setCmd.Flags().Bool("disable-snat", false, "Disable Source NAT on external gateway")
	router_setCmd.Flags().Bool("distributed", false, "Set router to distributed mode (disabled router only)")
	router_setCmd.Flags().Bool("enable", false, "Enable router")
	router_setCmd.Flags().Bool("enable-default-route-bfd", false, "Enable BFD sessions for default routes inferred from the external gateway port subnets for this router")
	router_setCmd.Flags().Bool("enable-default-route-ecmp", false, "Add ECMP default routes if multiple are available via different gateway ports")
	router_setCmd.Flags().Bool("enable-ndp-proxy", false, "Enable IPv6 NDP proxy on external gateway")
	router_setCmd.Flags().Bool("enable-snat", false, "Enable Source NAT on external gateway")
	router_setCmd.Flags().String("external-gateway", "", "External Network used as router's gateway (name or ID) (repeat option to set multiple gateways per router if the L3 service plugin in use supports it).")
	router_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	router_setCmd.Flags().String("fixed-ip", "", "Desired IP and/or subnet (name or ID) on external gateway: subnet=<subnet>,ip-address=<ip-address> (repeat option to set multiple fixed IP addresses)")
	router_setCmd.Flags().Bool("ha", false, "Set the router as highly available (disabled router only)")
	router_setCmd.Flags().String("name", "", "Set router name")
	router_setCmd.Flags().Bool("no-ha", false, "Clear high availability attribute of the router (disabled router only)")
	router_setCmd.Flags().Bool("no-qos-policy", false, "Remove QoS policy from router gateway IPs")
	router_setCmd.Flags().Bool("no-route", false, "Clear routes associated with the router.")
	router_setCmd.Flags().Bool("no-tag", false, "Clear tags associated with the router.")
	router_setCmd.Flags().String("qos-policy", "", "Attach QoS policy to router gateway IPs")
	router_setCmd.Flags().String("route", "", "Add routes to the router.")
	router_setCmd.Flags().String("tag", "", "Tag to be added to the router (repeat option to set multiple tags)")
	routerCmd.AddCommand(router_setCmd)
}
