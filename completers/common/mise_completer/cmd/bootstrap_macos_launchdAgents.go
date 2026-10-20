// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_macos_launchdAgentsCmd = &cobra.Command{
	Use:     "launchd-agents",
	Short:   "Manage macOS LaunchAgents from `[bootstrap.macos.launchd.agents]`",
	Aliases: []string{"launchd"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_macos_launchdAgentsCmd).Standalone()

	bootstrap_macos_launchdAgentsCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_macosCmd.AddCommand(bootstrap_macos_launchdAgentsCmd)
}
