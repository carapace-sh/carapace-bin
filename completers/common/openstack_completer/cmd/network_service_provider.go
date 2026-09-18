package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_service_providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_service_providerCmd).Standalone()

	network_serviceCmd.AddCommand(network_service_providerCmd)
}
