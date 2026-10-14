package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var port_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(port_createCmd).Standalone()

	port_createCmd.Flags().String("allowed-address", "", "Add allowed-address pair associated with this port: ip-address=<ip-address>[,mac-address=<mac-address>] (repeat option to set multiple allowed-address pairs)")
	port_createCmd.Flags().String("binding-profile", "", "Custom data to be passed as binding:profile.")
	port_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	port_createCmd.Flags().String("description", "", "Description of this port")
	port_createCmd.Flags().String("device", "", "Port device ID")
	port_createCmd.Flags().String("device-owner", "", "Device owner of this port.")
	port_createCmd.Flags().String("device-profile", "", "Port device profile")
	port_createCmd.Flags().Bool("disable", false, "Disable port")
	port_createCmd.Flags().Bool("disable-port-security", false, "Disable port security for this port")
	port_createCmd.Flags().Bool("disable-uplink-status-propagation", false, "Disable uplink status propagation")
	port_createCmd.Flags().String("dns-domain", "", "Set DNS domain to this port (requires dns_domain extension for ports)")
	port_createCmd.Flags().String("dns-name", "", "Set DNS name for this port (requires DNS integration extension)")
	port_createCmd.Flags().Bool("enable", false, "Enable port (default)")
	port_createCmd.Flags().Bool("enable-port-security", false, "Enable port security for this port (default)")
	port_createCmd.Flags().Bool("enable-uplink-status-propagation", false, "Enable uplink status propagation (default)")
	port_createCmd.Flags().String("extra-dhcp-option", "", "Extra DHCP options to be assigned to this port: name=<name>[,value=<value>,ip-version={4,6}] (repeat option to set multiple extra DHCP options)")
	port_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	port_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	port_createCmd.Flags().String("fixed-ip", "", "Desired IP and/or subnet for this port (name or ID): subnet=<subnet>,ip-address=<ip-address> (repeat option to set multiple fixed IP addresses)")
	port_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	port_createCmd.Flags().String("hardware-offload-type", "", "Hardware offload type this port will request when attached to the network backend")
	port_createCmd.Flags().String("hint", "", "Port hints as ALIAS=VALUE or as JSON.")
	port_createCmd.Flags().String("host", "", "Allocate port on host <host-id> (ID only)")
	port_createCmd.Flags().String("mac-address", "", "MAC address of this port")
	port_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	port_createCmd.Flags().String("network", "", "Network this port belongs to (name or ID)")
	port_createCmd.Flags().Bool("no-fixed-ip", false, "No IP or subnet set for this port")
	port_createCmd.Flags().Bool("no-security-group", false, "Associate no security groups with this port")
	port_createCmd.Flags().Bool("no-tag", false, "No tags associated with the port")
	port_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	port_createCmd.Flags().Bool("not-trusted", false, "Set port to be not trusted.")
	port_createCmd.Flags().Bool("numa-policy-legacy", false, "NUMA affinity policy using legacy mode to schedule this port")
	port_createCmd.Flags().Bool("numa-policy-preferred", false, "NUMA affinity policy preferred to schedule this port")
	port_createCmd.Flags().Bool("numa-policy-required", false, "NUMA affinity policy required to schedule this port")
	port_createCmd.Flags().Bool("numa-policy-socket", false, "NUMA affinity policy socket to schedule this port")
	port_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	port_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	port_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	port_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	port_createCmd.Flags().String("pvlan-community", "", "Set PVLAN community name for this port.")
	port_createCmd.Flags().String("pvlan-type", "", "Set Private VLAN type for this port.")
	port_createCmd.Flags().String("qos-policy", "", "Attach QoS policy to this port (name or ID)")
	port_createCmd.Flags().String("security-group", "", "Security group to associate with this port (name or ID) (repeat option to set multiple security groups)")
	port_createCmd.Flags().String("tag", "", "Tag to be added to the port (repeat option to set multiple tags)")
	port_createCmd.Flags().Bool("trusted", false, "Set port to be trusted.")
	port_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	port_createCmd.Flags().String("vnic-type", "", "VNIC type for this port (direct | direct-physical | macvtap | normal | baremetal | virtio-forwarder | vdpa | remote-managed) (default: normal)")
	port_createCmd.MarkFlagRequired("network")
	portCmd.AddCommand(port_createCmd)
}
