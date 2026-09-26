package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pip"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var tool_uvxCmd = &cobra.Command{
	Use:    "uvx",
	Short:  "Run a command provided by a Python package.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tool_uvxCmd).Standalone()

	tool_uvxCmd.Flags().Bool("binary", false, "")
	tool_uvxCmd.Flags().Bool("build", false, "")
	tool_uvxCmd.Flags().StringSlice("build-constraint", nil, "Constrain build dependencies using the given requirements files when building source distributions")
	tool_uvxCmd.Flags().StringSliceP("build-constraints", "b", nil, "Constrain build dependencies using the given requirements files when building source distributions")
	tool_uvxCmd.Flags().Bool("build-isolation", false, "")
	tool_uvxCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	tool_uvxCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	tool_uvxCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	tool_uvxCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	tool_uvxCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	tool_uvxCmd.Flags().StringSlice("constraint", nil, "Constrain versions using the given requirements files")
	tool_uvxCmd.Flags().StringSliceP("constraints", "c", nil, "Constrain versions using the given requirements files")
	tool_uvxCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	tool_uvxCmd.Flags().StringSlice("env-file", nil, "Load environment variables from a `.env` file")
	tool_uvxCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	tool_uvxCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	tool_uvxCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	tool_uvxCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	tool_uvxCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	tool_uvxCmd.Flags().String("fork-strategy", "", "The strategy to use when selecting multiple versions of a given package across Python versions and platforms")
	tool_uvxCmd.Flags().String("from", "", "Use the given package to provide the command")
	tool_uvxCmd.Flags().String("generate-shell-completion", "", "")
	tool_uvxCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	tool_uvxCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	tool_uvxCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	tool_uvxCmd.Flags().Bool("isolated", false, "Run the tool in an isolated virtual environment, ignoring any already-installed tools [env: UV_ISOLATED=]")
	tool_uvxCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	tool_uvxCmd.Flags().Bool("lfs", false, "Whether to use Git LFS when adding a dependency from Git")
	tool_uvxCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	tool_uvxCmd.Flags().Bool("no-binary", false, "Don't install pre-built wheels")
	tool_uvxCmd.Flags().StringSlice("no-binary-package", nil, "Don't install pre-built wheels for a specific package [env: `UV_NO_BINARY_PACKAGE`=]")
	tool_uvxCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	tool_uvxCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	tool_uvxCmd.Flags().StringSlice("no-build-isolation-package", nil, "Disable isolation when building source distributions for a specific package")
	tool_uvxCmd.Flags().StringSlice("no-build-package", nil, "Don't build source distributions for a specific package [env: `UV_NO_BUILD_PACKAGE`=]")
	tool_uvxCmd.Flags().Bool("no-compile", false, "")
	tool_uvxCmd.Flags().Bool("no-compile-bytecode", false, "")
	tool_uvxCmd.Flags().Bool("no-env-file", false, "Avoid reading environment variables from a `.env` file [env: UV_NO_ENV_FILE=]")
	tool_uvxCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	tool_uvxCmd.Flags().Bool("no-refresh", false, "")
	tool_uvxCmd.Flags().Bool("no-reinstall", false, "")
	tool_uvxCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	tool_uvxCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	tool_uvxCmd.Flags().Bool("no-upgrade", false, "")
	tool_uvxCmd.Flags().StringSlice("override", nil, "Override versions using the given requirements files")
	tool_uvxCmd.Flags().StringSlice("overrides", nil, "Override versions using the given requirements files")
	tool_uvxCmd.Flags().Bool("pre", false, "")
	tool_uvxCmd.Flags().String("prerelease", "", "The strategy to use when considering pre-release versions")
	tool_uvxCmd.Flags().StringSlice("prerelease-package", nil, "The strategy to use when considering pre-release versions for a specific package")
	tool_uvxCmd.Flags().StringP("python", "p", "", "The Python interpreter to use to build the run environment.")
	tool_uvxCmd.Flags().String("python-platform", "", "The platform for which requirements should be installed")
	tool_uvxCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	tool_uvxCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	tool_uvxCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	tool_uvxCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	tool_uvxCmd.Flags().String("resolution", "", "The strategy to use when selecting between the different compatible versions for a given package requirement")
	tool_uvxCmd.Flags().Bool("show-resolution", false, "Whether to show resolver and installer output from any environment modifications [env: UV_SHOW_RESOLUTION=]")
	tool_uvxCmd.Flags().String("torch-backend", "", "The backend to use when fetching packages in the PyTorch ecosystem (e.g., `cpu`, `cu126`, or `auto`)")
	tool_uvxCmd.Flags().BoolP("upgrade", "U", false, "Allow package upgrades, ignoring pinned versions in any existing output file. Implies `--refresh`")
	tool_uvxCmd.Flags().StringSlice("upgrade-group", nil, "Allow upgrades for all packages in a dependency group, ignoring pinned versions in any existing output file")
	tool_uvxCmd.Flags().StringSliceP("upgrade-package", "P", nil, "Allow upgrades for a specific package, ignoring pinned versions in any existing output file. Implies `--refresh-package`")
	tool_uvxCmd.Flags().BoolP("version", "V", false, "Display the uvx version")
	tool_uvxCmd.Flags().StringSliceP("with", "w", nil, "Run with the given packages installed")
	tool_uvxCmd.Flags().StringSlice("with-editable", nil, "Run with the given packages installed in editable mode")
	tool_uvxCmd.Flags().StringSlice("with-requirements", nil, "Run with the packages listed in the given files")
	tool_uvxCmd.Flag("binary").Hidden = true
	tool_uvxCmd.Flag("build").Hidden = true
	tool_uvxCmd.Flag("build-constraint").Hidden = true
	tool_uvxCmd.Flag("build-isolation").Hidden = true
	tool_uvxCmd.Flag("compile").Hidden = true
	tool_uvxCmd.Flag("config-settings").Hidden = true
	tool_uvxCmd.Flag("constraint").Hidden = true
	tool_uvxCmd.Flag("force-reinstall").Hidden = true
	tool_uvxCmd.Flag("generate-shell-completion").Hidden = true
	tool_uvxCmd.Flag("no-compile").Hidden = true
	tool_uvxCmd.Flag("no-compile-bytecode").Hidden = true
	tool_uvxCmd.Flag("no-refresh").Hidden = true
	tool_uvxCmd.Flag("no-reinstall").Hidden = true
	tool_uvxCmd.Flag("no-upgrade").Hidden = true
	tool_uvxCmd.Flag("override").Hidden = true
	tool_uvxCmd.Flag("pre").Hidden = true
	tool_uvxCmd.Flag("show-resolution").Hidden = true
	toolCmd.AddCommand(tool_uvxCmd)
	carapace.Gen(tool_uvxCmd).FlagCompletion(carapace.ActionMap{
		"build-constraint":  carapace.ActionFiles(),
		"build-constraints": carapace.ActionFiles(),
		"constraint":        carapace.ActionFiles(),
		"constraints":       carapace.ActionFiles(),
		"env-file":          carapace.ActionFiles(),
		"fork-strategy":     uv.ActionForkStrategies(),
		"from":              pip.ActionPackageSearch(),
		"generate-shell-completion": carapace.ActionValues(
			"bash",
			"elvish",
			"fish",
			"nushell",
			"powershell",
			"zsh",
		),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"override":         carapace.ActionFiles(),
		"overrides":        carapace.ActionFiles(),
		"prerelease":       uv.ActionPreReleases(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"python-platform":  uv.ActionPythonPlatforms(),
		"resolution":       uv.ActionResolutions(),
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
	carapace.Gen(tool_uvxCmd).PositionalCompletion(carapace.Batch(uv.ActionTools(), pip.ActionPackageSearch()).ToA())
}
