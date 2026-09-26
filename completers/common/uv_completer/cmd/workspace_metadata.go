package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var workspace_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "View metadata about the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_metadataCmd).Standalone()

	workspace_metadataCmd.Flags().Bool("active", false, "Sync dependencies to the active virtual environment")
	workspace_metadataCmd.Flags().Bool("binary", false, "")
	workspace_metadataCmd.Flags().Bool("build", false, "")
	workspace_metadataCmd.Flags().Bool("build-isolation", false, "")
	workspace_metadataCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	workspace_metadataCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	workspace_metadataCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	workspace_metadataCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	workspace_metadataCmd.Flags().Bool("exact", false, "Perform an exact sync, removing extraneous packages")
	workspace_metadataCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	workspace_metadataCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	workspace_metadataCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	workspace_metadataCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	workspace_metadataCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	workspace_metadataCmd.Flags().Bool("frozen", false, "Assert that a `uv.lock` exists without checking if it is up-to-date [env: UV_FROZEN=]")
	workspace_metadataCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	workspace_metadataCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	workspace_metadataCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	workspace_metadataCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	workspace_metadataCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	workspace_metadataCmd.Flags().Bool("locked", false, "Check if the lockfile is up-to-date [env: UV_LOCKED=]")
	workspace_metadataCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	workspace_metadataCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	workspace_metadataCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	workspace_metadataCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	workspace_metadataCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	workspace_metadataCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	workspace_metadataCmd.Flags().Bool("no-frozen", false, "Disable frozen mode, overriding `UV_FROZEN`")
	workspace_metadataCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	workspace_metadataCmd.Flags().Bool("no-locked", false, "Disable locked mode, overriding `UV_LOCKED`")
	workspace_metadataCmd.Flags().Bool("no-refresh", false, "")
	workspace_metadataCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	workspace_metadataCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	workspace_metadataCmd.Flags().Bool("no-upgrade", false, "")
	workspace_metadataCmd.Flags().Bool("pre", false, "")
	workspace_metadataCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	workspace_metadataCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	workspace_metadataCmd.Flags().StringP("python", "p", "", "The Python interpreter to use during resolution.")
	workspace_metadataCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	workspace_metadataCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	workspace_metadataCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	workspace_metadataCmd.Flags().String("script", "", "View metadata for the specified PEP 723 Python script, rather than the current workspace")
	workspace_metadataCmd.Flags().Bool("sync", false, "Sync the environment to include module ownership metadata in the output")
	workspace_metadataCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	workspace_metadataCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	workspace_metadataCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	workspace_metadataCmd.Flag("binary").Hidden = true
	workspace_metadataCmd.Flag("build").Hidden = true
	workspace_metadataCmd.Flag("build-isolation").Hidden = true
	workspace_metadataCmd.Flag("config-settings").Hidden = true
	workspace_metadataCmd.Flag("config-settings-package").Hidden = true
	workspace_metadataCmd.Flag("no-frozen").Hidden = true
	workspace_metadataCmd.Flag("no-locked").Hidden = true
	workspace_metadataCmd.Flag("no-refresh").Hidden = true
	workspace_metadataCmd.Flag("no-upgrade").Hidden = true
	workspace_metadataCmd.Flag("pre").Hidden = true
	workspaceCmd.AddCommand(workspace_metadataCmd)
	carapace.Gen(workspace_metadataCmd).FlagCompletion(carapace.ActionMap{
		"fork-strategy":    uv.ActionForkStrategies(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"prerelease":       uv.ActionPreReleases(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"resolution":       uv.ActionResolutions(),
		"script":           carapace.ActionFiles(),
	})
}
