// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade installed packages from `[bootstrap.packages]`",
	Aliases: []string{"up"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_upgradeCmd).Standalone()

	bootstrap_packages_upgradeCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrap_packages_upgradeCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packages_upgradeCmd.Flags().StringP("manager", "m", "", "Only upgrade packages for this built-in or plugin manager")
	bootstrap_packages_upgradeCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_packagesCmd.AddCommand(bootstrap_packages_upgradeCmd)
}
