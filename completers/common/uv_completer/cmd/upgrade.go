package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:    "upgrade",
	Short:  "Upgrade a dependency in the project",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	upgradeCmd.Flags().StringSlice("exclude", nil, "Exclude the named package from upgrades")
	upgradeCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	upgradeCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	upgradeCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	upgradeCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	upgradeCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	upgradeCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	upgradeCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	rootCmd.AddCommand(upgradeCmd)
	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
	})
}
