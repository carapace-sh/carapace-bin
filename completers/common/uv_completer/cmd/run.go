package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command or script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("active", false, "Prefer the active virtual environment over the project's virtual environment")
	runCmd.Flags().Bool("all-extras", false, "Include all optional dependencies")
	runCmd.Flags().Bool("all-groups", false, "Include dependencies from all dependency groups")
	runCmd.Flags().Bool("all-packages", false, "Run the command with all workspace members installed")
	runCmd.Flags().Bool("binary", false, "")
	runCmd.Flags().Bool("build", false, "")
	runCmd.Flags().Bool("build-isolation", false, "")
	runCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	runCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	runCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	runCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	runCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	runCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	runCmd.Flags().Bool("dev", false, "Include the development dependency group [env: UV_DEV=]")
	runCmd.Flags().Bool("editable", false, "Install any non-editable dependencies, including the project and any workspace members, as editable")
	runCmd.Flags().StringSlice("env-file", nil, "Load environment variables from a `.env` file")
	runCmd.Flags().Bool("exact", false, "Perform an exact sync, removing extraneous packages")
	runCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	runCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	runCmd.Flags().StringSlice("extra", nil, "Include optional dependencies from the specified extra name")
	runCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	runCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	runCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	runCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	runCmd.Flags().Bool("frozen", false, "Run without updating the `uv.lock` file [env: UV_FROZEN=]")
	runCmd.Flags().StringSlice("group", nil, "Include dependencies from the specified dependency group")
	runCmd.Flags().Bool("gui-script", false, "Run the given path as a Python GUI script")
	runCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	runCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	runCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	runCmd.Flags().Bool("inexact", false, "Do not remove extraneous packages present in the environment")
	runCmd.Flags().Bool("isolated", false, "Run the command in an isolated virtual environment [env: UV_ISOLATED=]")
	runCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	runCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	runCmd.Flags().Bool("locked", false, "Assert that the `uv.lock` will remain unchanged [env: UV_LOCKED=]")
	runCmd.Flags().String("max-recursion-depth", "", "Number of times that `uv run` will allow recursive invocations")
	runCmd.Flags().BoolP("module", "m", false, "Run a Python module")
	runCmd.Flags().Bool("no-active", false, "Prefer project's virtual environment over an active environment")
	runCmd.Flags().Bool("no-all-extras", false, "")
	runCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	runCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	runCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	runCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	runCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	runCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	runCmd.Flags().Bool("no-compile", false, "")
	runCmd.Flags().Bool("no-compile-bytecode", false, "")
	runCmd.Flags().Bool("no-default-groups", false, "Ignore the default dependency groups")
	runCmd.Flags().Bool("no-dev", false, "Disable the development dependency group [env: UV_NO_DEV=]")
	runCmd.Flags().Bool("no-editable", false, "Install any editable dependencies, including the project and any workspace members, as non-editable [env: UV_NO_EDITABLE=]")
	runCmd.Flags().StringSlice("no-editable-package", nil, "Install the specified editable packages as non-editable")
	runCmd.Flags().Bool("no-env-file", false, "Avoid reading environment variables from a `.env` file [env: UV_NO_ENV_FILE=]")
	runCmd.Flags().Bool("no-exact", false, "Do not remove extraneous packages present in the environment")
	runCmd.Flags().StringSlice("no-extra", nil, "Exclude the specified optional dependencies, if `--all-extras` is supplied")
	runCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	runCmd.Flags().StringSlice("no-group", nil, "Disable the specified dependency group [env: `UV_NO_GROUP`=]")
	runCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	runCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	runCmd.Flags().Bool("no-project", false, "Avoid discovering the project or workspace")
	runCmd.Flags().Bool("no-refresh", false, "")
	runCmd.Flags().Bool("no-reinstall", false, "")
	runCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	runCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	runCmd.Flags().Bool("no-sync", false, "Avoid syncing the virtual environment [env: UV_NO_SYNC=]")
	runCmd.Flags().Bool("no-upgrade", false, "")
	runCmd.Flags().Bool("no_workspace", false, "Avoid discovering the project or workspace")
	runCmd.Flags().Bool("only-dev", false, "Only include the development dependency group")
	runCmd.Flags().StringSlice("only-group", nil, "Only include dependencies from the specified dependency group")
	runCmd.Flags().String("package", "", "Run the command in a specific package in the workspace")
	runCmd.Flags().Bool("pre", false, "")
	runCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	runCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	runCmd.Flags().StringP("python", "p", "", "The Python interpreter to use for the run environment.")
	runCmd.Flags().String("python-platform", "", "The platform for which requirements should be installed")
	runCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	runCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	runCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	runCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	runCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	runCmd.Flags().BoolP("script", "s", false, "Run the given path as a Python script")
	runCmd.Flags().Bool("show-resolution", false, "Whether to show resolver and installer output from any environment modifications [env: UV_SHOW_RESOLUTION=]")
	runCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	runCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	runCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	runCmd.Flags().StringSliceP("with", "w", nil, "Run with the given packages installed")
	runCmd.Flags().StringSlice("with-editable", nil, "Run with the given packages installed in editable mode")
	runCmd.Flags().StringSlice("with-requirements", nil, "Run with the packages listed in the given files")
	runCmd.Flag("binary").Hidden = true
	runCmd.Flag("build").Hidden = true
	runCmd.Flag("build-isolation").Hidden = true
	runCmd.Flag("compile").Hidden = true
	runCmd.Flag("config-settings").Hidden = true
	runCmd.Flag("config-settings-package").Hidden = true
	runCmd.Flag("dev").Hidden = true
	runCmd.Flag("editable").Hidden = true
	runCmd.Flag("force-reinstall").Hidden = true
	runCmd.Flag("inexact").Hidden = true
	runCmd.Flag("max-recursion-depth").Hidden = true
	runCmd.Flag("no-active").Hidden = true
	runCmd.Flag("no-all-extras").Hidden = true
	runCmd.Flag("no-compile").Hidden = true
	runCmd.Flag("no-compile-bytecode").Hidden = true
	runCmd.Flag("no-exact").Hidden = true
	runCmd.Flag("no-frozen").Hidden = true
	runCmd.Flag("no-locked").Hidden = true
	runCmd.Flag("no-refresh").Hidden = true
	runCmd.Flag("no-reinstall").Hidden = true
	runCmd.Flag("no-upgrade").Hidden = true
	runCmd.Flag("no_workspace").Hidden = true
	runCmd.Flag("pre").Hidden = true
	runCmd.Flag("show-resolution").Hidden = true
	rootCmd.AddCommand(runCmd)
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"env-file":          carapace.ActionFiles(),
		"extra":             uv.ActionExtras(),
		"fork-strategy":     uv.ActionForkStrategies(),
		"group":             uv.ActionDependencyGroups(),
		"index-strategy":    uv.ActionIndexStrategies(),
		"keyring-provider":  uv.ActionKeyringProviders(),
		"link-mode":         uv.ActionLinkModes(),
		"prerelease":        uv.ActionPreReleases(),
		"python":            uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"python-platform":   uv.ActionPythonPlatforms(),
		"resolution":        uv.ActionResolutions(),
		"with-editable":     carapace.ActionDirectories(),
		"with-requirements": carapace.ActionFiles(),
	})
	carapace.Gen(runCmd).PositionalCompletion(carapace.Batch(uv.ActionScripts(), carapace.ActionExecutables()).ToA())
}
