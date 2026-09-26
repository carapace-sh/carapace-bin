package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var pip_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information about one or more installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pip_showCmd).Standalone()

	pip_showCmd.Flags().Bool("disable-pip-version-check", false, "")
	pip_showCmd.Flags().BoolP("files", "f", false, "Show the full list of installed files for each package")
	pip_showCmd.Flags().Bool("no-strict", false, "")
	pip_showCmd.Flags().Bool("no-system", false, "")
	pip_showCmd.Flags().String("prefix", "", "Show a package from the specified `--prefix` directory")
	pip_showCmd.Flags().StringP("python", "p", "", "The Python interpreter to find the package in.")
	pip_showCmd.Flags().Bool("strict", false, "Validate the Python environment, to detect packages with missing dependencies and other issues")
	pip_showCmd.Flags().Bool("system", false, "Show a package in the system Python environment")
	pip_showCmd.Flags().StringP("target", "t", "", "Show a package from the specified `--target` directory")
	pip_showCmd.Flag("disable-pip-version-check").Hidden = true
	pip_showCmd.Flag("no-strict").Hidden = true
	pip_showCmd.Flag("no-system").Hidden = true
	pipCmd.AddCommand(pip_showCmd)
	carapace.Gen(pip_showCmd).FlagCompletion(carapace.ActionMap{
		"prefix": carapace.ActionDirectories(),
		"python": uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"target": carapace.ActionDirectories(),
	})
	carapace.Gen(pip_showCmd).PositionalAnyCompletion(uv.ActionInstalledPackages())
}
