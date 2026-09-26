package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var pip_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall packages from an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pip_uninstallCmd).Standalone()

	pip_uninstallCmd.Flags().Bool("break-system-packages", false, "Allow uv to modify an `EXTERNALLY-MANAGED` Python installation")
	pip_uninstallCmd.Flags().Bool("disable-pip-version-check", false, "")
	pip_uninstallCmd.Flags().Bool("dry-run", false, "Perform a dry run, i.e., don't actually uninstall anything but print the resulting plan")
	pip_uninstallCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for remote requirements files")
	pip_uninstallCmd.Flags().Bool("no-break-system-packages", false, "")
	pip_uninstallCmd.Flags().Bool("no-system", false, "")
	pip_uninstallCmd.Flags().String("prefix", "", "Uninstall packages from the specified `--prefix` directory")
	pip_uninstallCmd.Flags().StringP("python", "p", "", "The Python interpreter from which packages should be uninstalled.")
	pip_uninstallCmd.Flags().StringSlice("requirement", nil, "Uninstall the packages listed in the given files")
	pip_uninstallCmd.Flags().StringSliceP("requirements", "r", nil, "Uninstall the packages listed in the given files")
	pip_uninstallCmd.Flags().Bool("system", false, "Use the system Python to uninstall packages")
	pip_uninstallCmd.Flags().StringP("target", "t", "", "Uninstall packages from the specified `--target` directory")
	pip_uninstallCmd.Flags().BoolP("yes", "y", false, "Don't ask for confirmation of uninstall deletions")
	pip_uninstallCmd.Flag("disable-pip-version-check").Hidden = true
	pip_uninstallCmd.Flag("no-system").Hidden = true
	pip_uninstallCmd.Flag("requirement").Hidden = true
	pip_uninstallCmd.Flag("yes").Hidden = true
	pipCmd.AddCommand(pip_uninstallCmd)
	carapace.Gen(pip_uninstallCmd).FlagCompletion(carapace.ActionMap{
		"keyring-provider": uv.ActionKeyringProviders(),
		"prefix":           carapace.ActionDirectories(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"requirement":      carapace.ActionFiles(),
		"requirements":     carapace.ActionFiles(),
		"target":           carapace.ActionDirectories(),
	})
	carapace.Gen(pip_uninstallCmd).PositionalAnyCompletion(uv.ActionInstalledPackages())
}
