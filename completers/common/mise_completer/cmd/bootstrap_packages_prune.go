// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Uninstall packages that no config declares",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_pruneCmd).Standalone()

	bootstrap_packages_pruneCmd.Flags().BoolP("dry-run", "n", false, "Show what would be removed without removing anything")
	bootstrap_packages_pruneCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packages_pruneCmd.Flags().StringP("manager", "m", "brew", "Package manager to prune: `brew`, `brew-cask`, or a package plugin")
	bootstrap_packages_pruneCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_packagesCmd.AddCommand(bootstrap_packages_pruneCmd)
}
