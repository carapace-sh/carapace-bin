package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_service_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Enable/Disable share service (Admin only).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_service_setCmd).Standalone()

	share_service_setCmd.Flags().Bool("disable", false, "Disable share service")
	share_service_setCmd.Flags().String("disable-reason", "", "Reason for disabling the service (should be used with --disable option)")
	share_service_setCmd.Flags().Bool("enable", false, "Enable share service")
	share_serviceCmd.AddCommand(share_service_setCmd)
}
