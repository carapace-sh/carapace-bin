package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var python_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Download and install Python versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(python_installCmd).Standalone()

	python_installCmd.Flags().Bool("bin", false, "Install a Python executable into the `bin` directory")
	python_installCmd.Flags().Bool("compile", false, "Compile Python's standard library to bytecode after installation")
	python_installCmd.Flags().Bool("compile-bytecode", false, "Compile Python's standard library to bytecode after installation")
	python_installCmd.Flags().Bool("default", false, "Use as the default Python version")
	python_installCmd.Flags().BoolP("force", "f", false, "Replace existing Python executables during installation")
	python_installCmd.Flags().StringP("install-dir", "i", "", "The directory to store the Python installation in")
	python_installCmd.Flags().String("mirror", "", "Set the URL to use as the source for downloading Python installations")
	python_installCmd.Flags().Bool("no-bin", false, "Do not install a Python executable into the `bin` directory")
	python_installCmd.Flags().Bool("no-compile", false, "")
	python_installCmd.Flags().Bool("no-compile-bytecode", false, "")
	python_installCmd.Flags().Bool("no-registry", false, "Do not register the Python installation in the Windows registry")
	python_installCmd.Flags().String("pypy-mirror", "", "Set the URL to use as the source for downloading PyPy installations")
	python_installCmd.Flags().String("python-downloads-json-url", "", "URL pointing to JSON of custom Python installations")
	python_installCmd.Flags().Bool("registry", false, "Register the Python installation in the Windows registry")
	python_installCmd.Flags().BoolP("reinstall", "r", false, "Reinstall the requested Python version, if it's already installed")
	python_installCmd.Flags().BoolP("upgrade", "U", false, "Upgrade existing Python installations to the latest patch version")
	python_installCmd.Flag("bin").Hidden = true
	python_installCmd.Flag("compile").Hidden = true
	python_installCmd.Flag("no-compile").Hidden = true
	python_installCmd.Flag("no-compile-bytecode").Hidden = true
	python_installCmd.Flag("registry").Hidden = true
	pythonCmd.AddCommand(python_installCmd)
	carapace.Gen(python_installCmd).FlagCompletion(carapace.ActionMap{
		"install-dir": carapace.ActionDirectories(),
	})
	carapace.Gen(python_installCmd).PositionalAnyCompletion(uv.ActionPythonInstallations(uv.InstallationsOpts{}))
}
