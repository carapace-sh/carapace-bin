package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Update the project's lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockCmd).Standalone()

	lockCmd.Flags().Bool("binary", false, "")
	lockCmd.Flags().Bool("build", false, "")
	lockCmd.Flags().Bool("build-isolation", false, "")
	lockCmd.Flags().Bool("check", false, "Check if the lockfile is up-to-date")
	lockCmd.Flags().Bool("check-exists", false, "Assert that a `uv.lock` exists without checking if it is up-to-date [env: UV_FROZEN=]")
	lockCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	lockCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	lockCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	lockCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	lockCmd.Flags().Bool("dry-run", false, "Perform a dry run, without writing the lockfile")
	lockCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	lockCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	lockCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	lockCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	lockCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	lockCmd.Flags().Bool("frozen", false, "Equivalent to `--check-exists`")
	lockCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	lockCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	lockCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	lockCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	lockCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	lockCmd.Flags().Bool("locked", false, "Check if the lockfile is up-to-date [env: UV_LOCKED=]")
	lockCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	lockCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	lockCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	lockCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	lockCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	lockCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	lockCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	lockCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	lockCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	lockCmd.Flags().Bool("no-refresh", false, "")
	lockCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	lockCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	lockCmd.Flags().Bool("no-upgrade", false, "")
	lockCmd.Flags().Bool("pre", false, "")
	lockCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	lockCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	lockCmd.Flags().StringP("python", "p", "", "The Python interpreter to use during resolution.")
	lockCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	lockCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	lockCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	lockCmd.Flags().String("script", "", "Lock the specified Python script, rather than the current project")
	lockCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	lockCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	lockCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	lockCmd.Flag("binary").Hidden = true
	lockCmd.Flag("build").Hidden = true
	lockCmd.Flag("build-isolation").Hidden = true
	lockCmd.Flag("config-settings").Hidden = true
	lockCmd.Flag("config-settings-package").Hidden = true
	lockCmd.Flag("frozen").Hidden = true
	lockCmd.Flag("locked").Hidden = true
	lockCmd.Flag("no-frozen").Hidden = true
	lockCmd.Flag("no-locked").Hidden = true
	lockCmd.Flag("no-refresh").Hidden = true
	lockCmd.Flag("no-upgrade").Hidden = true
	lockCmd.Flag("pre").Hidden = true
	rootCmd.AddCommand(lockCmd)
	carapace.Gen(lockCmd).FlagCompletion(carapace.ActionMap{
		"fork-strategy":    uv.ActionForkStrategies(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"prerelease":       uv.ActionPreReleases(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"resolution":       uv.ActionResolutions(),
		"script":           carapace.ActionFiles(),
	})
}
