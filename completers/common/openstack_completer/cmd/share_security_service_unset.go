package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_security_service_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset security service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_security_service_unsetCmd).Standalone()

	share_security_service_unsetCmd.Flags().Bool("default-ad-site", false, "Default AD site.")
	share_security_service_unsetCmd.Flags().Bool("description", false, "Unset security service description.")
	share_security_service_unsetCmd.Flags().Bool("dns-ip", false, "Unset DNS IP address used inside project's network.")
	share_security_service_unsetCmd.Flags().Bool("domain", false, "Unset security service domain.")
	share_security_service_unsetCmd.Flags().Bool("name", false, "Unset security service name.")
	share_security_service_unsetCmd.Flags().Bool("ou", false, "Unset security service OU (Organizational Unit).")
	share_security_service_unsetCmd.Flags().Bool("password", false, "Unset password used by user.")
	share_security_service_unsetCmd.Flags().Bool("server", false, "Unset security service IP address or hostname.")
	share_security_service_unsetCmd.Flags().Bool("user", false, "Unset security service user or group used by project.")
	share_security_serviceCmd.AddCommand(share_security_service_unsetCmd)
}
