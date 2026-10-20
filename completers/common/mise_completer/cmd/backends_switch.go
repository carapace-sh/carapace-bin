// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var backends_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch tools to the backend the registry now installs them from",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backends_switchCmd).Standalone()

	backends_switchCmd.Flags().BoolP("dry-run", "n", false, "Show what would switch without changing the lockfile or installs")
	backends_switchCmd.Flags().BoolP("global", "g", false, "Switch tools in the global config's lockfile instead of the project's")
	backends_switchCmd.Flags().BoolP("help", "h", false, "Print help")
	backendsCmd.AddCommand(backends_switchCmd)

	carapace.Gen(backends_switchCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
