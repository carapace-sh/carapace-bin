// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:     "deps",
	Short:   "[experimental] Manage project dependencies",
	Aliases: []string{"dep", "prepare"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(depsCmd).Standalone()

	depsCmd.Flags().BoolP("dry-run", "n", false, "Show which providers would run, without running them")
	depsCmd.Flags().Bool("explain", false, "Show why a provider is fresh or stale, and exit non-zero unless it is fresh (requires a PROVIDER argument)")
	depsCmd.Flags().BoolP("force", "f", false, "Run providers even when their outputs are up to date")
	depsCmd.Flags().BoolP("help", "h", false, "Print help")
	depsCmd.Flags().Bool("list", false, "List configured providers with their status, sources, and outputs")
	depsCmd.Flags().Bool("monorepo", false, "Also run providers from every config root in `[monorepo].config_roots`")
	depsCmd.Flags().StringSlice("only", nil, "Run only these providers")
	depsCmd.Flags().StringSlice("skip", nil, "Skip these providers")
	rootCmd.AddCommand(depsCmd)
}
