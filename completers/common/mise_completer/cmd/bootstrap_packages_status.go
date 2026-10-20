// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show the state of packages from `[bootstrap.packages]`",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_statusCmd).Standalone()

	bootstrap_packages_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packages_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_packages_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if any package differs from the config (all are still listed)")
	bootstrap_packagesCmd.AddCommand(bootstrap_packages_statusCmd)
}
