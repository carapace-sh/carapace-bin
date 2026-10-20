// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
)

var bootstrap_packages_brew_tapCmd = &cobra.Command{
	Use:   "tap",
	Short: "Add a Homebrew tap URL to the global `[bootstrap.brew.taps]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_brew_tapCmd).Standalone()

	bootstrap_packages_brew_tapCmd.Flags().BoolP("dry-run", "n", false, "Show the config change without writing it")
	bootstrap_packages_brew_tapCmd.Flags().String("file", "", "Write to this config file or directory")
	bootstrap_packages_brew_tapCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packages_brew_tapCmd.Flags().BoolP("local", "l", false, "Write to the local config instead of the global config")
	bootstrap_packages_brew_tapCmd.Flags().StringP("path", "p", "", "Write to this config file or directory")
	bootstrap_packages_brewCmd.AddCommand(bootstrap_packages_brew_tapCmd)

	carapace.Gen(bootstrap_packages_brew_tapCmd).PositionalCompletion(
		carapace.ActionValues(),
		http.ActionUrls(),
	)
}
