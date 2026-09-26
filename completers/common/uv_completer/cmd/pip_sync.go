package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var pip_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync an environment with a `requirements.txt` or `pylock.toml` file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pip_syncCmd).Standalone()

	pip_syncCmd.Flags().Bool("all-extras", false, "Include all optional dependencies")
	pip_syncCmd.Flags().Bool("allow-empty-requirements", false, "Allow sync of empty requirements, which will clear the environment of all packages")
	pip_syncCmd.Flags().BoolP("ask", "a", false, "")
	pip_syncCmd.Flags().Bool("break-system-packages", false, "Allow uv to modify an `EXTERNALLY-MANAGED` Python installation")
	pip_syncCmd.Flags().Bool("build", false, "")
	pip_syncCmd.Flags().StringSlice("build-constraint", nil, "Constrain build dependencies using the given requirements files when building source distributions")
	pip_syncCmd.Flags().StringSliceP("build-constraints", "b", nil, "Constrain build dependencies using the given requirements files when building source distributions")
	pip_syncCmd.Flags().Bool("build-isolation", false, "")
	pip_syncCmd.Flags().Bool("check", false, "Check whether the environment matches the requirements without modifying it")
	pip_syncCmd.Flags().String("client-cert", "", "")
	pip_syncCmd.Flags().Bool("compile", false, "Compile Python files to bytecode after installation")
	pip_syncCmd.Flags().Bool("compile-bytecode", false, "Compile Python files to bytecode after installation")
	pip_syncCmd.Flags().String("config", "", "")
	pip_syncCmd.Flags().StringSliceP("config-setting", "C", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	pip_syncCmd.Flags().StringSlice("config-settings", nil, "Settings to pass to the PEP 517 build backend, specified as `KEY=VALUE` pairs")
	pip_syncCmd.Flags().StringSlice("config-settings-package", nil, "Settings to pass to the PEP 517 build backend for a specific package, specified as `PACKAGE:KEY=VALUE` pairs")
	pip_syncCmd.Flags().StringSlice("constraint", nil, "Constrain versions using the given requirements files")
	pip_syncCmd.Flags().StringSliceP("constraints", "c", nil, "Constrain versions using the given requirements files")
	pip_syncCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	pip_syncCmd.Flags().Bool("dry-run", false, "Perform a dry run, i.e., don't actually install anything but resolve the dependencies and print the resulting plan")
	pip_syncCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	pip_syncCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	pip_syncCmd.Flags().StringSlice("extra", nil, "Include optional dependencies from the specified extra name; may be provided more than once")
	pip_syncCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	pip_syncCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	pip_syncCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	pip_syncCmd.Flags().StringSlice("group", nil, "Install the specified dependency group from a `pylock.toml` or `pyproject.toml`")
	pip_syncCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	pip_syncCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	pip_syncCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	pip_syncCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	pip_syncCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	pip_syncCmd.Flags().Bool("no-all-extras", false, "")
	pip_syncCmd.Flags().Bool("no-allow-empty-requirements", false, "")
	pip_syncCmd.Flags().StringSlice("no-binary", nil, "Don't install pre-built wheels")
	pip_syncCmd.Flags().Bool("no-break-system-packages", false, "")
	pip_syncCmd.Flags().Bool("no-build", false, "Don't build source distributions")
	pip_syncCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building source distributions")
	pip_syncCmd.Flags().Bool("no-compile", false, "")
	pip_syncCmd.Flags().Bool("no-compile-bytecode", false, "")
	pip_syncCmd.Flags().Bool("no-config", false, "")
	pip_syncCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	pip_syncCmd.Flags().Bool("no-refresh", false, "")
	pip_syncCmd.Flags().Bool("no-reinstall", false, "")
	pip_syncCmd.Flags().Bool("no-require-hashes", false, "")
	pip_syncCmd.Flags().Bool("no-sources", false, "Ignore the `tool.uv.sources` table when resolving dependencies. Used to lock against the standards-compliant, publishable package metadata, as opposed to using any workspace, Git, URL, or local path sources")
	pip_syncCmd.Flags().StringSlice("no-sources-package", nil, "Don't use sources from the `tool.uv.sources` table for the specified packages [env: `UV_NO_SOURCES_PACKAGE`=]")
	pip_syncCmd.Flags().Bool("no-strict", false, "")
	pip_syncCmd.Flags().Bool("no-system", false, "")
	pip_syncCmd.Flags().Bool("no-verify-hashes", false, "Disable validation of hashes in the requirements file")
	pip_syncCmd.Flags().StringSlice("only-binary", nil, "Only use pre-built wheels; don't build source distributions")
	pip_syncCmd.Flags().String("output-format", "text", "Select the output format")
	pip_syncCmd.Flags().String("pip-args", "", "")
	pip_syncCmd.Flags().String("prefix", "", "Install packages into `lib`, `bin`, and other top-level folders under the specified directory, as if a virtual environment were present at that location")
	pip_syncCmd.Flags().StringP("python", "p", "", "The Python interpreter into which packages should be installed.")
	pip_syncCmd.Flags().String("python-executable", "", "")
	pip_syncCmd.Flags().String("python-platform", "", "The platform for which requirements should be installed")
	pip_syncCmd.Flags().String("python-version", "", "The minimum Python version that should be supported by the requirements (e.g., `3.7` or `3.7.9`)")
	pip_syncCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	pip_syncCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	pip_syncCmd.Flags().Bool("reinstall", false, "Reinstall all packages, regardless of whether they're already installed. Implies `--refresh`")
	pip_syncCmd.Flags().StringSlice("reinstall-package", nil, "Reinstall a specific package, regardless of whether it's already installed. Implies `--refresh-package`")
	pip_syncCmd.Flags().Bool("require-hashes", false, "Require a matching hash for each requirement")
	pip_syncCmd.Flags().Bool("strict", false, "Validate the Python environment after completing the installation, to detect packages with missing dependencies or other issues")
	pip_syncCmd.Flags().Bool("system", false, "Install packages into the system Python environment")
	pip_syncCmd.Flags().StringP("target", "t", "", "Install packages into the specified directory, rather than into the virtual or system Python environment. The packages will be installed at the top-level of the directory")
	pip_syncCmd.Flags().String("torch-backend", "", "The backend to use when fetching packages in the PyTorch ecosystem (e.g., `cpu`, `cu126`, or `auto`)")
	pip_syncCmd.Flags().Bool("user", false, "")
	pip_syncCmd.Flags().Bool("verify-hashes", false, "")
	pip_syncCmd.Flag("ask").Hidden = true
	pip_syncCmd.Flag("build").Hidden = true
	pip_syncCmd.Flag("build-constraint").Hidden = true
	pip_syncCmd.Flag("build-isolation").Hidden = true
	pip_syncCmd.Flag("client-cert").Hidden = true
	pip_syncCmd.Flag("compile").Hidden = true
	pip_syncCmd.Flag("config").Hidden = true
	pip_syncCmd.Flag("config-settings").Hidden = true
	pip_syncCmd.Flag("constraint").Hidden = true
	pip_syncCmd.Flag("force-reinstall").Hidden = true
	pip_syncCmd.Flag("no-all-extras").Hidden = true
	pip_syncCmd.Flag("no-compile").Hidden = true
	pip_syncCmd.Flag("no-compile-bytecode").Hidden = true
	pip_syncCmd.Flag("no-config").Hidden = true
	pip_syncCmd.Flag("no-refresh").Hidden = true
	pip_syncCmd.Flag("no-reinstall").Hidden = true
	pip_syncCmd.Flag("no-require-hashes").Hidden = true
	pip_syncCmd.Flag("no-strict").Hidden = true
	pip_syncCmd.Flag("no-system").Hidden = true
	pip_syncCmd.Flag("pip-args").Hidden = true
	pip_syncCmd.Flag("python-executable").Hidden = true
	pip_syncCmd.Flag("user").Hidden = true
	pip_syncCmd.Flag("verify-hashes").Hidden = true
	pipCmd.AddCommand(pip_syncCmd)
	carapace.Gen(pip_syncCmd).FlagCompletion(carapace.ActionMap{
		"build-constraint":  carapace.ActionFiles(),
		"build-constraints": carapace.ActionFiles(),
		"constraint":        carapace.ActionFiles(),
		"constraints":       carapace.ActionFiles(),
		"extra":             uv.ActionExtras(),
		"group":             uv.ActionDependencyGroups(),
		"index-strategy":    uv.ActionIndexStrategies(),
		"keyring-provider":  uv.ActionKeyringProviders(),
		"link-mode":         uv.ActionLinkModes(),
		"output-format": carapace.ActionValuesDescribed(
			"text", "Display the result in a human-readable format",
			"json", "Display the result in JSON format",
		),
		"prefix":          carapace.ActionDirectories(),
		"python":          uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"python-platform": uv.ActionPythonPlatforms(),
		"target":          carapace.ActionDirectories(),
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
	})
	carapace.Gen(pip_syncCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
