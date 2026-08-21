package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_port_forwarding_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set floating IP Port Forwarding Properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_port_forwarding_setCmd).Standalone()

	floating_ip_port_forwarding_setCmd.Flags().String("description", "", "Text to describe/contextualize the use of the port forwarding configuration")
	floating_ip_port_forwarding_setCmd.Flags().String("external-protocol-port", "", "The TCP/UDP/other protocol port number of the port forwarding's floating IP address")
	floating_ip_port_forwarding_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	floating_ip_port_forwarding_setCmd.Flags().String("internal-ip-address", "", "The fixed IPv4 address of the network port associated to the floating IP port forwarding")
	floating_ip_port_forwarding_setCmd.Flags().String("internal-protocol-port", "", "The TCP/UDP/other protocol port number of the network port fixed IPv4 address associated to the floating IP port forwarding")
	floating_ip_port_forwarding_setCmd.Flags().String("port", "", "The ID of the network port associated to the floating IP port forwarding")
	floating_ip_port_forwarding_setCmd.Flags().String("protocol", "", "The IP protocol used in the floating IP port forwarding")
	floating_ip_port_forwardingCmd.AddCommand(floating_ip_port_forwarding_setCmd)
}
