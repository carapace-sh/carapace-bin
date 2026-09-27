package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pip"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var pip_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install packages into an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pip_installCmd).Standalone()

	pip_installCmd.Flags().Bool("all-extras", false, "Include all optional dependencies")
	pip_installCmd.Flags().Bool("break-system-packages", false, "Allow uv to modify an `EXTERNALLY-MANAGED` Python installation")
	pip_installCmd.Flags().Bool("build", false, "")
	pip_installCmd.Flags().StringSlice("build-constraint", nil, "Constrain build dependencies using the given requirements files when building source distributions")
	pip_installCmd.Flags().StringSliceP("build-constraints", "b", nil, "Constrain build dependencies using the given requirements files when building source distributions")
	pip_installCmd.Flags().Bool("build-isolation", false, "")
	pip_installCmd.Flags().Bool("check", false, "Check whether the environment satisfies the requirements without modifying it")
	pip_installCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	pip_installCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	pip_installCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	pip_installCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	pip_installCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	pip_installCmd.Flags().StringSlice("constraint", nil, "Constrain versions using the given requirements files")
	pip_installCmd.Flags().StringSliceP("constraints", "c", nil, "Constrain versions using the given requirements files")
	pip_installCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	pip_installCmd.Flags().Bool("deps", false, "")
	pip_installCmd.Flags().Bool("disable-pip-version-check", false, "")
	pip_installCmd.Flags().Bool("dry-run", false, "Perform a dry run, i.e., don't actually install anything but resolve the dependencies and print the resulting plan")
	pip_installCmd.Flags().StringSliceP("editable", "e", nil, "Install the editable package based on the provided local file path")
	pip_installCmd.Flags().Bool("exact", false, "Perform an exact sync, removing extraneous packages")
	pip_installCmd.Flags().StringSlice("exclude", nil, "Exclude packages from resolution using the given requirements files")
	pip_installCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	pip_installCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	pip_installCmd.Flags().StringSlice("excludes", nil, "Exclude packages from resolution using the given requirements files")
	pip_installCmd.Flags().StringSlice("extra", nil, "Include optional dependencies from the specified extra name; may be provided more than once")
	pip_installCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	pip_installCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	pip_installCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	pip_installCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	pip_installCmd.Flags().StringSlice("group", nil, "Install the specified dependency group from a `pylock.toml` or `pyproject.toml`")
	pip_installCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	pip_installCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	pip_installCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	pip_installCmd.Flags().Bool("inexact", false, "Do not remove extraneous packages present in the environment")
	pip_installCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	pip_installCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	pip_installCmd.Flags().Bool("no-all-extras", false, "")
	pip_installCmd.Flags().StringSlice("no-binary", nil, "Don't install pre-built wheels")
	pip_installCmd.Flags().Bool("no-break-system-packages", false, "")
	pip_installCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	pip_installCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	pip_installCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	pip_installCmd.Flags().Bool("no-compile", false, "")
	pip_installCmd.Flags().Bool("no-compile-bytecode", false, "")
	pip_installCmd.Flags().Bool("no-deps", false, "Ignore package dependencies, instead only installing those packages explicitly listed on the command line or in the requirements files")
	pip_installCmd.Flags().Bool("no-editable", false, "Install any editable dependencies as non-editable [env: UV_NO_EDITABLE=]")
	pip_installCmd.Flags().StringSlice("no-editable-package", nil, "Install the specified editable packages as non-editable")
	pip_installCmd.Flags().Bool("no-exact", false, "Do not remove extraneous packages present in the environment")
	pip_installCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	pip_installCmd.Flags().Bool("no-refresh", false, "")
	pip_installCmd.Flags().Bool("no-reinstall", false, "")
	pip_installCmd.Flags().Bool("no-require-hashes", false, "")
	pip_installCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	pip_installCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	pip_installCmd.Flags().Bool("no-strict", false, "")
	pip_installCmd.Flags().Bool("no-system", false, "")
	pip_installCmd.Flags().Bool("no-upgrade", false, "")
	pip_installCmd.Flags().Bool("no-verify-hashes", false, "Disable validation of hashes in the requirements file")
	pip_installCmd.Flags().StringSlice("only-binary", nil, "Only use pre-built wheels; don't build source distributions")
	pip_installCmd.Flags().String("output-format", "text", "Select the output format")
	pip_installCmd.Flags().StringSlice("override", nil, "Override versions using the given requirements files")
	pip_installCmd.Flags().StringSlice("overrides", nil, "Override versions using the given requirements files")
	pip_installCmd.Flags().Bool("pre", false, "")
	pip_installCmd.Flags().String("prefix", "", "Install packages into `lib`, `bin`, and other top-level folders under the specified directory, as if a virtual environment were present at that location")
	pip_installCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	pip_installCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	pip_installCmd.Flags().StringP("python", "p", "", "The Python interpreter into which packages should be installed.")
	pip_installCmd.Flags().String("python-platform", "", "The platform for which requirements should be installed")
	pip_installCmd.Flags().String("python-version", "", "The minimum Python version that should be supported by the requirements (e.g., `3.7` or `3.7.9`)")
	pip_installCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	pip_installCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	pip_installCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	pip_installCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	pip_installCmd.Flags().Bool("require-hashes", false, "Require a matching hash for each requirement")
	pip_installCmd.Flags().StringSlice("requirement", nil, "Install the packages listed in the given files")
	pip_installCmd.Flags().StringSliceP("requirements", "r", nil, "Install the packages listed in the given files")
	pip_installCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	pip_installCmd.Flags().Bool("strict", false, "Validate the Python environment after completing the installation, to detect packages with missing dependencies or other issues")
	pip_installCmd.Flags().Bool("system", false, "Install packages into the system Python environment")
	pip_installCmd.Flags().StringP("target", "t", "", "Install packages into the specified directory, rather than into the virtual or system Python environment. The packages will be installed at the top-level of the directory")
	pip_installCmd.Flags().String("torch-backend", "", "The backend to use when fetching packages in the PyTorch ecosystem (e.g., `cpu`, `cu126`, or `auto`)")
	pip_installCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	pip_installCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	pip_installCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	pip_installCmd.Flags().Bool("user", false, "")
	pip_installCmd.Flags().Bool("verify-hashes", false, "")

	pip_installCmd.Flag("build").Hidden = true
	pip_installCmd.Flag("build-constraint").Hidden = true
	pip_installCmd.Flag("build-isolation").Hidden = true
	pip_installCmd.Flag("compile").Hidden = true
	pip_installCmd.Flag("config-settings").Hidden = true
	pip_installCmd.Flag("constraint").Hidden = true
	pip_installCmd.Flag("deps").Hidden = true
	pip_installCmd.Flag("disable-pip-version-check").Hidden = true
	pip_installCmd.Flag("exclude").Hidden = true
	pip_installCmd.Flag("force-reinstall").Hidden = true
	pip_installCmd.Flag("inexact").Hidden = true
	pip_installCmd.Flag("no-all-extras").Hidden = true
	pip_installCmd.Flag("no-compile").Hidden = true
	pip_installCmd.Flag("no-compile-bytecode").Hidden = true
	pip_installCmd.Flag("no-exact").Hidden = true
	pip_installCmd.Flag("no-refresh").Hidden = true
	pip_installCmd.Flag("no-reinstall").Hidden = true
	pip_installCmd.Flag("no-require-hashes").Hidden = true
	pip_installCmd.Flag("no-strict").Hidden = true
	pip_installCmd.Flag("no-system").Hidden = true
	pip_installCmd.Flag("no-upgrade").Hidden = true
	pip_installCmd.Flag("override").Hidden = true
	pip_installCmd.Flag("pre").Hidden = true
	pip_installCmd.Flag("requirement").Hidden = true
	pip_installCmd.Flag("verify-hashes").Hidden = true

	pipCmd.AddCommand(pip_installCmd)

	carapace.Gen(pip_installCmd).FlagCompletion(carapace.ActionMap{
		"build-constraint":  carapace.ActionFiles(),
		"build-constraints": carapace.ActionFiles(),
		"constraint":        carapace.ActionFiles(),
		"constraints":       carapace.ActionFiles(),
		"exclude":           carapace.ActionFiles(),
		"excludes":          carapace.ActionFiles(),
		"extra":             uv.ActionExtras(),
		"fork-strategy":     uv.ActionForkStrategies(),
		"group":             uv.ActionDependencyGroups(),
		"index-strategy":    uv.ActionIndexStrategies(),
		"keyring-provider":  uv.ActionKeyringProviders(),
		"link-mode":         uv.ActionLinkModes(),
		"output-format": carapace.ActionValuesDescribed(
			"text", "Display the result in a human-readable format",
			"json", "Display the result in JSON format",
		),
		"override":        carapace.ActionFiles(),
		"overrides":       carapace.ActionFiles(),
		"prefix":          carapace.ActionDirectories(),
		"prerelease":      uv.ActionPreReleases(),
		"python":          uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"python-platform": uv.ActionPythonPlatforms(),
		"requirement":     carapace.ActionFiles(),
		"requirements":    carapace.ActionFiles(),
		"resolution":      uv.ActionResolutions(),
		"target":          carapace.ActionDirectories(),
		"torch-backend":   uv.ActionTorchBackends(),
	})

	carapace.Gen(pip_installCmd).PositionalAnyCompletion(pip.ActionPackageSearch())
}
