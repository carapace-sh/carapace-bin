package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Read or update the project's version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()

	versionCmd.Flags().Bool("active", false, "Prefer the active virtual environment over the project's virtual environment")
	versionCmd.Flags().Bool("binary", false, "")
	versionCmd.Flags().Bool("build", false, "")
	versionCmd.Flags().Bool("build-isolation", false, "")
	versionCmd.Flags().StringSlice("bump", nil, "Update the project version using the given semantics")
	versionCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	versionCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	versionCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	versionCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	versionCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	versionCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	versionCmd.Flags().Bool("dry-run", false, "Don't write a new version to the `pyproject.toml`")
	versionCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	versionCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	versionCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	versionCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	versionCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	versionCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	versionCmd.Flags().Bool("frozen", false, "Update the version without re-locking the project [env: UV_FROZEN=]")
	versionCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	versionCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	versionCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	versionCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	versionCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	versionCmd.Flags().Bool("locked", false, "Assert that the `uv.lock` will remain unchanged [env: UV_LOCKED=]")
	versionCmd.Flags().Bool("no-active", false, "Prefer project's virtual environment over an active environment")
	versionCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	versionCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	versionCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	versionCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	versionCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	versionCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	versionCmd.Flags().Bool("no-compile", false, "")
	versionCmd.Flags().Bool("no-compile-bytecode", false, "")
	versionCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	versionCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	versionCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	versionCmd.Flags().Bool("no-refresh", false, "")
	versionCmd.Flags().Bool("no-reinstall", false, "")
	versionCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	versionCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	versionCmd.Flags().Bool("no-sync", false, "Avoid syncing the virtual environment after re-locking the project [env: UV_NO_SYNC=]")
	versionCmd.Flags().Bool("no-upgrade", false, "")
	versionCmd.Flags().String("output-format", "text", "The format of the output")
	versionCmd.Flags().String("package", "", "Update the version of a specific package in the workspace")
	versionCmd.Flags().Bool("pre", false, "")
	versionCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	versionCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	versionCmd.Flags().StringP("python", "p", "", "The Python interpreter to use for resolving and syncing.")
	versionCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	versionCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	versionCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	versionCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	versionCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	versionCmd.Flags().Bool("short", false, "Only show the version")
	versionCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	versionCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	versionCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	versionCmd.Flag("binary").Hidden = true
	versionCmd.Flag("build").Hidden = true
	versionCmd.Flag("build-isolation").Hidden = true
	versionCmd.Flag("compile").Hidden = true
	versionCmd.Flag("config-settings").Hidden = true
	versionCmd.Flag("force-reinstall").Hidden = true
	versionCmd.Flag("no-active").Hidden = true
	versionCmd.Flag("no-compile").Hidden = true
	versionCmd.Flag("no-compile-bytecode").Hidden = true
	versionCmd.Flag("no-frozen").Hidden = true
	versionCmd.Flag("no-locked").Hidden = true
	versionCmd.Flag("no-refresh").Hidden = true
	versionCmd.Flag("no-reinstall").Hidden = true
	versionCmd.Flag("no-upgrade").Hidden = true
	versionCmd.Flag("pre").Hidden = true
	rootCmd.AddCommand(versionCmd)
	carapace.Gen(versionCmd).FlagCompletion(carapace.ActionMap{
		"bump": carapace.ActionValuesDescribed(
			"major", "Increase the major version (e.g., 1.2.3 => 2.0.0)",
			"minor", "Increase the minor version (e.g., 1.2.3 => 1.3.0)",
			"patch", "Increase the patch version (e.g., 1.2.3 => 1.2.4)",
			"stable", "Move from a pre-release to stable version (e.g., 1.2.3b4.post5.dev6 => 1.2.3)",
			"alpha", "Increase the alpha version (e.g., 1.2.3a4 => 1.2.3a5)",
			"beta", "Increase the beta version (e.g., 1.2.3b4 => 1.2.3b5)",
			"rc", "Increase the rc version (e.g., 1.2.3rc4 => 1.2.3rc5)",
			"post", "Increase the post version (e.g., 1.2.3.post5 => 1.2.3.post6)",
			"dev", "Increase the dev version (e.g., 1.2.3a4.dev6 => 1.2.3.dev7)",
		),
		"fork-strategy":    uv.ActionForkStrategies(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"output-format": carapace.ActionValuesDescribed(
			"text", "Display the version as plain text",
			"json", "Display the version as JSON",
		),
		"prerelease": uv.ActionPreReleases(),
		"python":     uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"resolution": uv.ActionResolutions(),
	})
}
