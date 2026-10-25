package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var integration_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install an integration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(integration_installCmd).Standalone()

	integrationCmd.AddCommand(integration_installCmd)

	carapace.Gen(integration_installCmd).PositionalCompletion(
		carapace.ActionValues("antigravity-cli", "claude", "codex", "copilot", "cursor", "devin", "droid", "grok", "hermes", "kilo", "kimi", "mastracode", "omp", "opencode", "pi", "qodercli"),
	)
}
