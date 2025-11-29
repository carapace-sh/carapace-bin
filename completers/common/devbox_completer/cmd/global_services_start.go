package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_services_startCmd = &cobra.Command{
	Use:   "start [service]...",
	Short: "Start service. If no service is specified, starts all services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_services_startCmd).Standalone()

	global_servicesCmd.AddCommand(global_services_startCmd)

	// TODO positional completion
}
