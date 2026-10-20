// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var binPathsCmd = &cobra.Command{
	Use:   "bin-paths",
	Short: "List the bin directories of the active tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(binPathsCmd).Standalone()

	binPathsCmd.Flags().Bool("bin-names", false, "Print executable names instead of bin directories")
	binPathsCmd.Flags().BoolP("help", "h", false, "Print help")
	binPathsCmd.Flags().BoolP("json", "J", false, "Print executables as JSON objects with `name`, `path`, and `symlink` (implies --bin-names)")
	rootCmd.AddCommand(binPathsCmd)

	carapace.Gen(binPathsCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
