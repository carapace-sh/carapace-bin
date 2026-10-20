// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_services_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Uninstall a user-scope service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_services_removeCmd).Standalone()

	bootstrap_services_removeCmd.Flags().BoolP("dry-run", "n", false, "Show what would be removed without removing it")
	bootstrap_services_removeCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_servicesCmd.AddCommand(bootstrap_services_removeCmd)
}
