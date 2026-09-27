package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_providerCmd).Standalone()

	serviceCmd.AddCommand(service_providerCmd)
}
