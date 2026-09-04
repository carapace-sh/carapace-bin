package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_service_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete VPN service(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_service_deleteCmd).Standalone()

	vpn_serviceCmd.AddCommand(vpn_service_deleteCmd)
}
