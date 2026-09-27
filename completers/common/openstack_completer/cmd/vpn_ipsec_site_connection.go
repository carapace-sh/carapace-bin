package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_site_connectionCmd = &cobra.Command{
	Use:   "connection",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_site_connectionCmd).Standalone()

	vpn_ipsec_siteCmd.AddCommand(vpn_ipsec_site_connectionCmd)
}
