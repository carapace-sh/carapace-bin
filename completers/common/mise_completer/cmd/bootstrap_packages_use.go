// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_useCmd = &cobra.Command{
	Use:     "use",
	Short:   "Add packages to `[bootstrap.packages]` and install them",
	Aliases: []string{"u"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_useCmd).Standalone()

	bootstrap_packages_useCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without writing config or installing")
	bootstrap_packages_useCmd.Flags().StringP("env", "e", "", "Write to the config file for this environment (mise.<ENV>.toml)")
	bootstrap_packages_useCmd.Flags().String("file", "", "Write to this config file or directory")
	bootstrap_packages_useCmd.Flags().BoolP("global", "g", false, "Write to the global config (~/.config/mise/config.toml) instead of the local one")
	bootstrap_packages_useCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packages_useCmd.Flags().Bool("no-install", false, "Write the package declarations without checking or installing packages")
	bootstrap_packages_useCmd.Flags().StringP("path", "p", "", "Write to this config file or directory")
	bootstrap_packages_useCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_packagesCmd.AddCommand(bootstrap_packages_useCmd)
}
