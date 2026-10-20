// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade outdated tools",
	Aliases: []string{"up"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().String("before", "", "Only upgrade to versions released before this date or older than this duration")
	upgradeCmd.Flags().BoolP("bump", "b", false, "Upgrade past the configured range to the newest release, and update the config to match")
	upgradeCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	upgradeCmd.Flags().Bool("dry-run-code", false, "Like --dry-run, but exit with code 1 if there are outdated tools")
	upgradeCmd.Flags().StringSliceP("exclude", "x", nil, "Tools to exclude from upgrading, such as go python")
	upgradeCmd.Flags().Bool("global", false, "Only upgrade tools defined in the global config")
	upgradeCmd.Flags().BoolP("help", "h", false, "Print help")
	upgradeCmd.Flags().Bool("inactive", false, "Also upgrade installed tools that the current config does not request")
	upgradeCmd.Flags().BoolP("interactive", "i", false, "Choose which tools to upgrade from a multiselect menu")
	upgradeCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel (default: the `jobs` setting)")
	upgradeCmd.Flags().BoolS("l", "l", false, "Deprecated shorthand for --bump")
	upgradeCmd.Flags().Bool("local", false, "Only upgrade tools defined in project config files")
	upgradeCmd.Flags().String("minimum-release-age", "", "Only upgrade to versions released before this date or older than this duration")
	upgradeCmd.Flags().Bool("monorepo", false, "Placeholder for future monorepo upgrades; `mise upgrade --monorepo` is not implemented yet.")
	upgradeCmd.Flags().Bool("no-prune", false, "Do not uninstall the versions that were upgraded away from")
	upgradeCmd.Flags().Bool("prune", false, "Immediately uninstall the versions that were upgraded away from")
	upgradeCmd.Flags().Bool("raw", false, "Connect the install commands' stdin, stdout, and stderr to the terminal")
	upgradeCmd.Flag("before").Hidden = true
	upgradeCmd.Flag("l").Hidden = true
	upgradeCmd.Flag("monorepo").Hidden = true
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"exclude": mise.ActionInstalledTools(),
	})

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		mise.ActionInstalledToolVersions(),
	)
}
