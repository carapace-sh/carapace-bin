package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serve_setConfigCmd = &cobra.Command{
	Use:   "set-config",
	Short: "Define service configuration from a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serve_setConfigCmd).Standalone()

	serve_setConfigCmd.Flags().Bool("all", false, "apply config to all services")
	serve_setConfigCmd.Flags().String("service", "", "apply config to a particular service")
	serveCmd.AddCommand(serve_setConfigCmd)

	carapace.Gen(serve_setConfigCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
