package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pip"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var tool_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install commands provided by a Python package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tool_installCmd).Standalone()

	tool_installCmd.Flags().Bool("binary", false, "")
	tool_installCmd.Flags().Bool("build", false, "")
	tool_installCmd.Flags().StringSlice("build-constraint", nil, "Constrain build dependencies using the given requirements files when building source distributions")
	tool_installCmd.Flags().StringSliceP("build-constraints", "b", nil, "Constrain build dependencies using the given requirements files when building source distributions")
	tool_installCmd.Flags().Bool("build-isolation", false, "")
	tool_installCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	tool_installCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	tool_installCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	tool_installCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	tool_installCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	tool_installCmd.Flags().StringSlice("constraint", nil, "Constrain versions using the given requirements files")
	tool_installCmd.Flags().StringSliceP("constraints", "c", nil, "Constrain versions using the given requirements files")
	tool_installCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	tool_installCmd.Flags().BoolP("editable", "e", false, "Install the target package in editable mode, such that changes in the package's source directory are reflected without reinstallation")
	tool_installCmd.Flags().StringSlice("exclude", nil, "Exclude packages from resolution using the given requirements files")
	tool_installCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	tool_installCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	tool_installCmd.Flags().StringSlice("excludes", nil, "Exclude packages from resolution using the given requirements files")
	tool_installCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	tool_installCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	tool_installCmd.Flags().Bool("force", false, "Force installation of the tool")
	tool_installCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	tool_installCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	tool_installCmd.Flags().String("from", "", "The package to install commands from")
	tool_installCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	tool_installCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	tool_installCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	tool_installCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	tool_installCmd.Flags().Bool("lfs", false, "Whether to use Git LFS when adding a dependency from Git")
	tool_installCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	tool_installCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	tool_installCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	tool_installCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	tool_installCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	tool_installCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	tool_installCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	tool_installCmd.Flags().Bool("no-compile", false, "")
	tool_installCmd.Flags().Bool("no-compile-bytecode", false, "")
	tool_installCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	tool_installCmd.Flags().Bool("no-refresh", false, "")
	tool_installCmd.Flags().Bool("no-reinstall", false, "")
	tool_installCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	tool_installCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	tool_installCmd.Flags().Bool("no-upgrade", false, "")
	tool_installCmd.Flags().StringSlice("override", nil, "Override versions using the given requirements files")
	tool_installCmd.Flags().StringSlice("overrides", nil, "Override versions using the given requirements files")
	tool_installCmd.Flags().Bool("pre", false, "")
	tool_installCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	tool_installCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	tool_installCmd.Flags().StringP("python", "p", "", "The Python interpreter to use to build the tool environment.")
	tool_installCmd.Flags().String("python-platform", "", "The platform for which requirements should be installed")
	tool_installCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	tool_installCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	tool_installCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	tool_installCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	tool_installCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	tool_installCmd.Flags().String("torch-backend", "", "The backend to use when fetching packages in the PyTorch ecosystem (e.g., `cpu`, `cu126`, or `auto`)")
	tool_installCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	tool_installCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	tool_installCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	tool_installCmd.Flags().StringSliceP("with", "w", nil, "Include the following additional requirements")
	tool_installCmd.Flags().StringSlice("with-editable", nil, "Include the given packages in editable mode")
	tool_installCmd.Flags().StringSlice("with-executables-from", nil, "Install executables from the following packages")
	tool_installCmd.Flags().StringSlice("with-requirements", nil, "Run with the packages listed in the given files")
	tool_installCmd.Flag("binary").Hidden = true
	tool_installCmd.Flag("build").Hidden = true
	tool_installCmd.Flag("build-constraint").Hidden = true
	tool_installCmd.Flag("build-isolation").Hidden = true
	tool_installCmd.Flag("compile").Hidden = true
	tool_installCmd.Flag("config-settings").Hidden = true
	tool_installCmd.Flag("config-settings-package").Hidden = true
	tool_installCmd.Flag("constraint").Hidden = true
	tool_installCmd.Flag("exclude").Hidden = true
	tool_installCmd.Flag("force-reinstall").Hidden = true
	tool_installCmd.Flag("from").Hidden = true
	tool_installCmd.Flag("no-compile").Hidden = true
	tool_installCmd.Flag("no-compile-bytecode").Hidden = true
	tool_installCmd.Flag("no-refresh").Hidden = true
	tool_installCmd.Flag("no-reinstall").Hidden = true
	tool_installCmd.Flag("no-upgrade").Hidden = true
	tool_installCmd.Flag("override").Hidden = true
	tool_installCmd.Flag("pre").Hidden = true
	toolCmd.AddCommand(tool_installCmd)
	carapace.Gen(tool_installCmd).FlagCompletion(carapace.ActionMap{
		"build-constraint":  carapace.ActionFiles(),
		"build-constraints": carapace.ActionFiles(),
		"constraint":        carapace.ActionFiles(),
		"constraints":       carapace.ActionFiles(),
		"exclude":           carapace.ActionFiles(),
		"excludes":          carapace.ActionFiles(),
		"fork-strategy":     uv.ActionForkStrategies(),
		"from":              pip.ActionPackageSearch(),
		"index-strategy":    uv.ActionIndexStrategies(),
		"keyring-provider":  uv.ActionKeyringProviders(),
		"link-mode":         uv.ActionLinkModes(),
		"override":          carapace.ActionFiles(),
		"overrides":         carapace.ActionFiles(),
		"prerelease":        uv.ActionPreReleases(),
		"python":            uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"python-platform":   uv.ActionPythonPlatforms(),
		"resolution":        uv.ActionResolutions(),
		"torch-backend": carapace.ActionValuesDescribed(
			"auto", "Select the appropriate PyTorch index based on the operating system and CUDA driver version",
			"cpu", "Use the CPU-only PyTorch index",
			"cu132", "Use the PyTorch index for CUDA 13.2",
			"cu130", "Use the PyTorch index for CUDA 13.0",
			"cu129", "Use the PyTorch index for CUDA 12.9",
			"cu128", "Use the PyTorch index for CUDA 12.8",
			"cu126", "Use the PyTorch index for CUDA 12.6",
			"cu125", "Use the PyTorch index for CUDA 12.5",
			"cu124", "Use the PyTorch index for CUDA 12.4",
			"cu123", "Use the PyTorch index for CUDA 12.3",
			"cu122", "Use the PyTorch index for CUDA 12.2",
			"cu121", "Use the PyTorch index for CUDA 12.1",
			"cu120", "Use the PyTorch index for CUDA 12.0",
			"cu118", "Use the PyTorch index for CUDA 11.8",
			"cu117", "Use the PyTorch index for CUDA 11.7",
			"cu116", "Use the PyTorch index for CUDA 11.6",
			"cu115", "Use the PyTorch index for CUDA 11.5",
			"cu114", "Use the PyTorch index for CUDA 11.4",
			"cu113", "Use the PyTorch index for CUDA 11.3",
			"cu112", "Use the PyTorch index for CUDA 11.2",
			"cu111", "Use the PyTorch index for CUDA 11.1",
			"cu110", "Use the PyTorch index for CUDA 11.0",
			"cu102", "Use the PyTorch index for CUDA 10.2",
			"cu101", "Use the PyTorch index for CUDA 10.1",
			"cu100", "Use the PyTorch index for CUDA 10.0",
			"cu92", "Use the PyTorch index for CUDA 9.2",
			"cu91", "Use the PyTorch index for CUDA 9.1",
			"cu90", "Use the PyTorch index for CUDA 9.0",
			"cu80", "Use the PyTorch index for CUDA 8.0",
			"rocm7.2", "Use the PyTorch index for ROCm 7.2",
			"rocm7.1", "Use the PyTorch index for ROCm 7.1",
			"rocm7.0", "Use the PyTorch index for ROCm 7.0",
			"rocm6.4", "Use the PyTorch index for ROCm 6.4",
			"rocm6.3", "Use the PyTorch index for ROCm 6.3",
			"rocm6.2.4", "Use the PyTorch index for ROCm 6.2.4",
			"rocm6.2", "Use the PyTorch index for ROCm 6.2",
			"rocm6.1", "Use the PyTorch index for ROCm 6.1",
			"rocm6.0", "Use the PyTorch index for ROCm 6.0",
			"rocm5.7", "Use the PyTorch index for ROCm 5.7",
			"rocm5.6", "Use the PyTorch index for ROCm 5.6",
			"rocm5.5", "Use the PyTorch index for ROCm 5.5",
			"rocm5.4.2", "Use the PyTorch index for ROCm 5.4.2",
			"rocm5.4", "Use the PyTorch index for ROCm 5.4",
			"rocm5.3", "Use the PyTorch index for ROCm 5.3",
			"rocm5.2", "Use the PyTorch index for ROCm 5.2",
			"rocm5.1.1", "Use the PyTorch index for ROCm 5.1.1",
			"rocm4.2", "Use the PyTorch index for ROCm 4.2",
			"rocm4.1", "Use the PyTorch index for ROCm 4.1",
			"rocm4.0.1", "Use the PyTorch index for ROCm 4.0.1",
			"xpu", "Use the PyTorch index for Intel XPU",
		),
		"with-editable":     carapace.ActionDirectories(),
		"with-requirements": carapace.ActionFiles(),
	})
	carapace.Gen(tool_installCmd).PositionalCompletion(pip.ActionPackageSearch())
}
