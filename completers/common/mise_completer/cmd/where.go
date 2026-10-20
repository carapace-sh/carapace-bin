// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Show the install directory of a tool version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whereCmd).Standalone()

	whereCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(whereCmd)

	carapace.Gen(whereCmd).PositionalCompletion(
		mise.ActionToolVersions(),
	)
}
