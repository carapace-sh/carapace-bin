package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_site_connection_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete IPsec site connection(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_site_connection_deleteCmd).Standalone()

	vpn_ipsec_site_connectionCmd.AddCommand(vpn_ipsec_site_connection_deleteCmd)
}
