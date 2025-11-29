package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_services_stopCmd = &cobra.Command{
	Use:     "stop [service]...",
	Short:   "Stop one or more services in the current project. If no service is specified, stops all services in the current project.",
	Aliases: []string{"down"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_services_stopCmd).Standalone()

	global_services_stopCmd.Flags().Bool("all-projects", false, "stop all running services across all your projects.")
	global_servicesCmd.AddCommand(global_services_stopCmd)

	// TODO positional completion
}
