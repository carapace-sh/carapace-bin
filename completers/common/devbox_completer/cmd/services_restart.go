package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var services_restartCmd = &cobra.Command{
	Use:   "restart [service]...",
	Short: "Restart service. If no service is specified, restarts all services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_restartCmd).Standalone()

	servicesCmd.AddCommand(services_restartCmd)

	// TODO positional completion
}
