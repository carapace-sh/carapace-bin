// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Record installed Homebrew formulae in `[bootstrap.packages]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_importCmd).Standalone()

	bootstrap_packages_importCmd.Flags().Bool("all", false, "Import every linked formula, including dependencies")
	bootstrap_packages_importCmd.Flags().BoolP("dry-run", "n", false, "Show the entries that would be written without writing them")
	bootstrap_packages_importCmd.Flags().StringP("env", "e", "", "Write to the config file for this environment (mise.<ENV>.toml)")
	bootstrap_packages_importCmd.Flags().String("file", "", "Write to this config file or directory")
	bootstrap_packages_importCmd.Flags().BoolP("global", "g", false, "Write to the global config (~/.config/mise/config.toml)")
	bootstrap_packages_importCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packages_importCmd.Flags().StringP("manager", "m", "brew", "Package manager to import from")
	bootstrap_packages_importCmd.Flags().StringP("path", "p", "", "Write to this config file or directory")
	bootstrap_packagesCmd.AddCommand(bootstrap_packages_importCmd)

	carapace.Gen(bootstrap_packages_importCmd).FlagCompletion(carapace.ActionMap{
		"manager": carapace.ActionValues("brew"),
	})
}
