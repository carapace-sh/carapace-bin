// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_launchdCmd = &cobra.Command{
	Use:    "launchd",
	Short:  "Manage macOS LaunchAgents (use `mise bootstrap macos launchd-agents`)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_launchdCmd).Standalone()

	bootstrap_launchdCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_launchdCmd)
}
