package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var tool_upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade installed tools",
	Aliases: []string{"update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tool_upgradeCmd).Standalone()

	tool_upgradeCmd.Flags().Bool("all", false, "Upgrade all tools")
	tool_upgradeCmd.Flags().Bool("binary", false, "")
	tool_upgradeCmd.Flags().Bool("build", false, "")
	tool_upgradeCmd.Flags().Bool("build-isolation", false, "")
	tool_upgradeCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	tool_upgradeCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	tool_upgradeCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	tool_upgradeCmd.Flags().StringSlice("config-setting-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	tool_upgradeCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	tool_upgradeCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	tool_upgradeCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	tool_upgradeCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	tool_upgradeCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	tool_upgradeCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	tool_upgradeCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	tool_upgradeCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	tool_upgradeCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	tool_upgradeCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	tool_upgradeCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	tool_upgradeCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	tool_upgradeCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	tool_upgradeCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	tool_upgradeCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	tool_upgradeCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	tool_upgradeCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	tool_upgradeCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	tool_upgradeCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	tool_upgradeCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	tool_upgradeCmd.Flags().Bool("no-compile", false, "")
	tool_upgradeCmd.Flags().Bool("no-compile-bytecode", false, "")
	tool_upgradeCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	tool_upgradeCmd.Flags().Bool("no-reinstall", false, "")
	tool_upgradeCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	tool_upgradeCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	tool_upgradeCmd.Flags().Bool("pre", false, "")
	tool_upgradeCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	tool_upgradeCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	tool_upgradeCmd.Flags().StringP("python", "p", "", "Upgrade a tool, and specify it to use the given Python interpreter to build its environment.")
	tool_upgradeCmd.Flags().String("python-platform", "", "The platform for which requirements should be installed")
	tool_upgradeCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	tool_upgradeCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	tool_upgradeCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	tool_upgradeCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	tool_upgradeCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	tool_upgradeCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	tool_upgradeCmd.Flag("binary").Hidden = true
	tool_upgradeCmd.Flag("build").Hidden = true
	tool_upgradeCmd.Flag("build-isolation").Hidden = true
	tool_upgradeCmd.Flag("compile").Hidden = true
	tool_upgradeCmd.Flag("config-settings").Hidden = true
	tool_upgradeCmd.Flag("config-settings-package").Hidden = true
	tool_upgradeCmd.Flag("force-reinstall").Hidden = true
	tool_upgradeCmd.Flag("no-compile").Hidden = true
	tool_upgradeCmd.Flag("no-compile-bytecode").Hidden = true
	tool_upgradeCmd.Flag("no-reinstall").Hidden = true
	tool_upgradeCmd.Flag("pre").Hidden = true
	tool_upgradeCmd.Flag("upgrade").Hidden = true
	tool_upgradeCmd.Flag("upgrade-group").Hidden = true
	tool_upgradeCmd.Flag("upgrade-package").Hidden = true
	toolCmd.AddCommand(tool_upgradeCmd)
	carapace.Gen(tool_upgradeCmd).FlagCompletion(carapace.ActionMap{
		"fork-strategy":    uv.ActionForkStrategies(),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"prerelease":       uv.ActionPreReleases(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"python-platform":  uv.ActionPythonPlatforms(),
		"resolution":       uv.ActionResolutions(),
	})
	carapace.Gen(tool_upgradeCmd).PositionalAnyCompletion(uv.ActionTools())
}
