package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_applyCmd = &cobra.Command{
	Use:     "apply",
	Short:   "Install or remove packages to match `[bootstrap.packages]`",
	Aliases: []string{"i", "install"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_applyCmd).Standalone()

	bootstrap_packages_applyCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrap_packages_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packages_applyCmd.Flags().StringP("manager", "m", "", "Only apply packages for this built-in or plugin manager")
	bootstrap_packages_applyCmd.Flags().Bool("update", false, "Refresh package manager metadata first, where the manager supports it")
	bootstrap_packages_applyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_packagesCmd.AddCommand(bootstrap_packages_applyCmd)
}
