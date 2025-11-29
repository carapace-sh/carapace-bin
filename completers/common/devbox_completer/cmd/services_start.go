package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var services_startCmd = &cobra.Command{
	Use:   "start [service]...",
	Short: "Start service. If no service is specified, starts all services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_startCmd).Standalone()

	servicesCmd.AddCommand(services_startCmd)

	// TODO positional completion
}
