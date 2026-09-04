package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set network properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_setCmd).Standalone()

	network_setCmd.Flags().Bool("default", false, "Set the network as the default external network")
	network_setCmd.Flags().String("description", "", "Set network description")
	network_setCmd.Flags().Bool("disable", false, "Disable network")
	network_setCmd.Flags().Bool("disable-port-security", false, "Disable port security by default for ports created on this network")
	network_setCmd.Flags().String("dns-domain", "", "Set DNS domain for this network (requires DNS integration extension)")
	network_setCmd.Flags().Bool("enable", false, "Enable network")
	network_setCmd.Flags().Bool("enable-port-security", false, "Enable port security by default for ports created on this network")
	network_setCmd.Flags().Bool("external", false, "The network has an external routing facility that is not managed by Neutron and can be used.")
	network_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_setCmd.Flags().Bool("internal", false, "Opposite of '--external'")
	network_setCmd.Flags().String("mtu", "", "Set network mtu")
	network_setCmd.Flags().String("name", "", "Set network name")
	network_setCmd.Flags().Bool("no-default", false, "Do not use the network as the default external network")
	network_setCmd.Flags().Bool("no-pvlan", false, "Disable Private VLAN for the network (Default).")
	network_setCmd.Flags().Bool("no-qos-policy", false, "Remove the QoS policy attached to this network")
	network_setCmd.Flags().Bool("no-share", false, "Do not share the network between projects")
	network_setCmd.Flags().Bool("no-tag", false, "Clear tags associated with the network.")
	network_setCmd.Flags().String("provider-network-type", "", "The physical mechanism by which the virtual network is implemented.")
	network_setCmd.Flags().String("provider-physical-network", "", "Name of the physical network over which the virtual network is implemented")
	network_setCmd.Flags().String("provider-segment", "", "VLAN ID for VLAN networks or Tunnel ID for GENEVE/GRE/VXLAN networks")
	network_setCmd.Flags().Bool("pvlan", false, "Enable Private VLAN for the network.")
	network_setCmd.Flags().String("qos-policy", "", "QoS policy to attach to this network (name or ID)")
	network_setCmd.Flags().Bool("share", false, "Share the network between projects")
	network_setCmd.Flags().String("tag", "", "Tag to be added to the network (repeat option to set multiple tags)")
	networkCmd.AddCommand(network_setCmd)
}
