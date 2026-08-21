package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset subnet properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_unsetCmd).Standalone()

	subnet_unsetCmd.Flags().Bool("all-tag", false, "Clear all tags associated with the subnet")
	subnet_unsetCmd.Flags().String("allocation-pool", "", "Allocation pool IP addresses to be removed from this subnet, for example, start=192.168.199.2,end=192.168.199.254 (repeat option to unset multiple allocation pools)")
	subnet_unsetCmd.Flags().String("dns-nameserver", "", "DNS server to be removed from this subnet (repeat option to unset multiple DNS servers)")
	subnet_unsetCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	subnet_unsetCmd.Flags().Bool("gateway", false, "Remove gateway IP from this subnet")
	subnet_unsetCmd.Flags().String("host-route", "", "Route to be removed from this subnet, for example, destination=10.10.0.0/16,gateway=192.168.71.254 destination: destination subnet (in CIDR notation) gateway: next-hop IP address (repeat option to unset multiple host routes)")
	subnet_unsetCmd.Flags().String("service-type", "", "Service type to be removed from this subnet, for example, network:floatingip_agent_gateway.")
	subnet_unsetCmd.Flags().String("tag", "", "Tag to be removed from the subnet (repeat option to remove multiple tags)")
	subnetCmd.AddCommand(subnet_unsetCmd)
}
