package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set service properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_setCmd).Standalone()

	service_setCmd.Flags().String("description", "", "New service description")
	service_setCmd.Flags().Bool("disable", false, "Disable service")
	service_setCmd.Flags().Bool("enable", false, "Enable service")
	service_setCmd.Flags().String("name", "", "New service name")
	service_setCmd.Flags().String("type", "", "New service type (compute, image, identity, volume, etc)")
	serviceCmd.AddCommand(service_setCmd)
}
