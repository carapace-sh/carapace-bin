package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set subnet properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_setCmd).Standalone()

	subnet_setCmd.Flags().String("allocation-pool", "", "Allocation pool IP addresses for this subnet, for example, start=192.168.199.2,end=192.168.199.254 (repeat option to add multiple IP addresses)")
	subnet_setCmd.Flags().String("description", "", "Set subnet description")
	subnet_setCmd.Flags().Bool("dhcp", false, "Enable DHCP")
	subnet_setCmd.Flags().String("dns-nameserver", "", "DNS server for this subnet (repeat option to set multiple DNS servers)")
	subnet_setCmd.Flags().Bool("dns-publish-fixed-ip", false, "Enable publishing fixed IPs in DNS")
	subnet_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	subnet_setCmd.Flags().String("gateway", "", "Specify a gateway for the subnet.")
	subnet_setCmd.Flags().String("host-route", "", "Additional route for this subnet, for example, destination=10.10.0.0/16,gateway=192.168.71.254 destination: destination subnet (in CIDR notation) gateway: next-hop IP address (repeat option to add multiple routes)")
	subnet_setCmd.Flags().String("name", "", "Updated name of the subnet")
	subnet_setCmd.Flags().String("network-segment", "", "Network segment to associate with this subnet (name or ID).")
	subnet_setCmd.Flags().Bool("no-allocation-pool", false, "Clear associated allocation-pools from the subnet.")
	subnet_setCmd.Flags().Bool("no-dhcp", false, "Disable DHCP")
	subnet_setCmd.Flags().Bool("no-dns-nameservers", false, "Clear existing information of DNS Nameservers.")
	subnet_setCmd.Flags().Bool("no-dns-publish-fixed-ip", false, "Disable publishing fixed IPs in DNS")
	subnet_setCmd.Flags().Bool("no-host-route", false, "Clear associated host-routes from the subnet.")
	subnet_setCmd.Flags().Bool("no-tag", false, "Clear tags associated with the subnet.")
	subnet_setCmd.Flags().String("service-type", "", "Service type for this subnet, for example, network:floatingip_agent_gateway.")
	subnet_setCmd.Flags().String("tag", "", "Tag to be added to the subnet (repeat option to set multiple tags)")
	subnetCmd.AddCommand(subnet_setCmd)
}
