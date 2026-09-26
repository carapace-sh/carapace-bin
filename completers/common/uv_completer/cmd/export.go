package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the project's lockfile to an alternate format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()

	exportCmd.Flags().Bool("all-extras", false, "Include all optional dependencies")
	exportCmd.Flags().Bool("all-groups", false, "Include dependencies from all dependency groups")
	exportCmd.Flags().Bool("all-packages", false, "Export the entire workspace")
	exportCmd.Flags().Bool("annotate", false, "")
	exportCmd.Flags().String("batch", "", "Export multiple selections from a TOML manifest containing `[[export]]` entries")
	exportCmd.Flags().Bool("binary", false, "")
	exportCmd.Flags().Bool("build", false, "")
	exportCmd.Flags().Bool("build-isolation", false, "")
	exportCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	exportCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	exportCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	exportCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	exportCmd.Flags().Bool("dev", false, "Include the development dependency group [env: UV_DEV=]")
	exportCmd.Flags().Bool("editable", false, "Export any non-editable dependencies, including the project and any workspace members, as editable")
	exportCmd.Flags().Bool("emit-find-links", false, "Include `--find-links` entries in the generated output file")
	exportCmd.Flags().Bool("emit-index-url", false, "Include `--index-url` and `--extra-index-url` entries in the generated output file")
	exportCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	exportCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	exportCmd.Flags().StringSlice("extra", nil, "Include optional dependencies from the specified extra name")
	exportCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	exportCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	exportCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	exportCmd.Flags().String("format", "", "The format to which `uv.lock` should be exported")
	exportCmd.Flags().Bool("frozen", false, "Do not update the `uv.lock` before exporting [env: UV_FROZEN=]")
	exportCmd.Flags().StringSlice("group", nil, "Include dependencies from the specified dependency group")
	exportCmd.Flags().Bool("hashes", false, "Include hashes for all dependencies")
	exportCmd.Flags().Bool("header", false, "")
	exportCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	exportCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	exportCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	exportCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	exportCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	exportCmd.Flags().Bool("locked", false, "Assert that the `uv.lock` will remain unchanged [env: UV_LOCKED=]")
	exportCmd.Flags().Bool("no-all-extras", false, "")
	exportCmd.Flags().Bool("no-annotate", false, "Exclude comment annotations indicating the source of each package")
	exportCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	exportCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	exportCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	exportCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	exportCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	exportCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	exportCmd.Flags().Bool("no-default-groups", false, "Ignore the default dependency groups")
	exportCmd.Flags().Bool("no-dev", false, "Disable the development dependency group [env: UV_NO_DEV=]")
	exportCmd.Flags().Bool("no-editable", false, "Export any editable dependencies, including the project and any workspace members, as non-editable [env: UV_NO_EDITABLE=]")
	exportCmd.Flags().StringSlice("no-editable-package", nil, "Export the specified editable packages as non-editable")
	exportCmd.Flags().Bool("no-emit-find-links", false, "")
	exportCmd.Flags().Bool("no-emit-index-url", false, "")
	exportCmd.Flags().Bool("no-emit-local", false, "Do not include local path dependencies in the exported requirements")
	exportCmd.Flags().StringSlice("no-emit-package", nil, "Do not emit the given package(s)")
	exportCmd.Flags().Bool("no-emit-project", false, "Do not emit the current project")
	exportCmd.Flags().Bool("no-emit-workspace", false, "Do not emit any workspace members, including the root project")
	exportCmd.Flags().StringSlice("no-extra", nil, "Exclude the specified optional dependencies, if `--all-extras` is supplied")
	exportCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	exportCmd.Flags().StringSlice("no-group", nil, "Disable the specified dependency group [env: `UV_NO_GROUP`=]")
	exportCmd.Flags().Bool("no-hashes", false, "Omit hashes in the generated output")
	exportCmd.Flags().Bool("no-header", false, "Exclude the comment header at the top of the generated output file")
	exportCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	exportCmd.Flags().Bool("no-install-local", false, "Do not include local path dependencies in the exported requirements")
	exportCmd.Flags().StringSlice("no-install-package", nil, "Do not emit the given package(s)")
	exportCmd.Flags().Bool("no-install-project", false, "Do not emit the current project")
	exportCmd.Flags().Bool("no-install-workspace", false, "Do not emit any workspace members, including the root project")
	exportCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	exportCmd.Flags().Bool("no-refresh", false, "")
	exportCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	exportCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	exportCmd.Flags().Bool("no-upgrade", false, "")
	exportCmd.Flags().Bool("only-dev", false, "Only include the development dependency group")
	exportCmd.Flags().Bool("only-emit-local", false, "Only include local path dependencies in the exported requirements")
	exportCmd.Flags().StringSlice("only-emit-package", nil, "Only emit the given package(s)")
	exportCmd.Flags().Bool("only-emit-project", false, "Only emit the current project")
	exportCmd.Flags().Bool("only-emit-workspace", false, "Only emit workspace members, including the root project")
	exportCmd.Flags().StringSlice("only-group", nil, "Only include dependencies from the specified dependency group")
	exportCmd.Flags().Bool("only-install-local", false, "Only include local path dependencies in the exported requirements")
	exportCmd.Flags().StringSlice("only-install-package", nil, "Only emit the given package(s)")
	exportCmd.Flags().Bool("only-install-project", false, "Only emit the current project")
	exportCmd.Flags().Bool("only-install-workspace", false, "Only emit workspace members, including the root project")
	exportCmd.Flags().StringP("output-file", "o", "", "Write the exported requirements to the given file")
	exportCmd.Flags().StringSlice("package", nil, "Export the dependencies for specific packages in the workspace")
	exportCmd.Flags().Bool("pre", false, "")
	exportCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	exportCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	exportCmd.Flags().StringSlice("prune", nil, "Prune the given package from the dependency tree")
	exportCmd.Flags().StringP("python", "p", "", "The Python interpreter to use during resolution.")
	exportCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	exportCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	exportCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	exportCmd.Flags().String("script", "", "Export the dependencies for the specified PEP 723 Python script, rather than the current project")
	exportCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	exportCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	exportCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	exportCmd.Flag("annotate").Hidden = true
	exportCmd.Flag("batch").Hidden = true
	exportCmd.Flag("binary").Hidden = true
	exportCmd.Flag("build").Hidden = true
	exportCmd.Flag("build-isolation").Hidden = true
	exportCmd.Flag("config-settings").Hidden = true
	exportCmd.Flag("dev").Hidden = true
	exportCmd.Flag("editable").Hidden = true
	exportCmd.Flag("hashes").Hidden = true
	exportCmd.Flag("header").Hidden = true
	exportCmd.Flag("no-all-extras").Hidden = true
	exportCmd.Flag("no-emit-find-links").Hidden = true
	exportCmd.Flag("no-emit-index-url").Hidden = true
	exportCmd.Flag("no-frozen").Hidden = true
	exportCmd.Flag("no-install-local").Hidden = true
	exportCmd.Flag("no-install-package").Hidden = true
	exportCmd.Flag("no-install-project").Hidden = true
	exportCmd.Flag("no-install-workspace").Hidden = true
	exportCmd.Flag("no-locked").Hidden = true
	exportCmd.Flag("no-refresh").Hidden = true
	exportCmd.Flag("no-upgrade").Hidden = true
	exportCmd.Flag("only-emit-local").Hidden = true
	exportCmd.Flag("only-emit-package").Hidden = true
	exportCmd.Flag("only-emit-project").Hidden = true
	exportCmd.Flag("only-emit-workspace").Hidden = true
	exportCmd.Flag("only-install-local").Hidden = true
	exportCmd.Flag("only-install-package").Hidden = true
	exportCmd.Flag("only-install-project").Hidden = true
	exportCmd.Flag("only-install-workspace").Hidden = true
	exportCmd.Flag("pre").Hidden = true
	rootCmd.AddCommand(exportCmd)
	carapace.Gen(exportCmd).FlagCompletion(carapace.ActionMap{
		"batch":         carapace.ActionFiles(),
		"extra":         uv.ActionExtras(),
		"fork-strategy": uv.ActionForkStrategies(),
		"format": carapace.ActionValuesDescribed(
			"requirements.txt", "Export in `requirements.txt` format",
			"pylock.toml", "Export in `pylock.toml` format",
			"cyclonedx1.5", "Export in `CycloneDX` v1.5 JSON format",
		),
		"group":            uv.ActionDependencyGroups(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"output-file":      carapace.ActionFiles(),
		"prerelease":       uv.ActionPreReleases(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"resolution":       uv.ActionResolutions(),
		"script":           carapace.ActionFiles(),
	})
}
