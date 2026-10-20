package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deps_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install project dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deps_installCmd).Standalone()

	deps_installCmd.Flags().BoolP("dry-run", "n", false, "Show which providers would run, without running them")
	deps_installCmd.Flags().Bool("explain", false, "Show why a provider is fresh or stale, and exit non-zero unless it is fresh (requires a PROVIDER argument)")
	deps_installCmd.Flags().BoolP("force", "f", false, "Run providers even when their outputs are up to date")
	deps_installCmd.Flags().BoolP("help", "h", false, "Print help")
	deps_installCmd.Flags().Bool("list", false, "List configured providers with their status, sources, and outputs")
	deps_installCmd.Flags().Bool("monorepo", false, "Also run providers from every config root in `[monorepo].config_roots`")
	deps_installCmd.Flags().StringSlice("only", nil, "Run only these providers")
	deps_installCmd.Flags().StringSlice("skip", nil, "Skip these providers")
	depsCmd.AddCommand(deps_installCmd)
}
