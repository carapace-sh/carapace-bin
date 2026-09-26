package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pip"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add dependencies to the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().Bool("active", false, "Prefer the active virtual environment over the project's virtual environment")
	addCmd.Flags().Bool("binary", false, "")
	addCmd.Flags().String("bounds", "", "The kind of version specifier to use when adding dependencies")
	addCmd.Flags().String("branch", "", "Branch to use when adding a dependency from Git")
	addCmd.Flags().Bool("build", false, "")
	addCmd.Flags().Bool("build-isolation", false, "")
	addCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	addCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	addCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	addCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	addCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	addCmd.Flags().StringSlice("constraint", nil, "Constrain versions using the given requirements files")
	addCmd.Flags().StringSliceP("constraints", "c", nil, "Constrain versions using the given requirements files")
	addCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	addCmd.Flags().Bool("dev", false, "Add the requirements to the development dependency group [env: UV_DEV=]")
	addCmd.Flags().Bool("editable", false, "Add the requirements as editable")
	addCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	addCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	addCmd.Flags().StringSlice("extra", nil, "Extras to enable for the dependency")
	addCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	addCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	addCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	addCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	addCmd.Flags().Bool("frozen", false, "Add dependencies without re-locking the project [env: UV_FROZEN=]")
	addCmd.Flags().String("group", "", "Add the requirements to the specified dependency group")
	addCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	addCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	addCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	addCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	addCmd.Flags().Bool("lfs", false, "Whether to use Git LFS when adding a dependency from Git")
	addCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	addCmd.Flags().Bool("locked", false, "Assert that the `uv.lock` will remain unchanged [env: UV_LOCKED=]")
	addCmd.Flags().StringP("marker", "m", "", "Apply this marker to all added packages")
	addCmd.Flags().Bool("no-active", false, "Prefer project's virtual environment over an active environment")
	addCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	addCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	addCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	addCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	addCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	addCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	addCmd.Flags().Bool("no-compile", false, "")
	addCmd.Flags().Bool("no-compile-bytecode", false, "")
	addCmd.Flags().Bool("no-editable", false, "Don't add the requirements as editable [env: UV_NO_EDITABLE=]")
	addCmd.Flags().StringSlice("no-editable-package", nil, "Don't add the specified requirements as editable")
	addCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	addCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	addCmd.Flags().Bool("no-install-local", false, "Do not install local path dependencies [env: UV_NO_INSTALL_LOCAL=]")
	addCmd.Flags().StringSlice("no-install-package", nil, "Do not install the given package(s)")
	addCmd.Flags().Bool("no-install-project", false, "Do not install the current project [env: UV_NO_INSTALL_PROJECT=]")
	addCmd.Flags().Bool("no-install-workspace", false, "Do not install any workspace members, including the current project [env: UV_NO_INSTALL_WORKSPACE=]")
	addCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	addCmd.Flags().Bool("no-refresh", false, "")
	addCmd.Flags().Bool("no-reinstall", false, "")
	addCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	addCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	addCmd.Flags().Bool("no-sync", false, "Avoid syncing the virtual environment [env: UV_NO_SYNC=]")
	addCmd.Flags().Bool("no-upgrade", false, "")
	addCmd.Flags().Bool("no-workspace", false, "Don't add the dependency as a workspace member")
	addCmd.Flags().Bool("only-install-local", false, "Only install local path dependencies")
	addCmd.Flags().StringSlice("only-install-package", nil, "Only install the given package(s)")
	addCmd.Flags().Bool("only-install-project", false, "Only install the current project")
	addCmd.Flags().Bool("only-install-workspace", false, "Only install workspace members, including the current project")
	addCmd.Flags().String("optional", "", "Add the requirements to the package's optional dependencies for the specified extra")
	addCmd.Flags().String("package", "", "Add the dependency to a specific package in the workspace")
	addCmd.Flags().Bool("pre", false, "")
	addCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	addCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	addCmd.Flags().StringP("python", "p", "", "The Python interpreter to use for resolving and syncing.")
	addCmd.Flags().Bool("raw", false, "Add a dependency as provided")
	addCmd.Flags().Bool("raw-sources", false, "Add a dependency as provided")
	addCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	addCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	addCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	addCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	addCmd.Flags().StringSlice("requirement", nil, "Add the packages listed in the given files")
	addCmd.Flags().StringSliceP("requirements", "r", nil, "Add the packages listed in the given files")
	addCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	addCmd.Flags().String("rev", "", "Commit to use when adding a dependency from Git")
	addCmd.Flags().String("script", "", "Add the dependency to the specified Python script, rather than to a project")
	addCmd.Flags().String("tag", "", "Tag to use when adding a dependency from Git")
	addCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	addCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	addCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	addCmd.Flags().Bool("workspace", false, "Add the dependency as a workspace member")
	addCmd.Flag("binary").Hidden = true
	addCmd.Flag("build").Hidden = true
	addCmd.Flag("build-isolation").Hidden = true
	addCmd.Flag("compile").Hidden = true
	addCmd.Flag("config-settings").Hidden = true
	addCmd.Flag("config-settings-package").Hidden = true
	addCmd.Flag("constraint").Hidden = true
	addCmd.Flag("force-reinstall").Hidden = true
	addCmd.Flag("no-active").Hidden = true
	addCmd.Flag("no-compile").Hidden = true
	addCmd.Flag("no-compile-bytecode").Hidden = true
	addCmd.Flag("no-editable").Hidden = true
	addCmd.Flag("no-editable-package").Hidden = true
	addCmd.Flag("no-frozen").Hidden = true
	addCmd.Flag("no-locked").Hidden = true
	addCmd.Flag("no-refresh").Hidden = true
	addCmd.Flag("no-reinstall").Hidden = true
	addCmd.Flag("no-upgrade").Hidden = true
	addCmd.Flag("only-install-local").Hidden = true
	addCmd.Flag("only-install-package").Hidden = true
	addCmd.Flag("only-install-project").Hidden = true
	addCmd.Flag("only-install-workspace").Hidden = true
	addCmd.Flag("pre").Hidden = true
	addCmd.Flag("raw-sources").Hidden = true
	addCmd.Flag("requirement").Hidden = true
	rootCmd.AddCommand(addCmd)
	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"bounds": carapace.ActionValuesDescribed(
			"lower", "Only a lower bound, e.g., `>=1.2.3`",
			"major", "Allow the same major version, similar to the semver caret, e.g., `>=1.2.3, <2.0.0`",
			"minor", "Allow the same minor version, similar to the semver tilde, e.g., `>=1.2.3, <1.3.0`",
			"exact", "Pin the exact version, e.g., `==1.2.3`",
		),
		"constraint":       carapace.ActionFiles(),
		"constraints":      carapace.ActionFiles(),
		"extra":            uv.ActionExtras(),
		"fork-strategy":    uv.ActionForkStrategies(),
		"group":            uv.ActionDependencyGroups(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"prerelease":       uv.ActionPreReleases(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"requirement":      carapace.ActionFiles(),
		"requirements":     carapace.ActionFiles(),
		"resolution":       uv.ActionResolutions(),
		"script":           carapace.ActionFiles(),
	})
	carapace.Gen(addCmd).PositionalAnyCompletion(pip.ActionPackageSearch())
}
