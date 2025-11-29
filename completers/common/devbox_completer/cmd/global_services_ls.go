package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_services_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_services_lsCmd).Standalone()

	global_servicesCmd.AddCommand(global_services_lsCmd)
}
