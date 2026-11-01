package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compute_service_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set compute service properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_service_setCmd).Standalone()

	compute_service_setCmd.Flags().Bool("disable", false, "Disable service")
	compute_service_setCmd.Flags().String("disable-reason", "", "Reason for disabling the service (in quotes).")
	compute_service_setCmd.Flags().Bool("down", false, "Force down service.")
	compute_service_setCmd.Flags().Bool("enable", false, "Enable service")
	compute_service_setCmd.Flags().Bool("up", false, "Force up service.")
	compute_serviceCmd.AddCommand(compute_service_setCmd)
}
