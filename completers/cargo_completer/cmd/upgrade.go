package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Update dependencies as recorded in the local lock file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().Bool("all", false, "[deprecated in favor of `--workspace`]")
	upgradeCmd.Flags().Bool("allow-prerelease", false, "Include prerelease versions when fetching from crates.io (e.g. 0.6.0-alpha')")
	upgradeCmd.Flags().Bool("dry-run", false, "Print changes to be made without making them")
	upgradeCmd.Flags().String("exclude", "", "Crates to exclude and not upgrade")
	upgradeCmd.Flags().BoolP("help", "h", false, "Prints help information")
	upgradeCmd.Flags().String("manifest-path", "", "Path to the manifest to upgrade")
	upgradeCmd.Flags().Bool("offline", false, "Run without accessing the network")
	upgradeCmd.Flags().StringP("package", "p", "", "Package id of the crate to add this dependency to")
	upgradeCmd.Flags().Bool("skip-compatible", false, "Only update a dependency if the new version is semver incompatible")
	upgradeCmd.Flags().Bool("to-lockfile", false, "Upgrade all packages to the version in the lockfile")
	upgradeCmd.Flags().BoolP("version", "V", false, "Prints version information")
	upgradeCmd.Flags().Bool("workspace", false, "Upgrade all packages in the workspace")
	rootCmd.AddCommand(upgradeCmd)

	// TODO exclude / package
	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(".toml"),
	})

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionDependencies(rmCmd, false).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
