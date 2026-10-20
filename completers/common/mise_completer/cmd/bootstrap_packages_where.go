// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Print the `opt` path of an installed Homebrew formula",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_whereCmd).Standalone()

	bootstrap_packages_whereCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packagesCmd.AddCommand(bootstrap_packages_whereCmd)
}
