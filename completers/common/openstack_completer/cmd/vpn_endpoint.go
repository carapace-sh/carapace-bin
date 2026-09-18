package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_endpointCmd = &cobra.Command{
	Use:   "endpoint",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_endpointCmd).Standalone()

	vpnCmd.AddCommand(vpn_endpointCmd)
}
