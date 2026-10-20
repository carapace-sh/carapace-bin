// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_macosCmd = &cobra.Command{
	Use:   "macos",
	Short: "Manage macOS preferences and LaunchAgents from `[bootstrap.macos]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_macosCmd).Standalone()

	bootstrap_macosCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_macosCmd)
}
