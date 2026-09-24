package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_port_forwarding_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete floating IP port forwarding",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_port_forwarding_deleteCmd).Standalone()

	floating_ip_port_forwardingCmd.AddCommand(floating_ip_port_forwarding_deleteCmd)
}
