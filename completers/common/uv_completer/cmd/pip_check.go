package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var pip_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Verify installed packages have compatible dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pip_checkCmd).Standalone()

	pip_checkCmd.Flags().Bool("no-system", false, "")
	pip_checkCmd.Flags().StringP("python", "p", "", "The Python interpreter for which packages should be checked.")
	pip_checkCmd.Flags().String("python-platform", "", "The platform for which packages should be checked")
	pip_checkCmd.Flags().String("python-version", "", "The Python version against which packages should be checked")
	pip_checkCmd.Flags().Bool("system", false, "Check packages in the system Python environment")
	pip_checkCmd.Flag("no-system").Hidden = true
	pipCmd.AddCommand(pip_checkCmd)
	carapace.Gen(pip_checkCmd).FlagCompletion(carapace.ActionMap{
		"python":          uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"python-platform": uv.ActionPythonPlatforms(),
	})
}
