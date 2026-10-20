// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Create or refresh lockfile versions, checksums, and download URLs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockCmd).Standalone()

	lockCmd.Flags().String("before", "", "Only lock versions released before a date or at least a duration ago")
	lockCmd.Flags().Bool("bump", false, "Re-resolve version requests against the latest available versions")
	lockCmd.Flags().BoolP("dry-run", "n", false, "Show what would change, without writing anything")
	lockCmd.Flags().BoolP("global", "g", false, "Lock only the global and system config lockfiles")
	lockCmd.Flags().BoolP("help", "h", false, "Print help")
	lockCmd.Flags().StringP("jobs", "j", "", "How many jobs to run in parallel; defaults to the `jobs` setting")
	lockCmd.Flags().Bool("json", false, "Output version changes as JSON")
	lockCmd.Flags().Bool("local", false, "Update mise.local.lock, for tools in mise.local.toml, instead of mise.lock")
	lockCmd.Flags().String("minimum-release-age", "", "Only lock versions released before a date or at least a duration ago")
	lockCmd.Flags().StringSliceP("platform", "p", nil, "Platforms to lock, comma-separated, such as `linux-x64,macos-arm64`")
	lockCmd.Flags().Bool("sidecars", false, "List native dependency sidecars instead of updating lockfiles")
	lockCmd.Flags().Bool("upgrade", false, "Upgrade lockfiles to the latest format")
	lockCmd.Flag("before").Hidden = true
	rootCmd.AddCommand(lockCmd)

	carapace.Gen(lockCmd).PositionalAnyCompletion(
		mise.ActionTools(),
	)
}
