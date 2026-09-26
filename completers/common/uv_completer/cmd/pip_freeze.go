package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var pip_freezeCmd = &cobra.Command{
	Use:   "freeze",
	Short: "List, in requirements format, packages installed in an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pip_freezeCmd).Standalone()

	pip_freezeCmd.Flags().Bool("disable-pip-version-check", false, "")
	pip_freezeCmd.Flags().StringSlice("exclude", nil, "Exclude the specified package(s) from the output")
	pip_freezeCmd.Flags().Bool("exclude-editable", false, "Exclude any editable packages from output")
	pip_freezeCmd.Flags().Bool("no-strict", false, "")
	pip_freezeCmd.Flags().Bool("no-system", false, "")
	pip_freezeCmd.Flags().StringSlice("path", nil, "Restrict to the specified installation path for listing packages (can be used multiple times)")
	pip_freezeCmd.Flags().String("prefix", "", "List packages from the specified `--prefix` directory")
	pip_freezeCmd.Flags().StringP("python", "p", "", "The Python interpreter for which packages should be listed.")
	pip_freezeCmd.Flags().Bool("strict", false, "Validate the Python environment, to detect packages with missing dependencies and other issues")
	pip_freezeCmd.Flags().Bool("system", false, "List packages in the system Python environment")
	pip_freezeCmd.Flags().StringP("target", "t", "", "List packages from the specified `--target` directory")
	pip_freezeCmd.Flag("disable-pip-version-check").Hidden = true
	pip_freezeCmd.Flag("no-strict").Hidden = true
	pip_freezeCmd.Flag("no-system").Hidden = true
	pipCmd.AddCommand(pip_freezeCmd)
	carapace.Gen(pip_freezeCmd).FlagCompletion(carapace.ActionMap{
		"path":   carapace.ActionDirectories(),
		"prefix": carapace.ActionDirectories(),
		"python": uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"target": carapace.ActionDirectories(),
	})
}
