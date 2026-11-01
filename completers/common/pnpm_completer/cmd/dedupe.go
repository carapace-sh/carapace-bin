package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dedupeCmd = &cobra.Command{
	Use:   "dedupe",
	Short: "Deduplicate packages in the lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dedupeCmd).Standalone()

	dedupeCmd.Flags().Bool("check", false, "Check if running dedupe would result in changes without installing packages or editing the lockfile. Exits with a non-zero status code if changes are possible")
	dedupeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	dedupeCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in `.pnpmfile.cjs`, including the pnpmfiles of config dependencies")
	dedupeCmd.Flags().Bool("ignore-scripts", false, "Don't run lifecycle scripts of the project or its dependencies. Packages are still installed; only their build scripts are skipped, and the install won't fail because of it")
	dedupeCmd.Flags().Bool("lockfile-only", false, "Only update `pnpm-lock.yaml`. Don't download packages or write `node_modules`")
	dedupeCmd.Flags().Bool("no-ignore-scripts", false, "Run lifecycle scripts even when the configuration disables them")
	dedupeCmd.Flags().Bool("no-offline", false, "Allow network fetches even when the configuration enables offline mode")
	dedupeCmd.Flags().Bool("no-prefer-offline", false, "Don't prefer cached packages even when the configuration enables it")
	dedupeCmd.Flags().Bool("offline", false, "Fail on a cache miss instead of fetching from the registry, using only packages already in the store")
	dedupeCmd.Flags().Bool("prefer-offline", false, "Prefer packages already in the cache over the network, even past their freshness window")
	rootCmd.AddCommand(dedupeCmd)
}
