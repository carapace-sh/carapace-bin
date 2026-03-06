package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_services_pcportCmd = &cobra.Command{
	Use:   "pcport",
	Short: "Display the port that process-compose is running on",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_services_pcportCmd).Standalone()

	global_servicesCmd.AddCommand(global_services_pcportCmd)
}
