package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_providers_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "[experimental] Restart shared servers with their current configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_providers_restartCmd).Standalone()

	daemons_providers_restartCmd.Flags().BoolP("help", "h", false, "Print help")
	daemons_providersCmd.AddCommand(daemons_providers_restartCmd)
}
