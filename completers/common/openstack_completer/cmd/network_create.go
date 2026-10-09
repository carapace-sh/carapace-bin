package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_createCmd).Standalone()

	network_createCmd.Flags().String("availability-zone-hint", "", "Availability Zone in which to create this network (Network Availability Zone extension required, repeat option to set multiple availability zones)")
	network_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	network_createCmd.Flags().Bool("default", false, "Specify if this network should be used as the default external network")
	network_createCmd.Flags().String("description", "", "Set network description")
	network_createCmd.Flags().Bool("disable", false, "Disable network")
	network_createCmd.Flags().Bool("disable-port-security", false, "Disable port security by default for ports created on this network")
	network_createCmd.Flags().String("dns-domain", "", "Set DNS domain for this network (requires DNS integration extension)")
	network_createCmd.Flags().Bool("enable", false, "Enable network (default)")
	network_createCmd.Flags().Bool("enable-port-security", false, "Enable port security by default for ports created on this network (default)")
	network_createCmd.Flags().Bool("external", false, "The network has an external routing facility that is not managed by Neutron and can be used.")
	network_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	network_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	network_createCmd.Flags().Bool("internal", false, "Opposite of '--external' (default)")
	network_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	network_createCmd.Flags().String("mtu", "", "Set network mtu")
	network_createCmd.Flags().Bool("no-default", false, "Do not use the network as the default external network (default)")
	network_createCmd.Flags().Bool("no-pvlan", false, "Disable Private VLAN for the network (PVLAN extension required)")
	network_createCmd.Flags().Bool("no-qinq-vlan", false, "Disable VLAN QinQ (S-Tag ethtype 0x8a88) for the network")
	network_createCmd.Flags().Bool("no-share", false, "Do not share the network between projects")
	network_createCmd.Flags().Bool("no-tag", false, "No tags associated with the network")
	network_createCmd.Flags().Bool("no-transparent-vlan", false, "Do not make the network VLAN transparent")
	network_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	network_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	network_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	network_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	network_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	network_createCmd.Flags().String("provider-network-type", "", "The physical mechanism by which the virtual network is implemented.")
	network_createCmd.Flags().String("provider-physical-network", "", "Name of the physical network over which the virtual network is implemented")
	network_createCmd.Flags().String("provider-segment", "", "VLAN ID for VLAN networks or Tunnel ID for GENEVE/GRE/VXLAN networks")
	network_createCmd.Flags().Bool("pvlan", false, "Enable Private VLAN for the network (PVLAN extension required)")
	network_createCmd.Flags().Bool("qinq-vlan", false, "Enable VLAN QinQ (S-Tag ethtype 0x8a88) for the network")
	network_createCmd.Flags().String("qos-policy", "", "QoS policy to attach to this network (name or ID)")
	network_createCmd.Flags().Bool("share", false, "Share the network between projects")
	network_createCmd.Flags().String("tag", "", "Tag to be added to the network (repeat option to set multiple tags)")
	network_createCmd.Flags().Bool("transparent-vlan", false, "Make the network VLAN transparent")
	network_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	networkCmd.AddCommand(network_createCmd)
}
