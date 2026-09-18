package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_service_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set volume service properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_service_setCmd).Standalone()

	volume_service_setCmd.Flags().Bool("disable", false, "Disable volume service")
	volume_service_setCmd.Flags().String("disable-reason", "", "Reason for disabling the service (should be used with --disable option)")
	volume_service_setCmd.Flags().Bool("enable", false, "Enable volume service")
	volume_serviceCmd.AddCommand(volume_service_setCmd)
}
