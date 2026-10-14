package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_policyCmd = &cobra.Command{
	Use:   "policy",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_policyCmd).Standalone()

	vpn_ipsecCmd.AddCommand(vpn_ipsec_policyCmd)
}
