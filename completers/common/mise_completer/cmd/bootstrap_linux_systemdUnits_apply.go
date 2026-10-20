// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_linux_systemdUnits_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Install and start the configured systemd user units",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_linux_systemdUnits_applyCmd).Standalone()

	bootstrap_linux_systemdUnits_applyCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrap_linux_systemdUnits_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_linux_systemdUnits_applyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_linux_systemdUnitsCmd.AddCommand(bootstrap_linux_systemdUnits_applyCmd)
}
