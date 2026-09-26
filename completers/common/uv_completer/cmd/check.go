package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Run checks on the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().Bool("all-extras", false, "Include all optional dependencies")
	checkCmd.Flags().Bool("all-groups", false, "Include dependencies from all dependency groups")
	checkCmd.Flags().Bool("all-packages", false, "Check all packages in the workspace")
	checkCmd.Flags().Bool("binary", false, "")
	checkCmd.Flags().Bool("build", false, "")
	checkCmd.Flags().Bool("build-isolation", false, "")
	checkCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	checkCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	checkCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	checkCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	checkCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	checkCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	checkCmd.Flags().Bool("dev", false, "Include the development dependency group [env: UV_DEV=]")
	checkCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	checkCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	checkCmd.Flags().StringSlice("extra", nil, "Include optional dependencies from the specified extra name")
	checkCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	checkCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	checkCmd.Flags().Bool("fix", false, "Apply safe fixes to resolve type-checking errors")
	checkCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	checkCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	checkCmd.Flags().Bool("frozen", false, "Sync without updating the `uv.lock` file [env: UV_FROZEN=]")
	checkCmd.Flags().StringSlice("group", nil, "Include dependencies from the specified dependency group")
	checkCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	checkCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	checkCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	checkCmd.Flags().Bool("isolated", false, "Run checks without mutating project state [env: UV_ISOLATED=]")
	checkCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	checkCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	checkCmd.Flags().Bool("locked", false, "Assert that the `uv.lock` will remain unchanged [env: UV_LOCKED=]")
	checkCmd.Flags().Bool("no-all-extras", false, "")
	checkCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	checkCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	checkCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	checkCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	checkCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	checkCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	checkCmd.Flags().Bool("no-compile", false, "")
	checkCmd.Flags().Bool("no-compile-bytecode", false, "")
	checkCmd.Flags().Bool("no-default-groups", false, "Ignore the default dependency groups")
	checkCmd.Flags().Bool("no-dev", false, "Disable the development dependency group [env: UV_NO_DEV=]")
	checkCmd.Flags().StringSlice("no-extra", nil, "Exclude the specified optional dependencies, if `--all-extras` is supplied")
	checkCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	checkCmd.Flags().StringSlice("no-group", nil, "Disable the specified dependency group [env: `UV_NO_GROUP`=]")
	checkCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	checkCmd.Flags().Bool("no-install-project", false, "Do not install the current project [env: UV_NO_INSTALL_PROJECT=]")
	checkCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	checkCmd.Flags().Bool("no-project", false, "Avoid discovering a project or workspace")
	checkCmd.Flags().Bool("no-refresh", false, "")
	checkCmd.Flags().Bool("no-reinstall", false, "")
	checkCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	checkCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	checkCmd.Flags().Bool("no-sync", false, "Avoid syncing the virtual environment [env: UV_NO_SYNC=]")
	checkCmd.Flags().Bool("no-upgrade", false, "")
	checkCmd.Flags().Bool("only-dev", false, "Only include the development dependency group")
	checkCmd.Flags().StringSlice("only-group", nil, "Only include dependencies from the specified dependency group")
	checkCmd.Flags().StringSlice("package", nil, "Check specific packages in the workspace")
	checkCmd.Flags().Bool("pre", false, "")
	checkCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	checkCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	checkCmd.Flags().StringP("python", "p", "", "The Python interpreter to use for the project environment")
	checkCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	checkCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	checkCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	checkCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	checkCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	checkCmd.Flags().String("script", "", "Run checks for the specified PEP 723 Python script, rather than the current project")
	checkCmd.Flags().Bool("show-command", false, "Display the ty command that will be used for type checking")
	checkCmd.Flags().Bool("show-version", false, "Display the version of ty that will be used for type checking")
	checkCmd.Flags().String("ty-version", "", "The version of ty to use for type checking")
	checkCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	checkCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	checkCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	checkCmd.Flag("binary").Hidden = true
	checkCmd.Flag("build").Hidden = true
	checkCmd.Flag("build-isolation").Hidden = true
	checkCmd.Flag("compile").Hidden = true
	checkCmd.Flag("config-settings").Hidden = true
	checkCmd.Flag("dev").Hidden = true
	checkCmd.Flag("force-reinstall").Hidden = true
	checkCmd.Flag("no-all-extras").Hidden = true
	checkCmd.Flag("no-compile").Hidden = true
	checkCmd.Flag("no-compile-bytecode").Hidden = true
	checkCmd.Flag("no-frozen").Hidden = true
	checkCmd.Flag("no-locked").Hidden = true
	checkCmd.Flag("no-refresh").Hidden = true
	checkCmd.Flag("no-reinstall").Hidden = true
	checkCmd.Flag("no-upgrade").Hidden = true
	checkCmd.Flag("pre").Hidden = true
	checkCmd.Flag("show-command").Hidden = true
	checkCmd.Flag("show-version").Hidden = true
	rootCmd.AddCommand(checkCmd)
	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"extra":            uv.ActionExtras(),
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
}
