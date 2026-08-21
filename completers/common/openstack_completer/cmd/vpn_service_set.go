package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_service_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set VPN service properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_service_setCmd).Standalone()

	vpn_service_setCmd.Flags().String("description", "", "Description for the VPN service")
	vpn_service_setCmd.Flags().Bool("disable", false, "Disable VPN service")
	vpn_service_setCmd.Flags().Bool("enable", false, "Enable VPN service")
	vpn_service_setCmd.Flags().String("flavor", "", "Flavor for the VPN service (name or ID)")
	vpn_service_setCmd.Flags().String("name", "", "Name for the VPN service")
	vpn_service_setCmd.Flags().String("subnet", "", "Local private subnet (name or ID)")
	vpn_serviceCmd.AddCommand(vpn_service_setCmd)
}
