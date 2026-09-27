package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgpvpn_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete BGP VPN resource(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgpvpn_deleteCmd).Standalone()

	bgpvpnCmd.AddCommand(bgpvpn_deleteCmd)
}
