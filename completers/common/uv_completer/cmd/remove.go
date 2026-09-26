package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove dependencies from the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("active", false, "Prefer the active virtual environment over the project's virtual environment")
	removeCmd.Flags().Bool("binary", false, "")
	removeCmd.Flags().Bool("build", false, "")
	removeCmd.Flags().Bool("build-isolation", false, "")
	removeCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	removeCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	removeCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	removeCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	removeCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	removeCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	removeCmd.Flags().Bool("dev", false, "Remove the packages from the development dependency group [env: UV_DEV=]")
	removeCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	removeCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	removeCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	removeCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	removeCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	removeCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	removeCmd.Flags().Bool("frozen", false, "Remove dependencies without re-locking the project [env: UV_FROZEN=]")
	removeCmd.Flags().String("group", "", "Remove the packages from the specified dependency group")
	removeCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	removeCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	removeCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	removeCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	removeCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	removeCmd.Flags().Bool("locked", false, "Assert that the `uv.lock` will remain unchanged [env: UV_LOCKED=]")
	removeCmd.Flags().Bool("no-active", false, "Prefer project's virtual environment over an active environment")
	removeCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	removeCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	removeCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	removeCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	removeCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	removeCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	removeCmd.Flags().Bool("no-compile", false, "")
	removeCmd.Flags().Bool("no-compile-bytecode", false, "")
	removeCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	removeCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	removeCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	removeCmd.Flags().Bool("no-refresh", false, "")
	removeCmd.Flags().Bool("no-reinstall", false, "")
	removeCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	removeCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	removeCmd.Flags().Bool("no-sync", false, "Avoid syncing the virtual environment after re-locking the project [env: UV_NO_SYNC=]")
	removeCmd.Flags().Bool("no-upgrade", false, "")
	removeCmd.Flags().String("optional", "", "Remove the packages from the project's optional dependencies for the specified extra")
	removeCmd.Flags().String("package", "", "Remove the dependencies from a specific package in the workspace")
	removeCmd.Flags().Bool("pre", false, "")
	removeCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	removeCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	removeCmd.Flags().StringP("python", "p", "", "The Python interpreter to use for resolving and syncing.")
	removeCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	removeCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	removeCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	removeCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	removeCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	removeCmd.Flags().String("script", "", "Remove the dependency from the specified Python script, rather than from a project")
	removeCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	removeCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	removeCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	removeCmd.Flag("binary").Hidden = true
	removeCmd.Flag("build").Hidden = true
	removeCmd.Flag("build-isolation").Hidden = true
	removeCmd.Flag("compile").Hidden = true
	removeCmd.Flag("config-settings").Hidden = true
	removeCmd.Flag("force-reinstall").Hidden = true
	removeCmd.Flag("no-active").Hidden = true
	removeCmd.Flag("no-compile").Hidden = true
	removeCmd.Flag("no-compile-bytecode").Hidden = true
	removeCmd.Flag("no-frozen").Hidden = true
	removeCmd.Flag("no-locked").Hidden = true
	removeCmd.Flag("no-refresh").Hidden = true
	removeCmd.Flag("no-reinstall").Hidden = true
	removeCmd.Flag("no-upgrade").Hidden = true
	removeCmd.Flag("pre").Hidden = true
	rootCmd.AddCommand(removeCmd)
	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"fork-strategy":    uv.ActionForkStrategies(),
		"group":            uv.ActionDependencyGroups(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"prerelease":       uv.ActionPreReleases(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"resolution":       uv.ActionResolutions(),
		"script":           carapace.ActionFiles(),
	})
	carapace.Gen(removeCmd).PositionalAnyCompletion(uv.ActionProjectDependencies())
}
