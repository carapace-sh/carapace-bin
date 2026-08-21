package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_serviceCmd).Standalone()

	vpnCmd.AddCommand(vpn_serviceCmd)
}
