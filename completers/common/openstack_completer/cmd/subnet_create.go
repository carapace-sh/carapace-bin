package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var subnet_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a subnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(subnet_createCmd).Standalone()

	subnet_createCmd.Flags().String("allocation-pool", "", "Allocation pool IP addresses for this subnet, for example, start=192.168.199.2,end=192.168.199.254 (repeat option to add multiple IP addresses)")
	subnet_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	subnet_createCmd.Flags().String("description", "", "Set subnet description")
	subnet_createCmd.Flags().Bool("dhcp", false, "Enable DHCP (default)")
	subnet_createCmd.Flags().String("dns-nameserver", "", "DNS server for this subnet (repeat option to set multiple DNS servers)")
	subnet_createCmd.Flags().Bool("dns-publish-fixed-ip", false, "Enable publishing fixed IPs in DNS")
	subnet_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	subnet_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	subnet_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	subnet_createCmd.Flags().String("gateway", "", "Specify a gateway for the subnet.")
	subnet_createCmd.Flags().String("host-route", "", "Additional route for this subnet, for example, destination=10.10.0.0/16,gateway=192.168.71.254 destination: destination subnet (in CIDR notation) gateway: next-hop IP address (repeat option to add multiple routes)")
	subnet_createCmd.Flags().String("ip-version", "", "IP version (default is 4).")
	subnet_createCmd.Flags().String("ipv6-address-mode", "", "IPv6 address mode, valid modes: [dhcpv6-stateful, dhcpv6-stateless, slaac]")
	subnet_createCmd.Flags().String("ipv6-ra-mode", "", "IPv6 RA (Router Advertisement) mode, valid modes: [dhcpv6-stateful, dhcpv6-stateless, slaac]")
	subnet_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	subnet_createCmd.Flags().String("network", "", "Network this subnet belongs to (name or ID)")
	subnet_createCmd.Flags().String("network-segment", "", "Network segment to associate with this subnet (name or ID)")
	subnet_createCmd.Flags().Bool("no-dhcp", false, "Disable DHCP")
	subnet_createCmd.Flags().Bool("no-dns-publish-fixed-ip", false, "Disable publishing fixed IPs in DNS (default)")
	subnet_createCmd.Flags().Bool("no-tag", false, "No tags associated with the subnet")
	subnet_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	subnet_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	subnet_createCmd.Flags().String("prefix-length", "", "Prefix length for subnet allocation from subnet pool")
	subnet_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	subnet_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	subnet_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	subnet_createCmd.Flags().String("service-type", "", "Service type for this subnet, for example, network:floatingip_agent_gateway.")
	subnet_createCmd.Flags().String("subnet-pool", "", "Subnet pool from which this subnet will obtain a CIDR (Name or ID)")
	subnet_createCmd.Flags().String("subnet-range", "", "Subnet range in CIDR notation (required if --subnet-pool is not specified, optional otherwise)")
	subnet_createCmd.Flags().String("tag", "", "Tag to be added to the subnet (repeat option to set multiple tags)")
	subnet_createCmd.Flags().Bool("use-default-subnet-pool", false, "Use default subnet pool for --ip-version")
	subnet_createCmd.Flags().Bool("use-prefix-delegation", false, "Use 'prefix-delegation' if IP is IPv6 format and IP would be delegated externally")
	subnet_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	subnet_createCmd.MarkFlagRequired("network")
	subnetCmd.AddCommand(subnet_createCmd)
}
