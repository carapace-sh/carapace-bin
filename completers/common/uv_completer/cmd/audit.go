package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit the project's dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditCmd).Standalone()

	auditCmd.Flags().Bool("binary", false, "")
	auditCmd.Flags().Bool("build", false, "")
	auditCmd.Flags().Bool("build-isolation", false, "")
	auditCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	auditCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	auditCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	auditCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	auditCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	auditCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	auditCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	auditCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	auditCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	auditCmd.Flags().Bool("frozen", false, "Audit the requirements without locking the project [env: UV_FROZEN=]")
	auditCmd.Flags().StringSlice("ignore", nil, "Ignore a vulnerability by ID")
	auditCmd.Flags().StringSlice("ignore-until-fixed", nil, "Ignore a vulnerability by ID, but only while no fix is available")
	auditCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	auditCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	auditCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	auditCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	auditCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	auditCmd.Flags().Bool("locked", false, "Assert that the `uv.lock` will remain unchanged [env: UV_LOCKED=]")
	auditCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	auditCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	auditCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	auditCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	auditCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	auditCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	auditCmd.Flags().Bool("no-default-groups", false, "Don't audit the default dependency groups")
	auditCmd.Flags().Bool("no-dev", false, "Don't audit the development dependency group [env: UV_NO_DEV=]")
	auditCmd.Flags().StringSlice("no-extra", nil, "Don't audit the specified optional dependencies")
	auditCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	auditCmd.Flags().StringSlice("no-group", nil, "Don't audit the specified dependency group [env: `UV_NO_GROUP`=]")
	auditCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	auditCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	auditCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	auditCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	auditCmd.Flags().Bool("no-upgrade", false, "")
	auditCmd.Flags().Bool("only-dev", false, "Only audit the development dependency group")
	auditCmd.Flags().StringSlice("only-group", nil, "Only audit dependencies from the specified dependency group")
	auditCmd.Flags().String("output-format", "text", "Select the output format")
	auditCmd.Flags().Bool("pre", false, "")
	auditCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	auditCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	auditCmd.Flags().String("python-platform", "", "The platform to use when auditing")
	auditCmd.Flags().String("python-version", "", "The Python version to use when auditing")
	auditCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	auditCmd.Flags().String("script", "", "Audit the specified PEP 723 Python script, rather than the current project")
	auditCmd.Flags().String("service-format", "osv", "The service format to use for vulnerability lookups")
	auditCmd.Flags().String("service-url", "", "The URL to vulnerability service API endpoint")
	auditCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	auditCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	auditCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	auditCmd.Flag("binary").Hidden = true
	auditCmd.Flag("build").Hidden = true
	auditCmd.Flag("build-isolation").Hidden = true
	auditCmd.Flag("config-settings").Hidden = true
	auditCmd.Flag("config-settings-package").Hidden = true
	auditCmd.Flag("no-frozen").Hidden = true
	auditCmd.Flag("no-locked").Hidden = true
	auditCmd.Flag("no-upgrade").Hidden = true
	auditCmd.Flag("pre").Hidden = true
	rootCmd.AddCommand(auditCmd)
	carapace.Gen(auditCmd).FlagCompletion(carapace.ActionMap{
		"fork-strategy":    uv.ActionForkStrategies(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"output-format": carapace.ActionValuesDescribed(
			"text", "Display the result in a human-readable format",
			"json", "Display the result in JSON format",
			"sarif", "Display the result in SARIF format",
		),
		"prerelease":      uv.ActionPreReleases(),
		"python-platform": uv.ActionPythonPlatforms(),
		"resolution":      uv.ActionResolutions(),
		"script":          carapace.ActionFiles(),
		"service-format":  carapace.ActionValues("osv"),
	})
}
