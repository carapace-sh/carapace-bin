package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_services_restartCmd = &cobra.Command{
	Use:   "restart [service]...",
	Short: "Restart service. If no service is specified, restarts all services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_services_restartCmd).Standalone()

	global_servicesCmd.AddCommand(global_services_restartCmd)

	// TODO positional completion
}
