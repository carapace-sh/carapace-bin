// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:    "current",
	Short:  "Show the active tool versions",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(currentCmd).Standalone()

	currentCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(currentCmd)

	carapace.Gen(currentCmd).PositionalCompletion(
		mise.ActionPlugins(),
	)
}
