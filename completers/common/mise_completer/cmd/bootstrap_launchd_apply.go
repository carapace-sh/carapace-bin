package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_launchd_applyCmd = &cobra.Command{
	Use:    "apply",
	Short:  "Install and load LaunchAgents from `[bootstrap.macos.launchd.agents]`",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_launchd_applyCmd).Standalone()

	bootstrap_launchd_applyCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrap_launchd_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_launchd_applyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_launchdCmd.AddCommand(bootstrap_launchd_applyCmd)
}
