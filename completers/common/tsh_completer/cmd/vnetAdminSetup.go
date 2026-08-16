package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vnetAdminSetupCmd = &cobra.Command{
	Use:    "vnet-admin-setup",
	Short:  "Start the VNet admin subprocess.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vnetAdminSetupCmd).Standalone()

	vnetAdminSetupCmd.Flags().String("socket", "", "Client application service unix socket path.")
	vnetAdminSetupCmd.MarkFlagRequired("socket")
	rootCmd.AddCommand(vnetAdminSetupCmd)
}
