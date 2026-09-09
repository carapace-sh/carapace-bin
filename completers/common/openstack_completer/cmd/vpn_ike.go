package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ikeCmd = &cobra.Command{
	Use:   "ike",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ikeCmd).Standalone()

	vpnCmd.AddCommand(vpn_ikeCmd)
}
