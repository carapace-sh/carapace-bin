package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_provider_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete service provider(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_provider_deleteCmd).Standalone()

	service_providerCmd.AddCommand(service_provider_deleteCmd)
}
