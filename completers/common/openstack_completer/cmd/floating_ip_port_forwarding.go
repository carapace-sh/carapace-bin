package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_port_forwardingCmd = &cobra.Command{
	Use:   "forwarding",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_port_forwardingCmd).Standalone()

	floating_ip_portCmd.AddCommand(floating_ip_port_forwardingCmd)
}
