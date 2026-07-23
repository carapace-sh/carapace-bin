package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var integration_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall an integration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(integration_uninstallCmd).Standalone()

	integrationCmd.AddCommand(integration_uninstallCmd)

	carapace.Gen(integration_uninstallCmd).PositionalCompletion(
		carapace.ActionValues("pi", "omp", "claude", "codex", "copilot", "devin", "droid", "kimi", "opencode", "kilo", "hermes", "qodercli", "cursor"),
	)
}
