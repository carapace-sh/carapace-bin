package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_providers_startCmd = &cobra.Command{
	Use:   "start",
	Short: "[experimental] Start shared servers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_providers_startCmd).Standalone()

	daemons_providers_startCmd.Flags().BoolP("help", "h", false, "Print help")
	daemons_providersCmd.AddCommand(daemons_providers_startCmd)
}
