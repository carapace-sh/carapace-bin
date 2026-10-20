// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemons_providers_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "[experimental] Stop shared servers (data is kept)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemons_providers_stopCmd).Standalone()

	daemons_providers_stopCmd.Flags().BoolP("help", "h", false, "Print help")
	daemons_providersCmd.AddCommand(daemons_providers_stopCmd)
}
