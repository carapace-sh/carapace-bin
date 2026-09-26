package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Display the project's dependency tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(treeCmd).Standalone()

	treeCmd.Flags().Bool("all-groups", false, "Include dependencies from all dependency groups")
	treeCmd.Flags().Bool("binary", false, "")
	treeCmd.Flags().Bool("build", false, "")
	treeCmd.Flags().Bool("build-isolation", false, "")
	treeCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	treeCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	treeCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	treeCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	treeCmd.Flags().StringP("depth", "d", "255", "Maximum display depth of the dependency tree")
	treeCmd.Flags().Bool("dev", false, "Include the development dependency group [env: UV_DEV=]")
	treeCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	treeCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	treeCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	treeCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	treeCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	treeCmd.Flags().String("format", "text", "The format in which to display the dependency graph")
	treeCmd.Flags().Bool("frozen", false, "Display the requirements without locking the project [env: UV_FROZEN=]")
	treeCmd.Flags().StringSlice("group", nil, "Include dependencies from the specified dependency group")
	treeCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	treeCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	treeCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	treeCmd.Flags().Bool("invert", false, "Show the reverse dependencies for the given package. This flag will invert the tree and display the packages that depend on the given package")
	treeCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	treeCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	treeCmd.Flags().Bool("locked", false, "Assert that the `uv.lock` will remain unchanged [env: UV_LOCKED=]")
	treeCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	treeCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	treeCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	treeCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	treeCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	treeCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	treeCmd.Flags().Bool("no-dedupe", false, "Do not de-duplicate repeated dependencies. Usually, when a package has already displayed its dependencies, further occurrences will not re-display its dependencies, and will include a (*) to indicate it has already been shown. This flag will cause those duplicates to be repeated")
	treeCmd.Flags().Bool("no-default-groups", false, "Ignore the default dependency groups")
	treeCmd.Flags().Bool("no-dev", false, "Disable the development dependency group [env: UV_NO_DEV=]")
	treeCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	treeCmd.Flags().StringSlice("no-group", nil, "Disable the specified dependency group [env: `UV_NO_GROUP`=]")
	treeCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	treeCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	treeCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	treeCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	treeCmd.Flags().Bool("no-upgrade", false, "")
	treeCmd.Flags().Bool("only-dev", false, "Only include the development dependency group")
	treeCmd.Flags().StringSlice("only-group", nil, "Only include dependencies from the specified dependency group")
	treeCmd.Flags().Bool("outdated", false, "Show the latest available version of each package in the tree")
	treeCmd.Flags().StringSlice("package", nil, "Display only the specified packages")
	treeCmd.Flags().Bool("pre", false, "")
	treeCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	treeCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	treeCmd.Flags().StringSlice("prune", nil, "Prune the given package from the display of the dependency tree")
	treeCmd.Flags().StringP("python", "p", "", "The Python interpreter to use for locking and filtering.")
	treeCmd.Flags().String("python-platform", "", "The platform to use when filtering the tree")
	treeCmd.Flags().String("python-version", "", "The Python version to use when filtering the tree")
	treeCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	treeCmd.Flags().Bool("reverse", false, "Show the reverse dependencies for the given package. This flag will invert the tree and display the packages that depend on the given package")
	treeCmd.Flags().String("script", "", "Show the dependency tree the specified PEP 723 Python script, rather than the current project")
	treeCmd.Flags().Bool("show-sizes", false, "Show compressed wheel sizes for packages in the tree")
	treeCmd.Flags().Bool("universal", false, "Show a platform-independent dependency tree")
	treeCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	treeCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	treeCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	treeCmd.Flag("binary").Hidden = true
	treeCmd.Flag("build").Hidden = true
	treeCmd.Flag("build-isolation").Hidden = true
	treeCmd.Flag("config-settings").Hidden = true
	treeCmd.Flag("dev").Hidden = true
	treeCmd.Flag("no-frozen").Hidden = true
	treeCmd.Flag("no-locked").Hidden = true
	treeCmd.Flag("no-upgrade").Hidden = true
	treeCmd.Flag("pre").Hidden = true
	treeCmd.Flag("reverse").Hidden = true
	rootCmd.AddCommand(treeCmd)
	carapace.Gen(treeCmd).FlagCompletion(carapace.ActionMap{
		"fork-strategy": uv.ActionForkStrategies(),
		"format": carapace.ActionValuesDescribed(
			"text", "Display the dependency graph as a human-readable tree",
			"json", "Display the dependency graph as JSON",
		),
		"group":            uv.ActionDependencyGroups(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"prerelease":       uv.ActionPreReleases(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"python-platform":  uv.ActionPythonPlatforms(),
		"resolution":       uv.ActionResolutions(),
		"script":           carapace.ActionFiles(),
	})
}
