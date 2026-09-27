package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var port_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set port properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(port_setCmd).Standalone()

	port_setCmd.Flags().String("allowed-address", "", "Add allowed-address pair associated with this port: ip-address=<ip-address>[,mac-address=<mac-address>] (repeat option to set multiple allowed-address pairs)")
	port_setCmd.Flags().String("binding-profile", "", "Custom data to be passed as binding:profile.")
	port_setCmd.Flags().String("data-plane-status", "", "Set data plane status of this port (ACTIVE | DOWN).")
	port_setCmd.Flags().String("description", "", "Description of this port")
	port_setCmd.Flags().String("device", "", "Port device ID")
	port_setCmd.Flags().String("device-owner", "", "Device owner of this port.")
	port_setCmd.Flags().Bool("disable", false, "Disable port")
	port_setCmd.Flags().Bool("disable-port-security", false, "Disable port security for this port")
	port_setCmd.Flags().Bool("disable-uplink-status-propagation", false, "Disable uplink status propagation")
	port_setCmd.Flags().String("dns-domain", "", "Set DNS domain to this port (requires dns_domain extension for ports)")
	port_setCmd.Flags().String("dns-name", "", "Set DNS name for this port (requires DNS integration extension)")
	port_setCmd.Flags().Bool("enable", false, "Enable port")
	port_setCmd.Flags().Bool("enable-port-security", false, "Enable port security for this port")
	port_setCmd.Flags().Bool("enable-uplink-status-propagation", false, "Enable uplink status propagation")
	port_setCmd.Flags().String("extra-dhcp-option", "", "Extra DHCP options to be assigned to this port: name=<name>[,value=<value>,ip-version={4,6}] (repeat option to set multiple extra DHCP options)")
	port_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	port_setCmd.Flags().String("fixed-ip", "", "Desired IP and/or subnet for this port (name or ID): subnet=<subnet>,ip-address=<ip-address> (repeat option to set multiple fixed IP addresses)")
	port_setCmd.Flags().String("hint", "", "Port hints as ALIAS=VALUE or as JSON.")
	port_setCmd.Flags().String("host", "", "Allocate port on host <host-id> (ID only)")
	port_setCmd.Flags().String("mac-address", "", "MAC address of this port (admin only)")
	port_setCmd.Flags().String("name", "", "Set port name")
	port_setCmd.Flags().Bool("no-allowed-address", false, "Clear existing allowed-address pairs associated with this port.")
	port_setCmd.Flags().Bool("no-binding-profile", false, "Clear existing information of binding:profile.")
	port_setCmd.Flags().Bool("no-fixed-ip", false, "Clear existing information of fixed IP addresses.")
	port_setCmd.Flags().Bool("no-security-group", false, "Clear existing security groups associated with this port")
	port_setCmd.Flags().Bool("no-tag", false, "Clear tags associated with the port.")
	port_setCmd.Flags().Bool("not-trusted", false, "Set port to be not trusted.")
	port_setCmd.Flags().Bool("numa-policy-legacy", false, "NUMA affinity policy using legacy mode to schedule this port")
	port_setCmd.Flags().Bool("numa-policy-preferred", false, "NUMA affinity policy preferred to schedule this port")
	port_setCmd.Flags().Bool("numa-policy-required", false, "NUMA affinity policy required to schedule this port")
	port_setCmd.Flags().Bool("numa-policy-socket", false, "NUMA affinity policy socket to schedule this port")
	port_setCmd.Flags().String("pvlan-community", "", "Set PVLAN community name for this port.")
	port_setCmd.Flags().String("pvlan-type", "", "Set Private VLAN type for this port.")
	port_setCmd.Flags().String("qos-policy", "", "Attach QoS policy to this port (name or ID)")
	port_setCmd.Flags().String("security-group", "", "Security group to associate with this port (name or ID) (repeat option to set multiple security groups)")
	port_setCmd.Flags().String("tag", "", "Tag to be added to the port (repeat option to set multiple tags)")
	port_setCmd.Flags().Bool("trusted", false, "Set port to be trusted.")
	port_setCmd.Flags().String("vnic-type", "", "VNIC type for this port (direct | direct-physical | macvtap | normal | baremetal | virtio-forwarder | vdpa | remote-managed) (default: normal)")
	portCmd.AddCommand(port_setCmd)
}
