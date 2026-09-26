package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Update the project's environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().Bool("active", false, "Sync dependencies to the active virtual environment")
	syncCmd.Flags().Bool("all-extras", false, "Include all optional dependencies")
	syncCmd.Flags().Bool("all-groups", false, "Include dependencies from all dependency groups")
	syncCmd.Flags().Bool("all-packages", false, "Sync all packages in the workspace")
	syncCmd.Flags().Bool("binary", false, "")
	syncCmd.Flags().Bool("build", false, "")
	syncCmd.Flags().Bool("build-isolation", false, "")
	syncCmd.Flags().Bool("check", false, "Check if the Python environment is synchronized with the project")
	syncCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	syncCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	syncCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	syncCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	syncCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	syncCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	syncCmd.Flags().Bool("dev", false, "Include the development dependency group [env: UV_DEV=]")
	syncCmd.Flags().Bool("dry-run", false, "Perform a dry run, without writing the lockfile or modifying the project environment")
	syncCmd.Flags().Bool("editable", false, "Install any non-editable dependencies, including the project and any workspace members, as editable")
	syncCmd.Flags().Bool("exact", false, "Perform an exact sync, removing extraneous packages")
	syncCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	syncCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	syncCmd.Flags().StringSlice("extra", nil, "Include optional dependencies from the specified extra name")
	syncCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	syncCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	syncCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	syncCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	syncCmd.Flags().Bool("frozen", false, "Sync without updating the `uv.lock` file [env: UV_FROZEN=]")
	syncCmd.Flags().StringSlice("group", nil, "Include dependencies from the specified dependency group")
	syncCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	syncCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	syncCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	syncCmd.Flags().Bool("inexact", false, "Do not remove extraneous packages present in the environment")
	syncCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	syncCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	syncCmd.Flags().Bool("locked", false, "Assert that the `uv.lock` will remain unchanged [env: UV_LOCKED=]")
	syncCmd.Flags().Bool("no-active", false, "Prefer project's virtual environment over an active environment")
	syncCmd.Flags().Bool("no-all-extras", false, "")
	syncCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	syncCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	syncCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	syncCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	syncCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	syncCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	syncCmd.Flags().Bool("no-check", false, "")
	syncCmd.Flags().Bool("no-compile", false, "")
	syncCmd.Flags().Bool("no-compile-bytecode", false, "")
	syncCmd.Flags().Bool("no-default-groups", false, "Ignore the default dependency groups")
	syncCmd.Flags().Bool("no-dev", false, "Disable the development dependency group [env: UV_NO_DEV=]")
	syncCmd.Flags().Bool("no-editable", false, "Install any editable dependencies, including the project and any workspace members, as non-editable [env: UV_NO_EDITABLE=]")
	syncCmd.Flags().StringSlice("no-editable-package", nil, "Install the specified editable packages as non-editable")
	syncCmd.Flags().Bool("no-exact", false, "Do not remove extraneous packages present in the environment")
	syncCmd.Flags().StringSlice("no-extra", nil, "Exclude the specified optional dependencies, if `--all-extras` is supplied")
	syncCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	syncCmd.Flags().StringSlice("no-group", nil, "Disable the specified dependency group [env: `UV_NO_GROUP`=]")
	syncCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	syncCmd.Flags().Bool("no-install-local", false, "Do not install local path dependencies [env: UV_NO_INSTALL_LOCAL=]")
	syncCmd.Flags().StringSlice("no-install-package", nil, "Do not install the given package(s)")
	syncCmd.Flags().Bool("no-install-project", false, "Do not install the current project [env: UV_NO_INSTALL_PROJECT=]")
	syncCmd.Flags().Bool("no-install-workspace", false, "Do not install any workspace members, including the root project [env: UV_NO_INSTALL_WORKSPACE=]")
	syncCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	syncCmd.Flags().Bool("no-refresh", false, "")
	syncCmd.Flags().Bool("no-reinstall", false, "")
	syncCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	syncCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	syncCmd.Flags().Bool("no-upgrade", false, "")
	syncCmd.Flags().Bool("only-dev", false, "Only include the development dependency group")
	syncCmd.Flags().StringSlice("only-group", nil, "Only include dependencies from the specified dependency group")
	syncCmd.Flags().Bool("only-install-local", false, "Only install local path dependencies")
	syncCmd.Flags().StringSlice("only-install-package", nil, "Only install the given package(s)")
	syncCmd.Flags().Bool("only-install-project", false, "Only install the current project")
	syncCmd.Flags().Bool("only-install-workspace", false, "Only install workspace members, including the root project")
	syncCmd.Flags().String("output-format", "text", "Select the output format")
	syncCmd.Flags().StringSlice("package", nil, "Sync for specific packages in the workspace")
	syncCmd.Flags().Bool("pre", false, "")
	syncCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	syncCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	syncCmd.Flags().StringP("python", "p", "", "The Python interpreter to use for the project environment.")
	syncCmd.Flags().String("python-platform", "", "The platform for which requirements should be installed")
	syncCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	syncCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	syncCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	syncCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	syncCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	syncCmd.Flags().String("script", "", "Sync the environment for a Python script, rather than the current project")
	syncCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	syncCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	syncCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	syncCmd.Flag("binary").Hidden = true
	syncCmd.Flag("build").Hidden = true
	syncCmd.Flag("build-isolation").Hidden = true
	syncCmd.Flag("compile").Hidden = true
	syncCmd.Flag("config-settings").Hidden = true
	syncCmd.Flag("config-settings-package").Hidden = true
	syncCmd.Flag("dev").Hidden = true
	syncCmd.Flag("editable").Hidden = true
	syncCmd.Flag("exact").Hidden = true
	syncCmd.Flag("force-reinstall").Hidden = true
	syncCmd.Flag("no-active").Hidden = true
	syncCmd.Flag("no-all-extras").Hidden = true
	syncCmd.Flag("no-check").Hidden = true
	syncCmd.Flag("no-compile").Hidden = true
	syncCmd.Flag("no-compile-bytecode").Hidden = true
	syncCmd.Flag("no-exact").Hidden = true
	syncCmd.Flag("no-frozen").Hidden = true
	syncCmd.Flag("no-locked").Hidden = true
	syncCmd.Flag("no-refresh").Hidden = true
	syncCmd.Flag("no-reinstall").Hidden = true
	syncCmd.Flag("no-upgrade").Hidden = true
	syncCmd.Flag("only-install-local").Hidden = true
	syncCmd.Flag("only-install-package").Hidden = true
	syncCmd.Flag("only-install-project").Hidden = true
	syncCmd.Flag("only-install-workspace").Hidden = true
	syncCmd.Flag("pre").Hidden = true
	rootCmd.AddCommand(syncCmd)
	carapace.Gen(syncCmd).FlagCompletion(carapace.ActionMap{
		"extra":            uv.ActionExtras(),
		"fork-strategy":    uv.ActionForkStrategies(),
		"group":            uv.ActionDependencyGroups(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"output-format": carapace.ActionValuesDescribed(
			"text", "Display the result in a human-readable format",
			"json", "Display the result in JSON format",
		),
		"prerelease":      uv.ActionPreReleases(),
		"python":          uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"python-platform": uv.ActionPythonPlatforms(),
		"resolution":      uv.ActionResolutions(),
		"script":          carapace.ActionFiles(),
	})
}
