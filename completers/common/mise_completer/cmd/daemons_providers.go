// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "[experimental] Manage shared servers from global `[daemon_providers]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_providersCmd).Standalone()

	daemons_providersCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonsCmd.AddCommand(daemons_providersCmd)
}
