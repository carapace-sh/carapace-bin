package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Python packages into source distributions and wheels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("all", false, "Builds all packages in the workspace")
	buildCmd.Flags().Bool("all-packages", false, "Builds all packages in the workspace")
	buildCmd.Flags().Bool("binary", false, "")
	buildCmd.Flags().Bool("build", false, "")
	buildCmd.Flags().StringSlice("build-constraint", nil, "Constrain build dependencies using the given requirements files when building distributions")
	buildCmd.Flags().StringSliceP("build-constraints", "b", nil, "Constrain build dependencies using the given requirements files when building distributions")
	buildCmd.Flags().Bool("build-isolation", false, "")
	buildCmd.Flags().Bool("build-logs", false, "")
	buildCmd.Flags().Bool("clear", false, "Clear the output directory before the build, removing stale artifacts")
	buildCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	buildCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	buildCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	buildCmd.Flags().Bool("create-gitignore", false, "")
	buildCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	buildCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	buildCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	buildCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	buildCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	buildCmd.Flags().Bool("force-pep517", false, "Always build through PEP 517, don't use the fast path for the uv build backend")
	buildCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	buildCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	buildCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	buildCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	buildCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	buildCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	buildCmd.Flags().Bool("list", false, "When using the uv build backend, list the files that would be included when building")
	buildCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	buildCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	buildCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	buildCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	buildCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	buildCmd.Flags().Bool("no-build-logs", false, "Hide logs from the build backend")
	buildCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	buildCmd.Flags().Bool("no-create-gitignore", false, "Do not create a `.gitignore` file in the output directory")
	buildCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	buildCmd.Flags().Bool("no-refresh", false, "")
	buildCmd.Flags().Bool("no-require-hashes", false, "")
	buildCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	buildCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	buildCmd.Flags().Bool("no-upgrade", false, "")
	buildCmd.Flags().Bool("no-verify-hashes", false, "Disable validation of hashes in the requirements file")
	buildCmd.Flags().StringP("out-dir", "o", "", "The output directory to which distributions should be written")
	buildCmd.Flags().String("package", "", "Build a specific package in the workspace")
	buildCmd.Flags().Bool("pre", false, "")
	buildCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	buildCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	buildCmd.Flags().StringP("python", "p", "", "The Python interpreter to use for the build environment.")
	buildCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	buildCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	buildCmd.Flags().Bool("require-hashes", false, "Require a matching hash for each requirement")
	buildCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	buildCmd.Flags().Bool("sdist", false, "Build a source distribution (\"sdist\") from the given directory")
	buildCmd.Flags().Bool("skip-dependency-check", false, "Skip checking if build dependencies are satisfied when building without isolation")
	buildCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	buildCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	buildCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	buildCmd.Flags().Bool("verify-hashes", false, "")
	buildCmd.Flags().Bool("wheel", false, "Build a binary distribution (\"wheel\") from the given directory")
	buildCmd.Flag("all").Hidden = true
	buildCmd.Flag("binary").Hidden = true
	buildCmd.Flag("build").Hidden = true
	buildCmd.Flag("build-constraint").Hidden = true
	buildCmd.Flag("build-isolation").Hidden = true
	buildCmd.Flag("build-logs").Hidden = true
	buildCmd.Flag("config-settings").Hidden = true
	buildCmd.Flag("create-gitignore").Hidden = true
	buildCmd.Flag("list").Hidden = true
	buildCmd.Flag("no-refresh").Hidden = true
	buildCmd.Flag("no-require-hashes").Hidden = true
	buildCmd.Flag("no-upgrade").Hidden = true
	buildCmd.Flag("pre").Hidden = true
	buildCmd.Flag("skip-dependency-check").Hidden = true
	buildCmd.Flag("verify-hashes").Hidden = true
	rootCmd.AddCommand(buildCmd)
	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"build-constraint":  carapace.ActionFiles(),
		"build-constraints": carapace.ActionFiles(),
		"fork-strategy":     uv.ActionForkStrategies(),
		"index-strategy":    uv.ActionIndexStrategies(),
		"keyring-provider":  uv.ActionKeyringProviders(),
		"link-mode":         uv.ActionLinkModes(),
		"out-dir":           carapace.ActionDirectories(),
		"prerelease":        uv.ActionPreReleases(),
		"python":            uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"resolution":        uv.ActionResolutions(),
	})
	carapace.Gen(buildCmd).PositionalCompletion(carapace.ActionDirectories())
}
