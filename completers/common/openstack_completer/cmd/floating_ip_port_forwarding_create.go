package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_port_forwarding_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create floating IP port forwarding",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_port_forwarding_createCmd).Standalone()

	floating_ip_port_forwarding_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	floating_ip_port_forwarding_createCmd.Flags().String("description", "", "Text to describe/contextualize the use of the port forwarding configuration")
	floating_ip_port_forwarding_createCmd.Flags().String("external-protocol-port", "", "The protocol port number of the port forwarding's floating IP address")
	floating_ip_port_forwarding_createCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	floating_ip_port_forwarding_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	floating_ip_port_forwarding_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	floating_ip_port_forwarding_createCmd.Flags().String("internal-ip-address", "", "The fixed IPv4 address of the network port associated to the floating IP port forwarding")
	floating_ip_port_forwarding_createCmd.Flags().String("internal-protocol-port", "", "The protocol port number of the network port fixed IPv4 address associated to the floating IP port forwarding")
	floating_ip_port_forwarding_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	floating_ip_port_forwarding_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	floating_ip_port_forwarding_createCmd.Flags().String("port", "", "The name or ID of the network port associated to the floating IP port forwarding")
	floating_ip_port_forwarding_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	floating_ip_port_forwarding_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	floating_ip_port_forwarding_createCmd.Flags().String("protocol", "", "The protocol used in the floating IP port forwarding, for instance: TCP, UDP")
	floating_ip_port_forwarding_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	floating_ip_port_forwarding_createCmd.MarkFlagRequired("external-protocol-port")
	floating_ip_port_forwarding_createCmd.MarkFlagRequired("internal-ip-address")
	floating_ip_port_forwarding_createCmd.MarkFlagRequired("internal-protocol-port")
	floating_ip_port_forwarding_createCmd.MarkFlagRequired("port")
	floating_ip_port_forwarding_createCmd.MarkFlagRequired("protocol")
	floating_ip_port_forwardingCmd.AddCommand(floating_ip_port_forwarding_createCmd)
}
