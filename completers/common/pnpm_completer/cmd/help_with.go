package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_withCmd = &cobra.Command{
	Use:   "with",
	Short: "Runs pnpm at a specific version (or the currently running one) for a single invocation, ignoring the \"packageManager\" and \"devEngines.packageManager\" fields of the project's manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_withCmd).Standalone()

	helpCmd.AddCommand(help_withCmd)
}
