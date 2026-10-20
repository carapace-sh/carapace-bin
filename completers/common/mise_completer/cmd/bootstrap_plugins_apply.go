// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_plugins_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Install package manager plugins declared in `[bootstrap.plugins]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_plugins_applyCmd).Standalone()

	bootstrap_plugins_applyCmd.Flags().BoolP("dry-run", "n", false, "Show which plugins would be installed without installing them")
	bootstrap_plugins_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_pluginsCmd.AddCommand(bootstrap_plugins_applyCmd)
}
