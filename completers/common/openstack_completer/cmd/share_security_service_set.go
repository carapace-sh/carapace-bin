package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_security_service_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set security service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_security_service_setCmd).Standalone()

	share_security_service_setCmd.Flags().String("default-ad-site", "", "Default AD site.")
	share_security_service_setCmd.Flags().String("description", "", "Set security service description.")
	share_security_service_setCmd.Flags().String("dns-ip", "", "Set DNS IP address used inside project's network.")
	share_security_service_setCmd.Flags().String("domain", "", "Set security service domain.")
	share_security_service_setCmd.Flags().String("name", "", "Set security service name.")
	share_security_service_setCmd.Flags().String("ou", "", "Set security service OU (Organizational Unit).")
	share_security_service_setCmd.Flags().String("password", "", "Set password used by user.")
	share_security_service_setCmd.Flags().String("server", "", "Set security service IP address or hostname.")
	share_security_service_setCmd.Flags().String("user", "", "Set security service user or group used by project.")
	share_security_serviceCmd.AddCommand(share_security_service_setCmd)
}
