package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var port_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset port properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(port_unsetCmd).Standalone()

	port_unsetCmd.Flags().Bool("all-tag", false, "Clear all tags associated with the port")
	port_unsetCmd.Flags().String("allowed-address", "", "Desired allowed-address pair which should be removed from this port: ip-address=<ip-address>[,mac-address=<mac-address>] (repeat option to unset multiple allowed-address pairs)")
	port_unsetCmd.Flags().String("binding-profile", "", "Desired key which should be removed from binding:profile (repeat option to unset multiple binding:profile keys)")
	port_unsetCmd.Flags().Bool("data-plane-status", false, "Clear existing data plane status information")
	port_unsetCmd.Flags().Bool("device", false, "Clear device ID for the port.")
	port_unsetCmd.Flags().Bool("device-owner", false, "Clear device owner for the port.")
	port_unsetCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	port_unsetCmd.Flags().String("fixed-ip", "", "Desired IP and/or subnet which should be removed from this port (name or ID): subnet=<subnet>,ip-address=<ip-address> (repeat option to unset multiple fixed IP addresses)")
	port_unsetCmd.Flags().Bool("hints", false, "Clear hints for the port")
	port_unsetCmd.Flags().Bool("host", false, "Clear host binding for the port")
	port_unsetCmd.Flags().Bool("numa-policy", false, "Clear existing NUMA affinity policy")
	port_unsetCmd.Flags().Bool("pvlan-community", false, "Clear PVLAN community name for the port.")
	port_unsetCmd.Flags().Bool("qos-policy", false, "Remove the QoS policy attached to the port")
	port_unsetCmd.Flags().String("security-group", "", "Security group which should be removed this port (name or ID) (repeat option to unset multiple security groups)")
	port_unsetCmd.Flags().String("tag", "", "Tag to be removed from the port (repeat option to remove multiple tags)")
	portCmd.AddCommand(port_unsetCmd)
}
