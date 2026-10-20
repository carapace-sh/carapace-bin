// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Delete tool versions that nothing uses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	pruneCmd.Flags().Bool("configs", false, "Only forget tracked, trusted, and ignored config files that no longer exist")
	pruneCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	pruneCmd.Flags().Bool("dry-run-code", false, "Like --dry-run, but exit with code 1 if there are tools to prune")
	pruneCmd.Flags().BoolP("help", "h", false, "Print help")
	pruneCmd.Flags().Bool("monorepo", false, "Placeholder for future monorepo pruning; `mise prune --monorepo` is not implemented yet.")
	pruneCmd.Flags().Bool("tools", false, "Only delete unused tool versions")
	pruneCmd.Flag("monorepo").Hidden = true
	rootCmd.AddCommand(pruneCmd)

	carapace.Gen(pruneCmd).PositionalAnyCompletion(
		mise.ActionInstalledTools(),
	)
}
