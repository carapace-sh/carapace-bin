// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_plugins_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show whether declared package manager plugins are installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_plugins_statusCmd).Standalone()

	bootstrap_plugins_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_plugins_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if a declared plugin is not installed")
	bootstrap_pluginsCmd.AddCommand(bootstrap_plugins_statusCmd)
}
