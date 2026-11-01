package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serve_getConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "Get service configuration to save to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serve_getConfigCmd).Standalone()

	serve_getConfigCmd.Flags().Bool("all", false, "read config from all services")
	serve_getConfigCmd.Flags().String("service", "", "read config from a particular service")
	serveCmd.AddCommand(serve_getConfigCmd)

	carapace.Gen(serve_getConfigCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
