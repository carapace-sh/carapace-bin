package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_macVpnCmd = &cobra.Command{
	Use:   "mac-vpn",
	Short: "Manage the VPN configuration on macOS (App Store and Standalone variants)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_macVpnCmd).Standalone()

	configureCmd.AddCommand(configure_macVpnCmd)
}
